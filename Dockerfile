FROM golang:latest
LABEL authors="kirylputseyeu"

EXPOSE 9000

WORKDIR ./server

COPY . .

RUN go build -o main.go

CMD ["./main.go"]