![Commit activity](https://img.shields.io/github/commit-activity/w/thibault-cne/projet_nouveau_site_integration)
![License](https://img.shields.io/github/license/thibault-cne/projet_nouveau_site_integration)
![Go Version](https://img.shields.io/github/go-mod/go-version/thibault-cne/projet_nouveau_site_integration/master?filename=backend%2Fgo.mod)
# Projet d'un nouveau site web pour l'intégration
A project for a new site for the integration club of TELECOM Nancy. I work on this project with @BigBaz54.

# Table of Content

1. [Installation guide](#installation-guide)

## Installation guide

``` shell
$ git clone https://github.com/thibault-cne/projet_nouveau_site_integration
$ cd projet_nouveau_site_integration
```

### Install Docker

**MACOS**

``` shell
$ brew update
$ brew cask install docker
```

**Linux (Ubuntu, Debian, Linux Mint)**

``` shell
$ sudo apt update
$ sudo apt-install docker
```

### Install golang

You can install golang on the official page : [official golang page](https://go.dev/doc/install)

## Start the application

**Docker**

``` shell
$ sudo docker-compose build
$ sudo docker-compose up
```

**Separation of backend and frontend**

``` shell
$ cd backend
$ go mod download
$ go run main.go
```

``` shell
$ cd frontend
$ npm install package-lock.json
$ npm run serve
```
