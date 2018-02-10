package qtgui

// /usr/include/qt/QtGui/qpaintengine.h
// #include <qpaintengine.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 203
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin

type QTextItem struct {
	*qtrt.CObject
}

func (this *QTextItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextItemFromPointer(cthis unsafe.Pointer) *QTextItem {
	return &QTextItem{&qtrt.CObject{cthis}}
}
func (*QTextItem) NewFromPointer(cthis unsafe.Pointer) *QTextItem {
	return NewQTextItemFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpaintengine.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal descent()
func (this *QTextItem) Descent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextItem7descentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpaintengine.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal ascent()
func (this *QTextItem) Ascent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextItem6ascentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpaintengine.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal width()
func (this *QTextItem) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextItem5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpaintengine.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextItem::RenderFlags renderFlags()
func (this *QTextItem) RenderFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextItem11renderFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QTextItem) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextItem4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qpaintengine.h:81
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font()
func (this *QTextItem) Font() *QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextItem4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

func DeleteQTextItem(this *QTextItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QTextItem__RenderFlag = int

const QTextItem__RightToLeft QTextItem__RenderFlag = 1
const QTextItem__Overline QTextItem__RenderFlag = 16
const QTextItem__Underline QTextItem__RenderFlag = 32
const QTextItem__StrikeOut QTextItem__RenderFlag = 64
const QTextItem__Dummy QTextItem__RenderFlag = -1

//  body block end