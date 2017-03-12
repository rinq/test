FROM scratch
ENTRYPOINT ["/app/bin/rinq-test-server"]

EXPOSE 8080

ENV GODEBUG           netdns=cgo
ENV RINQ_AMQP_DSN     ""

COPY artifacts/build/release/linux/amd64/rinq-test-server /app/bin/
