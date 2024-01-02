FROM golang:1.20.4-alpine3.17

WORKDIR /app/

COPY ./ /tmp/app/

RUN apk add just
RUN (cd /tmp/app/; just build && cp ./build/server-example /app/server) && rm -rf /tmp/app
RUN apk del just

EXPOSE 8088
CMD ["./server"]
