FROM golang:1.17

WORKDIR /go/src/github.com/D-Undefined/hack-camp_vol13_server

COPY . .

RUN go get github.com/cosmtrek/air
RUN go mod download
