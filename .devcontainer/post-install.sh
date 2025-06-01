#!/bin/bash

# Setup Docker group permissions
sudo usermod -aG docker $USER
sudo chmod 666 /var/run/docker.sock

# Add bin directory to PATH
echo "export PATH=$PATH:$PWD/bin" >> ~/.bashrc
source ~/.bashrc

# Install latest stable kubectl
echo "Installing latest stable kubectl..."
KUBECTL_VERSION=$(curl -sSL https://dl.k8s.io/release/stable.txt)
ARCH=$(uname -m)
case "$ARCH" in
  x86_64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH" && exit 1 ;;
esac
curl -LO "https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/$(uname | tr '[:upper:]' '[:lower:]')/${ARCH}/kubectl"
chmod +x kubectl
sudo mv kubectl /usr/local/bin/

# Install latest kind
echo "Installing latest kind..."
KIND_VERSION=$(curl -s https://api.github.com/repos/kubernetes-sigs/kind/releases/latest | grep tag_name | cut -d '"' -f 4)
curl -Lo kind "https://kind.sigs.k8s.io/dl/${KIND_VERSION}/kind-$(uname | tr '[:upper:]' '[:lower:]')-${ARCH}"
chmod +x kind
sudo mv kind /usr/local/bin/

docker network create -d=bridge --subnet=172.19.0.0/24 kind || true

# Install local tools
make tools

# Verify installation
go version
kind version
kubectl version --client=true