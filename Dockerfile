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
# (disable bellow line for debugging, enable this for production)
# FROM scratch   

# Container exposed port
EXPOSE 8888 

# Finally get the exe to target from previous stage
# (disable bellow line for debugging, enable this for production)
# COPY --from=build /bin/datatx /bin/datatx

# Start the exe
CMD [ "/bin/datatx"]