# Specify the base image for the go app.
FROM golang:1.15
WORKDIR /app
COPY . .
COPY go.mod go.sum ./
RUN go get -u github.com/lib/pq
RUN go build -o main .
CMD ["./main"]