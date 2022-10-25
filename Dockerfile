FROM golang
WORKDIR /build
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

WORKDIR /dist
RUN cp /build/main /dist/main


ENV TARGET_DIRECTORY=/target
VOLUME [ "/target" ]
CMD ["/dist/main"]