FROM golang:1.15-alpine

RUN apk add --no-cache git

WORKDIR /golang_jobs

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/graphql-api ./cmd/api/graphql.go
RUN go build -o ./out/scraper ./cmd/scraper/scraper.go


# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./out/graphql-api"]

