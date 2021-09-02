# [Choice] https://github.com/KEINOS/VSCode-Dev-Container-Go/pkgs/container/vscode-dev-container-go
ARG VARIANT="latest"
FROM ghcr.io/keinos/vscode-dev-container-go:${VARIANT}

USER root

# [Optional] Uncomment this section to install additional OS packages.
# RUN apk add --no-cache <your-package-list-here>

USER vscode

# [Optional] Uncomment the next line to use go get to install anything else you need
RUN cd /tmp && go install github.com/KEINOS/go-prettybench@latest
