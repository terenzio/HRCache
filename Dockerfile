# Use an official Golang runtime as a parent image
FROM golang:1.22.1

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace.
ADD . /app

# Build your program
RUN go build -o main .

# Run the program when the container launches
CMD ["./main"]
