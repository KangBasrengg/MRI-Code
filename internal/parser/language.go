package parser

import (
	"path/filepath"
	"strings"
)

// DetectLanguage determines the primary programming or markup language of a file by extension and path.
func DetectLanguage(filePath string) Language {
	ext := strings.ToLower(filepath.Ext(filePath))
	base := strings.ToLower(filepath.Base(filePath))

	// Special filenames
	if base == "dockerfile" || strings.HasPrefix(base, "dockerfile.") {
		return "docker"
	}
	if base == "makefile" || base == "gnumakefile" {
		return "make"
	}

	switch ext {
	case ".go":
		return LangGo
	case ".ts", ".tsx":
		return LangTypeScript
	case ".js", ".jsx", ".mjs", ".cjs":
		return LangJavaScript
	case ".py", ".pyw":
		return LangPython
	case ".java":
		return LangJava
	case ".php", ".phtml":
		return LangPHP
	case ".sql":
		return LangSQL
	case ".rs":
		return LangRust
	case ".html", ".htm":
		return LangHTML
	case ".css", ".scss", ".sass", ".less":
		return LangCSS
	case ".md", ".markdown":
		return LangMarkdown
	case ".json":
		return LangJSON
	case ".yml", ".yaml":
		return LangYAML
	case ".sh", ".bash", ".zsh":
		return LangShell
	default:
		return LangUnknown
	}
}

// IsSourceCode Returns true if the language represents executable program logic or architecture schema.
func IsSourceCode(lang Language) bool {
	switch lang {
	case LangGo, LangTypeScript, LangJavaScript, LangPython, LangJava, LangPHP, LangSQL, LangRust:
		return true
	default:
		return false
	}
}
