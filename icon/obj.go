package icon

import (
	"strings"

	"github.com/Git-So/white-noise/audio"
	"github.com/Git-So/white-noise/logger"
	"github.com/Git-So/white-noise/toast"
	"github.com/therecipe/qt/core"
)

var (
	// 图标数据
	iconList []Info
	// 选中栈： 一开始是用栈的，后来发现思维错误
	selectedStack    []map[string]*core.QVariant
	SelectedStackObj *SelectedStackModel
)

// IconListModel 场景对象
type IconListModel struct {
	core.QAbstractListModel

	_ func()                   `constructor:"init"`
	_ func(idx *core.QVariant) `signal:"change,auto"`

	modelData []Info
}

// IconStack 选中场景栈
type IconStack struct {
	core.QObject

	_ string `property:"colorParam"`
	_ bool   `property:"isActive"`
	_ string `property:"title"`

	_ func() `constructor:"init"`

	_ func(index *core.QVariant)           `signal:"pop,auto"`
	_ func(data map[string]*core.QVariant) `signal:"push,auto"`
	_ func()                               `signal:"play,auto"`
}

// SelectedStackModel 选中音频栈
type SelectedStackModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ func(idx *core.QVariant, volume *core.QVariant) `signal:"write,auto"`

	modelData []map[string]*core.QVariant
}

// init IconListModel
func (o *IconListModel) init() {
	o.modelData = iconList

	o.ConnectRowCount(o.rowCount)
	o.ConnectData(o.data)
}

func (o *IconListModel) rowCount(*core.QModelIndex) int {
	return len(o.modelData)
}

func (o *IconListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}

	item := o.modelData[index.Row()]
	// return core.NewQVariant()
	return core.NewQVariant23(map[string]*core.QVariant{
		"index":      core.NewQVariant5(index.Row()),
		"name":       core.NewQVariant12(item.Title),
		"colorParam": core.NewQVariant12(item.ColorParam),
		"icon":       core.NewQVariant12(item.Icon),
		"isActive":   core.NewQVariant9(item.IsActive),
		"volume":     core.NewQVariant11(item.Volume),
	})
}

// change
func (o *IconListModel) change(idx *core.QVariant) {
	stat := true
	index := idx.ToInt(&stat)
	if o.modelData[index].IsActive {
		// 取消选中
		o.modelData[index].IsActive = false
	} else {
		// 最多同时播放三个音频
		if len(selectedStack) >= 3 {
			toast.Msg("最多可以选三个")
			return
		}
		o.modelData[index].IsActive = true
	}
	o.DataChanged(o.Index(index, 0, core.NewQModelIndex()), o.Index(index, 0, core.NewQModelIndex()), []int{int(core.Qt__DisplayRole)})
}

// init IconStack
func (cs *IconStack) init() {
	cs.SetColorParam("")
	cs.SetIsActive(false)
}

func (cs *IconStack) pop(index *core.QVariant) {
	if len(selectedStack) == 0 {
		return
	}

	logger.Debug("pop")
	stat := true

	for idx, val := range selectedStack {
		logger.Debug("idx", val["index"].ToInt(&stat))
		logger.Debug("index", index.ToInt(&stat))
		if val["index"].ToInt(&stat) == index.ToInt(&stat) {
			// 移除播放
			removePlayer(val["index"].ToInt(&stat))
			// 出栈
			selectedStack = append(selectedStack[:idx], selectedStack[idx+1:]...)
			break
		}
	}

	cs.update()
}

func (cs *IconStack) push(data map[string]*core.QVariant) {
	// 最多同时播放三个音频
	if len(selectedStack) >= 3 {
		toast.Msg("最多可以选三个")
		return
	}

	logger.Debug("push")
	selectedStack = append(selectedStack, data)
	stat := true
	logger.Debug("volume", data["volume"].ToFloat(&stat))
	cs.update()

	// 添加到播放
	addPlayer(data["index"].ToInt(&stat))
}

// update 栈数据更新后同步更新相关数据
func (cs *IconStack) update() {
	logger.Debug("update")
	stackLen := len(selectedStack)
	if stackLen == 0 {
		cs.SetIsActive(false)
	} else {
		// 标题
		var names []string
		for _, val := range selectedStack {
			names = append(names, val["name"].ToString())
		}
		title := strings.Join(names, "/")
		cs.SetTitle(title)

		// 其他信息
		selectedInfo := selectedStack[stackLen-1]
		cs.SetColorParam(selectedInfo["colorParam"].ToString())
		cs.SetIsActive(true)
	}
}

// play 播放音频
func (cs *IconStack) play() {
	logger.Debug("icons: play 播放音频")

	for _, info := range selectedStack {
		// 添加到播放
		if _, ok := info["index"]; ok {
			// 添加播放
			addPlayer(info["index"].ToInt(&ok))

			// 获取音频信息
			idx := info["index"].ToInt(&ok)
			iconInfo := iconList[idx]
			// stat := true
			for _, val := range iconInfo.AudioUrls {
				// 获取播放实例
				p := audio.GetPlayer(val.URL)
				// 开始播放
				p.SetState(audio.PLAYER_ENABLE)
			}
		}
	}
}

// init SelectedStackModel
func (ss *SelectedStackModel) init() {
	SelectedStackObj = ss
	ss.modelData = selectedStack

	ss.ConnectRowCount(ss.rowCount)
	ss.ConnectData(ss.data)
}

// rowCount SelectedStackModel
func (ss *SelectedStackModel) rowCount(*core.QModelIndex) int {
	return len(ss.modelData)
}

// data SelectedStackModel
func (ss *SelectedStackModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}

	item := ss.modelData[index.Row()]
	// return core.NewQVariant()
	return core.NewQVariant23(item)
}

// Write 写入音量
func (ss *SelectedStackModel) write(index *core.QVariant, volume *core.QVariant) {
	stat := true
	logger.Debug("index", index.ToInt(&stat))
	logger.Debug("index", volume.ToFloat(&stat))

	key := index.ToInt(&stat)
	for idx, val := range selectedStack {
		ValIndex := val["index"].ToInt(&stat)
		if ValIndex == key {
			selectedStack[idx]["volume"] = volume
			updateVolume(ValIndex, volume.ToFloat(&stat))
		}
	}

	// ss.DataChanged(ss.Index(key, 0, core.NewQModelIndex()), ss.Index(key, 0, core.NewQModelIndex()), []int{int(core.Qt__DisplayRole)})
}
