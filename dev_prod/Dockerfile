# References the base image
FROM golang:1.20.2-bullseye

# installs bee tool globaly, used to live-reload our code during development
RUN go install github.com/beego/bee/v2@latest

# Env variables. used to configure the use of our vendor location for the GO modules
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

# creates a folder for the code and make it the active directory
ENV APP_HOME /go/src/mathapp
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

# port of our port we work with
EXPOSE 8010
#uses the bee comand to start our application
CMD ["bee", "run"]

