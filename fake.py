#!/usr/bin/python
# coding=utf-8

"""
 88 const (
 89     HeaderLen    = 16
 90     VersionLen   = 2
 91     CodecTypeLen = 1
 92     MsgTypeLen   = 1
 93     DataLen      = 4 // 不包括头的长度
 94     RequextIdx   = 4
 95     ReservedLen  = 4
 96
 97     Version1              = uint16(1000)
 98     Version2              = uint16(2000)
 99     MsgTypeReq    MsgType = 0x00
100     MsgTypeStatus MsgType = 0x01
101     MsgTypeLog    MsgType = 0x02
102 )

104 type MsgHeader struct {
105     Version    uint16 /*版本*/
106     CodecType  uint8  /*消息编码类型，根据该值获取对应的解析器*/
107     MsgType    uint8  /*消息类型*/
108     DataLen    uint32 /*body数据长度*/
109     RequestIdx uint32 /*请求索引号*/
110     Reserved   uint32 // reserved
111 }
"""

import struct
import socket

#(1,0,0,0)-->(b'\x01\x00\x00\x00')
def BinaryDataToByteStr(data):
    content = ""
    for i in range(len(data)):
        content += struct.pack('<B', data[i])

    return content

#(b'\x01\x00\x00\x00')-->(1,0,0,0)
def ByteStrToBinaryData(data):
    dLen = len(data)
    fmtStr = "<"
    fmtStr += dLen * 'B'
    return struct.unpack(fmtStr, data), dLen

def BinaryDataToUint16Str(data):
    content = ""
    for i in range(len(data)):
        content += struct.pack('<H', data[i])

    return content

def BinaryDataToUint32Str(data):
    content = ""
    for i in range(len(data)):
        content += struct.pack('<I', data[i])

    return content

#52-->'4' 0x34-->'4'
def IntToChar(number):
    return chr(number)

#52-->'0x34'
def IntToHex(number):
    return hex(number)

#0x34-->'34'
def HexToHexStr(number):
    return hex(number)[2:]

#'a'-->97
def AsciiToInt(char):
    return ord(char)

#'0x34'-->0x34
def HexStrToHex(number):
    sLen = len(number)
    total = 0
    for i in range(sLen):
        total = total*16+int(number[i])

    return total

class PacketBuilder:
    payload = ""
    header = ""
    body = ""

    def setBinary(self, data, byteLen):
        byteArray = []
        if data <= 0:
            for i in range(0, byteLen):
                byteArray.append(0x00)
                return

        hexStr = hex(data)[2:]
        print "0x%s" % hexStr
       
        sLen = len(hexStr)
        i = 1
        for i in range(sLen-2, -1, -2):
            byte = int(hexStr[i], 16)*16+int(hexStr[i+1], 16)
            byteArray.append(byte)
            print "byte:0x%x" % byte
        if i == 1:
            byte = int(hexStr[0], 16)
            byteArray.append(byte)
            print "byte:0x%x" % byte

        print range(byteLen-len(byteArray), byteLen)
        for i in range(0, byteLen-len(byteArray)):
            byteArray.append(0x00)
        print byteArray

        for i in range(0, len(byteArray)):
            self.payload += struct.pack("<B", byteArray[i])

    def setVersion(self, version):
        self.setBinary(version, 2)

    def setCodecType(self, data):
        self.setBinary(0, 1)

    def setMsgType(self, data):
        self.setBinary(0, 1)

    def setDataLen(self, dLen):
        self.setBinary(dLen, 4)

    def setRequestIdx(self, idx):
        self.setBinary(idx, 4)

    def setPadding(self, padLen):
        self.setBinary(0, padLen)

    def setKeyValue(self, key, value):
        kLen = len(key)
        vLen = len(value)
        self.setBinary(kLen, 4)
        self.payload += key
        self.setBinary(vLen, 4)
        self.payload += value
        
        return 4 + 4 + kLen + vLen

    def string(self):
        return self.payload

def constructPacket():
    kvMap = {}
    kvMap["A"*0x31] = "B"*0x32
    bodyLen = 0
    for (k,v) in kvMap.items():
        bodyLen += 8
        bodyLen += len(k)
        bodyLen += len(v)
    print "bodyLen:%d" % bodyLen

    builder = PacketBuilder()
    builder.setVersion(1000)
    builder.setCodecType(0)
    builder.setMsgType(0)
    builder.setDataLen(bodyLen)
    builder.setRequestIdx(2)
    builder.setPadding(4)
    for (k,v) in kvMap.items():
        builder.setKeyValue(k, v)
    pkt = builder.string()
    print pkt
    return pkt

unix_sock_path = "/home/admin/warden/run/ngwaf.sock"

if __name__ == "__main__":
    pkt = constructPacket()
    s = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    try:
        s.connect(unix_sock_path)
        s.sendall(pkt)
        print s.recv()
    finally:
        s.close()
