package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qbuffer.h
// dst-file: /src/core/qbuffer.go
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
  // proto:  bool QBuffer::seek(qint64 off);
extern void _ZN7QBuffer4seekEx(void* qthis, long long arg0);
  // proto:  bool QBuffer::canReadLine();
extern void _ZNK7QBuffer11canReadLineEv(void* qthis);
  // proto:  void QBuffer::~QBuffer();
extern void _ZN7QBufferD0Ev(void* qthis);
  // proto:  void QBuffer::setData(const QByteArray & data);
extern void _ZN7QBuffer7setDataERK10QByteArray(void* qthis, void* arg0);
  // proto:  const QByteArray & QBuffer::data();
extern void _ZNK7QBuffer4dataEv(void* qthis);
  // proto:  void QBuffer::QBuffer(QObject * parent);
extern void* dector_ZN7QBufferC1EP7QObject(void* arg0);
extern void _ZN7QBufferC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QBuffer::setBuffer(QByteArray * a);
extern void _ZN7QBuffer9setBufferEP10QByteArray(void* qthis, void* arg0);
  // proto:  QByteArray & QBuffer::buffer();
extern void _ZN7QBuffer6bufferEv(void* qthis);
  // proto:  qint64 QBuffer::pos();
extern void _ZNK7QBuffer3posEv(void* qthis);
  // proto:  void QBuffer::QBuffer(const QBuffer & );
extern void* dector_ZN7QBufferC1ERKS_(void* arg0);
extern void _ZN7QBufferC1ERKS_(void* qthis, void* arg0);
  // proto:  void QBuffer::close();
extern void _ZN7QBuffer5closeEv(void* qthis);
  // proto:  const QMetaObject * QBuffer::metaObject();
extern void _ZNK7QBuffer10metaObjectEv(void* qthis);
  // proto:  qint64 QBuffer::size();
extern void _ZNK7QBuffer4sizeEv(void* qthis);
  // proto:  void QBuffer::QBuffer(QByteArray * buf, QObject * parent);
extern void* dector_ZN7QBufferC1EP10QByteArrayP7QObject(void* arg0, void* arg1);
extern void _ZN7QBufferC1EP10QByteArrayP7QObject(void* qthis, void* arg0, void* arg1);
  // proto:  bool QBuffer::atEnd();
extern void _ZNK7QBuffer5atEndEv(void* qthis);
  // proto:  void QBuffer::setData(const char * data, int len);
extern void demth_ZN7QBuffer7setDataEPKci(void* qthis, char* arg0, int arg1);
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

// class sizeof(QBuffer)=1
type QBuffer struct {
  /*qbase*/ QIODevice;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QBuffer::seek(qint64 off);
func (this *QBuffer) seek(args ...interface{}) () {
  // seek(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBuffer4seekEx
  default:
    qtrt.ErrorResolve("QBuffer", "seek", args)
  }

}

  // proto:  bool QBuffer::canReadLine();
func (this *QBuffer) canReadLine(args ...interface{}) () {
  // canReadLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer11canReadLineEv
  default:
    qtrt.ErrorResolve("QBuffer", "canReadLine", args)
  }

}

  // proto:  void QBuffer::~QBuffer();
func (this *QBuffer) FreeQBuffer(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QBuffer", "~QBuffer", args)
  }

}

  // proto:  void QBuffer::setData(const QByteArray & data);
func (this *QBuffer) setData(args ...interface{}) () {
  // setData(const class QByteArray &)
  // setData(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBuffer7setDataERK10QByteArray
  case 1:
    // invoke: _ZN7QBuffer7setDataEPKci
  default:
    qtrt.ErrorResolve("QBuffer", "setData", args)
  }

}

  // proto:  const QByteArray & QBuffer::data();
func (this *QBuffer) data(args ...interface{}) () {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer4dataEv
  default:
    qtrt.ErrorResolve("QBuffer", "data", args)
  }

}

  // proto:  void QBuffer::QBuffer(QObject * parent);
func NewQBuffer(args ...interface{}) QBuffer {
  return QBuffer{}
}

  // proto:  void QBuffer::setBuffer(QByteArray * a);
func (this *QBuffer) setBuffer(args ...interface{}) () {
  // setBuffer(class QByteArray *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "QByteArray *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBuffer9setBufferEP10QByteArray
  default:
    qtrt.ErrorResolve("QBuffer", "setBuffer", args)
  }

}

  // proto:  QByteArray & QBuffer::buffer();
func (this *QBuffer) buffer(args ...interface{}) () {
  // buffer()
  // buffer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer6bufferEv
  case 1:
    // invoke: _ZN7QBuffer6bufferEv
  default:
    qtrt.ErrorResolve("QBuffer", "buffer", args)
  }

}

  // proto:  qint64 QBuffer::pos();
func (this *QBuffer) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer3posEv
  default:
    qtrt.ErrorResolve("QBuffer", "pos", args)
  }

}

  // proto:  void QBuffer::close();
func (this *QBuffer) close(args ...interface{}) () {
  // close()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBuffer5closeEv
  default:
    qtrt.ErrorResolve("QBuffer", "close", args)
  }

}

  // proto:  const QMetaObject * QBuffer::metaObject();
func (this *QBuffer) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer10metaObjectEv
  default:
    qtrt.ErrorResolve("QBuffer", "metaObject", args)
  }

}

  // proto:  qint64 QBuffer::size();
func (this *QBuffer) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer4sizeEv
  default:
    qtrt.ErrorResolve("QBuffer", "size", args)
  }

}

  // proto:  bool QBuffer::atEnd();
func (this *QBuffer) atEnd(args ...interface{}) () {
  // atEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer5atEndEv
  default:
    qtrt.ErrorResolve("QBuffer", "atEnd", args)
  }

}

// <= body block end
