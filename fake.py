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

#for rule 111090
def constructHeaderPacket(ruleId):
    kvMap = {}
    kvMap["header.accept"] = "*/*" 
    kvMap["header.content-type"] = "text/html" 
    kvMap["header.user_agent"] = "tiger-browser" 
    kvMap["header.referer"] = "www.fake.com" 
    kvMap["header.cookie"] = "arg1=tiger&arg2=tiger2" 
    kvMap["first_pkt"] = "true" 
    kvMap["matched_host"] = "www.tiger.test.com" 
    kvMap["uid"] = "270828"
    if ruleId == 111044:
        kvMap["uri"] = "courseOpenId=8w52as6seodlxrtqdeyrbw&quesType=1&dataJson=%5B%7B%22SortOrder%22%3A0%2C%22Content%22%3A%22%3Cspan style%3D%5C%22color%3A rgb(0%2C     0%2C 0)%3B font-family%3A %E5%AE%8B%E4%BD%93%3B font-size%3A 14px%3B white-space%3A pre-wrap%3B background-color%3A rgb(255%2C 255%2C 255)%3B%5C%22%3Ea%3D1 b%    3D5 c%3D1%3C%2Fspan%3E%22%2C%22IsAnswer%22%3Afalse%7D%2C%7B%22SortOrder%22%3A1%2C%22Content%22%3A%22%3Cdiv%3E%3Cspan style%3D%5C%22background-color%3A rgb(255    %2C 255%2C 255)%3B color%3A rgb(0%2C 0%2C 0)%3B font-family%3A %E5%AE%8B%E4%BD%93%3B font-size%3A 14px%3B white-space%3A pre-wrap%3B%5C%22%3Ea%3D5 b%3D3 c%3D1    %3C%2Fspan%3E%3Cbr%3E%3C%2Fdiv%3E%22%2C%22IsAnswer%22%3Afalse%7D%2C%7B%22SortOrder%22%3A2%2C%22Content%22%3A%22%3Cspan style%3D%5C%22color%3A rgb(0%2C 0%2C 0)    %3B font-family%3A %E5%AE%8B%E4%BD%93%3B font-size%3A 14px%3B white-space%3A pre-wrap%3B background-color%3A rgb(255%2C 255%2C 255)%3B%5C%22%3Ea%3D1 b%3D3 c%3    D1%3C%2Fspan%3E%22%2C%22IsAnswer%22%3Atrue%7D%2C%7B%22SortOrder%22%3A3%2C%22Content%22%3A%22%3Cdiv%3E%3Cspan style%3D%5C%22background-color%3A rgb(255%2C 255%    2C 255)%3B color%3A rgb(0%2C 0%2C 0)%3B font-family%3A %E5%AE%8B%E4%BD%93%3B font-size%3A 14px%3B white-space%3A pre-wrap%3B%5C%22%3Ea%3D5 b%3D3 c%3D5%3C%2Fsp    an%3E%3Cbr%3E%3C%2Fdiv%3E%22%2C%22IsAnswer%22%3Afalse%7D%5D&questionId=&pageSize=10&page=1&DefficultyLevel=1&knowledgeTitle=%E9%80%89%E6%8B%A9%E7%BB%93%E6%9E%    84%E8%87%AA%E6%B5%8B&knowledgeIds=c7pqaiqsjitkyiy7p2edmw&Title=%3Cp%3E%E5%81%87%E5%AE%9A%E6%89%80%E6%9C%89%E5%8F%98%E9%87%8F%E5%9D%87%E5%B7%B2%E6%AD%A3%E7%A1%    AE%E8%AF%B4%E6%98%8E%EF%BC%8C%E4%BB%A5%E4%B8%8B%E7%A8%8B%E5%BA%8F%E6%AE%B5%E8%BF%90%E8%A1%8C%E5%90%8E%E7%9A%84%E8%BE%93%E5%87%BA%E7%BB%93%E6%9E%9C%E6%98%AF%26    nbsp%3B%26nbsp%3B%26nbsp%3B%26nbsp%3B%26nbsp%3B %E3%80%82%3C%2Fp%3E%3Cp%3Eint a%3D1%2Cb%3D5%2Cc%3D3%3B%3C%2Fp%3E%3Cp%3Eif(a%26gt%3Bb)%3C%2Fp%3E%3Cp%3Ea%3Db%3B    %3C%2Fp%3E%3Cp%3Eb%3Dc%3B%3C%2Fp%3E%3Cp%3Ec%3Da%3B%3C%2Fp%3E%3Cp%3Eprintf(%26quot%3Ba%3D%25d b%3D%25d c%3D%25d%5Cn%26quot%3B%2Ca%2Cb%2Cc)%3B%3C%2Fp%3E%3Cp%3E%    3Cbr%2F%3E%3C%2Fp%3E%3Cp%3E%3Cbr%2F%3E%3C%2Fp%3E&answer=2&ResultAnalysis="
    else:
        kvMap["uri"] = "/test.php?alertd()"
    if ruleId == 111090:
        kvMap["body"] = "search=select trunc( categoryid/pkg_roles.Fgetmancat(2)) as CATEGORYID, sum(nvl(a.salevalue,0)) as salevalue, sum(nvl(a.netsalevalue,0)) as netsalevalue, sum(nvl(a.salevalueratio,0)) as salevalueratio, sum(nvl(a.costvalue,0)) as costvalue, sum(nvl(a.netcostvalue,0)) as netcostvalue, sum(nvl(a.grossprofit,0)) as grossprofit, sum(nvl(a.grossprofit,0))/f0tonull( sum(nvl(a.salevalue,0)))*100 as grossmargin, sum(nvl(a.grossprofitratio,0)) as grossprofitratio,salechannel from TMP_RPT_M85151006 a where 1=1 group by salechannel, trunc( categoryid/pkg_roles.Fgetmancat(2)) order by trunc( categoryid/pkg_roles.Fgetmancat(2))"
    elif ruleId == 111044:
        kvMap["body"] = "courseOpenId=8w52as6seodlxrtqdeyrbw&quesType=1&dataJson=%5B%7B%22SortOrder%22%3A0%2C%22Content%22%3A%22%3Cspan style%3D%5C%22color%3A rgb(0%2C     0%2C 0)%3B font-family%3A %E5%AE%8B%E4%BD%93%3B font-size%3A 14px%3B white-space%3A pre-wrap%3B background-color%3A rgb(255%2C 255%2C 255)%3B%5C%22%3Ea%3D1 b%    3D5 c%3D1%3C%2Fspan%3E%22%2C%22IsAnswer%22%3Afalse%7D%2C%7B%22SortOrder%22%3A1%2C%22Content%22%3A%22%3Cdiv%3E%3Cspan style%3D%5C%22background-color%3A rgb(255    %2C 255%2C 255)%3B color%3A rgb(0%2C 0%2C 0)%3B font-family%3A %E5%AE%8B%E4%BD%93%3B font-size%3A 14px%3B white-space%3A pre-wrap%3B%5C%22%3Ea%3D5 b%3D3 c%3D1    %3C%2Fspan%3E%3Cbr%3E%3C%2Fdiv%3E%22%2C%22IsAnswer%22%3Afalse%7D%2C%7B%22SortOrder%22%3A2%2C%22Content%22%3A%22%3Cspan style%3D%5C%22color%3A rgb(0%2C 0%2C 0)    %3B font-family%3A %E5%AE%8B%E4%BD%93%3B font-size%3A 14px%3B white-space%3A pre-wrap%3B background-color%3A rgb(255%2C 255%2C 255)%3B%5C%22%3Ea%3D1 b%3D3 c%3    D1%3C%2Fspan%3E%22%2C%22IsAnswer%22%3Atrue%7D%2C%7B%22SortOrder%22%3A3%2C%22Content%22%3A%22%3Cdiv%3E%3Cspan style%3D%5C%22background-color%3A rgb(255%2C 255%    2C 255)%3B color%3A rgb(0%2C 0%2C 0)%3B font-family%3A %E5%AE%8B%E4%BD%93%3B font-size%3A 14px%3B white-space%3A pre-wrap%3B%5C%22%3Ea%3D5 b%3D3 c%3D5%3C%2Fsp    an%3E%3Cbr%3E%3C%2Fdiv%3E%22%2C%22IsAnswer%22%3Afalse%7D%5D&questionId=&pageSize=10&page=1&DefficultyLevel=1&knowledgeTitle=%E9%80%89%E6%8B%A9%E7%BB%93%E6%9E%    84%E8%87%AA%E6%B5%8B&knowledgeIds=c7pqaiqsjitkyiy7p2edmw&Title=%3Cp%3E%E5%81%87%E5%AE%9A%E6%89%80%E6%9C%89%E5%8F%98%E9%87%8F%E5%9D%87%E5%B7%B2%E6%AD%A3%E7%A1%    AE%E8%AF%B4%E6%98%8E%EF%BC%8C%E4%BB%A5%E4%B8%8B%E7%A8%8B%E5%BA%8F%E6%AE%B5%E8%BF%90%E8%A1%8C%E5%90%8E%E7%9A%84%E8%BE%93%E5%87%BA%E7%BB%93%E6%9E%9C%E6%98%AF%26    nbsp%3B%26nbsp%3B%26nbsp%3B%26nbsp%3B%26nbsp%3B %E3%80%82%3C%2Fp%3E%3Cp%3Eint a%3D1%2Cb%3D5%2Cc%3D3%3B%3C%2Fp%3E%3Cp%3Eif(a%26gt%3Bb)%3C%2Fp%3E%3Cp%3Ea%3Db%3B    %3C%2Fp%3E%3Cp%3Eb%3Dc%3B%3C%2Fp%3E%3Cp%3Ec%3Da%3B%3C%2Fp%3E%3Cp%3Eprintf(%26quot%3Ba%3D%25d b%3D%25d c%3D%25d%5Cn%26quot%3B%2Ca%2Cb%2Cc)%3B%3C%2Fp%3E%3Cp%3E%    3Cbr%2F%3E%3C%2Fp%3E%3Cp%3E%3Cbr%2F%3E%3C%2Fp%3E&answer=2&ResultAnalysis="
    else:
        kvMap["body"] = "<?xml version=\"1.0\" encoding=\"UTF-8\"?><note><to>Tove</to><from>Jani</from><heading>Reminder</heading><body>Don't forget me this weekend!</body></note>"
        #kvMap["body"] = "<html><head><title>tiger</title></head><body></body></html>"

    return constructPacket(kvMap)

def constructPacket(kvMap):
    bodyLen = 0
    for (k,v) in kvMap.items():
        bodyLen += 4
        bodyLen += len(k)
        bodyLen += 4
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
    ids = [0, 111044, 111090]
    for i in range(0, len(ids)):
        pkt = constructHeaderPacket(ids[i])
        s = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
        try:
            s.connect(unix_sock_path)
            s.sendall(pkt)
            print "Warden response:\n%s" % s.recv(1024)
        finally:
            s.close()
