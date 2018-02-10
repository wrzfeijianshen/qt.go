package qtcore

// /usr/include/qt/QtCore/qtextstream.h
// #include <qtextstream.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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

type QTextStream struct {
	*qtrt.CObject
}

func (this *QTextStream) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextStream) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextStreamFromPointer(cthis unsafe.Pointer) *QTextStream {
	return &QTextStream{&qtrt.CObject{cthis}}
}
func (*QTextStream) NewFromPointer(cthis unsafe.Pointer) *QTextStream {
	return NewQTextStreamFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtextstream.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextStream()
func NewQTextStream() *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:94
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextStream(QIODevice *)
func NewQTextStream_1(device *QIODevice /*777 QIODevice **/) *QTextStream {
	var convArg0 = device.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2EP9QIODevice", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:96
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextStream(QString *, QIODevice::OpenMode)
func NewQTextStream_2(string string, openMode int) *QTextStream {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2EP7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, openMode)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:97
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTextStream(QByteArray *, QIODevice::OpenMode)
func NewQTextStream_3(array *QByteArray /*777 QByteArray **/, openMode int) *QTextStream {
	var convArg0 = array.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2EP10QByteArray6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, openMode)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:98
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QTextStream(const QByteArray &, QIODevice::OpenMode)
func NewQTextStream_4(array *QByteArray, openMode int) *QTextStream {
	var convArg0 = array.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2ERK10QByteArray6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, openMode)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTextStream()
func DeleteQTextStream(this *QTextStream) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtextstream.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCodec(QTextCodec *)
func (this *QTextStream) SetCodec(codec *QTextCodec /*777 QTextCodec **/) {
	var convArg0 = codec.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream8setCodecEP10QTextCodec", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:103
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setCodec(const char *)
func (this *QTextStream) SetCodec_1(codecName string) {
	var convArg0 = qtrt.CString(codecName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream8setCodecEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCodec * codec()
func (this *QTextStream) Codec() *QTextCodec /*777 QTextCodec **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream5codecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qtextstream.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoDetectUnicode(_Bool)
func (this *QTextStream) SetAutoDetectUnicode(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream20setAutoDetectUnicodeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoDetectUnicode()
func (this *QTextStream) AutoDetectUnicode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream17autoDetectUnicodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGenerateByteOrderMark(_Bool)
func (this *QTextStream) SetGenerateByteOrderMark(generate bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream24setGenerateByteOrderMarkEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), generate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:108
// index:0
// Public Visibility=Default Availability=Available
// [1] bool generateByteOrderMark()
func (this *QTextStream) GenerateByteOrderMark() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream21generateByteOrderMarkEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocale(const QLocale &)
func (this *QTextStream) SetLocale(locale *QLocale) {
	var convArg0 = locale.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream9setLocaleERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QLocale locale()
func (this *QTextStream) Locale() *QLocale /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream6localeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLocale)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)
func (this *QTextStream) SetDevice(device *QIODevice /*777 QIODevice **/) {
	var convArg0 = device.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device()
func (this *QTextStream) Device() *QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qtextstream.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setString(QString *, QIODevice::OpenMode)
func (this *QTextStream) SetString(string string, openMode int) {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream9setStringEP7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, openMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QString * string()
func (this *QTextStream) String() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream6stringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextstream.h:120
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextStream::Status status()
func (this *QTextStream) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtextstream.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStatus(enum QTextStream::Status)
func (this *QTextStream) SetStatus(status int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream9setStatusENS_6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), status)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetStatus()
func (this *QTextStream) ResetStatus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream11resetStatusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:124
// index:0
// Public Visibility=Default Availability=Available
// [1] bool atEnd()
func (this *QTextStream) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()
func (this *QTextStream) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void flush()
func (this *QTextStream) Flush() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:127
// index:0
// Public Visibility=Default Availability=Available
// [1] bool seek(qint64)
func (this *QTextStream) Seek(pos int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream4seekEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 pos()
func (this *QTextStream) Pos() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qtextstream.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void skipWhiteSpace()
func (this *QTextStream) SkipWhiteSpace() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream14skipWhiteSpaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QString readLine(qint64)
func (this *QTextStream) ReadLine(maxlen int64) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream8readLineEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxlen)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextstream.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool readLineInto(QString *, qint64)
func (this *QTextStream) ReadLineInto(line string, maxlen int64) bool {
	var tmpArg0 = NewQString_5(line)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream12readLineIntoEP7QStringx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QString readAll()
func (this *QTextStream) ReadAll() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream7readAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextstream.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QString read(qint64)
func (this *QTextStream) Read(maxlen int64) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream4readEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxlen)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextstream.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFieldAlignment(enum QTextStream::FieldAlignment)
func (this *QTextStream) SetFieldAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream17setFieldAlignmentENS_14FieldAlignmentE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:138
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextStream::FieldAlignment fieldAlignment()
func (this *QTextStream) FieldAlignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream14fieldAlignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtextstream.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPadChar(QChar)
func (this *QTextStream) SetPadChar(ch *QChar /*123*/) {
	var convArg0 = ch.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream10setPadCharE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:141
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar padChar()
func (this *QTextStream) PadChar() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream7padCharEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFieldWidth(int)
func (this *QTextStream) SetFieldWidth(width int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream13setFieldWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:144
// index:0
// Public Visibility=Default Availability=Available
// [4] int fieldWidth()
func (this *QTextStream) FieldWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream10fieldWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextstream.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNumberFlags(QTextStream::NumberFlags)
func (this *QTextStream) SetNumberFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream14setNumberFlagsE6QFlagsINS_10NumberFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:147
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextStream::NumberFlags numberFlags()
func (this *QTextStream) NumberFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream11numberFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtextstream.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIntegerBase(int)
func (this *QTextStream) SetIntegerBase(base int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream14setIntegerBaseEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:150
// index:0
// Public Visibility=Default Availability=Available
// [4] int integerBase()
func (this *QTextStream) IntegerBase() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream11integerBaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextstream.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRealNumberNotation(enum QTextStream::RealNumberNotation)
func (this *QTextStream) SetRealNumberNotation(notation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream21setRealNumberNotationENS_18RealNumberNotationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), notation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:153
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextStream::RealNumberNotation realNumberNotation()
func (this *QTextStream) RealNumberNotation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream18realNumberNotationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtextstream.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRealNumberPrecision(int)
func (this *QTextStream) SetRealNumberPrecision(precision int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream22setRealNumberPrecisionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), precision)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:156
// index:0
// Public Visibility=Default Availability=Available
// [4] int realNumberPrecision()
func (this *QTextStream) RealNumberPrecision() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream19realNumberPrecisionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

type QTextStream__RealNumberNotation = int

const QTextStream__SmartNotation QTextStream__RealNumberNotation = 0
const QTextStream__FixedNotation QTextStream__RealNumberNotation = 1
const QTextStream__ScientificNotation QTextStream__RealNumberNotation = 2

type QTextStream__FieldAlignment = int

const QTextStream__AlignLeft QTextStream__FieldAlignment = 0
const QTextStream__AlignRight QTextStream__FieldAlignment = 1
const QTextStream__AlignCenter QTextStream__FieldAlignment = 2
const QTextStream__AlignAccountingStyle QTextStream__FieldAlignment = 3

type QTextStream__Status = int

const QTextStream__Ok QTextStream__Status = 0
const QTextStream__ReadPastEnd QTextStream__Status = 1
const QTextStream__ReadCorruptData QTextStream__Status = 2
const QTextStream__WriteFailed QTextStream__Status = 3

type QTextStream__NumberFlag = int

const QTextStream__ShowBase QTextStream__NumberFlag = 1
const QTextStream__ForcePoint QTextStream__NumberFlag = 2
const QTextStream__ForceSign QTextStream__NumberFlag = 4
const QTextStream__UppercaseBase QTextStream__NumberFlag = 8
const QTextStream__UppercaseDigits QTextStream__NumberFlag = 16

//  body block end