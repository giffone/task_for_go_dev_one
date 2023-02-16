FROM golang:1.19-bullseye

ENV PACKAGE_PATH=/web
WORKDIR $PACKAGE_PATH

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build cmd/main.go

ENTRYPOINT ["./main", "-addr=:8080"]