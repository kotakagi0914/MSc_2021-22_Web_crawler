# MSc_2021-22_Web_crawler
## Abstract
This is a repository for MSc Cyber Security individual project at City, University of London.
The code in the repository contains web crawlers and their prototypes that have the features to bypass or solve reCAPTCHA v3.

## How to run
### Preparation
Before running web crawlers, you need to run selenium servers.
Here are just two examples:
* Run a docker container individually depending on your browser choise.
```
## Run the docker image with GeckoDriver
$ docker run --rm -d -p 4444:4444 seleniarm/standalone-firefox:<tag name>

## Run the docker image with Chromium
$ docker run --rm -d -p 4445:4444 seleniarm/standalone-chromium:<tag name>
```

* Use `docker-compose` to run everything you need.
  You need to use this if you want to target the mock webserver (`https://github.com/sheva0914/MSc_2021-22_Mock_webserver`).
```
$ docker-compose up -d
```

### Run main crawlers
```
## Run the crawler for Firefox
$ go run cmd/crawler/firefox/crawler.go

## Run the crawler for Chrome
$ go run cmd/crawler/chrome/crawler.go
```


### Run Prototype crawlers
```
## Run the crawler for Firefox
$ go run prototypes/firefox/crawler.go

## Run the crawler for Chrome
$ go run prototypes/chrome/crawler.go
```
