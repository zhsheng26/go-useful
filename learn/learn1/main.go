package learn1

import (
	"fmt"
	"strings"
)

func main() {
	logger := NewLogger(UsLogger)
	logger.Debug("hello")
}

type Log func(info string) string

type Print interface {
	Debug(info string)
}

func (log Log) Debug(info string) {
	fmt.Println("debug level:", log(info))
}

func NewLogger(log Log) Print {
	return log
}

func UsLogger(info string) string {
	return strings.ToUpper(info)
}
