FROM golang:1.18.1-buster

WORKDIR /app

COPY . .

# RUN go mod download

RUN go build -o /godocker

EXPOSE 8080

CMD [ "/godocker" ]
