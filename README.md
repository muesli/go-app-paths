# go-app-paths

[![Latest Release](https://img.shields.io/github/release/muesli/go-app-paths.svg)](https://github.com/muesli/go-app-paths/releases)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/muesli/go-app-paths?tab=doc)
[![Build Status](https://github.com/muesli/go-app-paths/workflows/build/badge.svg)](https://github.com/muesli/go-app-paths/actions)
[![Coverage Status](https://coveralls.io/repos/github/muesli/go-app-paths/badge.svg?branch=master)](https://coveralls.io/github/muesli/go-app-paths?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/muesli/go-app-paths)](http://goreportcard.com/report/muesli/go-app-paths)

Lets you retrieve platform-specific paths (like directories for app-data, cache,
config, and logs). It is fully compliant with the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html)
on Unix, but also provides implementations for macOS and Windows systems.

## Installation

Make sure you have a working Go environment (Go 1.2 or higher is required).
See the [install instructions](http://golang.org/doc/install.html).

To install go-app-paths, simply run:

    go get github.com/muesli/go-app-paths

## Platform Support

| Unix        | User Scope                      | System Scope          |
| ----------- | ------------------------------- | --------------------- |
| Data Dir    | ${XDG_DATA_HOME}/app            | /usr/share/app        |
| Cache Dir   | ${XDG_CACHE_HOME}/app           | /var/cache/app        |
| Config Path | ${XDG_CONFIG_HOME}/app/filename | /etc/app/filename     |
| Log Path    | ${XDG_DATA_HOME}/app/filename   | /var/log/app/filename |

| macOS       | User Scope                         | System Scope                      |
| ----------- | ---------------------------------- | --------------------------------- |
| Data Dir    | ~/Library/Application Support/app  | /Library/Application Support/app  |
| Cache Dir   | ~/Library/Caches/app               | /Library/Caches/app               |
| Config Path | ~/Library/Preferences/app/filename | /Library/Preferences/app/filename |
| Log Path    | ~/Library/Logs/app/filename        | /Library/Logs/app/filename        |

| Windows     | User Scope                         | System Scope                      |
| ----------- | ---------------------------------- | --------------------------------- |
| Data Dir    | %LOCALAPPDATA%/app                 | %PROGRAMDATA%/app                 |
| Cache Dir   | %LOCALAPPDATA%/app/Cache           | %PROGRAMDATA%/app/Cache           |
| Config Path | %LOCALAPPDATA%/app/Config/filename | %PROGRAMDATA%/app/Config/filename |
| Log Path    | %LOCALAPPDATA%/app/Logs/filename   | %PROGRAMDATA%/app/Logs/filename   |

## Example

```go
import (
	gap "github.com/muesli/go-app-paths"
)
```

### User Scope

```go
scope := gap.NewScope(gap.User, "app")

scope.DataDirs()
// Unix: ["~/.local/share/app", "/usr/local/share/app", "/usr/share/app"]
// macOS: ["~/Library/Application Support/app"]
// Windows: ["%LOCALAPPDATA%/app"]

scope.ConfigDirs()
// Unix: ["~/.config/app", "/etc/xdg/app", "/etc/app"]
// macOS: ["~/Library/Preferences/app"]
// Windows: ["%LOCALAPPDATA%/app/Config"]

scope.CacheDir()
// Unix: ~/.cache/app
// macOS: ~/Library/Caches/app
// Windows: %LOCALAPPDATA%/app/Cache

scope.DataPath("filename")
// Unix: ~/.local/share/app/filename
// macOS: ~/Library/Application Support/app/filename
// Windows: %LOCALAPPDATA%/app/filename

scope.ConfigPath("filename.conf")
// Unix: ~/.config/app/filename.conf
// macOS: ~/Library/Preferences/app/filename.conf
// Windows: %LOCALAPPDATA%/app/Config/filename.conf

scope.LogPath("filename.log")
// Unix: ~/.local/share/app/filename.log
// macOS: ~/Library/Logs/app/filename.log
// Windows: %LOCALAPPDATA%/app/Logs/filename.log

scope.LookupData("filename")
// Unix: ["~/.local/share/app/filename", "/usr/local/share/app/filename", "/usr/share/app/filename"]
// macOS: ["~/Library/Application Support/app/filename"]
// Windows: ["%LOCALAPPDATA%/app/filename"]

scope.LookupConfig("filename.conf")
// Unix: ["~/.config/app/filename.conf", "/etc/xdg/app/filename.conf", "/etc/app/filename.conf"]
// macOS: ["~/Library/Preferences/app/filename.conf"]
// Windows: ["%LOCALAPPDATA%/app/Config/filename.conf"]
```

### System Scope

```go
scope := gap.NewScope(gap.System, "app")

scope.DataDirs()
// Unix: ["/usr/local/share/app", "/usr/share/app"]
// macOS: ["/Library/Application Support/app"]
// Windows: ["%PROGRAMDATA%/app"]

scope.ConfigDirs()
// Unix: ["/etc/xdg/app", "/etc/app"]
// macOS: ["/Library/Preferences/app"]
// Windows: ["%PROGRAMDATA%/app/Config"]

scope.CacheDir()
// Unix: /var/cache/app
// macOS: /Library/Caches/app
// Windows: %PROGRAMDATA%/app/Cache

scope.DataPath("filename")
// Unix: /usr/local/share/app/filename
// macOS: /Library/Application Support/app/filename
// Windows: %PROGRAMDATA%/app/filename

scope.ConfigPath("filename.conf")
// Unix: /etc/xdg/app/filename.conf
// macOS: /Library/Preferences/app/filename.conf
// Windows: %PROGRAMDATA%/app/Config/filename.conf

scope.LogPath("filename.log")
// Unix: /var/log/app/filename.log
// macOS: /Library/Logs/app/filename.log
// Windows: %PROGRAMDATA%/app/Logs/filename.log

scope.LookupData("filename")
// Unix: ["/usr/local/share/app/filename", "/usr/share/app/filename"]
// macOS: ["/Library/Application Support/app/filename"]
// Windows: ["%PROGRAMDATA%/app/filename"]

scope.LookupConfig("filename.conf")
// Unix: ["/etc/xdg/app/filename.conf", "/etc/app/filename.conf"]
// macOS: ["/Library/Preferences/app/filename.conf"]
// Windows: ["%PROGRAMDATA%/app/Config/filename.conf"]
```
