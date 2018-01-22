// +build darwin

package dirs

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

// configPath returns the full path to the config dir.
func (s *Scope) configPath() (string, error) {
	def := "/Library/Preferences"
	switch s.Type {
	case System:
		return def, nil
	case User:
		return expandUser("~" + def), nil
	}

	return "", ErrInvalidScope
}

// logPath returns the full path to the log dir.
func (s *Scope) logPath() (string, error) {
	def := "/Library/Logs"
	switch s.Type {
	case System:
		return def, nil
	case User:
		return expandUser("~" + def), nil
	}

	return "", ErrInvalidScope
}
