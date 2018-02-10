package qtgui

// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin

type QTextBlock struct {
	*qtrt.CObject
}

func (this *QTextBlock) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextBlock) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextBlockFromPointer(cthis unsafe.Pointer) *QTextBlock {
	return &QTextBlock{&qtrt.CObject{cthis}}
}
func (*QTextBlock) NewFromPointer(cthis unsafe.Pointer) *QTextBlock {
	return NewQTextBlockFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextobject.h:206
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTextBlock()
func NewQTextBlock() *QTextBlock {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlockC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextBlock)
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:210
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTextBlock) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:216
// index:0
// Public Visibility=Default Availability=Available
// [4] int position()
func (this *QTextBlock) Position() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:217
// index:0
// Public Visibility=Default Availability=Available
// [4] int length()
func (this *QTextBlock) Length() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:218
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(int)
func (this *QTextBlock) Contains(position int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock8containsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:220
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextLayout * layout()
func (this *QTextBlock) Layout() *QTextLayout /*777 QTextLayout **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock6layoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextobject.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearLayout()
func (this *QTextBlock) ClearLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlock11clearLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:222
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlockFormat blockFormat()
func (this *QTextBlock) BlockFormat() *QTextBlockFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock11blockFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlockFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:223
// index:0
// Public Visibility=Default Availability=Available
// [4] int blockFormatIndex()
func (this *QTextBlock) BlockFormatIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock16blockFormatIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:224
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat charFormat()
func (this *QTextBlock) CharFormat() *QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock10charFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:225
// index:0
// Public Visibility=Default Availability=Available
// [4] int charFormatIndex()
func (this *QTextBlock) CharFormatIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock15charFormatIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:227
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::LayoutDirection textDirection()
func (this *QTextBlock) TextDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock13textDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextobject.h:229
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QTextBlock) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextobject.h:233
// index:0
// Public Visibility=Default Availability=Available
// [8] const QTextDocument * document()
func (this *QTextBlock) Document() *QTextDocument /*777 const QTextDocument **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock8documentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextobject.h:235
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextList * textList()
func (this *QTextBlock) TextList() *QTextList /*777 QTextList **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock8textListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextobject.h:237
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextBlockUserData * userData()
func (this *QTextBlock) UserData() *QTextBlockUserData /*777 QTextBlockUserData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock8userDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQTextBlockUserDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextobject.h:238
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserData(QTextBlockUserData *)
func (this *QTextBlock) SetUserData(data *QTextBlockUserData /*777 QTextBlockUserData **/) {
	var convArg0 = data.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlock11setUserDataEP18QTextBlockUserData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:240
// index:0
// Public Visibility=Default Availability=Available
// [4] int userState()
func (this *QTextBlock) UserState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock9userStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:241
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserState(int)
func (this *QTextBlock) SetUserState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlock12setUserStateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:243
// index:0
// Public Visibility=Default Availability=Available
// [4] int revision()
func (this *QTextBlock) Revision() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock8revisionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:244
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRevision(int)
func (this *QTextBlock) SetRevision(rev int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlock11setRevisionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rev)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:246
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible()
func (this *QTextBlock) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:247
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QTextBlock) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlock10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:249
// index:0
// Public Visibility=Default Availability=Available
// [4] int blockNumber()
func (this *QTextBlock) BlockNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock11blockNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:250
// index:0
// Public Visibility=Default Availability=Available
// [4] int firstLineNumber()
func (this *QTextBlock) FirstLineNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock15firstLineNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:252
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLineCount(int)
func (this *QTextBlock) SetLineCount(count int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlock12setLineCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:253
// index:0
// Public Visibility=Default Availability=Available
// [4] int lineCount()
func (this *QTextBlock) LineCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock9lineCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:283
// index:0
// Public Visibility=Default Availability=Available
// [24] QTextBlock::iterator begin()
func (this *QTextBlock) Begin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qtextobject.h:284
// index:0
// Public Visibility=Default Availability=Available
// [24] QTextBlock::iterator end()
func (this *QTextBlock) End() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qtextobject.h:286
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock next()
func (this *QTextBlock) Next() *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock4nextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:287
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock previous()
func (this *QTextBlock) Previous() *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock8previousEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:290
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int fragmentIndex()
func (this *QTextBlock) FragmentIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock13fragmentIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQTextBlock(this *QTextBlock) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlockD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end