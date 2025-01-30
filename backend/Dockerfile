# Step 1: Use an official Golang image to build the app
FROM golang:1.23 AS builder

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the Go module files and download the dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Step 4: Copy the rest of the Go application code into the container
COPY . .

# Step 5: Build the Go application
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Step 6: Create a smaller image to run the app
FROM alpine:latest

# Step 7: Set the working directory inside the container
WORKDIR /root/

# Step 8: Copy the compiled binary from the builder image
COPY --from=builder /app/main .

# Step 9: Expose the port the app will run on
EXPOSE 8080

# Step 10: Command to run the Go server
CMD ["./main"]
