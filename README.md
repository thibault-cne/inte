# Projet d'un nouveau site web pour l'intégration
A project for a new site for the integration club of TELECOM Nancy

# Table of Content

1. [Installation guide](#installation-guide)

## Installation guide

``` shell
$ git clone https://github.com/thibault-cne/projet_nouveau_site_integration
$ cd project2-E19
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

### Partie website

**Docker**

``` shell
$ sudo docker-compose build
$ sudo docker-compose up
```

**Séparation des images**

``` shell
$ cd backend
$ go mod download
$ go run .
```

``` shell
$ cd frontend
$ npm install package-lock.json
$ npm run serve
```
