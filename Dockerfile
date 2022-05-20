#### We will use Multi-stage Docker Build ####
#### Build Executable binary ####
FROM golang:latest as builder

# Set Environment Variables
ENV CGO_ENABLED 0
ENV GOOS linux

# Copy the code
WORKDIR $GOPATH/src/go-rest-api/
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build binary
RUN go build -a -installsuffix cgo -o /app/main .

#### Build Small(Tiny) Image ####
FROM alpine:latest

# RUN apk --no-cache add ca-certificates
WORKDIR /root/

# We need quotes.json file. So lets include it as well.
COPY quotes.json ./

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Run the service
CMD [ "./main" ]