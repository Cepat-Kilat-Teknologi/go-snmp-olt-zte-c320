FROM golang:1.21-alpine as dev
RUN go install github.com/cosmtrek/air@latest
WORKDIR /app
COPY . /app/
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app ./cmd/api

FROM gcr.io/distroless/static-debian11 as prod
ENV APP_ENV=production
COPY --from=dev go/bin/app /
COPY --from=dev app/config/config-prod.yaml /config/config-prod.yaml
EXPOSE 8081
ENTRYPOINT ["/app"]