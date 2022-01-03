//go:build windows
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
		dataFile   string
		logFile    string
	}{
		{
			scope:      NewScope(System, "foobar"),
			dataDir:    "C:\\ProgramData\\foobar",
			cacheDir:   "C:\\ProgramData\\foobar\\Cache",
			configFile: "C:\\ProgramData\\foobar\\Config\\foobar.conf",
			dataFile:   "C:\\ProgramData\\foobar\\foobar.data",
			logFile:    "C:\\ProgramData\\foobar\\Logs\\foobar.log",
		},
		{
			scope:      NewVendorScope(System, "barcorp", "foobar"),
			dataDir:    "C:\\ProgramData\\barcorp\\foobar",
			cacheDir:   "C:\\ProgramData\\barcorp\\foobar\\Cache",
			configFile: "C:\\ProgramData\\barcorp\\foobar\\Config\\foobar.conf",
			dataFile:   "C:\\ProgramData\\barcorp\\foobar\\foobar.data",
			logFile:    "C:\\ProgramData\\barcorp\\foobar\\Logs\\foobar.log",
		},
		{
			scope:      NewScope(User, "foobar"),
			dataDir:    "C:\\Users\\runneradmin\\AppData\\Local\\foobar",
			cacheDir:   "C:\\Users\\runneradmin\\AppData\\Local\\foobar\\Cache",
			configFile: "C:\\Users\\runneradmin\\AppData\\Local\\foobar\\Config\\foobar.conf",
			dataFile:   "C:\\Users\\runneradmin\\AppData\\Local\\foobar\\foobar.data",
			logFile:    "C:\\Users\\runneradmin\\AppData\\Local\\foobar\\Logs\\foobar.log",
		},
		{
			scope:      NewCustomHomeScope("C:\\tmp", "", "foobar"),
			dataDir:    "C:\\tmp",
			cacheDir:   "C:\\tmp\\Cache",
			configFile: "C:\\tmp\\Config\\foobar.conf",
			dataFile:   "C:\\tmp\\foobar.data",
			logFile:    "C:\\tmp\\Logs\\foobar.log",
		},
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

		path, err = tt.scope.DataPath(tt.scope.App + ".data")
		if err != nil {
			t.Errorf("Error retrieving data path: %s", err)
		}
		if path != expandUser(tt.dataFile) {
			t.Errorf("Expected data path: %s - got: %s", tt.dataFile, path)
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
