package scene

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

//export callbackObjd4edec_Constructor
func callbackObjd4edec_Constructor(ptr unsafe.Pointer) {
	this := NewObjFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectNext(this.next)
	this.init()
}

//export callbackObjd4edec_Next
func callbackObjd4edec_Next(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "next"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *Obj) ConnectNext(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "next") {
			C.Objd4edec_ConnectNext(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "next"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "next", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "next", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Obj) DisconnectNext() {
	if ptr.Pointer() != nil {
		C.Objd4edec_DisconnectNext(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "next")
	}
}

func (ptr *Obj) Next() {
	if ptr.Pointer() != nil {
		C.Objd4edec_Next(ptr.Pointer())
	}
}

//export callbackObjd4edec_Title
func callbackObjd4edec_Title(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "title"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewObjFromPointer(ptr).TitleDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Obj) ConnectTitle(f func() string) {
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

func (ptr *Obj) DisconnectTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "title")
	}
}

func (ptr *Obj) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Objd4edec_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *Obj) TitleDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Objd4edec_TitleDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackObjd4edec_SetTitle
func callbackObjd4edec_SetTitle(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	} else {
		NewObjFromPointer(ptr).SetTitleDefault(cGoUnpackString(title))
	}
}

func (ptr *Obj) ConnectSetTitle(f func(title string)) {
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

func (ptr *Obj) DisconnectSetTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTitle")
	}
}

func (ptr *Obj) SetTitle(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.Objd4edec_SetTitle(ptr.Pointer(), C.struct_Moc_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

func (ptr *Obj) SetTitleDefault(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.Objd4edec_SetTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

//export callbackObjd4edec_TitleChanged
func callbackObjd4edec_TitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "titleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

func (ptr *Obj) ConnectTitleChanged(f func(title string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "titleChanged") {
			C.Objd4edec_ConnectTitleChanged(ptr.Pointer())
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

func (ptr *Obj) DisconnectTitleChanged() {
	if ptr.Pointer() != nil {
		C.Objd4edec_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "titleChanged")
	}
}

func (ptr *Obj) TitleChanged(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.Objd4edec_TitleChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

//export callbackObjd4edec_Desp
func callbackObjd4edec_Desp(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "desp"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewObjFromPointer(ptr).DespDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Obj) ConnectDesp(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "desp"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "desp", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "desp", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Obj) DisconnectDesp() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "desp")
	}
}

func (ptr *Obj) Desp() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Objd4edec_Desp(ptr.Pointer()))
	}
	return ""
}

func (ptr *Obj) DespDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Objd4edec_DespDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackObjd4edec_SetDesp
func callbackObjd4edec_SetDesp(ptr unsafe.Pointer, desp C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setDesp"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(desp))
	} else {
		NewObjFromPointer(ptr).SetDespDefault(cGoUnpackString(desp))
	}
}

func (ptr *Obj) ConnectSetDesp(f func(desp string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDesp"); signal != nil {
			f := func(desp string) {
				(*(*func(string))(signal))(desp)
				f(desp)
			}
			qt.ConnectSignal(ptr.Pointer(), "setDesp", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDesp", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Obj) DisconnectSetDesp() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDesp")
	}
}

func (ptr *Obj) SetDesp(desp string) {
	if ptr.Pointer() != nil {
		var despC *C.char
		if desp != "" {
			despC = C.CString(desp)
			defer C.free(unsafe.Pointer(despC))
		}
		C.Objd4edec_SetDesp(ptr.Pointer(), C.struct_Moc_PackedString{data: despC, len: C.longlong(len(desp))})
	}
}

func (ptr *Obj) SetDespDefault(desp string) {
	if ptr.Pointer() != nil {
		var despC *C.char
		if desp != "" {
			despC = C.CString(desp)
			defer C.free(unsafe.Pointer(despC))
		}
		C.Objd4edec_SetDespDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: despC, len: C.longlong(len(desp))})
	}
}

//export callbackObjd4edec_DespChanged
func callbackObjd4edec_DespChanged(ptr unsafe.Pointer, desp C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "despChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(desp))
	}

}

func (ptr *Obj) ConnectDespChanged(f func(desp string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "despChanged") {
			C.Objd4edec_ConnectDespChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "despChanged"); signal != nil {
			f := func(desp string) {
				(*(*func(string))(signal))(desp)
				f(desp)
			}
			qt.ConnectSignal(ptr.Pointer(), "despChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "despChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Obj) DisconnectDespChanged() {
	if ptr.Pointer() != nil {
		C.Objd4edec_DisconnectDespChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "despChanged")
	}
}

func (ptr *Obj) DespChanged(desp string) {
	if ptr.Pointer() != nil {
		var despC *C.char
		if desp != "" {
			despC = C.CString(desp)
			defer C.free(unsafe.Pointer(despC))
		}
		C.Objd4edec_DespChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: despC, len: C.longlong(len(desp))})
	}
}

//export callbackObjd4edec_ImagePath
func callbackObjd4edec_ImagePath(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "imagePath"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewObjFromPointer(ptr).ImagePathDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Obj) ConnectImagePath(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "imagePath"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "imagePath", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "imagePath", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Obj) DisconnectImagePath() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "imagePath")
	}
}

func (ptr *Obj) ImagePath() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Objd4edec_ImagePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *Obj) ImagePathDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Objd4edec_ImagePathDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackObjd4edec_SetImagePath
func callbackObjd4edec_SetImagePath(ptr unsafe.Pointer, imagePath C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setImagePath"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(imagePath))
	} else {
		NewObjFromPointer(ptr).SetImagePathDefault(cGoUnpackString(imagePath))
	}
}

func (ptr *Obj) ConnectSetImagePath(f func(imagePath string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setImagePath"); signal != nil {
			f := func(imagePath string) {
				(*(*func(string))(signal))(imagePath)
				f(imagePath)
			}
			qt.ConnectSignal(ptr.Pointer(), "setImagePath", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setImagePath", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Obj) DisconnectSetImagePath() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setImagePath")
	}
}

func (ptr *Obj) SetImagePath(imagePath string) {
	if ptr.Pointer() != nil {
		var imagePathC *C.char
		if imagePath != "" {
			imagePathC = C.CString(imagePath)
			defer C.free(unsafe.Pointer(imagePathC))
		}
		C.Objd4edec_SetImagePath(ptr.Pointer(), C.struct_Moc_PackedString{data: imagePathC, len: C.longlong(len(imagePath))})
	}
}

func (ptr *Obj) SetImagePathDefault(imagePath string) {
	if ptr.Pointer() != nil {
		var imagePathC *C.char
		if imagePath != "" {
			imagePathC = C.CString(imagePath)
			defer C.free(unsafe.Pointer(imagePathC))
		}
		C.Objd4edec_SetImagePathDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: imagePathC, len: C.longlong(len(imagePath))})
	}
}

//export callbackObjd4edec_ImagePathChanged
func callbackObjd4edec_ImagePathChanged(ptr unsafe.Pointer, imagePath C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "imagePathChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(imagePath))
	}

}

func (ptr *Obj) ConnectImagePathChanged(f func(imagePath string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "imagePathChanged") {
			C.Objd4edec_ConnectImagePathChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "imagePathChanged"); signal != nil {
			f := func(imagePath string) {
				(*(*func(string))(signal))(imagePath)
				f(imagePath)
			}
			qt.ConnectSignal(ptr.Pointer(), "imagePathChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "imagePathChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Obj) DisconnectImagePathChanged() {
	if ptr.Pointer() != nil {
		C.Objd4edec_DisconnectImagePathChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "imagePathChanged")
	}
}

func (ptr *Obj) ImagePathChanged(imagePath string) {
	if ptr.Pointer() != nil {
		var imagePathC *C.char
		if imagePath != "" {
			imagePathC = C.CString(imagePath)
			defer C.free(unsafe.Pointer(imagePathC))
		}
		C.Objd4edec_ImagePathChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: imagePathC, len: C.longlong(len(imagePath))})
	}
}

func Obj_QRegisterMetaType() int {
	return int(int32(C.Objd4edec_Objd4edec_QRegisterMetaType()))
}

func (ptr *Obj) QRegisterMetaType() int {
	return int(int32(C.Objd4edec_Objd4edec_QRegisterMetaType()))
}

func Obj_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Objd4edec_Objd4edec_QRegisterMetaType2(typeNameC)))
}

func (ptr *Obj) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Objd4edec_Objd4edec_QRegisterMetaType2(typeNameC)))
}

func Obj_QmlRegisterType() int {
	return int(int32(C.Objd4edec_Objd4edec_QmlRegisterType()))
}

func (ptr *Obj) QmlRegisterType() int {
	return int(int32(C.Objd4edec_Objd4edec_QmlRegisterType()))
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
	return int(int32(C.Objd4edec_Objd4edec_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.Objd4edec_Objd4edec_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Obj) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Objd4edec___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Obj) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Objd4edec___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Obj) __children_newList() unsafe.Pointer {
	return C.Objd4edec___children_newList(ptr.Pointer())
}

func (ptr *Obj) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.Objd4edec___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Obj) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Objd4edec___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Obj) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Objd4edec___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Obj) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Objd4edec___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Obj) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Objd4edec___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Obj) __findChildren_newList() unsafe.Pointer {
	return C.Objd4edec___findChildren_newList(ptr.Pointer())
}

func (ptr *Obj) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Objd4edec___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Obj) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Objd4edec___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Obj) __findChildren_newList3() unsafe.Pointer {
	return C.Objd4edec___findChildren_newList3(ptr.Pointer())
}

func (ptr *Obj) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Objd4edec___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Obj) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Objd4edec___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Obj) __qFindChildren_newList2() unsafe.Pointer {
	return C.Objd4edec___qFindChildren_newList2(ptr.Pointer())
}

func NewObj(parent std_core.QObject_ITF) *Obj {
	tmpValue := NewObjFromPointer(C.Objd4edec_NewObj(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackObjd4edec_DestroyObj
func callbackObjd4edec_DestroyObj(ptr unsafe.Pointer) {
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
		C.Objd4edec_DestroyObj(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Obj) DestroyObjDefault() {
	if ptr.Pointer() != nil {
		C.Objd4edec_DestroyObjDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackObjd4edec_ChildEvent
func callbackObjd4edec_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewObjFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Obj) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Objd4edec_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackObjd4edec_ConnectNotify
func callbackObjd4edec_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewObjFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Obj) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Objd4edec_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackObjd4edec_CustomEvent
func callbackObjd4edec_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewObjFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Obj) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Objd4edec_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackObjd4edec_DeleteLater
func callbackObjd4edec_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewObjFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Obj) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Objd4edec_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackObjd4edec_Destroyed
func callbackObjd4edec_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackObjd4edec_DisconnectNotify
func callbackObjd4edec_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewObjFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Obj) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Objd4edec_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackObjd4edec_Event
func callbackObjd4edec_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewObjFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Obj) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Objd4edec_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackObjd4edec_EventFilter
func callbackObjd4edec_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewObjFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Obj) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Objd4edec_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackObjd4edec_ObjectNameChanged
func callbackObjd4edec_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackObjd4edec_TimerEvent
func callbackObjd4edec_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewObjFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Obj) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Objd4edec_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
