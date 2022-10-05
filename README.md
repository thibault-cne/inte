![Commit activity](https://img.shields.io/github/commit-activity/w/thibault-cne/projet_nouveau_site_integration)
![License](https://img.shields.io/github/license/thibault-cne/projet_nouveau_site_integration)
![Go Version](https://img.shields.io/github/go-mod/go-version/thibault-cne/projet_nouveau_site_integration/master?filename=backend%2Fgo.mod)

# Projet d'un nouveau site web pour l'intégration

A project for a new site for the integration club of TELECOM Nancy. I work on this project with @BigBaz54.

# Table of Content

1. [Installation guide](#installation-guide)

## Installation guide

``` bash
git clone https://github.com/thibault-cne/projet_nouveau_site_integration
cd projet_nouveau_site_integration
```

### Install Docker

**MACOS**

``` bash
brew update
brew cask install docker
```

**Linux (Ubuntu, Debian, Linux Mint)**

``` bash
sudo apt update
sudo apt-install docker
```

### Install golang

You can install golang on the official page : [official golang page](https://go.dev/doc/install)

## Start the application

**Docker**

``` bash
sudo docker-compose build
sudo docker-compose up
```

**Separation of backend and frontend**

``` bash
cd backend
go mod download
go run main.go
```

``` bash
cd frontend
npm install package-lock.json
npm run serve
```

## Client OAuth

- Open the Google Developers Console and make sure that you are logged into your telecomnancy.net account.
- Use the arrow next to the Google APIs logo to open the project list and create a new project.
- Open the "OAuth consent screen" page and configure its appearance : it is the authentication page shown to users.
- Open the "credentials" tab and click "create credentials", choosing the "OAuth client ID" option and setting the application type as "Web application".
- Leave the "Authorised JavaScript origins" text box empty, and inside "Authorised redirect URIs" enter the URL that the web service will be accessible with (https://custom-subdomain.telecomnancy.net/path/to/redirect)
- Once your OAuth credentials are created, use the download icon to download the JSON file containing the credentials and move it into a file named client_secrets.json in the root of the project.

## Licence

This project is licensed under the MIT license. Please read the LICENSE file.
