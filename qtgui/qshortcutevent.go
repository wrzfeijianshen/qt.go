package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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

type QShortcutEvent struct {
	*qtcore.QEvent
}

func (this *QShortcutEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QShortcutEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQShortcutEventFromPointer(cthis unsafe.Pointer) *QShortcutEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QShortcutEvent{bcthis0}
}
func (*QShortcutEvent) NewFromPointer(cthis unsafe.Pointer) *QShortcutEvent {
	return NewQShortcutEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:767
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QShortcutEvent(const QKeySequence &, int, _Bool)
func NewQShortcutEvent(key *QKeySequence, id int, ambiguous bool) *QShortcutEvent {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QShortcutEventC2ERK12QKeySequenceib", qtrt.FFI_TYPE_POINTER, convArg0, id, ambiguous)
	gopp.ErrPrint(err, rv)
	gothis := NewQShortcutEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQShortcutEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:768
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QShortcutEvent()
func DeleteQShortcutEvent(this *QShortcutEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QShortcutEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:770
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QKeySequence & key()
func (this *QShortcutEvent) Key() *QKeySequence {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QShortcutEvent3keyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQKeySequence)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:771
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int shortcutId()
func (this *QShortcutEvent) ShortcutId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QShortcutEvent10shortcutIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:772
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAmbiguous()
func (this *QShortcutEvent) IsAmbiguous() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QShortcutEvent11isAmbiguousEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

//  body block end