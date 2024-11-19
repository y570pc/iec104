package worker

import (
	"encoding/json"
	"fmt"
	"github.com/9d77v/iec104"
)

// Task 数据处理任务
func Task(data *iec104.APDU) {
	//TODO 自定义数据处理
	for _, signal := range data.Signals {
		signalJSON, _ := json.Marshal(signal)
		fmt.Println(string(signalJSON))
	}

}
