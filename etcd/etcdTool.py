#!/usr/bin/python
# coding=utf-8

import etcd3
import etcd
import requests

class EtcdV3Instance:
    etcdConn = None
    host = ""
    port = 0

    def __init__(self, host, port):
        self.host = host
        self.port = port
        self.etcdConn = etcd3.client(host, port)

    def setValue(self, key, value):
        self.etcdConn.put(key, value)

    def getValue(self, key):
        return self.etcdConn.get(key)

    def deleteKey(self, key):
        self.etcdConn.delete(key)

    def getAllKeys(self):
        return self.etcdConn.get_all()

class EtcdV2Instance:
    etcdConn = None
    host = ""
    port = 0

    def __init__(self, host, port):
        self.host = host
        self.port = port
        self.etcdConn = etcd.Client(host, port)

    def setValue(self, key, value):
        self.etcdConn.write(key, value)

    def getValue(self, key):
        return self.etcdConn.read(key).value

    def deleteKey(self, key):
        self.etcdConn.delete(key)

class HttpParse:
    uri = ""

    def __init__(self, uri):
        self.uri = uri

if __name__ == "__main__":
    source = EtcdV2Instance("11.164.89.26", 4001)
    print source.getValue("/luffy")

    target = EtcdV3Instance("127.0.0.1", 2379)
    print target.getValue("tiger")
