FROM golang:1.22-alpine

# install dependencies
RUN apk update && apk add postgresql-client curl tzdata

# set timezone
ENV TZ=Europe/Moscow

# install goose
RUN curl -fsSL \
        https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
        sh -s v3.5.0

WORKDIR /go/src/app

COPY . .

# make start.sh executable
RUN chmod +x start.sh

# build go app
RUN go mod download && go build -o mailbomber ./cmd/apiserver

CMD ["./mailbomber"]