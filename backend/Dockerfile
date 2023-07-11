FROM golang:1.20
WORKDIR /backend
COPY . .
RUN go get ./...
RUN go build
EXPOSE 8888
CMD ["./backend"]