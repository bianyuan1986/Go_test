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
import sys

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

    def setBinary(self, data, byteLen):
        if data <= 0:
            for i in range(0, byteLen):
		c = struct.pack("<B", 0)
            	self.payload += c 
            return

        hexStr = hex(data)[2:]
        #print "0x%s" % hexStr
       
        sLen = len(hexStr)
        i = 1
        byteArray = []
        for i in range(sLen-2, -1, -2):
            byte = int(hexStr[i], 16)*16+int(hexStr[i+1], 16)
            byteArray.append(byte)
        if i == 1:
            byte = int(hexStr[0], 16)
            byteArray.append(byte)

        for i in range(0, byteLen-len(byteArray)):
            byteArray.append(0x00)

        for i in range(0, len(byteArray)):
            c = struct.pack("<B", byteArray[i])
            self.payload += c 

    def setVersion(self, version):
        self.setBinary(version, 2)

    def setCodecType(self, codecType):
        self.setBinary(codecType, 1)

    def setMsgType(self, msgType):
        self.setBinary(msgType, 1)

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
    kvMap["header.accept"] = "*/*" 
    kvMap["header.content-type"] = "text/html" 
    kvMap["header.user_agent"] = "tiger-browser" 
    kvMap["header.referer"] = "www.fake.com" 
    kvMap["header.cookie"] = "arg1=tiger&arg2=tiger2" 
    #kvMap["first_pkt"] = "true" 
    kvMap["matched_host"] = "www.tiger.test.com" 
    kvMap["uid"] = "270828"
    kvMap["uri"] = "/test.php?alert()"
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
    print "headerLen:%d" % len(builder.string())
    for (k,v) in kvMap.items():
        builder.setKeyValue(k, v)
    pkt = builder.string()
    print "packetLen:%d" % len(pkt)
    DumpPacket(pkt, len(pkt))

    return pkt

def DumpPacket(data, dLen):
    start = 0
    sys.stdout.write("0x{0:04x} ".format(0))
    for i in range(0, dLen):
        if (i != 0) and (i % 16) == 0:
            sys.stdout.write("  ")
            for j in range(0, 16):
                digit = ord(data[start+j])
                if (digit >= 33) and (digit <= 126):
                   sys.stdout.write(data[start+j])
                else:
                   sys.stdout.write(".")
            print "\n"
            sys.stdout.write("0x{0:04x} ".format(i))
            start = i
        str = "{0:02x} ".format(ord(data[i]))  
        sys.stdout.write(str)

    i = dLen
    dLen = (dLen + 15) & (~15) 
    for j in range(i, dLen):
	sys.stdout.write("   ")
    sys.stdout.write("  ")
    for j in range(0, i-start):
        digit = ord(data[start+j])
        if (digit >= 33) and (digit <= 126):
           sys.stdout.write(data[start+j])
        else:
           sys.stdout.write(".")
    print "\n"

unix_sock_path = "/home/admin/warden/run/ngwaf.sock"

if __name__ == "__main__":
    pkt = constructPacket()
    s = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    try:
        s.connect(unix_sock_path)
        s.sendall(pkt)
        print "Warden response:\n%s" % s.recv(1024)
    finally:
        s.close()
