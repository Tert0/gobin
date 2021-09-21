FROM alpine as BUILDER

WORKDIR /go/src/

RUN apk update && apk add --update --no-cache go gcc g++

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o /go/bin/main .

FROM alpine:latest

WORKDIR /go/bin

COPY --from=builder /go/bin/main .

CMD ["/go/bin/main"]