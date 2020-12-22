#!/bin/bash
set -euox pipefail

# install golang
if [ ! -d /usr/local/go ]; then
    pushd /tmp
    wget -nv https://golang.org/dl/go1.15.6.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz
    rm go1.15.6.linux-amd64.tar.gz
    popd

    echo "export PATH=\$PATH:/usr/local/go/bin" > /etc/profile.d/usr-local-golang-path.sh
    echo "export CGO_ENABLED=0" > /etc/profile.d/disable-cgo.sh
fi

/usr/local/go/bin/go version
