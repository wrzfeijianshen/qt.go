package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qdatawidgetmapper.h
// dst-file: /src/widgets/qdatawidgetmapper.go
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
  // proto:  void QDataWidgetMapper::QDataWidgetMapper(const QDataWidgetMapper & );
extern void* dector_ZN17QDataWidgetMapperC1ERKS_(void* arg0);
extern void _ZN17QDataWidgetMapperC1ERKS_(void* qthis, void* arg0);
  // proto:  int QDataWidgetMapper::currentIndex();
extern void _ZNK17QDataWidgetMapper12currentIndexEv(void* qthis);
  // proto:  void QDataWidgetMapper::addMapping(QWidget * widget, int section, const QByteArray & propertyName);
extern void _ZN17QDataWidgetMapper10addMappingEP7QWidgetiRK10QByteArray(void* qthis, void* arg0, int arg1, void* arg2);
  // proto:  QModelIndex QDataWidgetMapper::rootIndex();
extern void _ZNK17QDataWidgetMapper9rootIndexEv(void* qthis);
  // proto:  void QDataWidgetMapper::setCurrentIndex(int index);
extern void _ZN17QDataWidgetMapper15setCurrentIndexEi(void* qthis, int arg0);
  // proto:  void QDataWidgetMapper::setModel(QAbstractItemModel * model);
extern void _ZN17QDataWidgetMapper8setModelEP18QAbstractItemModel(void* qthis, void* arg0);
  // proto:  QWidget * QDataWidgetMapper::mappedWidgetAt(int section);
extern void _ZNK17QDataWidgetMapper14mappedWidgetAtEi(void* qthis, int arg0);
  // proto:  void QDataWidgetMapper::removeMapping(QWidget * widget);
extern void _ZN17QDataWidgetMapper13removeMappingEP7QWidget(void* qthis, void* arg0);
  // proto:  void QDataWidgetMapper::toFirst();
extern void _ZN17QDataWidgetMapper7toFirstEv(void* qthis);
  // proto:  void QDataWidgetMapper::toPrevious();
extern void _ZN17QDataWidgetMapper10toPreviousEv(void* qthis);
  // proto:  void QDataWidgetMapper::setRootIndex(const QModelIndex & index);
extern void _ZN17QDataWidgetMapper12setRootIndexERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QDataWidgetMapper::revert();
extern void _ZN17QDataWidgetMapper6revertEv(void* qthis);
  // proto:  void QDataWidgetMapper::clearMapping();
extern void _ZN17QDataWidgetMapper12clearMappingEv(void* qthis);
  // proto:  int QDataWidgetMapper::mappedSection(QWidget * widget);
extern void _ZNK17QDataWidgetMapper13mappedSectionEP7QWidget(void* qthis, void* arg0);
  // proto:  QByteArray QDataWidgetMapper::mappedPropertyName(QWidget * widget);
extern void _ZNK17QDataWidgetMapper18mappedPropertyNameEP7QWidget(void* qthis, void* arg0);
  // proto:  void QDataWidgetMapper::setItemDelegate(QAbstractItemDelegate * delegate);
extern void _ZN17QDataWidgetMapper15setItemDelegateEP21QAbstractItemDelegate(void* qthis, void* arg0);
  // proto:  void QDataWidgetMapper::setCurrentModelIndex(const QModelIndex & index);
extern void _ZN17QDataWidgetMapper20setCurrentModelIndexERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QDataWidgetMapper::~QDataWidgetMapper();
extern void _ZN17QDataWidgetMapperD0Ev(void* qthis);
  // proto:  void QDataWidgetMapper::addMapping(QWidget * widget, int section);
extern void _ZN17QDataWidgetMapper10addMappingEP7QWidgeti(void* qthis, void* arg0, int arg1);
  // proto:  void QDataWidgetMapper::QDataWidgetMapper(QObject * parent);
extern void* dector_ZN17QDataWidgetMapperC1EP7QObject(void* arg0);
extern void _ZN17QDataWidgetMapperC1EP7QObject(void* qthis, void* arg0);
  // proto:  const QMetaObject * QDataWidgetMapper::metaObject();
extern void _ZNK17QDataWidgetMapper10metaObjectEv(void* qthis);
  // proto:  void QDataWidgetMapper::toLast();
extern void _ZN17QDataWidgetMapper6toLastEv(void* qthis);
  // proto:  QAbstractItemModel * QDataWidgetMapper::model();
extern void _ZNK17QDataWidgetMapper5modelEv(void* qthis);
  // proto:  QAbstractItemDelegate * QDataWidgetMapper::itemDelegate();
extern void _ZNK17QDataWidgetMapper12itemDelegateEv(void* qthis);
  // proto:  bool QDataWidgetMapper::submit();
extern void _ZN17QDataWidgetMapper6submitEv(void* qthis);
  // proto:  void QDataWidgetMapper::toNext();
extern void _ZN17QDataWidgetMapper6toNextEv(void* qthis);
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

// class sizeof(QDataWidgetMapper)=1
type QDataWidgetMapper struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _currentIndexChanged QDataWidgetMapper_currentIndexChanged_signal;
}

  // proto:  void QDataWidgetMapper::QDataWidgetMapper(const QDataWidgetMapper & );
func NewQDataWidgetMapper(args ...interface{}) QDataWidgetMapper {
  return QDataWidgetMapper{}
}

  // proto:  int QDataWidgetMapper::currentIndex();
func (this *QDataWidgetMapper) currentIndex(args ...interface{}) () {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper12currentIndexEv
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "currentIndex", args)
  }

}

  // proto:  void QDataWidgetMapper::addMapping(QWidget * widget, int section, const QByteArray & propertyName);
func (this *QDataWidgetMapper) addMapping(args ...interface{}) () {
  // addMapping(class QWidget *, int, const class QByteArray &)
  // addMapping(class QWidget *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper10addMappingEP7QWidgetiRK10QByteArray
  case 1:
    // invoke: _ZN17QDataWidgetMapper10addMappingEP7QWidgeti
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "addMapping", args)
  }

}

  // proto:  QModelIndex QDataWidgetMapper::rootIndex();
func (this *QDataWidgetMapper) rootIndex(args ...interface{}) () {
  // rootIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper9rootIndexEv
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "rootIndex", args)
  }

}

  // proto:  void QDataWidgetMapper::setCurrentIndex(int index);
func (this *QDataWidgetMapper) setCurrentIndex(args ...interface{}) () {
  // setCurrentIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper15setCurrentIndexEi
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setCurrentIndex", args)
  }

}

  // proto:  void QDataWidgetMapper::setModel(QAbstractItemModel * model);
func (this *QDataWidgetMapper) setModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper8setModelEP18QAbstractItemModel
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setModel", args)
  }

}

  // proto:  QWidget * QDataWidgetMapper::mappedWidgetAt(int section);
func (this *QDataWidgetMapper) mappedWidgetAt(args ...interface{}) () {
  // mappedWidgetAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper14mappedWidgetAtEi
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "mappedWidgetAt", args)
  }

}

  // proto:  void QDataWidgetMapper::removeMapping(QWidget * widget);
func (this *QDataWidgetMapper) removeMapping(args ...interface{}) () {
  // removeMapping(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper13removeMappingEP7QWidget
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "removeMapping", args)
  }

}

  // proto:  void QDataWidgetMapper::toFirst();
func (this *QDataWidgetMapper) toFirst(args ...interface{}) () {
  // toFirst()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper7toFirstEv
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toFirst", args)
  }

}

  // proto:  void QDataWidgetMapper::toPrevious();
func (this *QDataWidgetMapper) toPrevious(args ...interface{}) () {
  // toPrevious()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper10toPreviousEv
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toPrevious", args)
  }

}

  // proto:  void QDataWidgetMapper::setRootIndex(const QModelIndex & index);
func (this *QDataWidgetMapper) setRootIndex(args ...interface{}) () {
  // setRootIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper12setRootIndexERK11QModelIndex
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setRootIndex", args)
  }

}

  // proto:  void QDataWidgetMapper::revert();
func (this *QDataWidgetMapper) revert(args ...interface{}) () {
  // revert()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper6revertEv
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "revert", args)
  }

}

  // proto:  void QDataWidgetMapper::clearMapping();
func (this *QDataWidgetMapper) clearMapping(args ...interface{}) () {
  // clearMapping()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper12clearMappingEv
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "clearMapping", args)
  }

}

  // proto:  int QDataWidgetMapper::mappedSection(QWidget * widget);
func (this *QDataWidgetMapper) mappedSection(args ...interface{}) () {
  // mappedSection(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper13mappedSectionEP7QWidget
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "mappedSection", args)
  }

}

  // proto:  QByteArray QDataWidgetMapper::mappedPropertyName(QWidget * widget);
func (this *QDataWidgetMapper) mappedPropertyName(args ...interface{}) () {
  // mappedPropertyName(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper18mappedPropertyNameEP7QWidget
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "mappedPropertyName", args)
  }

}

  // proto:  void QDataWidgetMapper::setItemDelegate(QAbstractItemDelegate * delegate);
func (this *QDataWidgetMapper) setItemDelegate(args ...interface{}) () {
  // setItemDelegate(class QAbstractItemDelegate *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemDelegate{}) // "QAbstractItemDelegate *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper15setItemDelegateEP21QAbstractItemDelegate
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setItemDelegate", args)
  }

}

  // proto:  void QDataWidgetMapper::setCurrentModelIndex(const QModelIndex & index);
func (this *QDataWidgetMapper) setCurrentModelIndex(args ...interface{}) () {
  // setCurrentModelIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper20setCurrentModelIndexERK11QModelIndex
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setCurrentModelIndex", args)
  }

}

  // proto:  void QDataWidgetMapper::~QDataWidgetMapper();
func (this *QDataWidgetMapper) FreeQDataWidgetMapper(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "~QDataWidgetMapper", args)
  }

}

  // proto:  const QMetaObject * QDataWidgetMapper::metaObject();
func (this *QDataWidgetMapper) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper10metaObjectEv
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "metaObject", args)
  }

}

  // proto:  void QDataWidgetMapper::toLast();
func (this *QDataWidgetMapper) toLast(args ...interface{}) () {
  // toLast()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper6toLastEv
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toLast", args)
  }

}

  // proto:  QAbstractItemModel * QDataWidgetMapper::model();
func (this *QDataWidgetMapper) model(args ...interface{}) () {
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper5modelEv
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "model", args)
  }

}

  // proto:  QAbstractItemDelegate * QDataWidgetMapper::itemDelegate();
func (this *QDataWidgetMapper) itemDelegate(args ...interface{}) () {
  // itemDelegate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper12itemDelegateEv
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "itemDelegate", args)
  }

}

  // proto:  bool QDataWidgetMapper::submit();
func (this *QDataWidgetMapper) submit(args ...interface{}) () {
  // submit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper6submitEv
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "submit", args)
  }

}

  // proto:  void QDataWidgetMapper::toNext();
func (this *QDataWidgetMapper) toNext(args ...interface{}) () {
  // toNext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper6toNextEv
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toNext", args)
  }

}

// <= body block end
