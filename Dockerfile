FROM alpine:3.21.3
WORKDIR /app

# TODO change me
COPY go-starter .

# TODO change me
CMD [ "./go-starter" ]