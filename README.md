# Projet d'un nouveau site web pour l'intégration
A project for a new site for the integration club of TELECOM Nancy


# Guide d'installation

``` shell
$ git clone https://github.com/thibault-cne/projet_nouveau_site_integration
$ cd project2-E19
```

## Installer Docker

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

# Lancer les applications

## Partie website

**Docker**

``` shell
$ cd Website
$ sudo docker-compose build
$ sudo docker-compose up
```

**Séparation des images**

``` shell
$ cd Website/backend
$ python3 -m venv venv
$ source venv/bin/activate
$ pip install -r requirements.txt
$ python3 tazmouApp.py
```

``` shell
$ cd Website/frontend
$ npm install package-lock.json
$ npm run serve
```