FROM golang:latest as builder

RUN mkdir /app
ENV GOPATH /app

RUN go get -u gopkg.in/mgo.v2 \
    && go get -u github.com/gin-gonic/gin \ 
    && go get -u gopkg.in/check.v1 \
    && go get -u github.com/vjeantet/jodaTime \
    && go get -u github.com/swaggo/gin-swagger \
    && go get -u github.com/swaggo/gin-swagger/swaggerFiles

COPY ./src /app/src
WORKDIR /app/src
RUN CGO_ENABLED=0 go build -a -ldflags '-s' -o cruddy 

FROM scratch
COPY --from=builder /app/src/cruddy /cruddy
EXPOSE 8080 
CMD ["./cruddy"]


