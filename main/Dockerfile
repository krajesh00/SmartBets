# FROM golang:1.18.1-buster

# WORKDIR /app

# COPY go.mod ./

# RUN go mod download

# COPY main.go .

# RUN go build -o /godocker

# EXPOSE 8080

# CMD [ "/godocker" ]

FROM golang:1.18.1-buster AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY *.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o /opt/go-docker-multistage

FROM alpine:latest
COPY --from=builder /opt/go-docker-multistage /opt/go-docker-multistage
EXPOSE 8080
ENTRYPOINT [ "/opt/go-docker-multistage" ]
