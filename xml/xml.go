// +build !minimal

package xml

//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QDomAttr struct {
	QDomNode
}

type QDomAttr_ITF interface {
	QDomNode_ITF
	QDomAttr_PTR() *QDomAttr
}

func (p *QDomAttr) QDomAttr_PTR() *QDomAttr {
	return p
}

func (p *QDomAttr) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDomNode_PTR().Pointer()
	}
	return nil
}

func (p *QDomAttr) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDomNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQDomAttr(ptr QDomAttr_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomAttr_PTR().Pointer()
	}
	return nil
}

func NewQDomAttrFromPointer(ptr unsafe.Pointer) *QDomAttr {
	var n = new(QDomAttr)
	n.SetPointer(ptr)
	return n
}

func newQDomAttrFromPointer(ptr unsafe.Pointer) *QDomAttr {
	var n = NewQDomAttrFromPointer(ptr)
	return n
}

func NewQDomAttr() *QDomAttr {
	defer qt.Recovering("QDomAttr::QDomAttr")

	return newQDomAttrFromPointer(C.QDomAttr_NewQDomAttr())
}

func NewQDomAttr2(x QDomAttr_ITF) *QDomAttr {
	defer qt.Recovering("QDomAttr::QDomAttr")

	return newQDomAttrFromPointer(C.QDomAttr_NewQDomAttr2(PointerFromQDomAttr(x)))
}

func (ptr *QDomAttr) Name() string {
	defer qt.Recovering("QDomAttr::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomAttr_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomAttr) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomAttr::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomAttr_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomAttr) OwnerElement() *QDomElement {
	defer qt.Recovering("QDomAttr::ownerElement")

	if ptr.Pointer() != nil {
		return NewQDomElementFromPointer(C.QDomAttr_OwnerElement(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomAttr) SetValue(v string) {
	defer qt.Recovering("QDomAttr::setValue")

	if ptr.Pointer() != nil {
		C.QDomAttr_SetValue(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QDomAttr) Specified() bool {
	defer qt.Recovering("QDomAttr::specified")

	if ptr.Pointer() != nil {
		return C.QDomAttr_Specified(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomAttr) Value() string {
	defer qt.Recovering("QDomAttr::value")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomAttr_Value(ptr.Pointer()))
	}
	return ""
}

type QDomCDATASection struct {
	QDomText
}

type QDomCDATASection_ITF interface {
	QDomText_ITF
	QDomCDATASection_PTR() *QDomCDATASection
}

func (p *QDomCDATASection) QDomCDATASection_PTR() *QDomCDATASection {
	return p
}

func (p *QDomCDATASection) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDomText_PTR().Pointer()
	}
	return nil
}

func (p *QDomCDATASection) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDomText_PTR().SetPointer(ptr)
	}
}

func PointerFromQDomCDATASection(ptr QDomCDATASection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomCDATASection_PTR().Pointer()
	}
	return nil
}

func NewQDomCDATASectionFromPointer(ptr unsafe.Pointer) *QDomCDATASection {
	var n = new(QDomCDATASection)
	n.SetPointer(ptr)
	return n
}

func newQDomCDATASectionFromPointer(ptr unsafe.Pointer) *QDomCDATASection {
	var n = NewQDomCDATASectionFromPointer(ptr)
	return n
}

func NewQDomCDATASection() *QDomCDATASection {
	defer qt.Recovering("QDomCDATASection::QDomCDATASection")

	return newQDomCDATASectionFromPointer(C.QDomCDATASection_NewQDomCDATASection())
}

func NewQDomCDATASection2(x QDomCDATASection_ITF) *QDomCDATASection {
	defer qt.Recovering("QDomCDATASection::QDomCDATASection")

	return newQDomCDATASectionFromPointer(C.QDomCDATASection_NewQDomCDATASection2(PointerFromQDomCDATASection(x)))
}

func (ptr *QDomCDATASection) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomCDATASection::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomCDATASection_NodeType(ptr.Pointer()))
	}
	return 0
}

type QDomCharacterData struct {
	QDomNode
}

type QDomCharacterData_ITF interface {
	QDomNode_ITF
	QDomCharacterData_PTR() *QDomCharacterData
}

func (p *QDomCharacterData) QDomCharacterData_PTR() *QDomCharacterData {
	return p
}

func (p *QDomCharacterData) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDomNode_PTR().Pointer()
	}
	return nil
}

func (p *QDomCharacterData) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDomNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQDomCharacterData(ptr QDomCharacterData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomCharacterData_PTR().Pointer()
	}
	return nil
}

func NewQDomCharacterDataFromPointer(ptr unsafe.Pointer) *QDomCharacterData {
	var n = new(QDomCharacterData)
	n.SetPointer(ptr)
	return n
}

func newQDomCharacterDataFromPointer(ptr unsafe.Pointer) *QDomCharacterData {
	var n = NewQDomCharacterDataFromPointer(ptr)
	return n
}

func NewQDomCharacterData() *QDomCharacterData {
	defer qt.Recovering("QDomCharacterData::QDomCharacterData")

	return newQDomCharacterDataFromPointer(C.QDomCharacterData_NewQDomCharacterData())
}

func NewQDomCharacterData2(x QDomCharacterData_ITF) *QDomCharacterData {
	defer qt.Recovering("QDomCharacterData::QDomCharacterData")

	return newQDomCharacterDataFromPointer(C.QDomCharacterData_NewQDomCharacterData2(PointerFromQDomCharacterData(x)))
}

func (ptr *QDomCharacterData) AppendData(arg string) {
	defer qt.Recovering("QDomCharacterData::appendData")

	if ptr.Pointer() != nil {
		C.QDomCharacterData_AppendData(ptr.Pointer(), C.CString(arg))
	}
}

func (ptr *QDomCharacterData) Data() string {
	defer qt.Recovering("QDomCharacterData::data")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomCharacterData_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomCharacterData) Length() int {
	defer qt.Recovering("QDomCharacterData::length")

	if ptr.Pointer() != nil {
		return int(C.QDomCharacterData_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomCharacterData) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomCharacterData::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomCharacterData_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomCharacterData) SetData(v string) {
	defer qt.Recovering("QDomCharacterData::setData")

	if ptr.Pointer() != nil {
		C.QDomCharacterData_SetData(ptr.Pointer(), C.CString(v))
	}
}

type QDomComment struct {
	QDomCharacterData
}

type QDomComment_ITF interface {
	QDomCharacterData_ITF
	QDomComment_PTR() *QDomComment
}

func (p *QDomComment) QDomComment_PTR() *QDomComment {
	return p
}

func (p *QDomComment) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDomCharacterData_PTR().Pointer()
	}
	return nil
}

func (p *QDomComment) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDomCharacterData_PTR().SetPointer(ptr)
	}
}

func PointerFromQDomComment(ptr QDomComment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomComment_PTR().Pointer()
	}
	return nil
}

func NewQDomCommentFromPointer(ptr unsafe.Pointer) *QDomComment {
	var n = new(QDomComment)
	n.SetPointer(ptr)
	return n
}

func newQDomCommentFromPointer(ptr unsafe.Pointer) *QDomComment {
	var n = NewQDomCommentFromPointer(ptr)
	return n
}

func NewQDomComment() *QDomComment {
	defer qt.Recovering("QDomComment::QDomComment")

	return newQDomCommentFromPointer(C.QDomComment_NewQDomComment())
}

func NewQDomComment2(x QDomComment_ITF) *QDomComment {
	defer qt.Recovering("QDomComment::QDomComment")

	return newQDomCommentFromPointer(C.QDomComment_NewQDomComment2(PointerFromQDomComment(x)))
}

func (ptr *QDomComment) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomComment::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomComment_NodeType(ptr.Pointer()))
	}
	return 0
}

type QDomDocument struct {
	QDomNode
}

type QDomDocument_ITF interface {
	QDomNode_ITF
	QDomDocument_PTR() *QDomDocument
}

func (p *QDomDocument) QDomDocument_PTR() *QDomDocument {
	return p
}

func (p *QDomDocument) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDomNode_PTR().Pointer()
	}
	return nil
}

func (p *QDomDocument) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDomNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQDomDocument(ptr QDomDocument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomDocument_PTR().Pointer()
	}
	return nil
}

func NewQDomDocumentFromPointer(ptr unsafe.Pointer) *QDomDocument {
	var n = new(QDomDocument)
	n.SetPointer(ptr)
	return n
}

func newQDomDocumentFromPointer(ptr unsafe.Pointer) *QDomDocument {
	var n = NewQDomDocumentFromPointer(ptr)
	return n
}

func NewQDomDocument() *QDomDocument {
	defer qt.Recovering("QDomDocument::QDomDocument")

	return newQDomDocumentFromPointer(C.QDomDocument_NewQDomDocument())
}

func NewQDomDocument4(x QDomDocument_ITF) *QDomDocument {
	defer qt.Recovering("QDomDocument::QDomDocument")

	return newQDomDocumentFromPointer(C.QDomDocument_NewQDomDocument4(PointerFromQDomDocument(x)))
}

func NewQDomDocument3(doctype QDomDocumentType_ITF) *QDomDocument {
	defer qt.Recovering("QDomDocument::QDomDocument")

	return newQDomDocumentFromPointer(C.QDomDocument_NewQDomDocument3(PointerFromQDomDocumentType(doctype)))
}

func NewQDomDocument2(name string) *QDomDocument {
	defer qt.Recovering("QDomDocument::QDomDocument")

	return newQDomDocumentFromPointer(C.QDomDocument_NewQDomDocument2(C.CString(name)))
}

func (ptr *QDomDocument) CreateAttribute(name string) *QDomAttr {
	defer qt.Recovering("QDomDocument::createAttribute")

	if ptr.Pointer() != nil {
		return NewQDomAttrFromPointer(C.QDomDocument_CreateAttribute(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QDomDocument) CreateAttributeNS(nsURI string, qName string) *QDomAttr {
	defer qt.Recovering("QDomDocument::createAttributeNS")

	if ptr.Pointer() != nil {
		return NewQDomAttrFromPointer(C.QDomDocument_CreateAttributeNS(ptr.Pointer(), C.CString(nsURI), C.CString(qName)))
	}
	return nil
}

func (ptr *QDomDocument) CreateCDATASection(value string) *QDomCDATASection {
	defer qt.Recovering("QDomDocument::createCDATASection")

	if ptr.Pointer() != nil {
		return NewQDomCDATASectionFromPointer(C.QDomDocument_CreateCDATASection(ptr.Pointer(), C.CString(value)))
	}
	return nil
}

func (ptr *QDomDocument) CreateComment(value string) *QDomComment {
	defer qt.Recovering("QDomDocument::createComment")

	if ptr.Pointer() != nil {
		return NewQDomCommentFromPointer(C.QDomDocument_CreateComment(ptr.Pointer(), C.CString(value)))
	}
	return nil
}

func (ptr *QDomDocument) CreateDocumentFragment() *QDomDocumentFragment {
	defer qt.Recovering("QDomDocument::createDocumentFragment")

	if ptr.Pointer() != nil {
		return NewQDomDocumentFragmentFromPointer(C.QDomDocument_CreateDocumentFragment(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomDocument) CreateElement(tagName string) *QDomElement {
	defer qt.Recovering("QDomDocument::createElement")

	if ptr.Pointer() != nil {
		return NewQDomElementFromPointer(C.QDomDocument_CreateElement(ptr.Pointer(), C.CString(tagName)))
	}
	return nil
}

func (ptr *QDomDocument) CreateElementNS(nsURI string, qName string) *QDomElement {
	defer qt.Recovering("QDomDocument::createElementNS")

	if ptr.Pointer() != nil {
		return NewQDomElementFromPointer(C.QDomDocument_CreateElementNS(ptr.Pointer(), C.CString(nsURI), C.CString(qName)))
	}
	return nil
}

func (ptr *QDomDocument) CreateEntityReference(name string) *QDomEntityReference {
	defer qt.Recovering("QDomDocument::createEntityReference")

	if ptr.Pointer() != nil {
		return NewQDomEntityReferenceFromPointer(C.QDomDocument_CreateEntityReference(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QDomDocument) CreateProcessingInstruction(target string, data string) *QDomProcessingInstruction {
	defer qt.Recovering("QDomDocument::createProcessingInstruction")

	if ptr.Pointer() != nil {
		return NewQDomProcessingInstructionFromPointer(C.QDomDocument_CreateProcessingInstruction(ptr.Pointer(), C.CString(target), C.CString(data)))
	}
	return nil
}

func (ptr *QDomDocument) CreateTextNode(value string) *QDomText {
	defer qt.Recovering("QDomDocument::createTextNode")

	if ptr.Pointer() != nil {
		return NewQDomTextFromPointer(C.QDomDocument_CreateTextNode(ptr.Pointer(), C.CString(value)))
	}
	return nil
}

func (ptr *QDomDocument) Doctype() *QDomDocumentType {
	defer qt.Recovering("QDomDocument::doctype")

	if ptr.Pointer() != nil {
		return NewQDomDocumentTypeFromPointer(C.QDomDocument_Doctype(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomDocument) DocumentElement() *QDomElement {
	defer qt.Recovering("QDomDocument::documentElement")

	if ptr.Pointer() != nil {
		return NewQDomElementFromPointer(C.QDomDocument_DocumentElement(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomDocument) ElementById(elementId string) *QDomElement {
	defer qt.Recovering("QDomDocument::elementById")

	if ptr.Pointer() != nil {
		return NewQDomElementFromPointer(C.QDomDocument_ElementById(ptr.Pointer(), C.CString(elementId)))
	}
	return nil
}

func (ptr *QDomDocument) ElementsByTagName(tagname string) *QDomNodeList {
	defer qt.Recovering("QDomDocument::elementsByTagName")

	if ptr.Pointer() != nil {
		return NewQDomNodeListFromPointer(C.QDomDocument_ElementsByTagName(ptr.Pointer(), C.CString(tagname)))
	}
	return nil
}

func (ptr *QDomDocument) ElementsByTagNameNS(nsURI string, localName string) *QDomNodeList {
	defer qt.Recovering("QDomDocument::elementsByTagNameNS")

	if ptr.Pointer() != nil {
		return NewQDomNodeListFromPointer(C.QDomDocument_ElementsByTagNameNS(ptr.Pointer(), C.CString(nsURI), C.CString(localName)))
	}
	return nil
}

func (ptr *QDomDocument) Implementation() *QDomImplementation {
	defer qt.Recovering("QDomDocument::implementation")

	if ptr.Pointer() != nil {
		return NewQDomImplementationFromPointer(C.QDomDocument_Implementation(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomDocument) ImportNode(importedNode QDomNode_ITF, deep bool) *QDomNode {
	defer qt.Recovering("QDomDocument::importNode")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomDocument_ImportNode(ptr.Pointer(), PointerFromQDomNode(importedNode), C.int(qt.GoBoolToInt(deep))))
	}
	return nil
}

func (ptr *QDomDocument) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomDocument::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomDocument_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomDocument) SetContent7(dev core.QIODevice_ITF, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent7(ptr.Pointer(), core.PointerFromQIODevice(dev), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent3(dev core.QIODevice_ITF, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent3(ptr.Pointer(), core.PointerFromQIODevice(dev), C.int(qt.GoBoolToInt(namespaceProcessing)), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent8(source QXmlInputSource_ITF, reader QXmlReader_ITF, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent8(ptr.Pointer(), PointerFromQXmlInputSource(source), PointerFromQXmlReader(reader), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent4(source QXmlInputSource_ITF, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent4(ptr.Pointer(), PointerFromQXmlInputSource(source), C.int(qt.GoBoolToInt(namespaceProcessing)), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent5(buffer string, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent5(ptr.Pointer(), C.CString(buffer), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent(data string, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent(ptr.Pointer(), C.CString(data), C.int(qt.GoBoolToInt(namespaceProcessing)), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent6(text string, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent6(ptr.Pointer(), C.CString(text), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent2(text string, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent2(ptr.Pointer(), C.CString(text), C.int(qt.GoBoolToInt(namespaceProcessing)), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) ToByteArray(indent int) string {
	defer qt.Recovering("QDomDocument::toByteArray")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocument_ToByteArray(ptr.Pointer(), C.int(indent)))
	}
	return ""
}

func (ptr *QDomDocument) ToString(indent int) string {
	defer qt.Recovering("QDomDocument::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocument_ToString(ptr.Pointer(), C.int(indent)))
	}
	return ""
}

func (ptr *QDomDocument) DestroyQDomDocument() {
	defer qt.Recovering("QDomDocument::~QDomDocument")

	if ptr.Pointer() != nil {
		C.QDomDocument_DestroyQDomDocument(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDomDocumentFragment struct {
	QDomNode
}

type QDomDocumentFragment_ITF interface {
	QDomNode_ITF
	QDomDocumentFragment_PTR() *QDomDocumentFragment
}

func (p *QDomDocumentFragment) QDomDocumentFragment_PTR() *QDomDocumentFragment {
	return p
}

func (p *QDomDocumentFragment) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDomNode_PTR().Pointer()
	}
	return nil
}

func (p *QDomDocumentFragment) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDomNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQDomDocumentFragment(ptr QDomDocumentFragment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomDocumentFragment_PTR().Pointer()
	}
	return nil
}

func NewQDomDocumentFragmentFromPointer(ptr unsafe.Pointer) *QDomDocumentFragment {
	var n = new(QDomDocumentFragment)
	n.SetPointer(ptr)
	return n
}

func newQDomDocumentFragmentFromPointer(ptr unsafe.Pointer) *QDomDocumentFragment {
	var n = NewQDomDocumentFragmentFromPointer(ptr)
	return n
}

func NewQDomDocumentFragment() *QDomDocumentFragment {
	defer qt.Recovering("QDomDocumentFragment::QDomDocumentFragment")

	return newQDomDocumentFragmentFromPointer(C.QDomDocumentFragment_NewQDomDocumentFragment())
}

func NewQDomDocumentFragment2(x QDomDocumentFragment_ITF) *QDomDocumentFragment {
	defer qt.Recovering("QDomDocumentFragment::QDomDocumentFragment")

	return newQDomDocumentFragmentFromPointer(C.QDomDocumentFragment_NewQDomDocumentFragment2(PointerFromQDomDocumentFragment(x)))
}

func (ptr *QDomDocumentFragment) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomDocumentFragment::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomDocumentFragment_NodeType(ptr.Pointer()))
	}
	return 0
}

type QDomDocumentType struct {
	QDomNode
}

type QDomDocumentType_ITF interface {
	QDomNode_ITF
	QDomDocumentType_PTR() *QDomDocumentType
}

func (p *QDomDocumentType) QDomDocumentType_PTR() *QDomDocumentType {
	return p
}

func (p *QDomDocumentType) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDomNode_PTR().Pointer()
	}
	return nil
}

func (p *QDomDocumentType) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDomNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQDomDocumentType(ptr QDomDocumentType_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomDocumentType_PTR().Pointer()
	}
	return nil
}

func NewQDomDocumentTypeFromPointer(ptr unsafe.Pointer) *QDomDocumentType {
	var n = new(QDomDocumentType)
	n.SetPointer(ptr)
	return n
}

func newQDomDocumentTypeFromPointer(ptr unsafe.Pointer) *QDomDocumentType {
	var n = NewQDomDocumentTypeFromPointer(ptr)
	return n
}

func NewQDomDocumentType() *QDomDocumentType {
	defer qt.Recovering("QDomDocumentType::QDomDocumentType")

	return newQDomDocumentTypeFromPointer(C.QDomDocumentType_NewQDomDocumentType())
}

func NewQDomDocumentType2(n QDomDocumentType_ITF) *QDomDocumentType {
	defer qt.Recovering("QDomDocumentType::QDomDocumentType")

	return newQDomDocumentTypeFromPointer(C.QDomDocumentType_NewQDomDocumentType2(PointerFromQDomDocumentType(n)))
}

func (ptr *QDomDocumentType) Entities() *QDomNamedNodeMap {
	defer qt.Recovering("QDomDocumentType::entities")

	if ptr.Pointer() != nil {
		return NewQDomNamedNodeMapFromPointer(C.QDomDocumentType_Entities(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomDocumentType) InternalSubset() string {
	defer qt.Recovering("QDomDocumentType::internalSubset")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_InternalSubset(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomDocumentType) Name() string {
	defer qt.Recovering("QDomDocumentType::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomDocumentType) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomDocumentType::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomDocumentType_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomDocumentType) Notations() *QDomNamedNodeMap {
	defer qt.Recovering("QDomDocumentType::notations")

	if ptr.Pointer() != nil {
		return NewQDomNamedNodeMapFromPointer(C.QDomDocumentType_Notations(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomDocumentType) PublicId() string {
	defer qt.Recovering("QDomDocumentType::publicId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomDocumentType) SystemId() string {
	defer qt.Recovering("QDomDocumentType::systemId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_SystemId(ptr.Pointer()))
	}
	return ""
}

type QDomElement struct {
	QDomNode
}

type QDomElement_ITF interface {
	QDomNode_ITF
	QDomElement_PTR() *QDomElement
}

func (p *QDomElement) QDomElement_PTR() *QDomElement {
	return p
}

func (p *QDomElement) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDomNode_PTR().Pointer()
	}
	return nil
}

func (p *QDomElement) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDomNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQDomElement(ptr QDomElement_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomElement_PTR().Pointer()
	}
	return nil
}

func NewQDomElementFromPointer(ptr unsafe.Pointer) *QDomElement {
	var n = new(QDomElement)
	n.SetPointer(ptr)
	return n
}

func newQDomElementFromPointer(ptr unsafe.Pointer) *QDomElement {
	var n = NewQDomElementFromPointer(ptr)
	return n
}

func NewQDomElement() *QDomElement {
	defer qt.Recovering("QDomElement::QDomElement")

	return newQDomElementFromPointer(C.QDomElement_NewQDomElement())
}

func NewQDomElement2(x QDomElement_ITF) *QDomElement {
	defer qt.Recovering("QDomElement::QDomElement")

	return newQDomElementFromPointer(C.QDomElement_NewQDomElement2(PointerFromQDomElement(x)))
}

func (ptr *QDomElement) Attribute(name string, defValue string) string {
	defer qt.Recovering("QDomElement::attribute")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomElement_Attribute(ptr.Pointer(), C.CString(name), C.CString(defValue)))
	}
	return ""
}

func (ptr *QDomElement) AttributeNS(nsURI string, localName string, defValue string) string {
	defer qt.Recovering("QDomElement::attributeNS")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomElement_AttributeNS(ptr.Pointer(), C.CString(nsURI), C.CString(localName), C.CString(defValue)))
	}
	return ""
}

func (ptr *QDomElement) AttributeNode(name string) *QDomAttr {
	defer qt.Recovering("QDomElement::attributeNode")

	if ptr.Pointer() != nil {
		return NewQDomAttrFromPointer(C.QDomElement_AttributeNode(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QDomElement) AttributeNodeNS(nsURI string, localName string) *QDomAttr {
	defer qt.Recovering("QDomElement::attributeNodeNS")

	if ptr.Pointer() != nil {
		return NewQDomAttrFromPointer(C.QDomElement_AttributeNodeNS(ptr.Pointer(), C.CString(nsURI), C.CString(localName)))
	}
	return nil
}

func (ptr *QDomElement) Attributes() *QDomNamedNodeMap {
	defer qt.Recovering("QDomElement::attributes")

	if ptr.Pointer() != nil {
		return NewQDomNamedNodeMapFromPointer(C.QDomElement_Attributes(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomElement) ElementsByTagName(tagname string) *QDomNodeList {
	defer qt.Recovering("QDomElement::elementsByTagName")

	if ptr.Pointer() != nil {
		return NewQDomNodeListFromPointer(C.QDomElement_ElementsByTagName(ptr.Pointer(), C.CString(tagname)))
	}
	return nil
}

func (ptr *QDomElement) ElementsByTagNameNS(nsURI string, localName string) *QDomNodeList {
	defer qt.Recovering("QDomElement::elementsByTagNameNS")

	if ptr.Pointer() != nil {
		return NewQDomNodeListFromPointer(C.QDomElement_ElementsByTagNameNS(ptr.Pointer(), C.CString(nsURI), C.CString(localName)))
	}
	return nil
}

func (ptr *QDomElement) HasAttribute(name string) bool {
	defer qt.Recovering("QDomElement::hasAttribute")

	if ptr.Pointer() != nil {
		return C.QDomElement_HasAttribute(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QDomElement) HasAttributeNS(nsURI string, localName string) bool {
	defer qt.Recovering("QDomElement::hasAttributeNS")

	if ptr.Pointer() != nil {
		return C.QDomElement_HasAttributeNS(ptr.Pointer(), C.CString(nsURI), C.CString(localName)) != 0
	}
	return false
}

func (ptr *QDomElement) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomElement::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomElement_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomElement) RemoveAttribute(name string) {
	defer qt.Recovering("QDomElement::removeAttribute")

	if ptr.Pointer() != nil {
		C.QDomElement_RemoveAttribute(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QDomElement) RemoveAttributeNS(nsURI string, localName string) {
	defer qt.Recovering("QDomElement::removeAttributeNS")

	if ptr.Pointer() != nil {
		C.QDomElement_RemoveAttributeNS(ptr.Pointer(), C.CString(nsURI), C.CString(localName))
	}
}

func (ptr *QDomElement) RemoveAttributeNode(oldAttr QDomAttr_ITF) *QDomAttr {
	defer qt.Recovering("QDomElement::removeAttributeNode")

	if ptr.Pointer() != nil {
		return NewQDomAttrFromPointer(C.QDomElement_RemoveAttributeNode(ptr.Pointer(), PointerFromQDomAttr(oldAttr)))
	}
	return nil
}

func (ptr *QDomElement) SetAttribute(name string, value string) {
	defer qt.Recovering("QDomElement::setAttribute")

	if ptr.Pointer() != nil {
		C.QDomElement_SetAttribute(ptr.Pointer(), C.CString(name), C.CString(value))
	}
}

func (ptr *QDomElement) SetAttribute4(name string, value int) {
	defer qt.Recovering("QDomElement::setAttribute")

	if ptr.Pointer() != nil {
		C.QDomElement_SetAttribute4(ptr.Pointer(), C.CString(name), C.int(value))
	}
}

func (ptr *QDomElement) SetAttributeNS(nsURI string, qName string, value string) {
	defer qt.Recovering("QDomElement::setAttributeNS")

	if ptr.Pointer() != nil {
		C.QDomElement_SetAttributeNS(ptr.Pointer(), C.CString(nsURI), C.CString(qName), C.CString(value))
	}
}

func (ptr *QDomElement) SetAttributeNS2(nsURI string, qName string, value int) {
	defer qt.Recovering("QDomElement::setAttributeNS")

	if ptr.Pointer() != nil {
		C.QDomElement_SetAttributeNS2(ptr.Pointer(), C.CString(nsURI), C.CString(qName), C.int(value))
	}
}

func (ptr *QDomElement) SetAttributeNode(newAttr QDomAttr_ITF) *QDomAttr {
	defer qt.Recovering("QDomElement::setAttributeNode")

	if ptr.Pointer() != nil {
		return NewQDomAttrFromPointer(C.QDomElement_SetAttributeNode(ptr.Pointer(), PointerFromQDomAttr(newAttr)))
	}
	return nil
}

func (ptr *QDomElement) SetAttributeNodeNS(newAttr QDomAttr_ITF) *QDomAttr {
	defer qt.Recovering("QDomElement::setAttributeNodeNS")

	if ptr.Pointer() != nil {
		return NewQDomAttrFromPointer(C.QDomElement_SetAttributeNodeNS(ptr.Pointer(), PointerFromQDomAttr(newAttr)))
	}
	return nil
}

func (ptr *QDomElement) SetTagName(name string) {
	defer qt.Recovering("QDomElement::setTagName")

	if ptr.Pointer() != nil {
		C.QDomElement_SetTagName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QDomElement) TagName() string {
	defer qt.Recovering("QDomElement::tagName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomElement_TagName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomElement) Text() string {
	defer qt.Recovering("QDomElement::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomElement_Text(ptr.Pointer()))
	}
	return ""
}

type QDomEntity struct {
	QDomNode
}

type QDomEntity_ITF interface {
	QDomNode_ITF
	QDomEntity_PTR() *QDomEntity
}

func (p *QDomEntity) QDomEntity_PTR() *QDomEntity {
	return p
}

func (p *QDomEntity) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDomNode_PTR().Pointer()
	}
	return nil
}

func (p *QDomEntity) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDomNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQDomEntity(ptr QDomEntity_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomEntity_PTR().Pointer()
	}
	return nil
}

func NewQDomEntityFromPointer(ptr unsafe.Pointer) *QDomEntity {
	var n = new(QDomEntity)
	n.SetPointer(ptr)
	return n
}

func newQDomEntityFromPointer(ptr unsafe.Pointer) *QDomEntity {
	var n = NewQDomEntityFromPointer(ptr)
	return n
}

func NewQDomEntity() *QDomEntity {
	defer qt.Recovering("QDomEntity::QDomEntity")

	return newQDomEntityFromPointer(C.QDomEntity_NewQDomEntity())
}

func NewQDomEntity2(x QDomEntity_ITF) *QDomEntity {
	defer qt.Recovering("QDomEntity::QDomEntity")

	return newQDomEntityFromPointer(C.QDomEntity_NewQDomEntity2(PointerFromQDomEntity(x)))
}

func (ptr *QDomEntity) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomEntity::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomEntity_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomEntity) NotationName() string {
	defer qt.Recovering("QDomEntity::notationName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomEntity_NotationName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomEntity) PublicId() string {
	defer qt.Recovering("QDomEntity::publicId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomEntity_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomEntity) SystemId() string {
	defer qt.Recovering("QDomEntity::systemId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomEntity_SystemId(ptr.Pointer()))
	}
	return ""
}

type QDomEntityReference struct {
	QDomNode
}

type QDomEntityReference_ITF interface {
	QDomNode_ITF
	QDomEntityReference_PTR() *QDomEntityReference
}

func (p *QDomEntityReference) QDomEntityReference_PTR() *QDomEntityReference {
	return p
}

func (p *QDomEntityReference) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDomNode_PTR().Pointer()
	}
	return nil
}

func (p *QDomEntityReference) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDomNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQDomEntityReference(ptr QDomEntityReference_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomEntityReference_PTR().Pointer()
	}
	return nil
}

func NewQDomEntityReferenceFromPointer(ptr unsafe.Pointer) *QDomEntityReference {
	var n = new(QDomEntityReference)
	n.SetPointer(ptr)
	return n
}

func newQDomEntityReferenceFromPointer(ptr unsafe.Pointer) *QDomEntityReference {
	var n = NewQDomEntityReferenceFromPointer(ptr)
	return n
}

func NewQDomEntityReference() *QDomEntityReference {
	defer qt.Recovering("QDomEntityReference::QDomEntityReference")

	return newQDomEntityReferenceFromPointer(C.QDomEntityReference_NewQDomEntityReference())
}

func NewQDomEntityReference2(x QDomEntityReference_ITF) *QDomEntityReference {
	defer qt.Recovering("QDomEntityReference::QDomEntityReference")

	return newQDomEntityReferenceFromPointer(C.QDomEntityReference_NewQDomEntityReference2(PointerFromQDomEntityReference(x)))
}

func (ptr *QDomEntityReference) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomEntityReference::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomEntityReference_NodeType(ptr.Pointer()))
	}
	return 0
}

//QDomImplementation::InvalidDataPolicy
type QDomImplementation__InvalidDataPolicy int64

const (
	QDomImplementation__AcceptInvalidChars = QDomImplementation__InvalidDataPolicy(0)
	QDomImplementation__DropInvalidChars   = QDomImplementation__InvalidDataPolicy(1)
	QDomImplementation__ReturnNullNode     = QDomImplementation__InvalidDataPolicy(2)
)

type QDomImplementation struct {
	ptr unsafe.Pointer
}

type QDomImplementation_ITF interface {
	QDomImplementation_PTR() *QDomImplementation
}

func (p *QDomImplementation) QDomImplementation_PTR() *QDomImplementation {
	return p
}

func (p *QDomImplementation) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDomImplementation) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDomImplementation(ptr QDomImplementation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomImplementation_PTR().Pointer()
	}
	return nil
}

func NewQDomImplementationFromPointer(ptr unsafe.Pointer) *QDomImplementation {
	var n = new(QDomImplementation)
	n.SetPointer(ptr)
	return n
}

func newQDomImplementationFromPointer(ptr unsafe.Pointer) *QDomImplementation {
	var n = NewQDomImplementationFromPointer(ptr)
	return n
}

func NewQDomImplementation() *QDomImplementation {
	defer qt.Recovering("QDomImplementation::QDomImplementation")

	return newQDomImplementationFromPointer(C.QDomImplementation_NewQDomImplementation())
}

func NewQDomImplementation2(x QDomImplementation_ITF) *QDomImplementation {
	defer qt.Recovering("QDomImplementation::QDomImplementation")

	return newQDomImplementationFromPointer(C.QDomImplementation_NewQDomImplementation2(PointerFromQDomImplementation(x)))
}

func (ptr *QDomImplementation) CreateDocument(nsURI string, qName string, doctype QDomDocumentType_ITF) *QDomDocument {
	defer qt.Recovering("QDomImplementation::createDocument")

	if ptr.Pointer() != nil {
		return NewQDomDocumentFromPointer(C.QDomImplementation_CreateDocument(ptr.Pointer(), C.CString(nsURI), C.CString(qName), PointerFromQDomDocumentType(doctype)))
	}
	return nil
}

func (ptr *QDomImplementation) CreateDocumentType(qName string, publicId string, systemId string) *QDomDocumentType {
	defer qt.Recovering("QDomImplementation::createDocumentType")

	if ptr.Pointer() != nil {
		return NewQDomDocumentTypeFromPointer(C.QDomImplementation_CreateDocumentType(ptr.Pointer(), C.CString(qName), C.CString(publicId), C.CString(systemId)))
	}
	return nil
}

func (ptr *QDomImplementation) HasFeature(feature string, version string) bool {
	defer qt.Recovering("QDomImplementation::hasFeature")

	if ptr.Pointer() != nil {
		return C.QDomImplementation_HasFeature(ptr.Pointer(), C.CString(feature), C.CString(version)) != 0
	}
	return false
}

func QDomImplementation_InvalidDataPolicy() QDomImplementation__InvalidDataPolicy {
	defer qt.Recovering("QDomImplementation::invalidDataPolicy")

	return QDomImplementation__InvalidDataPolicy(C.QDomImplementation_QDomImplementation_InvalidDataPolicy())
}

func (ptr *QDomImplementation) InvalidDataPolicy() QDomImplementation__InvalidDataPolicy {
	defer qt.Recovering("QDomImplementation::invalidDataPolicy")

	return QDomImplementation__InvalidDataPolicy(C.QDomImplementation_QDomImplementation_InvalidDataPolicy())
}

func (ptr *QDomImplementation) IsNull() bool {
	defer qt.Recovering("QDomImplementation::isNull")

	if ptr.Pointer() != nil {
		return C.QDomImplementation_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func QDomImplementation_SetInvalidDataPolicy(policy QDomImplementation__InvalidDataPolicy) {
	defer qt.Recovering("QDomImplementation::setInvalidDataPolicy")

	C.QDomImplementation_QDomImplementation_SetInvalidDataPolicy(C.int(policy))
}

func (ptr *QDomImplementation) SetInvalidDataPolicy(policy QDomImplementation__InvalidDataPolicy) {
	defer qt.Recovering("QDomImplementation::setInvalidDataPolicy")

	C.QDomImplementation_QDomImplementation_SetInvalidDataPolicy(C.int(policy))
}

func (ptr *QDomImplementation) DestroyQDomImplementation() {
	defer qt.Recovering("QDomImplementation::~QDomImplementation")

	if ptr.Pointer() != nil {
		C.QDomImplementation_DestroyQDomImplementation(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDomNamedNodeMap struct {
	ptr unsafe.Pointer
}

type QDomNamedNodeMap_ITF interface {
	QDomNamedNodeMap_PTR() *QDomNamedNodeMap
}

func (p *QDomNamedNodeMap) QDomNamedNodeMap_PTR() *QDomNamedNodeMap {
	return p
}

func (p *QDomNamedNodeMap) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDomNamedNodeMap) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDomNamedNodeMap(ptr QDomNamedNodeMap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNamedNodeMap_PTR().Pointer()
	}
	return nil
}

func NewQDomNamedNodeMapFromPointer(ptr unsafe.Pointer) *QDomNamedNodeMap {
	var n = new(QDomNamedNodeMap)
	n.SetPointer(ptr)
	return n
}

func newQDomNamedNodeMapFromPointer(ptr unsafe.Pointer) *QDomNamedNodeMap {
	var n = NewQDomNamedNodeMapFromPointer(ptr)
	return n
}

func NewQDomNamedNodeMap() *QDomNamedNodeMap {
	defer qt.Recovering("QDomNamedNodeMap::QDomNamedNodeMap")

	return newQDomNamedNodeMapFromPointer(C.QDomNamedNodeMap_NewQDomNamedNodeMap())
}

func NewQDomNamedNodeMap2(n QDomNamedNodeMap_ITF) *QDomNamedNodeMap {
	defer qt.Recovering("QDomNamedNodeMap::QDomNamedNodeMap")

	return newQDomNamedNodeMapFromPointer(C.QDomNamedNodeMap_NewQDomNamedNodeMap2(PointerFromQDomNamedNodeMap(n)))
}

func (ptr *QDomNamedNodeMap) Contains(name string) bool {
	defer qt.Recovering("QDomNamedNodeMap::contains")

	if ptr.Pointer() != nil {
		return C.QDomNamedNodeMap_Contains(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QDomNamedNodeMap) Count() int {
	defer qt.Recovering("QDomNamedNodeMap::count")

	if ptr.Pointer() != nil {
		return int(C.QDomNamedNodeMap_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNamedNodeMap) IsEmpty() bool {
	defer qt.Recovering("QDomNamedNodeMap::isEmpty")

	if ptr.Pointer() != nil {
		return C.QDomNamedNodeMap_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNamedNodeMap) Item(index int) *QDomNode {
	defer qt.Recovering("QDomNamedNodeMap::item")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNamedNodeMap_Item(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QDomNamedNodeMap) Length() int {
	defer qt.Recovering("QDomNamedNodeMap::length")

	if ptr.Pointer() != nil {
		return int(C.QDomNamedNodeMap_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNamedNodeMap) NamedItem(name string) *QDomNode {
	defer qt.Recovering("QDomNamedNodeMap::namedItem")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNamedNodeMap_NamedItem(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QDomNamedNodeMap) NamedItemNS(nsURI string, localName string) *QDomNode {
	defer qt.Recovering("QDomNamedNodeMap::namedItemNS")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNamedNodeMap_NamedItemNS(ptr.Pointer(), C.CString(nsURI), C.CString(localName)))
	}
	return nil
}

func (ptr *QDomNamedNodeMap) RemoveNamedItem(name string) *QDomNode {
	defer qt.Recovering("QDomNamedNodeMap::removeNamedItem")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNamedNodeMap_RemoveNamedItem(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QDomNamedNodeMap) RemoveNamedItemNS(nsURI string, localName string) *QDomNode {
	defer qt.Recovering("QDomNamedNodeMap::removeNamedItemNS")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNamedNodeMap_RemoveNamedItemNS(ptr.Pointer(), C.CString(nsURI), C.CString(localName)))
	}
	return nil
}

func (ptr *QDomNamedNodeMap) SetNamedItem(newNode QDomNode_ITF) *QDomNode {
	defer qt.Recovering("QDomNamedNodeMap::setNamedItem")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNamedNodeMap_SetNamedItem(ptr.Pointer(), PointerFromQDomNode(newNode)))
	}
	return nil
}

func (ptr *QDomNamedNodeMap) SetNamedItemNS(newNode QDomNode_ITF) *QDomNode {
	defer qt.Recovering("QDomNamedNodeMap::setNamedItemNS")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNamedNodeMap_SetNamedItemNS(ptr.Pointer(), PointerFromQDomNode(newNode)))
	}
	return nil
}

func (ptr *QDomNamedNodeMap) Size() int {
	defer qt.Recovering("QDomNamedNodeMap::size")

	if ptr.Pointer() != nil {
		return int(C.QDomNamedNodeMap_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNamedNodeMap) DestroyQDomNamedNodeMap() {
	defer qt.Recovering("QDomNamedNodeMap::~QDomNamedNodeMap")

	if ptr.Pointer() != nil {
		C.QDomNamedNodeMap_DestroyQDomNamedNodeMap(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QDomNode::EncodingPolicy
type QDomNode__EncodingPolicy int64

const (
	QDomNode__EncodingFromDocument   = QDomNode__EncodingPolicy(1)
	QDomNode__EncodingFromTextStream = QDomNode__EncodingPolicy(2)
)

//QDomNode::NodeType
type QDomNode__NodeType int64

const (
	QDomNode__ElementNode               = QDomNode__NodeType(1)
	QDomNode__AttributeNode             = QDomNode__NodeType(2)
	QDomNode__TextNode                  = QDomNode__NodeType(3)
	QDomNode__CDATASectionNode          = QDomNode__NodeType(4)
	QDomNode__EntityReferenceNode       = QDomNode__NodeType(5)
	QDomNode__EntityNode                = QDomNode__NodeType(6)
	QDomNode__ProcessingInstructionNode = QDomNode__NodeType(7)
	QDomNode__CommentNode               = QDomNode__NodeType(8)
	QDomNode__DocumentNode              = QDomNode__NodeType(9)
	QDomNode__DocumentTypeNode          = QDomNode__NodeType(10)
	QDomNode__DocumentFragmentNode      = QDomNode__NodeType(11)
	QDomNode__NotationNode              = QDomNode__NodeType(12)
	QDomNode__BaseNode                  = QDomNode__NodeType(21)
	QDomNode__CharacterDataNode         = QDomNode__NodeType(22)
)

type QDomNode struct {
	ptr unsafe.Pointer
}

type QDomNode_ITF interface {
	QDomNode_PTR() *QDomNode
}

func (p *QDomNode) QDomNode_PTR() *QDomNode {
	return p
}

func (p *QDomNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDomNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDomNode(ptr QDomNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNode_PTR().Pointer()
	}
	return nil
}

func NewQDomNodeFromPointer(ptr unsafe.Pointer) *QDomNode {
	var n = new(QDomNode)
	n.SetPointer(ptr)
	return n
}

func newQDomNodeFromPointer(ptr unsafe.Pointer) *QDomNode {
	var n = NewQDomNodeFromPointer(ptr)
	return n
}

func NewQDomNode() *QDomNode {
	defer qt.Recovering("QDomNode::QDomNode")

	return newQDomNodeFromPointer(C.QDomNode_NewQDomNode())
}

func NewQDomNode2(n QDomNode_ITF) *QDomNode {
	defer qt.Recovering("QDomNode::QDomNode")

	return newQDomNodeFromPointer(C.QDomNode_NewQDomNode2(PointerFromQDomNode(n)))
}

func (ptr *QDomNode) AppendChild(newChild QDomNode_ITF) *QDomNode {
	defer qt.Recovering("QDomNode::appendChild")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNode_AppendChild(ptr.Pointer(), PointerFromQDomNode(newChild)))
	}
	return nil
}

func (ptr *QDomNode) Attributes() *QDomNamedNodeMap {
	defer qt.Recovering("QDomNode::attributes")

	if ptr.Pointer() != nil {
		return NewQDomNamedNodeMapFromPointer(C.QDomNode_Attributes(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) ChildNodes() *QDomNodeList {
	defer qt.Recovering("QDomNode::childNodes")

	if ptr.Pointer() != nil {
		return NewQDomNodeListFromPointer(C.QDomNode_ChildNodes(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) Clear() {
	defer qt.Recovering("QDomNode::clear")

	if ptr.Pointer() != nil {
		C.QDomNode_Clear(ptr.Pointer())
	}
}

func (ptr *QDomNode) CloneNode(deep bool) *QDomNode {
	defer qt.Recovering("QDomNode::cloneNode")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNode_CloneNode(ptr.Pointer(), C.int(qt.GoBoolToInt(deep))))
	}
	return nil
}

func (ptr *QDomNode) ColumnNumber() int {
	defer qt.Recovering("QDomNode::columnNumber")

	if ptr.Pointer() != nil {
		return int(C.QDomNode_ColumnNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNode) FirstChild() *QDomNode {
	defer qt.Recovering("QDomNode::firstChild")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNode_FirstChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) FirstChildElement(tagName string) *QDomElement {
	defer qt.Recovering("QDomNode::firstChildElement")

	if ptr.Pointer() != nil {
		return NewQDomElementFromPointer(C.QDomNode_FirstChildElement(ptr.Pointer(), C.CString(tagName)))
	}
	return nil
}

func (ptr *QDomNode) HasAttributes() bool {
	defer qt.Recovering("QDomNode::hasAttributes")

	if ptr.Pointer() != nil {
		return C.QDomNode_HasAttributes(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) HasChildNodes() bool {
	defer qt.Recovering("QDomNode::hasChildNodes")

	if ptr.Pointer() != nil {
		return C.QDomNode_HasChildNodes(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) InsertAfter(newChild QDomNode_ITF, refChild QDomNode_ITF) *QDomNode {
	defer qt.Recovering("QDomNode::insertAfter")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNode_InsertAfter(ptr.Pointer(), PointerFromQDomNode(newChild), PointerFromQDomNode(refChild)))
	}
	return nil
}

func (ptr *QDomNode) InsertBefore(newChild QDomNode_ITF, refChild QDomNode_ITF) *QDomNode {
	defer qt.Recovering("QDomNode::insertBefore")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNode_InsertBefore(ptr.Pointer(), PointerFromQDomNode(newChild), PointerFromQDomNode(refChild)))
	}
	return nil
}

func (ptr *QDomNode) IsAttr() bool {
	defer qt.Recovering("QDomNode::isAttr")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsAttr(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsCDATASection() bool {
	defer qt.Recovering("QDomNode::isCDATASection")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsCDATASection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsCharacterData() bool {
	defer qt.Recovering("QDomNode::isCharacterData")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsCharacterData(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsComment() bool {
	defer qt.Recovering("QDomNode::isComment")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsComment(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsDocument() bool {
	defer qt.Recovering("QDomNode::isDocument")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsDocumentFragment() bool {
	defer qt.Recovering("QDomNode::isDocumentFragment")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsDocumentFragment(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsDocumentType() bool {
	defer qt.Recovering("QDomNode::isDocumentType")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsDocumentType(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsElement() bool {
	defer qt.Recovering("QDomNode::isElement")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsElement(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsEntity() bool {
	defer qt.Recovering("QDomNode::isEntity")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsEntity(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsEntityReference() bool {
	defer qt.Recovering("QDomNode::isEntityReference")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsEntityReference(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsNotation() bool {
	defer qt.Recovering("QDomNode::isNotation")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsNotation(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsNull() bool {
	defer qt.Recovering("QDomNode::isNull")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsProcessingInstruction() bool {
	defer qt.Recovering("QDomNode::isProcessingInstruction")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsProcessingInstruction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsSupported(feature string, version string) bool {
	defer qt.Recovering("QDomNode::isSupported")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsSupported(ptr.Pointer(), C.CString(feature), C.CString(version)) != 0
	}
	return false
}

func (ptr *QDomNode) IsText() bool {
	defer qt.Recovering("QDomNode::isText")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsText(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) LastChild() *QDomNode {
	defer qt.Recovering("QDomNode::lastChild")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNode_LastChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) LastChildElement(tagName string) *QDomElement {
	defer qt.Recovering("QDomNode::lastChildElement")

	if ptr.Pointer() != nil {
		return NewQDomElementFromPointer(C.QDomNode_LastChildElement(ptr.Pointer(), C.CString(tagName)))
	}
	return nil
}

func (ptr *QDomNode) LineNumber() int {
	defer qt.Recovering("QDomNode::lineNumber")

	if ptr.Pointer() != nil {
		return int(C.QDomNode_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNode) LocalName() string {
	defer qt.Recovering("QDomNode::localName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNode_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) NamedItem(name string) *QDomNode {
	defer qt.Recovering("QDomNode::namedItem")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNode_NamedItem(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QDomNode) NamespaceURI() string {
	defer qt.Recovering("QDomNode::namespaceURI")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNode_NamespaceURI(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) NextSibling() *QDomNode {
	defer qt.Recovering("QDomNode::nextSibling")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNode_NextSibling(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) NextSiblingElement(tagName string) *QDomElement {
	defer qt.Recovering("QDomNode::nextSiblingElement")

	if ptr.Pointer() != nil {
		return NewQDomElementFromPointer(C.QDomNode_NextSiblingElement(ptr.Pointer(), C.CString(tagName)))
	}
	return nil
}

func (ptr *QDomNode) NodeName() string {
	defer qt.Recovering("QDomNode::nodeName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNode_NodeName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomNode::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomNode_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNode) NodeValue() string {
	defer qt.Recovering("QDomNode::nodeValue")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNode_NodeValue(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) Normalize() {
	defer qt.Recovering("QDomNode::normalize")

	if ptr.Pointer() != nil {
		C.QDomNode_Normalize(ptr.Pointer())
	}
}

func (ptr *QDomNode) OwnerDocument() *QDomDocument {
	defer qt.Recovering("QDomNode::ownerDocument")

	if ptr.Pointer() != nil {
		return NewQDomDocumentFromPointer(C.QDomNode_OwnerDocument(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) ParentNode() *QDomNode {
	defer qt.Recovering("QDomNode::parentNode")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNode_ParentNode(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) Prefix() string {
	defer qt.Recovering("QDomNode::prefix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNode_Prefix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) PreviousSibling() *QDomNode {
	defer qt.Recovering("QDomNode::previousSibling")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNode_PreviousSibling(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) PreviousSiblingElement(tagName string) *QDomElement {
	defer qt.Recovering("QDomNode::previousSiblingElement")

	if ptr.Pointer() != nil {
		return NewQDomElementFromPointer(C.QDomNode_PreviousSiblingElement(ptr.Pointer(), C.CString(tagName)))
	}
	return nil
}

func (ptr *QDomNode) RemoveChild(oldChild QDomNode_ITF) *QDomNode {
	defer qt.Recovering("QDomNode::removeChild")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNode_RemoveChild(ptr.Pointer(), PointerFromQDomNode(oldChild)))
	}
	return nil
}

func (ptr *QDomNode) ReplaceChild(newChild QDomNode_ITF, oldChild QDomNode_ITF) *QDomNode {
	defer qt.Recovering("QDomNode::replaceChild")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNode_ReplaceChild(ptr.Pointer(), PointerFromQDomNode(newChild), PointerFromQDomNode(oldChild)))
	}
	return nil
}

func (ptr *QDomNode) Save(stream core.QTextStream_ITF, indent int, encodingPolicy QDomNode__EncodingPolicy) {
	defer qt.Recovering("QDomNode::save")

	if ptr.Pointer() != nil {
		C.QDomNode_Save(ptr.Pointer(), core.PointerFromQTextStream(stream), C.int(indent), C.int(encodingPolicy))
	}
}

func (ptr *QDomNode) SetNodeValue(v string) {
	defer qt.Recovering("QDomNode::setNodeValue")

	if ptr.Pointer() != nil {
		C.QDomNode_SetNodeValue(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QDomNode) SetPrefix(pre string) {
	defer qt.Recovering("QDomNode::setPrefix")

	if ptr.Pointer() != nil {
		C.QDomNode_SetPrefix(ptr.Pointer(), C.CString(pre))
	}
}

func (ptr *QDomNode) ToAttr() *QDomAttr {
	defer qt.Recovering("QDomNode::toAttr")

	if ptr.Pointer() != nil {
		return NewQDomAttrFromPointer(C.QDomNode_ToAttr(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) ToCDATASection() *QDomCDATASection {
	defer qt.Recovering("QDomNode::toCDATASection")

	if ptr.Pointer() != nil {
		return NewQDomCDATASectionFromPointer(C.QDomNode_ToCDATASection(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) ToCharacterData() *QDomCharacterData {
	defer qt.Recovering("QDomNode::toCharacterData")

	if ptr.Pointer() != nil {
		return NewQDomCharacterDataFromPointer(C.QDomNode_ToCharacterData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) ToComment() *QDomComment {
	defer qt.Recovering("QDomNode::toComment")

	if ptr.Pointer() != nil {
		return NewQDomCommentFromPointer(C.QDomNode_ToComment(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) ToDocument() *QDomDocument {
	defer qt.Recovering("QDomNode::toDocument")

	if ptr.Pointer() != nil {
		return NewQDomDocumentFromPointer(C.QDomNode_ToDocument(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) ToDocumentFragment() *QDomDocumentFragment {
	defer qt.Recovering("QDomNode::toDocumentFragment")

	if ptr.Pointer() != nil {
		return NewQDomDocumentFragmentFromPointer(C.QDomNode_ToDocumentFragment(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) ToDocumentType() *QDomDocumentType {
	defer qt.Recovering("QDomNode::toDocumentType")

	if ptr.Pointer() != nil {
		return NewQDomDocumentTypeFromPointer(C.QDomNode_ToDocumentType(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) ToElement() *QDomElement {
	defer qt.Recovering("QDomNode::toElement")

	if ptr.Pointer() != nil {
		return NewQDomElementFromPointer(C.QDomNode_ToElement(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) ToEntity() *QDomEntity {
	defer qt.Recovering("QDomNode::toEntity")

	if ptr.Pointer() != nil {
		return NewQDomEntityFromPointer(C.QDomNode_ToEntity(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) ToEntityReference() *QDomEntityReference {
	defer qt.Recovering("QDomNode::toEntityReference")

	if ptr.Pointer() != nil {
		return NewQDomEntityReferenceFromPointer(C.QDomNode_ToEntityReference(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) ToNotation() *QDomNotation {
	defer qt.Recovering("QDomNode::toNotation")

	if ptr.Pointer() != nil {
		return NewQDomNotationFromPointer(C.QDomNode_ToNotation(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) ToProcessingInstruction() *QDomProcessingInstruction {
	defer qt.Recovering("QDomNode::toProcessingInstruction")

	if ptr.Pointer() != nil {
		return NewQDomProcessingInstructionFromPointer(C.QDomNode_ToProcessingInstruction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) ToText() *QDomText {
	defer qt.Recovering("QDomNode::toText")

	if ptr.Pointer() != nil {
		return NewQDomTextFromPointer(C.QDomNode_ToText(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDomNode) DestroyQDomNode() {
	defer qt.Recovering("QDomNode::~QDomNode")

	if ptr.Pointer() != nil {
		C.QDomNode_DestroyQDomNode(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDomNodeList struct {
	ptr unsafe.Pointer
}

type QDomNodeList_ITF interface {
	QDomNodeList_PTR() *QDomNodeList
}

func (p *QDomNodeList) QDomNodeList_PTR() *QDomNodeList {
	return p
}

func (p *QDomNodeList) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDomNodeList) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDomNodeList(ptr QDomNodeList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNodeList_PTR().Pointer()
	}
	return nil
}

func NewQDomNodeListFromPointer(ptr unsafe.Pointer) *QDomNodeList {
	var n = new(QDomNodeList)
	n.SetPointer(ptr)
	return n
}

func newQDomNodeListFromPointer(ptr unsafe.Pointer) *QDomNodeList {
	var n = NewQDomNodeListFromPointer(ptr)
	return n
}

func NewQDomNodeList() *QDomNodeList {
	defer qt.Recovering("QDomNodeList::QDomNodeList")

	return newQDomNodeListFromPointer(C.QDomNodeList_NewQDomNodeList())
}

func NewQDomNodeList2(n QDomNodeList_ITF) *QDomNodeList {
	defer qt.Recovering("QDomNodeList::QDomNodeList")

	return newQDomNodeListFromPointer(C.QDomNodeList_NewQDomNodeList2(PointerFromQDomNodeList(n)))
}

func (ptr *QDomNodeList) At(index int) *QDomNode {
	defer qt.Recovering("QDomNodeList::at")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNodeList_At(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QDomNodeList) Count() int {
	defer qt.Recovering("QDomNodeList::count")

	if ptr.Pointer() != nil {
		return int(C.QDomNodeList_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNodeList) IsEmpty() bool {
	defer qt.Recovering("QDomNodeList::isEmpty")

	if ptr.Pointer() != nil {
		return C.QDomNodeList_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNodeList) Item(index int) *QDomNode {
	defer qt.Recovering("QDomNodeList::item")

	if ptr.Pointer() != nil {
		return NewQDomNodeFromPointer(C.QDomNodeList_Item(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QDomNodeList) Length() int {
	defer qt.Recovering("QDomNodeList::length")

	if ptr.Pointer() != nil {
		return int(C.QDomNodeList_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNodeList) Size() int {
	defer qt.Recovering("QDomNodeList::size")

	if ptr.Pointer() != nil {
		return int(C.QDomNodeList_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNodeList) DestroyQDomNodeList() {
	defer qt.Recovering("QDomNodeList::~QDomNodeList")

	if ptr.Pointer() != nil {
		C.QDomNodeList_DestroyQDomNodeList(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDomNotation struct {
	QDomNode
}

type QDomNotation_ITF interface {
	QDomNode_ITF
	QDomNotation_PTR() *QDomNotation
}

func (p *QDomNotation) QDomNotation_PTR() *QDomNotation {
	return p
}

func (p *QDomNotation) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDomNode_PTR().Pointer()
	}
	return nil
}

func (p *QDomNotation) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDomNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQDomNotation(ptr QDomNotation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNotation_PTR().Pointer()
	}
	return nil
}

func NewQDomNotationFromPointer(ptr unsafe.Pointer) *QDomNotation {
	var n = new(QDomNotation)
	n.SetPointer(ptr)
	return n
}

func newQDomNotationFromPointer(ptr unsafe.Pointer) *QDomNotation {
	var n = NewQDomNotationFromPointer(ptr)
	return n
}

func NewQDomNotation() *QDomNotation {
	defer qt.Recovering("QDomNotation::QDomNotation")

	return newQDomNotationFromPointer(C.QDomNotation_NewQDomNotation())
}

func NewQDomNotation2(x QDomNotation_ITF) *QDomNotation {
	defer qt.Recovering("QDomNotation::QDomNotation")

	return newQDomNotationFromPointer(C.QDomNotation_NewQDomNotation2(PointerFromQDomNotation(x)))
}

func (ptr *QDomNotation) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomNotation::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomNotation_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNotation) PublicId() string {
	defer qt.Recovering("QDomNotation::publicId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNotation_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNotation) SystemId() string {
	defer qt.Recovering("QDomNotation::systemId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNotation_SystemId(ptr.Pointer()))
	}
	return ""
}

type QDomProcessingInstruction struct {
	QDomNode
}

type QDomProcessingInstruction_ITF interface {
	QDomNode_ITF
	QDomProcessingInstruction_PTR() *QDomProcessingInstruction
}

func (p *QDomProcessingInstruction) QDomProcessingInstruction_PTR() *QDomProcessingInstruction {
	return p
}

func (p *QDomProcessingInstruction) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDomNode_PTR().Pointer()
	}
	return nil
}

func (p *QDomProcessingInstruction) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDomNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQDomProcessingInstruction(ptr QDomProcessingInstruction_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomProcessingInstruction_PTR().Pointer()
	}
	return nil
}

func NewQDomProcessingInstructionFromPointer(ptr unsafe.Pointer) *QDomProcessingInstruction {
	var n = new(QDomProcessingInstruction)
	n.SetPointer(ptr)
	return n
}

func newQDomProcessingInstructionFromPointer(ptr unsafe.Pointer) *QDomProcessingInstruction {
	var n = NewQDomProcessingInstructionFromPointer(ptr)
	return n
}

func NewQDomProcessingInstruction() *QDomProcessingInstruction {
	defer qt.Recovering("QDomProcessingInstruction::QDomProcessingInstruction")

	return newQDomProcessingInstructionFromPointer(C.QDomProcessingInstruction_NewQDomProcessingInstruction())
}

func NewQDomProcessingInstruction2(x QDomProcessingInstruction_ITF) *QDomProcessingInstruction {
	defer qt.Recovering("QDomProcessingInstruction::QDomProcessingInstruction")

	return newQDomProcessingInstructionFromPointer(C.QDomProcessingInstruction_NewQDomProcessingInstruction2(PointerFromQDomProcessingInstruction(x)))
}

func (ptr *QDomProcessingInstruction) Data() string {
	defer qt.Recovering("QDomProcessingInstruction::data")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomProcessingInstruction_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomProcessingInstruction) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomProcessingInstruction::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomProcessingInstruction_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomProcessingInstruction) SetData(d string) {
	defer qt.Recovering("QDomProcessingInstruction::setData")

	if ptr.Pointer() != nil {
		C.QDomProcessingInstruction_SetData(ptr.Pointer(), C.CString(d))
	}
}

func (ptr *QDomProcessingInstruction) Target() string {
	defer qt.Recovering("QDomProcessingInstruction::target")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomProcessingInstruction_Target(ptr.Pointer()))
	}
	return ""
}

type QDomText struct {
	QDomCharacterData
}

type QDomText_ITF interface {
	QDomCharacterData_ITF
	QDomText_PTR() *QDomText
}

func (p *QDomText) QDomText_PTR() *QDomText {
	return p
}

func (p *QDomText) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDomCharacterData_PTR().Pointer()
	}
	return nil
}

func (p *QDomText) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDomCharacterData_PTR().SetPointer(ptr)
	}
}

func PointerFromQDomText(ptr QDomText_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomText_PTR().Pointer()
	}
	return nil
}

func NewQDomTextFromPointer(ptr unsafe.Pointer) *QDomText {
	var n = new(QDomText)
	n.SetPointer(ptr)
	return n
}

func newQDomTextFromPointer(ptr unsafe.Pointer) *QDomText {
	var n = NewQDomTextFromPointer(ptr)
	return n
}

func NewQDomText() *QDomText {
	defer qt.Recovering("QDomText::QDomText")

	return newQDomTextFromPointer(C.QDomText_NewQDomText())
}

func NewQDomText2(x QDomText_ITF) *QDomText {
	defer qt.Recovering("QDomText::QDomText")

	return newQDomTextFromPointer(C.QDomText_NewQDomText2(PointerFromQDomText(x)))
}

func (ptr *QDomText) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomText::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomText_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomText) SplitText(offset int) *QDomText {
	defer qt.Recovering("QDomText::splitText")

	if ptr.Pointer() != nil {
		return NewQDomTextFromPointer(C.QDomText_SplitText(ptr.Pointer(), C.int(offset)))
	}
	return nil
}

type QXmlAttributes struct {
	ptr unsafe.Pointer
}

type QXmlAttributes_ITF interface {
	QXmlAttributes_PTR() *QXmlAttributes
}

func (p *QXmlAttributes) QXmlAttributes_PTR() *QXmlAttributes {
	return p
}

func (p *QXmlAttributes) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlAttributes) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlAttributes(ptr QXmlAttributes_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlAttributes_PTR().Pointer()
	}
	return nil
}

func NewQXmlAttributesFromPointer(ptr unsafe.Pointer) *QXmlAttributes {
	var n = new(QXmlAttributes)
	n.SetPointer(ptr)
	return n
}

func newQXmlAttributesFromPointer(ptr unsafe.Pointer) *QXmlAttributes {
	var n = NewQXmlAttributesFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlAttributes_") {
		n.SetObjectNameAbs("QXmlAttributes_" + qt.Identifier())
	}
	return n
}

func NewQXmlAttributes() *QXmlAttributes {
	defer qt.Recovering("QXmlAttributes::QXmlAttributes")

	return newQXmlAttributesFromPointer(C.QXmlAttributes_NewQXmlAttributes())
}

func (ptr *QXmlAttributes) DestroyQXmlAttributes() {
	defer qt.Recovering("QXmlAttributes::~QXmlAttributes")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QXmlAttributes_DestroyQXmlAttributes(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlAttributes) Append(qName string, uri string, localPart string, value string) {
	defer qt.Recovering("QXmlAttributes::append")

	if ptr.Pointer() != nil {
		C.QXmlAttributes_Append(ptr.Pointer(), C.CString(qName), C.CString(uri), C.CString(localPart), C.CString(value))
	}
}

func (ptr *QXmlAttributes) Clear() {
	defer qt.Recovering("QXmlAttributes::clear")

	if ptr.Pointer() != nil {
		C.QXmlAttributes_Clear(ptr.Pointer())
	}
}

func (ptr *QXmlAttributes) Count() int {
	defer qt.Recovering("QXmlAttributes::count")

	if ptr.Pointer() != nil {
		return int(C.QXmlAttributes_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlAttributes) Index2(qName core.QLatin1String_ITF) int {
	defer qt.Recovering("QXmlAttributes::index")

	if ptr.Pointer() != nil {
		return int(C.QXmlAttributes_Index2(ptr.Pointer(), core.PointerFromQLatin1String(qName)))
	}
	return 0
}

func (ptr *QXmlAttributes) Index(qName string) int {
	defer qt.Recovering("QXmlAttributes::index")

	if ptr.Pointer() != nil {
		return int(C.QXmlAttributes_Index(ptr.Pointer(), C.CString(qName)))
	}
	return 0
}

func (ptr *QXmlAttributes) Index3(uri string, localPart string) int {
	defer qt.Recovering("QXmlAttributes::index")

	if ptr.Pointer() != nil {
		return int(C.QXmlAttributes_Index3(ptr.Pointer(), C.CString(uri), C.CString(localPart)))
	}
	return 0
}

func (ptr *QXmlAttributes) Length() int {
	defer qt.Recovering("QXmlAttributes::length")

	if ptr.Pointer() != nil {
		return int(C.QXmlAttributes_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlAttributes) LocalName(index int) string {
	defer qt.Recovering("QXmlAttributes::localName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_LocalName(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QXmlAttributes) QName(index int) string {
	defer qt.Recovering("QXmlAttributes::qName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_QName(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QXmlAttributes) Type2(qName string) string {
	defer qt.Recovering("QXmlAttributes::type")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Type2(ptr.Pointer(), C.CString(qName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Type3(uri string, localName string) string {
	defer qt.Recovering("QXmlAttributes::type")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Type3(ptr.Pointer(), C.CString(uri), C.CString(localName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Type(index int) string {
	defer qt.Recovering("QXmlAttributes::type")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Type(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QXmlAttributes) Uri(index int) string {
	defer qt.Recovering("QXmlAttributes::uri")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Uri(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QXmlAttributes) Value3(qName core.QLatin1String_ITF) string {
	defer qt.Recovering("QXmlAttributes::value")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Value3(ptr.Pointer(), core.PointerFromQLatin1String(qName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Value2(qName string) string {
	defer qt.Recovering("QXmlAttributes::value")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Value2(ptr.Pointer(), C.CString(qName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Value4(uri string, localName string) string {
	defer qt.Recovering("QXmlAttributes::value")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Value4(ptr.Pointer(), C.CString(uri), C.CString(localName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Value(index int) string {
	defer qt.Recovering("QXmlAttributes::value")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Value(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QXmlAttributes) ObjectNameAbs() string {
	defer qt.Recovering("QXmlAttributes::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlAttributes) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlAttributes::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlAttributes_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlContentHandler struct {
	ptr unsafe.Pointer
}

type QXmlContentHandler_ITF interface {
	QXmlContentHandler_PTR() *QXmlContentHandler
}

func (p *QXmlContentHandler) QXmlContentHandler_PTR() *QXmlContentHandler {
	return p
}

func (p *QXmlContentHandler) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlContentHandler) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlContentHandler(ptr QXmlContentHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlContentHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlContentHandlerFromPointer(ptr unsafe.Pointer) *QXmlContentHandler {
	var n = new(QXmlContentHandler)
	n.SetPointer(ptr)
	return n
}

func newQXmlContentHandlerFromPointer(ptr unsafe.Pointer) *QXmlContentHandler {
	var n = NewQXmlContentHandlerFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlContentHandler_") {
		n.SetObjectNameAbs("QXmlContentHandler_" + qt.Identifier())
	}
	return n
}

//export callbackQXmlContentHandler_Characters
func callbackQXmlContentHandler_Characters(ptr unsafe.Pointer, ptrName *C.char, ch *C.char) C.int {
	defer qt.Recovering("callback QXmlContentHandler::characters")

	if signal := qt.GetSignal(C.GoString(ptrName), "characters"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(ch))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlContentHandler) ConnectCharacters(f func(ch string) bool) {
	defer qt.Recovering("connect QXmlContentHandler::characters")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "characters", f)
	}
}

func (ptr *QXmlContentHandler) DisconnectCharacters(ch string) {
	defer qt.Recovering("disconnect QXmlContentHandler::characters")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "characters")
	}
}

func (ptr *QXmlContentHandler) Characters(ch string) bool {
	defer qt.Recovering("QXmlContentHandler::characters")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_Characters(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

//export callbackQXmlContentHandler_EndDocument
func callbackQXmlContentHandler_EndDocument(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QXmlContentHandler::endDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "endDocument"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlContentHandler) ConnectEndDocument(f func() bool) {
	defer qt.Recovering("connect QXmlContentHandler::endDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endDocument", f)
	}
}

func (ptr *QXmlContentHandler) DisconnectEndDocument() {
	defer qt.Recovering("disconnect QXmlContentHandler::endDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endDocument")
	}
}

func (ptr *QXmlContentHandler) EndDocument() bool {
	defer qt.Recovering("QXmlContentHandler::endDocument")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndDocument(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQXmlContentHandler_EndElement
func callbackQXmlContentHandler_EndElement(ptr unsafe.Pointer, ptrName *C.char, namespaceURI *C.char, localName *C.char, qName *C.char) C.int {
	defer qt.Recovering("callback QXmlContentHandler::endElement")

	if signal := qt.GetSignal(C.GoString(ptrName), "endElement"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string, string) bool)(C.GoString(namespaceURI), C.GoString(localName), C.GoString(qName))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlContentHandler) ConnectEndElement(f func(namespaceURI string, localName string, qName string) bool) {
	defer qt.Recovering("connect QXmlContentHandler::endElement")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endElement", f)
	}
}

func (ptr *QXmlContentHandler) DisconnectEndElement(namespaceURI string, localName string, qName string) {
	defer qt.Recovering("disconnect QXmlContentHandler::endElement")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endElement")
	}
}

func (ptr *QXmlContentHandler) EndElement(namespaceURI string, localName string, qName string) bool {
	defer qt.Recovering("QXmlContentHandler::endElement")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndElement(ptr.Pointer(), C.CString(namespaceURI), C.CString(localName), C.CString(qName)) != 0
	}
	return false
}

//export callbackQXmlContentHandler_EndPrefixMapping
func callbackQXmlContentHandler_EndPrefixMapping(ptr unsafe.Pointer, ptrName *C.char, prefix *C.char) C.int {
	defer qt.Recovering("callback QXmlContentHandler::endPrefixMapping")

	if signal := qt.GetSignal(C.GoString(ptrName), "endPrefixMapping"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(prefix))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlContentHandler) ConnectEndPrefixMapping(f func(prefix string) bool) {
	defer qt.Recovering("connect QXmlContentHandler::endPrefixMapping")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endPrefixMapping", f)
	}
}

func (ptr *QXmlContentHandler) DisconnectEndPrefixMapping(prefix string) {
	defer qt.Recovering("disconnect QXmlContentHandler::endPrefixMapping")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endPrefixMapping")
	}
}

func (ptr *QXmlContentHandler) EndPrefixMapping(prefix string) bool {
	defer qt.Recovering("QXmlContentHandler::endPrefixMapping")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndPrefixMapping(ptr.Pointer(), C.CString(prefix)) != 0
	}
	return false
}

//export callbackQXmlContentHandler_ErrorString
func callbackQXmlContentHandler_ErrorString(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QXmlContentHandler::errorString")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorString"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QXmlContentHandler) ConnectErrorString(f func() string) {
	defer qt.Recovering("connect QXmlContentHandler::errorString")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "errorString", f)
	}
}

func (ptr *QXmlContentHandler) DisconnectErrorString() {
	defer qt.Recovering("disconnect QXmlContentHandler::errorString")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "errorString")
	}
}

func (ptr *QXmlContentHandler) ErrorString() string {
	defer qt.Recovering("QXmlContentHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlContentHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQXmlContentHandler_IgnorableWhitespace
func callbackQXmlContentHandler_IgnorableWhitespace(ptr unsafe.Pointer, ptrName *C.char, ch *C.char) C.int {
	defer qt.Recovering("callback QXmlContentHandler::ignorableWhitespace")

	if signal := qt.GetSignal(C.GoString(ptrName), "ignorableWhitespace"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(ch))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlContentHandler) ConnectIgnorableWhitespace(f func(ch string) bool) {
	defer qt.Recovering("connect QXmlContentHandler::ignorableWhitespace")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "ignorableWhitespace", f)
	}
}

func (ptr *QXmlContentHandler) DisconnectIgnorableWhitespace(ch string) {
	defer qt.Recovering("disconnect QXmlContentHandler::ignorableWhitespace")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "ignorableWhitespace")
	}
}

func (ptr *QXmlContentHandler) IgnorableWhitespace(ch string) bool {
	defer qt.Recovering("QXmlContentHandler::ignorableWhitespace")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_IgnorableWhitespace(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

//export callbackQXmlContentHandler_ProcessingInstruction
func callbackQXmlContentHandler_ProcessingInstruction(ptr unsafe.Pointer, ptrName *C.char, target *C.char, data *C.char) C.int {
	defer qt.Recovering("callback QXmlContentHandler::processingInstruction")

	if signal := qt.GetSignal(C.GoString(ptrName), "processingInstruction"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string) bool)(C.GoString(target), C.GoString(data))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlContentHandler) ConnectProcessingInstruction(f func(target string, data string) bool) {
	defer qt.Recovering("connect QXmlContentHandler::processingInstruction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "processingInstruction", f)
	}
}

func (ptr *QXmlContentHandler) DisconnectProcessingInstruction(target string, data string) {
	defer qt.Recovering("disconnect QXmlContentHandler::processingInstruction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "processingInstruction")
	}
}

func (ptr *QXmlContentHandler) ProcessingInstruction(target string, data string) bool {
	defer qt.Recovering("QXmlContentHandler::processingInstruction")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_ProcessingInstruction(ptr.Pointer(), C.CString(target), C.CString(data)) != 0
	}
	return false
}

//export callbackQXmlContentHandler_SetDocumentLocator
func callbackQXmlContentHandler_SetDocumentLocator(ptr unsafe.Pointer, ptrName *C.char, locator unsafe.Pointer) {
	defer qt.Recovering("callback QXmlContentHandler::setDocumentLocator")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDocumentLocator"); signal != nil {
		signal.(func(*QXmlLocator))(NewQXmlLocatorFromPointer(locator))
	}

}

func (ptr *QXmlContentHandler) ConnectSetDocumentLocator(f func(locator *QXmlLocator)) {
	defer qt.Recovering("connect QXmlContentHandler::setDocumentLocator")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setDocumentLocator", f)
	}
}

func (ptr *QXmlContentHandler) DisconnectSetDocumentLocator(locator QXmlLocator_ITF) {
	defer qt.Recovering("disconnect QXmlContentHandler::setDocumentLocator")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setDocumentLocator")
	}
}

func (ptr *QXmlContentHandler) SetDocumentLocator(locator QXmlLocator_ITF) {
	defer qt.Recovering("QXmlContentHandler::setDocumentLocator")

	if ptr.Pointer() != nil {
		C.QXmlContentHandler_SetDocumentLocator(ptr.Pointer(), PointerFromQXmlLocator(locator))
	}
}

//export callbackQXmlContentHandler_SkippedEntity
func callbackQXmlContentHandler_SkippedEntity(ptr unsafe.Pointer, ptrName *C.char, name *C.char) C.int {
	defer qt.Recovering("callback QXmlContentHandler::skippedEntity")

	if signal := qt.GetSignal(C.GoString(ptrName), "skippedEntity"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlContentHandler) ConnectSkippedEntity(f func(name string) bool) {
	defer qt.Recovering("connect QXmlContentHandler::skippedEntity")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "skippedEntity", f)
	}
}

func (ptr *QXmlContentHandler) DisconnectSkippedEntity(name string) {
	defer qt.Recovering("disconnect QXmlContentHandler::skippedEntity")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "skippedEntity")
	}
}

func (ptr *QXmlContentHandler) SkippedEntity(name string) bool {
	defer qt.Recovering("QXmlContentHandler::skippedEntity")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_SkippedEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

//export callbackQXmlContentHandler_StartDocument
func callbackQXmlContentHandler_StartDocument(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QXmlContentHandler::startDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDocument"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlContentHandler) ConnectStartDocument(f func() bool) {
	defer qt.Recovering("connect QXmlContentHandler::startDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startDocument", f)
	}
}

func (ptr *QXmlContentHandler) DisconnectStartDocument() {
	defer qt.Recovering("disconnect QXmlContentHandler::startDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startDocument")
	}
}

func (ptr *QXmlContentHandler) StartDocument() bool {
	defer qt.Recovering("QXmlContentHandler::startDocument")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartDocument(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQXmlContentHandler_StartPrefixMapping
func callbackQXmlContentHandler_StartPrefixMapping(ptr unsafe.Pointer, ptrName *C.char, prefix *C.char, uri *C.char) C.int {
	defer qt.Recovering("callback QXmlContentHandler::startPrefixMapping")

	if signal := qt.GetSignal(C.GoString(ptrName), "startPrefixMapping"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string) bool)(C.GoString(prefix), C.GoString(uri))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlContentHandler) ConnectStartPrefixMapping(f func(prefix string, uri string) bool) {
	defer qt.Recovering("connect QXmlContentHandler::startPrefixMapping")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startPrefixMapping", f)
	}
}

func (ptr *QXmlContentHandler) DisconnectStartPrefixMapping(prefix string, uri string) {
	defer qt.Recovering("disconnect QXmlContentHandler::startPrefixMapping")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startPrefixMapping")
	}
}

func (ptr *QXmlContentHandler) StartPrefixMapping(prefix string, uri string) bool {
	defer qt.Recovering("QXmlContentHandler::startPrefixMapping")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartPrefixMapping(ptr.Pointer(), C.CString(prefix), C.CString(uri)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) DestroyQXmlContentHandler() {
	defer qt.Recovering("QXmlContentHandler::~QXmlContentHandler")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QXmlContentHandler_DestroyQXmlContentHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlContentHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlContentHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlContentHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlContentHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlContentHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlContentHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlDTDHandler struct {
	ptr unsafe.Pointer
}

type QXmlDTDHandler_ITF interface {
	QXmlDTDHandler_PTR() *QXmlDTDHandler
}

func (p *QXmlDTDHandler) QXmlDTDHandler_PTR() *QXmlDTDHandler {
	return p
}

func (p *QXmlDTDHandler) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlDTDHandler) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlDTDHandler(ptr QXmlDTDHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDTDHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlDTDHandlerFromPointer(ptr unsafe.Pointer) *QXmlDTDHandler {
	var n = new(QXmlDTDHandler)
	n.SetPointer(ptr)
	return n
}

func newQXmlDTDHandlerFromPointer(ptr unsafe.Pointer) *QXmlDTDHandler {
	var n = NewQXmlDTDHandlerFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlDTDHandler_") {
		n.SetObjectNameAbs("QXmlDTDHandler_" + qt.Identifier())
	}
	return n
}

//export callbackQXmlDTDHandler_ErrorString
func callbackQXmlDTDHandler_ErrorString(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QXmlDTDHandler::errorString")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorString"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QXmlDTDHandler) ConnectErrorString(f func() string) {
	defer qt.Recovering("connect QXmlDTDHandler::errorString")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "errorString", f)
	}
}

func (ptr *QXmlDTDHandler) DisconnectErrorString() {
	defer qt.Recovering("disconnect QXmlDTDHandler::errorString")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "errorString")
	}
}

func (ptr *QXmlDTDHandler) ErrorString() string {
	defer qt.Recovering("QXmlDTDHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDTDHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQXmlDTDHandler_NotationDecl
func callbackQXmlDTDHandler_NotationDecl(ptr unsafe.Pointer, ptrName *C.char, name *C.char, publicId *C.char, systemId *C.char) C.int {
	defer qt.Recovering("callback QXmlDTDHandler::notationDecl")

	if signal := qt.GetSignal(C.GoString(ptrName), "notationDecl"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string, string) bool)(C.GoString(name), C.GoString(publicId), C.GoString(systemId))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlDTDHandler) ConnectNotationDecl(f func(name string, publicId string, systemId string) bool) {
	defer qt.Recovering("connect QXmlDTDHandler::notationDecl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "notationDecl", f)
	}
}

func (ptr *QXmlDTDHandler) DisconnectNotationDecl(name string, publicId string, systemId string) {
	defer qt.Recovering("disconnect QXmlDTDHandler::notationDecl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "notationDecl")
	}
}

func (ptr *QXmlDTDHandler) NotationDecl(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlDTDHandler::notationDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDTDHandler_NotationDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

//export callbackQXmlDTDHandler_UnparsedEntityDecl
func callbackQXmlDTDHandler_UnparsedEntityDecl(ptr unsafe.Pointer, ptrName *C.char, name *C.char, publicId *C.char, systemId *C.char, notationName *C.char) C.int {
	defer qt.Recovering("callback QXmlDTDHandler::unparsedEntityDecl")

	if signal := qt.GetSignal(C.GoString(ptrName), "unparsedEntityDecl"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string, string, string) bool)(C.GoString(name), C.GoString(publicId), C.GoString(systemId), C.GoString(notationName))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlDTDHandler) ConnectUnparsedEntityDecl(f func(name string, publicId string, systemId string, notationName string) bool) {
	defer qt.Recovering("connect QXmlDTDHandler::unparsedEntityDecl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "unparsedEntityDecl", f)
	}
}

func (ptr *QXmlDTDHandler) DisconnectUnparsedEntityDecl(name string, publicId string, systemId string, notationName string) {
	defer qt.Recovering("disconnect QXmlDTDHandler::unparsedEntityDecl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "unparsedEntityDecl")
	}
}

func (ptr *QXmlDTDHandler) UnparsedEntityDecl(name string, publicId string, systemId string, notationName string) bool {
	defer qt.Recovering("QXmlDTDHandler::unparsedEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDTDHandler_UnparsedEntityDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId), C.CString(notationName)) != 0
	}
	return false
}

func (ptr *QXmlDTDHandler) DestroyQXmlDTDHandler() {
	defer qt.Recovering("QXmlDTDHandler::~QXmlDTDHandler")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QXmlDTDHandler_DestroyQXmlDTDHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlDTDHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlDTDHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDTDHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDTDHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlDTDHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlDTDHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlDeclHandler struct {
	ptr unsafe.Pointer
}

type QXmlDeclHandler_ITF interface {
	QXmlDeclHandler_PTR() *QXmlDeclHandler
}

func (p *QXmlDeclHandler) QXmlDeclHandler_PTR() *QXmlDeclHandler {
	return p
}

func (p *QXmlDeclHandler) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlDeclHandler) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlDeclHandler(ptr QXmlDeclHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDeclHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlDeclHandlerFromPointer(ptr unsafe.Pointer) *QXmlDeclHandler {
	var n = new(QXmlDeclHandler)
	n.SetPointer(ptr)
	return n
}

func newQXmlDeclHandlerFromPointer(ptr unsafe.Pointer) *QXmlDeclHandler {
	var n = NewQXmlDeclHandlerFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlDeclHandler_") {
		n.SetObjectNameAbs("QXmlDeclHandler_" + qt.Identifier())
	}
	return n
}

//export callbackQXmlDeclHandler_AttributeDecl
func callbackQXmlDeclHandler_AttributeDecl(ptr unsafe.Pointer, ptrName *C.char, eName *C.char, aName *C.char, ty *C.char, valueDefault *C.char, value *C.char) C.int {
	defer qt.Recovering("callback QXmlDeclHandler::attributeDecl")

	if signal := qt.GetSignal(C.GoString(ptrName), "attributeDecl"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string, string, string, string) bool)(C.GoString(eName), C.GoString(aName), C.GoString(ty), C.GoString(valueDefault), C.GoString(value))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlDeclHandler) ConnectAttributeDecl(f func(eName string, aName string, ty string, valueDefault string, value string) bool) {
	defer qt.Recovering("connect QXmlDeclHandler::attributeDecl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "attributeDecl", f)
	}
}

func (ptr *QXmlDeclHandler) DisconnectAttributeDecl(eName string, aName string, ty string, valueDefault string, value string) {
	defer qt.Recovering("disconnect QXmlDeclHandler::attributeDecl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "attributeDecl")
	}
}

func (ptr *QXmlDeclHandler) AttributeDecl(eName string, aName string, ty string, valueDefault string, value string) bool {
	defer qt.Recovering("QXmlDeclHandler::attributeDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDeclHandler_AttributeDecl(ptr.Pointer(), C.CString(eName), C.CString(aName), C.CString(ty), C.CString(valueDefault), C.CString(value)) != 0
	}
	return false
}

//export callbackQXmlDeclHandler_ErrorString
func callbackQXmlDeclHandler_ErrorString(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QXmlDeclHandler::errorString")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorString"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QXmlDeclHandler) ConnectErrorString(f func() string) {
	defer qt.Recovering("connect QXmlDeclHandler::errorString")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "errorString", f)
	}
}

func (ptr *QXmlDeclHandler) DisconnectErrorString() {
	defer qt.Recovering("disconnect QXmlDeclHandler::errorString")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "errorString")
	}
}

func (ptr *QXmlDeclHandler) ErrorString() string {
	defer qt.Recovering("QXmlDeclHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDeclHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQXmlDeclHandler_ExternalEntityDecl
func callbackQXmlDeclHandler_ExternalEntityDecl(ptr unsafe.Pointer, ptrName *C.char, name *C.char, publicId *C.char, systemId *C.char) C.int {
	defer qt.Recovering("callback QXmlDeclHandler::externalEntityDecl")

	if signal := qt.GetSignal(C.GoString(ptrName), "externalEntityDecl"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string, string) bool)(C.GoString(name), C.GoString(publicId), C.GoString(systemId))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlDeclHandler) ConnectExternalEntityDecl(f func(name string, publicId string, systemId string) bool) {
	defer qt.Recovering("connect QXmlDeclHandler::externalEntityDecl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "externalEntityDecl", f)
	}
}

func (ptr *QXmlDeclHandler) DisconnectExternalEntityDecl(name string, publicId string, systemId string) {
	defer qt.Recovering("disconnect QXmlDeclHandler::externalEntityDecl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "externalEntityDecl")
	}
}

func (ptr *QXmlDeclHandler) ExternalEntityDecl(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlDeclHandler::externalEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDeclHandler_ExternalEntityDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

//export callbackQXmlDeclHandler_InternalEntityDecl
func callbackQXmlDeclHandler_InternalEntityDecl(ptr unsafe.Pointer, ptrName *C.char, name *C.char, value *C.char) C.int {
	defer qt.Recovering("callback QXmlDeclHandler::internalEntityDecl")

	if signal := qt.GetSignal(C.GoString(ptrName), "internalEntityDecl"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string) bool)(C.GoString(name), C.GoString(value))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlDeclHandler) ConnectInternalEntityDecl(f func(name string, value string) bool) {
	defer qt.Recovering("connect QXmlDeclHandler::internalEntityDecl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "internalEntityDecl", f)
	}
}

func (ptr *QXmlDeclHandler) DisconnectInternalEntityDecl(name string, value string) {
	defer qt.Recovering("disconnect QXmlDeclHandler::internalEntityDecl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "internalEntityDecl")
	}
}

func (ptr *QXmlDeclHandler) InternalEntityDecl(name string, value string) bool {
	defer qt.Recovering("QXmlDeclHandler::internalEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDeclHandler_InternalEntityDecl(ptr.Pointer(), C.CString(name), C.CString(value)) != 0
	}
	return false
}

func (ptr *QXmlDeclHandler) DestroyQXmlDeclHandler() {
	defer qt.Recovering("QXmlDeclHandler::~QXmlDeclHandler")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QXmlDeclHandler_DestroyQXmlDeclHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlDeclHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlDeclHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDeclHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDeclHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlDeclHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlDeclHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlDefaultHandler struct {
	QXmlContentHandler
	QXmlErrorHandler
	QXmlDTDHandler
	QXmlEntityResolver
	QXmlLexicalHandler
	QXmlDeclHandler
}

type QXmlDefaultHandler_ITF interface {
	QXmlContentHandler_ITF
	QXmlErrorHandler_ITF
	QXmlDTDHandler_ITF
	QXmlEntityResolver_ITF
	QXmlLexicalHandler_ITF
	QXmlDeclHandler_ITF
	QXmlDefaultHandler_PTR() *QXmlDefaultHandler
}

func (p *QXmlDefaultHandler) QXmlDefaultHandler_PTR() *QXmlDefaultHandler {
	return p
}

func (p *QXmlDefaultHandler) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QXmlContentHandler_PTR().Pointer()
	}
	return nil
}

func (p *QXmlDefaultHandler) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QXmlContentHandler_PTR().SetPointer(ptr)
		p.QXmlErrorHandler_PTR().SetPointer(ptr)
		p.QXmlDTDHandler_PTR().SetPointer(ptr)
		p.QXmlEntityResolver_PTR().SetPointer(ptr)
		p.QXmlLexicalHandler_PTR().SetPointer(ptr)
		p.QXmlDeclHandler_PTR().SetPointer(ptr)
	}
}

func PointerFromQXmlDefaultHandler(ptr QXmlDefaultHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDefaultHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlDefaultHandlerFromPointer(ptr unsafe.Pointer) *QXmlDefaultHandler {
	var n = new(QXmlDefaultHandler)
	n.SetPointer(ptr)
	return n
}

func newQXmlDefaultHandlerFromPointer(ptr unsafe.Pointer) *QXmlDefaultHandler {
	var n = NewQXmlDefaultHandlerFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlDefaultHandler_") {
		n.SetObjectNameAbs("QXmlDefaultHandler_" + qt.Identifier())
	}
	return n
}

func NewQXmlDefaultHandler() *QXmlDefaultHandler {
	defer qt.Recovering("QXmlDefaultHandler::QXmlDefaultHandler")

	return newQXmlDefaultHandlerFromPointer(C.QXmlDefaultHandler_NewQXmlDefaultHandler())
}

func (ptr *QXmlDefaultHandler) DestroyQXmlDefaultHandler() {
	defer qt.Recovering("QXmlDefaultHandler::~QXmlDefaultHandler")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QXmlDefaultHandler_DestroyQXmlDefaultHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQXmlDefaultHandler_AttributeDecl
func callbackQXmlDefaultHandler_AttributeDecl(ptr unsafe.Pointer, ptrName *C.char, eName *C.char, aName *C.char, ty *C.char, valueDefault *C.char, value *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::attributeDecl")

	if signal := qt.GetSignal(C.GoString(ptrName), "attributeDecl"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string, string, string, string) bool)(C.GoString(eName), C.GoString(aName), C.GoString(ty), C.GoString(valueDefault), C.GoString(value))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).AttributeDeclDefault(C.GoString(eName), C.GoString(aName), C.GoString(ty), C.GoString(valueDefault), C.GoString(value))))
}

func (ptr *QXmlDefaultHandler) ConnectAttributeDecl(f func(eName string, aName string, ty string, valueDefault string, value string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::attributeDecl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "attributeDecl", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectAttributeDecl() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::attributeDecl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "attributeDecl")
	}
}

func (ptr *QXmlDefaultHandler) AttributeDecl(eName string, aName string, ty string, valueDefault string, value string) bool {
	defer qt.Recovering("QXmlDefaultHandler::attributeDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_AttributeDecl(ptr.Pointer(), C.CString(eName), C.CString(aName), C.CString(ty), C.CString(valueDefault), C.CString(value)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) AttributeDeclDefault(eName string, aName string, ty string, valueDefault string, value string) bool {
	defer qt.Recovering("QXmlDefaultHandler::attributeDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_AttributeDeclDefault(ptr.Pointer(), C.CString(eName), C.CString(aName), C.CString(ty), C.CString(valueDefault), C.CString(value)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_Characters
func callbackQXmlDefaultHandler_Characters(ptr unsafe.Pointer, ptrName *C.char, ch *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::characters")

	if signal := qt.GetSignal(C.GoString(ptrName), "characters"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(ch))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).CharactersDefault(C.GoString(ch))))
}

func (ptr *QXmlDefaultHandler) ConnectCharacters(f func(ch string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::characters")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "characters", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectCharacters() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::characters")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "characters")
	}
}

func (ptr *QXmlDefaultHandler) Characters(ch string) bool {
	defer qt.Recovering("QXmlDefaultHandler::characters")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_Characters(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) CharactersDefault(ch string) bool {
	defer qt.Recovering("QXmlDefaultHandler::characters")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_CharactersDefault(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_Comment
func callbackQXmlDefaultHandler_Comment(ptr unsafe.Pointer, ptrName *C.char, ch *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::comment")

	if signal := qt.GetSignal(C.GoString(ptrName), "comment"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(ch))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).CommentDefault(C.GoString(ch))))
}

func (ptr *QXmlDefaultHandler) ConnectComment(f func(ch string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::comment")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "comment", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectComment() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::comment")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "comment")
	}
}

func (ptr *QXmlDefaultHandler) Comment(ch string) bool {
	defer qt.Recovering("QXmlDefaultHandler::comment")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_Comment(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) CommentDefault(ch string) bool {
	defer qt.Recovering("QXmlDefaultHandler::comment")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_CommentDefault(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_EndCDATA
func callbackQXmlDefaultHandler_EndCDATA(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::endCDATA")

	if signal := qt.GetSignal(C.GoString(ptrName), "endCDATA"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).EndCDATADefault()))
}

func (ptr *QXmlDefaultHandler) ConnectEndCDATA(f func() bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::endCDATA")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endCDATA", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectEndCDATA() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::endCDATA")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endCDATA")
	}
}

func (ptr *QXmlDefaultHandler) EndCDATA() bool {
	defer qt.Recovering("QXmlDefaultHandler::endCDATA")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndCDATA(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndCDATADefault() bool {
	defer qt.Recovering("QXmlDefaultHandler::endCDATA")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndCDATADefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_EndDTD
func callbackQXmlDefaultHandler_EndDTD(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::endDTD")

	if signal := qt.GetSignal(C.GoString(ptrName), "endDTD"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).EndDTDDefault()))
}

func (ptr *QXmlDefaultHandler) ConnectEndDTD(f func() bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::endDTD")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endDTD", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectEndDTD() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::endDTD")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endDTD")
	}
}

func (ptr *QXmlDefaultHandler) EndDTD() bool {
	defer qt.Recovering("QXmlDefaultHandler::endDTD")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndDTD(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndDTDDefault() bool {
	defer qt.Recovering("QXmlDefaultHandler::endDTD")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndDTDDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_EndDocument
func callbackQXmlDefaultHandler_EndDocument(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::endDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "endDocument"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).EndDocumentDefault()))
}

func (ptr *QXmlDefaultHandler) ConnectEndDocument(f func() bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::endDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endDocument", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectEndDocument() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::endDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endDocument")
	}
}

func (ptr *QXmlDefaultHandler) EndDocument() bool {
	defer qt.Recovering("QXmlDefaultHandler::endDocument")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndDocumentDefault() bool {
	defer qt.Recovering("QXmlDefaultHandler::endDocument")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndDocumentDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_EndElement
func callbackQXmlDefaultHandler_EndElement(ptr unsafe.Pointer, ptrName *C.char, namespaceURI *C.char, localName *C.char, qName *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::endElement")

	if signal := qt.GetSignal(C.GoString(ptrName), "endElement"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string, string) bool)(C.GoString(namespaceURI), C.GoString(localName), C.GoString(qName))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).EndElementDefault(C.GoString(namespaceURI), C.GoString(localName), C.GoString(qName))))
}

func (ptr *QXmlDefaultHandler) ConnectEndElement(f func(namespaceURI string, localName string, qName string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::endElement")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endElement", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectEndElement() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::endElement")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endElement")
	}
}

func (ptr *QXmlDefaultHandler) EndElement(namespaceURI string, localName string, qName string) bool {
	defer qt.Recovering("QXmlDefaultHandler::endElement")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndElement(ptr.Pointer(), C.CString(namespaceURI), C.CString(localName), C.CString(qName)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndElementDefault(namespaceURI string, localName string, qName string) bool {
	defer qt.Recovering("QXmlDefaultHandler::endElement")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndElementDefault(ptr.Pointer(), C.CString(namespaceURI), C.CString(localName), C.CString(qName)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_EndEntity
func callbackQXmlDefaultHandler_EndEntity(ptr unsafe.Pointer, ptrName *C.char, name *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::endEntity")

	if signal := qt.GetSignal(C.GoString(ptrName), "endEntity"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).EndEntityDefault(C.GoString(name))))
}

func (ptr *QXmlDefaultHandler) ConnectEndEntity(f func(name string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::endEntity")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endEntity", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectEndEntity() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::endEntity")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endEntity")
	}
}

func (ptr *QXmlDefaultHandler) EndEntity(name string) bool {
	defer qt.Recovering("QXmlDefaultHandler::endEntity")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndEntityDefault(name string) bool {
	defer qt.Recovering("QXmlDefaultHandler::endEntity")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndEntityDefault(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_EndPrefixMapping
func callbackQXmlDefaultHandler_EndPrefixMapping(ptr unsafe.Pointer, ptrName *C.char, prefix *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::endPrefixMapping")

	if signal := qt.GetSignal(C.GoString(ptrName), "endPrefixMapping"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(prefix))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).EndPrefixMappingDefault(C.GoString(prefix))))
}

func (ptr *QXmlDefaultHandler) ConnectEndPrefixMapping(f func(prefix string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::endPrefixMapping")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endPrefixMapping", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectEndPrefixMapping() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::endPrefixMapping")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endPrefixMapping")
	}
}

func (ptr *QXmlDefaultHandler) EndPrefixMapping(prefix string) bool {
	defer qt.Recovering("QXmlDefaultHandler::endPrefixMapping")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndPrefixMapping(ptr.Pointer(), C.CString(prefix)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndPrefixMappingDefault(prefix string) bool {
	defer qt.Recovering("QXmlDefaultHandler::endPrefixMapping")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndPrefixMappingDefault(ptr.Pointer(), C.CString(prefix)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_Error
func callbackQXmlDefaultHandler_Error(ptr unsafe.Pointer, ptrName *C.char, exception unsafe.Pointer) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QXmlParseException) bool)(NewQXmlParseExceptionFromPointer(exception))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).ErrorDefault(NewQXmlParseExceptionFromPointer(exception))))
}

func (ptr *QXmlDefaultHandler) ConnectError(f func(exception *QXmlParseException) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::error")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "error", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectError() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::error")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "error")
	}
}

func (ptr *QXmlDefaultHandler) Error(exception QXmlParseException_ITF) bool {
	defer qt.Recovering("QXmlDefaultHandler::error")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_Error(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) ErrorDefault(exception QXmlParseException_ITF) bool {
	defer qt.Recovering("QXmlDefaultHandler::error")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_ErrorDefault(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_ErrorString
func callbackQXmlDefaultHandler_ErrorString(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QXmlDefaultHandler::errorString")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorString"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQXmlDefaultHandlerFromPointer(ptr).ErrorStringDefault())
}

func (ptr *QXmlDefaultHandler) ConnectErrorString(f func() string) {
	defer qt.Recovering("connect QXmlDefaultHandler::errorString")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "errorString", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectErrorString() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::errorString")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "errorString")
	}
}

func (ptr *QXmlDefaultHandler) ErrorString() string {
	defer qt.Recovering("QXmlDefaultHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDefaultHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDefaultHandler) ErrorStringDefault() string {
	defer qt.Recovering("QXmlDefaultHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDefaultHandler_ErrorStringDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQXmlDefaultHandler_ExternalEntityDecl
func callbackQXmlDefaultHandler_ExternalEntityDecl(ptr unsafe.Pointer, ptrName *C.char, name *C.char, publicId *C.char, systemId *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::externalEntityDecl")

	if signal := qt.GetSignal(C.GoString(ptrName), "externalEntityDecl"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string, string) bool)(C.GoString(name), C.GoString(publicId), C.GoString(systemId))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).ExternalEntityDeclDefault(C.GoString(name), C.GoString(publicId), C.GoString(systemId))))
}

func (ptr *QXmlDefaultHandler) ConnectExternalEntityDecl(f func(name string, publicId string, systemId string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::externalEntityDecl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "externalEntityDecl", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectExternalEntityDecl() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::externalEntityDecl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "externalEntityDecl")
	}
}

func (ptr *QXmlDefaultHandler) ExternalEntityDecl(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlDefaultHandler::externalEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_ExternalEntityDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) ExternalEntityDeclDefault(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlDefaultHandler::externalEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_ExternalEntityDeclDefault(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_FatalError
func callbackQXmlDefaultHandler_FatalError(ptr unsafe.Pointer, ptrName *C.char, exception unsafe.Pointer) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::fatalError")

	if signal := qt.GetSignal(C.GoString(ptrName), "fatalError"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QXmlParseException) bool)(NewQXmlParseExceptionFromPointer(exception))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).FatalErrorDefault(NewQXmlParseExceptionFromPointer(exception))))
}

func (ptr *QXmlDefaultHandler) ConnectFatalError(f func(exception *QXmlParseException) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::fatalError")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "fatalError", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectFatalError() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::fatalError")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "fatalError")
	}
}

func (ptr *QXmlDefaultHandler) FatalError(exception QXmlParseException_ITF) bool {
	defer qt.Recovering("QXmlDefaultHandler::fatalError")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_FatalError(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) FatalErrorDefault(exception QXmlParseException_ITF) bool {
	defer qt.Recovering("QXmlDefaultHandler::fatalError")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_FatalErrorDefault(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_IgnorableWhitespace
func callbackQXmlDefaultHandler_IgnorableWhitespace(ptr unsafe.Pointer, ptrName *C.char, ch *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::ignorableWhitespace")

	if signal := qt.GetSignal(C.GoString(ptrName), "ignorableWhitespace"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(ch))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).IgnorableWhitespaceDefault(C.GoString(ch))))
}

func (ptr *QXmlDefaultHandler) ConnectIgnorableWhitespace(f func(ch string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::ignorableWhitespace")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "ignorableWhitespace", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectIgnorableWhitespace() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::ignorableWhitespace")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "ignorableWhitespace")
	}
}

func (ptr *QXmlDefaultHandler) IgnorableWhitespace(ch string) bool {
	defer qt.Recovering("QXmlDefaultHandler::ignorableWhitespace")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_IgnorableWhitespace(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) IgnorableWhitespaceDefault(ch string) bool {
	defer qt.Recovering("QXmlDefaultHandler::ignorableWhitespace")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_IgnorableWhitespaceDefault(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_InternalEntityDecl
func callbackQXmlDefaultHandler_InternalEntityDecl(ptr unsafe.Pointer, ptrName *C.char, name *C.char, value *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::internalEntityDecl")

	if signal := qt.GetSignal(C.GoString(ptrName), "internalEntityDecl"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string) bool)(C.GoString(name), C.GoString(value))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).InternalEntityDeclDefault(C.GoString(name), C.GoString(value))))
}

func (ptr *QXmlDefaultHandler) ConnectInternalEntityDecl(f func(name string, value string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::internalEntityDecl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "internalEntityDecl", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectInternalEntityDecl() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::internalEntityDecl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "internalEntityDecl")
	}
}

func (ptr *QXmlDefaultHandler) InternalEntityDecl(name string, value string) bool {
	defer qt.Recovering("QXmlDefaultHandler::internalEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_InternalEntityDecl(ptr.Pointer(), C.CString(name), C.CString(value)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) InternalEntityDeclDefault(name string, value string) bool {
	defer qt.Recovering("QXmlDefaultHandler::internalEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_InternalEntityDeclDefault(ptr.Pointer(), C.CString(name), C.CString(value)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_NotationDecl
func callbackQXmlDefaultHandler_NotationDecl(ptr unsafe.Pointer, ptrName *C.char, name *C.char, publicId *C.char, systemId *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::notationDecl")

	if signal := qt.GetSignal(C.GoString(ptrName), "notationDecl"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string, string) bool)(C.GoString(name), C.GoString(publicId), C.GoString(systemId))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).NotationDeclDefault(C.GoString(name), C.GoString(publicId), C.GoString(systemId))))
}

func (ptr *QXmlDefaultHandler) ConnectNotationDecl(f func(name string, publicId string, systemId string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::notationDecl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "notationDecl", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectNotationDecl() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::notationDecl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "notationDecl")
	}
}

func (ptr *QXmlDefaultHandler) NotationDecl(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlDefaultHandler::notationDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_NotationDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) NotationDeclDefault(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlDefaultHandler::notationDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_NotationDeclDefault(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_ProcessingInstruction
func callbackQXmlDefaultHandler_ProcessingInstruction(ptr unsafe.Pointer, ptrName *C.char, target *C.char, data *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::processingInstruction")

	if signal := qt.GetSignal(C.GoString(ptrName), "processingInstruction"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string) bool)(C.GoString(target), C.GoString(data))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).ProcessingInstructionDefault(C.GoString(target), C.GoString(data))))
}

func (ptr *QXmlDefaultHandler) ConnectProcessingInstruction(f func(target string, data string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::processingInstruction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "processingInstruction", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectProcessingInstruction() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::processingInstruction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "processingInstruction")
	}
}

func (ptr *QXmlDefaultHandler) ProcessingInstruction(target string, data string) bool {
	defer qt.Recovering("QXmlDefaultHandler::processingInstruction")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_ProcessingInstruction(ptr.Pointer(), C.CString(target), C.CString(data)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) ProcessingInstructionDefault(target string, data string) bool {
	defer qt.Recovering("QXmlDefaultHandler::processingInstruction")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_ProcessingInstructionDefault(ptr.Pointer(), C.CString(target), C.CString(data)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_ResolveEntity
func callbackQXmlDefaultHandler_ResolveEntity(ptr unsafe.Pointer, ptrName *C.char, publicId *C.char, systemId *C.char, ret unsafe.Pointer) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::resolveEntity")

	if signal := qt.GetSignal(C.GoString(ptrName), "resolveEntity"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string, *QXmlInputSource) bool)(C.GoString(publicId), C.GoString(systemId), NewQXmlInputSourceFromPointer(ret))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).ResolveEntityDefault(C.GoString(publicId), C.GoString(systemId), NewQXmlInputSourceFromPointer(ret))))
}

func (ptr *QXmlDefaultHandler) ConnectResolveEntity(f func(publicId string, systemId string, ret *QXmlInputSource) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::resolveEntity")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "resolveEntity", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectResolveEntity() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::resolveEntity")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "resolveEntity")
	}
}

func (ptr *QXmlDefaultHandler) ResolveEntity(publicId string, systemId string, ret QXmlInputSource_ITF) bool {
	defer qt.Recovering("QXmlDefaultHandler::resolveEntity")

	if ptr.Pointer() != nil {

	}
	return false
}

func (ptr *QXmlDefaultHandler) ResolveEntityDefault(publicId string, systemId string, ret QXmlInputSource_ITF) bool {
	defer qt.Recovering("QXmlDefaultHandler::resolveEntity")

	if ptr.Pointer() != nil {

	}
	return false
}

//export callbackQXmlDefaultHandler_SetDocumentLocator
func callbackQXmlDefaultHandler_SetDocumentLocator(ptr unsafe.Pointer, ptrName *C.char, locator unsafe.Pointer) {
	defer qt.Recovering("callback QXmlDefaultHandler::setDocumentLocator")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDocumentLocator"); signal != nil {
		signal.(func(*QXmlLocator))(NewQXmlLocatorFromPointer(locator))
	} else {
		NewQXmlDefaultHandlerFromPointer(ptr).SetDocumentLocatorDefault(NewQXmlLocatorFromPointer(locator))
	}
}

func (ptr *QXmlDefaultHandler) ConnectSetDocumentLocator(f func(locator *QXmlLocator)) {
	defer qt.Recovering("connect QXmlDefaultHandler::setDocumentLocator")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setDocumentLocator", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectSetDocumentLocator() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::setDocumentLocator")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setDocumentLocator")
	}
}

func (ptr *QXmlDefaultHandler) SetDocumentLocator(locator QXmlLocator_ITF) {
	defer qt.Recovering("QXmlDefaultHandler::setDocumentLocator")

	if ptr.Pointer() != nil {
		C.QXmlDefaultHandler_SetDocumentLocator(ptr.Pointer(), PointerFromQXmlLocator(locator))
	}
}

func (ptr *QXmlDefaultHandler) SetDocumentLocatorDefault(locator QXmlLocator_ITF) {
	defer qt.Recovering("QXmlDefaultHandler::setDocumentLocator")

	if ptr.Pointer() != nil {
		C.QXmlDefaultHandler_SetDocumentLocatorDefault(ptr.Pointer(), PointerFromQXmlLocator(locator))
	}
}

//export callbackQXmlDefaultHandler_SkippedEntity
func callbackQXmlDefaultHandler_SkippedEntity(ptr unsafe.Pointer, ptrName *C.char, name *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::skippedEntity")

	if signal := qt.GetSignal(C.GoString(ptrName), "skippedEntity"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).SkippedEntityDefault(C.GoString(name))))
}

func (ptr *QXmlDefaultHandler) ConnectSkippedEntity(f func(name string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::skippedEntity")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "skippedEntity", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectSkippedEntity() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::skippedEntity")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "skippedEntity")
	}
}

func (ptr *QXmlDefaultHandler) SkippedEntity(name string) bool {
	defer qt.Recovering("QXmlDefaultHandler::skippedEntity")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_SkippedEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) SkippedEntityDefault(name string) bool {
	defer qt.Recovering("QXmlDefaultHandler::skippedEntity")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_SkippedEntityDefault(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_StartCDATA
func callbackQXmlDefaultHandler_StartCDATA(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::startCDATA")

	if signal := qt.GetSignal(C.GoString(ptrName), "startCDATA"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).StartCDATADefault()))
}

func (ptr *QXmlDefaultHandler) ConnectStartCDATA(f func() bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::startCDATA")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startCDATA", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectStartCDATA() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::startCDATA")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startCDATA")
	}
}

func (ptr *QXmlDefaultHandler) StartCDATA() bool {
	defer qt.Recovering("QXmlDefaultHandler::startCDATA")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartCDATA(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartCDATADefault() bool {
	defer qt.Recovering("QXmlDefaultHandler::startCDATA")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartCDATADefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_StartDTD
func callbackQXmlDefaultHandler_StartDTD(ptr unsafe.Pointer, ptrName *C.char, name *C.char, publicId *C.char, systemId *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::startDTD")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDTD"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string, string) bool)(C.GoString(name), C.GoString(publicId), C.GoString(systemId))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).StartDTDDefault(C.GoString(name), C.GoString(publicId), C.GoString(systemId))))
}

func (ptr *QXmlDefaultHandler) ConnectStartDTD(f func(name string, publicId string, systemId string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::startDTD")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startDTD", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectStartDTD() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::startDTD")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startDTD")
	}
}

func (ptr *QXmlDefaultHandler) StartDTD(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlDefaultHandler::startDTD")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartDTD(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartDTDDefault(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlDefaultHandler::startDTD")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartDTDDefault(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_StartDocument
func callbackQXmlDefaultHandler_StartDocument(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::startDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDocument"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).StartDocumentDefault()))
}

func (ptr *QXmlDefaultHandler) ConnectStartDocument(f func() bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::startDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startDocument", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectStartDocument() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::startDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startDocument")
	}
}

func (ptr *QXmlDefaultHandler) StartDocument() bool {
	defer qt.Recovering("QXmlDefaultHandler::startDocument")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartDocumentDefault() bool {
	defer qt.Recovering("QXmlDefaultHandler::startDocument")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartDocumentDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_StartEntity
func callbackQXmlDefaultHandler_StartEntity(ptr unsafe.Pointer, ptrName *C.char, name *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::startEntity")

	if signal := qt.GetSignal(C.GoString(ptrName), "startEntity"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).StartEntityDefault(C.GoString(name))))
}

func (ptr *QXmlDefaultHandler) ConnectStartEntity(f func(name string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::startEntity")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startEntity", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectStartEntity() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::startEntity")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startEntity")
	}
}

func (ptr *QXmlDefaultHandler) StartEntity(name string) bool {
	defer qt.Recovering("QXmlDefaultHandler::startEntity")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartEntityDefault(name string) bool {
	defer qt.Recovering("QXmlDefaultHandler::startEntity")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartEntityDefault(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_StartPrefixMapping
func callbackQXmlDefaultHandler_StartPrefixMapping(ptr unsafe.Pointer, ptrName *C.char, prefix *C.char, uri *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::startPrefixMapping")

	if signal := qt.GetSignal(C.GoString(ptrName), "startPrefixMapping"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string) bool)(C.GoString(prefix), C.GoString(uri))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).StartPrefixMappingDefault(C.GoString(prefix), C.GoString(uri))))
}

func (ptr *QXmlDefaultHandler) ConnectStartPrefixMapping(f func(prefix string, uri string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::startPrefixMapping")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startPrefixMapping", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectStartPrefixMapping() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::startPrefixMapping")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startPrefixMapping")
	}
}

func (ptr *QXmlDefaultHandler) StartPrefixMapping(prefix string, uri string) bool {
	defer qt.Recovering("QXmlDefaultHandler::startPrefixMapping")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartPrefixMapping(ptr.Pointer(), C.CString(prefix), C.CString(uri)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartPrefixMappingDefault(prefix string, uri string) bool {
	defer qt.Recovering("QXmlDefaultHandler::startPrefixMapping")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartPrefixMappingDefault(ptr.Pointer(), C.CString(prefix), C.CString(uri)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_UnparsedEntityDecl
func callbackQXmlDefaultHandler_UnparsedEntityDecl(ptr unsafe.Pointer, ptrName *C.char, name *C.char, publicId *C.char, systemId *C.char, notationName *C.char) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::unparsedEntityDecl")

	if signal := qt.GetSignal(C.GoString(ptrName), "unparsedEntityDecl"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string, string, string) bool)(C.GoString(name), C.GoString(publicId), C.GoString(systemId), C.GoString(notationName))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).UnparsedEntityDeclDefault(C.GoString(name), C.GoString(publicId), C.GoString(systemId), C.GoString(notationName))))
}

func (ptr *QXmlDefaultHandler) ConnectUnparsedEntityDecl(f func(name string, publicId string, systemId string, notationName string) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::unparsedEntityDecl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "unparsedEntityDecl", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectUnparsedEntityDecl() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::unparsedEntityDecl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "unparsedEntityDecl")
	}
}

func (ptr *QXmlDefaultHandler) UnparsedEntityDecl(name string, publicId string, systemId string, notationName string) bool {
	defer qt.Recovering("QXmlDefaultHandler::unparsedEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_UnparsedEntityDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId), C.CString(notationName)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) UnparsedEntityDeclDefault(name string, publicId string, systemId string, notationName string) bool {
	defer qt.Recovering("QXmlDefaultHandler::unparsedEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_UnparsedEntityDeclDefault(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId), C.CString(notationName)) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_Warning
func callbackQXmlDefaultHandler_Warning(ptr unsafe.Pointer, ptrName *C.char, exception unsafe.Pointer) C.int {
	defer qt.Recovering("callback QXmlDefaultHandler::warning")

	if signal := qt.GetSignal(C.GoString(ptrName), "warning"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QXmlParseException) bool)(NewQXmlParseExceptionFromPointer(exception))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).WarningDefault(NewQXmlParseExceptionFromPointer(exception))))
}

func (ptr *QXmlDefaultHandler) ConnectWarning(f func(exception *QXmlParseException) bool) {
	defer qt.Recovering("connect QXmlDefaultHandler::warning")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "warning", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectWarning() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::warning")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "warning")
	}
}

func (ptr *QXmlDefaultHandler) Warning(exception QXmlParseException_ITF) bool {
	defer qt.Recovering("QXmlDefaultHandler::warning")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_Warning(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) WarningDefault(exception QXmlParseException_ITF) bool {
	defer qt.Recovering("QXmlDefaultHandler::warning")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_WarningDefault(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlDefaultHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDefaultHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDefaultHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlDefaultHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlDefaultHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlEntityResolver struct {
	ptr unsafe.Pointer
}

type QXmlEntityResolver_ITF interface {
	QXmlEntityResolver_PTR() *QXmlEntityResolver
}

func (p *QXmlEntityResolver) QXmlEntityResolver_PTR() *QXmlEntityResolver {
	return p
}

func (p *QXmlEntityResolver) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlEntityResolver) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlEntityResolver(ptr QXmlEntityResolver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlEntityResolver_PTR().Pointer()
	}
	return nil
}

func NewQXmlEntityResolverFromPointer(ptr unsafe.Pointer) *QXmlEntityResolver {
	var n = new(QXmlEntityResolver)
	n.SetPointer(ptr)
	return n
}

func newQXmlEntityResolverFromPointer(ptr unsafe.Pointer) *QXmlEntityResolver {
	var n = NewQXmlEntityResolverFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlEntityResolver_") {
		n.SetObjectNameAbs("QXmlEntityResolver_" + qt.Identifier())
	}
	return n
}

//export callbackQXmlEntityResolver_ErrorString
func callbackQXmlEntityResolver_ErrorString(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QXmlEntityResolver::errorString")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorString"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QXmlEntityResolver) ConnectErrorString(f func() string) {
	defer qt.Recovering("connect QXmlEntityResolver::errorString")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "errorString", f)
	}
}

func (ptr *QXmlEntityResolver) DisconnectErrorString() {
	defer qt.Recovering("disconnect QXmlEntityResolver::errorString")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "errorString")
	}
}

func (ptr *QXmlEntityResolver) ErrorString() string {
	defer qt.Recovering("QXmlEntityResolver::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlEntityResolver_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQXmlEntityResolver_ResolveEntity
func callbackQXmlEntityResolver_ResolveEntity(ptr unsafe.Pointer, ptrName *C.char, publicId *C.char, systemId *C.char, ret unsafe.Pointer) C.int {
	defer qt.Recovering("callback QXmlEntityResolver::resolveEntity")

	if signal := qt.GetSignal(C.GoString(ptrName), "resolveEntity"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string, *QXmlInputSource) bool)(C.GoString(publicId), C.GoString(systemId), NewQXmlInputSourceFromPointer(ret))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlEntityResolver) ConnectResolveEntity(f func(publicId string, systemId string, ret *QXmlInputSource) bool) {
	defer qt.Recovering("connect QXmlEntityResolver::resolveEntity")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "resolveEntity", f)
	}
}

func (ptr *QXmlEntityResolver) DisconnectResolveEntity(publicId string, systemId string, ret QXmlInputSource_ITF) {
	defer qt.Recovering("disconnect QXmlEntityResolver::resolveEntity")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "resolveEntity")
	}
}

func (ptr *QXmlEntityResolver) ResolveEntity(publicId string, systemId string, ret QXmlInputSource_ITF) bool {
	defer qt.Recovering("QXmlEntityResolver::resolveEntity")

	if ptr.Pointer() != nil {

	}
	return false
}

func (ptr *QXmlEntityResolver) DestroyQXmlEntityResolver() {
	defer qt.Recovering("QXmlEntityResolver::~QXmlEntityResolver")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QXmlEntityResolver_DestroyQXmlEntityResolver(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlEntityResolver) ObjectNameAbs() string {
	defer qt.Recovering("QXmlEntityResolver::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlEntityResolver_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlEntityResolver) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlEntityResolver::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlEntityResolver_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlErrorHandler struct {
	ptr unsafe.Pointer
}

type QXmlErrorHandler_ITF interface {
	QXmlErrorHandler_PTR() *QXmlErrorHandler
}

func (p *QXmlErrorHandler) QXmlErrorHandler_PTR() *QXmlErrorHandler {
	return p
}

func (p *QXmlErrorHandler) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlErrorHandler) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlErrorHandler(ptr QXmlErrorHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlErrorHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlErrorHandlerFromPointer(ptr unsafe.Pointer) *QXmlErrorHandler {
	var n = new(QXmlErrorHandler)
	n.SetPointer(ptr)
	return n
}

func newQXmlErrorHandlerFromPointer(ptr unsafe.Pointer) *QXmlErrorHandler {
	var n = NewQXmlErrorHandlerFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlErrorHandler_") {
		n.SetObjectNameAbs("QXmlErrorHandler_" + qt.Identifier())
	}
	return n
}

//export callbackQXmlErrorHandler_Error
func callbackQXmlErrorHandler_Error(ptr unsafe.Pointer, ptrName *C.char, exception unsafe.Pointer) C.int {
	defer qt.Recovering("callback QXmlErrorHandler::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QXmlParseException) bool)(NewQXmlParseExceptionFromPointer(exception))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlErrorHandler) ConnectError(f func(exception *QXmlParseException) bool) {
	defer qt.Recovering("connect QXmlErrorHandler::error")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "error", f)
	}
}

func (ptr *QXmlErrorHandler) DisconnectError(exception QXmlParseException_ITF) {
	defer qt.Recovering("disconnect QXmlErrorHandler::error")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "error")
	}
}

func (ptr *QXmlErrorHandler) Error(exception QXmlParseException_ITF) bool {
	defer qt.Recovering("QXmlErrorHandler::error")

	if ptr.Pointer() != nil {
		return C.QXmlErrorHandler_Error(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

//export callbackQXmlErrorHandler_ErrorString
func callbackQXmlErrorHandler_ErrorString(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QXmlErrorHandler::errorString")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorString"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QXmlErrorHandler) ConnectErrorString(f func() string) {
	defer qt.Recovering("connect QXmlErrorHandler::errorString")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "errorString", f)
	}
}

func (ptr *QXmlErrorHandler) DisconnectErrorString() {
	defer qt.Recovering("disconnect QXmlErrorHandler::errorString")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "errorString")
	}
}

func (ptr *QXmlErrorHandler) ErrorString() string {
	defer qt.Recovering("QXmlErrorHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlErrorHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQXmlErrorHandler_FatalError
func callbackQXmlErrorHandler_FatalError(ptr unsafe.Pointer, ptrName *C.char, exception unsafe.Pointer) C.int {
	defer qt.Recovering("callback QXmlErrorHandler::fatalError")

	if signal := qt.GetSignal(C.GoString(ptrName), "fatalError"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QXmlParseException) bool)(NewQXmlParseExceptionFromPointer(exception))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlErrorHandler) ConnectFatalError(f func(exception *QXmlParseException) bool) {
	defer qt.Recovering("connect QXmlErrorHandler::fatalError")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "fatalError", f)
	}
}

func (ptr *QXmlErrorHandler) DisconnectFatalError(exception QXmlParseException_ITF) {
	defer qt.Recovering("disconnect QXmlErrorHandler::fatalError")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "fatalError")
	}
}

func (ptr *QXmlErrorHandler) FatalError(exception QXmlParseException_ITF) bool {
	defer qt.Recovering("QXmlErrorHandler::fatalError")

	if ptr.Pointer() != nil {
		return C.QXmlErrorHandler_FatalError(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

//export callbackQXmlErrorHandler_Warning
func callbackQXmlErrorHandler_Warning(ptr unsafe.Pointer, ptrName *C.char, exception unsafe.Pointer) C.int {
	defer qt.Recovering("callback QXmlErrorHandler::warning")

	if signal := qt.GetSignal(C.GoString(ptrName), "warning"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QXmlParseException) bool)(NewQXmlParseExceptionFromPointer(exception))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlErrorHandler) ConnectWarning(f func(exception *QXmlParseException) bool) {
	defer qt.Recovering("connect QXmlErrorHandler::warning")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "warning", f)
	}
}

func (ptr *QXmlErrorHandler) DisconnectWarning(exception QXmlParseException_ITF) {
	defer qt.Recovering("disconnect QXmlErrorHandler::warning")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "warning")
	}
}

func (ptr *QXmlErrorHandler) Warning(exception QXmlParseException_ITF) bool {
	defer qt.Recovering("QXmlErrorHandler::warning")

	if ptr.Pointer() != nil {
		return C.QXmlErrorHandler_Warning(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

func (ptr *QXmlErrorHandler) DestroyQXmlErrorHandler() {
	defer qt.Recovering("QXmlErrorHandler::~QXmlErrorHandler")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QXmlErrorHandler_DestroyQXmlErrorHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlErrorHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlErrorHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlErrorHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlErrorHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlErrorHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlErrorHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlInputSource struct {
	ptr unsafe.Pointer
}

type QXmlInputSource_ITF interface {
	QXmlInputSource_PTR() *QXmlInputSource
}

func (p *QXmlInputSource) QXmlInputSource_PTR() *QXmlInputSource {
	return p
}

func (p *QXmlInputSource) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlInputSource) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlInputSource(ptr QXmlInputSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlInputSource_PTR().Pointer()
	}
	return nil
}

func NewQXmlInputSourceFromPointer(ptr unsafe.Pointer) *QXmlInputSource {
	var n = new(QXmlInputSource)
	n.SetPointer(ptr)
	return n
}

func newQXmlInputSourceFromPointer(ptr unsafe.Pointer) *QXmlInputSource {
	var n = NewQXmlInputSourceFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlInputSource_") {
		n.SetObjectNameAbs("QXmlInputSource_" + qt.Identifier())
	}
	return n
}

func NewQXmlInputSource() *QXmlInputSource {
	defer qt.Recovering("QXmlInputSource::QXmlInputSource")

	return newQXmlInputSourceFromPointer(C.QXmlInputSource_NewQXmlInputSource())
}

func NewQXmlInputSource2(dev core.QIODevice_ITF) *QXmlInputSource {
	defer qt.Recovering("QXmlInputSource::QXmlInputSource")

	return newQXmlInputSourceFromPointer(C.QXmlInputSource_NewQXmlInputSource2(core.PointerFromQIODevice(dev)))
}

//export callbackQXmlInputSource_Data
func callbackQXmlInputSource_Data(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QXmlInputSource::data")

	if signal := qt.GetSignal(C.GoString(ptrName), "data"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQXmlInputSourceFromPointer(ptr).DataDefault())
}

func (ptr *QXmlInputSource) ConnectData(f func() string) {
	defer qt.Recovering("connect QXmlInputSource::data")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "data", f)
	}
}

func (ptr *QXmlInputSource) DisconnectData() {
	defer qt.Recovering("disconnect QXmlInputSource::data")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "data")
	}
}

func (ptr *QXmlInputSource) Data() string {
	defer qt.Recovering("QXmlInputSource::data")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlInputSource_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlInputSource) DataDefault() string {
	defer qt.Recovering("QXmlInputSource::data")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlInputSource_DataDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQXmlInputSource_FetchData
func callbackQXmlInputSource_FetchData(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlInputSource::fetchData")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchData"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlInputSourceFromPointer(ptr).FetchDataDefault()
	}
}

func (ptr *QXmlInputSource) ConnectFetchData(f func()) {
	defer qt.Recovering("connect QXmlInputSource::fetchData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "fetchData", f)
	}
}

func (ptr *QXmlInputSource) DisconnectFetchData() {
	defer qt.Recovering("disconnect QXmlInputSource::fetchData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "fetchData")
	}
}

func (ptr *QXmlInputSource) FetchData() {
	defer qt.Recovering("QXmlInputSource::fetchData")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_FetchData(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) FetchDataDefault() {
	defer qt.Recovering("QXmlInputSource::fetchData")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_FetchDataDefault(ptr.Pointer())
	}
}

//export callbackQXmlInputSource_FromRawData
func callbackQXmlInputSource_FromRawData(ptr unsafe.Pointer, ptrName *C.char, data *C.char, beginning C.int) *C.char {
	defer qt.Recovering("callback QXmlInputSource::fromRawData")

	if signal := qt.GetSignal(C.GoString(ptrName), "fromRawData"); signal != nil {
		return C.CString(signal.(func(string, bool) string)(C.GoString(data), int(beginning) != 0))
	}

	return C.CString(NewQXmlInputSourceFromPointer(ptr).FromRawDataDefault(C.GoString(data), int(beginning) != 0))
}

func (ptr *QXmlInputSource) ConnectFromRawData(f func(data string, beginning bool) string) {
	defer qt.Recovering("connect QXmlInputSource::fromRawData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "fromRawData", f)
	}
}

func (ptr *QXmlInputSource) DisconnectFromRawData() {
	defer qt.Recovering("disconnect QXmlInputSource::fromRawData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "fromRawData")
	}
}

func (ptr *QXmlInputSource) FromRawData(data string, beginning bool) string {
	defer qt.Recovering("QXmlInputSource::fromRawData")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlInputSource_FromRawData(ptr.Pointer(), C.CString(data), C.int(qt.GoBoolToInt(beginning))))
	}
	return ""
}

func (ptr *QXmlInputSource) FromRawDataDefault(data string, beginning bool) string {
	defer qt.Recovering("QXmlInputSource::fromRawData")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlInputSource_FromRawDataDefault(ptr.Pointer(), C.CString(data), C.int(qt.GoBoolToInt(beginning))))
	}
	return ""
}

//export callbackQXmlInputSource_Next
func callbackQXmlInputSource_Next(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QXmlInputSource::next")

	if signal := qt.GetSignal(C.GoString(ptrName), "next"); signal != nil {
		return core.PointerFromQChar(signal.(func() *core.QChar)())
	}

	return core.PointerFromQChar(NewQXmlInputSourceFromPointer(ptr).NextDefault())
}

func (ptr *QXmlInputSource) ConnectNext(f func() *core.QChar) {
	defer qt.Recovering("connect QXmlInputSource::next")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "next", f)
	}
}

func (ptr *QXmlInputSource) DisconnectNext() {
	defer qt.Recovering("disconnect QXmlInputSource::next")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "next")
	}
}

func (ptr *QXmlInputSource) Next() *core.QChar {
	defer qt.Recovering("QXmlInputSource::next")

	if ptr.Pointer() != nil {

	}
	return nil
}

func (ptr *QXmlInputSource) NextDefault() *core.QChar {
	defer qt.Recovering("QXmlInputSource::next")

	if ptr.Pointer() != nil {

	}
	return nil
}

//export callbackQXmlInputSource_Reset
func callbackQXmlInputSource_Reset(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlInputSource::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlInputSourceFromPointer(ptr).ResetDefault()
	}
}

func (ptr *QXmlInputSource) ConnectReset(f func()) {
	defer qt.Recovering("connect QXmlInputSource::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "reset", f)
	}
}

func (ptr *QXmlInputSource) DisconnectReset() {
	defer qt.Recovering("disconnect QXmlInputSource::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "reset")
	}
}

func (ptr *QXmlInputSource) Reset() {
	defer qt.Recovering("QXmlInputSource::reset")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_Reset(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) ResetDefault() {
	defer qt.Recovering("QXmlInputSource::reset")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_ResetDefault(ptr.Pointer())
	}
}

//export callbackQXmlInputSource_SetData2
func callbackQXmlInputSource_SetData2(ptr unsafe.Pointer, ptrName *C.char, dat *C.char) {
	defer qt.Recovering("callback QXmlInputSource::setData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setData2"); signal != nil {
		signal.(func(string))(C.GoString(dat))
	} else {
		NewQXmlInputSourceFromPointer(ptr).SetData2Default(C.GoString(dat))
	}
}

func (ptr *QXmlInputSource) ConnectSetData2(f func(dat string)) {
	defer qt.Recovering("connect QXmlInputSource::setData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setData2", f)
	}
}

func (ptr *QXmlInputSource) DisconnectSetData2() {
	defer qt.Recovering("disconnect QXmlInputSource::setData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setData2")
	}
}

func (ptr *QXmlInputSource) SetData2(dat string) {
	defer qt.Recovering("QXmlInputSource::setData")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetData2(ptr.Pointer(), C.CString(dat))
	}
}

func (ptr *QXmlInputSource) SetData2Default(dat string) {
	defer qt.Recovering("QXmlInputSource::setData")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetData2Default(ptr.Pointer(), C.CString(dat))
	}
}

//export callbackQXmlInputSource_SetData
func callbackQXmlInputSource_SetData(ptr unsafe.Pointer, ptrName *C.char, dat *C.char) {
	defer qt.Recovering("callback QXmlInputSource::setData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setData"); signal != nil {
		signal.(func(string))(C.GoString(dat))
	} else {
		NewQXmlInputSourceFromPointer(ptr).SetDataDefault(C.GoString(dat))
	}
}

func (ptr *QXmlInputSource) ConnectSetData(f func(dat string)) {
	defer qt.Recovering("connect QXmlInputSource::setData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setData", f)
	}
}

func (ptr *QXmlInputSource) DisconnectSetData() {
	defer qt.Recovering("disconnect QXmlInputSource::setData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setData")
	}
}

func (ptr *QXmlInputSource) SetData(dat string) {
	defer qt.Recovering("QXmlInputSource::setData")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetData(ptr.Pointer(), C.CString(dat))
	}
}

func (ptr *QXmlInputSource) SetDataDefault(dat string) {
	defer qt.Recovering("QXmlInputSource::setData")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetDataDefault(ptr.Pointer(), C.CString(dat))
	}
}

func (ptr *QXmlInputSource) DestroyQXmlInputSource() {
	defer qt.Recovering("QXmlInputSource::~QXmlInputSource")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QXmlInputSource_DestroyQXmlInputSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlInputSource) ObjectNameAbs() string {
	defer qt.Recovering("QXmlInputSource::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlInputSource_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlInputSource) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlInputSource::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlLexicalHandler struct {
	ptr unsafe.Pointer
}

type QXmlLexicalHandler_ITF interface {
	QXmlLexicalHandler_PTR() *QXmlLexicalHandler
}

func (p *QXmlLexicalHandler) QXmlLexicalHandler_PTR() *QXmlLexicalHandler {
	return p
}

func (p *QXmlLexicalHandler) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlLexicalHandler) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlLexicalHandler(ptr QXmlLexicalHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlLexicalHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlLexicalHandlerFromPointer(ptr unsafe.Pointer) *QXmlLexicalHandler {
	var n = new(QXmlLexicalHandler)
	n.SetPointer(ptr)
	return n
}

func newQXmlLexicalHandlerFromPointer(ptr unsafe.Pointer) *QXmlLexicalHandler {
	var n = NewQXmlLexicalHandlerFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlLexicalHandler_") {
		n.SetObjectNameAbs("QXmlLexicalHandler_" + qt.Identifier())
	}
	return n
}

//export callbackQXmlLexicalHandler_Comment
func callbackQXmlLexicalHandler_Comment(ptr unsafe.Pointer, ptrName *C.char, ch *C.char) C.int {
	defer qt.Recovering("callback QXmlLexicalHandler::comment")

	if signal := qt.GetSignal(C.GoString(ptrName), "comment"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(ch))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlLexicalHandler) ConnectComment(f func(ch string) bool) {
	defer qt.Recovering("connect QXmlLexicalHandler::comment")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "comment", f)
	}
}

func (ptr *QXmlLexicalHandler) DisconnectComment(ch string) {
	defer qt.Recovering("disconnect QXmlLexicalHandler::comment")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "comment")
	}
}

func (ptr *QXmlLexicalHandler) Comment(ch string) bool {
	defer qt.Recovering("QXmlLexicalHandler::comment")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_Comment(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

//export callbackQXmlLexicalHandler_EndCDATA
func callbackQXmlLexicalHandler_EndCDATA(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QXmlLexicalHandler::endCDATA")

	if signal := qt.GetSignal(C.GoString(ptrName), "endCDATA"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlLexicalHandler) ConnectEndCDATA(f func() bool) {
	defer qt.Recovering("connect QXmlLexicalHandler::endCDATA")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endCDATA", f)
	}
}

func (ptr *QXmlLexicalHandler) DisconnectEndCDATA() {
	defer qt.Recovering("disconnect QXmlLexicalHandler::endCDATA")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endCDATA")
	}
}

func (ptr *QXmlLexicalHandler) EndCDATA() bool {
	defer qt.Recovering("QXmlLexicalHandler::endCDATA")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_EndCDATA(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQXmlLexicalHandler_EndDTD
func callbackQXmlLexicalHandler_EndDTD(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QXmlLexicalHandler::endDTD")

	if signal := qt.GetSignal(C.GoString(ptrName), "endDTD"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlLexicalHandler) ConnectEndDTD(f func() bool) {
	defer qt.Recovering("connect QXmlLexicalHandler::endDTD")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endDTD", f)
	}
}

func (ptr *QXmlLexicalHandler) DisconnectEndDTD() {
	defer qt.Recovering("disconnect QXmlLexicalHandler::endDTD")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endDTD")
	}
}

func (ptr *QXmlLexicalHandler) EndDTD() bool {
	defer qt.Recovering("QXmlLexicalHandler::endDTD")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_EndDTD(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQXmlLexicalHandler_EndEntity
func callbackQXmlLexicalHandler_EndEntity(ptr unsafe.Pointer, ptrName *C.char, name *C.char) C.int {
	defer qt.Recovering("callback QXmlLexicalHandler::endEntity")

	if signal := qt.GetSignal(C.GoString(ptrName), "endEntity"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlLexicalHandler) ConnectEndEntity(f func(name string) bool) {
	defer qt.Recovering("connect QXmlLexicalHandler::endEntity")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endEntity", f)
	}
}

func (ptr *QXmlLexicalHandler) DisconnectEndEntity(name string) {
	defer qt.Recovering("disconnect QXmlLexicalHandler::endEntity")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endEntity")
	}
}

func (ptr *QXmlLexicalHandler) EndEntity(name string) bool {
	defer qt.Recovering("QXmlLexicalHandler::endEntity")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_EndEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

//export callbackQXmlLexicalHandler_ErrorString
func callbackQXmlLexicalHandler_ErrorString(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QXmlLexicalHandler::errorString")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorString"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QXmlLexicalHandler) ConnectErrorString(f func() string) {
	defer qt.Recovering("connect QXmlLexicalHandler::errorString")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "errorString", f)
	}
}

func (ptr *QXmlLexicalHandler) DisconnectErrorString() {
	defer qt.Recovering("disconnect QXmlLexicalHandler::errorString")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "errorString")
	}
}

func (ptr *QXmlLexicalHandler) ErrorString() string {
	defer qt.Recovering("QXmlLexicalHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlLexicalHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQXmlLexicalHandler_StartCDATA
func callbackQXmlLexicalHandler_StartCDATA(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QXmlLexicalHandler::startCDATA")

	if signal := qt.GetSignal(C.GoString(ptrName), "startCDATA"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlLexicalHandler) ConnectStartCDATA(f func() bool) {
	defer qt.Recovering("connect QXmlLexicalHandler::startCDATA")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startCDATA", f)
	}
}

func (ptr *QXmlLexicalHandler) DisconnectStartCDATA() {
	defer qt.Recovering("disconnect QXmlLexicalHandler::startCDATA")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startCDATA")
	}
}

func (ptr *QXmlLexicalHandler) StartCDATA() bool {
	defer qt.Recovering("QXmlLexicalHandler::startCDATA")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_StartCDATA(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQXmlLexicalHandler_StartDTD
func callbackQXmlLexicalHandler_StartDTD(ptr unsafe.Pointer, ptrName *C.char, name *C.char, publicId *C.char, systemId *C.char) C.int {
	defer qt.Recovering("callback QXmlLexicalHandler::startDTD")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDTD"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, string, string) bool)(C.GoString(name), C.GoString(publicId), C.GoString(systemId))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlLexicalHandler) ConnectStartDTD(f func(name string, publicId string, systemId string) bool) {
	defer qt.Recovering("connect QXmlLexicalHandler::startDTD")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startDTD", f)
	}
}

func (ptr *QXmlLexicalHandler) DisconnectStartDTD(name string, publicId string, systemId string) {
	defer qt.Recovering("disconnect QXmlLexicalHandler::startDTD")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startDTD")
	}
}

func (ptr *QXmlLexicalHandler) StartDTD(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlLexicalHandler::startDTD")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_StartDTD(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

//export callbackQXmlLexicalHandler_StartEntity
func callbackQXmlLexicalHandler_StartEntity(ptr unsafe.Pointer, ptrName *C.char, name *C.char) C.int {
	defer qt.Recovering("callback QXmlLexicalHandler::startEntity")

	if signal := qt.GetSignal(C.GoString(ptrName), "startEntity"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlLexicalHandler) ConnectStartEntity(f func(name string) bool) {
	defer qt.Recovering("connect QXmlLexicalHandler::startEntity")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startEntity", f)
	}
}

func (ptr *QXmlLexicalHandler) DisconnectStartEntity(name string) {
	defer qt.Recovering("disconnect QXmlLexicalHandler::startEntity")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startEntity")
	}
}

func (ptr *QXmlLexicalHandler) StartEntity(name string) bool {
	defer qt.Recovering("QXmlLexicalHandler::startEntity")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_StartEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) DestroyQXmlLexicalHandler() {
	defer qt.Recovering("QXmlLexicalHandler::~QXmlLexicalHandler")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QXmlLexicalHandler_DestroyQXmlLexicalHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlLexicalHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlLexicalHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlLexicalHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlLexicalHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlLexicalHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlLexicalHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlLocator struct {
	ptr unsafe.Pointer
}

type QXmlLocator_ITF interface {
	QXmlLocator_PTR() *QXmlLocator
}

func (p *QXmlLocator) QXmlLocator_PTR() *QXmlLocator {
	return p
}

func (p *QXmlLocator) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlLocator) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlLocator(ptr QXmlLocator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlLocator_PTR().Pointer()
	}
	return nil
}

func NewQXmlLocatorFromPointer(ptr unsafe.Pointer) *QXmlLocator {
	var n = new(QXmlLocator)
	n.SetPointer(ptr)
	return n
}

func newQXmlLocatorFromPointer(ptr unsafe.Pointer) *QXmlLocator {
	var n = NewQXmlLocatorFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlLocator_") {
		n.SetObjectNameAbs("QXmlLocator_" + qt.Identifier())
	}
	return n
}

func NewQXmlLocator() *QXmlLocator {
	defer qt.Recovering("QXmlLocator::QXmlLocator")

	return newQXmlLocatorFromPointer(C.QXmlLocator_NewQXmlLocator())
}

//export callbackQXmlLocator_ColumnNumber
func callbackQXmlLocator_ColumnNumber(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QXmlLocator::columnNumber")

	if signal := qt.GetSignal(C.GoString(ptrName), "columnNumber"); signal != nil {
		return C.int(signal.(func() int)())
	}

	return C.int(0)
}

func (ptr *QXmlLocator) ConnectColumnNumber(f func() int) {
	defer qt.Recovering("connect QXmlLocator::columnNumber")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "columnNumber", f)
	}
}

func (ptr *QXmlLocator) DisconnectColumnNumber() {
	defer qt.Recovering("disconnect QXmlLocator::columnNumber")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "columnNumber")
	}
}

func (ptr *QXmlLocator) ColumnNumber() int {
	defer qt.Recovering("QXmlLocator::columnNumber")

	if ptr.Pointer() != nil {
		return int(C.QXmlLocator_ColumnNumber(ptr.Pointer()))
	}
	return 0
}

//export callbackQXmlLocator_LineNumber
func callbackQXmlLocator_LineNumber(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QXmlLocator::lineNumber")

	if signal := qt.GetSignal(C.GoString(ptrName), "lineNumber"); signal != nil {
		return C.int(signal.(func() int)())
	}

	return C.int(0)
}

func (ptr *QXmlLocator) ConnectLineNumber(f func() int) {
	defer qt.Recovering("connect QXmlLocator::lineNumber")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "lineNumber", f)
	}
}

func (ptr *QXmlLocator) DisconnectLineNumber() {
	defer qt.Recovering("disconnect QXmlLocator::lineNumber")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "lineNumber")
	}
}

func (ptr *QXmlLocator) LineNumber() int {
	defer qt.Recovering("QXmlLocator::lineNumber")

	if ptr.Pointer() != nil {
		return int(C.QXmlLocator_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlLocator) DestroyQXmlLocator() {
	defer qt.Recovering("QXmlLocator::~QXmlLocator")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QXmlLocator_DestroyQXmlLocator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlLocator) ObjectNameAbs() string {
	defer qt.Recovering("QXmlLocator::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlLocator_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlLocator) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlLocator::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlLocator_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlNamespaceSupport struct {
	ptr unsafe.Pointer
}

type QXmlNamespaceSupport_ITF interface {
	QXmlNamespaceSupport_PTR() *QXmlNamespaceSupport
}

func (p *QXmlNamespaceSupport) QXmlNamespaceSupport_PTR() *QXmlNamespaceSupport {
	return p
}

func (p *QXmlNamespaceSupport) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlNamespaceSupport) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlNamespaceSupport(ptr QXmlNamespaceSupport_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNamespaceSupport_PTR().Pointer()
	}
	return nil
}

func NewQXmlNamespaceSupportFromPointer(ptr unsafe.Pointer) *QXmlNamespaceSupport {
	var n = new(QXmlNamespaceSupport)
	n.SetPointer(ptr)
	return n
}

func newQXmlNamespaceSupportFromPointer(ptr unsafe.Pointer) *QXmlNamespaceSupport {
	var n = NewQXmlNamespaceSupportFromPointer(ptr)
	return n
}

func NewQXmlNamespaceSupport() *QXmlNamespaceSupport {
	defer qt.Recovering("QXmlNamespaceSupport::QXmlNamespaceSupport")

	return newQXmlNamespaceSupportFromPointer(C.QXmlNamespaceSupport_NewQXmlNamespaceSupport())
}

func (ptr *QXmlNamespaceSupport) PopContext() {
	defer qt.Recovering("QXmlNamespaceSupport::popContext")

	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_PopContext(ptr.Pointer())
	}
}

func (ptr *QXmlNamespaceSupport) Prefix(uri string) string {
	defer qt.Recovering("QXmlNamespaceSupport::prefix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlNamespaceSupport_Prefix(ptr.Pointer(), C.CString(uri)))
	}
	return ""
}

func (ptr *QXmlNamespaceSupport) Prefixes() []string {
	defer qt.Recovering("QXmlNamespaceSupport::prefixes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QXmlNamespaceSupport_Prefixes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QXmlNamespaceSupport) Prefixes2(uri string) []string {
	defer qt.Recovering("QXmlNamespaceSupport::prefixes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QXmlNamespaceSupport_Prefixes2(ptr.Pointer(), C.CString(uri))), "|")
	}
	return make([]string, 0)
}

func (ptr *QXmlNamespaceSupport) ProcessName(qname string, isAttribute bool, nsuri string, localname string) {
	defer qt.Recovering("QXmlNamespaceSupport::processName")

	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_ProcessName(ptr.Pointer(), C.CString(qname), C.int(qt.GoBoolToInt(isAttribute)), C.CString(nsuri), C.CString(localname))
	}
}

func (ptr *QXmlNamespaceSupport) PushContext() {
	defer qt.Recovering("QXmlNamespaceSupport::pushContext")

	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_PushContext(ptr.Pointer())
	}
}

func (ptr *QXmlNamespaceSupport) Reset() {
	defer qt.Recovering("QXmlNamespaceSupport::reset")

	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_Reset(ptr.Pointer())
	}
}

func (ptr *QXmlNamespaceSupport) SetPrefix(pre string, uri string) {
	defer qt.Recovering("QXmlNamespaceSupport::setPrefix")

	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_SetPrefix(ptr.Pointer(), C.CString(pre), C.CString(uri))
	}
}

func (ptr *QXmlNamespaceSupport) SplitName(qname string, prefix string, localname string) {
	defer qt.Recovering("QXmlNamespaceSupport::splitName")

	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_SplitName(ptr.Pointer(), C.CString(qname), C.CString(prefix), C.CString(localname))
	}
}

func (ptr *QXmlNamespaceSupport) Uri(prefix string) string {
	defer qt.Recovering("QXmlNamespaceSupport::uri")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlNamespaceSupport_Uri(ptr.Pointer(), C.CString(prefix)))
	}
	return ""
}

func (ptr *QXmlNamespaceSupport) DestroyQXmlNamespaceSupport() {
	defer qt.Recovering("QXmlNamespaceSupport::~QXmlNamespaceSupport")

	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_DestroyQXmlNamespaceSupport(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXmlParseException struct {
	ptr unsafe.Pointer
}

type QXmlParseException_ITF interface {
	QXmlParseException_PTR() *QXmlParseException
}

func (p *QXmlParseException) QXmlParseException_PTR() *QXmlParseException {
	return p
}

func (p *QXmlParseException) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlParseException) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlParseException(ptr QXmlParseException_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlParseException_PTR().Pointer()
	}
	return nil
}

func NewQXmlParseExceptionFromPointer(ptr unsafe.Pointer) *QXmlParseException {
	var n = new(QXmlParseException)
	n.SetPointer(ptr)
	return n
}

func newQXmlParseExceptionFromPointer(ptr unsafe.Pointer) *QXmlParseException {
	var n = NewQXmlParseExceptionFromPointer(ptr)
	return n
}

func NewQXmlParseException(name string, c int, l int, p string, s string) *QXmlParseException {
	defer qt.Recovering("QXmlParseException::QXmlParseException")

	return newQXmlParseExceptionFromPointer(C.QXmlParseException_NewQXmlParseException(C.CString(name), C.int(c), C.int(l), C.CString(p), C.CString(s)))
}

func NewQXmlParseException2(other QXmlParseException_ITF) *QXmlParseException {
	defer qt.Recovering("QXmlParseException::QXmlParseException")

	return newQXmlParseExceptionFromPointer(C.QXmlParseException_NewQXmlParseException2(PointerFromQXmlParseException(other)))
}

func (ptr *QXmlParseException) ColumnNumber() int {
	defer qt.Recovering("QXmlParseException::columnNumber")

	if ptr.Pointer() != nil {
		return int(C.QXmlParseException_ColumnNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlParseException) LineNumber() int {
	defer qt.Recovering("QXmlParseException::lineNumber")

	if ptr.Pointer() != nil {
		return int(C.QXmlParseException_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlParseException) Message() string {
	defer qt.Recovering("QXmlParseException::message")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlParseException_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlParseException) PublicId() string {
	defer qt.Recovering("QXmlParseException::publicId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlParseException_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlParseException) SystemId() string {
	defer qt.Recovering("QXmlParseException::systemId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlParseException_SystemId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlParseException) DestroyQXmlParseException() {
	defer qt.Recovering("QXmlParseException::~QXmlParseException")

	if ptr.Pointer() != nil {
		C.QXmlParseException_DestroyQXmlParseException(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXmlReader struct {
	ptr unsafe.Pointer
}

type QXmlReader_ITF interface {
	QXmlReader_PTR() *QXmlReader
}

func (p *QXmlReader) QXmlReader_PTR() *QXmlReader {
	return p
}

func (p *QXmlReader) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlReader) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlReader(ptr QXmlReader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlReader_PTR().Pointer()
	}
	return nil
}

func NewQXmlReaderFromPointer(ptr unsafe.Pointer) *QXmlReader {
	var n = new(QXmlReader)
	n.SetPointer(ptr)
	return n
}

func newQXmlReaderFromPointer(ptr unsafe.Pointer) *QXmlReader {
	var n = NewQXmlReaderFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlReader_") {
		n.SetObjectNameAbs("QXmlReader_" + qt.Identifier())
	}
	return n
}

//export callbackQXmlReader_DTDHandler
func callbackQXmlReader_DTDHandler(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QXmlReader::DTDHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "DTDHandler"); signal != nil {
		return PointerFromQXmlDTDHandler(signal.(func() *QXmlDTDHandler)())
	}

	return PointerFromQXmlDTDHandler(nil)
}

func (ptr *QXmlReader) ConnectDTDHandler(f func() *QXmlDTDHandler) {
	defer qt.Recovering("connect QXmlReader::DTDHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "DTDHandler", f)
	}
}

func (ptr *QXmlReader) DisconnectDTDHandler() {
	defer qt.Recovering("disconnect QXmlReader::DTDHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "DTDHandler")
	}
}

func (ptr *QXmlReader) DTDHandler() *QXmlDTDHandler {
	defer qt.Recovering("QXmlReader::DTDHandler")

	if ptr.Pointer() != nil {
		return NewQXmlDTDHandlerFromPointer(C.QXmlReader_DTDHandler(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlReader_ContentHandler
func callbackQXmlReader_ContentHandler(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QXmlReader::contentHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentHandler"); signal != nil {
		return PointerFromQXmlContentHandler(signal.(func() *QXmlContentHandler)())
	}

	return PointerFromQXmlContentHandler(nil)
}

func (ptr *QXmlReader) ConnectContentHandler(f func() *QXmlContentHandler) {
	defer qt.Recovering("connect QXmlReader::contentHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "contentHandler", f)
	}
}

func (ptr *QXmlReader) DisconnectContentHandler() {
	defer qt.Recovering("disconnect QXmlReader::contentHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "contentHandler")
	}
}

func (ptr *QXmlReader) ContentHandler() *QXmlContentHandler {
	defer qt.Recovering("QXmlReader::contentHandler")

	if ptr.Pointer() != nil {
		return NewQXmlContentHandlerFromPointer(C.QXmlReader_ContentHandler(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlReader_DeclHandler
func callbackQXmlReader_DeclHandler(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QXmlReader::declHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "declHandler"); signal != nil {
		return PointerFromQXmlDeclHandler(signal.(func() *QXmlDeclHandler)())
	}

	return PointerFromQXmlDeclHandler(nil)
}

func (ptr *QXmlReader) ConnectDeclHandler(f func() *QXmlDeclHandler) {
	defer qt.Recovering("connect QXmlReader::declHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "declHandler", f)
	}
}

func (ptr *QXmlReader) DisconnectDeclHandler() {
	defer qt.Recovering("disconnect QXmlReader::declHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "declHandler")
	}
}

func (ptr *QXmlReader) DeclHandler() *QXmlDeclHandler {
	defer qt.Recovering("QXmlReader::declHandler")

	if ptr.Pointer() != nil {
		return NewQXmlDeclHandlerFromPointer(C.QXmlReader_DeclHandler(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlReader_EntityResolver
func callbackQXmlReader_EntityResolver(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QXmlReader::entityResolver")

	if signal := qt.GetSignal(C.GoString(ptrName), "entityResolver"); signal != nil {
		return PointerFromQXmlEntityResolver(signal.(func() *QXmlEntityResolver)())
	}

	return PointerFromQXmlEntityResolver(nil)
}

func (ptr *QXmlReader) ConnectEntityResolver(f func() *QXmlEntityResolver) {
	defer qt.Recovering("connect QXmlReader::entityResolver")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "entityResolver", f)
	}
}

func (ptr *QXmlReader) DisconnectEntityResolver() {
	defer qt.Recovering("disconnect QXmlReader::entityResolver")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "entityResolver")
	}
}

func (ptr *QXmlReader) EntityResolver() *QXmlEntityResolver {
	defer qt.Recovering("QXmlReader::entityResolver")

	if ptr.Pointer() != nil {
		return NewQXmlEntityResolverFromPointer(C.QXmlReader_EntityResolver(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlReader_ErrorHandler
func callbackQXmlReader_ErrorHandler(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QXmlReader::errorHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorHandler"); signal != nil {
		return PointerFromQXmlErrorHandler(signal.(func() *QXmlErrorHandler)())
	}

	return PointerFromQXmlErrorHandler(nil)
}

func (ptr *QXmlReader) ConnectErrorHandler(f func() *QXmlErrorHandler) {
	defer qt.Recovering("connect QXmlReader::errorHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "errorHandler", f)
	}
}

func (ptr *QXmlReader) DisconnectErrorHandler() {
	defer qt.Recovering("disconnect QXmlReader::errorHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "errorHandler")
	}
}

func (ptr *QXmlReader) ErrorHandler() *QXmlErrorHandler {
	defer qt.Recovering("QXmlReader::errorHandler")

	if ptr.Pointer() != nil {
		return NewQXmlErrorHandlerFromPointer(C.QXmlReader_ErrorHandler(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlReader_Feature
func callbackQXmlReader_Feature(ptr unsafe.Pointer, ptrName *C.char, name *C.char, ok C.int) C.int {
	defer qt.Recovering("callback QXmlReader::feature")

	if signal := qt.GetSignal(C.GoString(ptrName), "feature"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, bool) bool)(C.GoString(name), int(ok) != 0)))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlReader) ConnectFeature(f func(name string, ok bool) bool) {
	defer qt.Recovering("connect QXmlReader::feature")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "feature", f)
	}
}

func (ptr *QXmlReader) DisconnectFeature(name string, ok bool) {
	defer qt.Recovering("disconnect QXmlReader::feature")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "feature")
	}
}

func (ptr *QXmlReader) Feature(name string, ok bool) bool {
	defer qt.Recovering("QXmlReader::feature")

	if ptr.Pointer() != nil {
		return C.QXmlReader_Feature(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))) != 0
	}
	return false
}

//export callbackQXmlReader_HasFeature
func callbackQXmlReader_HasFeature(ptr unsafe.Pointer, ptrName *C.char, name *C.char) C.int {
	defer qt.Recovering("callback QXmlReader::hasFeature")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasFeature"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlReader) ConnectHasFeature(f func(name string) bool) {
	defer qt.Recovering("connect QXmlReader::hasFeature")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "hasFeature", f)
	}
}

func (ptr *QXmlReader) DisconnectHasFeature(name string) {
	defer qt.Recovering("disconnect QXmlReader::hasFeature")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "hasFeature")
	}
}

func (ptr *QXmlReader) HasFeature(name string) bool {
	defer qt.Recovering("QXmlReader::hasFeature")

	if ptr.Pointer() != nil {
		return C.QXmlReader_HasFeature(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

//export callbackQXmlReader_HasProperty
func callbackQXmlReader_HasProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char) C.int {
	defer qt.Recovering("callback QXmlReader::hasProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasProperty"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlReader) ConnectHasProperty(f func(name string) bool) {
	defer qt.Recovering("connect QXmlReader::hasProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "hasProperty", f)
	}
}

func (ptr *QXmlReader) DisconnectHasProperty(name string) {
	defer qt.Recovering("disconnect QXmlReader::hasProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "hasProperty")
	}
}

func (ptr *QXmlReader) HasProperty(name string) bool {
	defer qt.Recovering("QXmlReader::hasProperty")

	if ptr.Pointer() != nil {
		return C.QXmlReader_HasProperty(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

//export callbackQXmlReader_LexicalHandler
func callbackQXmlReader_LexicalHandler(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QXmlReader::lexicalHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "lexicalHandler"); signal != nil {
		return PointerFromQXmlLexicalHandler(signal.(func() *QXmlLexicalHandler)())
	}

	return PointerFromQXmlLexicalHandler(nil)
}

func (ptr *QXmlReader) ConnectLexicalHandler(f func() *QXmlLexicalHandler) {
	defer qt.Recovering("connect QXmlReader::lexicalHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "lexicalHandler", f)
	}
}

func (ptr *QXmlReader) DisconnectLexicalHandler() {
	defer qt.Recovering("disconnect QXmlReader::lexicalHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "lexicalHandler")
	}
}

func (ptr *QXmlReader) LexicalHandler() *QXmlLexicalHandler {
	defer qt.Recovering("QXmlReader::lexicalHandler")

	if ptr.Pointer() != nil {
		return NewQXmlLexicalHandlerFromPointer(C.QXmlReader_LexicalHandler(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlReader_Parse
func callbackQXmlReader_Parse(ptr unsafe.Pointer, ptrName *C.char, input unsafe.Pointer) C.int {
	defer qt.Recovering("callback QXmlReader::parse")

	if signal := qt.GetSignal(C.GoString(ptrName), "parse"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QXmlInputSource) bool)(NewQXmlInputSourceFromPointer(input))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QXmlReader) ConnectParse(f func(input *QXmlInputSource) bool) {
	defer qt.Recovering("connect QXmlReader::parse")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "parse", f)
	}
}

func (ptr *QXmlReader) DisconnectParse(input QXmlInputSource_ITF) {
	defer qt.Recovering("disconnect QXmlReader::parse")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "parse")
	}
}

func (ptr *QXmlReader) Parse(input QXmlInputSource_ITF) bool {
	defer qt.Recovering("QXmlReader::parse")

	if ptr.Pointer() != nil {
		return C.QXmlReader_Parse(ptr.Pointer(), PointerFromQXmlInputSource(input)) != 0
	}
	return false
}

//export callbackQXmlReader_Property
func callbackQXmlReader_Property(ptr unsafe.Pointer, ptrName *C.char, name *C.char, ok C.int) unsafe.Pointer {
	defer qt.Recovering("callback QXmlReader::property")

	if signal := qt.GetSignal(C.GoString(ptrName), "property"); signal != nil {
		return signal.(func(string, bool) unsafe.Pointer)(C.GoString(name), int(ok) != 0)
	}

	return nil
}

func (ptr *QXmlReader) ConnectProperty(f func(name string, ok bool) unsafe.Pointer) {
	defer qt.Recovering("connect QXmlReader::property")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "property", f)
	}
}

func (ptr *QXmlReader) DisconnectProperty(name string, ok bool) {
	defer qt.Recovering("disconnect QXmlReader::property")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "property")
	}
}

func (ptr *QXmlReader) Property(name string, ok bool) unsafe.Pointer {
	defer qt.Recovering("QXmlReader::property")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXmlReader_Property(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))))
	}
	return nil
}

//export callbackQXmlReader_SetContentHandler
func callbackQXmlReader_SetContentHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlReader::setContentHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setContentHandler"); signal != nil {
		signal.(func(*QXmlContentHandler))(NewQXmlContentHandlerFromPointer(handler))
	}

}

func (ptr *QXmlReader) ConnectSetContentHandler(f func(handler *QXmlContentHandler)) {
	defer qt.Recovering("connect QXmlReader::setContentHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setContentHandler", f)
	}
}

func (ptr *QXmlReader) DisconnectSetContentHandler(handler QXmlContentHandler_ITF) {
	defer qt.Recovering("disconnect QXmlReader::setContentHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setContentHandler")
	}
}

func (ptr *QXmlReader) SetContentHandler(handler QXmlContentHandler_ITF) {
	defer qt.Recovering("QXmlReader::setContentHandler")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetContentHandler(ptr.Pointer(), PointerFromQXmlContentHandler(handler))
	}
}

//export callbackQXmlReader_SetDTDHandler
func callbackQXmlReader_SetDTDHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlReader::setDTDHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDTDHandler"); signal != nil {
		signal.(func(*QXmlDTDHandler))(NewQXmlDTDHandlerFromPointer(handler))
	}

}

func (ptr *QXmlReader) ConnectSetDTDHandler(f func(handler *QXmlDTDHandler)) {
	defer qt.Recovering("connect QXmlReader::setDTDHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setDTDHandler", f)
	}
}

func (ptr *QXmlReader) DisconnectSetDTDHandler(handler QXmlDTDHandler_ITF) {
	defer qt.Recovering("disconnect QXmlReader::setDTDHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setDTDHandler")
	}
}

func (ptr *QXmlReader) SetDTDHandler(handler QXmlDTDHandler_ITF) {
	defer qt.Recovering("QXmlReader::setDTDHandler")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetDTDHandler(ptr.Pointer(), PointerFromQXmlDTDHandler(handler))
	}
}

//export callbackQXmlReader_SetDeclHandler
func callbackQXmlReader_SetDeclHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlReader::setDeclHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDeclHandler"); signal != nil {
		signal.(func(*QXmlDeclHandler))(NewQXmlDeclHandlerFromPointer(handler))
	}

}

func (ptr *QXmlReader) ConnectSetDeclHandler(f func(handler *QXmlDeclHandler)) {
	defer qt.Recovering("connect QXmlReader::setDeclHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setDeclHandler", f)
	}
}

func (ptr *QXmlReader) DisconnectSetDeclHandler(handler QXmlDeclHandler_ITF) {
	defer qt.Recovering("disconnect QXmlReader::setDeclHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setDeclHandler")
	}
}

func (ptr *QXmlReader) SetDeclHandler(handler QXmlDeclHandler_ITF) {
	defer qt.Recovering("QXmlReader::setDeclHandler")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetDeclHandler(ptr.Pointer(), PointerFromQXmlDeclHandler(handler))
	}
}

//export callbackQXmlReader_SetEntityResolver
func callbackQXmlReader_SetEntityResolver(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlReader::setEntityResolver")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEntityResolver"); signal != nil {
		signal.(func(*QXmlEntityResolver))(NewQXmlEntityResolverFromPointer(handler))
	}

}

func (ptr *QXmlReader) ConnectSetEntityResolver(f func(handler *QXmlEntityResolver)) {
	defer qt.Recovering("connect QXmlReader::setEntityResolver")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setEntityResolver", f)
	}
}

func (ptr *QXmlReader) DisconnectSetEntityResolver(handler QXmlEntityResolver_ITF) {
	defer qt.Recovering("disconnect QXmlReader::setEntityResolver")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setEntityResolver")
	}
}

func (ptr *QXmlReader) SetEntityResolver(handler QXmlEntityResolver_ITF) {
	defer qt.Recovering("QXmlReader::setEntityResolver")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetEntityResolver(ptr.Pointer(), PointerFromQXmlEntityResolver(handler))
	}
}

//export callbackQXmlReader_SetErrorHandler
func callbackQXmlReader_SetErrorHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlReader::setErrorHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setErrorHandler"); signal != nil {
		signal.(func(*QXmlErrorHandler))(NewQXmlErrorHandlerFromPointer(handler))
	}

}

func (ptr *QXmlReader) ConnectSetErrorHandler(f func(handler *QXmlErrorHandler)) {
	defer qt.Recovering("connect QXmlReader::setErrorHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setErrorHandler", f)
	}
}

func (ptr *QXmlReader) DisconnectSetErrorHandler(handler QXmlErrorHandler_ITF) {
	defer qt.Recovering("disconnect QXmlReader::setErrorHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setErrorHandler")
	}
}

func (ptr *QXmlReader) SetErrorHandler(handler QXmlErrorHandler_ITF) {
	defer qt.Recovering("QXmlReader::setErrorHandler")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetErrorHandler(ptr.Pointer(), PointerFromQXmlErrorHandler(handler))
	}
}

//export callbackQXmlReader_SetFeature
func callbackQXmlReader_SetFeature(ptr unsafe.Pointer, ptrName *C.char, name *C.char, value C.int) {
	defer qt.Recovering("callback QXmlReader::setFeature")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFeature"); signal != nil {
		signal.(func(string, bool))(C.GoString(name), int(value) != 0)
	}

}

func (ptr *QXmlReader) ConnectSetFeature(f func(name string, value bool)) {
	defer qt.Recovering("connect QXmlReader::setFeature")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setFeature", f)
	}
}

func (ptr *QXmlReader) DisconnectSetFeature(name string, value bool) {
	defer qt.Recovering("disconnect QXmlReader::setFeature")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setFeature")
	}
}

func (ptr *QXmlReader) SetFeature(name string, value bool) {
	defer qt.Recovering("QXmlReader::setFeature")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetFeature(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(value)))
	}
}

//export callbackQXmlReader_SetLexicalHandler
func callbackQXmlReader_SetLexicalHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlReader::setLexicalHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setLexicalHandler"); signal != nil {
		signal.(func(*QXmlLexicalHandler))(NewQXmlLexicalHandlerFromPointer(handler))
	}

}

func (ptr *QXmlReader) ConnectSetLexicalHandler(f func(handler *QXmlLexicalHandler)) {
	defer qt.Recovering("connect QXmlReader::setLexicalHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setLexicalHandler", f)
	}
}

func (ptr *QXmlReader) DisconnectSetLexicalHandler(handler QXmlLexicalHandler_ITF) {
	defer qt.Recovering("disconnect QXmlReader::setLexicalHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setLexicalHandler")
	}
}

func (ptr *QXmlReader) SetLexicalHandler(handler QXmlLexicalHandler_ITF) {
	defer qt.Recovering("QXmlReader::setLexicalHandler")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetLexicalHandler(ptr.Pointer(), PointerFromQXmlLexicalHandler(handler))
	}
}

//export callbackQXmlReader_SetProperty
func callbackQXmlReader_SetProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QXmlReader::setProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "setProperty"); signal != nil {
		signal.(func(string, unsafe.Pointer))(C.GoString(name), value)
	}

}

func (ptr *QXmlReader) ConnectSetProperty(f func(name string, value unsafe.Pointer)) {
	defer qt.Recovering("connect QXmlReader::setProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setProperty", f)
	}
}

func (ptr *QXmlReader) DisconnectSetProperty(name string, value unsafe.Pointer) {
	defer qt.Recovering("disconnect QXmlReader::setProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setProperty")
	}
}

func (ptr *QXmlReader) SetProperty(name string, value unsafe.Pointer) {
	defer qt.Recovering("QXmlReader::setProperty")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetProperty(ptr.Pointer(), C.CString(name), value)
	}
}

func (ptr *QXmlReader) DestroyQXmlReader() {
	defer qt.Recovering("QXmlReader::~QXmlReader")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QXmlReader_DestroyQXmlReader(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlReader) ObjectNameAbs() string {
	defer qt.Recovering("QXmlReader::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlReader_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlReader) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlReader::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlSimpleReader struct {
	QXmlReader
}

type QXmlSimpleReader_ITF interface {
	QXmlReader_ITF
	QXmlSimpleReader_PTR() *QXmlSimpleReader
}

func (p *QXmlSimpleReader) QXmlSimpleReader_PTR() *QXmlSimpleReader {
	return p
}

func (p *QXmlSimpleReader) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QXmlReader_PTR().Pointer()
	}
	return nil
}

func (p *QXmlSimpleReader) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QXmlReader_PTR().SetPointer(ptr)
	}
}

func PointerFromQXmlSimpleReader(ptr QXmlSimpleReader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSimpleReader_PTR().Pointer()
	}
	return nil
}

func NewQXmlSimpleReaderFromPointer(ptr unsafe.Pointer) *QXmlSimpleReader {
	var n = new(QXmlSimpleReader)
	n.SetPointer(ptr)
	return n
}

func newQXmlSimpleReaderFromPointer(ptr unsafe.Pointer) *QXmlSimpleReader {
	var n = NewQXmlSimpleReaderFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlSimpleReader_") {
		n.SetObjectNameAbs("QXmlSimpleReader_" + qt.Identifier())
	}
	return n
}

//export callbackQXmlSimpleReader_DTDHandler
func callbackQXmlSimpleReader_DTDHandler(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QXmlSimpleReader::DTDHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "DTDHandler"); signal != nil {
		return PointerFromQXmlDTDHandler(signal.(func() *QXmlDTDHandler)())
	}

	return PointerFromQXmlDTDHandler(NewQXmlSimpleReaderFromPointer(ptr).DTDHandlerDefault())
}

func (ptr *QXmlSimpleReader) ConnectDTDHandler(f func() *QXmlDTDHandler) {
	defer qt.Recovering("connect QXmlSimpleReader::DTDHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "DTDHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectDTDHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::DTDHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "DTDHandler")
	}
}

func (ptr *QXmlSimpleReader) DTDHandler() *QXmlDTDHandler {
	defer qt.Recovering("QXmlSimpleReader::DTDHandler")

	if ptr.Pointer() != nil {
		return NewQXmlDTDHandlerFromPointer(C.QXmlSimpleReader_DTDHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) DTDHandlerDefault() *QXmlDTDHandler {
	defer qt.Recovering("QXmlSimpleReader::DTDHandler")

	if ptr.Pointer() != nil {
		return NewQXmlDTDHandlerFromPointer(C.QXmlSimpleReader_DTDHandlerDefault(ptr.Pointer()))
	}
	return nil
}

func NewQXmlSimpleReader() *QXmlSimpleReader {
	defer qt.Recovering("QXmlSimpleReader::QXmlSimpleReader")

	return newQXmlSimpleReaderFromPointer(C.QXmlSimpleReader_NewQXmlSimpleReader())
}

//export callbackQXmlSimpleReader_ContentHandler
func callbackQXmlSimpleReader_ContentHandler(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QXmlSimpleReader::contentHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentHandler"); signal != nil {
		return PointerFromQXmlContentHandler(signal.(func() *QXmlContentHandler)())
	}

	return PointerFromQXmlContentHandler(NewQXmlSimpleReaderFromPointer(ptr).ContentHandlerDefault())
}

func (ptr *QXmlSimpleReader) ConnectContentHandler(f func() *QXmlContentHandler) {
	defer qt.Recovering("connect QXmlSimpleReader::contentHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "contentHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectContentHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::contentHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "contentHandler")
	}
}

func (ptr *QXmlSimpleReader) ContentHandler() *QXmlContentHandler {
	defer qt.Recovering("QXmlSimpleReader::contentHandler")

	if ptr.Pointer() != nil {
		return NewQXmlContentHandlerFromPointer(C.QXmlSimpleReader_ContentHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) ContentHandlerDefault() *QXmlContentHandler {
	defer qt.Recovering("QXmlSimpleReader::contentHandler")

	if ptr.Pointer() != nil {
		return NewQXmlContentHandlerFromPointer(C.QXmlSimpleReader_ContentHandlerDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlSimpleReader_DeclHandler
func callbackQXmlSimpleReader_DeclHandler(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QXmlSimpleReader::declHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "declHandler"); signal != nil {
		return PointerFromQXmlDeclHandler(signal.(func() *QXmlDeclHandler)())
	}

	return PointerFromQXmlDeclHandler(NewQXmlSimpleReaderFromPointer(ptr).DeclHandlerDefault())
}

func (ptr *QXmlSimpleReader) ConnectDeclHandler(f func() *QXmlDeclHandler) {
	defer qt.Recovering("connect QXmlSimpleReader::declHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "declHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectDeclHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::declHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "declHandler")
	}
}

func (ptr *QXmlSimpleReader) DeclHandler() *QXmlDeclHandler {
	defer qt.Recovering("QXmlSimpleReader::declHandler")

	if ptr.Pointer() != nil {
		return NewQXmlDeclHandlerFromPointer(C.QXmlSimpleReader_DeclHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) DeclHandlerDefault() *QXmlDeclHandler {
	defer qt.Recovering("QXmlSimpleReader::declHandler")

	if ptr.Pointer() != nil {
		return NewQXmlDeclHandlerFromPointer(C.QXmlSimpleReader_DeclHandlerDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlSimpleReader_EntityResolver
func callbackQXmlSimpleReader_EntityResolver(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QXmlSimpleReader::entityResolver")

	if signal := qt.GetSignal(C.GoString(ptrName), "entityResolver"); signal != nil {
		return PointerFromQXmlEntityResolver(signal.(func() *QXmlEntityResolver)())
	}

	return PointerFromQXmlEntityResolver(NewQXmlSimpleReaderFromPointer(ptr).EntityResolverDefault())
}

func (ptr *QXmlSimpleReader) ConnectEntityResolver(f func() *QXmlEntityResolver) {
	defer qt.Recovering("connect QXmlSimpleReader::entityResolver")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "entityResolver", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectEntityResolver() {
	defer qt.Recovering("disconnect QXmlSimpleReader::entityResolver")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "entityResolver")
	}
}

func (ptr *QXmlSimpleReader) EntityResolver() *QXmlEntityResolver {
	defer qt.Recovering("QXmlSimpleReader::entityResolver")

	if ptr.Pointer() != nil {
		return NewQXmlEntityResolverFromPointer(C.QXmlSimpleReader_EntityResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) EntityResolverDefault() *QXmlEntityResolver {
	defer qt.Recovering("QXmlSimpleReader::entityResolver")

	if ptr.Pointer() != nil {
		return NewQXmlEntityResolverFromPointer(C.QXmlSimpleReader_EntityResolverDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlSimpleReader_ErrorHandler
func callbackQXmlSimpleReader_ErrorHandler(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QXmlSimpleReader::errorHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorHandler"); signal != nil {
		return PointerFromQXmlErrorHandler(signal.(func() *QXmlErrorHandler)())
	}

	return PointerFromQXmlErrorHandler(NewQXmlSimpleReaderFromPointer(ptr).ErrorHandlerDefault())
}

func (ptr *QXmlSimpleReader) ConnectErrorHandler(f func() *QXmlErrorHandler) {
	defer qt.Recovering("connect QXmlSimpleReader::errorHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "errorHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectErrorHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::errorHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "errorHandler")
	}
}

func (ptr *QXmlSimpleReader) ErrorHandler() *QXmlErrorHandler {
	defer qt.Recovering("QXmlSimpleReader::errorHandler")

	if ptr.Pointer() != nil {
		return NewQXmlErrorHandlerFromPointer(C.QXmlSimpleReader_ErrorHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) ErrorHandlerDefault() *QXmlErrorHandler {
	defer qt.Recovering("QXmlSimpleReader::errorHandler")

	if ptr.Pointer() != nil {
		return NewQXmlErrorHandlerFromPointer(C.QXmlSimpleReader_ErrorHandlerDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlSimpleReader_Feature
func callbackQXmlSimpleReader_Feature(ptr unsafe.Pointer, ptrName *C.char, name *C.char, ok C.int) C.int {
	defer qt.Recovering("callback QXmlSimpleReader::feature")

	if signal := qt.GetSignal(C.GoString(ptrName), "feature"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, bool) bool)(C.GoString(name), int(ok) != 0)))
	}

	return C.int(qt.GoBoolToInt(NewQXmlSimpleReaderFromPointer(ptr).FeatureDefault(C.GoString(name), int(ok) != 0)))
}

func (ptr *QXmlSimpleReader) ConnectFeature(f func(name string, ok bool) bool) {
	defer qt.Recovering("connect QXmlSimpleReader::feature")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "feature", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectFeature() {
	defer qt.Recovering("disconnect QXmlSimpleReader::feature")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "feature")
	}
}

func (ptr *QXmlSimpleReader) Feature(name string, ok bool) bool {
	defer qt.Recovering("QXmlSimpleReader::feature")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Feature(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) FeatureDefault(name string, ok bool) bool {
	defer qt.Recovering("QXmlSimpleReader::feature")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_FeatureDefault(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))) != 0
	}
	return false
}

//export callbackQXmlSimpleReader_HasFeature
func callbackQXmlSimpleReader_HasFeature(ptr unsafe.Pointer, ptrName *C.char, name *C.char) C.int {
	defer qt.Recovering("callback QXmlSimpleReader::hasFeature")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasFeature"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlSimpleReaderFromPointer(ptr).HasFeatureDefault(C.GoString(name))))
}

func (ptr *QXmlSimpleReader) ConnectHasFeature(f func(name string) bool) {
	defer qt.Recovering("connect QXmlSimpleReader::hasFeature")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "hasFeature", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectHasFeature() {
	defer qt.Recovering("disconnect QXmlSimpleReader::hasFeature")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "hasFeature")
	}
}

func (ptr *QXmlSimpleReader) HasFeature(name string) bool {
	defer qt.Recovering("QXmlSimpleReader::hasFeature")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_HasFeature(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) HasFeatureDefault(name string) bool {
	defer qt.Recovering("QXmlSimpleReader::hasFeature")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_HasFeatureDefault(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

//export callbackQXmlSimpleReader_HasProperty
func callbackQXmlSimpleReader_HasProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char) C.int {
	defer qt.Recovering("callback QXmlSimpleReader::hasProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasProperty"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlSimpleReaderFromPointer(ptr).HasPropertyDefault(C.GoString(name))))
}

func (ptr *QXmlSimpleReader) ConnectHasProperty(f func(name string) bool) {
	defer qt.Recovering("connect QXmlSimpleReader::hasProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "hasProperty", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectHasProperty() {
	defer qt.Recovering("disconnect QXmlSimpleReader::hasProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "hasProperty")
	}
}

func (ptr *QXmlSimpleReader) HasProperty(name string) bool {
	defer qt.Recovering("QXmlSimpleReader::hasProperty")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_HasProperty(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) HasPropertyDefault(name string) bool {
	defer qt.Recovering("QXmlSimpleReader::hasProperty")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_HasPropertyDefault(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

//export callbackQXmlSimpleReader_LexicalHandler
func callbackQXmlSimpleReader_LexicalHandler(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QXmlSimpleReader::lexicalHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "lexicalHandler"); signal != nil {
		return PointerFromQXmlLexicalHandler(signal.(func() *QXmlLexicalHandler)())
	}

	return PointerFromQXmlLexicalHandler(NewQXmlSimpleReaderFromPointer(ptr).LexicalHandlerDefault())
}

func (ptr *QXmlSimpleReader) ConnectLexicalHandler(f func() *QXmlLexicalHandler) {
	defer qt.Recovering("connect QXmlSimpleReader::lexicalHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "lexicalHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectLexicalHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::lexicalHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "lexicalHandler")
	}
}

func (ptr *QXmlSimpleReader) LexicalHandler() *QXmlLexicalHandler {
	defer qt.Recovering("QXmlSimpleReader::lexicalHandler")

	if ptr.Pointer() != nil {
		return NewQXmlLexicalHandlerFromPointer(C.QXmlSimpleReader_LexicalHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) LexicalHandlerDefault() *QXmlLexicalHandler {
	defer qt.Recovering("QXmlSimpleReader::lexicalHandler")

	if ptr.Pointer() != nil {
		return NewQXmlLexicalHandlerFromPointer(C.QXmlSimpleReader_LexicalHandlerDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlSimpleReader_Parse2
func callbackQXmlSimpleReader_Parse2(ptr unsafe.Pointer, ptrName *C.char, input unsafe.Pointer) C.int {
	defer qt.Recovering("callback QXmlSimpleReader::parse")

	if signal := qt.GetSignal(C.GoString(ptrName), "parse2"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QXmlInputSource) bool)(NewQXmlInputSourceFromPointer(input))))
	}

	return C.int(qt.GoBoolToInt(NewQXmlSimpleReaderFromPointer(ptr).Parse2Default(NewQXmlInputSourceFromPointer(input))))
}

func (ptr *QXmlSimpleReader) ConnectParse2(f func(input *QXmlInputSource) bool) {
	defer qt.Recovering("connect QXmlSimpleReader::parse")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "parse2", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectParse2() {
	defer qt.Recovering("disconnect QXmlSimpleReader::parse")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "parse2")
	}
}

func (ptr *QXmlSimpleReader) Parse2(input QXmlInputSource_ITF) bool {
	defer qt.Recovering("QXmlSimpleReader::parse")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Parse2(ptr.Pointer(), PointerFromQXmlInputSource(input)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) Parse2Default(input QXmlInputSource_ITF) bool {
	defer qt.Recovering("QXmlSimpleReader::parse")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Parse2Default(ptr.Pointer(), PointerFromQXmlInputSource(input)) != 0
	}
	return false
}

//export callbackQXmlSimpleReader_Parse3
func callbackQXmlSimpleReader_Parse3(ptr unsafe.Pointer, ptrName *C.char, input unsafe.Pointer, incremental C.int) C.int {
	defer qt.Recovering("callback QXmlSimpleReader::parse")

	if signal := qt.GetSignal(C.GoString(ptrName), "parse3"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QXmlInputSource, bool) bool)(NewQXmlInputSourceFromPointer(input), int(incremental) != 0)))
	}

	return C.int(qt.GoBoolToInt(NewQXmlSimpleReaderFromPointer(ptr).Parse3Default(NewQXmlInputSourceFromPointer(input), int(incremental) != 0)))
}

func (ptr *QXmlSimpleReader) ConnectParse3(f func(input *QXmlInputSource, incremental bool) bool) {
	defer qt.Recovering("connect QXmlSimpleReader::parse")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "parse3", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectParse3() {
	defer qt.Recovering("disconnect QXmlSimpleReader::parse")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "parse3")
	}
}

func (ptr *QXmlSimpleReader) Parse3(input QXmlInputSource_ITF, incremental bool) bool {
	defer qt.Recovering("QXmlSimpleReader::parse")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Parse3(ptr.Pointer(), PointerFromQXmlInputSource(input), C.int(qt.GoBoolToInt(incremental))) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) Parse3Default(input QXmlInputSource_ITF, incremental bool) bool {
	defer qt.Recovering("QXmlSimpleReader::parse")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Parse3Default(ptr.Pointer(), PointerFromQXmlInputSource(input), C.int(qt.GoBoolToInt(incremental))) != 0
	}
	return false
}

//export callbackQXmlSimpleReader_ParseContinue
func callbackQXmlSimpleReader_ParseContinue(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QXmlSimpleReader::parseContinue")

	if signal := qt.GetSignal(C.GoString(ptrName), "parseContinue"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQXmlSimpleReaderFromPointer(ptr).ParseContinueDefault()))
}

func (ptr *QXmlSimpleReader) ConnectParseContinue(f func() bool) {
	defer qt.Recovering("connect QXmlSimpleReader::parseContinue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "parseContinue", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectParseContinue() {
	defer qt.Recovering("disconnect QXmlSimpleReader::parseContinue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "parseContinue")
	}
}

func (ptr *QXmlSimpleReader) ParseContinue() bool {
	defer qt.Recovering("QXmlSimpleReader::parseContinue")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_ParseContinue(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) ParseContinueDefault() bool {
	defer qt.Recovering("QXmlSimpleReader::parseContinue")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_ParseContinueDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQXmlSimpleReader_Property
func callbackQXmlSimpleReader_Property(ptr unsafe.Pointer, ptrName *C.char, name *C.char, ok C.int) unsafe.Pointer {
	defer qt.Recovering("callback QXmlSimpleReader::property")

	if signal := qt.GetSignal(C.GoString(ptrName), "property"); signal != nil {
		return signal.(func(string, bool) unsafe.Pointer)(C.GoString(name), int(ok) != 0)
	}

	return NewQXmlSimpleReaderFromPointer(ptr).PropertyDefault(C.GoString(name), int(ok) != 0)
}

func (ptr *QXmlSimpleReader) ConnectProperty(f func(name string, ok bool) unsafe.Pointer) {
	defer qt.Recovering("connect QXmlSimpleReader::property")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "property", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectProperty() {
	defer qt.Recovering("disconnect QXmlSimpleReader::property")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "property")
	}
}

func (ptr *QXmlSimpleReader) Property(name string, ok bool) unsafe.Pointer {
	defer qt.Recovering("QXmlSimpleReader::property")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXmlSimpleReader_Property(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))))
	}
	return nil
}

func (ptr *QXmlSimpleReader) PropertyDefault(name string, ok bool) unsafe.Pointer {
	defer qt.Recovering("QXmlSimpleReader::property")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXmlSimpleReader_PropertyDefault(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))))
	}
	return nil
}

//export callbackQXmlSimpleReader_SetContentHandler
func callbackQXmlSimpleReader_SetContentHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setContentHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setContentHandler"); signal != nil {
		signal.(func(*QXmlContentHandler))(NewQXmlContentHandlerFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetContentHandlerDefault(NewQXmlContentHandlerFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetContentHandler(f func(handler *QXmlContentHandler)) {
	defer qt.Recovering("connect QXmlSimpleReader::setContentHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setContentHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetContentHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setContentHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setContentHandler")
	}
}

func (ptr *QXmlSimpleReader) SetContentHandler(handler QXmlContentHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setContentHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetContentHandler(ptr.Pointer(), PointerFromQXmlContentHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetContentHandlerDefault(handler QXmlContentHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setContentHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetContentHandlerDefault(ptr.Pointer(), PointerFromQXmlContentHandler(handler))
	}
}

//export callbackQXmlSimpleReader_SetDTDHandler
func callbackQXmlSimpleReader_SetDTDHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setDTDHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDTDHandler"); signal != nil {
		signal.(func(*QXmlDTDHandler))(NewQXmlDTDHandlerFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetDTDHandlerDefault(NewQXmlDTDHandlerFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetDTDHandler(f func(handler *QXmlDTDHandler)) {
	defer qt.Recovering("connect QXmlSimpleReader::setDTDHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setDTDHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetDTDHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setDTDHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setDTDHandler")
	}
}

func (ptr *QXmlSimpleReader) SetDTDHandler(handler QXmlDTDHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setDTDHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDTDHandler(ptr.Pointer(), PointerFromQXmlDTDHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetDTDHandlerDefault(handler QXmlDTDHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setDTDHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDTDHandlerDefault(ptr.Pointer(), PointerFromQXmlDTDHandler(handler))
	}
}

//export callbackQXmlSimpleReader_SetDeclHandler
func callbackQXmlSimpleReader_SetDeclHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setDeclHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDeclHandler"); signal != nil {
		signal.(func(*QXmlDeclHandler))(NewQXmlDeclHandlerFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetDeclHandlerDefault(NewQXmlDeclHandlerFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetDeclHandler(f func(handler *QXmlDeclHandler)) {
	defer qt.Recovering("connect QXmlSimpleReader::setDeclHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setDeclHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetDeclHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setDeclHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setDeclHandler")
	}
}

func (ptr *QXmlSimpleReader) SetDeclHandler(handler QXmlDeclHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setDeclHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDeclHandler(ptr.Pointer(), PointerFromQXmlDeclHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetDeclHandlerDefault(handler QXmlDeclHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setDeclHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDeclHandlerDefault(ptr.Pointer(), PointerFromQXmlDeclHandler(handler))
	}
}

//export callbackQXmlSimpleReader_SetEntityResolver
func callbackQXmlSimpleReader_SetEntityResolver(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setEntityResolver")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEntityResolver"); signal != nil {
		signal.(func(*QXmlEntityResolver))(NewQXmlEntityResolverFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetEntityResolverDefault(NewQXmlEntityResolverFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetEntityResolver(f func(handler *QXmlEntityResolver)) {
	defer qt.Recovering("connect QXmlSimpleReader::setEntityResolver")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setEntityResolver", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetEntityResolver() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setEntityResolver")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setEntityResolver")
	}
}

func (ptr *QXmlSimpleReader) SetEntityResolver(handler QXmlEntityResolver_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setEntityResolver")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetEntityResolver(ptr.Pointer(), PointerFromQXmlEntityResolver(handler))
	}
}

func (ptr *QXmlSimpleReader) SetEntityResolverDefault(handler QXmlEntityResolver_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setEntityResolver")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetEntityResolverDefault(ptr.Pointer(), PointerFromQXmlEntityResolver(handler))
	}
}

//export callbackQXmlSimpleReader_SetErrorHandler
func callbackQXmlSimpleReader_SetErrorHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setErrorHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setErrorHandler"); signal != nil {
		signal.(func(*QXmlErrorHandler))(NewQXmlErrorHandlerFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetErrorHandlerDefault(NewQXmlErrorHandlerFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetErrorHandler(f func(handler *QXmlErrorHandler)) {
	defer qt.Recovering("connect QXmlSimpleReader::setErrorHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setErrorHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetErrorHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setErrorHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setErrorHandler")
	}
}

func (ptr *QXmlSimpleReader) SetErrorHandler(handler QXmlErrorHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setErrorHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetErrorHandler(ptr.Pointer(), PointerFromQXmlErrorHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetErrorHandlerDefault(handler QXmlErrorHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setErrorHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetErrorHandlerDefault(ptr.Pointer(), PointerFromQXmlErrorHandler(handler))
	}
}

//export callbackQXmlSimpleReader_SetFeature
func callbackQXmlSimpleReader_SetFeature(ptr unsafe.Pointer, ptrName *C.char, name *C.char, enable C.int) {
	defer qt.Recovering("callback QXmlSimpleReader::setFeature")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFeature"); signal != nil {
		signal.(func(string, bool))(C.GoString(name), int(enable) != 0)
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetFeatureDefault(C.GoString(name), int(enable) != 0)
	}
}

func (ptr *QXmlSimpleReader) ConnectSetFeature(f func(name string, enable bool)) {
	defer qt.Recovering("connect QXmlSimpleReader::setFeature")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setFeature", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetFeature() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setFeature")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setFeature")
	}
}

func (ptr *QXmlSimpleReader) SetFeature(name string, enable bool) {
	defer qt.Recovering("QXmlSimpleReader::setFeature")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetFeature(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QXmlSimpleReader) SetFeatureDefault(name string, enable bool) {
	defer qt.Recovering("QXmlSimpleReader::setFeature")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetFeatureDefault(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(enable)))
	}
}

//export callbackQXmlSimpleReader_SetLexicalHandler
func callbackQXmlSimpleReader_SetLexicalHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setLexicalHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setLexicalHandler"); signal != nil {
		signal.(func(*QXmlLexicalHandler))(NewQXmlLexicalHandlerFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetLexicalHandlerDefault(NewQXmlLexicalHandlerFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetLexicalHandler(f func(handler *QXmlLexicalHandler)) {
	defer qt.Recovering("connect QXmlSimpleReader::setLexicalHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setLexicalHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetLexicalHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setLexicalHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setLexicalHandler")
	}
}

func (ptr *QXmlSimpleReader) SetLexicalHandler(handler QXmlLexicalHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setLexicalHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetLexicalHandler(ptr.Pointer(), PointerFromQXmlLexicalHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetLexicalHandlerDefault(handler QXmlLexicalHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setLexicalHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetLexicalHandlerDefault(ptr.Pointer(), PointerFromQXmlLexicalHandler(handler))
	}
}

//export callbackQXmlSimpleReader_SetProperty
func callbackQXmlSimpleReader_SetProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "setProperty"); signal != nil {
		signal.(func(string, unsafe.Pointer))(C.GoString(name), value)
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetPropertyDefault(C.GoString(name), value)
	}
}

func (ptr *QXmlSimpleReader) ConnectSetProperty(f func(name string, value unsafe.Pointer)) {
	defer qt.Recovering("connect QXmlSimpleReader::setProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setProperty", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetProperty() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setProperty")
	}
}

func (ptr *QXmlSimpleReader) SetProperty(name string, value unsafe.Pointer) {
	defer qt.Recovering("QXmlSimpleReader::setProperty")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetProperty(ptr.Pointer(), C.CString(name), value)
	}
}

func (ptr *QXmlSimpleReader) SetPropertyDefault(name string, value unsafe.Pointer) {
	defer qt.Recovering("QXmlSimpleReader::setProperty")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetPropertyDefault(ptr.Pointer(), C.CString(name), value)
	}
}

func (ptr *QXmlSimpleReader) DestroyQXmlSimpleReader() {
	defer qt.Recovering("QXmlSimpleReader::~QXmlSimpleReader")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QXmlSimpleReader_DestroyQXmlSimpleReader(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlSimpleReader) ObjectNameAbs() string {
	defer qt.Recovering("QXmlSimpleReader::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlSimpleReader_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlSimpleReader) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlSimpleReader::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
