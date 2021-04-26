# Dockerfile for building the application and creating the runtime container.
FROM golang:1.16 as builder

RUN go get github.com/revel/cmd/revel && \
  go get github.com/lib/pq
COPY ./ /accounting

RUN revel package -a /accounting -m prod

FROM debian

ENV DB_DRIVER=sqlite3
ENV DB_HOST=/tmp/testdb.db
ENV DB_NAME=accounting
ENV DB_USERNAME=
ENV DB_PASSWORD=
ENV PORT=80

COPY --from=builder /accounting/accounting.tar.gz /accounting/accounting.tar.gz
COPY ./entrypoint.sh /accounting/entrypoint.sh

RUN tar -xzf /accounting/accounting.tar.gz -C /accounting

CMD /accounting/entrypoint.sh
