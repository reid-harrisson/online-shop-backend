# syntax=docker/dockerfile:1

FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Prepare swag cli
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN export PATH=$(go env GOPATH)/bin:$PATH

# Copy the source code. Note the slash at the end, as explained in
COPY . .

RUN swag init -g ./cmd/main.go

# Download Go modules
RUN go mod download

RUN go build -o ./main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]