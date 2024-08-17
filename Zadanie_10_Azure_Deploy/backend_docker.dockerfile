# Use Golang image as base
FROM golang:alpine
WORKDIR /app
COPY ../Zadanie_4_Go .  
RUN go build -o main .
CMD ["./main"]
