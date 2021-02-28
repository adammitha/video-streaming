FROM golang:1.16-alpine

WORKDIR /src

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o google-storage

EXPOSE 8080

CMD ["./google-storage"]