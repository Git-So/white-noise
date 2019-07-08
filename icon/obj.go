package icon

import (
	"fmt"

	"github.com/therecipe/qt/core"
)

var (
	iconList   []Info
	colorstack []string
	viewColor  string
)

// IconListModel 场景对象
type IconListModel struct {
	core.QAbstractListModel

	_ func()        `constructor:"init"`
	_ func(idx int) `signal:"change,auto"`

	modelData []Info
}

// IconColorStack 选中场景栈
type IconColorStack struct {
	core.QObject

	_ string `property:"colorParam"`
	_ bool   `property:"isActive"`

	_ func() `constructor:"init"`

	_ func() `signal:"update,auto"`
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
		"title":      core.NewQVariant12(item.Title),
		"colorParam": core.NewQVariant12(item.ColorParam),
		"icon":       core.NewQVariant12(item.Icon),
	})
}

// change
func (o *IconListModel) change(idx int) {
	color := iconList[idx].ColorParam
	var isExist bool
	var index int
	for key, val := range colorstack {
		if val == color {
			isExist = true
			index = key
		}
	}
	if isExist {
		colorstack = append(colorstack[:index], colorstack[index+1:]...)
	} else {
		colorstack = append(colorstack, color)
	}

	viewColor = colorstack[len(colorstack)-1]
}

// init IconColorStack
func (cs *IconColorStack) init() {
	cs.SetColorParam("")
	cs.SetIsActive(false)
}

func (cs *IconColorStack) update() {
	fmt.Println("update")
}
