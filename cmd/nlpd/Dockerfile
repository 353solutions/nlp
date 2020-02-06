# Docker container for nlpd
# Run
#   docker run -p8080:8080 tdg/nlp


FROM golang:1.13-buster as build

WORKDIR /nlp
COPY . .
ENV CGO_ENABLED=0
ARG VERSION=unknown
# -ldflags="-X ..." injects version to the executable
RUN go build -ldflags="-X main.Version=${VERSION}" ./cmd/nlpd

FROM debian:buster-slim
COPY --from=build /nlp/nlpd /usr/local/bin

ENV NLP_PORT 8080
CMD nlpd
