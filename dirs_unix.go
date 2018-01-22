// +build !darwin,!windows

package dirs

import (
	"os"
)

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

// configPath returns the full path to the config dir.
func (s *Scope) configPath() (string, error) {
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

// logPath returns the full path to the log dir.
func (s *Scope) logPath() (string, error) {
	switch s.Type {
	case System:
		return "/var/log", nil
	case User:
		return s.dataDir()
	}

	return "", ErrInvalidScope
}
