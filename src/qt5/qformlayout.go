package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qformlayout.h
// dst-file: /src/widgets/qformlayout.go
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
  // proto:  int QFormLayout::horizontalSpacing();
extern void _ZNK11QFormLayout17horizontalSpacingEv(void* qthis);
  // proto:  int QFormLayout::rowCount();
extern void _ZNK11QFormLayout8rowCountEv(void* qthis);
  // proto:  QWidget * QFormLayout::labelForField(QLayout * field);
extern void _ZNK11QFormLayout13labelForFieldEP7QLayout(void* qthis, void* arg0);
  // proto:  void QFormLayout::addRow(const QString & labelText, QLayout * field);
extern void _ZN11QFormLayout6addRowERK7QStringP7QLayout(void* qthis, void* arg0, void* arg1);
  // proto:  void QFormLayout::insertRow(int row, const QString & labelText, QLayout * field);
extern void _ZN11QFormLayout9insertRowEiRK7QStringP7QLayout(void* qthis, int arg0, void* arg1, void* arg2);
  // proto:  QWidget * QFormLayout::labelForField(QWidget * field);
extern void _ZNK11QFormLayout13labelForFieldEP7QWidget(void* qthis, void* arg0);
  // proto:  void QFormLayout::insertRow(int row, QWidget * label, QLayout * field);
extern void _ZN11QFormLayout9insertRowEiP7QWidgetP7QLayout(void* qthis, int arg0, void* arg1, void* arg2);
  // proto:  int QFormLayout::count();
extern void _ZNK11QFormLayout5countEv(void* qthis);
  // proto:  int QFormLayout::spacing();
extern void _ZNK11QFormLayout7spacingEv(void* qthis);
  // proto:  void QFormLayout::QFormLayout(QWidget * parent);
extern void* dector_ZN11QFormLayoutC1EP7QWidget(void* arg0);
extern void _ZN11QFormLayoutC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QFormLayout::insertRow(int row, QLayout * layout);
extern void _ZN11QFormLayout9insertRowEiP7QLayout(void* qthis, int arg0, void* arg1);
  // proto:  void QFormLayout::setGeometry(const QRect & rect);
extern void _ZN11QFormLayout11setGeometryERK5QRect(void* qthis, void* arg0);
  // proto:  void QFormLayout::setVerticalSpacing(int spacing);
extern void _ZN11QFormLayout18setVerticalSpacingEi(void* qthis, int arg0);
  // proto:  void QFormLayout::setHorizontalSpacing(int spacing);
extern void _ZN11QFormLayout20setHorizontalSpacingEi(void* qthis, int arg0);
  // proto:  void QFormLayout::insertRow(int row, const QString & labelText, QWidget * field);
extern void _ZN11QFormLayout9insertRowEiRK7QStringP7QWidget(void* qthis, int arg0, void* arg1, void* arg2);
  // proto:  const QMetaObject * QFormLayout::metaObject();
extern void _ZNK11QFormLayout10metaObjectEv(void* qthis);
  // proto:  void QFormLayout::insertRow(int row, QWidget * label, QWidget * field);
extern void _ZN11QFormLayout9insertRowEiP7QWidgetS1_(void* qthis, int arg0, void* arg1, void* arg2);
  // proto:  void QFormLayout::setSpacing(int );
extern void _ZN11QFormLayout10setSpacingEi(void* qthis, int arg0);
  // proto:  void QFormLayout::~QFormLayout();
extern void _ZN11QFormLayoutD0Ev(void* qthis);
  // proto:  void QFormLayout::addRow(QLayout * layout);
extern void _ZN11QFormLayout6addRowEP7QLayout(void* qthis, void* arg0);
  // proto:  QSize QFormLayout::sizeHint();
extern void _ZNK11QFormLayout8sizeHintEv(void* qthis);
  // proto:  void QFormLayout::invalidate();
extern void _ZN11QFormLayout10invalidateEv(void* qthis);
  // proto:  QLayoutItem * QFormLayout::itemAt(int index);
extern void _ZNK11QFormLayout6itemAtEi(void* qthis, int arg0);
  // proto:  QLayoutItem * QFormLayout::takeAt(int index);
extern void _ZN11QFormLayout6takeAtEi(void* qthis, int arg0);
  // proto:  void QFormLayout::addRow(const QString & labelText, QWidget * field);
extern void _ZN11QFormLayout6addRowERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  QSize QFormLayout::minimumSize();
extern void _ZNK11QFormLayout11minimumSizeEv(void* qthis);
  // proto:  void QFormLayout::addRow(QWidget * widget);
extern void _ZN11QFormLayout6addRowEP7QWidget(void* qthis, void* arg0);
  // proto:  void QFormLayout::addRow(QWidget * label, QLayout * field);
extern void _ZN11QFormLayout6addRowEP7QWidgetP7QLayout(void* qthis, void* arg0, void* arg1);
  // proto:  int QFormLayout::verticalSpacing();
extern void _ZNK11QFormLayout15verticalSpacingEv(void* qthis);
  // proto:  int QFormLayout::heightForWidth(int width);
extern void _ZNK11QFormLayout14heightForWidthEi(void* qthis, int arg0);
  // proto:  void QFormLayout::addItem(QLayoutItem * item);
extern void _ZN11QFormLayout7addItemEP11QLayoutItem(void* qthis, void* arg0);
  // proto:  bool QFormLayout::hasHeightForWidth();
extern void _ZNK11QFormLayout17hasHeightForWidthEv(void* qthis);
  // proto:  void QFormLayout::insertRow(int row, QWidget * widget);
extern void _ZN11QFormLayout9insertRowEiP7QWidget(void* qthis, int arg0, void* arg1);
  // proto:  void QFormLayout::addRow(QWidget * label, QWidget * field);
extern void _ZN11QFormLayout6addRowEP7QWidgetS1_(void* qthis, void* arg0, void* arg1);
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

// class sizeof(QFormLayout)=1
type QFormLayout struct {
  /*qbase*/ QLayout;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  int QFormLayout::horizontalSpacing();
func (this *QFormLayout) horizontalSpacing(args ...interface{}) () {
  // horizontalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout17horizontalSpacingEv
  default:
    qtrt.ErrorResolve("QFormLayout", "horizontalSpacing", args)
  }

}

  // proto:  int QFormLayout::rowCount();
func (this *QFormLayout) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout8rowCountEv
  default:
    qtrt.ErrorResolve("QFormLayout", "rowCount", args)
  }

}

  // proto:  QWidget * QFormLayout::labelForField(QLayout * field);
func (this *QFormLayout) labelForField(args ...interface{}) () {
  // labelForField(class QLayout *)
  // labelForField(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout13labelForFieldEP7QLayout
  case 1:
    // invoke: _ZNK11QFormLayout13labelForFieldEP7QWidget
  default:
    qtrt.ErrorResolve("QFormLayout", "labelForField", args)
  }

}

  // proto:  void QFormLayout::addRow(const QString & labelText, QLayout * field);
func (this *QFormLayout) addRow(args ...interface{}) () {
  // addRow(const class QString &, class QLayout *)
  // addRow(class QLayout *)
  // addRow(const class QString &, class QWidget *)
  // addRow(class QWidget *)
  // addRow(class QWidget *, class QLayout *)
  // addRow(class QWidget *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[4][1] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[5][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout6addRowERK7QStringP7QLayout
  case 1:
    // invoke: _ZN11QFormLayout6addRowEP7QLayout
  case 2:
    // invoke: _ZN11QFormLayout6addRowERK7QStringP7QWidget
  case 3:
    // invoke: _ZN11QFormLayout6addRowEP7QWidget
  case 4:
    // invoke: _ZN11QFormLayout6addRowEP7QWidgetP7QLayout
  case 5:
    // invoke: _ZN11QFormLayout6addRowEP7QWidgetS1_
  default:
    qtrt.ErrorResolve("QFormLayout", "addRow", args)
  }

}

  // proto:  void QFormLayout::insertRow(int row, const QString & labelText, QLayout * field);
func (this *QFormLayout) insertRow(args ...interface{}) () {
  // insertRow(int, const class QString &, class QLayout *)
  // insertRow(int, class QWidget *, class QLayout *)
  // insertRow(int, class QLayout *)
  // insertRow(int, const class QString &, class QWidget *)
  // insertRow(int, class QWidget *, class QWidget *)
  // insertRow(int, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][2] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][2] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[4][2] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout9insertRowEiRK7QStringP7QLayout
  case 1:
    // invoke: _ZN11QFormLayout9insertRowEiP7QWidgetP7QLayout
  case 2:
    // invoke: _ZN11QFormLayout9insertRowEiP7QLayout
  case 3:
    // invoke: _ZN11QFormLayout9insertRowEiRK7QStringP7QWidget
  case 4:
    // invoke: _ZN11QFormLayout9insertRowEiP7QWidgetS1_
  case 5:
    // invoke: _ZN11QFormLayout9insertRowEiP7QWidget
  default:
    qtrt.ErrorResolve("QFormLayout", "insertRow", args)
  }

}

  // proto:  int QFormLayout::count();
func (this *QFormLayout) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout5countEv
  default:
    qtrt.ErrorResolve("QFormLayout", "count", args)
  }

}

  // proto:  int QFormLayout::spacing();
func (this *QFormLayout) spacing(args ...interface{}) () {
  // spacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout7spacingEv
  default:
    qtrt.ErrorResolve("QFormLayout", "spacing", args)
  }

}

  // proto:  void QFormLayout::QFormLayout(QWidget * parent);
func NewQFormLayout(args ...interface{}) QFormLayout {
  return QFormLayout{}
}

  // proto:  void QFormLayout::setGeometry(const QRect & rect);
func (this *QFormLayout) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout11setGeometryERK5QRect
  default:
    qtrt.ErrorResolve("QFormLayout", "setGeometry", args)
  }

}

  // proto:  void QFormLayout::setVerticalSpacing(int spacing);
func (this *QFormLayout) setVerticalSpacing(args ...interface{}) () {
  // setVerticalSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout18setVerticalSpacingEi
  default:
    qtrt.ErrorResolve("QFormLayout", "setVerticalSpacing", args)
  }

}

  // proto:  void QFormLayout::setHorizontalSpacing(int spacing);
func (this *QFormLayout) setHorizontalSpacing(args ...interface{}) () {
  // setHorizontalSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout20setHorizontalSpacingEi
  default:
    qtrt.ErrorResolve("QFormLayout", "setHorizontalSpacing", args)
  }

}

  // proto:  const QMetaObject * QFormLayout::metaObject();
func (this *QFormLayout) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout10metaObjectEv
  default:
    qtrt.ErrorResolve("QFormLayout", "metaObject", args)
  }

}

  // proto:  void QFormLayout::setSpacing(int );
func (this *QFormLayout) setSpacing(args ...interface{}) () {
  // setSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout10setSpacingEi
  default:
    qtrt.ErrorResolve("QFormLayout", "setSpacing", args)
  }

}

  // proto:  void QFormLayout::~QFormLayout();
func (this *QFormLayout) FreeQFormLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFormLayout", "~QFormLayout", args)
  }

}

  // proto:  QSize QFormLayout::sizeHint();
func (this *QFormLayout) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout8sizeHintEv
  default:
    qtrt.ErrorResolve("QFormLayout", "sizeHint", args)
  }

}

  // proto:  void QFormLayout::invalidate();
func (this *QFormLayout) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout10invalidateEv
  default:
    qtrt.ErrorResolve("QFormLayout", "invalidate", args)
  }

}

  // proto:  QLayoutItem * QFormLayout::itemAt(int index);
func (this *QFormLayout) itemAt(args ...interface{}) () {
  // itemAt(int, enum QFormLayout::ItemRole)
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "QFormLayout::ItemRole"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout6itemAtEiNS_8ItemRoleE
  case 1:
    // invoke: _ZNK11QFormLayout6itemAtEi
  default:
    qtrt.ErrorResolve("QFormLayout", "itemAt", args)
  }

}

  // proto:  QLayoutItem * QFormLayout::takeAt(int index);
func (this *QFormLayout) takeAt(args ...interface{}) () {
  // takeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout6takeAtEi
  default:
    qtrt.ErrorResolve("QFormLayout", "takeAt", args)
  }

}

  // proto:  QSize QFormLayout::minimumSize();
func (this *QFormLayout) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout11minimumSizeEv
  default:
    qtrt.ErrorResolve("QFormLayout", "minimumSize", args)
  }

}

  // proto:  int QFormLayout::verticalSpacing();
func (this *QFormLayout) verticalSpacing(args ...interface{}) () {
  // verticalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout15verticalSpacingEv
  default:
    qtrt.ErrorResolve("QFormLayout", "verticalSpacing", args)
  }

}

  // proto:  int QFormLayout::heightForWidth(int width);
func (this *QFormLayout) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout14heightForWidthEi
  default:
    qtrt.ErrorResolve("QFormLayout", "heightForWidth", args)
  }

}

  // proto:  void QFormLayout::addItem(QLayoutItem * item);
func (this *QFormLayout) addItem(args ...interface{}) () {
  // addItem(class QLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout7addItemEP11QLayoutItem
  default:
    qtrt.ErrorResolve("QFormLayout", "addItem", args)
  }

}

  // proto:  bool QFormLayout::hasHeightForWidth();
func (this *QFormLayout) hasHeightForWidth(args ...interface{}) () {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout17hasHeightForWidthEv
  default:
    qtrt.ErrorResolve("QFormLayout", "hasHeightForWidth", args)
  }

}

// <= body block end
