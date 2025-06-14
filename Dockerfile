# STAGE 1

# Pick a base image
FROM golang:latest AS build

# Configure target location in the container
WORKDIR /src/

# Copy source code from local machine to deployment machine and build executable
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/datatx


# STAGE 2

# Build two stage image
FROM scratch

# Container expose port
EXPOSE 8888 

# Finally get the exe to target from previous stage
COPY --from=build /bin/datatx /bin/datatx

# Start the exe
CMD [ "/bin/datatx"]