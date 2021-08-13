FROM golang:1.16

RUN mkdir /back

WORKDIR /back

COPY . .

RUN go build -o main 

EXPOSE 8080

CMD ["/back/main"]

# CMD ["go", "run", "server.go"]
