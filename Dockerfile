FROM golang:1.23-alpine

ADD . /home

WORKDIR /home

CMD ["go", "run", "main.go"]