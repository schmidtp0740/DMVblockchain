# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.10.0-stretch
WORKDIR /usr/src/app
COPY . .
RUN go get github.com/gorilla/mux 
RUN go build main.go
CMD [ "./main" ]