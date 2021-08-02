FROM golang:1.16

RUN mkdir /back

ADD . /back

WORKDIR /back

RUN go build -o main 

EXPOSE 8080

CMD ["/back/main"]

