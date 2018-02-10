package qtcore

// /usr/include/qt/QtCore/qeventloop.h
// #include <qeventloop.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin

type QEventLoop struct {
	*QObject
}

func (this *QEventLoop) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QEventLoop) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQEventLoopFromPointer(cthis unsafe.Pointer) *QEventLoop {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QEventLoop{bcthis0}
}
func (*QEventLoop) NewFromPointer(cthis unsafe.Pointer) *QEventLoop {
	return NewQEventLoopFromPointer(cthis)
}

// /usr/include/qt/QtCore/qeventloop.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QEventLoop) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QEventLoop10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qeventloop.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QEventLoop(QObject *)
func NewQEventLoop(parent *QObject /*777 QObject **/) *QEventLoop {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoopC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQEventLoopFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qeventloop.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QEventLoop()
func DeleteQEventLoop(this *QEventLoop) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoopD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qeventloop.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool processEvents(QEventLoop::ProcessEventsFlags)
func (this *QEventLoop) ProcessEvents(flags int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop13processEventsE6QFlagsINS_17ProcessEventsFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qeventloop.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void processEvents(QEventLoop::ProcessEventsFlags, int)
func (this *QEventLoop) ProcessEvents_1(flags int, maximumTime int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop13processEventsE6QFlagsINS_17ProcessEventsFlagEEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags, maximumTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventloop.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] int exec(QEventLoop::ProcessEventsFlags)
func (this *QEventLoop) Exec(flags int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop4execE6QFlagsINS_17ProcessEventsFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qeventloop.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void exit(int)
func (this *QEventLoop) Exit(returnCode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop4exitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), returnCode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventloop.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRunning()
func (this *QEventLoop) IsRunning() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QEventLoop9isRunningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qeventloop.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void wakeUp()
func (this *QEventLoop) WakeUp() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop6wakeUpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventloop.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QEventLoop) Event(event *QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qeventloop.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void quit()
func (this *QEventLoop) Quit() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop4quitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

type QEventLoop__ProcessEventsFlag = int

const QEventLoop__AllEvents QEventLoop__ProcessEventsFlag = 0
const QEventLoop__ExcludeUserInputEvents QEventLoop__ProcessEventsFlag = 1
const QEventLoop__ExcludeSocketNotifiers QEventLoop__ProcessEventsFlag = 2
const QEventLoop__WaitForMoreEvents QEventLoop__ProcessEventsFlag = 4
const QEventLoop__X11ExcludeTimers QEventLoop__ProcessEventsFlag = 8
const QEventLoop__EventLoopExec QEventLoop__ProcessEventsFlag = 32
const QEventLoop__DialogExec QEventLoop__ProcessEventsFlag = 64

//  body block end