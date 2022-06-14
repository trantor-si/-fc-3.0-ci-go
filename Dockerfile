FROM golang:latest AS go-builder

WORKDIR /app
COPY . .

RUN go build -o math

FROM scratch

COPY --from=go-builder /app/math .

CMD [ "./math" ]