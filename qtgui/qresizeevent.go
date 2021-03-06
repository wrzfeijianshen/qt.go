package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QResizeEvent struct {
	*qtcore.QEvent
}
type QResizeEvent_ITF interface {
	qtcore.QEvent_ITF
	QResizeEvent_PTR() *QResizeEvent
}

func (ptr *QResizeEvent) QResizeEvent_PTR() *QResizeEvent { return ptr }

func (this *QResizeEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QResizeEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQResizeEventFromPointer(cthis unsafe.Pointer) *QResizeEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QResizeEvent{bcthis0}
}
func (*QResizeEvent) NewFromPointer(cthis unsafe.Pointer) *QResizeEvent {
	return NewQResizeEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:463
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QResizeEvent(const QSize &, const QSize &)

/*

 */
func (*QResizeEvent) NewForInherit(size qtcore.QSize_ITF, oldSize qtcore.QSize_ITF) *QResizeEvent {
	return NewQResizeEvent(size, oldSize)
}
func NewQResizeEvent(size qtcore.QSize_ITF, oldSize qtcore.QSize_ITF) *QResizeEvent {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if oldSize != nil && oldSize.QSize_PTR() != nil {
		convArg1 = oldSize.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QResizeEventC2ERK5QSizeS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQResizeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQResizeEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:464
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QResizeEvent()

/*

 */
func DeleteQResizeEvent(this *QResizeEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QResizeEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:466
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QSize & size() const

/*

 */
func (this *QResizeEvent) Size() *qtcore.QSize {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QResizeEvent4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:467
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QSize & oldSize() const

/*

 */
func (this *QResizeEvent) OldSize() *qtcore.QSize {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QResizeEvent7oldSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
