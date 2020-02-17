#!/bin/bash

# This script will set up the core part for both master and worker node
echo Setting up Kubernetes Node...

# Update the apt package index
apt-get update

# Docker installation per guidance from docs.docker.com
# Install packages to allow apt to use a repository over HTTPS
apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

# Add Docker’s official GPG key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# Set up the stable repository
add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

# Install the latest version of Docker Engine - Community and containerd
apt-get install docker-ce docker-ce-cli containerd.io

# Enable Docker service
systemctl enable docker
systemctl start docker

# Kubernetes installation per guidance from kubernetes.io
# Update repo list with kubernetes tools and install the 3 universal tools
# all nodes are expected to have.
sudo apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl

# Mark up kublet, kubeadm and kubectl to prevent from auto updating
sudo apt-mark hold kubelet kubeadm kubectl

# Enable kublet service
systemctl enable kubelet
systemctl start kubelet