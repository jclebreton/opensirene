# Multi-stage : Part 1
# Build Stage
FROM golang:alpine AS build

RUN mkdir -p $GOPATH/src/github.com/jclebreton/opensirene
RUN mkdir /output
ADD . $GOPATH/src/github.com/jclebreton/opensirene
WORKDIR $GOPATH/src/github.com/jclebreton/opensirene
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o opensirene
RUN cp opensirene /output/

# Multi-stage : Part 2
# Final Stage
FROM alpine

# Base packages
RUN apk update
RUN apk upgrade
RUN apk add ca-certificates && update-ca-certificates
RUN apk add --update tzdata
RUN rm -rf /var/cache/apk/*

# Copy binary
COPY --from=build /output/opensirene /home/

# Define timezone
ENV TZ=Europe/Paris

# Define the ENTRYPOINT
WORKDIR /home
ENTRYPOINT ./opensirene

# Document that the service listens on port 8080.
EXPOSE 8080