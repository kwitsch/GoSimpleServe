FROM --platform=$BUILDPLATFORM ghcr.io/euantorano/zig:master AS zig-env
FROM --platform=$BUILDPLATFORM golang:1-alpine AS build
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

WORKDIR /go/src

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg \
    go mod download

COPY . .

COPY --from=zig-env /usr/local/bin/zig /usr/local/bin/zig
ENV PATH="/usr/local/bin/zig:${PATH}" \
    CC="zigcc" \
    CXX="zigcpp" \
    CGO_ENABLED=0 \
    GOOS="linux" \
    GOARCH=$TARGETARCH

RUN apk add --no-cache make

RUN --mount=type=bind,target=. \
    --mount=type=cache,target=/root/.cache/go-build \ 
    --mount=type=cache,target=/go/pkg \
    make build GOARM=${TARGETVARIANT##*v} && \
    chown 1000 /tmp/gss

FROM scratch

USER 1000
WORKDIR /app

COPY --from=build /tmp/gss /app/gss

ENV VERBOSE=False \
    ENDPOINT_FILES=False

ENTRYPOINT [ "/app/gss" ]