FROM scratch

# grab binary
WORKDIR /gopenfusion
COPY --chmod=0755 ./bin/server ./

ENTRYPOINT [ "/gopenfusion/server" ]