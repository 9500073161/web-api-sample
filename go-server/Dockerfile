FROM golang:1.22

WORKDIR /WEB-API-SAMPLE/go-server/

COPY go.mod .
COPY main.go .

RUN go get
RUN go build -o /WEB-API-SAMPLE/go-server/build/myapp .

ENTRYPOINT [ "/WEB-API-SAMPLE/go-server/build/myapp" ]