package icon

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/core"
)

var (
	// 图标数据
	iconList []Info
	// 选中栈
	selectedStack []map[string]*core.QVariant
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
	})
}

// change
func (o *IconListModel) change(idx *core.QVariant) {
	stat := true
	index := idx.ToInt(&stat)
	if o.modelData[index].IsActive {
		o.modelData[index].IsActive = false
	} else {
		o.modelData[index].IsActive = true
	}
	// o.ConnectRowCount(o.rowCount)
	// o.ConnectData(o.data)
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

	fmt.Println("pop")
	stat := true
	for idx, val := range selectedStack {
		fmt.Println("idx", val["index"].ToInt(&stat))
		fmt.Println("index", index.ToInt(&stat))
		if val["index"].ToInt(&stat) == index.ToInt(&stat) {
			selectedStack = append(selectedStack[:idx], selectedStack[idx+1:]...)
			break
		}
	}

	cs.update()
}

func (cs *IconStack) push(data map[string]*core.QVariant) {
	fmt.Println("push")
	selectedStack = append(selectedStack, data)
	cs.update()
}

// update 栈数据更新后同步更新相关数据
func (cs *IconStack) update() {
	fmt.Println("update")
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
