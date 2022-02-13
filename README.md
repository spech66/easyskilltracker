[![Go Report Card](https://goreportcard.com/badge/github.com/spech66/easyskilltracker)](https://goreportcard.com/report/github.com/spech66/easyskilltracker) [![GoDoc](https://godoc.org/github.com/spech66/easyskilltracker?status.svg)](https://godoc.org/github.com/spech66/easyskilltracker) 

# EasySkillTracker

Super easy flat file based selfhosted skill tracker (learn new skills from multiple resources).

## Features

* Have all related skills in one file
* Share files (courses) with others
* Easy to use interface
* Mobile/Tablet ready (large buttons, ...)
* Selfhosted
* Free and open

## Screenshots

[![Start](https://raw.githubusercontent.com/spech66/easyskilltracker/master/_screenshots/s_001_start.png "Start")](https://raw.githubusercontent.com/spech66/easyskilltracker/master/_screenshots/001_start.png)
[![Skills](https://raw.githubusercontent.com/spech66/easyskilltracker/master/_screenshots/s_002_skills_01.png "Skills")](https://raw.githubusercontent.com/spech66/easyskilltracker/master/_screenshots/002_skills_01.png)
[![Edit](https://raw.githubusercontent.com/spech66/easyskilltracker/master/_screenshots/s_003_edit_01.png "Edit")](https://raw.githubusercontent.com/spech66/easyskilltracker/master/_screenshots/003_edit_01.png)

[More screenshots](https://github.com/spech66/easyskilltracker/tree/master/_screenshots)

## Security information

This daemon is developed for use in a private network. For securing the access use a proxy like nginx (see below)!

## Build and run from source

Make sure you have the [Go Programming Language](https://golang.org/) tools set up an ready.

### Linux

Checkout the code to your `GOPATH` directory, build the server and run it.

```bash
go get github.com/spech66/easyskilltracker
cd $GOPATH/src/github.com/spech66/easyskilltracker
go build
./easyskilltracker -data data/example
```

### Windows

Checkout the code to your `GOPATH` directory.

```cmd
go get github.com/spech66/easyskilltracker
cd %GOPATH%\src\github.com\spech66\easyskilltracker
go build
easyskilltracker.exe -data data\example
```

## nginx reverse proxy with basic auth

This daemon is intend for use in a local network. Make sure that the server port (default 8045) is not exposed to the internet.
In case you want to host this deamon in the internet set up nginx as a revery proxy for it using basic authentication.

```bash
sudo sh -c "echo -n 'admin:' >> /etc/nginx/.htpasswd"
sudo sh -c "openssl passwd -apr1 >> /etc/nginx/.htpasswd"
sudo vi /etc/nginx/sites-enabled/default
```

```nginx
server {
        listen 8645;
        server_name _;
        auth_basic "EasySkillTracker";
        auth_basic_user_file /etc/nginx/.htpasswd;

        location / {
                proxy_pass http://localhost:8045;
        }
}
```

## Dependencies

* [Go echo web framework](github.com/labstack/echo) - High performance, minimalist Go web framework

## Development

* Update Modules: `go mod tidy`
* Run: `go run . -data data\example`
