FROM golang:1.15 as build-stage

RUN mkdir -p /mocker
WORKDIR /mocker

ADD go.* /mocker/
RUN go mod download

ADD . /mocker/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mocker cmd/mocker/main.go


FROM alpine:3

ADD mocker /usr/local/bin/
COPY --from=build-stage /mocker/mocker /usr/local/bin/
RUN mkdir -p /mocker
WORKDIR /mocker

CMD [ "mocker" ]
