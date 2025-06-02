FROM alpine:3.22.0
WORKDIR /app

# TODO change me
COPY go-starter .

# TODO change me
CMD [ "./go-starter" ]