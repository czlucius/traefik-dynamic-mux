# Build stage
FROM --platform=$BUILDPLATFORM golang:alpine
WORKDIR /build
COPY . .
ARG TARGETOS TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /bin/traefik-dynamic-mux
# Deploy stage
FROM alpine
COPY --from=0 /bin/traefik-dynamic-mux /bin/traefik-dynamic-mux
CMD ["/bin/traefik-dynamic-mux"]
