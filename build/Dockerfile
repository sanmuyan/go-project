FROM  golang:1.18 as builder
ENV GO111MODULE=on

WORKDIR /workspace

COPY ./ ./

RUN --mount=type=cache,target=/root/.cache \
    --mount=type=cache,target=/go \
    go build -o go-project ./cmd/go-project

FROM centos:7
ENV LANG=en_US.utf8

WORKDIR /workspace
COPY --from=builder /workspace /workspace

CMD ["/bin/bash", "-c", "go-project"]