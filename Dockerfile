FROM golang:1.10-alpine as builder
WORKDIR /go/src/github.com/brunsgaard/img2tiff
COPY . .
RUN apk update; apk add git
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build cmd/img2tiff/img2tiff.go
RUN chmod a+x /go/src/github.com/brunsgaard/img2tiff/img2tiff

#FROM alpine
#COPY --from=builder /go/src/github.com/brunsgaard/img2tiff/img2tiff /usr/bin/img2tiff
#CMD ["/usr/bin/img2tiff"]
