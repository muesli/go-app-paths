// +build !darwin,!windows

package apppaths

import (
	"os"
	"path/filepath"
)

// appendPaths appends the app-name and further variadic parts to a path
func (s *Scope) appendPaths(path string, parts ...string) string {
	paths := []string{path, s.App}
	paths = append(paths, parts...)
	return filepath.Join(paths...)
}

// dataDir returns the full path to the data directory.
func (s *Scope) dataDir() (string, error) {
	switch s.Type {
	case System:
		return "/usr/share", nil
	case User:
		path := os.Getenv("XDG_DATA_HOME")
		if path == "" {
			return expandUser("~/.local/share"), nil
		}
		return path, nil
	}

	return "", ErrInvalidScope
}

// cacheDir returns the full path to the cache directory.
func (s *Scope) cacheDir() (string, error) {
	switch s.Type {
	case System:
		return "/var/cache", nil
	case User:
		path := os.Getenv("XDG_CACHE_HOME")
		if path == "" {
			return expandUser("~/.cache"), nil
		}
		return path, nil
	}

	return "", ErrInvalidScope
}

// configDir returns the full path to the config dir.
func (s *Scope) configDir() (string, error) {
	switch s.Type {
	case System:
		return "/etc", nil
	case User:
		path := os.Getenv("XDG_CONFIG_HOME")
		if path == "" {
			return expandUser("~/.config"), nil
		}
		return path, nil
	}

	return "", ErrInvalidScope
}

// logDir returns the full path to the log dir.
func (s *Scope) logDir() (string, error) {
	switch s.Type {
	case System:
		return "/var/log", nil
	case User:
		return s.dataDir()
	}

	return "", ErrInvalidScope
}
