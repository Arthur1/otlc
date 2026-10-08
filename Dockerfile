FROM --platform=$BUILDPLATFORM golang:bookworm@sha256:5cf287a799e6b94384bad13d16b14904c531f51ba65792237e122ce42b392f61 AS builder
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
