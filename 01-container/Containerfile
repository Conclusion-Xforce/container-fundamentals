### This is a multi-stage build. Completing it is left as an exercise for the reader.
### Stage 1: Build go binary
### Stage 2: Build the final image

## Stage 1
# Use the most recent 'go-toolset' from 'registry.access.redhat.com/ubi9'. Make sure to name the stage so it can be referenced later
TODO

# Copy the go source code into the container
TODO

# Build the go binary. The binary output is stored in '/opt/app-root/fibonacci-golang'
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /opt/app-root/fibonacci-golang

# Stage 2
# Use the most recent version of 'ubi-micro' from 'registry.access.redhat.com/ubi9'
TODO

# Ensure the container can be run as a non-root user by copying the '/etc/passwd' file from the previous stage
TODO

# Copy the go binary from the builder stage
TODO

# Ensure the binary is executable
TODO

# Set the user to the 'default' user
TODO

# Ensure the container listens on port 8080
TODO

# Set the default command for the container
TODO
