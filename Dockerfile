FROM golang:1.14.0-alpine
RUN mkdir proc-manager
WORKDIR proc-manager
COPY . .
RUN apk add make
RUN make build
CMD ["./bin/process-manager"]
