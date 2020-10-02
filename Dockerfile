FROM golang:latest as builder

ENV APP_HOME /app
ENV APP_NAME myapp

ENV GOPATH $APP_HOME

RUN mkdir $APP_HOME

RUN go get -u gopkg.in/mgo.v2 \
    && go get -u github.com/gin-gonic/gin \
    && go get -u gopkg.in/check.v1 \
    && go get -u github.com/vjeantet/jodaTime

COPY ./src/ $APP_HOME

WORKDIR $APP_HOME

RUN CGO_ENABLED=0 go build -a -ldflags '-s' -o $APP_NAME

FROM scratch

ENV APP_HOME /app
ENV APP_NAME myapp

COPY --from=builder $APP_HOME/$APP_NAME /$APP_NAME
EXPOSE 8080
CMD ["./myapp"]
