# 1.step-build the executable binary
FROM  sha256:1db127655b32aa559e32ed3754ed2ea735204d967a433e4b605aed1dd44c5084 as builder
#FROM golang:alpine AS builder

# Git is required for fetching the dependencies, so we install it
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Create a User with Name/PW stored in envs and set him up
ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

# Set a workingdirectory
WORKDIR $GOPATH/src/mypackage/myapp
COPY . .

# Fetch dependencies
RUN go get -d -v

# Build the binary
RUN go build -o /go/bin/hello

# ---------------------- #

# 2.step build a minimal image

FROM scratch

# Import the user and group files from the builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# copy our binary, from the build-step, to our new image
COPY --from=builder /go/bin/hello /go/bin/hello

# use the new unprivileged user from here on
USER appuser:appuser

# port to expose
EXPOSE 9876

# run the binary
ENTRYPOINT ["/go/bin/hello"]