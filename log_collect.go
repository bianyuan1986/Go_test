package log_collect

import (
	"encoding/binary"
	"fmt"
	"github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	_AT_CLKTCK = 17 // Frequency of times()

	_SYSTEM_CLK_TCK = 100

	uintSize uint = 32 << (^uint(0) >> 63)

	_FIELD_UTIME_IDX = 13
	_FIELD_STIME_IDX = 14
)

type FieldOps func(fieldName *map[string]bool, data interface{}) map[string]string

type FieldOpT struct {
	ops  FieldOps
	data interface{}
}

type Config struct {
	Enable          bool   `toml:"enable"`
	Endpoint        string `toml:"endpoint"`
	AccessKeyID     string `toml:"access_key_id"`
	AccessKeySecret string `toml:"access_key_secret"`
	LogFileName     string `toml:"log_file_name"`
	LogCompress     bool   `toml:"log_compress"`
	Retries         int    `toml:"retries"`
	ProjectName     string `toml:"project_name"`
	LogStoreName    string `toml:"log_store_name"`
	MaxIoWorkerCnt  int64  `toml:"max_ioworker_cnt"`
	ReportInterval  int64  `toml:"report_interval"`
}

type LogModule struct {
	fields sync.Map
}

var logModuleInit bool
var logModule *LogModule

var tickFreq int64
var processTime [2]int64
var idx int
var interval int64

const (
	ReportInterval = 10
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

func readFile(name string) (string, error) {
	fi, err := os.Open(name)
	if err != nil {
		fmt.Println("Open file failed!")
		return "", err
	}
	defer fi.Close()
	content, err := ioutil.ReadAll(fi)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func New() *LogModule {
	return &LogModule{fields: sync.Map{}}
}

func init() {
	if logModuleInit == false {
		logModule = New()
	}
	logModuleInit = true
}

func getLogModule() *LogModule {
	return logModule
}

func RegisterField(name string, ops FieldOps, data interface{}) {
	m := getLogModule()
	if m == nil {
		return
	}
	v := FieldOpT{ops: ops, data: data}
	fMap := make(map[string]bool)
	fMap[name] = true
	m.fields.Store(&fMap, v)
}

func RegisterMultiField(names map[string]bool, ops FieldOps, data interface{}) {
	m := getLogModule()
	if m == nil {
		return
	}
	v := FieldOpT{ops: ops, data: data}
	m.fields.Store(&names, v)
}

func (m *LogModule) generateLog(logTime uint32) *sls.Log {
	content := []*sls.LogContent{}
	m.fields.Range(func(k, v interface{}) bool {
		opT := v.(FieldOpT)
		kvPairs := opT.ops(k.(*map[string]bool), opT.data)
		for field, value := range kvPairs {
			content = append(content, &sls.LogContent{
				Key:   proto.String(field),
				Value: proto.String(value),
			})
		}

		return true
	})

	return &sls.Log{
		Time:     proto.Uint32(logTime),
		Contents: content,
	}
}

func newProducer(cfg *Config) *producer.Producer {
	producerConfig := producer.GetDefaultProducerConfig()

	producerConfig.Endpoint = cfg.Endpoint
	producerConfig.AccessKeyID = cfg.AccessKeyID
	producerConfig.AccessKeySecret = cfg.AccessKeySecret
	producerConfig.LogFileName = cfg.LogFileName

	if cfg.Retries > 0 && cfg.Retries <= 5 {
		producerConfig.Retries = cfg.Retries
	}
	if cfg.MaxIoWorkerCnt > 0 && cfg.MaxIoWorkerCnt <= 100 {
		producerConfig.MaxIoWorkerCount = cfg.MaxIoWorkerCnt
	}
	producerInstance := producer.InitProducer(producerConfig)

	return producerInstance
}

func checkProjectAndLogStore(cfg *Config) bool {
	client := sls.CreateNormalInterface(cfg.Endpoint, cfg.AccessKeyID, cfg.AccessKeySecret, "")
	_, err := client.GetProject(cfg.ProjectName)
	if err != nil {
		e := err.(*sls.Error)
		fmt.Printf("Get Project %s failed! HTTPCode:%d Code:%s Message:%s\n", cfg.ProjectName, e.HTTPCode, e.Code, e.Message)
		return false
	}

	_, err = client.GetLogStore(cfg.ProjectName, cfg.LogStoreName)
	if err != nil {
		e := err.(*sls.Error)
		fmt.Printf("Get LogStore %s failed! HTTPCode:%d Code:%s Message:%s\n", cfg.LogStoreName, e.HTTPCode, e.Code, e.Message)
		return false
	}

	return true
}

func GetGoroutineCount(name string) string {
	return strconv.Itoa(runtime.NumGoroutine())
}

func GetCpuUsage(name string) string {
	content, err := readFile("/proc/self/stat")
	if err != nil {
		return "0"
	}
	flds := strings.Fields(content)
	if len(flds) <= _FIELD_STIME_IDX {
		return "0"
	}
	utime, err := strconv.ParseInt(flds[_FIELD_UTIME_IDX], 10, 0)
	stime, err := strconv.ParseInt(flds[_FIELD_STIME_IDX], 10, 0)
	processTime[idx] = utime + stime
	previous := (idx + 1) & 0x1
	if processTime[previous] <= 0 {
		idx = previous
		return "0"
	}
	cpuUsage := float64(processTime[idx]-processTime[previous]) / float64(tickFreq*interval) * 100
	idx = previous

	return fmt.Sprintf("%.1f%%", cpuUsage)
}

func GetMemoryUsage(name string) string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return strconv.FormatInt(int64(m.Sys), 10)
}

func LogRoutine(cfg *Config) {
	if cfg.Enable == false {
		return
	}
	ready := checkProjectAndLogStore(cfg)
	if ready == false {
		fmt.Println("Project and Logstore is not ready!")
		return
	}

	tickFreq = getclktck()
	interval = cfg.ReportInterval
	fmt.Printf("Tick frequent:%d\n", tickFreq)
	logM := getLogModule()
	producerInstance := newProducer(cfg)
	producerInstance.Start()
	defer producerInstance.SafeClose()
	ticker := time.NewTicker(time.Duration(cfg.ReportInterval) * time.Second)

	for {
		select {
		case <-ticker.C:
			/*optimize this method*/
			log := logM.generateLog(uint32(time.Now().Unix()))
			//fmt.Printf("Send log:%s\n", log.String())
			err := producerInstance.SendLog(cfg.ProjectName, cfg.LogStoreName, "topic", "127.0.0.1", log)
			if err != nil {
				fmt.Printf("Send log error!\n")
			}
		}
	}
}
