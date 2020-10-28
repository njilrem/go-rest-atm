FROM golang:1.15 AS build-env

ADD . /atm
WORKDIR /atm

RUN git config --global url."https://f7e15181a9ad18f4fa36edfd11dd2cf8a9c39664:@github.com/".insteadOf "https://github.com/"

RUN go build -o /server

FROM debian:buster

EXPOSE 8080

WORKDIR /
COPY --from=build-env /server /

CMD ["/server"]