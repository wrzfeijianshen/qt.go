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
// extern C begin: 7
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
type QDragEnterEvent struct {
	*QDragMoveEvent
}
type QDragEnterEvent_ITF interface {
	QDragMoveEvent_ITF
	QDragEnterEvent_PTR() *QDragEnterEvent
}

func (ptr *QDragEnterEvent) QDragEnterEvent_PTR() *QDragEnterEvent { return ptr }

func (this *QDragEnterEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDragMoveEvent.GetCthis()
	}
}
func (this *QDragEnterEvent) SetCthis(cthis unsafe.Pointer) {
	this.QDragMoveEvent = NewQDragMoveEventFromPointer(cthis)
}
func NewQDragEnterEventFromPointer(cthis unsafe.Pointer) *QDragEnterEvent {
	bcthis0 := NewQDragMoveEventFromPointer(cthis)
	return &QDragEnterEvent{bcthis0}
}
func (*QDragEnterEvent) NewFromPointer(cthis unsafe.Pointer) *QDragEnterEvent {
	return NewQDragEnterEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:662
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDragEnterEvent(const QPoint &, Qt::DropActions, const QMimeData *, Qt::MouseButtons, Qt::KeyboardModifiers)

/*

 */
func (*QDragEnterEvent) NewForInherit(pos qtcore.QPoint_ITF, actions int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, buttons int, modifiers int) *QDragEnterEvent {
	return NewQDragEnterEvent(pos, actions, data, buttons, modifiers)
}
func NewQDragEnterEvent(pos qtcore.QPoint_ITF, actions int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, buttons int, modifiers int) *QDragEnterEvent {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg2 = data.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QDragEnterEventC2ERK6QPoint6QFlagsIN2Qt10DropActionEEPK9QMimeDataS3_INS4_11MouseButtonEES3_INS4_16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, convArg0, actions, convArg2, buttons, modifiers)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDragEnterEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDragEnterEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:664
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDragEnterEvent()

/*

 */
func DeleteQDragEnterEvent(this *QDragEnterEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QDragEnterEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 88)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
