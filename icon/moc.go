package icon

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type SelectedStackModel_ITF interface {
	std_core.QAbstractListModel_ITF
	SelectedStackModel_PTR() *SelectedStackModel
}

func (ptr *SelectedStackModel) SelectedStackModel_PTR() *SelectedStackModel {
	return ptr
}

func (ptr *SelectedStackModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *SelectedStackModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromSelectedStackModel(ptr SelectedStackModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SelectedStackModel_PTR().Pointer()
	}
	return nil
}

func NewSelectedStackModelFromPointer(ptr unsafe.Pointer) (n *SelectedStackModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(SelectedStackModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *SelectedStackModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &SelectedStackModel{QAbstractListModel: *deduced}

		default:
			n = new(SelectedStackModel)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *SelectedStackModel) Init() { this.init() }

//export callbackSelectedStackModeld54ccb_Constructor
func callbackSelectedStackModeld54ccb_Constructor(ptr unsafe.Pointer) {
	this := NewSelectedStackModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectWrite(this.write)
	this.init()
}

//export callbackSelectedStackModeld54ccb_Write
func callbackSelectedStackModeld54ccb_Write(ptr unsafe.Pointer, idx unsafe.Pointer, volume unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "write"); signal != nil {
		(*(*func(*std_core.QVariant, *std_core.QVariant))(signal))(std_core.NewQVariantFromPointer(idx), std_core.NewQVariantFromPointer(volume))
	}

}

func (ptr *SelectedStackModel) ConnectWrite(f func(idx *std_core.QVariant, volume *std_core.QVariant)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "write") {
			C.SelectedStackModeld54ccb_ConnectWrite(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "write"); signal != nil {
			f := func(idx *std_core.QVariant, volume *std_core.QVariant) {
				(*(*func(*std_core.QVariant, *std_core.QVariant))(signal))(idx, volume)
				f(idx, volume)
			}
			qt.ConnectSignal(ptr.Pointer(), "write", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "write", unsafe.Pointer(&f))
		}
	}
}

func (ptr *SelectedStackModel) DisconnectWrite() {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_DisconnectWrite(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "write")
	}
}

func (ptr *SelectedStackModel) Write(idx std_core.QVariant_ITF, volume std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_Write(ptr.Pointer(), std_core.PointerFromQVariant(idx), std_core.PointerFromQVariant(volume))
	}
}

func SelectedStackModel_QRegisterMetaType() int {
	return int(int32(C.SelectedStackModeld54ccb_SelectedStackModeld54ccb_QRegisterMetaType()))
}

func (ptr *SelectedStackModel) QRegisterMetaType() int {
	return int(int32(C.SelectedStackModeld54ccb_SelectedStackModeld54ccb_QRegisterMetaType()))
}

func SelectedStackModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.SelectedStackModeld54ccb_SelectedStackModeld54ccb_QRegisterMetaType2(typeNameC)))
}

func (ptr *SelectedStackModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.SelectedStackModeld54ccb_SelectedStackModeld54ccb_QRegisterMetaType2(typeNameC)))
}

func SelectedStackModel_QmlRegisterType() int {
	return int(int32(C.SelectedStackModeld54ccb_SelectedStackModeld54ccb_QmlRegisterType()))
}

func (ptr *SelectedStackModel) QmlRegisterType() int {
	return int(int32(C.SelectedStackModeld54ccb_SelectedStackModeld54ccb_QmlRegisterType()))
}

func SelectedStackModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.SelectedStackModeld54ccb_SelectedStackModeld54ccb_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *SelectedStackModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.SelectedStackModeld54ccb_SelectedStackModeld54ccb_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *SelectedStackModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.SelectedStackModeld54ccb_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *SelectedStackModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *SelectedStackModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.SelectedStackModeld54ccb_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *SelectedStackModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *SelectedStackModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.SelectedStackModeld54ccb_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *SelectedStackModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *SelectedStackModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.SelectedStackModeld54ccb___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *SelectedStackModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *SelectedStackModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.SelectedStackModeld54ccb___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *SelectedStackModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *SelectedStackModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.SelectedStackModeld54ccb___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *SelectedStackModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *SelectedStackModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.SelectedStackModeld54ccb___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *SelectedStackModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *SelectedStackModel) __itemData_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___itemData_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewSelectedStackModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.SelectedStackModeld54ccb___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *SelectedStackModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.SelectedStackModeld54ccb___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *SelectedStackModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *SelectedStackModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.SelectedStackModeld54ccb___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *SelectedStackModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *SelectedStackModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.SelectedStackModeld54ccb___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *SelectedStackModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *SelectedStackModel) __match_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___match_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.SelectedStackModeld54ccb___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *SelectedStackModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *SelectedStackModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.SelectedStackModeld54ccb___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *SelectedStackModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *SelectedStackModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.SelectedStackModeld54ccb___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *SelectedStackModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *SelectedStackModel) __roleNames_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___roleNames_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewSelectedStackModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.SelectedStackModeld54ccb___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *SelectedStackModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.SelectedStackModeld54ccb___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *SelectedStackModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *SelectedStackModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewSelectedStackModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.SelectedStackModeld54ccb___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *SelectedStackModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.SelectedStackModeld54ccb_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *SelectedStackModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *SelectedStackModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.SelectedStackModeld54ccb_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *SelectedStackModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *SelectedStackModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SelectedStackModeld54ccb___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SelectedStackModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SelectedStackModel) __children_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___children_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.SelectedStackModeld54ccb___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *SelectedStackModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *SelectedStackModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SelectedStackModeld54ccb___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SelectedStackModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SelectedStackModel) __findChildren_newList() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___findChildren_newList(ptr.Pointer())
}

func (ptr *SelectedStackModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SelectedStackModeld54ccb___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SelectedStackModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SelectedStackModel) __findChildren_newList3() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___findChildren_newList3(ptr.Pointer())
}

func (ptr *SelectedStackModel) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SelectedStackModeld54ccb___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SelectedStackModel) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SelectedStackModel) __qFindChildren_newList2() unsafe.Pointer {
	return C.SelectedStackModeld54ccb___qFindChildren_newList2(ptr.Pointer())
}

func NewSelectedStackModel(parent std_core.QObject_ITF) *SelectedStackModel {
	tmpValue := NewSelectedStackModelFromPointer(C.SelectedStackModeld54ccb_NewSelectedStackModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackSelectedStackModeld54ccb_DestroySelectedStackModel
func callbackSelectedStackModeld54ccb_DestroySelectedStackModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~SelectedStackModel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSelectedStackModelFromPointer(ptr).DestroySelectedStackModelDefault()
	}
}

func (ptr *SelectedStackModel) ConnectDestroySelectedStackModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~SelectedStackModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~SelectedStackModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~SelectedStackModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *SelectedStackModel) DisconnectDestroySelectedStackModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~SelectedStackModel")
	}
}

func (ptr *SelectedStackModel) DestroySelectedStackModel() {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_DestroySelectedStackModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *SelectedStackModel) DestroySelectedStackModelDefault() {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_DestroySelectedStackModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackSelectedStackModeld54ccb_DropMimeData
func callbackSelectedStackModeld54ccb_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *SelectedStackModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_Flags
func callbackSelectedStackModeld54ccb_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewSelectedStackModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *SelectedStackModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.SelectedStackModeld54ccb_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackSelectedStackModeld54ccb_Index
func callbackSelectedStackModeld54ccb_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewSelectedStackModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *SelectedStackModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.SelectedStackModeld54ccb_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackSelectedStackModeld54ccb_Sibling
func callbackSelectedStackModeld54ccb_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewSelectedStackModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *SelectedStackModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.SelectedStackModeld54ccb_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackSelectedStackModeld54ccb_Buddy
func callbackSelectedStackModeld54ccb_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewSelectedStackModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *SelectedStackModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.SelectedStackModeld54ccb_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackSelectedStackModeld54ccb_CanDropMimeData
func callbackSelectedStackModeld54ccb_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *SelectedStackModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_CanFetchMore
func callbackSelectedStackModeld54ccb_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *SelectedStackModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_ColumnCount
func callbackSelectedStackModeld54ccb_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewSelectedStackModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *SelectedStackModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.SelectedStackModeld54ccb_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackSelectedStackModeld54ccb_ColumnsAboutToBeInserted
func callbackSelectedStackModeld54ccb_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackSelectedStackModeld54ccb_ColumnsAboutToBeMoved
func callbackSelectedStackModeld54ccb_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackSelectedStackModeld54ccb_ColumnsAboutToBeRemoved
func callbackSelectedStackModeld54ccb_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackSelectedStackModeld54ccb_ColumnsInserted
func callbackSelectedStackModeld54ccb_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackSelectedStackModeld54ccb_ColumnsMoved
func callbackSelectedStackModeld54ccb_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackSelectedStackModeld54ccb_ColumnsRemoved
func callbackSelectedStackModeld54ccb_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackSelectedStackModeld54ccb_Data
func callbackSelectedStackModeld54ccb_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(*std_core.QModelIndex, int) *std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewSelectedStackModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *SelectedStackModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.SelectedStackModeld54ccb_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackSelectedStackModeld54ccb_DataChanged
func callbackSelectedStackModeld54ccb_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		(*(*func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(signal))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewSelectedStackModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackSelectedStackModeld54ccb_FetchMore
func callbackSelectedStackModeld54ccb_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*std_core.QModelIndex))(signal))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewSelectedStackModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *SelectedStackModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackSelectedStackModeld54ccb_HasChildren
func callbackSelectedStackModeld54ccb_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *SelectedStackModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_HeaderData
func callbackSelectedStackModeld54ccb_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewSelectedStackModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *SelectedStackModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.SelectedStackModeld54ccb_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackSelectedStackModeld54ccb_HeaderDataChanged
func callbackSelectedStackModeld54ccb_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(std_core.Qt__Orientation, int, int))(signal))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackSelectedStackModeld54ccb_InsertColumns
func callbackSelectedStackModeld54ccb_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *SelectedStackModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_InsertRows
func callbackSelectedStackModeld54ccb_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *SelectedStackModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_ItemData
func callbackSelectedStackModeld54ccb_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewSelectedStackModelFromPointer(NewSelectedStackModelFromPointer(nil).__itemData_newList())
			for k, v := range (*(*func(*std_core.QModelIndex) map[int]*std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewSelectedStackModelFromPointer(NewSelectedStackModelFromPointer(nil).__itemData_newList())
		for k, v := range NewSelectedStackModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *SelectedStackModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewSelectedStackModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.SelectedStackModeld54ccb_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackSelectedStackModeld54ccb_LayoutAboutToBeChanged
func callbackSelectedStackModeld54ccb_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewSelectedStackModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackSelectedStackModeld54ccb_LayoutChanged
func callbackSelectedStackModeld54ccb_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewSelectedStackModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackSelectedStackModeld54ccb_Match
func callbackSelectedStackModeld54ccb_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewSelectedStackModelFromPointer(NewSelectedStackModelFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewSelectedStackModelFromPointer(NewSelectedStackModelFromPointer(nil).__match_newList())
		for _, v := range NewSelectedStackModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *SelectedStackModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewSelectedStackModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.SelectedStackModeld54ccb_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackSelectedStackModeld54ccb_MimeData
func callbackSelectedStackModeld54ccb_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData((*(*func([]*std_core.QModelIndex) *std_core.QMimeData)(signal))(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewSelectedStackModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewSelectedStackModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewSelectedStackModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *SelectedStackModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.SelectedStackModeld54ccb_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewSelectedStackModelFromPointer(NewSelectedStackModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackSelectedStackModeld54ccb_MimeTypes
func callbackSelectedStackModeld54ccb_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewSelectedStackModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *SelectedStackModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.SelectedStackModeld54ccb_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackSelectedStackModeld54ccb_ModelAboutToBeReset
func callbackSelectedStackModeld54ccb_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackSelectedStackModeld54ccb_ModelReset
func callbackSelectedStackModeld54ccb_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackSelectedStackModeld54ccb_MoveColumns
func callbackSelectedStackModeld54ccb_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *SelectedStackModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_MoveRows
func callbackSelectedStackModeld54ccb_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *SelectedStackModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_Parent
func callbackSelectedStackModeld54ccb_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewSelectedStackModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *SelectedStackModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.SelectedStackModeld54ccb_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackSelectedStackModeld54ccb_RemoveColumns
func callbackSelectedStackModeld54ccb_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *SelectedStackModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_RemoveRows
func callbackSelectedStackModeld54ccb_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *SelectedStackModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_ResetInternalData
func callbackSelectedStackModeld54ccb_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSelectedStackModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *SelectedStackModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackSelectedStackModeld54ccb_Revert
func callbackSelectedStackModeld54ccb_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSelectedStackModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *SelectedStackModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_RevertDefault(ptr.Pointer())
	}
}

//export callbackSelectedStackModeld54ccb_RoleNames
func callbackSelectedStackModeld54ccb_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewSelectedStackModelFromPointer(NewSelectedStackModelFromPointer(nil).__roleNames_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewSelectedStackModelFromPointer(NewSelectedStackModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewSelectedStackModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *SelectedStackModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewSelectedStackModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.SelectedStackModeld54ccb_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackSelectedStackModeld54ccb_RowCount
func callbackSelectedStackModeld54ccb_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewSelectedStackModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *SelectedStackModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.SelectedStackModeld54ccb_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackSelectedStackModeld54ccb_RowsAboutToBeInserted
func callbackSelectedStackModeld54ccb_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackSelectedStackModeld54ccb_RowsAboutToBeMoved
func callbackSelectedStackModeld54ccb_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackSelectedStackModeld54ccb_RowsAboutToBeRemoved
func callbackSelectedStackModeld54ccb_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackSelectedStackModeld54ccb_RowsInserted
func callbackSelectedStackModeld54ccb_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackSelectedStackModeld54ccb_RowsMoved
func callbackSelectedStackModeld54ccb_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackSelectedStackModeld54ccb_RowsRemoved
func callbackSelectedStackModeld54ccb_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackSelectedStackModeld54ccb_SetData
func callbackSelectedStackModeld54ccb_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(signal))(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *SelectedStackModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_SetHeaderData
func callbackSelectedStackModeld54ccb_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *SelectedStackModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_SetItemData
func callbackSelectedStackModeld54ccb_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(signal))(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewSelectedStackModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewSelectedStackModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *SelectedStackModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewSelectedStackModelFromPointer(NewSelectedStackModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_Sort
func callbackSelectedStackModeld54ccb_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, std_core.Qt__SortOrder))(signal))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewSelectedStackModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *SelectedStackModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackSelectedStackModeld54ccb_Span
func callbackSelectedStackModeld54ccb_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize((*(*func(*std_core.QModelIndex) *std_core.QSize)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewSelectedStackModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *SelectedStackModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.SelectedStackModeld54ccb_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackSelectedStackModeld54ccb_Submit
func callbackSelectedStackModeld54ccb_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *SelectedStackModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_SupportedDragActions
func callbackSelectedStackModeld54ccb_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewSelectedStackModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *SelectedStackModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.SelectedStackModeld54ccb_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackSelectedStackModeld54ccb_SupportedDropActions
func callbackSelectedStackModeld54ccb_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewSelectedStackModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *SelectedStackModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.SelectedStackModeld54ccb_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackSelectedStackModeld54ccb_ChildEvent
func callbackSelectedStackModeld54ccb_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewSelectedStackModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *SelectedStackModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackSelectedStackModeld54ccb_ConnectNotify
func callbackSelectedStackModeld54ccb_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewSelectedStackModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *SelectedStackModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackSelectedStackModeld54ccb_CustomEvent
func callbackSelectedStackModeld54ccb_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewSelectedStackModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *SelectedStackModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackSelectedStackModeld54ccb_DeleteLater
func callbackSelectedStackModeld54ccb_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSelectedStackModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *SelectedStackModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackSelectedStackModeld54ccb_Destroyed
func callbackSelectedStackModeld54ccb_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackSelectedStackModeld54ccb_DisconnectNotify
func callbackSelectedStackModeld54ccb_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewSelectedStackModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *SelectedStackModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackSelectedStackModeld54ccb_Event
func callbackSelectedStackModeld54ccb_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *SelectedStackModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_EventFilter
func callbackSelectedStackModeld54ccb_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSelectedStackModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *SelectedStackModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SelectedStackModeld54ccb_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackSelectedStackModeld54ccb_ObjectNameChanged
func callbackSelectedStackModeld54ccb_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackSelectedStackModeld54ccb_TimerEvent
func callbackSelectedStackModeld54ccb_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewSelectedStackModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *SelectedStackModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SelectedStackModeld54ccb_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type IconListModel_ITF interface {
	std_core.QAbstractListModel_ITF
	IconListModel_PTR() *IconListModel
}

func (ptr *IconListModel) IconListModel_PTR() *IconListModel {
	return ptr
}

func (ptr *IconListModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *IconListModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromIconListModel(ptr IconListModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.IconListModel_PTR().Pointer()
	}
	return nil
}

func NewIconListModelFromPointer(ptr unsafe.Pointer) (n *IconListModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(IconListModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *IconListModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &IconListModel{QAbstractListModel: *deduced}

		default:
			n = new(IconListModel)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *IconListModel) Init() { this.init() }

//export callbackIconListModeld54ccb_Constructor
func callbackIconListModeld54ccb_Constructor(ptr unsafe.Pointer) {
	this := NewIconListModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectChange(this.change)
	this.init()
}

//export callbackIconListModeld54ccb_Change
func callbackIconListModeld54ccb_Change(ptr unsafe.Pointer, idx unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "change"); signal != nil {
		(*(*func(*std_core.QVariant))(signal))(std_core.NewQVariantFromPointer(idx))
	}

}

func (ptr *IconListModel) ConnectChange(f func(idx *std_core.QVariant)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "change") {
			C.IconListModeld54ccb_ConnectChange(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "change"); signal != nil {
			f := func(idx *std_core.QVariant) {
				(*(*func(*std_core.QVariant))(signal))(idx)
				f(idx)
			}
			qt.ConnectSignal(ptr.Pointer(), "change", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "change", unsafe.Pointer(&f))
		}
	}
}

func (ptr *IconListModel) DisconnectChange() {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_DisconnectChange(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "change")
	}
}

func (ptr *IconListModel) Change(idx std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_Change(ptr.Pointer(), std_core.PointerFromQVariant(idx))
	}
}

func IconListModel_QRegisterMetaType() int {
	return int(int32(C.IconListModeld54ccb_IconListModeld54ccb_QRegisterMetaType()))
}

func (ptr *IconListModel) QRegisterMetaType() int {
	return int(int32(C.IconListModeld54ccb_IconListModeld54ccb_QRegisterMetaType()))
}

func IconListModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.IconListModeld54ccb_IconListModeld54ccb_QRegisterMetaType2(typeNameC)))
}

func (ptr *IconListModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.IconListModeld54ccb_IconListModeld54ccb_QRegisterMetaType2(typeNameC)))
}

func IconListModel_QmlRegisterType() int {
	return int(int32(C.IconListModeld54ccb_IconListModeld54ccb_QmlRegisterType()))
}

func (ptr *IconListModel) QmlRegisterType() int {
	return int(int32(C.IconListModeld54ccb_IconListModeld54ccb_QmlRegisterType()))
}

func IconListModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.IconListModeld54ccb_IconListModeld54ccb_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *IconListModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.IconListModeld54ccb_IconListModeld54ccb_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *IconListModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.IconListModeld54ccb_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *IconListModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *IconListModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.IconListModeld54ccb_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *IconListModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.IconListModeld54ccb_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *IconListModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *IconListModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.IconListModeld54ccb_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *IconListModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.IconListModeld54ccb_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *IconListModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *IconListModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.IconListModeld54ccb_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *IconListModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.IconListModeld54ccb___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *IconListModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *IconListModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.IconListModeld54ccb___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *IconListModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.IconListModeld54ccb___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *IconListModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *IconListModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.IconListModeld54ccb___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *IconListModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.IconListModeld54ccb___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *IconListModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *IconListModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.IconListModeld54ccb___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *IconListModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.IconListModeld54ccb___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *IconListModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *IconListModel) __itemData_newList() unsafe.Pointer {
	return C.IconListModeld54ccb___itemData_newList(ptr.Pointer())
}

func (ptr *IconListModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewIconListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.IconListModeld54ccb___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *IconListModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.IconListModeld54ccb___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *IconListModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *IconListModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.IconListModeld54ccb___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *IconListModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.IconListModeld54ccb___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *IconListModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *IconListModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.IconListModeld54ccb___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *IconListModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.IconListModeld54ccb___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *IconListModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *IconListModel) __match_newList() unsafe.Pointer {
	return C.IconListModeld54ccb___match_newList(ptr.Pointer())
}

func (ptr *IconListModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.IconListModeld54ccb___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *IconListModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *IconListModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.IconListModeld54ccb___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *IconListModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.IconListModeld54ccb___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *IconListModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *IconListModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.IconListModeld54ccb___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *IconListModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.IconListModeld54ccb___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *IconListModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *IconListModel) __roleNames_newList() unsafe.Pointer {
	return C.IconListModeld54ccb___roleNames_newList(ptr.Pointer())
}

func (ptr *IconListModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewIconListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.IconListModeld54ccb___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *IconListModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.IconListModeld54ccb___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *IconListModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *IconListModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.IconListModeld54ccb___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *IconListModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewIconListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.IconListModeld54ccb___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *IconListModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.IconListModeld54ccb_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *IconListModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *IconListModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.IconListModeld54ccb_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *IconListModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.IconListModeld54ccb_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *IconListModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *IconListModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.IconListModeld54ccb_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *IconListModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.IconListModeld54ccb___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *IconListModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *IconListModel) __children_newList() unsafe.Pointer {
	return C.IconListModeld54ccb___children_newList(ptr.Pointer())
}

func (ptr *IconListModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.IconListModeld54ccb___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *IconListModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *IconListModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.IconListModeld54ccb___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *IconListModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.IconListModeld54ccb___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *IconListModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *IconListModel) __findChildren_newList() unsafe.Pointer {
	return C.IconListModeld54ccb___findChildren_newList(ptr.Pointer())
}

func (ptr *IconListModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.IconListModeld54ccb___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *IconListModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *IconListModel) __findChildren_newList3() unsafe.Pointer {
	return C.IconListModeld54ccb___findChildren_newList3(ptr.Pointer())
}

func (ptr *IconListModel) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.IconListModeld54ccb___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *IconListModel) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *IconListModel) __qFindChildren_newList2() unsafe.Pointer {
	return C.IconListModeld54ccb___qFindChildren_newList2(ptr.Pointer())
}

func NewIconListModel(parent std_core.QObject_ITF) *IconListModel {
	tmpValue := NewIconListModelFromPointer(C.IconListModeld54ccb_NewIconListModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackIconListModeld54ccb_DestroyIconListModel
func callbackIconListModeld54ccb_DestroyIconListModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~IconListModel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewIconListModelFromPointer(ptr).DestroyIconListModelDefault()
	}
}

func (ptr *IconListModel) ConnectDestroyIconListModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~IconListModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~IconListModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~IconListModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *IconListModel) DisconnectDestroyIconListModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~IconListModel")
	}
}

func (ptr *IconListModel) DestroyIconListModel() {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_DestroyIconListModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *IconListModel) DestroyIconListModelDefault() {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_DestroyIconListModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackIconListModeld54ccb_DropMimeData
func callbackIconListModeld54ccb_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *IconListModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_Flags
func callbackIconListModeld54ccb_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewIconListModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *IconListModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.IconListModeld54ccb_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackIconListModeld54ccb_Index
func callbackIconListModeld54ccb_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewIconListModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *IconListModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.IconListModeld54ccb_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackIconListModeld54ccb_Sibling
func callbackIconListModeld54ccb_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewIconListModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *IconListModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.IconListModeld54ccb_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackIconListModeld54ccb_Buddy
func callbackIconListModeld54ccb_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewIconListModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *IconListModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.IconListModeld54ccb_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackIconListModeld54ccb_CanDropMimeData
func callbackIconListModeld54ccb_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *IconListModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_CanFetchMore
func callbackIconListModeld54ccb_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *IconListModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_ColumnCount
func callbackIconListModeld54ccb_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewIconListModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *IconListModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.IconListModeld54ccb_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackIconListModeld54ccb_ColumnsAboutToBeInserted
func callbackIconListModeld54ccb_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackIconListModeld54ccb_ColumnsAboutToBeMoved
func callbackIconListModeld54ccb_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackIconListModeld54ccb_ColumnsAboutToBeRemoved
func callbackIconListModeld54ccb_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackIconListModeld54ccb_ColumnsInserted
func callbackIconListModeld54ccb_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackIconListModeld54ccb_ColumnsMoved
func callbackIconListModeld54ccb_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackIconListModeld54ccb_ColumnsRemoved
func callbackIconListModeld54ccb_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackIconListModeld54ccb_Data
func callbackIconListModeld54ccb_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(*std_core.QModelIndex, int) *std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewIconListModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *IconListModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.IconListModeld54ccb_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackIconListModeld54ccb_DataChanged
func callbackIconListModeld54ccb_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		(*(*func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(signal))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewIconListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackIconListModeld54ccb_FetchMore
func callbackIconListModeld54ccb_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*std_core.QModelIndex))(signal))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewIconListModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *IconListModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackIconListModeld54ccb_HasChildren
func callbackIconListModeld54ccb_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *IconListModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_HeaderData
func callbackIconListModeld54ccb_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewIconListModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *IconListModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.IconListModeld54ccb_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackIconListModeld54ccb_HeaderDataChanged
func callbackIconListModeld54ccb_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(std_core.Qt__Orientation, int, int))(signal))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackIconListModeld54ccb_InsertColumns
func callbackIconListModeld54ccb_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *IconListModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_InsertRows
func callbackIconListModeld54ccb_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *IconListModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_ItemData
func callbackIconListModeld54ccb_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewIconListModelFromPointer(NewIconListModelFromPointer(nil).__itemData_newList())
			for k, v := range (*(*func(*std_core.QModelIndex) map[int]*std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewIconListModelFromPointer(NewIconListModelFromPointer(nil).__itemData_newList())
		for k, v := range NewIconListModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *IconListModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewIconListModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.IconListModeld54ccb_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackIconListModeld54ccb_LayoutAboutToBeChanged
func callbackIconListModeld54ccb_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewIconListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackIconListModeld54ccb_LayoutChanged
func callbackIconListModeld54ccb_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewIconListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackIconListModeld54ccb_Match
func callbackIconListModeld54ccb_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewIconListModelFromPointer(NewIconListModelFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewIconListModelFromPointer(NewIconListModelFromPointer(nil).__match_newList())
		for _, v := range NewIconListModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *IconListModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewIconListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.IconListModeld54ccb_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackIconListModeld54ccb_MimeData
func callbackIconListModeld54ccb_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData((*(*func([]*std_core.QModelIndex) *std_core.QMimeData)(signal))(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewIconListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewIconListModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewIconListModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *IconListModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.IconListModeld54ccb_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewIconListModelFromPointer(NewIconListModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackIconListModeld54ccb_MimeTypes
func callbackIconListModeld54ccb_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewIconListModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *IconListModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.IconListModeld54ccb_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackIconListModeld54ccb_ModelAboutToBeReset
func callbackIconListModeld54ccb_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackIconListModeld54ccb_ModelReset
func callbackIconListModeld54ccb_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackIconListModeld54ccb_MoveColumns
func callbackIconListModeld54ccb_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *IconListModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_MoveRows
func callbackIconListModeld54ccb_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *IconListModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_Parent
func callbackIconListModeld54ccb_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewIconListModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *IconListModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.IconListModeld54ccb_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackIconListModeld54ccb_RemoveColumns
func callbackIconListModeld54ccb_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *IconListModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_RemoveRows
func callbackIconListModeld54ccb_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *IconListModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_ResetInternalData
func callbackIconListModeld54ccb_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewIconListModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *IconListModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackIconListModeld54ccb_Revert
func callbackIconListModeld54ccb_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewIconListModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *IconListModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_RevertDefault(ptr.Pointer())
	}
}

//export callbackIconListModeld54ccb_RoleNames
func callbackIconListModeld54ccb_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewIconListModelFromPointer(NewIconListModelFromPointer(nil).__roleNames_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewIconListModelFromPointer(NewIconListModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewIconListModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *IconListModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewIconListModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.IconListModeld54ccb_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackIconListModeld54ccb_RowCount
func callbackIconListModeld54ccb_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewIconListModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *IconListModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.IconListModeld54ccb_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackIconListModeld54ccb_RowsAboutToBeInserted
func callbackIconListModeld54ccb_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackIconListModeld54ccb_RowsAboutToBeMoved
func callbackIconListModeld54ccb_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackIconListModeld54ccb_RowsAboutToBeRemoved
func callbackIconListModeld54ccb_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackIconListModeld54ccb_RowsInserted
func callbackIconListModeld54ccb_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackIconListModeld54ccb_RowsMoved
func callbackIconListModeld54ccb_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackIconListModeld54ccb_RowsRemoved
func callbackIconListModeld54ccb_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackIconListModeld54ccb_SetData
func callbackIconListModeld54ccb_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(signal))(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *IconListModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_SetHeaderData
func callbackIconListModeld54ccb_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *IconListModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_SetItemData
func callbackIconListModeld54ccb_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(signal))(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewIconListModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewIconListModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *IconListModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewIconListModelFromPointer(NewIconListModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_Sort
func callbackIconListModeld54ccb_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, std_core.Qt__SortOrder))(signal))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewIconListModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *IconListModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackIconListModeld54ccb_Span
func callbackIconListModeld54ccb_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize((*(*func(*std_core.QModelIndex) *std_core.QSize)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewIconListModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *IconListModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.IconListModeld54ccb_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackIconListModeld54ccb_Submit
func callbackIconListModeld54ccb_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *IconListModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_SupportedDragActions
func callbackIconListModeld54ccb_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewIconListModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *IconListModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.IconListModeld54ccb_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackIconListModeld54ccb_SupportedDropActions
func callbackIconListModeld54ccb_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewIconListModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *IconListModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.IconListModeld54ccb_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackIconListModeld54ccb_ChildEvent
func callbackIconListModeld54ccb_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewIconListModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *IconListModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackIconListModeld54ccb_ConnectNotify
func callbackIconListModeld54ccb_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewIconListModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *IconListModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackIconListModeld54ccb_CustomEvent
func callbackIconListModeld54ccb_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewIconListModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *IconListModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackIconListModeld54ccb_DeleteLater
func callbackIconListModeld54ccb_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewIconListModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *IconListModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackIconListModeld54ccb_Destroyed
func callbackIconListModeld54ccb_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackIconListModeld54ccb_DisconnectNotify
func callbackIconListModeld54ccb_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewIconListModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *IconListModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackIconListModeld54ccb_Event
func callbackIconListModeld54ccb_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *IconListModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_EventFilter
func callbackIconListModeld54ccb_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconListModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *IconListModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconListModeld54ccb_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackIconListModeld54ccb_ObjectNameChanged
func callbackIconListModeld54ccb_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackIconListModeld54ccb_TimerEvent
func callbackIconListModeld54ccb_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewIconListModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *IconListModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.IconListModeld54ccb_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type IconStack_ITF interface {
	std_core.QObject_ITF
	IconStack_PTR() *IconStack
}

func (ptr *IconStack) IconStack_PTR() *IconStack {
	return ptr
}

func (ptr *IconStack) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *IconStack) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromIconStack(ptr IconStack_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.IconStack_PTR().Pointer()
	}
	return nil
}

func NewIconStackFromPointer(ptr unsafe.Pointer) (n *IconStack) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(IconStack)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *IconStack:
			n = deduced

		case *std_core.QObject:
			n = &IconStack{QObject: *deduced}

		default:
			n = new(IconStack)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *IconStack) Init() { this.init() }

//export callbackIconStackd54ccb_Constructor
func callbackIconStackd54ccb_Constructor(ptr unsafe.Pointer) {
	this := NewIconStackFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectPop(this.pop)
	this.ConnectPush(this.push)
	this.init()
}

//export callbackIconStackd54ccb_Pop
func callbackIconStackd54ccb_Pop(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pop"); signal != nil {
		(*(*func(*std_core.QVariant))(signal))(std_core.NewQVariantFromPointer(index))
	}

}

func (ptr *IconStack) ConnectPop(f func(index *std_core.QVariant)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pop") {
			C.IconStackd54ccb_ConnectPop(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pop"); signal != nil {
			f := func(index *std_core.QVariant) {
				(*(*func(*std_core.QVariant))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "pop", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pop", unsafe.Pointer(&f))
		}
	}
}

func (ptr *IconStack) DisconnectPop() {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_DisconnectPop(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pop")
	}
}

func (ptr *IconStack) Pop(index std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_Pop(ptr.Pointer(), std_core.PointerFromQVariant(index))
	}
}

//export callbackIconStackd54ccb_Push
func callbackIconStackd54ccb_Push(ptr unsafe.Pointer, data C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "push"); signal != nil {
		(*(*func(map[string]*std_core.QVariant))(signal))(func(l C.struct_Moc_PackedList) map[string]*std_core.QVariant {
			out := make(map[string]*std_core.QVariant, int(l.len))
			tmpList := NewIconStackFromPointer(l.data)
			for i, v := range tmpList.__push_data_keyList() {
				out[v] = tmpList.__push_data_atList(v, i)
			}
			return out
		}(data))
	}

}

func (ptr *IconStack) ConnectPush(f func(data map[string]*std_core.QVariant)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "push") {
			C.IconStackd54ccb_ConnectPush(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "push"); signal != nil {
			f := func(data map[string]*std_core.QVariant) {
				(*(*func(map[string]*std_core.QVariant))(signal))(data)
				f(data)
			}
			qt.ConnectSignal(ptr.Pointer(), "push", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "push", unsafe.Pointer(&f))
		}
	}
}

func (ptr *IconStack) DisconnectPush() {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_DisconnectPush(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "push")
	}
}

func (ptr *IconStack) Push(data map[string]*std_core.QVariant) {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_Push(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewIconStackFromPointer(NewIconStackFromPointer(nil).__push_data_newList())
			for k, v := range data {
				tmpList.__push_data_setList(k, std_core.NewQVariant1(v))
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackIconStackd54ccb_ColorParam
func callbackIconStackd54ccb_ColorParam(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "colorParam"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewIconStackFromPointer(ptr).ColorParamDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *IconStack) ConnectColorParam(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "colorParam"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "colorParam", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "colorParam", unsafe.Pointer(&f))
		}
	}
}

func (ptr *IconStack) DisconnectColorParam() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "colorParam")
	}
}

func (ptr *IconStack) ColorParam() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.IconStackd54ccb_ColorParam(ptr.Pointer()))
	}
	return ""
}

func (ptr *IconStack) ColorParamDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.IconStackd54ccb_ColorParamDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackIconStackd54ccb_SetColorParam
func callbackIconStackd54ccb_SetColorParam(ptr unsafe.Pointer, colorParam C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setColorParam"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(colorParam))
	} else {
		NewIconStackFromPointer(ptr).SetColorParamDefault(cGoUnpackString(colorParam))
	}
}

func (ptr *IconStack) ConnectSetColorParam(f func(colorParam string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setColorParam"); signal != nil {
			f := func(colorParam string) {
				(*(*func(string))(signal))(colorParam)
				f(colorParam)
			}
			qt.ConnectSignal(ptr.Pointer(), "setColorParam", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setColorParam", unsafe.Pointer(&f))
		}
	}
}

func (ptr *IconStack) DisconnectSetColorParam() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setColorParam")
	}
}

func (ptr *IconStack) SetColorParam(colorParam string) {
	if ptr.Pointer() != nil {
		var colorParamC *C.char
		if colorParam != "" {
			colorParamC = C.CString(colorParam)
			defer C.free(unsafe.Pointer(colorParamC))
		}
		C.IconStackd54ccb_SetColorParam(ptr.Pointer(), C.struct_Moc_PackedString{data: colorParamC, len: C.longlong(len(colorParam))})
	}
}

func (ptr *IconStack) SetColorParamDefault(colorParam string) {
	if ptr.Pointer() != nil {
		var colorParamC *C.char
		if colorParam != "" {
			colorParamC = C.CString(colorParam)
			defer C.free(unsafe.Pointer(colorParamC))
		}
		C.IconStackd54ccb_SetColorParamDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: colorParamC, len: C.longlong(len(colorParam))})
	}
}

//export callbackIconStackd54ccb_ColorParamChanged
func callbackIconStackd54ccb_ColorParamChanged(ptr unsafe.Pointer, colorParam C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "colorParamChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(colorParam))
	}

}

func (ptr *IconStack) ConnectColorParamChanged(f func(colorParam string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "colorParamChanged") {
			C.IconStackd54ccb_ConnectColorParamChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "colorParamChanged"); signal != nil {
			f := func(colorParam string) {
				(*(*func(string))(signal))(colorParam)
				f(colorParam)
			}
			qt.ConnectSignal(ptr.Pointer(), "colorParamChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "colorParamChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *IconStack) DisconnectColorParamChanged() {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_DisconnectColorParamChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "colorParamChanged")
	}
}

func (ptr *IconStack) ColorParamChanged(colorParam string) {
	if ptr.Pointer() != nil {
		var colorParamC *C.char
		if colorParam != "" {
			colorParamC = C.CString(colorParam)
			defer C.free(unsafe.Pointer(colorParamC))
		}
		C.IconStackd54ccb_ColorParamChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: colorParamC, len: C.longlong(len(colorParam))})
	}
}

//export callbackIconStackd54ccb_IsActive
func callbackIconStackd54ccb_IsActive(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isActive"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconStackFromPointer(ptr).IsActiveDefault())))
}

func (ptr *IconStack) ConnectIsActive(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isActive"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "isActive", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isActive", unsafe.Pointer(&f))
		}
	}
}

func (ptr *IconStack) DisconnectIsActive() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isActive")
	}
}

func (ptr *IconStack) IsActive() bool {
	if ptr.Pointer() != nil {
		return int8(C.IconStackd54ccb_IsActive(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *IconStack) IsActiveDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.IconStackd54ccb_IsActiveDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackIconStackd54ccb_SetIsActive
func callbackIconStackd54ccb_SetIsActive(ptr unsafe.Pointer, isActive C.char) {
	if signal := qt.GetSignal(ptr, "setIsActive"); signal != nil {
		(*(*func(bool))(signal))(int8(isActive) != 0)
	} else {
		NewIconStackFromPointer(ptr).SetIsActiveDefault(int8(isActive) != 0)
	}
}

func (ptr *IconStack) ConnectSetIsActive(f func(isActive bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setIsActive"); signal != nil {
			f := func(isActive bool) {
				(*(*func(bool))(signal))(isActive)
				f(isActive)
			}
			qt.ConnectSignal(ptr.Pointer(), "setIsActive", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setIsActive", unsafe.Pointer(&f))
		}
	}
}

func (ptr *IconStack) DisconnectSetIsActive() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setIsActive")
	}
}

func (ptr *IconStack) SetIsActive(isActive bool) {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_SetIsActive(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isActive))))
	}
}

func (ptr *IconStack) SetIsActiveDefault(isActive bool) {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_SetIsActiveDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isActive))))
	}
}

//export callbackIconStackd54ccb_IsActiveChanged
func callbackIconStackd54ccb_IsActiveChanged(ptr unsafe.Pointer, isActive C.char) {
	if signal := qt.GetSignal(ptr, "isActiveChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(isActive) != 0)
	}

}

func (ptr *IconStack) ConnectIsActiveChanged(f func(isActive bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "isActiveChanged") {
			C.IconStackd54ccb_ConnectIsActiveChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "isActiveChanged"); signal != nil {
			f := func(isActive bool) {
				(*(*func(bool))(signal))(isActive)
				f(isActive)
			}
			qt.ConnectSignal(ptr.Pointer(), "isActiveChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isActiveChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *IconStack) DisconnectIsActiveChanged() {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_DisconnectIsActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "isActiveChanged")
	}
}

func (ptr *IconStack) IsActiveChanged(isActive bool) {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_IsActiveChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isActive))))
	}
}

//export callbackIconStackd54ccb_Title
func callbackIconStackd54ccb_Title(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "title"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewIconStackFromPointer(ptr).TitleDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *IconStack) ConnectTitle(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "title"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "title", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "title", unsafe.Pointer(&f))
		}
	}
}

func (ptr *IconStack) DisconnectTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "title")
	}
}

func (ptr *IconStack) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.IconStackd54ccb_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *IconStack) TitleDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.IconStackd54ccb_TitleDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackIconStackd54ccb_SetTitle
func callbackIconStackd54ccb_SetTitle(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	} else {
		NewIconStackFromPointer(ptr).SetTitleDefault(cGoUnpackString(title))
	}
}

func (ptr *IconStack) ConnectSetTitle(f func(title string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTitle"); signal != nil {
			f := func(title string) {
				(*(*func(string))(signal))(title)
				f(title)
			}
			qt.ConnectSignal(ptr.Pointer(), "setTitle", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTitle", unsafe.Pointer(&f))
		}
	}
}

func (ptr *IconStack) DisconnectSetTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTitle")
	}
}

func (ptr *IconStack) SetTitle(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.IconStackd54ccb_SetTitle(ptr.Pointer(), C.struct_Moc_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

func (ptr *IconStack) SetTitleDefault(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.IconStackd54ccb_SetTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

//export callbackIconStackd54ccb_TitleChanged
func callbackIconStackd54ccb_TitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "titleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

func (ptr *IconStack) ConnectTitleChanged(f func(title string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "titleChanged") {
			C.IconStackd54ccb_ConnectTitleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "titleChanged"); signal != nil {
			f := func(title string) {
				(*(*func(string))(signal))(title)
				f(title)
			}
			qt.ConnectSignal(ptr.Pointer(), "titleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "titleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *IconStack) DisconnectTitleChanged() {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "titleChanged")
	}
}

func (ptr *IconStack) TitleChanged(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.IconStackd54ccb_TitleChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

func IconStack_QRegisterMetaType() int {
	return int(int32(C.IconStackd54ccb_IconStackd54ccb_QRegisterMetaType()))
}

func (ptr *IconStack) QRegisterMetaType() int {
	return int(int32(C.IconStackd54ccb_IconStackd54ccb_QRegisterMetaType()))
}

func IconStack_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.IconStackd54ccb_IconStackd54ccb_QRegisterMetaType2(typeNameC)))
}

func (ptr *IconStack) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.IconStackd54ccb_IconStackd54ccb_QRegisterMetaType2(typeNameC)))
}

func IconStack_QmlRegisterType() int {
	return int(int32(C.IconStackd54ccb_IconStackd54ccb_QmlRegisterType()))
}

func (ptr *IconStack) QmlRegisterType() int {
	return int(int32(C.IconStackd54ccb_IconStackd54ccb_QmlRegisterType()))
}

func IconStack_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.IconStackd54ccb_IconStackd54ccb_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *IconStack) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.IconStackd54ccb_IconStackd54ccb_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *IconStack) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.IconStackd54ccb___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *IconStack) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *IconStack) __children_newList() unsafe.Pointer {
	return C.IconStackd54ccb___children_newList(ptr.Pointer())
}

func (ptr *IconStack) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.IconStackd54ccb___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *IconStack) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *IconStack) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.IconStackd54ccb___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *IconStack) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.IconStackd54ccb___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *IconStack) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *IconStack) __findChildren_newList() unsafe.Pointer {
	return C.IconStackd54ccb___findChildren_newList(ptr.Pointer())
}

func (ptr *IconStack) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.IconStackd54ccb___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *IconStack) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *IconStack) __findChildren_newList3() unsafe.Pointer {
	return C.IconStackd54ccb___findChildren_newList3(ptr.Pointer())
}

func (ptr *IconStack) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.IconStackd54ccb___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *IconStack) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *IconStack) __qFindChildren_newList2() unsafe.Pointer {
	return C.IconStackd54ccb___qFindChildren_newList2(ptr.Pointer())
}

func (ptr *IconStack) __push_data_atList(v string, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := std_core.NewQVariantFromPointer(C.IconStackd54ccb___push_data_atList(ptr.Pointer(), C.struct_Moc_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *IconStack) __push_data_setList(key string, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.IconStackd54ccb___push_data_setList(ptr.Pointer(), C.struct_Moc_PackedString{data: keyC, len: C.longlong(len(key))}, std_core.PointerFromQVariant(i))
	}
}

func (ptr *IconStack) __push_data_newList() unsafe.Pointer {
	return C.IconStackd54ccb___push_data_newList(ptr.Pointer())
}

func (ptr *IconStack) __push_data_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewIconStackFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____push_data_keyList_atList(i)
			}
			return out
		}(C.IconStackd54ccb___push_data_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *IconStack) ____push_data_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.IconStackd54ccb_____push_data_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *IconStack) ____push_data_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.IconStackd54ccb_____push_data_keyList_setList(ptr.Pointer(), C.struct_Moc_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *IconStack) ____push_data_keyList_newList() unsafe.Pointer {
	return C.IconStackd54ccb_____push_data_keyList_newList(ptr.Pointer())
}

func NewIconStack(parent std_core.QObject_ITF) *IconStack {
	tmpValue := NewIconStackFromPointer(C.IconStackd54ccb_NewIconStack(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackIconStackd54ccb_DestroyIconStack
func callbackIconStackd54ccb_DestroyIconStack(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~IconStack"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewIconStackFromPointer(ptr).DestroyIconStackDefault()
	}
}

func (ptr *IconStack) ConnectDestroyIconStack(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~IconStack"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~IconStack", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~IconStack", unsafe.Pointer(&f))
		}
	}
}

func (ptr *IconStack) DisconnectDestroyIconStack() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~IconStack")
	}
}

func (ptr *IconStack) DestroyIconStack() {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_DestroyIconStack(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *IconStack) DestroyIconStackDefault() {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_DestroyIconStackDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackIconStackd54ccb_ChildEvent
func callbackIconStackd54ccb_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewIconStackFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *IconStack) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackIconStackd54ccb_ConnectNotify
func callbackIconStackd54ccb_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewIconStackFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *IconStack) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackIconStackd54ccb_CustomEvent
func callbackIconStackd54ccb_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewIconStackFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *IconStack) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackIconStackd54ccb_DeleteLater
func callbackIconStackd54ccb_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewIconStackFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *IconStack) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackIconStackd54ccb_Destroyed
func callbackIconStackd54ccb_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackIconStackd54ccb_DisconnectNotify
func callbackIconStackd54ccb_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewIconStackFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *IconStack) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackIconStackd54ccb_Event
func callbackIconStackd54ccb_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconStackFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *IconStack) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconStackd54ccb_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackIconStackd54ccb_EventFilter
func callbackIconStackd54ccb_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewIconStackFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *IconStack) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.IconStackd54ccb_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackIconStackd54ccb_ObjectNameChanged
func callbackIconStackd54ccb_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackIconStackd54ccb_TimerEvent
func callbackIconStackd54ccb_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewIconStackFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *IconStack) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.IconStackd54ccb_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
