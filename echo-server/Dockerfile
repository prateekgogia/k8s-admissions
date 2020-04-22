FROM golang:1.14 AS build

# Set the Current Working Directory inside the container
WORKDIR /root/echo-server
RUN apt-get install -y ca-certificates

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Install the package
RUN CGO_ENABLED=0 go install -v ./...
CMD ["echo-server"]

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/bin/echo-server /root/
COPY --from=build /root/echo-server/keys/webhook-server-tls.crt /etc/certs/cert.pem
COPY --from=build /root/echo-server/keys/webhook-server-tls.key /etc/certs/key.pem
CMD ["/root/echo-server"]
