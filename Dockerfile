# Stage 1: Build
FROM docker.io/golang:1.21.4 AS builder

# Set the working directory in the Docker image
WORKDIR /app

# Copy go.mod and go.sum files from your 'src' directory
COPY src/go.mod src/go.sum* ./

# Fetch dependencies
RUN go mod download

# Copy the Go source files from the 'src' directory
COPY src/ .

# Build the application
# Adjust the build command if necessary, depending on the structure within your 'src' folder
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ballerapp .

# Stage 2: Runtime
FROM registry.access.redhat.com/ubi8/ubi-minimal

# Set the working directory in the runtime image
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/ballerapp .

# Copy templates and static directories if they exist
# Adjust these paths if your templates and static files are inside the 'src' directory
COPY src/templates/ templates/
COPY src/static/ static/

# Expose the port the app runs on
EXPOSE 8080

# Define the entry point
ENTRYPOINT ["./ballerapp"]

