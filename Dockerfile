FROM golang:1.16-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o unstable-api

FROM scratch

COPY --from=build /app/unstable-api /
COPY --from=build /app/config.yaml /

ENTRYPOINT ["/unstable-api"]