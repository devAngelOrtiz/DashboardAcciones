FROM golang

WORKDIR /api

COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download

RUN go install github.com/air-verse/air@latest

CMD ["air"]