package qtgui

// /usr/include/qt/QtGui/qpicture.h
// #include <qpicture.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPicture) InheritMetric(f func(m int) int) {
	qtrt.SetAllInheritCallback(this, "metric", f)
}

type QPicture struct {
	*QPaintDevice
}

func (this *QPicture) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QPaintDevice.GetCthis()
	}
}
func (this *QPicture) SetCthis(cthis unsafe.Pointer) {
	this.QPaintDevice = NewQPaintDeviceFromPointer(cthis)
}
func NewQPictureFromPointer(cthis unsafe.Pointer) *QPicture {
	bcthis0 := NewQPaintDeviceFromPointer(cthis)
	return &QPicture{bcthis0}
}
func (*QPicture) NewFromPointer(cthis unsafe.Pointer) *QPicture {
	return NewQPictureFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpicture.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPicture(int)
func NewQPicture(formatVersion int) *QPicture {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPictureC2Ei", qtrt.FFI_TYPE_POINTER, formatVersion)
	gopp.ErrPrint(err, rv)
	gothis := NewQPictureFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPicture)
	return gothis
}

// /usr/include/qt/QtGui/qpicture.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPicture()
func DeleteQPicture(this *QPicture) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPictureD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpicture.h:63
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QPicture) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int devType()
func (this *QPicture) DevType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture7devTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpicture.h:66
// index:0
// Public Visibility=Default Availability=Available
// [4] uint size()
func (this *QPicture) Size() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qpicture.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * data()
func (this *QPicture) Data() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtGui/qpicture.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setData(const char *, uint)
func (this *QPicture) SetData(data string, size uint) {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture7setDataEPKcj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool play(QPainter *)
func (this *QPicture) Play(p *QPainter /*777 QPainter **/) bool {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4playEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(QIODevice *, const char *)
func (this *QPicture) Load(dev *qtcore.QIODevice /*777 QIODevice **/, format string) bool {
	var convArg0 = dev.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4loadEP9QIODevicePKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:73
// index:1
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const char *)
func (this *QPicture) Load_1(fileName string, format string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4loadERK7QStringPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *)
func (this *QPicture) Save(dev *qtcore.QIODevice /*777 QIODevice **/, format string) bool {
	var convArg0 = dev.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4saveEP9QIODevicePKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:75
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(const QString &, const char *)
func (this *QPicture) Save_1(fileName string, format string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4saveERK7QStringPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:77
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect boundingRect()
func (this *QPicture) BoundingRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpicture.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBoundingRect(const QRect &)
func (this *QPicture) SetBoundingRect(r *qtcore.QRect) {
	var convArg0 = r.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture15setBoundingRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:85
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPicture &)
func (this *QPicture) Swap(other *QPicture) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detach()
func (this *QPicture) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached()
func (this *QPicture) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [8] const char * pictureFormat(const QString &)
func (this *QPicture) PictureFormat(fileName string) string {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture13pictureFormatERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}
func QPicture_PictureFormat(fileName string) string {
	var nilthis *QPicture
	rv := nilthis.PictureFormat(fileName)
	return rv
}

// /usr/include/qt/QtGui/qpicture.h:97
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList inputFormatList()
func (this *QPicture) InputFormatList() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture15inputFormatListEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QPicture_InputFormatList() *qtcore.QStringList /*123*/ {
	var nilthis *QPicture
	rv := nilthis.InputFormatList()
	return rv
}

// /usr/include/qt/QtGui/qpicture.h:98
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList outputFormatList()
func (this *QPicture) OutputFormatList() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture16outputFormatListEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QPicture_OutputFormatList() *qtcore.QStringList /*123*/ {
	var nilthis *QPicture
	rv := nilthis.OutputFormatList()
	return rv
}

// /usr/include/qt/QtGui/qpicture.h:101
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine()
func (this *QPicture) PaintEngine() *QPaintEngine /*777 QPaintEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture11paintEngineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpicture.h:106
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPicture) Metric(m int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture6metricEN12QPaintDevice17PaintDeviceMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), m)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end