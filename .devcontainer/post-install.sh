#!/bin/bash

# Setup Docker group permissions
sudo usermod -aG docker $USER
sudo chmod 666 /var/run/docker.sock

# Add bin directory to PATH
echo "export PATH=$PATH:$PWD/bin" >> ~/.bashrc
source ~/.bashrc

docker network create -d=bridge --subnet=172.19.0.0/24 kind || true

# Install local tools
make tools

# Verify installation
go version
kind version
kubectl version --client=true