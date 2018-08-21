FROM golang:alpine

RUN apk update && apk add git curl

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -ldflags \
  "-X main.CommitID=$(git rev-parse HEAD) \
  -X main.Branch=$(git branch | grep \* | cut -d ' ' -f2)" \
  -o /go/bin/app

ENV PORT 8981

EXPOSE 8981

CMD ["app"]
