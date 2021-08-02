FROM golang:alpine

RUN mkdir /back

ADD . /back

WORKDIR /back

RUN GO BUILD -o main .

CMD ["/back/main"]

