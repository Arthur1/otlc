FROM --platform=$BUILDPLATFORM golang:bookworm@sha256:00eccd446e023d3cd9566c25a6e6a02b90db3e1e0bbe26a48fc29cd96e800901 AS builder
ARG TARGETARCH
ARG VERSION=unknown

ENV GOTOOLCHAIN=auto
ENV CGO_ENABLED=0
ENV GOARCH=${TARGETARCH}

WORKDIR /opt/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN  go build -ldflags="-s -w -X github.com/Arthur1/otlc.Version=${VERSION}" -o otlc ./cmd/otlc

FROM gcr.io/distroless/base-debian12:nonroot@sha256:0a0dc2036b7c56d1a9b6b3eed67a974b6d5410187b88cbd6f1ef305697210ee2

COPY --from=builder --chown=nonroot:nonroot /opt/app/otlc /otlc

ENTRYPOINT ["/otlc"]
CMD ["-h"]
