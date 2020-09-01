package main

//#include <unistd.h>
//import "C"

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	_AT_CLKTCK = 17 // Frequency of times()

	_SYSTEM_CLK_TCK = 100

	uintSize uint = 32 << (^uint(0) >> 63)
)

func getclktck() int64 {
	buf, err := ioutil.ReadFile("/proc/self/auxv")
	if err == nil {
		pb := int(uintSize / 8)
		for i := 0; i < len(buf)-pb*2; i += pb * 2 {
			var tag, val uint
			switch uintSize {
			case 32:
				tag = uint(binary.LittleEndian.Uint32(buf[i:]))
				val = uint(binary.LittleEndian.Uint32(buf[i+pb:]))
			case 64:
				tag = uint(binary.LittleEndian.Uint64(buf[i:]))
				val = uint(binary.LittleEndian.Uint64(buf[i+pb:]))
			}

			switch tag {
			case _AT_CLKTCK:
				return int64(val)
			}
		}
	}

	return _SYSTEM_CLK_TCK
}

func readFile(name string) string {
	fi, err := os.Open(name)
	if err != nil {
		fmt.Println("Open file failed!")
		return "error"
	}
	defer fi.Close()
	content, err := ioutil.ReadAll(fi)
	return string(content)
}

func main() {
	statName := "/proc/"
	statName += os.Args[1]
	statName += "/stat"

	tickFre := getclktck()
	fmt.Printf("Tick frequent:%d\n", tickFre)

	var processTime [2]int64
	var idx int
	for i := 0; i < 60; i++ {

		content := readFile("/proc/uptime")
		//fmt.Printf("Read /proc/uptime\n%s\n", content)
		flds := strings.Fields(content)
		//fmt.Printf("Run time:%s(s)\n", flds[0])
		uptime, err := strconv.ParseFloat(flds[0], 0)
		if err == nil {
			fmt.Printf("Run time:%d(jiffies)\n", int64(uptime*100))
		}

		content = readFile(statName)
		//fmt.Printf("Read %s\n%s\n", statName, content)
		flds = strings.Fields(content)
		fmt.Printf("Userspace time:%s(jiffies)\n", flds[13])
		fmt.Printf("System time:%s(jiffies)\n", flds[14])
		fmt.Printf("Start time:%s(jiffies)\n", flds[21])
		fmt.Println()

		utime, err := strconv.ParseInt(flds[13], 10, 0)
		if err == nil {
			//fmt.Printf("Userspace time:%d(s)\n", utime)
		}
		stime, err := strconv.ParseInt(flds[14], 10, 0)
		if err == nil {
			//fmt.Printf("System time:%d(s)\n", stime)
		}
		_, err = strconv.ParseInt(flds[21], 10, 0)
		if err == nil {
			//fmt.Printf("Start time:%d(s)\n", btime)
		}

		processTime[idx] = utime + stime
		//total := int64(uptime*100) - btime
		//cpuUsage := float64(processTime) / float64(total)
		if i == 0 {
			idx = (idx + 1) & 0x1
			continue
		}
		fmt.Printf("%d %d\n", processTime[idx], processTime[(idx+1)&0x1])
		cpuUsage := float64(processTime[idx]-processTime[(idx+1)&0x1]) / float64(tickFre*10) * 100
		fmt.Printf("Cpu usage:%.1f%%\n", cpuUsage)
		idx = (idx + 1) & 0x1
		fmt.Printf("idx:%d\n", idx)
		time.Sleep(time.Duration(10) * time.Second)
	}
	//fmt.Printf("Get _SC_CLK_TCK %d", int(C.sysconf(C._SC_CLK_TCK)))
}
