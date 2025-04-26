FROM golang:latest
LABEL authors="kirysha"

EXPOSE 9000

WORKDIR ./server

COPY . .

RUN go build -o main.go

CMD ["./main.go"]