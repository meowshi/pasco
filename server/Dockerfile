FROM golang:1.21.3-alpine3.18 AS build

WORKDIR /server

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg

RUN CGO_ENABLED=0 GOOS=linux go build -o /pasco-server ./cmd/main.go 
# RUN go test -v ./...


FROM alpine:3.18.4 AS release

WORKDIR /

COPY --from=build /pasco-server /pasco-server
COPY credentials.json cert.pem key.unencrypted.pem /

RUN apk add tzdata

EXPOSE 443

ENTRYPOINT ["/pasco-server"]