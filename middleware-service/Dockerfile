FROM golang:1.14.2

ENV APP_NAME myproject
ENV PORT 8000
ENV APP_ORIGIN_URL http://andersonastrigi.com.br:3000

COPY . /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN go get -u github.com/jinzhu/gorm
RUN go get ./
RUN go build -o ${APP_NAME}

CMD ./${APP_NAME}

EXPOSE ${PORT}