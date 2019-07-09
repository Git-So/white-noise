package icon

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Git-So/white-noise/logger"
)



// Load 加载文件信息
func Load(file string) (infoList []Info) {
	// 读取文件
	data, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Error("加载自定义音频配置文件信息出错:", err)
		return
	}

	// 解析数据
	err = json.Unmarshal(data, &infoList)
	if err != nil {
		logger.Error("解析自定义音频配置文件文件出错:", err)
		return
	}

	return
}
