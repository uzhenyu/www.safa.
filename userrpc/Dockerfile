FROM golang:1.22.1

RUN mkdir /app

COPY ./ /app

WORKDIR /app

RUN go mdo tidy

RUN go bulid main.go

CMD ["/main"]
