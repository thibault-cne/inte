FROM golang:1.18.2

WORKDIR /backend

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

EXPOSE 5454

CMD [ "/docker-gs-ping" ]