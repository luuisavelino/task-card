FROM golang:1.20 AS builder

WORKDIR /go/src/github.com/luuisavelino/task-card-cards/

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

COPY main.go ./

RUN CGO_ENABLED=1 go build -a -installsuffix cgo -o app .

CMD ["./app"]