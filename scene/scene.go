package scene

import (
	"encoding/json"
	"io/ioutil"
)

// LoadFile 加载文件信息
func LoadFile(file string) (info *Info) {
	// 读取文件
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	// 解析数据
	err = json.Unmarshal(data, info)
	if err != nil {
		return
	}

	return
}
