FROM golang:alpine as builder

ARG version="0.0.1-dev"
ENV version=$version

WORKDIR /app

RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build \
    -ldflags="-X 'home-hap/internal/configs.Version=$version'" \
    -o ./home-hap ./cmd/server

FROM scratch

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/config.yml /app/config.yml
COPY --from=builder /app/home-hap /app/home-hap

CMD [ "/app/home-hap" ]