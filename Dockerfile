FROM golang:1.22.1

WORKDIR /app
COPY . .

RUN go mod download

ENV ADDR=:3001

# Build main.go
RUN go build -o main .

CMD ["/app/main"]