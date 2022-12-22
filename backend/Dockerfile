FROM golang:latest

WORKDIR /backend

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN make build

EXPOSE 5454

CMD exec ./backend