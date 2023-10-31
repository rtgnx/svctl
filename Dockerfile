ARG TARGET=svd
FROM node:18.13.0 as frontend
WORKDIR /src
COPY . .
WORKDIR /src/internal/srv/web/vue
RUN npm install
RUN npm run build

FROM golang:1.21 as build
WORKDIR /src
COPY . .
COPY --from=frontend /src/internal/srv/web/vue/ /src/internal/srv/web/vue/
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /$TARGET cmd/$TARGET/main.go

FROM scratch
LABEL org.opencontainers.image.source https://github.com/rtgnx/svctl
ENTRYPOINT ["/svd", "--addr", ":8080"]
COPY --from=build /$TARGET /$TARGET