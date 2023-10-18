FROM golang:1.21-alpine as dev
ENV config=dev
RUN go install github.com/cosmtrek/air@latest
WORKDIR /app
COPY . /app/
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app ./cmd/api

FROM gcr.io/distroless/static-debian11 as prod
ENV config=prod
COPY --from=dev go/bin/app /
COPY --from=dev app/config/config-prod.yml /config/config-prod.yml
EXPOSE 8080
ENTRYPOINT ["/app"]