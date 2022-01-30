FROM golang:latest

WORKDIR /app

COPY . .

RUN ls -lha

RUN go build -o math

CMD ["./math"]