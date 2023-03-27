FROM golang:1.20 AS builder

WORKDIR /go/src/github.com/luuisavelino/task-card-users/

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

COPY main.go ./

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /go/src/github.com/luuisavelino/task-card-users/app ./

CMD ["./app"]