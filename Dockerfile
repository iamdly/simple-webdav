FROM golang

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app ./...

VOLUME /data

EXPOSE 8602

ENV USERNAME user
ENV PASSWORD 123

CMD [ "app" ]