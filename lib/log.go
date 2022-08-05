package lib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	config "github.com/spf13/viper"
)

type Log struct {
	App         string                 // 服务名
	Url         string                 // 路由
	Data        map[string]interface{} // 数据传输参数
	RequestType string                 // request_start or request_end
	TraceId     string                 // 链路唯一标识
	Timestamp   time.Time              // 发生时间 精确到秒
}

func isExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// 返回字符串格式类似 - 20220804
func getDateString(year int, month time.Month, day int) string {
	res := fmt.Sprintf("%d%s%d", year, month, day)
	return res
}

// 当链路上有服务请求以及响应的时候，写入对应的日志文件 ~/logs/{app}/rp-request-{timestamp}.log
func (log Log) Info() error {
	dir := config.Get("log.path").(string)
	infoPath := fmt.Sprintf(dir+"/%s/info/rp-request-%s.log", log.App, getDateString(log.Timestamp.Date())) // 服务间请求调用日志文件
	if !isExists(infoPath) {
		_, err := os.Create(infoPath)
		if err != nil {
			return err
		}
	}

	file, err := os.OpenFile(infoPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)

	// 构造日志明细
	dataByte, err := json.Marshal(log.Data)
	if err != nil {
		return err
	}
	logSlice := []string{log.Timestamp.String(), log.RequestType, log.TraceId, log.Url, string(dataByte)}
	logStr := strings.Join(logSlice, "|")
	write.WriteString(logStr + " \n")

	//Flush将缓存的文件真正写入到文件中
	write.Flush()
	return nil
}
