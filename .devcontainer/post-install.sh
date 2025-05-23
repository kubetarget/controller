#!/bin/bash
set -x

# Add bin directory to PATH
echo 'export PATH=$PATH:$PWD/bin' >> ~/.bashrc
source ~/.bashrc

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
