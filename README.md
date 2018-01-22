go-app-paths
============

Lets you retrieve platform-specific paths (like directories for app-data, cache, config, and logs)

## Installation

Make sure you have a working Go environment (Go 1.2 or higher is required).
See the [install instructions](http://golang.org/doc/install.html).

To install go-app-paths, simply run:

    go get github.com/muesli/go-app-paths

To compile it from source:

    cd $GOPATH/src/github.com/muesli/go-app-paths
    go get -u -v
    go build && go test -v

## Platform Support

| Unix        | User Scope                             | System Scope                          |
| ----------- | -------------------------------------- | ------------------------------------- |
| Data Dir    | ${XDG_DATA_HOME}/appname               | /usr/share/appname                    |
| Cache Dir   | ${XDG_CACHE_HOME}/appname              | /var/cache/appname                    |
| Config Path | ${XDG_CONFIG_HOME}/appname/filename    | /etc/appname/filename                 |
| Log Path    | ${XDG_DATA_HOME}/appname/filename      | /var/log/appname/filename             |

| macOS       | User Scope                             | System Scope                          |
| ----------- | -------------------------------------- | ------------------------------------- |
| Data Dir    | ~/Library/Application Support/appname  | /Library/Application Support/appname  |
| Cache Dir   | ~/Library/Caches/appname               | /Library/Caches/appname               |
| Config Path | ~/Library/Preferences/appname/filename | /Library/Preferences/appname/filename |
| Log Path    | ~/Library/Logs/appname/filename        | /Library/Logs/appname/filename        |

| Windows     | User Scope                             | System Scope                          |
| ----------- | -------------------------------------- | ------------------------------------- |
| Data Dir    | %LOCALAPPDATA%/appname                 | %PROGRAMDATA%/appname                 |
| Cache Dir   | %LOCALAPPDATA%/appname/Cache           | %PROGRAMDATA%/appname/Cache           |
| Config Path | %LOCALAPPDATA%/appname/Config/filename | %PROGRAMDATA%/appname/Config/filename |
| Log Path    | %LOCALAPPDATA%/appname/Logs/filename   | %PROGRAMDATA%/appname/Logs/filename   |

## Example
```go
package main

import (
	"github.com/muesli/go-app-paths"
)

func main() {
	userScope := apppaths.NewScope(apppaths.User, "vendorname", "appname")
	userScope.DataDir()                   // => ~/.local/share/appname
	userScope.CacheDir()                  // => ~/.cache/appname
	userScope.ConfigPath("filename.conf") // => ~/.config/appname/filename.conf
	userScope.LogPath("filename.log")     // => ~/.local/share/appname/filename.log

	systemScope := apppaths.NewScope(apppaths.System, "vendorname", "appname")
	systemScope.DataDir()                   // => /usr/share/appname
	systemScope.CacheDir()                  // => /var/cache/appname
	systemScope.ConfigPath("filename.conf") // => /etc/appname/filename.conf
	systemScope.LogPath("filename.log")     // => /var/log/appname/filename.log
}
```

## Development

API docs can be found [here](http://godoc.org/github.com/muesli/go-app-paths).

[![Build Status](https://travis-ci.org/muesli/go-app-paths.svg?branch=master)](https://travis-ci.org/muesli/go-app-paths)
[![Coverage Status](https://coveralls.io/repos/github/muesli/go-app-paths/badge.svg?branch=master)](https://coveralls.io/github/muesli/go-app-paths?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/muesli/go-app-paths)](http://goreportcard.com/report/muesli/go-app-paths)
