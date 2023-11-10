FROM golang:1.21 as build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/server

FROM build as production

COPY --from=build /usr/local/bin/app .

ENV WEB_SERVER_PORT=8001
ENV DB_DRIVER=sqlite
ENV DB_NAME="./sqlite.db"
EXPOSE $WEB_SERVER_PORT

CMD ["app"]