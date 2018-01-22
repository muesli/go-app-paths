// +build darwin

package apppaths

import "path/filepath"

// appendPaths appends the app-name and further variadic parts to a path
func (s *Scope) appendPaths(path string, parts ...string) string {
	paths := []string{path, s.App}
	paths = append(paths, parts...)
	return filepath.Join(paths...)
}

// dataDir returns the full path to the data directory.
func (s *Scope) dataDir() (string, error) {
	def := "/Library/Application Support"
	switch s.Type {
	case System:
		return def, nil
	case User:
		return expandUser("~" + def), nil
	}

	return "", ErrInvalidScope
}

// cacheDir returns the full path to the cache directory.
func (s *Scope) cacheDir() (string, error) {
	def := "/Library/Caches"
	switch s.Type {
	case System:
		return def, nil
	case User:
		return expandUser("~" + def), nil
	}

	return "", ErrInvalidScope
}

// configDir returns the full path to the config dir.
func (s *Scope) configDir() (string, error) {
	def := "/Library/Preferences"
	switch s.Type {
	case System:
		return def, nil
	case User:
		return expandUser("~" + def), nil
	}

	return "", ErrInvalidScope
}

// logDir returns the full path to the log dir.
func (s *Scope) logDir() (string, error) {
	def := "/Library/Logs"
	switch s.Type {
	case System:
		return def, nil
	case User:
		return expandUser("~" + def), nil
	}

	return "", ErrInvalidScope
}
