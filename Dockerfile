FROM ubuntu:23.04

RUN apt update
RUN apt install -y curl wget

# TODO: install python, ruby, node, etc.

# Install go
RUN wget https://go.dev/dl/go1.21.1.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin:/root/go/bin
RUN rm go1.21.1.linux-amd64.tar.gz

# Initialize go workspace
RUN mkdir -p /root/go/src
WORKDIR /root/go/src
RUN go work init

# Copy code to image
COPY . /root/go/src/github.com/mikerybka/webmachine
WORKDIR /root/go/src/github.com/mikerybka/webmachine
RUN go work use .

# Build webmachine
RUN go install github.com/mikerybka/webmachine/cmd/webmachine
