# Copyright © Observerly Ltd.

# //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// #

# Building the binary of the App
FROM golang:latest AS base

# Set the working directory to /app
WORKDIR /usr/src/app

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// #

# Build the binary of the Tailscale VPN
FROM alpine:latest AS tailscale

# `boilerplate` should be replaced with your project name
WORKDIR /usr/src/app

# Set the environment variable for the Tailscale binary download file:
ENV TSFILE=tailscale_latest_amd64.tgz

RUN wget https://pkgs.tailscale.com/stable/${TSFILE} && \
  tar xzf ${TSFILE} --strip-components=1

COPY . ./

# //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// #

# Moving the binary to the 'final Image' to make it smaller
FROM golang:alpine as development

# Install gcc make and libc-dev to allow Makefile compilation
RUN apk update && apk add --no-cache gcc make libc-dev neovim neovim-doc ripgrep

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
RUN apk add --no-cache zsh git wget

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
