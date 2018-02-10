package qtandroidextras

import "unsafe"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

func init() {
	if false {
		_ = unsafe.Pointer(uintptr(0))
	}
	if false {
		qtrt.KeepMe()
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

//  header block end

//  body block begin
// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidfunctions.h:107
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool shouldShowRequestPermissionRationale(const QString &)
func ShouldShowRequestPermissionRationale(permission string) bool {
	var tmpArg0 = qtcore.NewQString_5(permission)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QtAndroid36shouldShowRequestPermissionRationaleERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidfunctions.h:104
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void requestPermissions(const QStringList &, const QtAndroid::PermissionResultCallback &)
func RequestPermissions(permissions *qtcore.QStringList, callbackFunc unsafe.Pointer /*555*/) {
	var convArg0 = permissions.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QtAndroid18requestPermissionsERK11QStringListRKSt8functionIFvRK5QHashI7QStringNS_16PermissionResultEEEE", qtrt.FFI_TYPE_POINTER, convArg0, callbackFunc)
	gopp.ErrPrint(err, rv)
}

//  body block end
