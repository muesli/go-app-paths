// +build !darwin,!windows

package apppaths

import (
	"testing"
)

func TestPaths(t *testing.T) {
	tests := []struct {
		scope      *Scope
		dataDir    string
		cacheDir   string
		configFile string
		logFile    string
	}{
		{NewScope(System, "", "foobar"), "/usr/share/foobar", "/var/cache/foobar", "/etc/foobar/foobar.conf", "/var/log/foobar/foobar.log"},
		{NewScope(User, "", "foobar"), "~/.local/share/foobar", "~/.cache/foobar", "~/.config/foobar/foobar.conf", "~/.local/share/foobar/foobar.log"},
		{NewCustomHomeScope("/tmp", "", "foobar"), "/tmp/.local/share/foobar", "/tmp/.cache/foobar", "/tmp/.config/foobar/foobar.conf", "/tmp/.local/share/foobar/foobar.log"},
	}
	for _, tt := range tests {
		path, err := tt.scope.DataDir()
		if err != nil {
			t.Errorf("Error retrieving data dir: %s", err)
		}
		if path != expandUser(tt.dataDir) {
			t.Errorf("Expected data dir: %s - got: %s", tt.dataDir, path)
		}

		path, err = tt.scope.CacheDir()
		if err != nil {
			t.Errorf("Error retrieving cache dir: %s", err)
		}
		if path != expandUser(tt.cacheDir) {
			t.Errorf("Expected cache dir: %s - got: %s", tt.cacheDir, path)
		}

		path, err = tt.scope.ConfigPath(tt.scope.App + ".conf")
		if err != nil {
			t.Errorf("Error retrieving config path: %s", err)
		}
		if path != expandUser(tt.configFile) {
			t.Errorf("Expected config path: %s - got: %s", tt.configFile, path)
		}

		path, err = tt.scope.LogPath(tt.scope.App + ".log")
		if err != nil {
			t.Errorf("Error retrieving log path: %s", err)
		}
		if path != expandUser(tt.logFile) {
			t.Errorf("Expected log path: %s - got: %s", tt.logFile, path)
		}
	}
}
