package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qtextcodec.h
// dst-file: /src/core/qtextcodec.go
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
  // proto:  void QTextEncoder::~QTextEncoder();
extern void _ZN12QTextEncoderD0Ev(void* qthis);
  // proto:  QByteArray QTextEncoder::fromUnicode(const QString & str);
extern void _ZN12QTextEncoder11fromUnicodeERK7QString(void* qthis, void* arg0);
  // proto:  bool QTextEncoder::hasFailure();
extern void _ZNK12QTextEncoder10hasFailureEv(void* qthis);
  // proto:  void QTextEncoder::QTextEncoder(const QTextCodec * codec);
extern void* dector_ZN12QTextEncoderC1EPK10QTextCodec(void* arg0);
extern void _ZN12QTextEncoderC1EPK10QTextCodec(void* qthis, void* arg0);
  // proto:  void QTextEncoder::QTextEncoder(const QTextEncoder & );
extern void* dector_ZN12QTextEncoderC1ERKS_(void* arg0);
extern void _ZN12QTextEncoderC1ERKS_(void* qthis, void* arg0);
  // proto:  QByteArray QTextEncoder::fromUnicode(const QChar * uc, int len);
extern void _ZN12QTextEncoder11fromUnicodeEPK5QChari(void* qthis, void* arg0, int arg1);
  // proto:  QByteArray QTextCodec::name();
extern void _ZNK10QTextCodec4nameEv(void* qthis);
  // proto:  QString QTextCodec::toUnicode(const QByteArray & );
extern void _ZNK10QTextCodec9toUnicodeERK10QByteArray(void* qthis, void* arg0);
  // proto:  QByteArray QTextCodec::fromUnicode(const QString & uc);
extern void _ZNK10QTextCodec11fromUnicodeERK7QString(void* qthis, void* arg0);
  // proto: static QTextCodec * QTextCodec::codecForLocale();
extern void _ZN10QTextCodec14codecForLocaleEv();
  // proto: static QList<int> QTextCodec::availableMibs();
extern void _ZN10QTextCodec13availableMibsEv();
  // proto: static QTextCodec * QTextCodec::codecForHtml(const QByteArray & ba);
extern void _ZN10QTextCodec12codecForHtmlERK10QByteArray(void* arg0);
  // proto:  void QTextCodec::QTextCodec(const QTextCodec & );
extern void* dector_ZN10QTextCodecC1ERKS_(void* arg0);
extern void _ZN10QTextCodecC1ERKS_(void* qthis, void* arg0);
  // proto: static void QTextCodec::setCodecForLocale(QTextCodec * c);
extern void _ZN10QTextCodec17setCodecForLocaleEPS_(void* arg0);
  // proto: static QTextCodec * QTextCodec::codecForUtfText(const QByteArray & ba);
extern void _ZN10QTextCodec15codecForUtfTextERK10QByteArray(void* arg0);
  // proto:  QString QTextCodec::toUnicode(const char * chars);
extern void _ZNK10QTextCodec9toUnicodeEPKc(void* qthis, char* arg0);
  // proto:  int QTextCodec::mibEnum();
extern void _ZNK10QTextCodec7mibEnumEv(void* qthis);
  // proto: static QTextCodec * QTextCodec::codecForName(const char * name);
extern void _ZN10QTextCodec12codecForNameEPKc(char* arg0);
  // proto:  bool QTextCodec::canEncode(const QString & );
extern void _ZNK10QTextCodec9canEncodeERK7QString(void* qthis, void* arg0);
  // proto:  QList<QByteArray> QTextCodec::aliases();
extern void _ZNK10QTextCodec7aliasesEv(void* qthis);
  // proto: static QTextCodec * QTextCodec::codecForName(const QByteArray & name);
extern void _ZN10QTextCodec12codecForNameERK10QByteArray(void* arg0);
  // proto: static QList<QByteArray> QTextCodec::availableCodecs();
extern void _ZN10QTextCodec15availableCodecsEv();
  // proto: static QTextCodec * QTextCodec::codecForHtml(const QByteArray & ba, QTextCodec * defaultCodec);
extern void _ZN10QTextCodec12codecForHtmlERK10QByteArrayPS_(void* arg0, void* arg1);
  // proto:  void QTextCodec::~QTextCodec();
extern void _ZN10QTextCodecD0Ev(void* qthis);
  // proto: static QTextCodec * QTextCodec::codecForMib(int mib);
extern void _ZN10QTextCodec11codecForMibEi(int arg0);
  // proto:  void QTextCodec::QTextCodec();
extern void* dector_ZN10QTextCodecC1Ev();
extern void _ZN10QTextCodecC1Ev(void* qthis);
  // proto: static QTextCodec * QTextCodec::codecForUtfText(const QByteArray & ba, QTextCodec * defaultCodec);
extern void _ZN10QTextCodec15codecForUtfTextERK10QByteArrayPS_(void* arg0, void* arg1);
  // proto:  bool QTextCodec::canEncode(QChar );
extern void _ZNK10QTextCodec9canEncodeE5QChar(void* qthis, void* arg0);
  // proto:  QString QTextDecoder::toUnicode(const char * chars, int len);
extern void _ZN12QTextDecoder9toUnicodeEPKci(void* qthis, char* arg0, int arg1);
  // proto:  void QTextDecoder::QTextDecoder(const QTextCodec * codec);
extern void* dector_ZN12QTextDecoderC1EPK10QTextCodec(void* arg0);
extern void _ZN12QTextDecoderC1EPK10QTextCodec(void* qthis, void* arg0);
  // proto:  bool QTextDecoder::hasFailure();
extern void _ZNK12QTextDecoder10hasFailureEv(void* qthis);
  // proto:  void QTextDecoder::~QTextDecoder();
extern void _ZN12QTextDecoderD0Ev(void* qthis);
  // proto:  QString QTextDecoder::toUnicode(const QByteArray & ba);
extern void _ZN12QTextDecoder9toUnicodeERK10QByteArray(void* qthis, void* arg0);
  // proto:  void QTextDecoder::QTextDecoder(const QTextDecoder & );
extern void* dector_ZN12QTextDecoderC1ERKS_(void* arg0);
extern void _ZN12QTextDecoderC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTextDecoder::toUnicode(QString * target, const char * chars, int len);
extern void _ZN12QTextDecoder9toUnicodeEP7QStringPKci(void* qthis, void* arg0, char* arg1, int arg2);
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

// class sizeof(QTextEncoder)=1
type QTextEncoder struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextCodec)=8
type QTextCodec struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextDecoder)=1
type QTextDecoder struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QTextEncoder::~QTextEncoder();
func (this *QTextEncoder) FreeQTextEncoder(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextEncoder", "~QTextEncoder", args)
  }

}

  // proto:  QByteArray QTextEncoder::fromUnicode(const QString & str);
func (this *QTextEncoder) fromUnicode(args ...interface{}) () {
  // fromUnicode(const class QString &)
  // fromUnicode(const class QChar *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextEncoder11fromUnicodeERK7QString
  case 1:
    // invoke: _ZN12QTextEncoder11fromUnicodeEPK5QChari
  default:
    qtrt.ErrorResolve("QTextEncoder", "fromUnicode", args)
  }

}

  // proto:  bool QTextEncoder::hasFailure();
func (this *QTextEncoder) hasFailure(args ...interface{}) () {
  // hasFailure()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTextEncoder10hasFailureEv
  default:
    qtrt.ErrorResolve("QTextEncoder", "hasFailure", args)
  }

}

  // proto:  void QTextEncoder::QTextEncoder(const QTextCodec * codec);
func NewQTextEncoder(args ...interface{}) QTextEncoder {
  return QTextEncoder{}
}

  // proto:  QByteArray QTextCodec::name();
func (this *QTextCodec) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextCodec4nameEv
  default:
    qtrt.ErrorResolve("QTextCodec", "name", args)
  }

}

  // proto:  QString QTextCodec::toUnicode(const QByteArray & );
func (this *QTextCodec) toUnicode(args ...interface{}) () {
  // toUnicode(const char *, int, struct QTextCodec::ConverterState *)
  // toUnicode(const class QByteArray &)
  // toUnicode(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  // vtys[0][2] = reflect.TypeOf(QTextCodec::ConverterState{}) // "QTextCodec::ConverterState *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextCodec9toUnicodeEPKciPNS_14ConverterStateE
  case 1:
    // invoke: _ZNK10QTextCodec9toUnicodeERK10QByteArray
  case 2:
    // invoke: _ZNK10QTextCodec9toUnicodeEPKc
  default:
    qtrt.ErrorResolve("QTextCodec", "toUnicode", args)
  }

}

  // proto:  QByteArray QTextCodec::fromUnicode(const QString & uc);
func (this *QTextCodec) fromUnicode(args ...interface{}) () {
  // fromUnicode(const class QString &)
  // fromUnicode(const class QChar *, int, struct QTextCodec::ConverterState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  // vtys[1][2] = reflect.TypeOf(QTextCodec::ConverterState{}) // "QTextCodec::ConverterState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextCodec11fromUnicodeERK7QString
  case 1:
    // invoke: _ZNK10QTextCodec11fromUnicodeEPK5QChariPNS_14ConverterStateE
  default:
    qtrt.ErrorResolve("QTextCodec", "fromUnicode", args)
  }

}

  // proto: static QTextCodec * QTextCodec::codecForLocale();
func (this *QTextCodec) codecForLocale_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForLocale", args)
  }

}

  // proto: static QList<int> QTextCodec::availableMibs();
func (this *QTextCodec) availableMibs_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "availableMibs", args)
  }

}

  // proto: static QTextCodec * QTextCodec::codecForHtml(const QByteArray & ba);
func (this *QTextCodec) codecForHtml_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForHtml", args)
  }

}

  // proto:  void QTextCodec::QTextCodec(const QTextCodec & );
func NewQTextCodec(args ...interface{}) QTextCodec {
  return QTextCodec{}
}

  // proto: static void QTextCodec::setCodecForLocale(QTextCodec * c);
func (this *QTextCodec) setCodecForLocale_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "setCodecForLocale", args)
  }

}

  // proto: static QTextCodec * QTextCodec::codecForUtfText(const QByteArray & ba);
func (this *QTextCodec) codecForUtfText_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForUtfText", args)
  }

}

  // proto:  int QTextCodec::mibEnum();
func (this *QTextCodec) mibEnum(args ...interface{}) () {
  // mibEnum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextCodec7mibEnumEv
  default:
    qtrt.ErrorResolve("QTextCodec", "mibEnum", args)
  }

}

  // proto: static QTextCodec * QTextCodec::codecForName(const char * name);
func (this *QTextCodec) codecForName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForName", args)
  }

}

  // proto:  bool QTextCodec::canEncode(const QString & );
func (this *QTextCodec) canEncode(args ...interface{}) () {
  // canEncode(const class QString &)
  // canEncode(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextCodec9canEncodeERK7QString
  case 1:
    // invoke: _ZNK10QTextCodec9canEncodeE5QChar
  default:
    qtrt.ErrorResolve("QTextCodec", "canEncode", args)
  }

}

  // proto:  QList<QByteArray> QTextCodec::aliases();
func (this *QTextCodec) aliases(args ...interface{}) () {
  // aliases()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextCodec7aliasesEv
  default:
    qtrt.ErrorResolve("QTextCodec", "aliases", args)
  }

}

  // proto: static QList<QByteArray> QTextCodec::availableCodecs();
func (this *QTextCodec) availableCodecs_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "availableCodecs", args)
  }

}

  // proto:  void QTextCodec::~QTextCodec();
func (this *QTextCodec) FreeQTextCodec(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "~QTextCodec", args)
  }

}

  // proto: static QTextCodec * QTextCodec::codecForMib(int mib);
func (this *QTextCodec) codecForMib_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForMib", args)
  }

}

  // proto:  QString QTextDecoder::toUnicode(const char * chars, int len);
func (this *QTextDecoder) toUnicode(args ...interface{}) () {
  // toUnicode(const char *, int)
  // toUnicode(const class QByteArray &)
  // toUnicode(class QString *, const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "QString *"
  vtys[2][1] = qtrt.ByteTy(true) // "const char *"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextDecoder9toUnicodeEPKci
  case 1:
    // invoke: _ZN12QTextDecoder9toUnicodeERK10QByteArray
  case 2:
    // invoke: _ZN12QTextDecoder9toUnicodeEP7QStringPKci
  default:
    qtrt.ErrorResolve("QTextDecoder", "toUnicode", args)
  }

}

  // proto:  void QTextDecoder::QTextDecoder(const QTextCodec * codec);
func NewQTextDecoder(args ...interface{}) QTextDecoder {
  return QTextDecoder{}
}

  // proto:  bool QTextDecoder::hasFailure();
func (this *QTextDecoder) hasFailure(args ...interface{}) () {
  // hasFailure()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTextDecoder10hasFailureEv
  default:
    qtrt.ErrorResolve("QTextDecoder", "hasFailure", args)
  }

}

  // proto:  void QTextDecoder::~QTextDecoder();
func (this *QTextDecoder) FreeQTextDecoder(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextDecoder", "~QTextDecoder", args)
  }

}

// <= body block end
