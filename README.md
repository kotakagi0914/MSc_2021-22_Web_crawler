# MSc_2021-22_Web_crawler
## Abstract
This is a repository for MSc Cyber Security individual project at City, University of London.
The code in the repository contains web crawlers and their prototypes that have the features to bypass or solve reCAPTCHA v3.

## How to run
### Preparation
Before running web crawlers, you need to run selenium servers.
Here are just two examples:
* Run a docker container individually depending on your browser choise
```
## Run the docker image with GeckoDriver
$ docker run --rm -d -p 4444:4444 seleniarm/standalone-firefox:<tag name>
## Run the docker image with Chromium
$ docker run --rm -d -p 4445:4444 seleniarm/standalone-chromium:<tag name>
```

* Use `docker-compose` to run everything you need
```
$ docker-compose up -d
```

### Run Prototype for Firefox
```
## Run the crawler For firefox
$ go run prototypes/firefox/crawler.go
```

### Run Prototype for Chromium
```
## Run the crawler for Chrome
$ go run prototypes/chrome/crawler.go
```
