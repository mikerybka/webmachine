FROM ubuntu:23.04

RUN apt update
RUN apt install -y curl

RUN curl https://webmachine.dev/bin/amd64-linux/webmachine > /usr/local/bin/webmachine
RUN chmod +x /usr/local/bin/webmachine

ENTRYPOINT ["/usr/local/bin/webmachine"]
