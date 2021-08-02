FROM golang:alpine

RUN mkdir /back

ADD . /back

WORKDIR /back

RUN GO BUILD -o main .

EXPOSE 8080

CMD ["/back/main"]

