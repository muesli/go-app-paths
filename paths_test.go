// +build !darwin,!windows

package apppaths

import (
	"testing"
)

func TestDirs(t *testing.T) {
	tests := []struct {
		scopeType  ScopeType
		app        string
		dataDir    string
		cacheDir   string
		configFile string
		logFile    string
	}{
		{System, "foobar", "/usr/share/foobar", "/var/cache/foobar", "/etc/foobar/foobar.conf", "/var/log/foobar/foobar.log"},
		{User, "foobar", "~/.local/share/foobar", "~/.cache/foobar", "~/.config/foobar/foobar.conf", "~/.local/share/foobar/foobar.log"},
	}
	for _, tt := range tests {
		s := NewScope(tt.scopeType, "", tt.app)

		path, err := s.DataDir()
		if err != nil {
			t.Errorf("Error retrieving data dir: %s", err)
		}
		if path != expandUser(tt.dataDir) {
			t.Errorf("Expected data dir: %s - got: %s", tt.dataDir, path)
		}

		path, err = s.CacheDir()
		if err != nil {
			t.Errorf("Error retrieving cache dir: %s", err)
		}
		if path != expandUser(tt.cacheDir) {
			t.Errorf("Expected cache dir: %s - got: %s", tt.cacheDir, path)
		}

		path, err = s.ConfigPath(tt.app + ".conf")
		if err != nil {
			t.Errorf("Error retrieving config path: %s", err)
		}
		if path != expandUser(tt.configFile) {
			t.Errorf("Expected config path: %s - got: %s", tt.configFile, path)
		}

		path, err = s.LogPath(tt.app + ".log")
		if err != nil {
			t.Errorf("Error retrieving log path: %s", err)
		}
		if path != expandUser(tt.logFile) {
			t.Errorf("Expected log path: %s - got: %s", tt.logFile, path)
		}
	}
}
