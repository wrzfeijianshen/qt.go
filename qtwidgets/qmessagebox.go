package qtwidgets

// /usr/include/qt/QtWidgets/qmessagebox.h
// #include <qmessagebox.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
// bool event(class QEvent *)
func (this *QMessageBox) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QMessageBox) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void showEvent(class QShowEvent *)
func (this *QMessageBox) InheritShowEvent(f func(event *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void closeEvent(class QCloseEvent *)
func (this *QMessageBox) InheritCloseEvent(f func(event *qtgui.QCloseEvent /*777 QCloseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "closeEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QMessageBox) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QMessageBox) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

type QMessageBox struct {
	*QDialog
}

func (this *QMessageBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDialog.GetCthis()
	}
}
func (this *QMessageBox) SetCthis(cthis unsafe.Pointer) {
	this.QDialog = NewQDialogFromPointer(cthis)
}
func NewQMessageBoxFromPointer(cthis unsafe.Pointer) *QMessageBox {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QMessageBox{bcthis0}
}
func (*QMessageBox) NewFromPointer(cthis unsafe.Pointer) *QMessageBox {
	return NewQMessageBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QMessageBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmessagebox.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMessageBox(QWidget *)
func NewQMessageBox(parent *QWidget /*777 QWidget **/) *QMessageBox {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qmessagebox.h:136
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMessageBox(enum QMessageBox::Icon, const QString &, const QString &, QMessageBox::StandardButtons, QWidget *, Qt::WindowFlags)
func NewQMessageBox_1(icon int, title string, text string, buttons int, parent *QWidget /*777 QWidget **/, flags int) *QMessageBox {
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBoxC2ENS_4IconERK7QStringS3_6QFlagsINS_14StandardButtonEEP7QWidgetS4_IN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, icon, convArg1, convArg2, buttons, convArg4, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qmessagebox.h:202
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QMessageBox(const QString &, const QString &, enum QMessageBox::Icon, int, int, int, QWidget *, Qt::WindowFlags)
func NewQMessageBox_2(title string, text string, icon int, button0 int, button1 int, button2 int, parent *QWidget /*777 QWidget **/, f int) *QMessageBox {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg6 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBoxC2ERK7QStringS2_NS_4IconEiiiP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, icon, button0, button1, button2, convArg6, f)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qmessagebox.h:139
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMessageBox()
func DeleteQMessageBox(this *QMessageBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addButton(QAbstractButton *, enum QMessageBox::ButtonRole)
func (this *QMessageBox) AddButton(button *QAbstractButton /*777 QAbstractButton **/, role int) {
	var convArg0 = button.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox9addButtonEP15QAbstractButtonNS_10ButtonRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:142
// index:1
// Public Visibility=Default Availability=Available
// [8] QPushButton * addButton(const QString &, enum QMessageBox::ButtonRole)
func (this *QMessageBox) AddButton_1(text string, role int) *QPushButton /*777 QPushButton **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox9addButtonERK7QStringNS_10ButtonRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmessagebox.h:143
// index:2
// Public Visibility=Default Availability=Available
// [8] QPushButton * addButton(enum QMessageBox::StandardButton)
func (this *QMessageBox) AddButton_2(button int) *QPushButton /*777 QPushButton **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox9addButtonENS_14StandardButtonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmessagebox.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeButton(QAbstractButton *)
func (this *QMessageBox) RemoveButton(button *QAbstractButton /*777 QAbstractButton **/) {
	var convArg0 = button.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox12removeButtonEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void open(QObject *, const char *)
func (this *QMessageBox) Open(receiver *qtcore.QObject /*777 QObject **/, member string) {
	var convArg0 = receiver.GetCthis()
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox4openEP7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:150
// index:0
// Public Visibility=Default Availability=Available
// [4] QMessageBox::ButtonRole buttonRole(QAbstractButton *)
func (this *QMessageBox) ButtonRole(button *QAbstractButton /*777 QAbstractButton **/) int {
	var convArg0 = button.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox10buttonRoleEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardButtons(QMessageBox::StandardButtons)
func (this *QMessageBox) SetStandardButtons(buttons int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox18setStandardButtonsE6QFlagsINS_14StandardButtonEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buttons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:153
// index:0
// Public Visibility=Default Availability=Available
// [4] QMessageBox::StandardButtons standardButtons()
func (this *QMessageBox) StandardButtons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox15standardButtonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:154
// index:0
// Public Visibility=Default Availability=Available
// [4] QMessageBox::StandardButton standardButton(QAbstractButton *)
func (this *QMessageBox) StandardButton(button *QAbstractButton /*777 QAbstractButton **/) int {
	var convArg0 = button.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox14standardButtonEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:155
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractButton * button(enum QMessageBox::StandardButton)
func (this *QMessageBox) Button(which int) *QAbstractButton /*777 QAbstractButton **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox6buttonENS_14StandardButtonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQAbstractButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmessagebox.h:157
// index:0
// Public Visibility=Default Availability=Available
// [8] QPushButton * defaultButton()
func (this *QMessageBox) DefaultButton() *QPushButton /*777 QPushButton **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox13defaultButtonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmessagebox.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultButton(QPushButton *)
func (this *QMessageBox) SetDefaultButton(button *QPushButton /*777 QPushButton **/) {
	var convArg0 = button.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox16setDefaultButtonEP11QPushButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:159
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setDefaultButton(enum QMessageBox::StandardButton)
func (this *QMessageBox) SetDefaultButton_1(button int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox16setDefaultButtonENS_14StandardButtonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:161
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractButton * escapeButton()
func (this *QMessageBox) EscapeButton() *QAbstractButton /*777 QAbstractButton **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox12escapeButtonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQAbstractButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmessagebox.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEscapeButton(QAbstractButton *)
func (this *QMessageBox) SetEscapeButton(button *QAbstractButton /*777 QAbstractButton **/) {
	var convArg0 = button.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox15setEscapeButtonEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:163
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setEscapeButton(enum QMessageBox::StandardButton)
func (this *QMessageBox) SetEscapeButton_1(button int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox15setEscapeButtonENS_14StandardButtonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:165
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractButton * clickedButton()
func (this *QMessageBox) ClickedButton() *QAbstractButton /*777 QAbstractButton **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox13clickedButtonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQAbstractButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmessagebox.h:167
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QMessageBox) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qmessagebox.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QMessageBox) SetText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:170
// index:0
// Public Visibility=Default Availability=Available
// [4] QMessageBox::Icon icon()
func (this *QMessageBox) Icon() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:171
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(enum QMessageBox::Icon)
func (this *QMessageBox) SetIcon(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox7setIconENS_4IconE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:173
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap iconPixmap()
func (this *QMessageBox) IconPixmap() *qtgui.QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox10iconPixmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qmessagebox.h:174
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconPixmap(const QPixmap &)
func (this *QMessageBox) SetIconPixmap(pixmap *qtgui.QPixmap) {
	var convArg0 = pixmap.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox13setIconPixmapERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:176
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextFormat textFormat()
func (this *QMessageBox) TextFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox10textFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextFormat(Qt::TextFormat)
func (this *QMessageBox) SetTextFormat(format int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox13setTextFormatEN2Qt10TextFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextInteractionFlags(Qt::TextInteractionFlags)
func (this *QMessageBox) SetTextInteractionFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:180
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextInteractionFlags textInteractionFlags()
func (this *QMessageBox) TextInteractionFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox20textInteractionFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCheckBox(QCheckBox *)
func (this *QMessageBox) SetCheckBox(cb *QCheckBox /*777 QCheckBox **/) {
	var convArg0 = cb.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox11setCheckBoxEP9QCheckBox", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:183
// index:0
// Public Visibility=Default Availability=Available
// [8] QCheckBox * checkBox()
func (this *QMessageBox) CheckBox() *QCheckBox /*777 QCheckBox **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox8checkBoxEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQCheckBoxFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmessagebox.h:185
// index:0
// Public static Visibility=Default Availability=Available
// [4] QMessageBox::StandardButton information(QWidget *, const QString &, const QString &, QMessageBox::StandardButtons, enum QMessageBox::StandardButton)
func (this *QMessageBox) Information(parent *QWidget /*777 QWidget **/, title string, text string, buttons int, defaultButton int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_6QFlagsINS_14StandardButtonEES6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, buttons, defaultButton)
	gopp.ErrPrint(err, rv)
	return int(rv)
}
func QMessageBox_Information(parent *QWidget /*777 QWidget **/, title string, text string, buttons int, defaultButton int) int {
	var nilthis *QMessageBox
	rv := nilthis.Information(parent, title, text, buttons, defaultButton)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:207
// index:1
// Public static Visibility=Default Availability=Available
// [4] int information(QWidget *, const QString &, const QString &, int, int, int)
func (this *QMessageBox) Information_1(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int, button2 int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_iii", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, button0, button1, button2)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QMessageBox_Information_1(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int, button2 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Information_1(parent, title, text, button0, button1, button2)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:210
// index:2
// Public static Visibility=Default Availability=Available
// [4] int information(QWidget *, const QString &, const QString &, const QString &, const QString &, const QString &, int, int)
func (this *QMessageBox) Information_2(parent *QWidget /*777 QWidget **/, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(button0Text)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(button1Text)
	var convArg4 = tmpArg4.GetCthis()
	var tmpArg5 = qtcore.NewQString_5(button2Text)
	var convArg5 = tmpArg5.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_S4_S4_S4_ii", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, defaultButtonNumber, escapeButtonNumber)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QMessageBox_Information_2(parent *QWidget /*777 QWidget **/, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	var nilthis *QMessageBox
	rv := nilthis.Information_2(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:217
// index:3
// Public static inline Visibility=Default Availability=Available
// [4] QMessageBox::StandardButton information(QWidget *, const QString &, const QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Information_3(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, button0, button1)
	gopp.ErrPrint(err, rv)
	return int(rv)
}
func QMessageBox_Information_3(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Information_3(parent, title, text, button0, button1)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:188
// index:0
// Public static Visibility=Default Availability=Available
// [4] QMessageBox::StandardButton question(QWidget *, const QString &, const QString &, QMessageBox::StandardButtons, enum QMessageBox::StandardButton)
func (this *QMessageBox) Question(parent *QWidget /*777 QWidget **/, title string, text string, buttons int, defaultButton int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_6QFlagsINS_14StandardButtonEES6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, buttons, defaultButton)
	gopp.ErrPrint(err, rv)
	return int(rv)
}
func QMessageBox_Question(parent *QWidget /*777 QWidget **/, title string, text string, buttons int, defaultButton int) int {
	var nilthis *QMessageBox
	rv := nilthis.Question(parent, title, text, buttons, defaultButton)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:222
// index:1
// Public static Visibility=Default Availability=Available
// [4] int question(QWidget *, const QString &, const QString &, int, int, int)
func (this *QMessageBox) Question_1(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int, button2 int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_iii", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, button0, button1, button2)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QMessageBox_Question_1(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int, button2 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Question_1(parent, title, text, button0, button1, button2)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:225
// index:2
// Public static Visibility=Default Availability=Available
// [4] int question(QWidget *, const QString &, const QString &, const QString &, const QString &, const QString &, int, int)
func (this *QMessageBox) Question_2(parent *QWidget /*777 QWidget **/, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(button0Text)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(button1Text)
	var convArg4 = tmpArg4.GetCthis()
	var tmpArg5 = qtcore.NewQString_5(button2Text)
	var convArg5 = tmpArg5.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_S4_S4_S4_ii", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, defaultButtonNumber, escapeButtonNumber)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QMessageBox_Question_2(parent *QWidget /*777 QWidget **/, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	var nilthis *QMessageBox
	rv := nilthis.Question_2(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:232
// index:3
// Public static inline Visibility=Default Availability=Available
// [4] int question(QWidget *, const QString &, const QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Question_3(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, button0, button1)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QMessageBox_Question_3(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Question_3(parent, title, text, button0, button1)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:191
// index:0
// Public static Visibility=Default Availability=Available
// [4] QMessageBox::StandardButton warning(QWidget *, const QString &, const QString &, QMessageBox::StandardButtons, enum QMessageBox::StandardButton)
func (this *QMessageBox) Warning(parent *QWidget /*777 QWidget **/, title string, text string, buttons int, defaultButton int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_6QFlagsINS_14StandardButtonEES6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, buttons, defaultButton)
	gopp.ErrPrint(err, rv)
	return int(rv)
}
func QMessageBox_Warning(parent *QWidget /*777 QWidget **/, title string, text string, buttons int, defaultButton int) int {
	var nilthis *QMessageBox
	rv := nilthis.Warning(parent, title, text, buttons, defaultButton)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:237
// index:1
// Public static Visibility=Default Availability=Available
// [4] int warning(QWidget *, const QString &, const QString &, int, int, int)
func (this *QMessageBox) Warning_1(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int, button2 int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_iii", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, button0, button1, button2)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QMessageBox_Warning_1(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int, button2 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Warning_1(parent, title, text, button0, button1, button2)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:240
// index:2
// Public static Visibility=Default Availability=Available
// [4] int warning(QWidget *, const QString &, const QString &, const QString &, const QString &, const QString &, int, int)
func (this *QMessageBox) Warning_2(parent *QWidget /*777 QWidget **/, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(button0Text)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(button1Text)
	var convArg4 = tmpArg4.GetCthis()
	var tmpArg5 = qtcore.NewQString_5(button2Text)
	var convArg5 = tmpArg5.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_S4_S4_S4_ii", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, defaultButtonNumber, escapeButtonNumber)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QMessageBox_Warning_2(parent *QWidget /*777 QWidget **/, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	var nilthis *QMessageBox
	rv := nilthis.Warning_2(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:247
// index:3
// Public static inline Visibility=Default Availability=Available
// [4] int warning(QWidget *, const QString &, const QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Warning_3(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, button0, button1)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QMessageBox_Warning_3(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Warning_3(parent, title, text, button0, button1)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:194
// index:0
// Public static Visibility=Default Availability=Available
// [4] QMessageBox::StandardButton critical(QWidget *, const QString &, const QString &, QMessageBox::StandardButtons, enum QMessageBox::StandardButton)
func (this *QMessageBox) Critical(parent *QWidget /*777 QWidget **/, title string, text string, buttons int, defaultButton int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_6QFlagsINS_14StandardButtonEES6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, buttons, defaultButton)
	gopp.ErrPrint(err, rv)
	return int(rv)
}
func QMessageBox_Critical(parent *QWidget /*777 QWidget **/, title string, text string, buttons int, defaultButton int) int {
	var nilthis *QMessageBox
	rv := nilthis.Critical(parent, title, text, buttons, defaultButton)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:252
// index:1
// Public static Visibility=Default Availability=Available
// [4] int critical(QWidget *, const QString &, const QString &, int, int, int)
func (this *QMessageBox) Critical_1(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int, button2 int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_iii", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, button0, button1, button2)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QMessageBox_Critical_1(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int, button2 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Critical_1(parent, title, text, button0, button1, button2)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:255
// index:2
// Public static Visibility=Default Availability=Available
// [4] int critical(QWidget *, const QString &, const QString &, const QString &, const QString &, const QString &, int, int)
func (this *QMessageBox) Critical_2(parent *QWidget /*777 QWidget **/, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(button0Text)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(button1Text)
	var convArg4 = tmpArg4.GetCthis()
	var tmpArg5 = qtcore.NewQString_5(button2Text)
	var convArg5 = tmpArg5.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_S4_S4_S4_ii", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, defaultButtonNumber, escapeButtonNumber)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QMessageBox_Critical_2(parent *QWidget /*777 QWidget **/, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	var nilthis *QMessageBox
	rv := nilthis.Critical_2(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:262
// index:3
// Public static inline Visibility=Default Availability=Available
// [4] int critical(QWidget *, const QString &, const QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Critical_3(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int) int {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, button0, button1)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QMessageBox_Critical_3(parent *QWidget /*777 QWidget **/, title string, text string, button0 int, button1 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Critical_3(parent, title, text, button0, button1)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:197
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void about(QWidget *, const QString &, const QString &)
func (this *QMessageBox) About(parent *QWidget /*777 QWidget **/, title string, text string) {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox5aboutEP7QWidgetRK7QStringS4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_About(parent *QWidget /*777 QWidget **/, title string, text string) {
	var nilthis *QMessageBox
	nilthis.About(parent, title, text)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:198
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void aboutQt(QWidget *, const QString &)
func (this *QMessageBox) AboutQt(parent *QWidget /*777 QWidget **/, title string) {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox7aboutQtEP7QWidgetRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_AboutQt(parent *QWidget /*777 QWidget **/, title string) {
	var nilthis *QMessageBox
	nilthis.AboutQt(parent, title)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:267
// index:0
// Public Visibility=Default Availability=Available
// [8] QString buttonText(int)
func (this *QMessageBox) ButtonText(button int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox10buttonTextEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qmessagebox.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButtonText(int, const QString &)
func (this *QMessageBox) SetButtonText(button int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox13setButtonTextEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), button, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:270
// index:0
// Public Visibility=Default Availability=Available
// [8] QString informativeText()
func (this *QMessageBox) InformativeText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox15informativeTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qmessagebox.h:271
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInformativeText(const QString &)
func (this *QMessageBox) SetInformativeText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox18setInformativeTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:274
// index:0
// Public Visibility=Default Availability=Available
// [8] QString detailedText()
func (this *QMessageBox) DetailedText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMessageBox12detailedTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qmessagebox.h:275
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDetailedText(const QString &)
func (this *QMessageBox) SetDetailedText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox15setDetailedTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:278
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowTitle(const QString &)
func (this *QMessageBox) SetWindowTitle(title string) {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox14setWindowTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:279
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowModality(Qt::WindowModality)
func (this *QMessageBox) SetWindowModality(windowModality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox17setWindowModalityEN2Qt14WindowModalityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), windowModality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:282
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap standardIcon(enum QMessageBox::Icon)
func (this *QMessageBox) StandardIcon(icon int) *qtgui.QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox12standardIconENS_4IconE", qtrt.FFI_TYPE_POINTER, icon)
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}
func QMessageBox_StandardIcon(icon int) *qtgui.QPixmap /*123*/ {
	var nilthis *QMessageBox
	rv := nilthis.StandardIcon(icon)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:285
// index:0
// Public Visibility=Default Availability=Available
// [-2] void buttonClicked(QAbstractButton *)
func (this *QMessageBox) ButtonClicked(button *QAbstractButton /*777 QAbstractButton **/) {
	var convArg0 = button.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox13buttonClickedEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:293
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QMessageBox) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmessagebox.h:294
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QMessageBox) ResizeEvent(event *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:295
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QMessageBox) ShowEvent(event *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:296
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEvent(QCloseEvent *)
func (this *QMessageBox) CloseEvent(event *qtgui.QCloseEvent /*777 QCloseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox10closeEventEP11QCloseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:297
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QMessageBox) KeyPressEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:298
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QMessageBox) ChangeEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMessageBox11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QMessageBox__Icon = int

const QMessageBox__NoIcon QMessageBox__Icon = 0
const QMessageBox__Information QMessageBox__Icon = 1
const QMessageBox__Warning QMessageBox__Icon = 2
const QMessageBox__Critical QMessageBox__Icon = 3
const QMessageBox__Question QMessageBox__Icon = 4

type QMessageBox__ButtonRole = int

const QMessageBox__InvalidRole QMessageBox__ButtonRole = -1
const QMessageBox__AcceptRole QMessageBox__ButtonRole = 0
const QMessageBox__RejectRole QMessageBox__ButtonRole = 1
const QMessageBox__DestructiveRole QMessageBox__ButtonRole = 2
const QMessageBox__ActionRole QMessageBox__ButtonRole = 3
const QMessageBox__HelpRole QMessageBox__ButtonRole = 4
const QMessageBox__YesRole QMessageBox__ButtonRole = 5
const QMessageBox__NoRole QMessageBox__ButtonRole = 6
const QMessageBox__ResetRole QMessageBox__ButtonRole = 7
const QMessageBox__ApplyRole QMessageBox__ButtonRole = 8
const QMessageBox__NRoles QMessageBox__ButtonRole = 9

type QMessageBox__StandardButton = int

const QMessageBox__NoButton QMessageBox__StandardButton = 0
const QMessageBox__Ok QMessageBox__StandardButton = 1024
const QMessageBox__Save QMessageBox__StandardButton = 2048
const QMessageBox__SaveAll QMessageBox__StandardButton = 4096
const QMessageBox__Open QMessageBox__StandardButton = 8192
const QMessageBox__Yes QMessageBox__StandardButton = 16384
const QMessageBox__YesToAll QMessageBox__StandardButton = 32768
const QMessageBox__No QMessageBox__StandardButton = 65536
const QMessageBox__NoToAll QMessageBox__StandardButton = 131072
const QMessageBox__Abort QMessageBox__StandardButton = 262144
const QMessageBox__Retry QMessageBox__StandardButton = 524288
const QMessageBox__Ignore QMessageBox__StandardButton = 1048576
const QMessageBox__Close QMessageBox__StandardButton = 2097152
const QMessageBox__Cancel QMessageBox__StandardButton = 4194304
const QMessageBox__Discard QMessageBox__StandardButton = 8388608
const QMessageBox__Help QMessageBox__StandardButton = 16777216
const QMessageBox__Apply QMessageBox__StandardButton = 33554432
const QMessageBox__Reset QMessageBox__StandardButton = 67108864
const QMessageBox__RestoreDefaults QMessageBox__StandardButton = 134217728
const QMessageBox__FirstButton QMessageBox__StandardButton = 1024
const QMessageBox__LastButton QMessageBox__StandardButton = 134217728
const QMessageBox__YesAll QMessageBox__StandardButton = 32768
const QMessageBox__NoAll QMessageBox__StandardButton = 131072
const QMessageBox__Default QMessageBox__StandardButton = 256
const QMessageBox__Escape QMessageBox__StandardButton = 512
const QMessageBox__FlagMask QMessageBox__StandardButton = 768
const QMessageBox__ButtonMask QMessageBox__StandardButton = -769

//  body block end