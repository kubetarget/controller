ARG GO_VERSION=1.23

FROM mcr.microsoft.com/devcontainers/go:${GO_VERSION}

# Install kubectl
RUN curl -LO "https://dl.k8s.io/release/v1.33.1/bin/linux/amd64/kubectl" && \
    chmod +x kubectl && \
    mv kubectl /usr/local/bin/

# Install kind
RUN curl -Lo ./kind https://kind.sigs.k8s.io/dl/latest/kind-linux-amd64 && \
    chmod +x ./kind && \
    mv ./kind /usr/local/bin/kind
