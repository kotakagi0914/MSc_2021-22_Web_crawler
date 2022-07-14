# MSc_2021-22_Web_crawler
## Abstract
This is a repository for MSc Cyber Security individual project at City, University of London.
The code in the repository contains web crawlers and their prototypes that have the features to bypass or solve reCAPTCHA v3.

## How to run
### Prototype for Firefox
```
## Run the docker image including GeckoDriver
$ docker run --rm -d -p 4444:4444 seleniarm/standalone-firefox:<tag name>

## Run the crawler
$ go run prototypes/firefox/crawler.go
```

### Prototype for Chromium
```
## Run the docker image including GeckoDriver
$ docker run --rm -d -p 4444:4444 seleniarm/standalone-chromium:<tag name>

## Run the crawler
$ go run prototypes/chrome/crawler.go
```
