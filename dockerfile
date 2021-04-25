FROM golang:1.16

ENV DB_DRIVER=sqlite3
ENV DB_HOST=/tmp/testdb.db
ENV DB_NAME=accounting
ENV DB_USERNAME=
ENV DB_PASSWORD=

RUN go get github.com/revel/cmd/revel && \
  go get github.com/lib/pq
COPY ./ /accounting

RUN revel build -a /accounting

EXPOSE 80

CMD revel run /accounting prod
