#!/bin/bash

./etcd &
export ETCDCTL_API=3
./etcdctl put tiger 1986
./etcdctl get tiger
curl -v http://localhost:2379/v2/keys
