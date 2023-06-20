FROM alpine:latest

WORKDIR /gopenfusion

# grab binary
COPY ./bin/server ./
RUN chmod +x ./server

CMD ["/gopenfusion/server"]