# Use Golang image as base
FROM golang:alpine
WORKDIR /app
COPY ../Zadanie_4_Go .  
RUN go build -o main .
ENV PORT=80
EXPOSE 80
CMD ["./main"]
