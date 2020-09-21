# Golang Microservice starter

Using go 1.11.5 version to build a Restful API Server.

Integration with bellow extensions.

### Used modules:
- Http server: [Gorilla Mux](https://github.com/gorilla/mux)

## Installation

Firstly you need to set `GOROOT` and `GOPATH` right path format. example bellow
```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
```
Second install all modules

```bash
go get -u github.com/gorilla/mux
```
#### Example Usage

```bash
go build -d micro
./micro
```
## Go Application Structure 
```
.
├── CODE_OF_CONDUCT.md
├── CONTRIBUTING.md
├── Dockerfile
├── LICENSE
├── README.md
├── config
│   └── config.go
├── contributors.txt
├── docker-compose.yml
├── helper
│   └── helper.go
├── main.go
├── main_test.go
├── micro.log
├── models
│   └── main.go
├── routes
│   ├── 404.go
│   ├── home.go
│   └── main.go
├── services
│   └── redis.go
└── utils
    ├── logger.go
    ├── middleware.go
    └── response.go

6 directories, 20 files
```

## Changelog

- Version 1.0 : Beta version

