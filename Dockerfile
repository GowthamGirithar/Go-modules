# Start from golang latest base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Gowtham Girithar Srirangasamy"

# Set the Current Working Directory inside the container
WORKDIR /go/src/employee-service

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .
# Download all the dependencies
RUN go get ./...
#build
RUN go build

EXPOSE 8080

# Run the executable
ENTRYPOINT "./employee-service"