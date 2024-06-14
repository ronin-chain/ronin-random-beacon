#### BASE ####
FROM golang:1.18.5-buster as builder

WORKDIR /opt

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM debian:buster

RUN apt-get update  && apt-get install -y ca-certificates gnupg lsb-release wget
RUN update-ca-certificates

WORKDIR /opt

ENV RONIN_RANDOM_BEACON_PARAMS ''

COPY --from=builder /go/pkg/mod/github.com/\!cosm\!wasm/wasmvm@v*/api/libwasmvm.so /usr/lib/libwasmvm.so
RUN chmod 755 /usr/lib/libwasmvm.so

COPY --from=builder /opt/ronin-random-beacon .
COPY --from=builder /opt/config ./config
COPY --from=builder /opt/entrypoint.sh .

ENTRYPOINT ["./entrypoint.sh"]
