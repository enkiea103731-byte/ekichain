FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go build -o eki main.go
EXPOSE 8080 6001
CMD ["./eki"]
