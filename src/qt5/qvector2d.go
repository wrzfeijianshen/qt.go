package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qvector2d.h
// dst-file: /src/gui/qvector2d.go
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
  // proto:  QPointF QVector2D::toPointF();
extern void _ZNK9QVector2D8toPointFEv(void* qthis);
  // proto:  void QVector2D::setX(float x);
extern void _ZN9QVector2D4setXEf(void* qthis, float arg0);
  // proto:  void QVector2D::QVector2D(const QVector4D & vector);
extern void* dector_ZN9QVector2DC1ERK9QVector4D(void* arg0);
extern void _ZN9QVector2DC1ERK9QVector4D(void* qthis, void* arg0);
  // proto:  QPoint QVector2D::toPoint();
extern void _ZNK9QVector2D7toPointEv(void* qthis);
  // proto:  float QVector2D::length();
extern void _ZNK9QVector2D6lengthEv(void* qthis);
  // proto:  void QVector2D::setY(float y);
extern void _ZN9QVector2D4setYEf(void* qthis, float arg0);
  // proto:  void QVector2D::QVector2D(const QPoint & point);
extern void* dector_ZN9QVector2DC1ERK6QPoint(void* arg0);
extern void _ZN9QVector2DC1ERK6QPoint(void* qthis, void* arg0);
  // proto:  void QVector2D::QVector2D(float xpos, float ypos);
extern void* dector_ZN9QVector2DC1Eff(float arg0, float arg1);
extern void _ZN9QVector2DC1Eff(void* qthis, float arg0, float arg1);
  // proto:  bool QVector2D::isNull();
extern void _ZNK9QVector2D6isNullEv(void* qthis);
  // proto:  float QVector2D::distanceToLine(const QVector2D & point, const QVector2D & direction);
extern void _ZNK9QVector2D14distanceToLineERKS_S1_(void* qthis, void* arg0, void* arg1);
  // proto:  void QVector2D::QVector2D();
extern void* dector_ZN9QVector2DC1Ev();
extern void _ZN9QVector2DC1Ev(void* qthis);
  // proto:  QVector3D QVector2D::toVector3D();
extern void _ZNK9QVector2D10toVector3DEv(void* qthis);
  // proto:  float QVector2D::lengthSquared();
extern void _ZNK9QVector2D13lengthSquaredEv(void* qthis);
  // proto:  float QVector2D::y();
extern void _ZNK9QVector2D1yEv(void* qthis);
  // proto:  void QVector2D::QVector2D(const QVector3D & vector);
extern void* dector_ZN9QVector2DC1ERK9QVector3D(void* arg0);
extern void _ZN9QVector2DC1ERK9QVector3D(void* qthis, void* arg0);
  // proto:  float QVector2D::x();
extern void _ZNK9QVector2D1xEv(void* qthis);
  // proto:  void QVector2D::QVector2D(const QPointF & point);
extern void* dector_ZN9QVector2DC1ERK7QPointF(void* arg0);
extern void _ZN9QVector2DC1ERK7QPointF(void* qthis, void* arg0);
  // proto:  float QVector2D::distanceToPoint(const QVector2D & point);
extern void _ZNK9QVector2D15distanceToPointERKS_(void* qthis, void* arg0);
  // proto:  QVector4D QVector2D::toVector4D();
extern void _ZNK9QVector2D10toVector4DEv(void* qthis);
  // proto:  QVector2D QVector2D::normalized();
extern void _ZNK9QVector2D10normalizedEv(void* qthis);
  // proto:  void QVector2D::normalize();
extern void _ZN9QVector2D9normalizeEv(void* qthis);
  // proto: static float QVector2D::dotProduct(const QVector2D & v1, const QVector2D & v2);
extern void _ZN9QVector2D10dotProductERKS_S1_(void* arg0, void* arg1);
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

// class sizeof(QVector2D)=8
type QVector2D struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QPointF QVector2D::toPointF();
func (this *QVector2D) toPointF(args ...interface{}) () {
  // toPointF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D8toPointFEv
  default:
    qtrt.ErrorResolve("QVector2D", "toPointF", args)
  }

}

  // proto:  void QVector2D::setX(float x);
func (this *QVector2D) setX(args ...interface{}) () {
  // setX(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2D4setXEf
  default:
    qtrt.ErrorResolve("QVector2D", "setX", args)
  }

}

  // proto:  void QVector2D::QVector2D(const QVector4D & vector);
func NewQVector2D(args ...interface{}) QVector2D {
  return QVector2D{}
}

  // proto:  QPoint QVector2D::toPoint();
func (this *QVector2D) toPoint(args ...interface{}) () {
  // toPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D7toPointEv
  default:
    qtrt.ErrorResolve("QVector2D", "toPoint", args)
  }

}

  // proto:  float QVector2D::length();
func (this *QVector2D) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D6lengthEv
  default:
    qtrt.ErrorResolve("QVector2D", "length", args)
  }

}

  // proto:  void QVector2D::setY(float y);
func (this *QVector2D) setY(args ...interface{}) () {
  // setY(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2D4setYEf
  default:
    qtrt.ErrorResolve("QVector2D", "setY", args)
  }

}

  // proto:  bool QVector2D::isNull();
func (this *QVector2D) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D6isNullEv
  default:
    qtrt.ErrorResolve("QVector2D", "isNull", args)
  }

}

  // proto:  float QVector2D::distanceToLine(const QVector2D & point, const QVector2D & direction);
func (this *QVector2D) distanceToLine(args ...interface{}) () {
  // distanceToLine(const class QVector2D &, const class QVector2D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[0][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D14distanceToLineERKS_S1_
  default:
    qtrt.ErrorResolve("QVector2D", "distanceToLine", args)
  }

}

  // proto:  QVector3D QVector2D::toVector3D();
func (this *QVector2D) toVector3D(args ...interface{}) () {
  // toVector3D()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D10toVector3DEv
  default:
    qtrt.ErrorResolve("QVector2D", "toVector3D", args)
  }

}

  // proto:  float QVector2D::lengthSquared();
func (this *QVector2D) lengthSquared(args ...interface{}) () {
  // lengthSquared()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D13lengthSquaredEv
  default:
    qtrt.ErrorResolve("QVector2D", "lengthSquared", args)
  }

}

  // proto:  float QVector2D::y();
func (this *QVector2D) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D1yEv
  default:
    qtrt.ErrorResolve("QVector2D", "y", args)
  }

}

  // proto:  float QVector2D::x();
func (this *QVector2D) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D1xEv
  default:
    qtrt.ErrorResolve("QVector2D", "x", args)
  }

}

  // proto:  float QVector2D::distanceToPoint(const QVector2D & point);
func (this *QVector2D) distanceToPoint(args ...interface{}) () {
  // distanceToPoint(const class QVector2D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D15distanceToPointERKS_
  default:
    qtrt.ErrorResolve("QVector2D", "distanceToPoint", args)
  }

}

  // proto:  QVector4D QVector2D::toVector4D();
func (this *QVector2D) toVector4D(args ...interface{}) () {
  // toVector4D()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D10toVector4DEv
  default:
    qtrt.ErrorResolve("QVector2D", "toVector4D", args)
  }

}

  // proto:  QVector2D QVector2D::normalized();
func (this *QVector2D) normalized(args ...interface{}) () {
  // normalized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D10normalizedEv
  default:
    qtrt.ErrorResolve("QVector2D", "normalized", args)
  }

}

  // proto:  void QVector2D::normalize();
func (this *QVector2D) normalize(args ...interface{}) () {
  // normalize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2D9normalizeEv
  default:
    qtrt.ErrorResolve("QVector2D", "normalize", args)
  }

}

  // proto: static float QVector2D::dotProduct(const QVector2D & v1, const QVector2D & v2);
func (this *QVector2D) dotProduct_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QVector2D", "dotProduct", args)
  }

}

// <= body block end
