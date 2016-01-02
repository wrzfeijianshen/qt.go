package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qcryptographichash.h
// dst-file: /src/core/qcryptographichash.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QCryptographicHash::addData(QIODevice * device);
extern void _ZN18QCryptographicHash7addDataEP9QIODevice(void* qthis, void* arg0);
  // proto:  void QCryptographicHash::~QCryptographicHash();
extern void _ZN18QCryptographicHashD0Ev(void* qthis);
  // proto:  void QCryptographicHash::reset();
extern void _ZN18QCryptographicHash5resetEv(void* qthis);
  // proto:  void QCryptographicHash::addData(const char * data, int length);
extern void _ZN18QCryptographicHash7addDataEPKci(void* qthis, char* arg0, int arg1);
  // proto:  QByteArray QCryptographicHash::result();
extern void _ZNK18QCryptographicHash6resultEv(void* qthis);
  // proto:  void QCryptographicHash::addData(const QByteArray & data);
extern void _ZN18QCryptographicHash7addDataERK10QByteArray(void* qthis, void* arg0);
  // proto:  void QCryptographicHash::QCryptographicHash(const QCryptographicHash & );
extern void* dector_ZN18QCryptographicHashC1ERKS_(void* arg0);
extern void _ZN18QCryptographicHashC1ERKS_(void* qthis, void* arg0);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QCryptographicHash)=8
type QCryptographicHash struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QCryptographicHash::addData(QIODevice * device);
func (this *QCryptographicHash) addData(args ...interface{}) () {
  // addData(class QIODevice *)
  // addData(const char *, int)
  // addData(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCryptographicHash7addDataEP9QIODevice
  case 1:
    // invoke: _ZN18QCryptographicHash7addDataEPKci
  case 2:
    // invoke: _ZN18QCryptographicHash7addDataERK10QByteArray
  default:
    qtrt.ErrorResolve("QCryptographicHash", "addData", args)
  }

}

  // proto:  void QCryptographicHash::~QCryptographicHash();
func (this *QCryptographicHash) FreeQCryptographicHash(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCryptographicHash", "~QCryptographicHash", args)
  }

}

  // proto:  void QCryptographicHash::reset();
func (this *QCryptographicHash) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCryptographicHash5resetEv
  default:
    qtrt.ErrorResolve("QCryptographicHash", "reset", args)
  }

}

  // proto:  QByteArray QCryptographicHash::result();
func (this *QCryptographicHash) result(args ...interface{}) () {
  // result()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCryptographicHash6resultEv
  default:
    qtrt.ErrorResolve("QCryptographicHash", "result", args)
  }

}

  // proto:  void QCryptographicHash::QCryptographicHash(const QCryptographicHash & );
func NewQCryptographicHash(args ...interface{}) QCryptographicHash {
  return QCryptographicHash{}
}

// <= body block end
