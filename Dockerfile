ARG GO_VERSION=1.17

FROM golang:${GO_VERSION}  AS builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /src

COPY ./ ./
RUN go get -d

RUN go build \
    -installsuffix 'static' \
    -o /app .

FROM scratch AS final

COPY --from=builder /app /app
COPY --from=builder /src/ui /ui

EXPOSE 4000

ENTRYPOINT ["/app"]
