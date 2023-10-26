FROM golang:1.21 as build
WORKDIR /src
COPY . .
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /svd cmd/svd/main.go

FROM scratch
LABEL org.opencontainers.image.source https://github.com/rtgnx/svctl
ENTRYPOINT ["/svd", "--addr", ":8080"]
COPY --from=build /svd /svd