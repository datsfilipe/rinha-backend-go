FROM golang:1.21.1

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd 
EXPOSE 80
CMD ["/api"]