FROM golang:1.19-alpine3.16
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
EXPOSE 3000
ENTRYPOINT ["/app/main"]
CMD ["serve"]