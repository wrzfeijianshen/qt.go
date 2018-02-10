package qtcore

// /usr/include/qt/QtCore/qcommandlineparser.h
// #include <qcommandlineparser.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 52
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

type QCommandLineParser struct {
	*qtrt.CObject
}

func (this *QCommandLineParser) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCommandLineParser) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCommandLineParserFromPointer(cthis unsafe.Pointer) *QCommandLineParser {
	return &QCommandLineParser{&qtrt.CObject{cthis}}
}
func (*QCommandLineParser) NewFromPointer(cthis unsafe.Pointer) *QCommandLineParser {
	return NewQCommandLineParserFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCommandLineParser()
func NewQCommandLineParser() *QCommandLineParser {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParserC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLineParserFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCommandLineParser)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineparser.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCommandLineParser()
func DeleteQCommandLineParser(this *QCommandLineParser) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParserD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSingleDashWordOptionMode(enum QCommandLineParser::SingleDashWordOptionMode)
func (this *QCommandLineParser) SetSingleDashWordOptionMode(parsingMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser27setSingleDashWordOptionModeENS_24SingleDashWordOptionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parsingMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptionsAfterPositionalArgumentsMode(enum QCommandLineParser::OptionsAfterPositionalArgumentsMode)
func (this *QCommandLineParser) SetOptionsAfterPositionalArgumentsMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser38setOptionsAfterPositionalArgumentsModeENS_35OptionsAfterPositionalArgumentsModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool addOption(const QCommandLineOption &)
func (this *QCommandLineParser) AddOption(commandLineOption *QCommandLineOption) bool {
	var convArg0 = commandLineOption.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser9addOptionERK18QCommandLineOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcommandlineparser.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QCommandLineOption addVersionOption()
func (this *QCommandLineParser) AddVersionOption() *QCommandLineOption /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser16addVersionOptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCommandLineOption)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QCommandLineOption addHelpOption()
func (this *QCommandLineParser) AddHelpOption() *QCommandLineOption /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser13addHelpOptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCommandLineOption)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setApplicationDescription(const QString &)
func (this *QCommandLineParser) SetApplicationDescription(description string) {
	var tmpArg0 = NewQString_5(description)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser25setApplicationDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QString applicationDescription()
func (this *QCommandLineParser) ApplicationDescription() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser22applicationDescriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcommandlineparser.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPositionalArgument(const QString &, const QString &, const QString &)
func (this *QCommandLineParser) AddPositionalArgument(name string, description string, syntax string) {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(description)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(syntax)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser21addPositionalArgumentERK7QStringS2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearPositionalArguments()
func (this *QCommandLineParser) ClearPositionalArguments() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser24clearPositionalArgumentsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void process(const QStringList &)
func (this *QCommandLineParser) Process(arguments *QStringList) {
	var convArg0 = arguments.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser7processERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:85
// index:1
// Public Visibility=Default Availability=Available
// [-2] void process(const QCoreApplication &)
func (this *QCommandLineParser) Process_1(app *QCoreApplication) {
	var convArg0 = app.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser7processERK16QCoreApplication", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:87
// index:0
// Public Visibility=Default Availability=Available
// [1] bool parse(const QStringList &)
func (this *QCommandLineParser) Parse(arguments *QStringList) bool {
	var convArg0 = arguments.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser5parseERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcommandlineparser.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorText()
func (this *QCommandLineParser) ErrorText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser9errorTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcommandlineparser.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSet(const QString &)
func (this *QCommandLineParser) IsSet(name string) bool {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser5isSetERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcommandlineparser.h:94
// index:1
// Public Visibility=Default Availability=Available
// [1] bool isSet(const QCommandLineOption &)
func (this *QCommandLineParser) IsSet_1(option *QCommandLineOption) bool {
	var convArg0 = option.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser5isSetERK18QCommandLineOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcommandlineparser.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QString value(const QString &)
func (this *QCommandLineParser) Value(name string) string {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser5valueERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcommandlineparser.h:95
// index:1
// Public Visibility=Default Availability=Available
// [8] QString value(const QCommandLineOption &)
func (this *QCommandLineParser) Value_1(option *QCommandLineOption) string {
	var convArg0 = option.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser5valueERK18QCommandLineOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcommandlineparser.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList values(const QString &)
func (this *QCommandLineParser) Values(name string) *QStringList /*123*/ {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser6valuesERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:96
// index:1
// Public Visibility=Default Availability=Available
// [8] QStringList values(const QCommandLineOption &)
func (this *QCommandLineParser) Values_1(option *QCommandLineOption) *QStringList /*123*/ {
	var convArg0 = option.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser6valuesERK18QCommandLineOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList positionalArguments()
func (this *QCommandLineParser) PositionalArguments() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser19positionalArgumentsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:99
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList optionNames()
func (this *QCommandLineParser) OptionNames() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser11optionNamesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList unknownOptionNames()
func (this *QCommandLineParser) UnknownOptionNames() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser18unknownOptionNamesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showVersion()
func (this *QCommandLineParser) ShowVersion() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser11showVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showHelp(int)
func (this *QCommandLineParser) ShowHelp(exitCode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser8showHelpEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), exitCode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QString helpText()
func (this *QCommandLineParser) HelpText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser8helpTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

type QCommandLineParser__SingleDashWordOptionMode = int

const QCommandLineParser__ParseAsCompactedShortOptions QCommandLineParser__SingleDashWordOptionMode = 0
const QCommandLineParser__ParseAsLongOptions QCommandLineParser__SingleDashWordOptionMode = 1

type QCommandLineParser__OptionsAfterPositionalArgumentsMode = int

const QCommandLineParser__ParseAsOptions QCommandLineParser__OptionsAfterPositionalArgumentsMode = 0
const QCommandLineParser__ParseAsPositionalArguments QCommandLineParser__OptionsAfterPositionalArgumentsMode = 1

//  body block end
