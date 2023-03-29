FROM golang:1.20-alpine AS builder

# Set necessary environmet variables needed for our image
ENV CGO_ENABLED=0
ENV config=docker

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
# go build -o [name] [path to file]
RUN go build -o server cmd/server/main.go

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY . .
COPY --from=builder /dist/app /
# Copy the code into the container

EXPOSE 5000
EXPOSE 5001
EXPOSE 5002
EXPOSE 7071

# Command to run the executable
ENTRYPOINT ["/server"]