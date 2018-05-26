FROM golang:alpine

RUN apk update && apk add git curl

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENV PORT 8981

EXPOSE 8981

CMD ["app"]
