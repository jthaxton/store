FROM golang:latest AS builder
ADD . /
WORKDIR /
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

FROM scratch
ENV MONGODB_USERNAME=root MONGODB_PASSWORD=example MONGODB_ENDPOINT=documents
COPY --from=builder /main ./
ENTRYPOINT ["./main"]
EXPOSE 8080
EXPOSE 8081