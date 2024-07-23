FROM golang:alpine AS build_base

RUN apk add --no-cache git

# Set Current Working Directory
WORKDIR /usr/neffable

# Populate module cache based
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY backend/ ./backend

WORKDIR /usr/neffable/backend
# Build the Go App
RUN go build -o /usr/neffable/backend/out/server .


# Start from fresh a smaller image
FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=build_base /usr/neffable/backend/out/server /backend/server

# Copy the .env file to the root directory of the container
COPY .env /neffable/.env

#  This container expose port 8080 to the outside world.
EXPOSE 8080

# Run the binary program produced by go install
CMD ["/backend/server"]