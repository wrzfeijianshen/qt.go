package qtwidgets

// /usr/include/qt/QtWidgets/qtooltip.h
// #include <qtooltip.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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

type QToolTip struct {
	*qtrt.CObject
}

func (this *QToolTip) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QToolTip) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQToolTipFromPointer(cthis unsafe.Pointer) *QToolTip {
	return &QToolTip{&qtrt.CObject{cthis}}
}
func (*QToolTip) NewFromPointer(cthis unsafe.Pointer) *QToolTip {
	return NewQToolTipFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtooltip.h:56
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void showText(const QPoint &, const QString &, QWidget *)
func (this *QToolTip) ShowText(pos *qtcore.QPoint, text string, w *QWidget /*777 QWidget **/) {
	var convArg0 = pos.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = w.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}
func QToolTip_ShowText(pos *qtcore.QPoint, text string, w *QWidget /*777 QWidget **/) {
	var nilthis *QToolTip
	nilthis.ShowText(pos, text, w)
}

// /usr/include/qt/QtWidgets/qtooltip.h:57
// index:1
// Public static Visibility=Default Availability=Available
// [-2] void showText(const QPoint &, const QString &, QWidget *, const QRect &)
func (this *QToolTip) ShowText_1(pos *qtcore.QPoint, text string, w *QWidget /*777 QWidget **/, rect *qtcore.QRect) {
	var convArg0 = pos.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = w.GetCthis()
	var convArg3 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidgetRK5QRect", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}
func QToolTip_ShowText_1(pos *qtcore.QPoint, text string, w *QWidget /*777 QWidget **/, rect *qtcore.QRect) {
	var nilthis *QToolTip
	nilthis.ShowText_1(pos, text, w, rect)
}

// /usr/include/qt/QtWidgets/qtooltip.h:58
// index:2
// Public static Visibility=Default Availability=Available
// [-2] void showText(const QPoint &, const QString &, QWidget *, const QRect &, int)
func (this *QToolTip) ShowText_2(pos *qtcore.QPoint, text string, w *QWidget /*777 QWidget **/, rect *qtcore.QRect, msecShowTime int) {
	var convArg0 = pos.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = w.GetCthis()
	var convArg3 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidgetRK5QRecti", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, msecShowTime)
	gopp.ErrPrint(err, rv)
}
func QToolTip_ShowText_2(pos *qtcore.QPoint, text string, w *QWidget /*777 QWidget **/, rect *qtcore.QRect, msecShowTime int) {
	var nilthis *QToolTip
	nilthis.ShowText_2(pos, text, w, rect, msecShowTime)
}

// /usr/include/qt/QtWidgets/qtooltip.h:59
// index:0
// Public static inline Visibility=Default Availability=Available
// [-2] void hideText()
func (this *QToolTip) HideText() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolTip8hideTextEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QToolTip_HideText() {
	var nilthis *QToolTip
	nilthis.HideText()
}

// /usr/include/qt/QtWidgets/qtooltip.h:61
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isVisible()
func (this *QToolTip) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolTip9isVisibleEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QToolTip_IsVisible() bool {
	var nilthis *QToolTip
	rv := nilthis.IsVisible()
	return rv
}

// /usr/include/qt/QtWidgets/qtooltip.h:62
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString text()
func (this *QToolTip) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolTip4textEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QToolTip_Text() string {
	var nilthis *QToolTip
	rv := nilthis.Text()
	return rv
}

// /usr/include/qt/QtWidgets/qtooltip.h:64
// index:0
// Public static Visibility=Default Availability=Available
// [16] QPalette palette()
func (this *QToolTip) Palette() *qtgui.QPalette /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolTip7paletteEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPalette)
	return rv2
}
func QToolTip_Palette() *qtgui.QPalette /*123*/ {
	var nilthis *QToolTip
	rv := nilthis.Palette()
	return rv
}

// /usr/include/qt/QtWidgets/qtooltip.h:65
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setPalette(const QPalette &)
func (this *QToolTip) SetPalette(arg0 *qtgui.QPalette) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolTip10setPaletteERK8QPalette", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QToolTip_SetPalette(arg0 *qtgui.QPalette) {
	var nilthis *QToolTip
	nilthis.SetPalette(arg0)
}

// /usr/include/qt/QtWidgets/qtooltip.h:66
// index:0
// Public static Visibility=Default Availability=Available
// [16] QFont font()
func (this *QToolTip) Font() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolTip4fontEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}
func QToolTip_Font() *qtgui.QFont /*123*/ {
	var nilthis *QToolTip
	rv := nilthis.Font()
	return rv
}

// /usr/include/qt/QtWidgets/qtooltip.h:67
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)
func (this *QToolTip) SetFont(arg0 *qtgui.QFont) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolTip7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QToolTip_SetFont(arg0 *qtgui.QFont) {
	var nilthis *QToolTip
	nilthis.SetFont(arg0)
}

func DeleteQToolTip(this *QToolTip) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolTipD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end