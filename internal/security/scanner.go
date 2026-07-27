// Package security implements Phase 07 (Shield) structural security intelligence.
// It scans source files for hardcoded secrets, SQL injection risks, XSS patterns,
// and unsafe dependency patterns without requiring external vulnerability databases.
package security

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// SeverityLevel classifies the criticality of a security finding.
type SeverityLevel string

const (
	Critical SeverityLevel = "CRITICAL"
	High     SeverityLevel = "HIGH"
	Medium   SeverityLevel = "MEDIUM"
	Low      SeverityLevel = "LOW"
	Info     SeverityLevel = "INFO"
)

// Finding represents a single security issue discovered during scanning.
type Finding struct {
	ID          string        `json:"id"`
	Category    string        `json:"category"`
	Severity    SeverityLevel `json:"severity"`
	FilePath    string        `json:"file_path"`
	LineNumber  int           `json:"line_number"`
	LineContent string        `json:"line_content"`
	Description string        `json:"description"`
	Suggestion  string        `json:"suggestion"`
}

// SecurityReport is the aggregate output of a Shield security scan.
type SecurityReport struct {
	ScannedAt    string    `json:"scanned_at"`
	TotalFiles   int       `json:"total_files_scanned"`
	TotalFindings int      `json:"total_findings"`
	Critical     int       `json:"critical_count"`
	High         int       `json:"high_count"`
	Medium       int       `json:"medium_count"`
	Low          int       `json:"low_count"`
	SecurityGrade string   `json:"security_grade"`
	Findings     []Finding `json:"findings"`
}

// secretPatterns defines regex patterns for detecting hardcoded secrets.
var secretPatterns = []struct {
	Name     string
	Pattern  *regexp.Regexp
	Severity SeverityLevel
}{
	{"AWS Access Key", regexp.MustCompile(`(?i)(AKIA[0-9A-Z]{16})`), Critical},
	{"AWS Secret Key", regexp.MustCompile(`(?i)aws_secret_access_key\s*[=:]\s*["']?([A-Za-z0-9/+=]{40})["']?`), Critical},
	{"Generic API Key", regexp.MustCompile(`(?i)(api[_-]?key|apikey)\s*[=:]\s*["']([A-Za-z0-9_\-]{20,})["']`), High},
	{"Generic Secret", regexp.MustCompile(`(?i)(secret|password|passwd|pwd|token)\s*[=:]\s*["']([^"']{8,})["']`), High},
	{"Private Key Block", regexp.MustCompile(`-----BEGIN (RSA |EC |DSA )?PRIVATE KEY-----`), Critical},
	{"JWT Token", regexp.MustCompile(`eyJ[A-Za-z0-9_-]{10,}\.[A-Za-z0-9_-]{10,}\.[A-Za-z0-9_-]{10,}`), High},
	{"GitHub Token", regexp.MustCompile(`(?i)(ghp_[A-Za-z0-9]{36}|github_pat_[A-Za-z0-9_]{22,})`), Critical},
	{"Slack Token", regexp.MustCompile(`xox[bpors]-[0-9a-zA-Z]{10,}`), High},
	{"Database URL", regexp.MustCompile(`(?i)(mysql|postgres|mongodb|redis)://[^\s"']+:[^\s"']+@`), Critical},
	{"Hardcoded IP", regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}:\d{2,5}\b`), Low},
}

// injectionPatterns detects SQL injection and XSS risks in source code.
var injectionPatterns = []struct {
	Name     string
	Pattern  *regexp.Regexp
	Severity SeverityLevel
	Category string
}{
	{"SQL String Concatenation", regexp.MustCompile(`(?i)(SELECT|INSERT|UPDATE|DELETE|DROP)\s+.*\+\s*(req\.|request\.|params\.|args\.|input)`), High, "SQL_INJECTION"},
	{"SQL Format String", regexp.MustCompile(`(?i)fmt\.Sprintf\(\s*["'].*?(SELECT|INSERT|UPDATE|DELETE|WHERE)`), Medium, "SQL_INJECTION"},
	{"Unsafe innerHTML", regexp.MustCompile(`(?i)\.innerHTML\s*=\s*[^'"]`), Medium, "XSS"},
	{"document.write", regexp.MustCompile(`(?i)document\.write\s*\(`), Medium, "XSS"},
	{"eval() Usage", regexp.MustCompile(`(?i)\beval\s*\(`), High, "CODE_INJECTION"},
	{"exec() Usage", regexp.MustCompile(`(?i)\bexec\s*\(\s*["']`), High, "CODE_INJECTION"},
	{"Unvalidated Redirect", regexp.MustCompile(`(?i)(redirect|location\.href)\s*=\s*(req\.|request\.|params\.)`), Medium, "OPEN_REDIRECT"},
	{"Disabled CSRF", regexp.MustCompile(`(?i)(csrf|xsrf).*disable|no.?csrf`), Medium, "CSRF"},
	{"Weak Crypto", regexp.MustCompile(`(?i)(md5|sha1)\s*[\.(]`), Low, "WEAK_CRYPTO"},
}

// skipExtensions are file extensions that should not be scanned for security issues.
var skipExtensions = map[string]bool{
	".png": true, ".jpg": true, ".jpeg": true, ".gif": true, ".ico": true,
	".svg": true, ".woff": true, ".woff2": true, ".ttf": true, ".eot": true,
	".mp3": true, ".mp4": true, ".avi": true, ".mov": true, ".pdf": true,
	".zip": true, ".tar": true, ".gz": true, ".db": true, ".sqlite": true,
	".exe": true, ".dll": true, ".so": true, ".dylib": true, ".wasm": true,
}

// skipDirs are directories that should be excluded from scanning.
var skipDirs = map[string]bool{
	".git": true, ".codemri": true, "node_modules": true, "vendor": true,
	"dist": true, "build": true, "__pycache__": true, ".venv": true,
	"bin": true, ".idea": true, ".vscode": true, "coverage": true,
}

// RunSecurityScan performs a full Shield security scan on the given directory.
func RunSecurityScan(rootDir string) (*SecurityReport, error) {
	report := &SecurityReport{
		ScannedAt: time.Now().Format(time.RFC3339),
		Findings:  make([]Finding, 0),
	}

	findingID := 0
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		// Skip directories
		if info.IsDir() {
			if skipDirs[info.Name()] {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip binary/media files
		ext := strings.ToLower(filepath.Ext(path))
		if skipExtensions[ext] {
			return nil
		}

		// Skip large files (>1MB likely not source code)
		if info.Size() > 1024*1024 {
			return nil
		}

		report.TotalFiles++

		file, err := os.Open(path)
		if err != nil {
			return nil
		}
		defer file.Close()

		relPath, _ := filepath.Rel(rootDir, path)
		relPath = filepath.ToSlash(relPath)

		scanner := bufio.NewScanner(file)
		lineNum := 0
		for scanner.Scan() {
			lineNum++
			line := scanner.Text()
			trimmed := strings.TrimSpace(line)

			// Skip comments and empty lines for performance
			if trimmed == "" || strings.HasPrefix(trimmed, "//") && !strings.Contains(strings.ToLower(trimmed), "password") {
				continue
			}

			// Check secret patterns
			for _, sp := range secretPatterns {
				if sp.Pattern.MatchString(line) {
					// Skip test files and example files for common patterns
					if sp.Severity == Low && (strings.Contains(relPath, "test") || strings.Contains(relPath, "example")) {
						continue
					}
					findingID++
					report.Findings = append(report.Findings, Finding{
						ID:          fmt.Sprintf("SHIELD-%04d", findingID),
						Category:    "HARDCODED_SECRET",
						Severity:    sp.Severity,
						FilePath:    relPath,
						LineNumber:  lineNum,
						LineContent: truncateLine(line, 120),
						Description: fmt.Sprintf("Potential %s detected in source code", sp.Name),
						Suggestion:  "Move secrets to environment variables or a secure vault. Never commit credentials to version control.",
					})
				}
			}

			// Check injection patterns
			for _, ip := range injectionPatterns {
				if ip.Pattern.MatchString(line) {
					findingID++
					report.Findings = append(report.Findings, Finding{
						ID:          fmt.Sprintf("SHIELD-%04d", findingID),
						Category:    ip.Category,
						Severity:    ip.Severity,
						FilePath:    relPath,
						LineNumber:  lineNum,
						LineContent: truncateLine(line, 120),
						Description: fmt.Sprintf("Potential %s risk: %s", ip.Category, ip.Name),
						Suggestion:  getSuggestion(ip.Category),
					})
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Calculate counts
	for _, f := range report.Findings {
		switch f.Severity {
		case Critical:
			report.Critical++
		case High:
			report.High++
		case Medium:
			report.Medium++
		case Low:
			report.Low++
		}
	}
	report.TotalFindings = len(report.Findings)
	report.SecurityGrade = calculateSecurityGrade(report)

	return report, nil
}

// SaveReport persists the security report as JSON to the .codemri directory.
func SaveReport(report *SecurityReport, outputPath string) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(outputPath, data, 0644)
}

func calculateSecurityGrade(r *SecurityReport) string {
	if r.Critical > 0 {
		return "F — Critical Vulnerabilities Detected"
	}
	if r.High > 3 {
		return "D — Multiple High-Severity Issues"
	}
	if r.High > 0 {
		return "C — High-Severity Issues Present"
	}
	if r.Medium > 5 {
		return "B- — Moderate Security Concerns"
	}
	if r.Medium > 0 {
		return "B — Minor Security Improvements Recommended"
	}
	if r.Low > 0 {
		return "A- — Low-Risk Items Only"
	}
	return "A+ — Excellent Security Posture"
}

func getSuggestion(category string) string {
	switch category {
	case "SQL_INJECTION":
		return "Use parameterized queries or prepared statements. Never concatenate user input into SQL strings."
	case "XSS":
		return "Sanitize all user input before rendering in HTML. Use framework-provided escaping functions."
	case "CODE_INJECTION":
		return "Avoid eval() and exec() with dynamic input. Use safe alternatives like JSON.parse() or structured dispatch."
	case "OPEN_REDIRECT":
		return "Validate redirect URLs against an allowlist. Never redirect to user-supplied URLs without validation."
	case "CSRF":
		return "Enable CSRF protection tokens on all state-changing endpoints."
	case "WEAK_CRYPTO":
		return "Replace MD5/SHA1 with SHA-256 or bcrypt for security-sensitive hashing."
	default:
		return "Review this code for potential security implications."
	}
}

func truncateLine(line string, maxLen int) string {
	trimmed := strings.TrimSpace(line)
	if len(trimmed) > maxLen {
		return trimmed[:maxLen] + "..."
	}
	return trimmed
}
