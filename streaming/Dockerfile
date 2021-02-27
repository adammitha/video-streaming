FROM golang:1.16-alpine

WORKDIR /src

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o streaming

EXPOSE 8080

CMD ["./streaming"]