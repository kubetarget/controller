#!/bin/bash
set -x

# Add bin directory to PATH
echo 'export PATH=$PATH:$PWD/bin' >> ~/.bashrc
source ~/.bashrc

# Install kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x kubectl
sudo mv kubectl /usr/local/bin/kubectl

# Install kind
curl -Lo ./kind https://kind.sigs.k8s.io/dl/latest/kind-linux-amd64
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind

docker network create -d=bridge --subnet=172.19.0.0/24 kind

# Install tools
make tools

# Verify installation
go version
kind version
kubectl version
kustomize version
helm version
kubebuilder version
controller-gen version
envtest version
golangci-lint version
