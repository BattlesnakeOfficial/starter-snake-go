FROM golang:1.18 as build

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/battlesnake ./...

FROM ubuntu
COPY --from=build /usr/local/bin/battlesnake /bin/
USER 65534
CMD ["battlesnake"]
