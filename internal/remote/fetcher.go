// Package remote provides auto-cloning and caching capabilities for online GitHub/GitLab repository URLs.
package remote

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// IsRemoteURL checks if the provided target string represents an online Git repository URL.
func IsRemoteURL(target string) bool {
	t := strings.ToLower(strings.TrimSpace(target))
	return strings.HasPrefix(t, "http://") ||
		strings.HasPrefix(t, "https://") ||
		strings.HasPrefix(t, "git@") ||
		strings.HasPrefix(t, "github.com/") ||
		strings.HasPrefix(t, "gitlab.com/")
}

// NormalizeURL ensures that bare hostnames like "github.com/user/repo" have an https:// prefix.
func NormalizeURL(target string) string {
	target = strings.TrimSpace(target)
	if strings.HasPrefix(target, "github.com/") || strings.HasPrefix(target, "gitlab.com/") || strings.HasPrefix(target, "bitbucket.org/") {
		return "https://" + target
	}
	return target
}

// GetCacheDir returns the absolute path to ~/.codemri/cache directory where remote repositories are cloned.
func GetCacheDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		home = os.TempDir()
	}
	cacheDir := filepath.Join(home, ".codemri", "cache", "repos")
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create cache directory: %w", err)
	}
	return cacheDir, nil
}

// FetchRepository clones or updates an online Git repository into a local cache directory and returns the absolute local path.
// It uses shallow cloning (--depth=1) for ultra-fast performance without downloading full commit histories.
func FetchRepository(url string) (string, error) {
	url = NormalizeURL(url)

	// Extract meaningful folder name from URL (e.g., https://github.com/torvalds/linux.git -> torvalds_linux)
	parts := strings.Split(strings.TrimSuffix(url, ".git"), "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid repository URL format: %s", url)
	}
	repoName := parts[len(parts)-1]
	ownerName := parts[len(parts)-2]
	dirName := fmt.Sprintf("%s_%s", ownerName, repoName)

	cacheDir, err := GetCacheDir()
	if err != nil {
		return "", err
	}
	targetPath := filepath.Join(cacheDir, dirName)

	// Check if repository already exists in cache
	if _, err := os.Stat(filepath.Join(targetPath, ".git")); err == nil {
		fmt.Printf("📦 Using cached repository at: %s\n", targetPath)
		// Optionally attempt git pull to sync newest changes
		cmd := exec.Command("git", "pull", "--depth", "1", "origin", "main")
		cmd.Dir = targetPath
		_ = cmd.Run() // Ignore errors if offline or branch differs
		return targetPath, nil
	}

	fmt.Printf("🌐 Remote Git repository detected! Performing high-speed shallow clone of %s...\n", url)
	cmd := exec.Command("git", "clone", "--depth", "1", url, targetPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		_ = os.RemoveAll(targetPath) // Clean up partial download on failure
		return "", fmt.Errorf("failed to clone repository from %s: %w", url, err)
	}

	fmt.Printf("✅ Repository cloned successfully to cache: %s\n", targetPath)
	return targetPath, nil
}

// ResolveTarget takes any user input (local directory or online URL) and resolves it to a valid local directory path.
func ResolveTarget(target string) (string, error) {
	if IsRemoteURL(target) {
		return FetchRepository(target)
	}
	// Local directory path
	abs, err := filepath.Abs(target)
	if err != nil {
		return "", fmt.Errorf("invalid path %s: %w", target, err)
	}
	if info, err := os.Stat(abs); err != nil || !info.IsDir() {
		return "", fmt.Errorf("directory does not exist: %s", abs)
	}
	return abs, nil
}
