FROM scratch
COPY ./bin/httpecho /opt/httpecho
ENTRYPOINT ["/opt/httpecho"]