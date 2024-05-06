### This is a multi-stage build. Completing it is left as an exercise for the reader.
### Stage 1: Build go binary
### Stage 2: Build the final image

## Stage 1
# Use the most recent 'go-toolset' from 'registry.access.redhat.com/ubi9/'. Make sure to name the stage so it can be referenced later
FROM registry.access.redhat.com/ubi9/go-toolset:latest AS builder

# Copy the go source code into the container
COPY . .

# Build the go binary. The binary output is stored in '/opt/app-root/fibonacci-golang'
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /opt/app-root/fibonacci-golang

# Stage 2
# Use the most recent version of 'ubi-micro' from 'registry.access.redhat.com/ubi9/'
FROM registry.access.redhat.com/ubi9/ubi-micro:latest

# Ensure the container can be run as a non-root user by copying the '/etc/passwd' file from the previous stage
COPY --from=builder /etc/passwd /etc/passwd

# Copy the go binary from the builder stage
COPY --from=builder /opt/app-root/fibonacci-golang /opt/app-root/fibonacci-golang

# Ensure the binary is executable
RUN chmod +x /opt/app-root/fibonacci-golang

# Set the user to the 'default' user
USER default

# Ensure the container listens on port 8080
EXPOSE 8080

# Set the default command for the container
CMD ["/opt/app-root/fibonacci-golang"]
