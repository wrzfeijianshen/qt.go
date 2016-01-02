package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qxmlstream.h
// dst-file: /src/core/qxmlstream.go
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
  // proto:  int QXmlStreamStringRef::size();
extern void demth_ZNK19QXmlStreamStringRef4sizeEv(void* qthis);
  // proto:  void QXmlStreamStringRef::clear();
extern void demth_ZN19QXmlStreamStringRef5clearEv(void* qthis);
  // proto:  const QString * QXmlStreamStringRef::string();
extern void demth_ZNK19QXmlStreamStringRef6stringEv(void* qthis);
  // proto:  void QXmlStreamStringRef::QXmlStreamStringRef(const QString & aString);
extern void* dector_ZN19QXmlStreamStringRefC1ERK7QString(void* arg0);
extern void demth_ZN19QXmlStreamStringRefC1ERK7QString(void* qthis, void* arg0);
  // proto:  void QXmlStreamStringRef::~QXmlStreamStringRef();
extern void demth_ZN19QXmlStreamStringRefD0Ev(void* qthis);
  // proto:  void QXmlStreamStringRef::QXmlStreamStringRef();
extern void* dector_ZN19QXmlStreamStringRefC1Ev();
extern void demth_ZN19QXmlStreamStringRefC1Ev(void* qthis);
  // proto:  int QXmlStreamStringRef::position();
extern void demth_ZNK19QXmlStreamStringRef8positionEv(void* qthis);
  // proto:  QStringRef QXmlStreamReader::name();
extern void _ZNK16QXmlStreamReader4nameEv(void* qthis);
  // proto:  QXmlStreamEntityResolver * QXmlStreamReader::entityResolver();
extern void _ZNK16QXmlStreamReader14entityResolverEv(void* qthis);
  // proto:  bool QXmlStreamReader::namespaceProcessing();
extern void _ZNK16QXmlStreamReader19namespaceProcessingEv(void* qthis);
  // proto:  bool QXmlStreamReader::isStartElement();
extern void demth_ZNK16QXmlStreamReader14isStartElementEv(void* qthis);
  // proto:  bool QXmlStreamReader::isStandaloneDocument();
extern void _ZNK16QXmlStreamReader20isStandaloneDocumentEv(void* qthis);
  // proto:  qint64 QXmlStreamReader::lineNumber();
extern void _ZNK16QXmlStreamReader10lineNumberEv(void* qthis);
  // proto:  void QXmlStreamReader::clear();
extern void _ZN16QXmlStreamReader5clearEv(void* qthis);
  // proto:  QStringRef QXmlStreamReader::processingInstructionData();
extern void _ZNK16QXmlStreamReader25processingInstructionDataEv(void* qthis);
  // proto:  void QXmlStreamReader::addData(const QString & data);
extern void _ZN16QXmlStreamReader7addDataERK7QString(void* qthis, void* arg0);
  // proto:  QStringRef QXmlStreamReader::dtdPublicId();
extern void _ZNK16QXmlStreamReader11dtdPublicIdEv(void* qthis);
  // proto:  QStringRef QXmlStreamReader::documentEncoding();
extern void _ZNK16QXmlStreamReader16documentEncodingEv(void* qthis);
  // proto:  qint64 QXmlStreamReader::characterOffset();
extern void _ZNK16QXmlStreamReader15characterOffsetEv(void* qthis);
  // proto:  QXmlStreamAttributes QXmlStreamReader::attributes();
extern void _ZNK16QXmlStreamReader10attributesEv(void* qthis);
  // proto:  QString QXmlStreamReader::tokenString();
extern void _ZNK16QXmlStreamReader11tokenStringEv(void* qthis);
  // proto:  void QXmlStreamReader::addExtraNamespaceDeclaration(const QXmlStreamNamespaceDeclaration & extraNamespaceDeclaraction);
extern void _ZN16QXmlStreamReader28addExtraNamespaceDeclarationERK30QXmlStreamNamespaceDeclaration(void* qthis, void* arg0);
  // proto:  void QXmlStreamReader::QXmlStreamReader(const QByteArray & data);
extern void* dector_ZN16QXmlStreamReaderC1ERK10QByteArray(void* arg0);
extern void _ZN16QXmlStreamReaderC1ERK10QByteArray(void* qthis, void* arg0);
  // proto:  QStringRef QXmlStreamReader::qualifiedName();
extern void _ZNK16QXmlStreamReader13qualifiedNameEv(void* qthis);
  // proto:  QIODevice * QXmlStreamReader::device();
extern void _ZNK16QXmlStreamReader6deviceEv(void* qthis);
  // proto:  QStringRef QXmlStreamReader::namespaceUri();
extern void _ZNK16QXmlStreamReader12namespaceUriEv(void* qthis);
  // proto:  QStringRef QXmlStreamReader::text();
extern void _ZNK16QXmlStreamReader4textEv(void* qthis);
  // proto:  void QXmlStreamReader::setDevice(QIODevice * device);
extern void _ZN16QXmlStreamReader9setDeviceEP9QIODevice(void* qthis, void* arg0);
  // proto:  void QXmlStreamReader::QXmlStreamReader(QIODevice * device);
extern void* dector_ZN16QXmlStreamReaderC1EP9QIODevice(void* arg0);
extern void _ZN16QXmlStreamReaderC1EP9QIODevice(void* qthis, void* arg0);
  // proto:  QStringRef QXmlStreamReader::documentVersion();
extern void _ZNK16QXmlStreamReader15documentVersionEv(void* qthis);
  // proto:  bool QXmlStreamReader::isDTD();
extern void demth_ZNK16QXmlStreamReader5isDTDEv(void* qthis);
  // proto:  bool QXmlStreamReader::isStartDocument();
extern void demth_ZNK16QXmlStreamReader15isStartDocumentEv(void* qthis);
  // proto:  QString QXmlStreamReader::errorString();
extern void _ZNK16QXmlStreamReader11errorStringEv(void* qthis);
  // proto:  bool QXmlStreamReader::isProcessingInstruction();
extern void demth_ZNK16QXmlStreamReader23isProcessingInstructionEv(void* qthis);
  // proto:  void QXmlStreamReader::setEntityResolver(QXmlStreamEntityResolver * resolver);
extern void _ZN16QXmlStreamReader17setEntityResolverEP24QXmlStreamEntityResolver(void* qthis, void* arg0);
  // proto:  bool QXmlStreamReader::isCharacters();
extern void demth_ZNK16QXmlStreamReader12isCharactersEv(void* qthis);
  // proto:  void QXmlStreamReader::QXmlStreamReader();
extern void* dector_ZN16QXmlStreamReaderC1Ev();
extern void _ZN16QXmlStreamReaderC1Ev(void* qthis);
  // proto:  void QXmlStreamReader::QXmlStreamReader(const QString & data);
extern void* dector_ZN16QXmlStreamReaderC1ERK7QString(void* arg0);
extern void _ZN16QXmlStreamReaderC1ERK7QString(void* qthis, void* arg0);
  // proto:  QXmlStreamEntityDeclarations QXmlStreamReader::entityDeclarations();
extern void _ZNK16QXmlStreamReader18entityDeclarationsEv(void* qthis);
  // proto:  bool QXmlStreamReader::isWhitespace();
extern void _ZNK16QXmlStreamReader12isWhitespaceEv(void* qthis);
  // proto:  void QXmlStreamReader::QXmlStreamReader(const QXmlStreamReader & );
extern void* dector_ZN16QXmlStreamReaderC1ERKS_(void* arg0);
extern void _ZN16QXmlStreamReaderC1ERKS_(void* qthis, void* arg0);
  // proto:  qint64 QXmlStreamReader::columnNumber();
extern void _ZNK16QXmlStreamReader12columnNumberEv(void* qthis);
  // proto:  bool QXmlStreamReader::hasError();
extern void demth_ZNK16QXmlStreamReader8hasErrorEv(void* qthis);
  // proto:  bool QXmlStreamReader::isCDATA();
extern void _ZNK16QXmlStreamReader7isCDATAEv(void* qthis);
  // proto:  void QXmlStreamReader::~QXmlStreamReader();
extern void _ZN16QXmlStreamReaderD0Ev(void* qthis);
  // proto:  QStringRef QXmlStreamReader::processingInstructionTarget();
extern void _ZNK16QXmlStreamReader27processingInstructionTargetEv(void* qthis);
  // proto:  void QXmlStreamReader::addData(const char * data);
extern void _ZN16QXmlStreamReader7addDataEPKc(void* qthis, char* arg0);
  // proto:  QStringRef QXmlStreamReader::dtdSystemId();
extern void _ZNK16QXmlStreamReader11dtdSystemIdEv(void* qthis);
  // proto:  QStringRef QXmlStreamReader::prefix();
extern void _ZNK16QXmlStreamReader6prefixEv(void* qthis);
  // proto:  bool QXmlStreamReader::isEndElement();
extern void demth_ZNK16QXmlStreamReader12isEndElementEv(void* qthis);
  // proto:  QXmlStreamNotationDeclarations QXmlStreamReader::notationDeclarations();
extern void _ZNK16QXmlStreamReader20notationDeclarationsEv(void* qthis);
  // proto:  QXmlStreamNamespaceDeclarations QXmlStreamReader::namespaceDeclarations();
extern void _ZNK16QXmlStreamReader21namespaceDeclarationsEv(void* qthis);
  // proto:  void QXmlStreamReader::setNamespaceProcessing(bool );
extern void _ZN16QXmlStreamReader22setNamespaceProcessingEb(void* qthis, bool arg0);
  // proto:  void QXmlStreamReader::raiseError(const QString & message);
extern void _ZN16QXmlStreamReader10raiseErrorERK7QString(void* qthis, void* arg0);
  // proto:  QStringRef QXmlStreamReader::dtdName();
extern void _ZNK16QXmlStreamReader7dtdNameEv(void* qthis);
  // proto:  bool QXmlStreamReader::isEndDocument();
extern void demth_ZNK16QXmlStreamReader13isEndDocumentEv(void* qthis);
  // proto:  bool QXmlStreamReader::readNextStartElement();
extern void _ZN16QXmlStreamReader20readNextStartElementEv(void* qthis);
  // proto:  bool QXmlStreamReader::isComment();
extern void demth_ZNK16QXmlStreamReader9isCommentEv(void* qthis);
  // proto:  void QXmlStreamReader::QXmlStreamReader(const char * data);
extern void* dector_ZN16QXmlStreamReaderC1EPKc(char* arg0);
extern void _ZN16QXmlStreamReaderC1EPKc(void* qthis, char* arg0);
  // proto:  void QXmlStreamReader::skipCurrentElement();
extern void _ZN16QXmlStreamReader18skipCurrentElementEv(void* qthis);
  // proto:  bool QXmlStreamReader::isEntityReference();
extern void demth_ZNK16QXmlStreamReader17isEntityReferenceEv(void* qthis);
  // proto:  void QXmlStreamReader::addData(const QByteArray & data);
extern void _ZN16QXmlStreamReader7addDataERK10QByteArray(void* qthis, void* arg0);
  // proto:  bool QXmlStreamReader::atEnd();
extern void _ZNK16QXmlStreamReader5atEndEv(void* qthis);
  // proto:  QString QXmlStreamEntityResolver::resolveEntity(const QString & publicId, const QString & systemId);
extern void _ZN24QXmlStreamEntityResolver13resolveEntityERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  QString QXmlStreamEntityResolver::resolveUndeclaredEntity(const QString & name);
extern void _ZN24QXmlStreamEntityResolver23resolveUndeclaredEntityERK7QString(void* qthis, void* arg0);
  // proto:  void QXmlStreamEntityResolver::~QXmlStreamEntityResolver();
extern void _ZN24QXmlStreamEntityResolverD0Ev(void* qthis);
  // proto:  QStringRef QXmlStreamNamespaceDeclaration::namespaceUri();
extern void demth_ZNK30QXmlStreamNamespaceDeclaration12namespaceUriEv(void* qthis);
  // proto:  void QXmlStreamNamespaceDeclaration::QXmlStreamNamespaceDeclaration(const QXmlStreamNamespaceDeclaration & );
extern void* dector_ZN30QXmlStreamNamespaceDeclarationC1ERKS_(void* arg0);
extern void _ZN30QXmlStreamNamespaceDeclarationC1ERKS_(void* qthis, void* arg0);
  // proto:  void QXmlStreamNamespaceDeclaration::QXmlStreamNamespaceDeclaration();
extern void* dector_ZN30QXmlStreamNamespaceDeclarationC1Ev();
extern void _ZN30QXmlStreamNamespaceDeclarationC1Ev(void* qthis);
  // proto:  QStringRef QXmlStreamNamespaceDeclaration::prefix();
extern void demth_ZNK30QXmlStreamNamespaceDeclaration6prefixEv(void* qthis);
  // proto:  void QXmlStreamNamespaceDeclaration::QXmlStreamNamespaceDeclaration(const QString & prefix, const QString & namespaceUri);
extern void* dector_ZN30QXmlStreamNamespaceDeclarationC1ERK7QStringS2_(void* arg0, void* arg1);
extern void _ZN30QXmlStreamNamespaceDeclarationC1ERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QXmlStreamNamespaceDeclaration::~QXmlStreamNamespaceDeclaration();
extern void _ZN30QXmlStreamNamespaceDeclarationD0Ev(void* qthis);
  // proto:  void QXmlStreamEntityDeclaration::~QXmlStreamEntityDeclaration();
extern void _ZN27QXmlStreamEntityDeclarationD0Ev(void* qthis);
  // proto:  QStringRef QXmlStreamEntityDeclaration::publicId();
extern void demth_ZNK27QXmlStreamEntityDeclaration8publicIdEv(void* qthis);
  // proto:  QStringRef QXmlStreamEntityDeclaration::name();
extern void demth_ZNK27QXmlStreamEntityDeclaration4nameEv(void* qthis);
  // proto:  void QXmlStreamEntityDeclaration::QXmlStreamEntityDeclaration();
extern void* dector_ZN27QXmlStreamEntityDeclarationC1Ev();
extern void _ZN27QXmlStreamEntityDeclarationC1Ev(void* qthis);
  // proto:  QStringRef QXmlStreamEntityDeclaration::value();
extern void demth_ZNK27QXmlStreamEntityDeclaration5valueEv(void* qthis);
  // proto:  QStringRef QXmlStreamEntityDeclaration::notationName();
extern void demth_ZNK27QXmlStreamEntityDeclaration12notationNameEv(void* qthis);
  // proto:  void QXmlStreamEntityDeclaration::QXmlStreamEntityDeclaration(const QXmlStreamEntityDeclaration & );
extern void* dector_ZN27QXmlStreamEntityDeclarationC1ERKS_(void* arg0);
extern void _ZN27QXmlStreamEntityDeclarationC1ERKS_(void* qthis, void* arg0);
  // proto:  QStringRef QXmlStreamEntityDeclaration::systemId();
extern void demth_ZNK27QXmlStreamEntityDeclaration8systemIdEv(void* qthis);
  // proto:  QStringRef QXmlStreamAttributes::value(const QString & qualifiedName);
extern void _ZNK20QXmlStreamAttributes5valueERK7QString(void* qthis, void* arg0);
  // proto:  void QXmlStreamAttributes::QXmlStreamAttributes();
extern void* dector_ZN20QXmlStreamAttributesC1Ev();
extern void demth_ZN20QXmlStreamAttributesC1Ev(void* qthis);
  // proto:  bool QXmlStreamAttributes::hasAttribute(const QString & qualifiedName);
extern void demth_ZNK20QXmlStreamAttributes12hasAttributeERK7QString(void* qthis, void* arg0);
  // proto:  bool QXmlStreamAttributes::hasAttribute(const QString & namespaceUri, const QString & name);
extern void demth_ZNK20QXmlStreamAttributes12hasAttributeERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QXmlStreamAttributes::append(const QString & namespaceUri, const QString & name, const QString & value);
extern void _ZN20QXmlStreamAttributes6appendERK7QStringS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QXmlStreamAttributes::append(const QString & qualifiedName, const QString & value);
extern void _ZN20QXmlStreamAttributes6appendERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  QStringRef QXmlStreamAttributes::value(const QString & namespaceUri, const QString & name);
extern void _ZNK20QXmlStreamAttributes5valueERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QXmlStreamWriter::writeEndElement();
extern void _ZN16QXmlStreamWriter15writeEndElementEv(void* qthis);
  // proto:  void QXmlStreamWriter::QXmlStreamWriter();
extern void* dector_ZN16QXmlStreamWriterC1Ev();
extern void _ZN16QXmlStreamWriterC1Ev(void* qthis);
  // proto:  void QXmlStreamWriter::writeEndDocument();
extern void _ZN16QXmlStreamWriter16writeEndDocumentEv(void* qthis);
  // proto:  bool QXmlStreamWriter::autoFormatting();
extern void _ZNK16QXmlStreamWriter14autoFormattingEv(void* qthis);
  // proto:  void QXmlStreamWriter::writeStartDocument(const QString & version, bool standalone);
extern void _ZN16QXmlStreamWriter18writeStartDocumentERK7QStringb(void* qthis, void* arg0, bool arg1);
  // proto:  void QXmlStreamWriter::setCodec(QTextCodec * codec);
extern void _ZN16QXmlStreamWriter8setCodecEP10QTextCodec(void* qthis, void* arg0);
  // proto:  void QXmlStreamWriter::writeProcessingInstruction(const QString & target, const QString & data);
extern void _ZN16QXmlStreamWriter26writeProcessingInstructionERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QXmlStreamWriter::writeCharacters(const QString & text);
extern void _ZN16QXmlStreamWriter15writeCharactersERK7QString(void* qthis, void* arg0);
  // proto:  void QXmlStreamWriter::setDevice(QIODevice * device);
extern void _ZN16QXmlStreamWriter9setDeviceEP9QIODevice(void* qthis, void* arg0);
  // proto:  void QXmlStreamWriter::writeStartDocument(const QString & version);
extern void _ZN16QXmlStreamWriter18writeStartDocumentERK7QString(void* qthis, void* arg0);
  // proto:  void QXmlStreamWriter::writeTextElement(const QString & namespaceUri, const QString & name, const QString & text);
extern void _ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QXmlStreamWriter::writeAttribute(const QString & qualifiedName, const QString & value);
extern void _ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QXmlStreamWriter::writeEmptyElement(const QString & qualifiedName);
extern void _ZN16QXmlStreamWriter17writeEmptyElementERK7QString(void* qthis, void* arg0);
  // proto:  void QXmlStreamWriter::writeDTD(const QString & dtd);
extern void _ZN16QXmlStreamWriter8writeDTDERK7QString(void* qthis, void* arg0);
  // proto:  void QXmlStreamWriter::setAutoFormattingIndent(int spacesOrTabs);
extern void _ZN16QXmlStreamWriter23setAutoFormattingIndentEi(void* qthis, int arg0);
  // proto:  void QXmlStreamWriter::writeAttribute(const QXmlStreamAttribute & attribute);
extern void _ZN16QXmlStreamWriter14writeAttributeERK19QXmlStreamAttribute(void* qthis, void* arg0);
  // proto:  void QXmlStreamWriter::writeStartElement(const QString & namespaceUri, const QString & name);
extern void _ZN16QXmlStreamWriter17writeStartElementERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QXmlStreamWriter::QXmlStreamWriter(QString * string);
extern void* dector_ZN16QXmlStreamWriterC1EP7QString(void* arg0);
extern void _ZN16QXmlStreamWriterC1EP7QString(void* qthis, void* arg0);
  // proto:  void QXmlStreamWriter::writeComment(const QString & text);
extern void _ZN16QXmlStreamWriter12writeCommentERK7QString(void* qthis, void* arg0);
  // proto:  QTextCodec * QXmlStreamWriter::codec();
extern void _ZNK16QXmlStreamWriter5codecEv(void* qthis);
  // proto:  void QXmlStreamWriter::writeAttribute(const QString & namespaceUri, const QString & name, const QString & value);
extern void _ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QXmlStreamWriter::writeNamespace(const QString & namespaceUri, const QString & prefix);
extern void _ZN16QXmlStreamWriter14writeNamespaceERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QXmlStreamWriter::QXmlStreamWriter(const QXmlStreamWriter & );
extern void* dector_ZN16QXmlStreamWriterC1ERKS_(void* arg0);
extern void _ZN16QXmlStreamWriterC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QXmlStreamWriter::hasError();
extern void _ZNK16QXmlStreamWriter8hasErrorEv(void* qthis);
  // proto:  void QXmlStreamWriter::writeCDATA(const QString & text);
extern void _ZN16QXmlStreamWriter10writeCDATAERK7QString(void* qthis, void* arg0);
  // proto:  void QXmlStreamWriter::writeStartDocument();
extern void _ZN16QXmlStreamWriter18writeStartDocumentEv(void* qthis);
  // proto:  void QXmlStreamWriter::writeEntityReference(const QString & name);
extern void _ZN16QXmlStreamWriter20writeEntityReferenceERK7QString(void* qthis, void* arg0);
  // proto:  void QXmlStreamWriter::setAutoFormatting(bool );
extern void _ZN16QXmlStreamWriter17setAutoFormattingEb(void* qthis, bool arg0);
  // proto:  void QXmlStreamWriter::setCodec(const char * codecName);
extern void _ZN16QXmlStreamWriter8setCodecEPKc(void* qthis, char* arg0);
  // proto:  int QXmlStreamWriter::autoFormattingIndent();
extern void _ZNK16QXmlStreamWriter20autoFormattingIndentEv(void* qthis);
  // proto:  void QXmlStreamWriter::~QXmlStreamWriter();
extern void _ZN16QXmlStreamWriterD0Ev(void* qthis);
  // proto:  void QXmlStreamWriter::writeAttributes(const QXmlStreamAttributes & attributes);
extern void _ZN16QXmlStreamWriter15writeAttributesERK20QXmlStreamAttributes(void* qthis, void* arg0);
  // proto:  void QXmlStreamWriter::writeDefaultNamespace(const QString & namespaceUri);
extern void _ZN16QXmlStreamWriter21writeDefaultNamespaceERK7QString(void* qthis, void* arg0);
  // proto:  QIODevice * QXmlStreamWriter::device();
extern void _ZNK16QXmlStreamWriter6deviceEv(void* qthis);
  // proto:  void QXmlStreamWriter::writeCurrentToken(const QXmlStreamReader & reader);
extern void _ZN16QXmlStreamWriter17writeCurrentTokenERK16QXmlStreamReader(void* qthis, void* arg0);
  // proto:  void QXmlStreamWriter::QXmlStreamWriter(QByteArray * array);
extern void* dector_ZN16QXmlStreamWriterC1EP10QByteArray(void* arg0);
extern void _ZN16QXmlStreamWriterC1EP10QByteArray(void* qthis, void* arg0);
  // proto:  void QXmlStreamWriter::writeTextElement(const QString & qualifiedName, const QString & text);
extern void _ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QXmlStreamWriter::writeEmptyElement(const QString & namespaceUri, const QString & name);
extern void _ZN16QXmlStreamWriter17writeEmptyElementERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QXmlStreamWriter::QXmlStreamWriter(QIODevice * device);
extern void* dector_ZN16QXmlStreamWriterC1EP9QIODevice(void* arg0);
extern void _ZN16QXmlStreamWriterC1EP9QIODevice(void* qthis, void* arg0);
  // proto:  void QXmlStreamWriter::writeStartElement(const QString & qualifiedName);
extern void _ZN16QXmlStreamWriter17writeStartElementERK7QString(void* qthis, void* arg0);
  // proto:  void QXmlStreamNotationDeclaration::QXmlStreamNotationDeclaration(const QXmlStreamNotationDeclaration & );
extern void* dector_ZN29QXmlStreamNotationDeclarationC1ERKS_(void* arg0);
extern void _ZN29QXmlStreamNotationDeclarationC1ERKS_(void* qthis, void* arg0);
  // proto:  QStringRef QXmlStreamNotationDeclaration::publicId();
extern void demth_ZNK29QXmlStreamNotationDeclaration8publicIdEv(void* qthis);
  // proto:  QStringRef QXmlStreamNotationDeclaration::name();
extern void demth_ZNK29QXmlStreamNotationDeclaration4nameEv(void* qthis);
  // proto:  void QXmlStreamNotationDeclaration::~QXmlStreamNotationDeclaration();
extern void _ZN29QXmlStreamNotationDeclarationD0Ev(void* qthis);
  // proto:  QStringRef QXmlStreamNotationDeclaration::systemId();
extern void demth_ZNK29QXmlStreamNotationDeclaration8systemIdEv(void* qthis);
  // proto:  void QXmlStreamNotationDeclaration::QXmlStreamNotationDeclaration();
extern void* dector_ZN29QXmlStreamNotationDeclarationC1Ev();
extern void _ZN29QXmlStreamNotationDeclarationC1Ev(void* qthis);
  // proto:  void QXmlStreamAttribute::QXmlStreamAttribute(const QString & qualifiedName, const QString & value);
extern void* dector_ZN19QXmlStreamAttributeC1ERK7QStringS2_(void* arg0, void* arg1);
extern void _ZN19QXmlStreamAttributeC1ERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  QStringRef QXmlStreamAttribute::qualifiedName();
extern void demth_ZNK19QXmlStreamAttribute13qualifiedNameEv(void* qthis);
  // proto:  void QXmlStreamAttribute::~QXmlStreamAttribute();
extern void _ZN19QXmlStreamAttributeD0Ev(void* qthis);
  // proto:  QStringRef QXmlStreamAttribute::value();
extern void demth_ZNK19QXmlStreamAttribute5valueEv(void* qthis);
  // proto:  QStringRef QXmlStreamAttribute::namespaceUri();
extern void demth_ZNK19QXmlStreamAttribute12namespaceUriEv(void* qthis);
  // proto:  void QXmlStreamAttribute::QXmlStreamAttribute();
extern void* dector_ZN19QXmlStreamAttributeC1Ev();
extern void _ZN19QXmlStreamAttributeC1Ev(void* qthis);
  // proto:  void QXmlStreamAttribute::QXmlStreamAttribute(const QXmlStreamAttribute & );
extern void* dector_ZN19QXmlStreamAttributeC1ERKS_(void* arg0);
extern void _ZN19QXmlStreamAttributeC1ERKS_(void* qthis, void* arg0);
  // proto:  void QXmlStreamAttribute::QXmlStreamAttribute(const QString & namespaceUri, const QString & name, const QString & value);
extern void* dector_ZN19QXmlStreamAttributeC1ERK7QStringS2_S2_(void* arg0, void* arg1, void* arg2);
extern void _ZN19QXmlStreamAttributeC1ERK7QStringS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  bool QXmlStreamAttribute::isDefault();
extern void demth_ZNK19QXmlStreamAttribute9isDefaultEv(void* qthis);
  // proto:  QStringRef QXmlStreamAttribute::prefix();
extern void demth_ZNK19QXmlStreamAttribute6prefixEv(void* qthis);
  // proto:  QStringRef QXmlStreamAttribute::name();
extern void demth_ZNK19QXmlStreamAttribute4nameEv(void* qthis);
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

// class sizeof(QXmlStreamStringRef)=16
type QXmlStreamStringRef struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QXmlStreamReader)=1
type QXmlStreamReader struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QXmlStreamEntityResolver)=8
type QXmlStreamEntityResolver struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QXmlStreamNamespaceDeclaration)=40
type QXmlStreamNamespaceDeclaration struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QXmlStreamEntityDeclaration)=88
type QXmlStreamEntityDeclaration struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QXmlStreamAttributes)=1
type QXmlStreamAttributes struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QXmlStreamWriter)=1
type QXmlStreamWriter struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QXmlStreamNotationDeclaration)=56
type QXmlStreamNotationDeclaration struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QXmlStreamAttribute)=80
type QXmlStreamAttribute struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  int QXmlStreamStringRef::size();
func (this *QXmlStreamStringRef) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamStringRef4sizeEv
  default:
    qtrt.ErrorResolve("QXmlStreamStringRef", "size", args)
  }

}

  // proto:  void QXmlStreamStringRef::clear();
func (this *QXmlStreamStringRef) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QXmlStreamStringRef5clearEv
  default:
    qtrt.ErrorResolve("QXmlStreamStringRef", "clear", args)
  }

}

  // proto:  const QString * QXmlStreamStringRef::string();
func (this *QXmlStreamStringRef) string(args ...interface{}) () {
  // string()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamStringRef6stringEv
  default:
    qtrt.ErrorResolve("QXmlStreamStringRef", "string", args)
  }

}

  // proto:  void QXmlStreamStringRef::QXmlStreamStringRef(const QString & aString);
func NewQXmlStreamStringRef(args ...interface{}) QXmlStreamStringRef {
  return QXmlStreamStringRef{}
}

  // proto:  void QXmlStreamStringRef::~QXmlStreamStringRef();
func (this *QXmlStreamStringRef) FreeQXmlStreamStringRef(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QXmlStreamStringRef", "~QXmlStreamStringRef", args)
  }

}

  // proto:  int QXmlStreamStringRef::position();
func (this *QXmlStreamStringRef) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamStringRef8positionEv
  default:
    qtrt.ErrorResolve("QXmlStreamStringRef", "position", args)
  }

}

  // proto:  QStringRef QXmlStreamReader::name();
func (this *QXmlStreamReader) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader4nameEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "name", args)
  }

}

  // proto:  QXmlStreamEntityResolver * QXmlStreamReader::entityResolver();
func (this *QXmlStreamReader) entityResolver(args ...interface{}) () {
  // entityResolver()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader14entityResolverEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "entityResolver", args)
  }

}

  // proto:  bool QXmlStreamReader::namespaceProcessing();
func (this *QXmlStreamReader) namespaceProcessing(args ...interface{}) () {
  // namespaceProcessing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader19namespaceProcessingEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "namespaceProcessing", args)
  }

}

  // proto:  bool QXmlStreamReader::isStartElement();
func (this *QXmlStreamReader) isStartElement(args ...interface{}) () {
  // isStartElement()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader14isStartElementEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isStartElement", args)
  }

}

  // proto:  bool QXmlStreamReader::isStandaloneDocument();
func (this *QXmlStreamReader) isStandaloneDocument(args ...interface{}) () {
  // isStandaloneDocument()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader20isStandaloneDocumentEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isStandaloneDocument", args)
  }

}

  // proto:  qint64 QXmlStreamReader::lineNumber();
func (this *QXmlStreamReader) lineNumber(args ...interface{}) () {
  // lineNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader10lineNumberEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "lineNumber", args)
  }

}

  // proto:  void QXmlStreamReader::clear();
func (this *QXmlStreamReader) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader5clearEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "clear", args)
  }

}

  // proto:  QStringRef QXmlStreamReader::processingInstructionData();
func (this *QXmlStreamReader) processingInstructionData(args ...interface{}) () {
  // processingInstructionData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader25processingInstructionDataEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "processingInstructionData", args)
  }

}

  // proto:  void QXmlStreamReader::addData(const QString & data);
func (this *QXmlStreamReader) addData(args ...interface{}) () {
  // addData(const class QString &)
  // addData(const char *)
  // addData(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader7addDataERK7QString
  case 1:
    // invoke: _ZN16QXmlStreamReader7addDataEPKc
  case 2:
    // invoke: _ZN16QXmlStreamReader7addDataERK10QByteArray
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "addData", args)
  }

}

  // proto:  QStringRef QXmlStreamReader::dtdPublicId();
func (this *QXmlStreamReader) dtdPublicId(args ...interface{}) () {
  // dtdPublicId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader11dtdPublicIdEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "dtdPublicId", args)
  }

}

  // proto:  QStringRef QXmlStreamReader::documentEncoding();
func (this *QXmlStreamReader) documentEncoding(args ...interface{}) () {
  // documentEncoding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader16documentEncodingEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "documentEncoding", args)
  }

}

  // proto:  qint64 QXmlStreamReader::characterOffset();
func (this *QXmlStreamReader) characterOffset(args ...interface{}) () {
  // characterOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader15characterOffsetEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "characterOffset", args)
  }

}

  // proto:  QXmlStreamAttributes QXmlStreamReader::attributes();
func (this *QXmlStreamReader) attributes(args ...interface{}) () {
  // attributes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader10attributesEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "attributes", args)
  }

}

  // proto:  QString QXmlStreamReader::tokenString();
func (this *QXmlStreamReader) tokenString(args ...interface{}) () {
  // tokenString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader11tokenStringEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "tokenString", args)
  }

}

  // proto:  void QXmlStreamReader::addExtraNamespaceDeclaration(const QXmlStreamNamespaceDeclaration & extraNamespaceDeclaraction);
func (this *QXmlStreamReader) addExtraNamespaceDeclaration(args ...interface{}) () {
  // addExtraNamespaceDeclaration(const class QXmlStreamNamespaceDeclaration &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QXmlStreamNamespaceDeclaration{}) // "const QXmlStreamNamespaceDeclaration &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader28addExtraNamespaceDeclarationERK30QXmlStreamNamespaceDeclaration
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "addExtraNamespaceDeclaration", args)
  }

}

  // proto:  void QXmlStreamReader::QXmlStreamReader(const QByteArray & data);
func NewQXmlStreamReader(args ...interface{}) QXmlStreamReader {
  return QXmlStreamReader{}
}

  // proto:  QStringRef QXmlStreamReader::qualifiedName();
func (this *QXmlStreamReader) qualifiedName(args ...interface{}) () {
  // qualifiedName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader13qualifiedNameEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "qualifiedName", args)
  }

}

  // proto:  QIODevice * QXmlStreamReader::device();
func (this *QXmlStreamReader) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader6deviceEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "device", args)
  }

}

  // proto:  QStringRef QXmlStreamReader::namespaceUri();
func (this *QXmlStreamReader) namespaceUri(args ...interface{}) () {
  // namespaceUri()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader12namespaceUriEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "namespaceUri", args)
  }

}

  // proto:  QStringRef QXmlStreamReader::text();
func (this *QXmlStreamReader) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader4textEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "text", args)
  }

}

  // proto:  void QXmlStreamReader::setDevice(QIODevice * device);
func (this *QXmlStreamReader) setDevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader9setDeviceEP9QIODevice
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "setDevice", args)
  }

}

  // proto:  QStringRef QXmlStreamReader::documentVersion();
func (this *QXmlStreamReader) documentVersion(args ...interface{}) () {
  // documentVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader15documentVersionEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "documentVersion", args)
  }

}

  // proto:  bool QXmlStreamReader::isDTD();
func (this *QXmlStreamReader) isDTD(args ...interface{}) () {
  // isDTD()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader5isDTDEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isDTD", args)
  }

}

  // proto:  bool QXmlStreamReader::isStartDocument();
func (this *QXmlStreamReader) isStartDocument(args ...interface{}) () {
  // isStartDocument()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader15isStartDocumentEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isStartDocument", args)
  }

}

  // proto:  QString QXmlStreamReader::errorString();
func (this *QXmlStreamReader) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader11errorStringEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "errorString", args)
  }

}

  // proto:  bool QXmlStreamReader::isProcessingInstruction();
func (this *QXmlStreamReader) isProcessingInstruction(args ...interface{}) () {
  // isProcessingInstruction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader23isProcessingInstructionEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isProcessingInstruction", args)
  }

}

  // proto:  void QXmlStreamReader::setEntityResolver(QXmlStreamEntityResolver * resolver);
func (this *QXmlStreamReader) setEntityResolver(args ...interface{}) () {
  // setEntityResolver(class QXmlStreamEntityResolver *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QXmlStreamEntityResolver{}) // "QXmlStreamEntityResolver *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader17setEntityResolverEP24QXmlStreamEntityResolver
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "setEntityResolver", args)
  }

}

  // proto:  bool QXmlStreamReader::isCharacters();
func (this *QXmlStreamReader) isCharacters(args ...interface{}) () {
  // isCharacters()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader12isCharactersEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isCharacters", args)
  }

}

  // proto:  QXmlStreamEntityDeclarations QXmlStreamReader::entityDeclarations();
func (this *QXmlStreamReader) entityDeclarations(args ...interface{}) () {
  // entityDeclarations()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader18entityDeclarationsEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "entityDeclarations", args)
  }

}

  // proto:  bool QXmlStreamReader::isWhitespace();
func (this *QXmlStreamReader) isWhitespace(args ...interface{}) () {
  // isWhitespace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader12isWhitespaceEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isWhitespace", args)
  }

}

  // proto:  qint64 QXmlStreamReader::columnNumber();
func (this *QXmlStreamReader) columnNumber(args ...interface{}) () {
  // columnNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader12columnNumberEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "columnNumber", args)
  }

}

  // proto:  bool QXmlStreamReader::hasError();
func (this *QXmlStreamReader) hasError(args ...interface{}) () {
  // hasError()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader8hasErrorEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "hasError", args)
  }

}

  // proto:  bool QXmlStreamReader::isCDATA();
func (this *QXmlStreamReader) isCDATA(args ...interface{}) () {
  // isCDATA()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader7isCDATAEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isCDATA", args)
  }

}

  // proto:  void QXmlStreamReader::~QXmlStreamReader();
func (this *QXmlStreamReader) FreeQXmlStreamReader(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "~QXmlStreamReader", args)
  }

}

  // proto:  QStringRef QXmlStreamReader::processingInstructionTarget();
func (this *QXmlStreamReader) processingInstructionTarget(args ...interface{}) () {
  // processingInstructionTarget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader27processingInstructionTargetEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "processingInstructionTarget", args)
  }

}

  // proto:  QStringRef QXmlStreamReader::dtdSystemId();
func (this *QXmlStreamReader) dtdSystemId(args ...interface{}) () {
  // dtdSystemId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader11dtdSystemIdEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "dtdSystemId", args)
  }

}

  // proto:  QStringRef QXmlStreamReader::prefix();
func (this *QXmlStreamReader) prefix(args ...interface{}) () {
  // prefix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader6prefixEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "prefix", args)
  }

}

  // proto:  bool QXmlStreamReader::isEndElement();
func (this *QXmlStreamReader) isEndElement(args ...interface{}) () {
  // isEndElement()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader12isEndElementEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isEndElement", args)
  }

}

  // proto:  QXmlStreamNotationDeclarations QXmlStreamReader::notationDeclarations();
func (this *QXmlStreamReader) notationDeclarations(args ...interface{}) () {
  // notationDeclarations()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader20notationDeclarationsEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "notationDeclarations", args)
  }

}

  // proto:  QXmlStreamNamespaceDeclarations QXmlStreamReader::namespaceDeclarations();
func (this *QXmlStreamReader) namespaceDeclarations(args ...interface{}) () {
  // namespaceDeclarations()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader21namespaceDeclarationsEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "namespaceDeclarations", args)
  }

}

  // proto:  void QXmlStreamReader::setNamespaceProcessing(bool );
func (this *QXmlStreamReader) setNamespaceProcessing(args ...interface{}) () {
  // setNamespaceProcessing(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader22setNamespaceProcessingEb
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "setNamespaceProcessing", args)
  }

}

  // proto:  void QXmlStreamReader::raiseError(const QString & message);
func (this *QXmlStreamReader) raiseError(args ...interface{}) () {
  // raiseError(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader10raiseErrorERK7QString
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "raiseError", args)
  }

}

  // proto:  QStringRef QXmlStreamReader::dtdName();
func (this *QXmlStreamReader) dtdName(args ...interface{}) () {
  // dtdName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader7dtdNameEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "dtdName", args)
  }

}

  // proto:  bool QXmlStreamReader::isEndDocument();
func (this *QXmlStreamReader) isEndDocument(args ...interface{}) () {
  // isEndDocument()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader13isEndDocumentEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isEndDocument", args)
  }

}

  // proto:  bool QXmlStreamReader::readNextStartElement();
func (this *QXmlStreamReader) readNextStartElement(args ...interface{}) () {
  // readNextStartElement()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader20readNextStartElementEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "readNextStartElement", args)
  }

}

  // proto:  bool QXmlStreamReader::isComment();
func (this *QXmlStreamReader) isComment(args ...interface{}) () {
  // isComment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader9isCommentEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isComment", args)
  }

}

  // proto:  void QXmlStreamReader::skipCurrentElement();
func (this *QXmlStreamReader) skipCurrentElement(args ...interface{}) () {
  // skipCurrentElement()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader18skipCurrentElementEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "skipCurrentElement", args)
  }

}

  // proto:  bool QXmlStreamReader::isEntityReference();
func (this *QXmlStreamReader) isEntityReference(args ...interface{}) () {
  // isEntityReference()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader17isEntityReferenceEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isEntityReference", args)
  }

}

  // proto:  bool QXmlStreamReader::atEnd();
func (this *QXmlStreamReader) atEnd(args ...interface{}) () {
  // atEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader5atEndEv
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "atEnd", args)
  }

}

  // proto:  QString QXmlStreamEntityResolver::resolveEntity(const QString & publicId, const QString & systemId);
func (this *QXmlStreamEntityResolver) resolveEntity(args ...interface{}) () {
  // resolveEntity(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QXmlStreamEntityResolver13resolveEntityERK7QStringS2_
  default:
    qtrt.ErrorResolve("QXmlStreamEntityResolver", "resolveEntity", args)
  }

}

  // proto:  QString QXmlStreamEntityResolver::resolveUndeclaredEntity(const QString & name);
func (this *QXmlStreamEntityResolver) resolveUndeclaredEntity(args ...interface{}) () {
  // resolveUndeclaredEntity(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QXmlStreamEntityResolver23resolveUndeclaredEntityERK7QString
  default:
    qtrt.ErrorResolve("QXmlStreamEntityResolver", "resolveUndeclaredEntity", args)
  }

}

  // proto:  void QXmlStreamEntityResolver::~QXmlStreamEntityResolver();
func (this *QXmlStreamEntityResolver) FreeQXmlStreamEntityResolver(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QXmlStreamEntityResolver", "~QXmlStreamEntityResolver", args)
  }

}

  // proto:  QStringRef QXmlStreamNamespaceDeclaration::namespaceUri();
func (this *QXmlStreamNamespaceDeclaration) namespaceUri(args ...interface{}) () {
  // namespaceUri()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK30QXmlStreamNamespaceDeclaration12namespaceUriEv
  default:
    qtrt.ErrorResolve("QXmlStreamNamespaceDeclaration", "namespaceUri", args)
  }

}

  // proto:  void QXmlStreamNamespaceDeclaration::QXmlStreamNamespaceDeclaration(const QXmlStreamNamespaceDeclaration & );
func NewQXmlStreamNamespaceDeclaration(args ...interface{}) QXmlStreamNamespaceDeclaration {
  return QXmlStreamNamespaceDeclaration{}
}

  // proto:  QStringRef QXmlStreamNamespaceDeclaration::prefix();
func (this *QXmlStreamNamespaceDeclaration) prefix(args ...interface{}) () {
  // prefix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK30QXmlStreamNamespaceDeclaration6prefixEv
  default:
    qtrt.ErrorResolve("QXmlStreamNamespaceDeclaration", "prefix", args)
  }

}

  // proto:  void QXmlStreamNamespaceDeclaration::~QXmlStreamNamespaceDeclaration();
func (this *QXmlStreamNamespaceDeclaration) FreeQXmlStreamNamespaceDeclaration(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QXmlStreamNamespaceDeclaration", "~QXmlStreamNamespaceDeclaration", args)
  }

}

  // proto:  void QXmlStreamEntityDeclaration::~QXmlStreamEntityDeclaration();
func (this *QXmlStreamEntityDeclaration) FreeQXmlStreamEntityDeclaration(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QXmlStreamEntityDeclaration", "~QXmlStreamEntityDeclaration", args)
  }

}

  // proto:  QStringRef QXmlStreamEntityDeclaration::publicId();
func (this *QXmlStreamEntityDeclaration) publicId(args ...interface{}) () {
  // publicId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QXmlStreamEntityDeclaration8publicIdEv
  default:
    qtrt.ErrorResolve("QXmlStreamEntityDeclaration", "publicId", args)
  }

}

  // proto:  QStringRef QXmlStreamEntityDeclaration::name();
func (this *QXmlStreamEntityDeclaration) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QXmlStreamEntityDeclaration4nameEv
  default:
    qtrt.ErrorResolve("QXmlStreamEntityDeclaration", "name", args)
  }

}

  // proto:  void QXmlStreamEntityDeclaration::QXmlStreamEntityDeclaration();
func NewQXmlStreamEntityDeclaration(args ...interface{}) QXmlStreamEntityDeclaration {
  return QXmlStreamEntityDeclaration{}
}

  // proto:  QStringRef QXmlStreamEntityDeclaration::value();
func (this *QXmlStreamEntityDeclaration) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QXmlStreamEntityDeclaration5valueEv
  default:
    qtrt.ErrorResolve("QXmlStreamEntityDeclaration", "value", args)
  }

}

  // proto:  QStringRef QXmlStreamEntityDeclaration::notationName();
func (this *QXmlStreamEntityDeclaration) notationName(args ...interface{}) () {
  // notationName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QXmlStreamEntityDeclaration12notationNameEv
  default:
    qtrt.ErrorResolve("QXmlStreamEntityDeclaration", "notationName", args)
  }

}

  // proto:  QStringRef QXmlStreamEntityDeclaration::systemId();
func (this *QXmlStreamEntityDeclaration) systemId(args ...interface{}) () {
  // systemId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QXmlStreamEntityDeclaration8systemIdEv
  default:
    qtrt.ErrorResolve("QXmlStreamEntityDeclaration", "systemId", args)
  }

}

  // proto:  QStringRef QXmlStreamAttributes::value(const QString & qualifiedName);
func (this *QXmlStreamAttributes) value(args ...interface{}) () {
  // value(const class QString &)
  // value(class QLatin1String, class QLatin1String)
  // value(const class QString &, class QLatin1String)
  // value(const class QString &, const class QString &)
  // value(class QLatin1String)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[1][1] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QXmlStreamAttributes5valueERK7QString
  case 1:
    // invoke: _ZNK20QXmlStreamAttributes5valueE13QLatin1StringS0_
  case 2:
    // invoke: _ZNK20QXmlStreamAttributes5valueERK7QString13QLatin1String
  case 3:
    // invoke: _ZNK20QXmlStreamAttributes5valueERK7QStringS2_
  case 4:
    // invoke: _ZNK20QXmlStreamAttributes5valueE13QLatin1String
  default:
    qtrt.ErrorResolve("QXmlStreamAttributes", "value", args)
  }

}

  // proto:  void QXmlStreamAttributes::QXmlStreamAttributes();
func NewQXmlStreamAttributes(args ...interface{}) QXmlStreamAttributes {
  return QXmlStreamAttributes{}
}

  // proto:  bool QXmlStreamAttributes::hasAttribute(const QString & qualifiedName);
func (this *QXmlStreamAttributes) hasAttribute(args ...interface{}) () {
  // hasAttribute(const class QString &)
  // hasAttribute(const class QString &, const class QString &)
  // hasAttribute(class QLatin1String)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QXmlStreamAttributes12hasAttributeERK7QString
  case 1:
    // invoke: _ZNK20QXmlStreamAttributes12hasAttributeERK7QStringS2_
  case 2:
    // invoke: _ZNK20QXmlStreamAttributes12hasAttributeE13QLatin1String
  default:
    qtrt.ErrorResolve("QXmlStreamAttributes", "hasAttribute", args)
  }

}

  // proto:  void QXmlStreamAttributes::append(const QString & namespaceUri, const QString & name, const QString & value);
func (this *QXmlStreamAttributes) append(args ...interface{}) () {
  // append(const class QString &, const class QString &, const class QString &)
  // append(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QXmlStreamAttributes6appendERK7QStringS2_S2_
  case 1:
    // invoke: _ZN20QXmlStreamAttributes6appendERK7QStringS2_
  default:
    qtrt.ErrorResolve("QXmlStreamAttributes", "append", args)
  }

}

  // proto:  void QXmlStreamWriter::writeEndElement();
func (this *QXmlStreamWriter) writeEndElement(args ...interface{}) () {
  // writeEndElement()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter15writeEndElementEv
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeEndElement", args)
  }

}

  // proto:  void QXmlStreamWriter::QXmlStreamWriter();
func NewQXmlStreamWriter(args ...interface{}) QXmlStreamWriter {
  return QXmlStreamWriter{}
}

  // proto:  void QXmlStreamWriter::writeEndDocument();
func (this *QXmlStreamWriter) writeEndDocument(args ...interface{}) () {
  // writeEndDocument()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter16writeEndDocumentEv
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeEndDocument", args)
  }

}

  // proto:  bool QXmlStreamWriter::autoFormatting();
func (this *QXmlStreamWriter) autoFormatting(args ...interface{}) () {
  // autoFormatting()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamWriter14autoFormattingEv
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "autoFormatting", args)
  }

}

  // proto:  void QXmlStreamWriter::writeStartDocument(const QString & version, bool standalone);
func (this *QXmlStreamWriter) writeStartDocument(args ...interface{}) () {
  // writeStartDocument(const class QString &, _Bool)
  // writeStartDocument(const class QString &)
  // writeStartDocument()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter18writeStartDocumentERK7QStringb
  case 1:
    // invoke: _ZN16QXmlStreamWriter18writeStartDocumentERK7QString
  case 2:
    // invoke: _ZN16QXmlStreamWriter18writeStartDocumentEv
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeStartDocument", args)
  }

}

  // proto:  void QXmlStreamWriter::setCodec(QTextCodec * codec);
func (this *QXmlStreamWriter) setCodec(args ...interface{}) () {
  // setCodec(class QTextCodec *)
  // setCodec(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter8setCodecEP10QTextCodec
  case 1:
    // invoke: _ZN16QXmlStreamWriter8setCodecEPKc
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "setCodec", args)
  }

}

  // proto:  void QXmlStreamWriter::writeProcessingInstruction(const QString & target, const QString & data);
func (this *QXmlStreamWriter) writeProcessingInstruction(args ...interface{}) () {
  // writeProcessingInstruction(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter26writeProcessingInstructionERK7QStringS2_
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeProcessingInstruction", args)
  }

}

  // proto:  void QXmlStreamWriter::writeCharacters(const QString & text);
func (this *QXmlStreamWriter) writeCharacters(args ...interface{}) () {
  // writeCharacters(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter15writeCharactersERK7QString
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeCharacters", args)
  }

}

  // proto:  void QXmlStreamWriter::setDevice(QIODevice * device);
func (this *QXmlStreamWriter) setDevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter9setDeviceEP9QIODevice
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "setDevice", args)
  }

}

  // proto:  void QXmlStreamWriter::writeTextElement(const QString & namespaceUri, const QString & name, const QString & text);
func (this *QXmlStreamWriter) writeTextElement(args ...interface{}) () {
  // writeTextElement(const class QString &, const class QString &, const class QString &)
  // writeTextElement(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_S2_
  case 1:
    // invoke: _ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeTextElement", args)
  }

}

  // proto:  void QXmlStreamWriter::writeAttribute(const QString & qualifiedName, const QString & value);
func (this *QXmlStreamWriter) writeAttribute(args ...interface{}) () {
  // writeAttribute(const class QString &, const class QString &)
  // writeAttribute(const class QXmlStreamAttribute &)
  // writeAttribute(const class QString &, const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QXmlStreamAttribute{}) // "const QXmlStreamAttribute &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_
  case 1:
    // invoke: _ZN16QXmlStreamWriter14writeAttributeERK19QXmlStreamAttribute
  case 2:
    // invoke: _ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_S2_
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeAttribute", args)
  }

}

  // proto:  void QXmlStreamWriter::writeEmptyElement(const QString & qualifiedName);
func (this *QXmlStreamWriter) writeEmptyElement(args ...interface{}) () {
  // writeEmptyElement(const class QString &)
  // writeEmptyElement(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter17writeEmptyElementERK7QString
  case 1:
    // invoke: _ZN16QXmlStreamWriter17writeEmptyElementERK7QStringS2_
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeEmptyElement", args)
  }

}

  // proto:  void QXmlStreamWriter::writeDTD(const QString & dtd);
func (this *QXmlStreamWriter) writeDTD(args ...interface{}) () {
  // writeDTD(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter8writeDTDERK7QString
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeDTD", args)
  }

}

  // proto:  void QXmlStreamWriter::setAutoFormattingIndent(int spacesOrTabs);
func (this *QXmlStreamWriter) setAutoFormattingIndent(args ...interface{}) () {
  // setAutoFormattingIndent(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter23setAutoFormattingIndentEi
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "setAutoFormattingIndent", args)
  }

}

  // proto:  void QXmlStreamWriter::writeStartElement(const QString & namespaceUri, const QString & name);
func (this *QXmlStreamWriter) writeStartElement(args ...interface{}) () {
  // writeStartElement(const class QString &, const class QString &)
  // writeStartElement(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter17writeStartElementERK7QStringS2_
  case 1:
    // invoke: _ZN16QXmlStreamWriter17writeStartElementERK7QString
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeStartElement", args)
  }

}

  // proto:  void QXmlStreamWriter::writeComment(const QString & text);
func (this *QXmlStreamWriter) writeComment(args ...interface{}) () {
  // writeComment(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter12writeCommentERK7QString
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeComment", args)
  }

}

  // proto:  QTextCodec * QXmlStreamWriter::codec();
func (this *QXmlStreamWriter) codec(args ...interface{}) () {
  // codec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamWriter5codecEv
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "codec", args)
  }

}

  // proto:  void QXmlStreamWriter::writeNamespace(const QString & namespaceUri, const QString & prefix);
func (this *QXmlStreamWriter) writeNamespace(args ...interface{}) () {
  // writeNamespace(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter14writeNamespaceERK7QStringS2_
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeNamespace", args)
  }

}

  // proto:  bool QXmlStreamWriter::hasError();
func (this *QXmlStreamWriter) hasError(args ...interface{}) () {
  // hasError()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamWriter8hasErrorEv
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "hasError", args)
  }

}

  // proto:  void QXmlStreamWriter::writeCDATA(const QString & text);
func (this *QXmlStreamWriter) writeCDATA(args ...interface{}) () {
  // writeCDATA(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter10writeCDATAERK7QString
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeCDATA", args)
  }

}

  // proto:  void QXmlStreamWriter::writeEntityReference(const QString & name);
func (this *QXmlStreamWriter) writeEntityReference(args ...interface{}) () {
  // writeEntityReference(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter20writeEntityReferenceERK7QString
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeEntityReference", args)
  }

}

  // proto:  void QXmlStreamWriter::setAutoFormatting(bool );
func (this *QXmlStreamWriter) setAutoFormatting(args ...interface{}) () {
  // setAutoFormatting(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter17setAutoFormattingEb
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "setAutoFormatting", args)
  }

}

  // proto:  int QXmlStreamWriter::autoFormattingIndent();
func (this *QXmlStreamWriter) autoFormattingIndent(args ...interface{}) () {
  // autoFormattingIndent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamWriter20autoFormattingIndentEv
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "autoFormattingIndent", args)
  }

}

  // proto:  void QXmlStreamWriter::~QXmlStreamWriter();
func (this *QXmlStreamWriter) FreeQXmlStreamWriter(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "~QXmlStreamWriter", args)
  }

}

  // proto:  void QXmlStreamWriter::writeAttributes(const QXmlStreamAttributes & attributes);
func (this *QXmlStreamWriter) writeAttributes(args ...interface{}) () {
  // writeAttributes(const class QXmlStreamAttributes &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QXmlStreamAttributes{}) // "const QXmlStreamAttributes &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter15writeAttributesERK20QXmlStreamAttributes
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeAttributes", args)
  }

}

  // proto:  void QXmlStreamWriter::writeDefaultNamespace(const QString & namespaceUri);
func (this *QXmlStreamWriter) writeDefaultNamespace(args ...interface{}) () {
  // writeDefaultNamespace(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter21writeDefaultNamespaceERK7QString
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeDefaultNamespace", args)
  }

}

  // proto:  QIODevice * QXmlStreamWriter::device();
func (this *QXmlStreamWriter) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamWriter6deviceEv
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "device", args)
  }

}

  // proto:  void QXmlStreamWriter::writeCurrentToken(const QXmlStreamReader & reader);
func (this *QXmlStreamWriter) writeCurrentToken(args ...interface{}) () {
  // writeCurrentToken(const class QXmlStreamReader &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QXmlStreamReader{}) // "const QXmlStreamReader &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter17writeCurrentTokenERK16QXmlStreamReader
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeCurrentToken", args)
  }

}

  // proto:  void QXmlStreamNotationDeclaration::QXmlStreamNotationDeclaration(const QXmlStreamNotationDeclaration & );
func NewQXmlStreamNotationDeclaration(args ...interface{}) QXmlStreamNotationDeclaration {
  return QXmlStreamNotationDeclaration{}
}

  // proto:  QStringRef QXmlStreamNotationDeclaration::publicId();
func (this *QXmlStreamNotationDeclaration) publicId(args ...interface{}) () {
  // publicId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QXmlStreamNotationDeclaration8publicIdEv
  default:
    qtrt.ErrorResolve("QXmlStreamNotationDeclaration", "publicId", args)
  }

}

  // proto:  QStringRef QXmlStreamNotationDeclaration::name();
func (this *QXmlStreamNotationDeclaration) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QXmlStreamNotationDeclaration4nameEv
  default:
    qtrt.ErrorResolve("QXmlStreamNotationDeclaration", "name", args)
  }

}

  // proto:  void QXmlStreamNotationDeclaration::~QXmlStreamNotationDeclaration();
func (this *QXmlStreamNotationDeclaration) FreeQXmlStreamNotationDeclaration(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QXmlStreamNotationDeclaration", "~QXmlStreamNotationDeclaration", args)
  }

}

  // proto:  QStringRef QXmlStreamNotationDeclaration::systemId();
func (this *QXmlStreamNotationDeclaration) systemId(args ...interface{}) () {
  // systemId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QXmlStreamNotationDeclaration8systemIdEv
  default:
    qtrt.ErrorResolve("QXmlStreamNotationDeclaration", "systemId", args)
  }

}

  // proto:  void QXmlStreamAttribute::QXmlStreamAttribute(const QString & qualifiedName, const QString & value);
func NewQXmlStreamAttribute(args ...interface{}) QXmlStreamAttribute {
  return QXmlStreamAttribute{}
}

  // proto:  QStringRef QXmlStreamAttribute::qualifiedName();
func (this *QXmlStreamAttribute) qualifiedName(args ...interface{}) () {
  // qualifiedName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamAttribute13qualifiedNameEv
  default:
    qtrt.ErrorResolve("QXmlStreamAttribute", "qualifiedName", args)
  }

}

  // proto:  void QXmlStreamAttribute::~QXmlStreamAttribute();
func (this *QXmlStreamAttribute) FreeQXmlStreamAttribute(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QXmlStreamAttribute", "~QXmlStreamAttribute", args)
  }

}

  // proto:  QStringRef QXmlStreamAttribute::value();
func (this *QXmlStreamAttribute) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamAttribute5valueEv
  default:
    qtrt.ErrorResolve("QXmlStreamAttribute", "value", args)
  }

}

  // proto:  QStringRef QXmlStreamAttribute::namespaceUri();
func (this *QXmlStreamAttribute) namespaceUri(args ...interface{}) () {
  // namespaceUri()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamAttribute12namespaceUriEv
  default:
    qtrt.ErrorResolve("QXmlStreamAttribute", "namespaceUri", args)
  }

}

  // proto:  bool QXmlStreamAttribute::isDefault();
func (this *QXmlStreamAttribute) isDefault(args ...interface{}) () {
  // isDefault()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamAttribute9isDefaultEv
  default:
    qtrt.ErrorResolve("QXmlStreamAttribute", "isDefault", args)
  }

}

  // proto:  QStringRef QXmlStreamAttribute::prefix();
func (this *QXmlStreamAttribute) prefix(args ...interface{}) () {
  // prefix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamAttribute6prefixEv
  default:
    qtrt.ErrorResolve("QXmlStreamAttribute", "prefix", args)
  }

}

  // proto:  QStringRef QXmlStreamAttribute::name();
func (this *QXmlStreamAttribute) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamAttribute4nameEv
  default:
    qtrt.ErrorResolve("QXmlStreamAttribute", "name", args)
  }

}

// <= body block end
