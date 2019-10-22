FROM golang:1.13.1 as builder
# docker build -t vanessa/nu-plugin-len .
WORKDIR /code
COPY . /code
RUN make
FROM quay.io/nushell/nu-base:devel
LABEL Maintainer vsochat@stanford.edu
COPY --from=builder /code/nu_plugin_len /usr/local/bin

