FROM golang:1.17 as build

WORKDIR /go/src/gindemo

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o gindemo .


FROM alpine
WORKDIR /usr/local/demo
COPY --from=build /go/src/gindemo/gindemo .
COPY config.yaml config.yaml
EXPOSE 80
CMD ["./gindemo"]