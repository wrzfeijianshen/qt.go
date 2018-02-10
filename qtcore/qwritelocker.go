package qtcore

// /usr/include/qt/QtCore/qreadwritelock.h
// #include <qreadwritelock.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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

type QWriteLocker struct {
	*qtrt.CObject
}

func (this *QWriteLocker) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWriteLocker) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWriteLockerFromPointer(cthis unsafe.Pointer) *QWriteLocker {
	return &QWriteLocker{&qtrt.CObject{cthis}}
}
func (*QWriteLocker) NewFromPointer(cthis unsafe.Pointer) *QWriteLocker {
	return NewQWriteLockerFromPointer(cthis)
}

// /usr/include/qt/QtCore/qreadwritelock.h:131
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QWriteLocker(QReadWriteLock *)
func NewQWriteLocker(readWriteLock *QReadWriteLock /*777 QReadWriteLock **/) *QWriteLocker {
	var convArg0 = readWriteLock.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QWriteLockerC2EP14QReadWriteLock", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQWriteLockerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWriteLocker)
	return gothis
}

// /usr/include/qt/QtCore/qreadwritelock.h:133
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QWriteLocker()
func DeleteQWriteLocker(this *QWriteLocker) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QWriteLockerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qreadwritelock.h:136
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void unlock()
func (this *QWriteLocker) Unlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QWriteLocker6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:146
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void relock()
func (this *QWriteLocker) Relock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QWriteLocker6relockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:156
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QReadWriteLock * readWriteLock()
func (this *QWriteLocker) ReadWriteLock() *QReadWriteLock /*777 QReadWriteLock **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QWriteLocker13readWriteLockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQReadWriteLockFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end