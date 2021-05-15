FROM golang:alpine as build
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine
COPY --from=build /go/bin/backend /
CMD ["./backend"]
