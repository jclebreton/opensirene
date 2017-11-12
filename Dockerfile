FROM golang:1.9 AS build
RUN mkdir -p $GOPATH/src/github.com/jclebreton/opensirene
ADD . $GOPATH/src/github.com/jclebreton/opensirene
WORKDIR $GOPATH/src/github.com/jclebreton/opensirene
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build -o opensirene
RUN cp opensirene /

FROM golang:1.9
COPY --from=build /opensirene /usr/bin/
ENTRYPOINT ["/usr/bin/opensirene"]
EXPOSE 8080