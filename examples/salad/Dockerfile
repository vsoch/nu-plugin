FROM golang:1.13.1 as builder
# docker build -t vanessa/nu-plugin-salad .
WORKDIR /code
COPY . /code
RUN go get github.com/vsoch/nu-plugin && make
FROM quay.io/nushell/nu-base:devel
LABEL Maintainer vsochat@stanford.edu
COPY --from=builder /code/nu_plugin_salad /usr/local/bin
ENTRYPOINT ["/bin/bash"]
