//go:build !darwin && !windows
// +build !darwin,!windows

package gap

import (
	"fmt"
	"testing"
)

func TestPaths(t *testing.T) {
	tests := []struct {
		scope      *Scope
		dataDirs   []string
		configDirs []string
		cacheDir   string
		configFile string
		dataFile   string
		logFile    string
	}{
		{
			scope:      NewScope(System, "foobar"),
			dataDirs:   []string{"/usr/local/share/foobar", "/usr/share/foobar"},
			configDirs: []string{"/etc/xdg/foobar", "/etc/foobar"},
			cacheDir:   "/var/cache/foobar",
			configFile: "/etc/xdg/foobar/foobar.conf",
			dataFile:   "/usr/local/share/foobar/foobar.data",
			logFile:    "/var/log/foobar/foobar.log",
		},
		{
			scope:      NewVendorScope(System, "barcorp", "foobar"),
			dataDirs:   []string{"/usr/local/share/barcorp/foobar", "/usr/share/barcorp/foobar"},
			configDirs: []string{"/etc/xdg/barcorp/foobar", "/etc/barcorp/foobar"},
			cacheDir:   "/var/cache/barcorp/foobar",
			configFile: "/etc/xdg/barcorp/foobar/foobar.conf",
			dataFile:   "/usr/local/share/barcorp/foobar/foobar.data",
			logFile:    "/var/log/barcorp/foobar/foobar.log",
		},
		{
			scope:      NewScope(User, "foobar"),
			dataDirs:   []string{"~/.local/share/foobar", "/usr/local/share/foobar", "/usr/share/foobar"},
			configDirs: []string{"~/.config/foobar", "/etc/xdg/foobar", "/etc/foobar"},
			cacheDir:   "~/.cache/foobar",
			configFile: "~/.config/foobar/foobar.conf",
			dataFile:   "~/.local/share/foobar/foobar.data",
			logFile:    "~/.local/share/foobar/foobar.log",
		},
		{
			scope:      NewCustomHomeScope("/tmp", "", "foobar"),
			dataDirs:   []string{"/tmp/.local/share/foobar"},
			configDirs: []string{"/tmp/.config/foobar"},
			cacheDir:   "/tmp/.cache/foobar",
			configFile: "/tmp/.config/foobar/foobar.conf",
			dataFile:   "/tmp/.local/share/foobar/foobar.data",
			logFile:    "/tmp/.local/share/foobar/foobar.log",
		},
	}

	for _, tt := range tests {
		paths, err := tt.scope.DataDirs()
		if err != nil {
			t.Errorf("Error retrieving data dir: %s", err)
		}

		if len(paths) != len(tt.dataDirs) {
			fmt.Println(paths)
			t.Fatalf("Expected %d results, got %d", len(tt.dataDirs), len(paths))
		}
		for i := range paths {
			if paths[i] != expandUser(tt.dataDirs[i]) {
				t.Errorf("Expected data dir: %s - got: %s", tt.dataDirs[i], paths[i])
			}
		}

		paths, err = tt.scope.ConfigDirs()
		if err != nil {
			t.Errorf("Error retrieving data dir: %s", err)
		}

		if len(paths) != len(tt.configDirs) {
			fmt.Println(paths)
			t.Fatalf("Expected %d results, got %d", len(tt.configDirs), len(paths))
		}
		for i := range paths {
			if paths[i] != expandUser(tt.configDirs[i]) {
				t.Errorf("Expected data dir: %s - got: %s", tt.configDirs[i], paths[i])
			}
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

func TestConfigLookups(t *testing.T) {
	tests := []struct {
		scope      *Scope
		configFile string
		result     []string
	}{
		{NewScope(System, "ssh"), "sshd_config", []string{"/etc/ssh/sshd_config"}},
		{NewScope(User, "ssh"), "sshd_config", []string{"/etc/ssh/sshd_config"}},
	}

	for _, tt := range tests {
		r, err := tt.scope.LookupConfig(tt.configFile)
		if err != nil {
			t.Errorf("Error looking up config: %s", err)
		}
		if len(r) != 1 {
			fmt.Println(r)
			t.Fatalf("Expected 1 result, got %d results", len(r))
		}
		if r[0] != tt.result[0] {
			t.Errorf("Expected config file: %s - got: %s", tt.result[0], r[0])
		}
	}
}

func TestDataLookups(t *testing.T) {
	tests := []struct {
		scope    *Scope
		dataFile string
		result   []string
	}{
		{NewVendorScope(System, "terminfo", "x"), "xterm+256color", []string{"/usr/share/terminfo/x/xterm+256color"}},
		{NewVendorScope(User, "terminfo", "x"), "xterm+256color", []string{"/usr/share/terminfo/x/xterm+256color"}},
	}

	for _, tt := range tests {
		r, err := tt.scope.LookupDataFile(tt.dataFile)
		if err != nil {
			t.Errorf("Error looking up data file: %s", err)
		}
		if len(r) != 1 {
			fmt.Println(r)
			t.Fatalf("Expected 1 result, got %d results", len(r))
		}
		if r[0] != tt.result[0] {
			t.Errorf("Expected data file: %s - got: %s", tt.result[0], r[0])
		}
	}
}
