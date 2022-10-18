FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY echo/*.go ./

RUN go build -o /echo

EXPOSE 8080

CMD [ "/echo" ] 
