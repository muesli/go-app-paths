// +build darwin

package gap

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
		{NewScope(System, "foobar"), "/Library/Application Support/foobar", "/Library/Caches/foobar", "/Library/Preferences/foobar/foobar.conf", "/Library/Logs/foobar/foobar.log"},
		{NewScope(User, "foobar"), "~/Library/Application Support/foobar", "~/Library/Caches/foobar", "~/Library/Preferences/foobar/foobar.conf", "~/Library/Logs/foobar/foobar.log"},
		{NewCustomHomeScope("/tmp", "", "foobar"), "/tmp/Library/Application Support/foobar", "/tmp/Library/Caches/foobar", "/tmp/Library/Preferences/foobar/foobar.conf", "/tmp/Library/Logs/foobar/foobar.log"},
	}
	for _, tt := range tests {
		paths, err := tt.scope.DataDirs()
		if err != nil {
			t.Errorf("Error retrieving data dir: %s", err)
		}
		if paths[0] != expandUser(tt.dataDir) {
			t.Errorf("Expected data dir: %s - got: %s", tt.dataDir, paths[0])
		}

		path, err := tt.scope.CacheDir()
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
