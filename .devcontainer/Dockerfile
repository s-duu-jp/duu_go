# ベースとなるDockerイメージを指定
FROM mcr.microsoft.com/vscode/devcontainers/base:ubuntu

# ユーザー情報の設定
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# UID/GIDがデフォルト値でない場合にユーザーとグループを更新または作成
RUN if [ "$USER_UID" != "1000" ] || [ "$USER_GID" != "1000" ]; then \
    groupmod -o -g "$USER_GID" "$USERNAME" || addgroup -g "$USER_GID" "$USERNAME"; \
    usermod -o -u "$USER_UID" -g "$USER_GID" "$USERNAME" || adduser --uid "$USER_UID" --gid "$USER_GID" --disabled-password --gecos "" "$USERNAME"; \
    chown -R $USER_UID:$USER_GID /home/$USERNAME; \
    fi

# ワークスペースディレクトリを設定
WORKDIR /workspace

# Switch User
USER ${USERNAME}

# Go 設定
ARG GO_VERSION=1.21
ARG GOPLS_VERSION=0.15.3
ARG GODLV_VERSION=1.22.1
ENV PATH="${PATH}:/usr/lib/go-${GO_VERSION}/bin:/home/${USERNAME}/go/bin"

# Go インストール
RUN sudo apt update \
  && sudo apt install -y golang-${GO_VERSION}

# Go Language Serverのインストール
RUN go install golang.org/x/tools/gopls@v${GOPLS_VERSION}

# Go dlvのインストール
RUN go env -w GOPROXY=https://goproxy.io,direct \
  && go install github.com/go-delve/delve/cmd/dlv@v${GODLV_VERSION}

# Go AIR インストール
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Go Atlas のインストール
RUN curl -sSfL https://atlasgo.sh | sh -s -- -b $(go env GOPATH)/bin