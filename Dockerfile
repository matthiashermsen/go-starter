FROM alpine:3.22.1
WORKDIR /app

# TODO change me
COPY go-starter .

# TODO change me
CMD [ "./go-starter" ]