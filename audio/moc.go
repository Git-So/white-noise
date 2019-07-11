package audio

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

type Play_ITF interface {
	std_core.QObject_ITF
	Play_PTR() *Play
}

func (ptr *Play) Play_PTR() *Play {
	return ptr
}

func (ptr *Play) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Play) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromPlay(ptr Play_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Play_PTR().Pointer()
	}
	return nil
}

func NewPlayFromPointer(ptr unsafe.Pointer) (n *Play) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Play)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Play:
			n = deduced

		case *std_core.QObject:
			n = &Play{QObject: *deduced}

		default:
			n = new(Play)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *Play) Init() { this.init() }

//export callbackPlayf024d8_Constructor
func callbackPlayf024d8_Constructor(ptr unsafe.Pointer) {
	this := NewPlayFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectStop(this.stop)
	this.ConnectStart(this.start)
	this.init()
}

//export callbackPlayf024d8_Stop
func callbackPlayf024d8_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stop"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *Play) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stop") {
			C.Playf024d8_ConnectStop(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stop"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "stop", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stop", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Play) DisconnectStop() {
	if ptr.Pointer() != nil {
		C.Playf024d8_DisconnectStop(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stop")
	}
}

func (ptr *Play) Stop() {
	if ptr.Pointer() != nil {
		C.Playf024d8_Stop(ptr.Pointer())
	}
}

//export callbackPlayf024d8_Start
func callbackPlayf024d8_Start(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "start"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *Play) ConnectStart(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "start") {
			C.Playf024d8_ConnectStart(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "start"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "start", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "start", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Play) DisconnectStart() {
	if ptr.Pointer() != nil {
		C.Playf024d8_DisconnectStart(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "start")
	}
}

func (ptr *Play) Start() {
	if ptr.Pointer() != nil {
		C.Playf024d8_Start(ptr.Pointer())
	}
}

//export callbackPlayf024d8_IsState
func callbackPlayf024d8_IsState(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isState"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewPlayFromPointer(ptr).IsStateDefault())))
}

func (ptr *Play) ConnectIsState(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isState"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "isState", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isState", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Play) DisconnectIsState() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isState")
	}
}

func (ptr *Play) IsState() bool {
	if ptr.Pointer() != nil {
		return int8(C.Playf024d8_IsState(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *Play) IsStateDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.Playf024d8_IsStateDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackPlayf024d8_SetState
func callbackPlayf024d8_SetState(ptr unsafe.Pointer, state C.char) {
	if signal := qt.GetSignal(ptr, "setState"); signal != nil {
		(*(*func(bool))(signal))(int8(state) != 0)
	} else {
		NewPlayFromPointer(ptr).SetStateDefault(int8(state) != 0)
	}
}

func (ptr *Play) ConnectSetState(f func(state bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setState"); signal != nil {
			f := func(state bool) {
				(*(*func(bool))(signal))(state)
				f(state)
			}
			qt.ConnectSignal(ptr.Pointer(), "setState", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setState", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Play) DisconnectSetState() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setState")
	}
}

func (ptr *Play) SetState(state bool) {
	if ptr.Pointer() != nil {
		C.Playf024d8_SetState(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(state))))
	}
}

func (ptr *Play) SetStateDefault(state bool) {
	if ptr.Pointer() != nil {
		C.Playf024d8_SetStateDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(state))))
	}
}

//export callbackPlayf024d8_StateChanged
func callbackPlayf024d8_StateChanged(ptr unsafe.Pointer, state C.char) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(state) != 0)
	}

}

func (ptr *Play) ConnectStateChanged(f func(state bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.Playf024d8_ConnectStateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			f := func(state bool) {
				(*(*func(bool))(signal))(state)
				f(state)
			}
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Play) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.Playf024d8_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
	}
}

func (ptr *Play) StateChanged(state bool) {
	if ptr.Pointer() != nil {
		C.Playf024d8_StateChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(state))))
	}
}

func Play_QRegisterMetaType() int {
	return int(int32(C.Playf024d8_Playf024d8_QRegisterMetaType()))
}

func (ptr *Play) QRegisterMetaType() int {
	return int(int32(C.Playf024d8_Playf024d8_QRegisterMetaType()))
}

func Play_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Playf024d8_Playf024d8_QRegisterMetaType2(typeNameC)))
}

func (ptr *Play) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Playf024d8_Playf024d8_QRegisterMetaType2(typeNameC)))
}

func Play_QmlRegisterType() int {
	return int(int32(C.Playf024d8_Playf024d8_QmlRegisterType()))
}

func (ptr *Play) QmlRegisterType() int {
	return int(int32(C.Playf024d8_Playf024d8_QmlRegisterType()))
}

func Play_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Playf024d8_Playf024d8_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Play) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Playf024d8_Playf024d8_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Play) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Playf024d8___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Play) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Playf024d8___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Play) __children_newList() unsafe.Pointer {
	return C.Playf024d8___children_newList(ptr.Pointer())
}

func (ptr *Play) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.Playf024d8___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Play) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Playf024d8___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Play) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Playf024d8___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Play) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Playf024d8___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Play) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Playf024d8___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Play) __findChildren_newList() unsafe.Pointer {
	return C.Playf024d8___findChildren_newList(ptr.Pointer())
}

func (ptr *Play) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Playf024d8___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Play) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Playf024d8___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Play) __findChildren_newList3() unsafe.Pointer {
	return C.Playf024d8___findChildren_newList3(ptr.Pointer())
}

func (ptr *Play) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Playf024d8___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Play) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Playf024d8___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Play) __qFindChildren_newList2() unsafe.Pointer {
	return C.Playf024d8___qFindChildren_newList2(ptr.Pointer())
}

func NewPlay(parent std_core.QObject_ITF) *Play {
	tmpValue := NewPlayFromPointer(C.Playf024d8_NewPlay(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackPlayf024d8_DestroyPlay
func callbackPlayf024d8_DestroyPlay(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Play"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlayFromPointer(ptr).DestroyPlayDefault()
	}
}

func (ptr *Play) ConnectDestroyPlay(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Play"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~Play", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Play", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Play) DisconnectDestroyPlay() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Play")
	}
}

func (ptr *Play) DestroyPlay() {
	if ptr.Pointer() != nil {
		C.Playf024d8_DestroyPlay(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Play) DestroyPlayDefault() {
	if ptr.Pointer() != nil {
		C.Playf024d8_DestroyPlayDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackPlayf024d8_ChildEvent
func callbackPlayf024d8_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewPlayFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Play) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Playf024d8_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackPlayf024d8_ConnectNotify
func callbackPlayf024d8_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewPlayFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Play) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Playf024d8_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackPlayf024d8_CustomEvent
func callbackPlayf024d8_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewPlayFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Play) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Playf024d8_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackPlayf024d8_DeleteLater
func callbackPlayf024d8_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlayFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Play) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Playf024d8_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackPlayf024d8_Destroyed
func callbackPlayf024d8_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackPlayf024d8_DisconnectNotify
func callbackPlayf024d8_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewPlayFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Play) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Playf024d8_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackPlayf024d8_Event
func callbackPlayf024d8_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPlayFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Play) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Playf024d8_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackPlayf024d8_EventFilter
func callbackPlayf024d8_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPlayFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Play) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Playf024d8_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackPlayf024d8_ObjectNameChanged
func callbackPlayf024d8_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackPlayf024d8_TimerEvent
func callbackPlayf024d8_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewPlayFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Play) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Playf024d8_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
