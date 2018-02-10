package qtwidgets

// /usr/include/qt/QtWidgets/qfontcombobox.h
// #include <qfontcombobox.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
// bool event(class QEvent *)
func (this *QFontComboBox) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QFontComboBox struct {
	*QComboBox
}

func (this *QFontComboBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QComboBox.GetCthis()
	}
}
func (this *QFontComboBox) SetCthis(cthis unsafe.Pointer) {
	this.QComboBox = NewQComboBoxFromPointer(cthis)
}
func NewQFontComboBoxFromPointer(cthis unsafe.Pointer) *QFontComboBox {
	bcthis0 := NewQComboBoxFromPointer(cthis)
	return &QFontComboBox{bcthis0}
}
func (*QFontComboBox) NewFromPointer(cthis unsafe.Pointer) *QFontComboBox {
	return NewQFontComboBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QFontComboBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontComboBox10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFontComboBox(QWidget *)
func NewQFontComboBox(parent *QWidget /*777 QWidget **/) *QFontComboBox {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontComboBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontComboBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFontComboBox()
func DeleteQFontComboBox(this *QFontComboBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontComboBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWritingSystem(QFontDatabase::WritingSystem)
func (this *QFontComboBox) SetWritingSystem(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontComboBox16setWritingSystemEN13QFontDatabase13WritingSystemE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:65
// index:0
// Public Visibility=Default Availability=Available
// [4] QFontDatabase::WritingSystem writingSystem()
func (this *QFontComboBox) WritingSystem() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontComboBox13writingSystemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFontFilters(QFontComboBox::FontFilters)
func (this *QFontComboBox) SetFontFilters(filters int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontComboBox14setFontFiltersE6QFlagsINS_10FontFilterEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] QFontComboBox::FontFilters fontFilters()
func (this *QFontComboBox) FontFilters() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontComboBox11fontFiltersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:80
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont currentFont()
func (this *QFontComboBox) CurrentFont() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontComboBox11currentFontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QFontComboBox) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontComboBox8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentFont(const QFont &)
func (this *QFontComboBox) SetCurrentFont(f *qtgui.QFont) {
	var convArg0 = f.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontComboBox14setCurrentFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentFontChanged(const QFont &)
func (this *QFontComboBox) CurrentFontChanged(f *qtgui.QFont) {
	var convArg0 = f.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontComboBox18currentFontChangedERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:90
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QFontComboBox) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontComboBox5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

type QFontComboBox__FontFilter = int

const QFontComboBox__AllFonts QFontComboBox__FontFilter = 0
const QFontComboBox__ScalableFonts QFontComboBox__FontFilter = 1
const QFontComboBox__NonScalableFonts QFontComboBox__FontFilter = 2
const QFontComboBox__MonospacedFonts QFontComboBox__FontFilter = 4
const QFontComboBox__ProportionalFonts QFontComboBox__FontFilter = 8

//  body block end