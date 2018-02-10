package qtwidgets

// /usr/include/qt/QtWidgets/qcolormap.h
// #include <qcolormap.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin

type QColormap struct {
	*qtrt.CObject
}

func (this *QColormap) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QColormap) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQColormapFromPointer(cthis unsafe.Pointer) *QColormap {
	return &QColormap{&qtrt.CObject{cthis}}
}
func (*QColormap) NewFromPointer(cthis unsafe.Pointer) *QColormap {
	return NewQColormapFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qcolormap.h:60
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void initialize()
func (this *QColormap) Initialize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QColormap10initializeEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QColormap_Initialize() {
	var nilthis *QColormap
	nilthis.Initialize()
}

// /usr/include/qt/QtWidgets/qcolormap.h:61
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void cleanup()
func (this *QColormap) Cleanup() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QColormap7cleanupEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QColormap_Cleanup() {
	var nilthis *QColormap
	nilthis.Cleanup()
}

// /usr/include/qt/QtWidgets/qcolormap.h:63
// index:0
// Public static Visibility=Default Availability=Available
// [8] QColormap instance(int)
func (this *QColormap) Instance(screen int) *QColormap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QColormap8instanceEi", qtrt.FFI_TYPE_POINTER, screen)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColormapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColormap)
	return rv2
}
func QColormap_Instance(screen int) *QColormap /*123*/ {
	var nilthis *QColormap
	rv := nilthis.Instance(screen)
	return rv
}

// /usr/include/qt/QtWidgets/qcolormap.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QColormap()
func DeleteQColormap(this *QColormap) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QColormapD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qcolormap.h:70
// index:0
// Public Visibility=Default Availability=Available
// [4] QColormap::Mode mode()
func (this *QColormap) Mode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QColormap4modeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcolormap.h:72
// index:0
// Public Visibility=Default Availability=Available
// [4] int depth()
func (this *QColormap) Depth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QColormap5depthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcolormap.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] int size()
func (this *QColormap) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QColormap4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcolormap.h:75
// index:0
// Public Visibility=Default Availability=Available
// [4] uint pixel(const QColor &)
func (this *QColormap) Pixel(color *qtgui.QColor) uint {
	var convArg0 = color.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QColormap5pixelERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtWidgets/qcolormap.h:76
// index:0
// Public Visibility=Default Availability=Available
// [16] const QColor colorAt(uint)
func (this *QColormap) ColorAt(pixel uint) *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QColormap7colorAtEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pixel)
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

type QColormap__Mode = int

const QColormap__Direct QColormap__Mode = 0
const QColormap__Indexed QColormap__Mode = 1
const QColormap__Gray QColormap__Mode = 2

//  body block end