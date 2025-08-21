# syntax=docker/dockerfile:1.7
FROM golang:1.22-alpine

ARG USERNAME=dev
ARG USER_UID=1000
ARG USER_GID=1000

# 必要なパッケージを追加
RUN apk add --no-cache \
    git \
    openssh-client \
    bash \
    make \
    curl \
    ca-certificates \
    build-base

# ユーザ作成（ホストとUID/GIDを合わせると便利）
RUN addgroup -g ${USER_GID} ${USERNAME} \
 && adduser -D -u ${USER_UID} -G ${USERNAME} ${USERNAME}

# 開発ユーザで作業
USER ${USERNAME}
WORKDIR /workspace

CMD ["bash"]
