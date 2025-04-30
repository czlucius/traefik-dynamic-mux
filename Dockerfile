# Build stage
FROM golang:alpine
WORKDIR /build
COPY . .
RUN go build -o /bin/traefik-dynamic-mux

FROM scratch
COPY --from=0 /bin/traefik-dynamic-mux /bin/traefik-dynamic-mux
CMD ["/bin/traefik-dynamic-mux"]
