# Copyright © Observerly Ltd.

# //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// #

ARG BUILDER_IMAGE=golang:1.22-alpine

# //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// #

# Building the binary of the App
FROM ${BUILDER_IMAGE} AS base

RUN apk update && apk add --no-cache git openssh-client ca-certificates tzdata && update-ca-certificates

# Set the working directory to /app
WORKDIR /usr/src/app

# Copy all the Code and stuff to compile everything
COPY . .

RUN apk update && apk add --no-cache git openssh-client ca-certificates tzdata && update-ca-certificates

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go build -o /server cmd/api/main.go

# //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// #

# Moving the binary to the 'final Image' to make it smaller
FROM golang:alpine as development

# Install gcc make and libc-dev to allow Makefile compilation
RUN apk update && apk add --no-cache curl git gcc jq libc-dev make ripgrep wget

# Set the working directory to /app
WORKDIR /usr/src/app

# Ensure staticcheck is executable and in the PATH
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

# Ensure go-critic is executable and in the PATH
RUN go install github.com/go-critic/go-critic/cmd/gocritic@latest

# Install buf CLI:
RUN go install github.com/bufbuild/buf/cmd/buf@latest

# Install protoc-gen-go and protoc-gen-go-grpc:
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Install protoc-gen-go connect:
RUN go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

# Install protect-gen-go-gorm
RUN go install github.com/infobloxopen/protoc-gen-gorm@latest

# Download and install the latest release of the Atlas CLI:
RUN curl -sSf https://atlasgo.sh | sh

# Set the GOPATH environment variable
ENV $PATH:$(go env GOPATH)/bin

# Install szh shell:
RUN apk add --no-cache zsh

# Install oh-my-zsh:
# Uses "Spaceship" theme with some customization. Uses some bundled plugins and installs some more from github
RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.5/zsh-in-docker.sh)" -- \
  -t https://github.com/denysdovhan/spaceship-prompt \
  -a 'SPACESHIP_PROMPT_ADD_NEWLINE="false"' \
  -a 'SPACESHIP_PROMPT_SEPARATE_LINE="false"' \
  -p git \
  -p ssh-agent \
  -p https://github.com/zsh-users/zsh-autosuggestions \
  -p https://github.com/zsh-users/zsh-completions

# Copy across all the files
COPY . /usr/src/app

# //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// #

# Moving the binary to the 'final Image' to make it smaller
FROM scratch AS production

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=base /server .

CMD ["/server"]

# //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// #
