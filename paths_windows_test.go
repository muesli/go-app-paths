// +build windows

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
		{NewScope(System, "", "foobar"), "C:\\ProgramData\\foobar", "C:\\ProgramData\\foobar\\Cache", "C:\\ProgramData\\foobar\\Config\\foobar.conf", "C:\\ProgramData\\foobar\\Logs\\foobar.log"},
		{NewScope(User, "", "foobar"), "C:\\Users\\runneradmin\\AppData\\Local\\foobar", "C:\\Users\\runneradmin\\AppData\\Local\\foobar\\Cache", "C:\\Users\\runneradmin\\AppData\\Local\\foobar\\Config\\foobar.conf", "C:\\Users\\runneradmin\\AppData\\Local\\foobar\\Logs\\foobar.log"},
		{NewCustomHomeScope("C:\\tmp", "", "foobar"), "C:\\tmp", "C:\\tmp\\Cache", "C:\\tmp\\Config\\foobar.conf", "C:\\tmp\\Logs\\foobar.log"},
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
