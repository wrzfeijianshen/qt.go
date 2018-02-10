package qtcore

// /usr/include/qt/QtCore/qtimezone.h
// #include <qtimezone.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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

type QTimeZone struct {
	*qtrt.CObject
}

func (this *QTimeZone) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTimeZone) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTimeZoneFromPointer(cthis unsafe.Pointer) *QTimeZone {
	return &QTimeZone{&qtrt.CObject{cthis}}
}
func (*QTimeZone) NewFromPointer(cthis unsafe.Pointer) *QTimeZone {
	return NewQTimeZoneFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtimezone.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTimeZone()
func NewQTimeZone() *QTimeZone {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZoneC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTimeZone)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:93
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTimeZone(const QByteArray &)
func NewQTimeZone_1(ianaId *QByteArray) *QTimeZone {
	var convArg0 = ianaId.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZoneC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTimeZone)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:94
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTimeZone(int)
func NewQTimeZone_2(offsetSeconds int) *QTimeZone {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZoneC2Ei", qtrt.FFI_TYPE_POINTER, offsetSeconds)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTimeZone)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:95
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTimeZone(const QByteArray &, int, const QString &, const QString &, QLocale::Country, const QString &)
func NewQTimeZone_3(zoneId *QByteArray, offsetSeconds int, name string, abbreviation string, country int, comment string) *QTimeZone {
	var convArg0 = zoneId.GetCthis()
	var tmpArg2 = NewQString_5(name)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString_5(abbreviation)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg5 = NewQString_5(comment)
	var convArg5 = tmpArg5.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZoneC2ERK10QByteArrayiRK7QStringS5_N7QLocale7CountryES5_", qtrt.FFI_TYPE_POINTER, convArg0, offsetSeconds, convArg2, convArg3, country, convArg5)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTimeZone)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTimeZone()
func DeleteQTimeZone(this *QTimeZone) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZoneD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtimezone.h:106
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QTimeZone &)
func (this *QTimeZone) Swap(other *QTimeZone) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTimeZone) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtimezone.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray id()
func (this *QTimeZone) Id() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone2idEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qtimezone.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] QLocale::Country country()
func (this *QTimeZone) Country() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone7countryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtimezone.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QString comment()
func (this *QTimeZone) Comment() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone7commentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtimezone.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QString displayName(const QDateTime &, QTimeZone::NameType, const QLocale &)
func (this *QTimeZone) DisplayName(atDateTime *QDateTime, nameType int, locale *QLocale) string {
	var convArg0 = atDateTime.GetCthis()
	var convArg2 = locale.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone11displayNameERK9QDateTimeNS_8NameTypeERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, nameType, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtimezone.h:121
// index:1
// Public Visibility=Default Availability=Available
// [8] QString displayName(QTimeZone::TimeType, QTimeZone::NameType, const QLocale &)
func (this *QTimeZone) DisplayName_1(timeType int, nameType int, locale *QLocale) string {
	var convArg2 = locale.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone11displayNameENS_8TimeTypeENS_8NameTypeERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeType, nameType, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtimezone.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] QString abbreviation(const QDateTime &)
func (this *QTimeZone) Abbreviation(atDateTime *QDateTime) string {
	var convArg0 = atDateTime.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone12abbreviationERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtimezone.h:126
// index:0
// Public Visibility=Default Availability=Available
// [4] int offsetFromUtc(const QDateTime &)
func (this *QTimeZone) OffsetFromUtc(atDateTime *QDateTime) int {
	var convArg0 = atDateTime.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone13offsetFromUtcERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimezone.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] int standardTimeOffset(const QDateTime &)
func (this *QTimeZone) StandardTimeOffset(atDateTime *QDateTime) int {
	var convArg0 = atDateTime.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone18standardTimeOffsetERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimezone.h:128
// index:0
// Public Visibility=Default Availability=Available
// [4] int daylightTimeOffset(const QDateTime &)
func (this *QTimeZone) DaylightTimeOffset(atDateTime *QDateTime) int {
	var convArg0 = atDateTime.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone18daylightTimeOffsetERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimezone.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasDaylightTime()
func (this *QTimeZone) HasDaylightTime() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone15hasDaylightTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtimezone.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDaylightTime(const QDateTime &)
func (this *QTimeZone) IsDaylightTime(atDateTime *QDateTime) bool {
	var convArg0 = atDateTime.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone14isDaylightTimeERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtimezone.h:133
// index:0
// Public Visibility=Default Availability=Available
// [32] QTimeZone::OffsetData offsetData(const QDateTime &)
func (this *QTimeZone) OffsetData(forDateTime *QDateTime) unsafe.Pointer /*444*/ {
	var convArg0 = forDateTime.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone10offsetDataERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qtimezone.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasTransitions()
func (this *QTimeZone) HasTransitions() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone14hasTransitionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtimezone.h:136
// index:0
// Public Visibility=Default Availability=Available
// [32] QTimeZone::OffsetData nextTransition(const QDateTime &)
func (this *QTimeZone) NextTransition(afterDateTime *QDateTime) unsafe.Pointer /*444*/ {
	var convArg0 = afterDateTime.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone14nextTransitionERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qtimezone.h:137
// index:0
// Public Visibility=Default Availability=Available
// [32] QTimeZone::OffsetData previousTransition(const QDateTime &)
func (this *QTimeZone) PreviousTransition(beforeDateTime *QDateTime) unsafe.Pointer /*444*/ {
	var convArg0 = beforeDateTime.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone18previousTransitionERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qtimezone.h:140
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray systemTimeZoneId()
func (this *QTimeZone) SystemTimeZoneId() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone16systemTimeZoneIdEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QTimeZone_SystemTimeZoneId() *QByteArray /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.SystemTimeZoneId()
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:141
// index:0
// Public static Visibility=Default Availability=Available
// [8] QTimeZone systemTimeZone()
func (this *QTimeZone) SystemTimeZone() *QTimeZone /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone14systemTimeZoneEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTimeZone)
	return rv2
}
func QTimeZone_SystemTimeZone() *QTimeZone /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.SystemTimeZone()
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:142
// index:0
// Public static Visibility=Default Availability=Available
// [8] QTimeZone utc()
func (this *QTimeZone) Utc() *QTimeZone /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone3utcEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTimeZone)
	return rv2
}
func QTimeZone_Utc() *QTimeZone /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.Utc()
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:144
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isTimeZoneIdAvailable(const QByteArray &)
func (this *QTimeZone) IsTimeZoneIdAvailable(ianaId *QByteArray) bool {
	var convArg0 = ianaId.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone21isTimeZoneIdAvailableERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QTimeZone_IsTimeZoneIdAvailable(ianaId *QByteArray) bool {
	var nilthis *QTimeZone
	rv := nilthis.IsTimeZoneIdAvailable(ianaId)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:150
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray ianaIdToWindowsId(const QByteArray &)
func (this *QTimeZone) IanaIdToWindowsId(ianaId *QByteArray) *QByteArray /*123*/ {
	var convArg0 = ianaId.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone17ianaIdToWindowsIdERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QTimeZone_IanaIdToWindowsId(ianaId *QByteArray) *QByteArray /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.IanaIdToWindowsId(ianaId)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:151
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray windowsIdToDefaultIanaId(const QByteArray &)
func (this *QTimeZone) WindowsIdToDefaultIanaId(windowsId *QByteArray) *QByteArray /*123*/ {
	var convArg0 = windowsId.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QTimeZone_WindowsIdToDefaultIanaId(windowsId *QByteArray) *QByteArray /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.WindowsIdToDefaultIanaId(windowsId)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:152
// index:1
// Public static Visibility=Default Availability=Available
// [8] QByteArray windowsIdToDefaultIanaId(const QByteArray &, QLocale::Country)
func (this *QTimeZone) WindowsIdToDefaultIanaId_1(windowsId *QByteArray, country int) *QByteArray /*123*/ {
	var convArg0 = windowsId.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArrayN7QLocale7CountryE", qtrt.FFI_TYPE_POINTER, convArg0, country)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QTimeZone_WindowsIdToDefaultIanaId_1(windowsId *QByteArray, country int) *QByteArray /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.WindowsIdToDefaultIanaId_1(windowsId, country)
	return rv
}

type QTimeZone__ = int

const QTimeZone__MinUtcOffsetSecs QTimeZone__ = -50400
const QTimeZone__MaxUtcOffsetSecs QTimeZone__ = 50400

type QTimeZone__TimeType = int

const QTimeZone__StandardTime QTimeZone__TimeType = 0
const QTimeZone__DaylightTime QTimeZone__TimeType = 1
const QTimeZone__GenericTime QTimeZone__TimeType = 2

type QTimeZone__NameType = int

const QTimeZone__DefaultName QTimeZone__NameType = 0
const QTimeZone__LongName QTimeZone__NameType = 1
const QTimeZone__ShortName QTimeZone__NameType = 2
const QTimeZone__OffsetName QTimeZone__NameType = 3

//  body block end