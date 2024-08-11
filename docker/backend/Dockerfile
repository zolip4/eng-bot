# Specify the base image with Go 1.22
FROM golang:1.22-alpine

# Set the working directory
WORKDIR /app

# Install Air with the correct import path
RUN go install github.com/air-verse/air@latest

# Copy the project files
COPY ./app .

# Install dependencies
RUN go mod tidy

# Expose port 80
EXPOSE 80

# Use Air for development
CMD ["air"]
