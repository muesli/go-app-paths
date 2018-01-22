package apppaths

import (
	"errors"
	"os/user"
	"path/filepath"
	"strings"
)

var (
	// ErrInvalidScope gets returned when an invalid scope type has been set.
	ErrInvalidScope = errors.New("Invalid scope type")
	// ErrRetrievingPath gets returned when the path could not be resolved.
	ErrRetrievingPath = errors.New("Could not retrieve path")
)

// ScopeType specifies whether returned paths are user-specific or system-wide.
type ScopeType int

const (
	// System is the system-wide scope.
	System ScopeType = iota
	// User is the user-specific scope.
	User
)

// Scope holds scope & app-specific information.
type Scope struct {
	Type   ScopeType
	Vendor string
	App    string
}

// NewScope returns a new Scope that lets you query app- & platform-specific
// paths.
func NewScope(t ScopeType, vendor, app string) *Scope {
	return &Scope{
		Type:   t,
		Vendor: vendor,
		App:    app,
	}
}

// DataDir returns the full path to the application's data dir.
func (s *Scope) DataDir() (string, error) {
	p, err := s.dataDir()
	if err != nil {
		return p, err
	}

	return s.appendPaths(p), nil
}

// CacheDir returns the full path to the application's cache dir.
func (s *Scope) CacheDir() (string, error) {
	p, err := s.cacheDir()
	if err != nil {
		return p, err
	}

	return s.appendPaths(p), nil
}

// ConfigPath returns the full path to the application's config file.
func (s *Scope) ConfigPath(filename string) (string, error) {
	p, err := s.configPath()
	if err != nil {
		return p, err
	}

	return s.appendPaths(p, filename), nil
}

// LogPath returns the full path to the application's log file.
func (s *Scope) LogPath(filename string) (string, error) {
	p, err := s.logPath()
	if err != nil {
		return p, err
	}

	return s.appendPaths(p, filename), nil
}

// appendPaths appends the app-name and further variadic parts to a path
func (s *Scope) appendPaths(path string, parts ...string) string {
	paths := []string{path, s.App}
	paths = append(paths, parts...)
	return filepath.Join(paths...)
}

// expandUser is a helper function that expands the first '~' it finds in the
// passed path with the home directory of the current user.
func expandUser(path string) string {
	if u, err := user.Current(); err == nil {
		return strings.Replace(path, "~", u.HomeDir, -1)
	}
	return path
}
