FROM --platform=$BUILDPLATFORM golang:bookworm@sha256:c4bc0741e3c79c0e2d47ca2505a06f5f2a44682ada94e1dba251a3854e60c2bd AS builder
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

FROM gcr.io/distroless/base-debian12:nonroot@sha256:fa5f94fa433728f8df3f63363ffc8dec4adcfb57e4d8c18b44bceccfea095ebc

COPY --from=builder --chown=nonroot:nonroot /opt/app/otlc /otlc

ENTRYPOINT ["/otlc"]
CMD ["-h"]
