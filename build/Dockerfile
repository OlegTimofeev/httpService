FROM golang:1.14 as builder

COPY . /build

WORKDIR /build

RUN make build

CMD ["/build/cmd"]