package qtcore

// /usr/include/qt/QtCore/qlinkedlist.h
// #include <qlinkedlist.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
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

type QLinkedListData struct {
	*qtrt.CObject
}

func (this *QLinkedListData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLinkedListData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQLinkedListDataFromPointer(cthis unsafe.Pointer) *QLinkedListData {
	return &QLinkedListData{&qtrt.CObject{cthis}}
}
func (*QLinkedListData) NewFromPointer(cthis unsafe.Pointer) *QLinkedListData {
	return NewQLinkedListDataFromPointer(cthis)
}

func DeleteQLinkedListData(this *QLinkedListData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QLinkedListDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end