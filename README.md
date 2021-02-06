# Golang job listing scraper
This is a program to scrape couple websites (golang.cafe, remoteok.io, welovegolang.com) for golang job listings, persist them to mysql and serve them trough graphql endpoint.

## Requirements
Docker and docker-compose

## Running
Clone this repository and run

```sh
docker-compose build
docker-compose up -d
```

To run scrapper you need attach to containers shell and run binary

```sh
docker-compose exec job-listings sh
./out/scraper
```
