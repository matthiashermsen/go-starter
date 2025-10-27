FROM alpine:3.22.2
WORKDIR /app

# TODO change me
COPY go-starter .

# TODO change me
CMD [ "./go-starter" ]