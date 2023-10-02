FROM golang:1.21.1

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN GOOS=linux go build -o /api ./cmd 
EXPOSE 80
CMD ["/api"]
