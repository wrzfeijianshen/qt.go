package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qtreewidget.h
// dst-file: /src/widgets/qtreewidget.go
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
  // proto:  void QTreeWidget::setColumnCount(int columns);
extern void _ZN11QTreeWidget14setColumnCountEi(void* qthis, int arg0);
  // proto:  void QTreeWidget::~QTreeWidget();
extern void _ZN11QTreeWidgetD0Ev(void* qthis);
  // proto:  QList<QTreeWidgetItem *> QTreeWidget::selectedItems();
extern void _ZNK11QTreeWidget13selectedItemsEv(void* qthis);
  // proto:  bool QTreeWidget::isItemExpanded(const QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget14isItemExpandedEPK15QTreeWidgetItem(void* qthis, void* arg0);
  // proto:  void QTreeWidget::QTreeWidget(const QTreeWidget & );
extern void* dector_ZN11QTreeWidgetC1ERKS_(void* arg0);
extern void _ZN11QTreeWidgetC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTreeWidget::setItemHidden(const QTreeWidgetItem * item, bool hide);
extern void _ZN11QTreeWidget13setItemHiddenEPK15QTreeWidgetItemb(void* qthis, void* arg0, bool arg1);
  // proto:  int QTreeWidget::indexOfTopLevelItem(QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget19indexOfTopLevelItemEP15QTreeWidgetItem(void* qthis, void* arg0);
  // proto:  void QTreeWidget::insertTopLevelItem(int index, QTreeWidgetItem * item);
extern void _ZN11QTreeWidget18insertTopLevelItemEiP15QTreeWidgetItem(void* qthis, int arg0, void* arg1);
  // proto:  void QTreeWidget::setItemWidget(QTreeWidgetItem * item, int column, QWidget * widget);
extern void _ZN11QTreeWidget13setItemWidgetEP15QTreeWidgetItemiP7QWidget(void* qthis, void* arg0, int arg1, void* arg2);
  // proto:  bool QTreeWidget::isItemSelected(const QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget14isItemSelectedEPK15QTreeWidgetItem(void* qthis, void* arg0);
  // proto:  int QTreeWidget::currentColumn();
extern void _ZNK11QTreeWidget13currentColumnEv(void* qthis);
  // proto:  bool QTreeWidget::isFirstItemColumnSpanned(const QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget24isFirstItemColumnSpannedEPK15QTreeWidgetItem(void* qthis, void* arg0);
  // proto:  void QTreeWidget::clear();
extern void _ZN11QTreeWidget5clearEv(void* qthis);
  // proto:  void QTreeWidget::setHeaderLabels(const QStringList & labels);
extern void _ZN11QTreeWidget15setHeaderLabelsERK11QStringList(void* qthis, void* arg0);
  // proto:  QTreeWidgetItem * QTreeWidget::invisibleRootItem();
extern void _ZNK11QTreeWidget17invisibleRootItemEv(void* qthis);
  // proto:  const QMetaObject * QTreeWidget::metaObject();
extern void _ZNK11QTreeWidget10metaObjectEv(void* qthis);
  // proto:  QTreeWidgetItem * QTreeWidget::itemBelow(const QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget9itemBelowEPK15QTreeWidgetItem(void* qthis, void* arg0);
  // proto:  int QTreeWidget::sortColumn();
extern void _ZNK11QTreeWidget10sortColumnEv(void* qthis);
  // proto:  QTreeWidgetItem * QTreeWidget::itemAt(int x, int y);
extern void demth_ZNK11QTreeWidget6itemAtEii(void* qthis, int arg0, int arg1);
  // proto:  QTreeWidgetItem * QTreeWidget::currentItem();
extern void _ZNK11QTreeWidget11currentItemEv(void* qthis);
  // proto:  QTreeWidgetItem * QTreeWidget::itemAt(const QPoint & p);
extern void _ZNK11QTreeWidget6itemAtERK6QPoint(void* qthis, void* arg0);
  // proto:  void QTreeWidget::setCurrentItem(QTreeWidgetItem * item, int column);
extern void _ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi(void* qthis, void* arg0, int arg1);
  // proto:  QTreeWidgetItem * QTreeWidget::topLevelItem(int index);
extern void _ZNK11QTreeWidget12topLevelItemEi(void* qthis, int arg0);
  // proto:  int QTreeWidget::topLevelItemCount();
extern void _ZNK11QTreeWidget17topLevelItemCountEv(void* qthis);
  // proto:  QTreeWidgetItem * QTreeWidget::headerItem();
extern void _ZNK11QTreeWidget10headerItemEv(void* qthis);
  // proto:  void QTreeWidget::setFirstItemColumnSpanned(const QTreeWidgetItem * item, bool span);
extern void _ZN11QTreeWidget25setFirstItemColumnSpannedEPK15QTreeWidgetItemb(void* qthis, void* arg0, bool arg1);
  // proto:  void QTreeWidget::removeItemWidget(QTreeWidgetItem * item, int column);
extern void demth_ZN11QTreeWidget16removeItemWidgetEP15QTreeWidgetItemi(void* qthis, void* arg0, int arg1);
  // proto:  QTreeWidgetItem * QTreeWidget::itemAbove(const QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget9itemAboveEPK15QTreeWidgetItem(void* qthis, void* arg0);
  // proto:  void QTreeWidget::expandItem(const QTreeWidgetItem * item);
extern void _ZN11QTreeWidget10expandItemEPK15QTreeWidgetItem(void* qthis, void* arg0);
  // proto:  void QTreeWidget::setHeaderItem(QTreeWidgetItem * item);
extern void _ZN11QTreeWidget13setHeaderItemEP15QTreeWidgetItem(void* qthis, void* arg0);
  // proto:  void QTreeWidget::collapseItem(const QTreeWidgetItem * item);
extern void _ZN11QTreeWidget12collapseItemEPK15QTreeWidgetItem(void* qthis, void* arg0);
  // proto:  QTreeWidgetItem * QTreeWidget::takeTopLevelItem(int index);
extern void _ZN11QTreeWidget16takeTopLevelItemEi(void* qthis, int arg0);
  // proto:  QWidget * QTreeWidget::itemWidget(QTreeWidgetItem * item, int column);
extern void _ZNK11QTreeWidget10itemWidgetEP15QTreeWidgetItemi(void* qthis, void* arg0, int arg1);
  // proto:  void QTreeWidget::editItem(QTreeWidgetItem * item, int column);
extern void _ZN11QTreeWidget8editItemEP15QTreeWidgetItemi(void* qthis, void* arg0, int arg1);
  // proto:  void QTreeWidget::setItemExpanded(const QTreeWidgetItem * item, bool expand);
extern void _ZN11QTreeWidget15setItemExpandedEPK15QTreeWidgetItemb(void* qthis, void* arg0, bool arg1);
  // proto:  void QTreeWidget::addTopLevelItem(QTreeWidgetItem * item);
extern void _ZN11QTreeWidget15addTopLevelItemEP15QTreeWidgetItem(void* qthis, void* arg0);
  // proto:  void QTreeWidget::closePersistentEditor(QTreeWidgetItem * item, int column);
extern void _ZN11QTreeWidget21closePersistentEditorEP15QTreeWidgetItemi(void* qthis, void* arg0, int arg1);
  // proto:  void QTreeWidget::QTreeWidget(QWidget * parent);
extern void* dector_ZN11QTreeWidgetC1EP7QWidget(void* arg0);
extern void _ZN11QTreeWidgetC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QTreeWidget::setSelectionModel(QItemSelectionModel * selectionModel);
extern void _ZN11QTreeWidget17setSelectionModelEP19QItemSelectionModel(void* qthis, void* arg0);
  // proto:  QRect QTreeWidget::visualItemRect(const QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget14visualItemRectEPK15QTreeWidgetItem(void* qthis, void* arg0);
  // proto:  void QTreeWidget::setHeaderLabel(const QString & label);
extern void demth_ZN11QTreeWidget14setHeaderLabelERK7QString(void* qthis, void* arg0);
  // proto:  bool QTreeWidget::isItemHidden(const QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget12isItemHiddenEPK15QTreeWidgetItem(void* qthis, void* arg0);
  // proto:  void QTreeWidget::openPersistentEditor(QTreeWidgetItem * item, int column);
extern void _ZN11QTreeWidget20openPersistentEditorEP15QTreeWidgetItemi(void* qthis, void* arg0, int arg1);
  // proto:  int QTreeWidget::columnCount();
extern void _ZNK11QTreeWidget11columnCountEv(void* qthis);
  // proto:  void QTreeWidget::setCurrentItem(QTreeWidgetItem * item);
extern void _ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItem(void* qthis, void* arg0);
  // proto:  void QTreeWidget::setItemSelected(const QTreeWidgetItem * item, bool select);
extern void _ZN11QTreeWidget15setItemSelectedEPK15QTreeWidgetItemb(void* qthis, void* arg0, bool arg1);
  // proto:  void QTreeWidgetItem::setFirstColumnSpanned(bool span);
extern void demth_ZN15QTreeWidgetItem21setFirstColumnSpannedEb(void* qthis, bool arg0);
  // proto:  int QTreeWidgetItem::indexOfChild(QTreeWidgetItem * child);
extern void demth_ZNK15QTreeWidgetItem12indexOfChildEPS_(void* qthis, void* arg0);
  // proto:  QVariant QTreeWidgetItem::data(int column, int role);
extern void _ZNK15QTreeWidgetItem4dataEii(void* qthis, int arg0, int arg1);
  // proto:  QTreeWidgetItem * QTreeWidgetItem::parent();
extern void demth_ZNK15QTreeWidgetItem6parentEv(void* qthis);
  // proto:  void QTreeWidgetItem::setFont(int column, const QFont & font);
extern void demth_ZN15QTreeWidgetItem7setFontEiRK5QFont(void* qthis, int arg0, void* arg1);
  // proto:  void QTreeWidgetItem::setData(int column, int role, const QVariant & value);
extern void _ZN15QTreeWidgetItem7setDataEiiRK8QVariant(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  QFont QTreeWidgetItem::font(int column);
extern void demth_ZNK15QTreeWidgetItem4fontEi(void* qthis, int arg0);
  // proto:  void QTreeWidgetItem::setStatusTip(int column, const QString & statusTip);
extern void demth_ZN15QTreeWidgetItem12setStatusTipEiRK7QString(void* qthis, int arg0, void* arg1);
  // proto:  void QTreeWidgetItem::setExpanded(bool expand);
extern void demth_ZN15QTreeWidgetItem11setExpandedEb(void* qthis, bool arg0);
  // proto:  void QTreeWidgetItem::write(QDataStream & out);
extern void _ZNK15QTreeWidgetItem5writeER11QDataStream(void* qthis, void* arg0);
  // proto:  bool QTreeWidgetItem::isExpanded();
extern void demth_ZNK15QTreeWidgetItem10isExpandedEv(void* qthis);
  // proto:  QList<QTreeWidgetItem *> QTreeWidgetItem::takeChildren();
extern void _ZN15QTreeWidgetItem12takeChildrenEv(void* qthis);
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidgetItem * parent, int type);
extern void* dector_ZN15QTreeWidgetItemC1EPS_i(void* arg0, int arg1);
extern void _ZN15QTreeWidgetItemC1EPS_i(void* qthis, void* arg0, int arg1);
  // proto:  void QTreeWidgetItem::setIcon(int column, const QIcon & icon);
extern void demth_ZN15QTreeWidgetItem7setIconEiRK5QIcon(void* qthis, int arg0, void* arg1);
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidgetItem * parent, QTreeWidgetItem * after, int type);
extern void* dector_ZN15QTreeWidgetItemC1EPS_S0_i(void* arg0, void* arg1, int arg2);
extern void _ZN15QTreeWidgetItemC1EPS_S0_i(void* qthis, void* arg0, void* arg1, int arg2);
  // proto:  QString QTreeWidgetItem::toolTip(int column);
extern void demth_ZNK15QTreeWidgetItem7toolTipEi(void* qthis, int arg0);
  // proto:  QColor QTreeWidgetItem::backgroundColor(int column);
extern void demth_ZNK15QTreeWidgetItem15backgroundColorEi(void* qthis, int arg0);
  // proto:  QString QTreeWidgetItem::text(int column);
extern void demth_ZNK15QTreeWidgetItem4textEi(void* qthis, int arg0);
  // proto:  bool QTreeWidgetItem::isHidden();
extern void demth_ZNK15QTreeWidgetItem8isHiddenEv(void* qthis);
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidget * view, QTreeWidgetItem * after, int type);
extern void* dector_ZN15QTreeWidgetItemC1EP11QTreeWidgetPS_i(void* arg0, void* arg1, int arg2);
extern void _ZN15QTreeWidgetItemC1EP11QTreeWidgetPS_i(void* qthis, void* arg0, void* arg1, int arg2);
  // proto:  void QTreeWidgetItem::setTextAlignment(int column, int alignment);
extern void demth_ZN15QTreeWidgetItem16setTextAlignmentEii(void* qthis, int arg0, int arg1);
  // proto:  void QTreeWidgetItem::insertChild(int index, QTreeWidgetItem * child);
extern void _ZN15QTreeWidgetItem11insertChildEiPS_(void* qthis, int arg0, void* arg1);
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(const QTreeWidgetItem & other);
extern void* dector_ZN15QTreeWidgetItemC1ERKS_(void* arg0);
extern void _ZN15QTreeWidgetItemC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QTreeWidgetItem::isDisabled();
extern void demth_ZNK15QTreeWidgetItem10isDisabledEv(void* qthis);
  // proto:  void QTreeWidgetItem::setText(int column, const QString & text);
extern void demth_ZN15QTreeWidgetItem7setTextEiRK7QString(void* qthis, int arg0, void* arg1);
  // proto:  void QTreeWidgetItem::setTextColor(int column, const QColor & color);
extern void demth_ZN15QTreeWidgetItem12setTextColorEiRK6QColor(void* qthis, int arg0, void* arg1);
  // proto:  QSize QTreeWidgetItem::sizeHint(int column);
extern void demth_ZNK15QTreeWidgetItem8sizeHintEi(void* qthis, int arg0);
  // proto:  QString QTreeWidgetItem::whatsThis(int column);
extern void demth_ZNK15QTreeWidgetItem9whatsThisEi(void* qthis, int arg0);
  // proto:  void QTreeWidgetItem::setWhatsThis(int column, const QString & whatsThis);
extern void demth_ZN15QTreeWidgetItem12setWhatsThisEiRK7QString(void* qthis, int arg0, void* arg1);
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(int type);
extern void* dector_ZN15QTreeWidgetItemC1Ei(int arg0);
extern void _ZN15QTreeWidgetItemC1Ei(void* qthis, int arg0);
  // proto:  QColor QTreeWidgetItem::textColor(int column);
extern void demth_ZNK15QTreeWidgetItem9textColorEi(void* qthis, int arg0);
  // proto:  QIcon QTreeWidgetItem::icon(int column);
extern void demth_ZNK15QTreeWidgetItem4iconEi(void* qthis, int arg0);
  // proto:  void QTreeWidgetItem::setToolTip(int column, const QString & toolTip);
extern void demth_ZN15QTreeWidgetItem10setToolTipEiRK7QString(void* qthis, int arg0, void* arg1);
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidget * view, const QStringList & strings, int type);
extern void* dector_ZN15QTreeWidgetItemC1EP11QTreeWidgetRK11QStringListi(void* arg0, void* arg1, int arg2);
extern void _ZN15QTreeWidgetItemC1EP11QTreeWidgetRK11QStringListi(void* qthis, void* arg0, void* arg1, int arg2);
  // proto:  bool QTreeWidgetItem::isFirstColumnSpanned();
extern void demth_ZNK15QTreeWidgetItem20isFirstColumnSpannedEv(void* qthis);
  // proto:  int QTreeWidgetItem::textAlignment(int column);
extern void demth_ZNK15QTreeWidgetItem13textAlignmentEi(void* qthis, int arg0);
  // proto:  QTreeWidgetItem * QTreeWidgetItem::child(int index);
extern void demth_ZNK15QTreeWidgetItem5childEi(void* qthis, int arg0);
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(const QStringList & strings, int type);
extern void* dector_ZN15QTreeWidgetItemC1ERK11QStringListi(void* arg0, int arg1);
extern void _ZN15QTreeWidgetItemC1ERK11QStringListi(void* qthis, void* arg0, int arg1);
  // proto:  void QTreeWidgetItem::setSelected(bool select);
extern void demth_ZN15QTreeWidgetItem11setSelectedEb(void* qthis, bool arg0);
  // proto:  void QTreeWidgetItem::~QTreeWidgetItem();
extern void _ZN15QTreeWidgetItemD0Ev(void* qthis);
  // proto:  void QTreeWidgetItem::setHidden(bool hide);
extern void demth_ZN15QTreeWidgetItem9setHiddenEb(void* qthis, bool arg0);
  // proto:  int QTreeWidgetItem::columnCount();
extern void demth_ZNK15QTreeWidgetItem11columnCountEv(void* qthis);
  // proto:  QTreeWidgetItem * QTreeWidgetItem::takeChild(int index);
extern void _ZN15QTreeWidgetItem9takeChildEi(void* qthis, int arg0);
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidgetItem * parent, const QStringList & strings, int type);
extern void* dector_ZN15QTreeWidgetItemC1EPS_RK11QStringListi(void* arg0, void* arg1, int arg2);
extern void _ZN15QTreeWidgetItemC1EPS_RK11QStringListi(void* qthis, void* arg0, void* arg1, int arg2);
  // proto:  void QTreeWidgetItem::setDisabled(bool disabled);
extern void demth_ZN15QTreeWidgetItem11setDisabledEb(void* qthis, bool arg0);
  // proto:  void QTreeWidgetItem::setBackground(int column, const QBrush & brush);
extern void demth_ZN15QTreeWidgetItem13setBackgroundEiRK6QBrush(void* qthis, int arg0, void* arg1);
  // proto:  void QTreeWidgetItem::addChild(QTreeWidgetItem * child);
extern void _ZN15QTreeWidgetItem8addChildEPS_(void* qthis, void* arg0);
  // proto:  void QTreeWidgetItem::removeChild(QTreeWidgetItem * child);
extern void _ZN15QTreeWidgetItem11removeChildEPS_(void* qthis, void* arg0);
  // proto:  QTreeWidgetItem * QTreeWidgetItem::clone();
extern void _ZNK15QTreeWidgetItem5cloneEv(void* qthis);
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidget * view, int type);
extern void* dector_ZN15QTreeWidgetItemC1EP11QTreeWidgeti(void* arg0, int arg1);
extern void _ZN15QTreeWidgetItemC1EP11QTreeWidgeti(void* qthis, void* arg0, int arg1);
  // proto:  void QTreeWidgetItem::setSizeHint(int column, const QSize & size);
extern void demth_ZN15QTreeWidgetItem11setSizeHintEiRK5QSize(void* qthis, int arg0, void* arg1);
  // proto:  QBrush QTreeWidgetItem::foreground(int column);
extern void demth_ZNK15QTreeWidgetItem10foregroundEi(void* qthis, int arg0);
  // proto:  int QTreeWidgetItem::childCount();
extern void demth_ZNK15QTreeWidgetItem10childCountEv(void* qthis);
  // proto:  void QTreeWidgetItem::setBackgroundColor(int column, const QColor & color);
extern void demth_ZN15QTreeWidgetItem18setBackgroundColorEiRK6QColor(void* qthis, int arg0, void* arg1);
  // proto:  QString QTreeWidgetItem::statusTip(int column);
extern void demth_ZNK15QTreeWidgetItem9statusTipEi(void* qthis, int arg0);
  // proto:  QBrush QTreeWidgetItem::background(int column);
extern void demth_ZNK15QTreeWidgetItem10backgroundEi(void* qthis, int arg0);
  // proto:  int QTreeWidgetItem::type();
extern void demth_ZNK15QTreeWidgetItem4typeEv(void* qthis);
  // proto:  QTreeWidget * QTreeWidgetItem::treeWidget();
extern void demth_ZNK15QTreeWidgetItem10treeWidgetEv(void* qthis);
  // proto:  void QTreeWidgetItem::read(QDataStream & in);
extern void _ZN15QTreeWidgetItem4readER11QDataStream(void* qthis, void* arg0);
  // proto:  void QTreeWidgetItem::setForeground(int column, const QBrush & brush);
extern void demth_ZN15QTreeWidgetItem13setForegroundEiRK6QBrush(void* qthis, int arg0, void* arg1);
  // proto:  bool QTreeWidgetItem::isSelected();
extern void demth_ZNK15QTreeWidgetItem10isSelectedEv(void* qthis);
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

// class sizeof(QTreeWidget)=1
type QTreeWidget struct {
  /*qbase*/ QTreeView;
  qclsinst uint64 /* *mut c_void*/;
//  _itemDoubleClicked QTreeWidget_itemDoubleClicked_signal;
//  _itemClicked QTreeWidget_itemClicked_signal;
//  _currentItemChanged QTreeWidget_currentItemChanged_signal;
//  _itemEntered QTreeWidget_itemEntered_signal;
//  _itemPressed QTreeWidget_itemPressed_signal;
//  _itemSelectionChanged QTreeWidget_itemSelectionChanged_signal;
//  _itemActivated QTreeWidget_itemActivated_signal;
//  _itemExpanded QTreeWidget_itemExpanded_signal;
//  _itemChanged QTreeWidget_itemChanged_signal;
//  _itemCollapsed QTreeWidget_itemCollapsed_signal;
}

// class sizeof(QTreeWidgetItem)=1
type QTreeWidgetItem struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QTreeWidget::setColumnCount(int columns);
func (this *QTreeWidget) setColumnCount(args ...interface{}) () {
  // setColumnCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget14setColumnCountEi
  default:
    qtrt.ErrorResolve("QTreeWidget", "setColumnCount", args)
  }

}

  // proto:  void QTreeWidget::~QTreeWidget();
func (this *QTreeWidget) FreeQTreeWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTreeWidget", "~QTreeWidget", args)
  }

}

  // proto:  QList<QTreeWidgetItem *> QTreeWidget::selectedItems();
func (this *QTreeWidget) selectedItems(args ...interface{}) () {
  // selectedItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget13selectedItemsEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "selectedItems", args)
  }

}

  // proto:  bool QTreeWidget::isItemExpanded(const QTreeWidgetItem * item);
func (this *QTreeWidget) isItemExpanded(args ...interface{}) () {
  // isItemExpanded(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget14isItemExpandedEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "isItemExpanded", args)
  }

}

  // proto:  void QTreeWidget::QTreeWidget(const QTreeWidget & );
func NewQTreeWidget(args ...interface{}) QTreeWidget {
  return QTreeWidget{}
}

  // proto:  void QTreeWidget::setItemHidden(const QTreeWidgetItem * item, bool hide);
func (this *QTreeWidget) setItemHidden(args ...interface{}) () {
  // setItemHidden(const class QTreeWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget13setItemHiddenEPK15QTreeWidgetItemb
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemHidden", args)
  }

}

  // proto:  int QTreeWidget::indexOfTopLevelItem(QTreeWidgetItem * item);
func (this *QTreeWidget) indexOfTopLevelItem(args ...interface{}) () {
  // indexOfTopLevelItem(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget19indexOfTopLevelItemEP15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "indexOfTopLevelItem", args)
  }

}

  // proto:  void QTreeWidget::insertTopLevelItem(int index, QTreeWidgetItem * item);
func (this *QTreeWidget) insertTopLevelItem(args ...interface{}) () {
  // insertTopLevelItem(int, class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget18insertTopLevelItemEiP15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "insertTopLevelItem", args)
  }

}

  // proto:  void QTreeWidget::setItemWidget(QTreeWidgetItem * item, int column, QWidget * widget);
func (this *QTreeWidget) setItemWidget(args ...interface{}) () {
  // setItemWidget(class QTreeWidgetItem *, int, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget13setItemWidgetEP15QTreeWidgetItemiP7QWidget
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemWidget", args)
  }

}

  // proto:  bool QTreeWidget::isItemSelected(const QTreeWidgetItem * item);
func (this *QTreeWidget) isItemSelected(args ...interface{}) () {
  // isItemSelected(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget14isItemSelectedEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "isItemSelected", args)
  }

}

  // proto:  int QTreeWidget::currentColumn();
func (this *QTreeWidget) currentColumn(args ...interface{}) () {
  // currentColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget13currentColumnEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "currentColumn", args)
  }

}

  // proto:  bool QTreeWidget::isFirstItemColumnSpanned(const QTreeWidgetItem * item);
func (this *QTreeWidget) isFirstItemColumnSpanned(args ...interface{}) () {
  // isFirstItemColumnSpanned(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget24isFirstItemColumnSpannedEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "isFirstItemColumnSpanned", args)
  }

}

  // proto:  void QTreeWidget::clear();
func (this *QTreeWidget) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget5clearEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "clear", args)
  }

}

  // proto:  void QTreeWidget::setHeaderLabels(const QStringList & labels);
func (this *QTreeWidget) setHeaderLabels(args ...interface{}) () {
  // setHeaderLabels(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget15setHeaderLabelsERK11QStringList
  default:
    qtrt.ErrorResolve("QTreeWidget", "setHeaderLabels", args)
  }

}

  // proto:  QTreeWidgetItem * QTreeWidget::invisibleRootItem();
func (this *QTreeWidget) invisibleRootItem(args ...interface{}) () {
  // invisibleRootItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget17invisibleRootItemEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "invisibleRootItem", args)
  }

}

  // proto:  const QMetaObject * QTreeWidget::metaObject();
func (this *QTreeWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget10metaObjectEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "metaObject", args)
  }

}

  // proto:  QTreeWidgetItem * QTreeWidget::itemBelow(const QTreeWidgetItem * item);
func (this *QTreeWidget) itemBelow(args ...interface{}) () {
  // itemBelow(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget9itemBelowEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemBelow", args)
  }

}

  // proto:  int QTreeWidget::sortColumn();
func (this *QTreeWidget) sortColumn(args ...interface{}) () {
  // sortColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget10sortColumnEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "sortColumn", args)
  }

}

  // proto:  QTreeWidgetItem * QTreeWidget::itemAt(int x, int y);
func (this *QTreeWidget) itemAt(args ...interface{}) () {
  // itemAt(int, int)
  // itemAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget6itemAtEii
  case 1:
    // invoke: _ZNK11QTreeWidget6itemAtERK6QPoint
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemAt", args)
  }

}

  // proto:  QTreeWidgetItem * QTreeWidget::currentItem();
func (this *QTreeWidget) currentItem(args ...interface{}) () {
  // currentItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget11currentItemEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "currentItem", args)
  }

}

  // proto:  void QTreeWidget::setCurrentItem(QTreeWidgetItem * item, int column);
func (this *QTreeWidget) setCurrentItem(args ...interface{}) () {
  // setCurrentItem(class QTreeWidgetItem *, int, class QItemSelectionModel::SelectionFlags)
  // setCurrentItem(class QTreeWidgetItem *, int)
  // setCurrentItem(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int64Ty(false) // "QItemSelectionModel::SelectionFlags"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi6QFlagsIN19QItemSelectionModel13SelectionFlagEE
  case 1:
    // invoke: _ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi
  case 2:
    // invoke: _ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "setCurrentItem", args)
  }

}

  // proto:  QTreeWidgetItem * QTreeWidget::topLevelItem(int index);
func (this *QTreeWidget) topLevelItem(args ...interface{}) () {
  // topLevelItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget12topLevelItemEi
  default:
    qtrt.ErrorResolve("QTreeWidget", "topLevelItem", args)
  }

}

  // proto:  int QTreeWidget::topLevelItemCount();
func (this *QTreeWidget) topLevelItemCount(args ...interface{}) () {
  // topLevelItemCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget17topLevelItemCountEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "topLevelItemCount", args)
  }

}

  // proto:  QTreeWidgetItem * QTreeWidget::headerItem();
func (this *QTreeWidget) headerItem(args ...interface{}) () {
  // headerItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget10headerItemEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "headerItem", args)
  }

}

  // proto:  void QTreeWidget::setFirstItemColumnSpanned(const QTreeWidgetItem * item, bool span);
func (this *QTreeWidget) setFirstItemColumnSpanned(args ...interface{}) () {
  // setFirstItemColumnSpanned(const class QTreeWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget25setFirstItemColumnSpannedEPK15QTreeWidgetItemb
  default:
    qtrt.ErrorResolve("QTreeWidget", "setFirstItemColumnSpanned", args)
  }

}

  // proto:  void QTreeWidget::removeItemWidget(QTreeWidgetItem * item, int column);
func (this *QTreeWidget) removeItemWidget(args ...interface{}) () {
  // removeItemWidget(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget16removeItemWidgetEP15QTreeWidgetItemi
  default:
    qtrt.ErrorResolve("QTreeWidget", "removeItemWidget", args)
  }

}

  // proto:  QTreeWidgetItem * QTreeWidget::itemAbove(const QTreeWidgetItem * item);
func (this *QTreeWidget) itemAbove(args ...interface{}) () {
  // itemAbove(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget9itemAboveEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemAbove", args)
  }

}

  // proto:  void QTreeWidget::expandItem(const QTreeWidgetItem * item);
func (this *QTreeWidget) expandItem(args ...interface{}) () {
  // expandItem(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget10expandItemEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "expandItem", args)
  }

}

  // proto:  void QTreeWidget::setHeaderItem(QTreeWidgetItem * item);
func (this *QTreeWidget) setHeaderItem(args ...interface{}) () {
  // setHeaderItem(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget13setHeaderItemEP15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "setHeaderItem", args)
  }

}

  // proto:  void QTreeWidget::collapseItem(const QTreeWidgetItem * item);
func (this *QTreeWidget) collapseItem(args ...interface{}) () {
  // collapseItem(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget12collapseItemEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "collapseItem", args)
  }

}

  // proto:  QTreeWidgetItem * QTreeWidget::takeTopLevelItem(int index);
func (this *QTreeWidget) takeTopLevelItem(args ...interface{}) () {
  // takeTopLevelItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget16takeTopLevelItemEi
  default:
    qtrt.ErrorResolve("QTreeWidget", "takeTopLevelItem", args)
  }

}

  // proto:  QWidget * QTreeWidget::itemWidget(QTreeWidgetItem * item, int column);
func (this *QTreeWidget) itemWidget(args ...interface{}) () {
  // itemWidget(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget10itemWidgetEP15QTreeWidgetItemi
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemWidget", args)
  }

}

  // proto:  void QTreeWidget::editItem(QTreeWidgetItem * item, int column);
func (this *QTreeWidget) editItem(args ...interface{}) () {
  // editItem(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget8editItemEP15QTreeWidgetItemi
  default:
    qtrt.ErrorResolve("QTreeWidget", "editItem", args)
  }

}

  // proto:  void QTreeWidget::setItemExpanded(const QTreeWidgetItem * item, bool expand);
func (this *QTreeWidget) setItemExpanded(args ...interface{}) () {
  // setItemExpanded(const class QTreeWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget15setItemExpandedEPK15QTreeWidgetItemb
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemExpanded", args)
  }

}

  // proto:  void QTreeWidget::addTopLevelItem(QTreeWidgetItem * item);
func (this *QTreeWidget) addTopLevelItem(args ...interface{}) () {
  // addTopLevelItem(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget15addTopLevelItemEP15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "addTopLevelItem", args)
  }

}

  // proto:  void QTreeWidget::closePersistentEditor(QTreeWidgetItem * item, int column);
func (this *QTreeWidget) closePersistentEditor(args ...interface{}) () {
  // closePersistentEditor(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget21closePersistentEditorEP15QTreeWidgetItemi
  default:
    qtrt.ErrorResolve("QTreeWidget", "closePersistentEditor", args)
  }

}

  // proto:  void QTreeWidget::setSelectionModel(QItemSelectionModel * selectionModel);
func (this *QTreeWidget) setSelectionModel(args ...interface{}) () {
  // setSelectionModel(class QItemSelectionModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelectionModel{}) // "QItemSelectionModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget17setSelectionModelEP19QItemSelectionModel
  default:
    qtrt.ErrorResolve("QTreeWidget", "setSelectionModel", args)
  }

}

  // proto:  QRect QTreeWidget::visualItemRect(const QTreeWidgetItem * item);
func (this *QTreeWidget) visualItemRect(args ...interface{}) () {
  // visualItemRect(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget14visualItemRectEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "visualItemRect", args)
  }

}

  // proto:  void QTreeWidget::setHeaderLabel(const QString & label);
func (this *QTreeWidget) setHeaderLabel(args ...interface{}) () {
  // setHeaderLabel(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget14setHeaderLabelERK7QString
  default:
    qtrt.ErrorResolve("QTreeWidget", "setHeaderLabel", args)
  }

}

  // proto:  bool QTreeWidget::isItemHidden(const QTreeWidgetItem * item);
func (this *QTreeWidget) isItemHidden(args ...interface{}) () {
  // isItemHidden(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget12isItemHiddenEPK15QTreeWidgetItem
  default:
    qtrt.ErrorResolve("QTreeWidget", "isItemHidden", args)
  }

}

  // proto:  void QTreeWidget::openPersistentEditor(QTreeWidgetItem * item, int column);
func (this *QTreeWidget) openPersistentEditor(args ...interface{}) () {
  // openPersistentEditor(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget20openPersistentEditorEP15QTreeWidgetItemi
  default:
    qtrt.ErrorResolve("QTreeWidget", "openPersistentEditor", args)
  }

}

  // proto:  int QTreeWidget::columnCount();
func (this *QTreeWidget) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget11columnCountEv
  default:
    qtrt.ErrorResolve("QTreeWidget", "columnCount", args)
  }

}

  // proto:  void QTreeWidget::setItemSelected(const QTreeWidgetItem * item, bool select);
func (this *QTreeWidget) setItemSelected(args ...interface{}) () {
  // setItemSelected(const class QTreeWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget15setItemSelectedEPK15QTreeWidgetItemb
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemSelected", args)
  }

}

  // proto:  void QTreeWidgetItem::setFirstColumnSpanned(bool span);
func (this *QTreeWidgetItem) setFirstColumnSpanned(args ...interface{}) () {
  // setFirstColumnSpanned(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem21setFirstColumnSpannedEb
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setFirstColumnSpanned", args)
  }

}

  // proto:  int QTreeWidgetItem::indexOfChild(QTreeWidgetItem * child);
func (this *QTreeWidgetItem) indexOfChild(args ...interface{}) () {
  // indexOfChild(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem12indexOfChildEPS_
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "indexOfChild", args)
  }

}

  // proto:  QVariant QTreeWidgetItem::data(int column, int role);
func (this *QTreeWidgetItem) data(args ...interface{}) () {
  // data(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem4dataEii
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "data", args)
  }

}

  // proto:  QTreeWidgetItem * QTreeWidgetItem::parent();
func (this *QTreeWidgetItem) parent(args ...interface{}) () {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem6parentEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "parent", args)
  }

}

  // proto:  void QTreeWidgetItem::setFont(int column, const QFont & font);
func (this *QTreeWidgetItem) setFont(args ...interface{}) () {
  // setFont(int, const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem7setFontEiRK5QFont
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setFont", args)
  }

}

  // proto:  void QTreeWidgetItem::setData(int column, int role, const QVariant & value);
func (this *QTreeWidgetItem) setData(args ...interface{}) () {
  // setData(int, int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem7setDataEiiRK8QVariant
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setData", args)
  }

}

  // proto:  QFont QTreeWidgetItem::font(int column);
func (this *QTreeWidgetItem) font(args ...interface{}) () {
  // font(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem4fontEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "font", args)
  }

}

  // proto:  void QTreeWidgetItem::setStatusTip(int column, const QString & statusTip);
func (this *QTreeWidgetItem) setStatusTip(args ...interface{}) () {
  // setStatusTip(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem12setStatusTipEiRK7QString
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setStatusTip", args)
  }

}

  // proto:  void QTreeWidgetItem::setExpanded(bool expand);
func (this *QTreeWidgetItem) setExpanded(args ...interface{}) () {
  // setExpanded(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11setExpandedEb
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setExpanded", args)
  }

}

  // proto:  void QTreeWidgetItem::write(QDataStream & out);
func (this *QTreeWidgetItem) write(args ...interface{}) () {
  // write(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem5writeER11QDataStream
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "write", args)
  }

}

  // proto:  bool QTreeWidgetItem::isExpanded();
func (this *QTreeWidgetItem) isExpanded(args ...interface{}) () {
  // isExpanded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10isExpandedEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isExpanded", args)
  }

}

  // proto:  QList<QTreeWidgetItem *> QTreeWidgetItem::takeChildren();
func (this *QTreeWidgetItem) takeChildren(args ...interface{}) () {
  // takeChildren()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem12takeChildrenEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "takeChildren", args)
  }

}

  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidgetItem * parent, int type);
func NewQTreeWidgetItem(args ...interface{}) QTreeWidgetItem {
  return QTreeWidgetItem{}
}

  // proto:  void QTreeWidgetItem::setIcon(int column, const QIcon & icon);
func (this *QTreeWidgetItem) setIcon(args ...interface{}) () {
  // setIcon(int, const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem7setIconEiRK5QIcon
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setIcon", args)
  }

}

  // proto:  QString QTreeWidgetItem::toolTip(int column);
func (this *QTreeWidgetItem) toolTip(args ...interface{}) () {
  // toolTip(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem7toolTipEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "toolTip", args)
  }

}

  // proto:  QColor QTreeWidgetItem::backgroundColor(int column);
func (this *QTreeWidgetItem) backgroundColor(args ...interface{}) () {
  // backgroundColor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem15backgroundColorEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "backgroundColor", args)
  }

}

  // proto:  QString QTreeWidgetItem::text(int column);
func (this *QTreeWidgetItem) text(args ...interface{}) () {
  // text(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem4textEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "text", args)
  }

}

  // proto:  bool QTreeWidgetItem::isHidden();
func (this *QTreeWidgetItem) isHidden(args ...interface{}) () {
  // isHidden()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem8isHiddenEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isHidden", args)
  }

}

  // proto:  void QTreeWidgetItem::setTextAlignment(int column, int alignment);
func (this *QTreeWidgetItem) setTextAlignment(args ...interface{}) () {
  // setTextAlignment(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem16setTextAlignmentEii
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setTextAlignment", args)
  }

}

  // proto:  void QTreeWidgetItem::insertChild(int index, QTreeWidgetItem * child);
func (this *QTreeWidgetItem) insertChild(args ...interface{}) () {
  // insertChild(int, class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11insertChildEiPS_
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "insertChild", args)
  }

}

  // proto:  bool QTreeWidgetItem::isDisabled();
func (this *QTreeWidgetItem) isDisabled(args ...interface{}) () {
  // isDisabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10isDisabledEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isDisabled", args)
  }

}

  // proto:  void QTreeWidgetItem::setText(int column, const QString & text);
func (this *QTreeWidgetItem) setText(args ...interface{}) () {
  // setText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem7setTextEiRK7QString
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setText", args)
  }

}

  // proto:  void QTreeWidgetItem::setTextColor(int column, const QColor & color);
func (this *QTreeWidgetItem) setTextColor(args ...interface{}) () {
  // setTextColor(int, const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem12setTextColorEiRK6QColor
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setTextColor", args)
  }

}

  // proto:  QSize QTreeWidgetItem::sizeHint(int column);
func (this *QTreeWidgetItem) sizeHint(args ...interface{}) () {
  // sizeHint(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem8sizeHintEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "sizeHint", args)
  }

}

  // proto:  QString QTreeWidgetItem::whatsThis(int column);
func (this *QTreeWidgetItem) whatsThis(args ...interface{}) () {
  // whatsThis(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem9whatsThisEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "whatsThis", args)
  }

}

  // proto:  void QTreeWidgetItem::setWhatsThis(int column, const QString & whatsThis);
func (this *QTreeWidgetItem) setWhatsThis(args ...interface{}) () {
  // setWhatsThis(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem12setWhatsThisEiRK7QString
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setWhatsThis", args)
  }

}

  // proto:  QColor QTreeWidgetItem::textColor(int column);
func (this *QTreeWidgetItem) textColor(args ...interface{}) () {
  // textColor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem9textColorEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "textColor", args)
  }

}

  // proto:  QIcon QTreeWidgetItem::icon(int column);
func (this *QTreeWidgetItem) icon(args ...interface{}) () {
  // icon(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem4iconEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "icon", args)
  }

}

  // proto:  void QTreeWidgetItem::setToolTip(int column, const QString & toolTip);
func (this *QTreeWidgetItem) setToolTip(args ...interface{}) () {
  // setToolTip(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem10setToolTipEiRK7QString
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setToolTip", args)
  }

}

  // proto:  bool QTreeWidgetItem::isFirstColumnSpanned();
func (this *QTreeWidgetItem) isFirstColumnSpanned(args ...interface{}) () {
  // isFirstColumnSpanned()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem20isFirstColumnSpannedEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isFirstColumnSpanned", args)
  }

}

  // proto:  int QTreeWidgetItem::textAlignment(int column);
func (this *QTreeWidgetItem) textAlignment(args ...interface{}) () {
  // textAlignment(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem13textAlignmentEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "textAlignment", args)
  }

}

  // proto:  QTreeWidgetItem * QTreeWidgetItem::child(int index);
func (this *QTreeWidgetItem) child(args ...interface{}) () {
  // child(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem5childEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "child", args)
  }

}

  // proto:  void QTreeWidgetItem::setSelected(bool select);
func (this *QTreeWidgetItem) setSelected(args ...interface{}) () {
  // setSelected(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11setSelectedEb
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setSelected", args)
  }

}

  // proto:  void QTreeWidgetItem::~QTreeWidgetItem();
func (this *QTreeWidgetItem) FreeQTreeWidgetItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "~QTreeWidgetItem", args)
  }

}

  // proto:  void QTreeWidgetItem::setHidden(bool hide);
func (this *QTreeWidgetItem) setHidden(args ...interface{}) () {
  // setHidden(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem9setHiddenEb
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setHidden", args)
  }

}

  // proto:  int QTreeWidgetItem::columnCount();
func (this *QTreeWidgetItem) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem11columnCountEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "columnCount", args)
  }

}

  // proto:  QTreeWidgetItem * QTreeWidgetItem::takeChild(int index);
func (this *QTreeWidgetItem) takeChild(args ...interface{}) () {
  // takeChild(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem9takeChildEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "takeChild", args)
  }

}

  // proto:  void QTreeWidgetItem::setDisabled(bool disabled);
func (this *QTreeWidgetItem) setDisabled(args ...interface{}) () {
  // setDisabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11setDisabledEb
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setDisabled", args)
  }

}

  // proto:  void QTreeWidgetItem::setBackground(int column, const QBrush & brush);
func (this *QTreeWidgetItem) setBackground(args ...interface{}) () {
  // setBackground(int, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem13setBackgroundEiRK6QBrush
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setBackground", args)
  }

}

  // proto:  void QTreeWidgetItem::addChild(QTreeWidgetItem * child);
func (this *QTreeWidgetItem) addChild(args ...interface{}) () {
  // addChild(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem8addChildEPS_
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "addChild", args)
  }

}

  // proto:  void QTreeWidgetItem::removeChild(QTreeWidgetItem * child);
func (this *QTreeWidgetItem) removeChild(args ...interface{}) () {
  // removeChild(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11removeChildEPS_
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "removeChild", args)
  }

}

  // proto:  QTreeWidgetItem * QTreeWidgetItem::clone();
func (this *QTreeWidgetItem) clone(args ...interface{}) () {
  // clone()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem5cloneEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "clone", args)
  }

}

  // proto:  void QTreeWidgetItem::setSizeHint(int column, const QSize & size);
func (this *QTreeWidgetItem) setSizeHint(args ...interface{}) () {
  // setSizeHint(int, const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11setSizeHintEiRK5QSize
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setSizeHint", args)
  }

}

  // proto:  QBrush QTreeWidgetItem::foreground(int column);
func (this *QTreeWidgetItem) foreground(args ...interface{}) () {
  // foreground(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10foregroundEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "foreground", args)
  }

}

  // proto:  int QTreeWidgetItem::childCount();
func (this *QTreeWidgetItem) childCount(args ...interface{}) () {
  // childCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10childCountEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "childCount", args)
  }

}

  // proto:  void QTreeWidgetItem::setBackgroundColor(int column, const QColor & color);
func (this *QTreeWidgetItem) setBackgroundColor(args ...interface{}) () {
  // setBackgroundColor(int, const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem18setBackgroundColorEiRK6QColor
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setBackgroundColor", args)
  }

}

  // proto:  QString QTreeWidgetItem::statusTip(int column);
func (this *QTreeWidgetItem) statusTip(args ...interface{}) () {
  // statusTip(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem9statusTipEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "statusTip", args)
  }

}

  // proto:  QBrush QTreeWidgetItem::background(int column);
func (this *QTreeWidgetItem) background(args ...interface{}) () {
  // background(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10backgroundEi
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "background", args)
  }

}

  // proto:  int QTreeWidgetItem::type();
func (this *QTreeWidgetItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "type", args)
  }

}

  // proto:  QTreeWidget * QTreeWidgetItem::treeWidget();
func (this *QTreeWidgetItem) treeWidget(args ...interface{}) () {
  // treeWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10treeWidgetEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "treeWidget", args)
  }

}

  // proto:  void QTreeWidgetItem::read(QDataStream & in);
func (this *QTreeWidgetItem) read(args ...interface{}) () {
  // read(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem4readER11QDataStream
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "read", args)
  }

}

  // proto:  void QTreeWidgetItem::setForeground(int column, const QBrush & brush);
func (this *QTreeWidgetItem) setForeground(args ...interface{}) () {
  // setForeground(int, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem13setForegroundEiRK6QBrush
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setForeground", args)
  }

}

  // proto:  bool QTreeWidgetItem::isSelected();
func (this *QTreeWidgetItem) isSelected(args ...interface{}) () {
  // isSelected()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10isSelectedEv
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isSelected", args)
  }

}

// <= body block end
