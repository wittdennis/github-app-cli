FROM golang:1.25.4 as builder
ARG VERSION
ARG GIT_COMMIT

ENV CGO_ENABLED=0

COPY ./src /app

WORKDIR /app

RUN go build -ldflags "-s -w -X github-app/cmd/version.appVersion=$VERSION -X github-app/cmd/version.gitCommit=${GIT_COMMIT} -X github-app/cmd/version.buildDate=$(date +"%Y-%m-%dT%H:%M:%SZ")" -trimpath -o 'bin/' ./...

FROM alpine:3.22.2

COPY --from=builder /app/bin/github-app /usr/local/bin/github-app

# hadolint ignore=DL3018
RUN set -eux; \
    apk add --no-cache "dumb-init";

RUN addgroup -S 1000 && adduser -S 1000 -G 1000

WORKDIR /workdir

RUN chown -R 1000:1000 /workdir

USER 1000

ENTRYPOINT ["dumb-init", "github-app"]
CMD [""]
