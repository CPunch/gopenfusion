FROM alpine:latest

WORKDIR /gopenfusion

# grab binary
COPY ./bin/server ./
RUN chmod +x ./server

ENTRYPOINT [ "/bin/sh", "-l", "-c" ]
CMD ["/gopenfusion/server"]