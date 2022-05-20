FROM golang:1.18.2

WORKDIR /backend

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /docker-gs-ping -v

EXPOSE 5454

CMD exec ./main