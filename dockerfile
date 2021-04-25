FROM golang:1.16

RUN go get github.com/revel/cmd/revel && \
  go get github.com/lib/pq
COPY ./ /accounting

EXPOSE 8888

ENTRYPOINT revel run /accounting
CMD prod
