FROM golang:1.20.4-alpine3.17

COPY ./ /tmp/app/

RUN apk add just

RUN (cd /tmp/app/; just build && mkdir /app && cp ./build/server /app/server) && rm -rf /tmp/app

WORKDIR /app/

EXPOSE 8088
CMD ["./server"]
