package toast

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

type Obj_ITF interface {
	std_core.QObject_ITF
	Obj_PTR() *Obj
}

func (ptr *Obj) Obj_PTR() *Obj {
	return ptr
}

func (ptr *Obj) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Obj) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromObj(ptr Obj_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Obj_PTR().Pointer()
	}
	return nil
}

func NewObjFromPointer(ptr unsafe.Pointer) (n *Obj) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Obj)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Obj:
			n = deduced

		case *std_core.QObject:
			n = &Obj{QObject: *deduced}

		default:
			n = new(Obj)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *Obj) Init() { this.init() }

//export callbackObj126d6c_Constructor
func callbackObj126d6c_Constructor(ptr unsafe.Pointer) {
	this := NewObjFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackObj126d6c_IsVisible
func callbackObj126d6c_IsVisible(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isVisible"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewObjFromPointer(ptr).IsVisibleDefault())))
}

func (ptr *Obj) ConnectIsVisible(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isVisible"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "isVisible", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isVisible", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Obj) DisconnectIsVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isVisible")
	}
}

func (ptr *Obj) IsVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.Obj126d6c_IsVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *Obj) IsVisibleDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.Obj126d6c_IsVisibleDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackObj126d6c_SetVisible
func callbackObj126d6c_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewObjFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *Obj) ConnectSetVisible(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setVisible"); signal != nil {
			f := func(visible bool) {
				(*(*func(bool))(signal))(visible)
				f(visible)
			}
			qt.ConnectSignal(ptr.Pointer(), "setVisible", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setVisible", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Obj) DisconnectSetVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setVisible")
	}
}

func (ptr *Obj) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.Obj126d6c_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *Obj) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.Obj126d6c_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackObj126d6c_VisibleChanged
func callbackObj126d6c_VisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "visibleChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	}

}

func (ptr *Obj) ConnectVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "visibleChanged") {
			C.Obj126d6c_ConnectVisibleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "visibleChanged"); signal != nil {
			f := func(visible bool) {
				(*(*func(bool))(signal))(visible)
				f(visible)
			}
			qt.ConnectSignal(ptr.Pointer(), "visibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "visibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Obj) DisconnectVisibleChanged() {
	if ptr.Pointer() != nil {
		C.Obj126d6c_DisconnectVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "visibleChanged")
	}
}

func (ptr *Obj) VisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.Obj126d6c_VisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackObj126d6c_Msg
func callbackObj126d6c_Msg(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "msg"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewObjFromPointer(ptr).MsgDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Obj) ConnectMsg(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "msg"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "msg", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "msg", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Obj) DisconnectMsg() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "msg")
	}
}

func (ptr *Obj) Msg() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Obj126d6c_Msg(ptr.Pointer()))
	}
	return ""
}

func (ptr *Obj) MsgDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Obj126d6c_MsgDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackObj126d6c_SetMsg
func callbackObj126d6c_SetMsg(ptr unsafe.Pointer, msg C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setMsg"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(msg))
	} else {
		NewObjFromPointer(ptr).SetMsgDefault(cGoUnpackString(msg))
	}
}

func (ptr *Obj) ConnectSetMsg(f func(msg string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setMsg"); signal != nil {
			f := func(msg string) {
				(*(*func(string))(signal))(msg)
				f(msg)
			}
			qt.ConnectSignal(ptr.Pointer(), "setMsg", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setMsg", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Obj) DisconnectSetMsg() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setMsg")
	}
}

func (ptr *Obj) SetMsg(msg string) {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		C.Obj126d6c_SetMsg(ptr.Pointer(), C.struct_Moc_PackedString{data: msgC, len: C.longlong(len(msg))})
	}
}

func (ptr *Obj) SetMsgDefault(msg string) {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		C.Obj126d6c_SetMsgDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: msgC, len: C.longlong(len(msg))})
	}
}

//export callbackObj126d6c_MsgChanged
func callbackObj126d6c_MsgChanged(ptr unsafe.Pointer, msg C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "msgChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(msg))
	}

}

func (ptr *Obj) ConnectMsgChanged(f func(msg string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "msgChanged") {
			C.Obj126d6c_ConnectMsgChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "msgChanged"); signal != nil {
			f := func(msg string) {
				(*(*func(string))(signal))(msg)
				f(msg)
			}
			qt.ConnectSignal(ptr.Pointer(), "msgChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "msgChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Obj) DisconnectMsgChanged() {
	if ptr.Pointer() != nil {
		C.Obj126d6c_DisconnectMsgChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "msgChanged")
	}
}

func (ptr *Obj) MsgChanged(msg string) {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		C.Obj126d6c_MsgChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: msgC, len: C.longlong(len(msg))})
	}
}

func Obj_QRegisterMetaType() int {
	return int(int32(C.Obj126d6c_Obj126d6c_QRegisterMetaType()))
}

func (ptr *Obj) QRegisterMetaType() int {
	return int(int32(C.Obj126d6c_Obj126d6c_QRegisterMetaType()))
}

func Obj_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Obj126d6c_Obj126d6c_QRegisterMetaType2(typeNameC)))
}

func (ptr *Obj) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Obj126d6c_Obj126d6c_QRegisterMetaType2(typeNameC)))
}

func Obj_QmlRegisterType() int {
	return int(int32(C.Obj126d6c_Obj126d6c_QmlRegisterType()))
}

func (ptr *Obj) QmlRegisterType() int {
	return int(int32(C.Obj126d6c_Obj126d6c_QmlRegisterType()))
}

func Obj_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Obj126d6c_Obj126d6c_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Obj) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Obj126d6c_Obj126d6c_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Obj) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Obj126d6c___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Obj) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Obj126d6c___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Obj) __children_newList() unsafe.Pointer {
	return C.Obj126d6c___children_newList(ptr.Pointer())
}

func (ptr *Obj) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.Obj126d6c___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Obj) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Obj126d6c___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Obj) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Obj126d6c___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Obj) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Obj126d6c___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Obj) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Obj126d6c___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Obj) __findChildren_newList() unsafe.Pointer {
	return C.Obj126d6c___findChildren_newList(ptr.Pointer())
}

func (ptr *Obj) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Obj126d6c___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Obj) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Obj126d6c___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Obj) __findChildren_newList3() unsafe.Pointer {
	return C.Obj126d6c___findChildren_newList3(ptr.Pointer())
}

func (ptr *Obj) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Obj126d6c___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Obj) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Obj126d6c___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Obj) __qFindChildren_newList2() unsafe.Pointer {
	return C.Obj126d6c___qFindChildren_newList2(ptr.Pointer())
}

func NewObj(parent std_core.QObject_ITF) *Obj {
	tmpValue := NewObjFromPointer(C.Obj126d6c_NewObj(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackObj126d6c_DestroyObj
func callbackObj126d6c_DestroyObj(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Obj"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewObjFromPointer(ptr).DestroyObjDefault()
	}
}

func (ptr *Obj) ConnectDestroyObj(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Obj"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~Obj", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Obj", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Obj) DisconnectDestroyObj() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Obj")
	}
}

func (ptr *Obj) DestroyObj() {
	if ptr.Pointer() != nil {
		C.Obj126d6c_DestroyObj(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Obj) DestroyObjDefault() {
	if ptr.Pointer() != nil {
		C.Obj126d6c_DestroyObjDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackObj126d6c_ChildEvent
func callbackObj126d6c_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewObjFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Obj) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Obj126d6c_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackObj126d6c_ConnectNotify
func callbackObj126d6c_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewObjFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Obj) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Obj126d6c_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackObj126d6c_CustomEvent
func callbackObj126d6c_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewObjFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Obj) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Obj126d6c_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackObj126d6c_DeleteLater
func callbackObj126d6c_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewObjFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Obj) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Obj126d6c_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackObj126d6c_Destroyed
func callbackObj126d6c_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackObj126d6c_DisconnectNotify
func callbackObj126d6c_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewObjFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Obj) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Obj126d6c_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackObj126d6c_Event
func callbackObj126d6c_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewObjFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Obj) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Obj126d6c_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackObj126d6c_EventFilter
func callbackObj126d6c_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewObjFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Obj) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Obj126d6c_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackObj126d6c_ObjectNameChanged
func callbackObj126d6c_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackObj126d6c_TimerEvent
func callbackObj126d6c_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewObjFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Obj) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Obj126d6c_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
