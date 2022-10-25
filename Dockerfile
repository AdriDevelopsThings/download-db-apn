FROM golang:alpine as build
RUN apk --no-cache add ca-certificates
WORKDIR /build
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

WORKDIR /dist
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /build/main /dist/main


ENV TARGET_DIRECTORY=/target
VOLUME [ "/target" ]
CMD ["/dist/main", "-t", "/target"]