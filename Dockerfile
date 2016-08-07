FROM golang:latest
MAINTAINER Nathan Osman <nathan@quickmediasolutions.com>

# Add the source files
ADD *.go /go/src/github.com/nathan-osman/go-shorten/

# Fetch dependencies
RUN go get ./...

# Build the application
RUN go install github.com/nathan-osman/go-shorten

# Add the configuration file
ADD config.json /etc/go-shorten/config.json

# Expose the HTTP port
EXPOSE 80

# Specify the command to run
CMD go-shorten /etc/go-shorten/config.json
