// +build !minimal

package serialport

//#include "serialport.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

//QSerialPort::BaudRate
type QSerialPort__BaudRate int64

const (
	QSerialPort__Baud1200    = QSerialPort__BaudRate(1200)
	QSerialPort__Baud2400    = QSerialPort__BaudRate(2400)
	QSerialPort__Baud4800    = QSerialPort__BaudRate(4800)
	QSerialPort__Baud9600    = QSerialPort__BaudRate(9600)
	QSerialPort__Baud19200   = QSerialPort__BaudRate(19200)
	QSerialPort__Baud38400   = QSerialPort__BaudRate(38400)
	QSerialPort__Baud57600   = QSerialPort__BaudRate(57600)
	QSerialPort__Baud115200  = QSerialPort__BaudRate(115200)
	QSerialPort__UnknownBaud = QSerialPort__BaudRate(-1)
)

//QSerialPort::DataBits
type QSerialPort__DataBits int64

const (
	QSerialPort__Data5           = QSerialPort__DataBits(5)
	QSerialPort__Data6           = QSerialPort__DataBits(6)
	QSerialPort__Data7           = QSerialPort__DataBits(7)
	QSerialPort__Data8           = QSerialPort__DataBits(8)
	QSerialPort__UnknownDataBits = QSerialPort__DataBits(-1)
)

//QSerialPort::Direction
type QSerialPort__Direction int64

const (
	QSerialPort__Input         = QSerialPort__Direction(1)
	QSerialPort__Output        = QSerialPort__Direction(2)
	QSerialPort__AllDirections = QSerialPort__Direction(QSerialPort__Input | QSerialPort__Output)
)

//QSerialPort::FlowControl
type QSerialPort__FlowControl int64

const (
	QSerialPort__NoFlowControl      = QSerialPort__FlowControl(0)
	QSerialPort__HardwareControl    = QSerialPort__FlowControl(1)
	QSerialPort__SoftwareControl    = QSerialPort__FlowControl(2)
	QSerialPort__UnknownFlowControl = QSerialPort__FlowControl(-1)
)

//QSerialPort::Parity
type QSerialPort__Parity int64

const (
	QSerialPort__NoParity      = QSerialPort__Parity(0)
	QSerialPort__EvenParity    = QSerialPort__Parity(2)
	QSerialPort__OddParity     = QSerialPort__Parity(3)
	QSerialPort__SpaceParity   = QSerialPort__Parity(4)
	QSerialPort__MarkParity    = QSerialPort__Parity(5)
	QSerialPort__UnknownParity = QSerialPort__Parity(-1)
)

//QSerialPort::PinoutSignal
type QSerialPort__PinoutSignal int64

const (
	QSerialPort__NoSignal                       = QSerialPort__PinoutSignal(0x00)
	QSerialPort__TransmittedDataSignal          = QSerialPort__PinoutSignal(0x01)
	QSerialPort__ReceivedDataSignal             = QSerialPort__PinoutSignal(0x02)
	QSerialPort__DataTerminalReadySignal        = QSerialPort__PinoutSignal(0x04)
	QSerialPort__DataCarrierDetectSignal        = QSerialPort__PinoutSignal(0x08)
	QSerialPort__DataSetReadySignal             = QSerialPort__PinoutSignal(0x10)
	QSerialPort__RingIndicatorSignal            = QSerialPort__PinoutSignal(0x20)
	QSerialPort__RequestToSendSignal            = QSerialPort__PinoutSignal(0x40)
	QSerialPort__ClearToSendSignal              = QSerialPort__PinoutSignal(0x80)
	QSerialPort__SecondaryTransmittedDataSignal = QSerialPort__PinoutSignal(0x100)
	QSerialPort__SecondaryReceivedDataSignal    = QSerialPort__PinoutSignal(0x200)
)

//QSerialPort::SerialPortError
type QSerialPort__SerialPortError int64

const (
	QSerialPort__NoError                   = QSerialPort__SerialPortError(0)
	QSerialPort__DeviceNotFoundError       = QSerialPort__SerialPortError(1)
	QSerialPort__PermissionError           = QSerialPort__SerialPortError(2)
	QSerialPort__OpenError                 = QSerialPort__SerialPortError(3)
	QSerialPort__ParityError               = QSerialPort__SerialPortError(4)
	QSerialPort__FramingError              = QSerialPort__SerialPortError(5)
	QSerialPort__BreakConditionError       = QSerialPort__SerialPortError(6)
	QSerialPort__WriteError                = QSerialPort__SerialPortError(7)
	QSerialPort__ReadError                 = QSerialPort__SerialPortError(8)
	QSerialPort__ResourceError             = QSerialPort__SerialPortError(9)
	QSerialPort__UnsupportedOperationError = QSerialPort__SerialPortError(10)
	QSerialPort__UnknownError              = QSerialPort__SerialPortError(11)
	QSerialPort__TimeoutError              = QSerialPort__SerialPortError(12)
	QSerialPort__NotOpenError              = QSerialPort__SerialPortError(13)
)

//QSerialPort::StopBits
type QSerialPort__StopBits int64

const (
	QSerialPort__OneStop         = QSerialPort__StopBits(1)
	QSerialPort__OneAndHalfStop  = QSerialPort__StopBits(3)
	QSerialPort__TwoStop         = QSerialPort__StopBits(2)
	QSerialPort__UnknownStopBits = QSerialPort__StopBits(-1)
)

type QSerialPort struct {
	core.QIODevice
}

type QSerialPort_ITF interface {
	core.QIODevice_ITF
	QSerialPort_PTR() *QSerialPort
}

func (p *QSerialPort) QSerialPort_PTR() *QSerialPort {
	return p
}

func (p *QSerialPort) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QIODevice_PTR().Pointer()
	}
	return nil
}

func (p *QSerialPort) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QIODevice_PTR().SetPointer(ptr)
	}
}

func PointerFromQSerialPort(ptr QSerialPort_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSerialPort_PTR().Pointer()
	}
	return nil
}

func NewQSerialPortFromPointer(ptr unsafe.Pointer) *QSerialPort {
	var n = new(QSerialPort)
	n.SetPointer(ptr)
	return n
}

func newQSerialPortFromPointer(ptr unsafe.Pointer) *QSerialPort {
	var n = NewQSerialPortFromPointer(ptr)
	for len(n.ObjectName()) < len("QSerialPort_") {
		n.SetObjectName("QSerialPort_" + qt.Identifier())
	}
	return n
}

func (ptr *QSerialPort) ClearError() {
	defer qt.Recovering("QSerialPort::clearError")

	if ptr.Pointer() != nil {
		C.QSerialPort_ClearError(ptr.Pointer())
	}
}

func (ptr *QSerialPort) DataBits() QSerialPort__DataBits {
	defer qt.Recovering("QSerialPort::dataBits")

	if ptr.Pointer() != nil {
		return QSerialPort__DataBits(C.QSerialPort_DataBits(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) Error() QSerialPort__SerialPortError {
	defer qt.Recovering("QSerialPort::error")

	if ptr.Pointer() != nil {
		return QSerialPort__SerialPortError(C.QSerialPort_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) FlowControl() QSerialPort__FlowControl {
	defer qt.Recovering("QSerialPort::flowControl")

	if ptr.Pointer() != nil {
		return QSerialPort__FlowControl(C.QSerialPort_FlowControl(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) IsBreakEnabled() bool {
	defer qt.Recovering("QSerialPort::isBreakEnabled")

	if ptr.Pointer() != nil {
		return C.QSerialPort_IsBreakEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPort) IsDataTerminalReady() bool {
	defer qt.Recovering("QSerialPort::isDataTerminalReady")

	if ptr.Pointer() != nil {
		return C.QSerialPort_IsDataTerminalReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPort) IsRequestToSend() bool {
	defer qt.Recovering("QSerialPort::isRequestToSend")

	if ptr.Pointer() != nil {
		return C.QSerialPort_IsRequestToSend(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPort) Parity() QSerialPort__Parity {
	defer qt.Recovering("QSerialPort::parity")

	if ptr.Pointer() != nil {
		return QSerialPort__Parity(C.QSerialPort_Parity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) SetBreakEnabled(set bool) bool {
	defer qt.Recovering("QSerialPort::setBreakEnabled")

	if ptr.Pointer() != nil {
		return C.QSerialPort_SetBreakEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(set))) != 0
	}
	return false
}

func (ptr *QSerialPort) SetDataBits(dataBits QSerialPort__DataBits) bool {
	defer qt.Recovering("QSerialPort::setDataBits")

	if ptr.Pointer() != nil {
		return C.QSerialPort_SetDataBits(ptr.Pointer(), C.int(dataBits)) != 0
	}
	return false
}

func (ptr *QSerialPort) SetDataTerminalReady(set bool) bool {
	defer qt.Recovering("QSerialPort::setDataTerminalReady")

	if ptr.Pointer() != nil {
		return C.QSerialPort_SetDataTerminalReady(ptr.Pointer(), C.int(qt.GoBoolToInt(set))) != 0
	}
	return false
}

func (ptr *QSerialPort) SetFlowControl(flowControl QSerialPort__FlowControl) bool {
	defer qt.Recovering("QSerialPort::setFlowControl")

	if ptr.Pointer() != nil {
		return C.QSerialPort_SetFlowControl(ptr.Pointer(), C.int(flowControl)) != 0
	}
	return false
}

func (ptr *QSerialPort) SetParity(parity QSerialPort__Parity) bool {
	defer qt.Recovering("QSerialPort::setParity")

	if ptr.Pointer() != nil {
		return C.QSerialPort_SetParity(ptr.Pointer(), C.int(parity)) != 0
	}
	return false
}

func (ptr *QSerialPort) SetRequestToSend(set bool) bool {
	defer qt.Recovering("QSerialPort::setRequestToSend")

	if ptr.Pointer() != nil {
		return C.QSerialPort_SetRequestToSend(ptr.Pointer(), C.int(qt.GoBoolToInt(set))) != 0
	}
	return false
}

func (ptr *QSerialPort) SetStopBits(stopBits QSerialPort__StopBits) bool {
	defer qt.Recovering("QSerialPort::setStopBits")

	if ptr.Pointer() != nil {
		return C.QSerialPort_SetStopBits(ptr.Pointer(), C.int(stopBits)) != 0
	}
	return false
}

func (ptr *QSerialPort) StopBits() QSerialPort__StopBits {
	defer qt.Recovering("QSerialPort::stopBits")

	if ptr.Pointer() != nil {
		return QSerialPort__StopBits(C.QSerialPort_StopBits(ptr.Pointer()))
	}
	return 0
}

func NewQSerialPort(parent core.QObject_ITF) *QSerialPort {
	defer qt.Recovering("QSerialPort::QSerialPort")

	return newQSerialPortFromPointer(C.QSerialPort_NewQSerialPort(core.PointerFromQObject(parent)))
}

func NewQSerialPort3(serialPortInfo QSerialPortInfo_ITF, parent core.QObject_ITF) *QSerialPort {
	defer qt.Recovering("QSerialPort::QSerialPort")

	return newQSerialPortFromPointer(C.QSerialPort_NewQSerialPort3(PointerFromQSerialPortInfo(serialPortInfo), core.PointerFromQObject(parent)))
}

func NewQSerialPort2(name string, parent core.QObject_ITF) *QSerialPort {
	defer qt.Recovering("QSerialPort::QSerialPort")

	return newQSerialPortFromPointer(C.QSerialPort_NewQSerialPort2(C.CString(name), core.PointerFromQObject(parent)))
}

func (ptr *QSerialPort) AtEnd() bool {
	defer qt.Recovering("QSerialPort::atEnd")

	if ptr.Pointer() != nil {
		return C.QSerialPort_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSerialPort_BreakEnabledChanged
func callbackQSerialPort_BreakEnabledChanged(ptr unsafe.Pointer, ptrName *C.char, set C.int) {
	defer qt.Recovering("callback QSerialPort::breakEnabledChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "breakEnabledChanged"); signal != nil {
		signal.(func(bool))(int(set) != 0)
	}

}

func (ptr *QSerialPort) ConnectBreakEnabledChanged(f func(set bool)) {
	defer qt.Recovering("connect QSerialPort::breakEnabledChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_ConnectBreakEnabledChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "breakEnabledChanged", f)
	}
}

func (ptr *QSerialPort) DisconnectBreakEnabledChanged() {
	defer qt.Recovering("disconnect QSerialPort::breakEnabledChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectBreakEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "breakEnabledChanged")
	}
}

func (ptr *QSerialPort) BreakEnabledChanged(set bool) {
	defer qt.Recovering("QSerialPort::breakEnabledChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_BreakEnabledChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(set)))
	}
}

func (ptr *QSerialPort) BytesAvailable() int64 {
	defer qt.Recovering("QSerialPort::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) BytesToWrite() int64 {
	defer qt.Recovering("QSerialPort::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) CanReadLine() bool {
	defer qt.Recovering("QSerialPort::canReadLine")

	if ptr.Pointer() != nil {
		return C.QSerialPort_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPort) Clear(directions QSerialPort__Direction) bool {
	defer qt.Recovering("QSerialPort::clear")

	if ptr.Pointer() != nil {
		return C.QSerialPort_Clear(ptr.Pointer(), C.int(directions)) != 0
	}
	return false
}

func (ptr *QSerialPort) Close() {
	defer qt.Recovering("QSerialPort::close")

	if ptr.Pointer() != nil {
		C.QSerialPort_Close(ptr.Pointer())
	}
}

//export callbackQSerialPort_DataBitsChanged
func callbackQSerialPort_DataBitsChanged(ptr unsafe.Pointer, ptrName *C.char, dataBits C.int) {
	defer qt.Recovering("callback QSerialPort::dataBitsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "dataBitsChanged"); signal != nil {
		signal.(func(QSerialPort__DataBits))(QSerialPort__DataBits(dataBits))
	}

}

func (ptr *QSerialPort) ConnectDataBitsChanged(f func(dataBits QSerialPort__DataBits)) {
	defer qt.Recovering("connect QSerialPort::dataBitsChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_ConnectDataBitsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "dataBitsChanged", f)
	}
}

func (ptr *QSerialPort) DisconnectDataBitsChanged() {
	defer qt.Recovering("disconnect QSerialPort::dataBitsChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectDataBitsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "dataBitsChanged")
	}
}

func (ptr *QSerialPort) DataBitsChanged(dataBits QSerialPort__DataBits) {
	defer qt.Recovering("QSerialPort::dataBitsChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_DataBitsChanged(ptr.Pointer(), C.int(dataBits))
	}
}

//export callbackQSerialPort_DataTerminalReadyChanged
func callbackQSerialPort_DataTerminalReadyChanged(ptr unsafe.Pointer, ptrName *C.char, set C.int) {
	defer qt.Recovering("callback QSerialPort::dataTerminalReadyChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "dataTerminalReadyChanged"); signal != nil {
		signal.(func(bool))(int(set) != 0)
	}

}

func (ptr *QSerialPort) ConnectDataTerminalReadyChanged(f func(set bool)) {
	defer qt.Recovering("connect QSerialPort::dataTerminalReadyChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_ConnectDataTerminalReadyChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "dataTerminalReadyChanged", f)
	}
}

func (ptr *QSerialPort) DisconnectDataTerminalReadyChanged() {
	defer qt.Recovering("disconnect QSerialPort::dataTerminalReadyChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectDataTerminalReadyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "dataTerminalReadyChanged")
	}
}

func (ptr *QSerialPort) DataTerminalReadyChanged(set bool) {
	defer qt.Recovering("QSerialPort::dataTerminalReadyChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_DataTerminalReadyChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(set)))
	}
}

//export callbackQSerialPort_Error2
func callbackQSerialPort_Error2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QSerialPort::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QSerialPort__SerialPortError))(QSerialPort__SerialPortError(error))
	}

}

func (ptr *QSerialPort) ConnectError2(f func(error QSerialPort__SerialPortError)) {
	defer qt.Recovering("connect QSerialPort::error")

	if ptr.Pointer() != nil {
		C.QSerialPort_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QSerialPort) DisconnectError2() {
	defer qt.Recovering("disconnect QSerialPort::error")

	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

func (ptr *QSerialPort) Error2(error QSerialPort__SerialPortError) {
	defer qt.Recovering("QSerialPort::error")

	if ptr.Pointer() != nil {
		C.QSerialPort_Error2(ptr.Pointer(), C.int(error))
	}
}

//export callbackQSerialPort_FlowControlChanged
func callbackQSerialPort_FlowControlChanged(ptr unsafe.Pointer, ptrName *C.char, flow C.int) {
	defer qt.Recovering("callback QSerialPort::flowControlChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "flowControlChanged"); signal != nil {
		signal.(func(QSerialPort__FlowControl))(QSerialPort__FlowControl(flow))
	}

}

func (ptr *QSerialPort) ConnectFlowControlChanged(f func(flow QSerialPort__FlowControl)) {
	defer qt.Recovering("connect QSerialPort::flowControlChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_ConnectFlowControlChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flowControlChanged", f)
	}
}

func (ptr *QSerialPort) DisconnectFlowControlChanged() {
	defer qt.Recovering("disconnect QSerialPort::flowControlChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectFlowControlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flowControlChanged")
	}
}

func (ptr *QSerialPort) FlowControlChanged(flow QSerialPort__FlowControl) {
	defer qt.Recovering("QSerialPort::flowControlChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_FlowControlChanged(ptr.Pointer(), C.int(flow))
	}
}

func (ptr *QSerialPort) Flush() bool {
	defer qt.Recovering("QSerialPort::flush")

	if ptr.Pointer() != nil {
		return C.QSerialPort_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPort) IsSequential() bool {
	defer qt.Recovering("QSerialPort::isSequential")

	if ptr.Pointer() != nil {
		return C.QSerialPort_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPort) Open(mode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QSerialPort::open")

	if ptr.Pointer() != nil {
		return C.QSerialPort_Open(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

//export callbackQSerialPort_ParityChanged
func callbackQSerialPort_ParityChanged(ptr unsafe.Pointer, ptrName *C.char, parity C.int) {
	defer qt.Recovering("callback QSerialPort::parityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "parityChanged"); signal != nil {
		signal.(func(QSerialPort__Parity))(QSerialPort__Parity(parity))
	}

}

func (ptr *QSerialPort) ConnectParityChanged(f func(parity QSerialPort__Parity)) {
	defer qt.Recovering("connect QSerialPort::parityChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_ConnectParityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "parityChanged", f)
	}
}

func (ptr *QSerialPort) DisconnectParityChanged() {
	defer qt.Recovering("disconnect QSerialPort::parityChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectParityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "parityChanged")
	}
}

func (ptr *QSerialPort) ParityChanged(parity QSerialPort__Parity) {
	defer qt.Recovering("QSerialPort::parityChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_ParityChanged(ptr.Pointer(), C.int(parity))
	}
}

func (ptr *QSerialPort) PinoutSignals() QSerialPort__PinoutSignal {
	defer qt.Recovering("QSerialPort::pinoutSignals")

	if ptr.Pointer() != nil {
		return QSerialPort__PinoutSignal(C.QSerialPort_PinoutSignals(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) PortName() string {
	defer qt.Recovering("QSerialPort::portName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSerialPort_PortName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSerialPort) ReadBufferSize() int64 {
	defer qt.Recovering("QSerialPort::readBufferSize")

	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) ReadLineData(data string, maxSize int64) int64 {
	defer qt.Recovering("QSerialPort::readLineData")

	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_ReadLineData(ptr.Pointer(), C.CString(data), C.longlong(maxSize)))
	}
	return 0
}

//export callbackQSerialPort_RequestToSendChanged
func callbackQSerialPort_RequestToSendChanged(ptr unsafe.Pointer, ptrName *C.char, set C.int) {
	defer qt.Recovering("callback QSerialPort::requestToSendChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestToSendChanged"); signal != nil {
		signal.(func(bool))(int(set) != 0)
	}

}

func (ptr *QSerialPort) ConnectRequestToSendChanged(f func(set bool)) {
	defer qt.Recovering("connect QSerialPort::requestToSendChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_ConnectRequestToSendChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "requestToSendChanged", f)
	}
}

func (ptr *QSerialPort) DisconnectRequestToSendChanged() {
	defer qt.Recovering("disconnect QSerialPort::requestToSendChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectRequestToSendChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "requestToSendChanged")
	}
}

func (ptr *QSerialPort) RequestToSendChanged(set bool) {
	defer qt.Recovering("QSerialPort::requestToSendChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_RequestToSendChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(set)))
	}
}

func (ptr *QSerialPort) SetPort(serialPortInfo QSerialPortInfo_ITF) {
	defer qt.Recovering("QSerialPort::setPort")

	if ptr.Pointer() != nil {
		C.QSerialPort_SetPort(ptr.Pointer(), PointerFromQSerialPortInfo(serialPortInfo))
	}
}

func (ptr *QSerialPort) SetPortName(name string) {
	defer qt.Recovering("QSerialPort::setPortName")

	if ptr.Pointer() != nil {
		C.QSerialPort_SetPortName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSerialPort) SetReadBufferSize(size int64) {
	defer qt.Recovering("QSerialPort::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QSerialPort_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

//export callbackQSerialPort_StopBitsChanged
func callbackQSerialPort_StopBitsChanged(ptr unsafe.Pointer, ptrName *C.char, stopBits C.int) {
	defer qt.Recovering("callback QSerialPort::stopBitsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stopBitsChanged"); signal != nil {
		signal.(func(QSerialPort__StopBits))(QSerialPort__StopBits(stopBits))
	}

}

func (ptr *QSerialPort) ConnectStopBitsChanged(f func(stopBits QSerialPort__StopBits)) {
	defer qt.Recovering("connect QSerialPort::stopBitsChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_ConnectStopBitsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stopBitsChanged", f)
	}
}

func (ptr *QSerialPort) DisconnectStopBitsChanged() {
	defer qt.Recovering("disconnect QSerialPort::stopBitsChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectStopBitsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stopBitsChanged")
	}
}

func (ptr *QSerialPort) StopBitsChanged(stopBits QSerialPort__StopBits) {
	defer qt.Recovering("QSerialPort::stopBitsChanged")

	if ptr.Pointer() != nil {
		C.QSerialPort_StopBitsChanged(ptr.Pointer(), C.int(stopBits))
	}
}

func (ptr *QSerialPort) WaitForBytesWritten(msecs int) bool {
	defer qt.Recovering("QSerialPort::waitForBytesWritten")

	if ptr.Pointer() != nil {
		return C.QSerialPort_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSerialPort) WaitForReadyRead(msecs int) bool {
	defer qt.Recovering("QSerialPort::waitForReadyRead")

	if ptr.Pointer() != nil {
		return C.QSerialPort_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSerialPort) WriteData(data string, maxSize int64) int64 {
	defer qt.Recovering("QSerialPort::writeData")

	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_WriteData(ptr.Pointer(), C.CString(data), C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QSerialPort) DestroyQSerialPort() {
	defer qt.Recovering("QSerialPort::~QSerialPort")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QSerialPort_DestroyQSerialPort(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSerialPort_Pos
func callbackQSerialPort_Pos(ptr unsafe.Pointer, ptrName *C.char) C.longlong {
	defer qt.Recovering("callback QSerialPort::pos")

	if signal := qt.GetSignal(C.GoString(ptrName), "pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQSerialPortFromPointer(ptr).PosDefault())
}

func (ptr *QSerialPort) ConnectPos(f func() int64) {
	defer qt.Recovering("connect QSerialPort::pos")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "pos", f)
	}
}

func (ptr *QSerialPort) DisconnectPos() {
	defer qt.Recovering("disconnect QSerialPort::pos")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "pos")
	}
}

func (ptr *QSerialPort) Pos() int64 {
	defer qt.Recovering("QSerialPort::pos")

	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) PosDefault() int64 {
	defer qt.Recovering("QSerialPort::pos")

	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSerialPort_Reset
func callbackQSerialPort_Reset(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSerialPort::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).ResetDefault()))
}

func (ptr *QSerialPort) ConnectReset(f func() bool) {
	defer qt.Recovering("connect QSerialPort::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QSerialPort) DisconnectReset() {
	defer qt.Recovering("disconnect QSerialPort::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

func (ptr *QSerialPort) Reset() bool {
	defer qt.Recovering("QSerialPort::reset")

	if ptr.Pointer() != nil {
		return C.QSerialPort_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPort) ResetDefault() bool {
	defer qt.Recovering("QSerialPort::reset")

	if ptr.Pointer() != nil {
		return C.QSerialPort_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSerialPort_Seek
func callbackQSerialPort_Seek(ptr unsafe.Pointer, ptrName *C.char, pos C.longlong) C.int {
	defer qt.Recovering("callback QSerialPort::seek")

	if signal := qt.GetSignal(C.GoString(ptrName), "seek"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos))))
	}

	return C.int(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).SeekDefault(int64(pos))))
}

func (ptr *QSerialPort) ConnectSeek(f func(pos int64) bool) {
	defer qt.Recovering("connect QSerialPort::seek")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "seek", f)
	}
}

func (ptr *QSerialPort) DisconnectSeek() {
	defer qt.Recovering("disconnect QSerialPort::seek")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "seek")
	}
}

func (ptr *QSerialPort) Seek(pos int64) bool {
	defer qt.Recovering("QSerialPort::seek")

	if ptr.Pointer() != nil {
		return C.QSerialPort_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QSerialPort) SeekDefault(pos int64) bool {
	defer qt.Recovering("QSerialPort::seek")

	if ptr.Pointer() != nil {
		return C.QSerialPort_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQSerialPort_Size
func callbackQSerialPort_Size(ptr unsafe.Pointer, ptrName *C.char) C.longlong {
	defer qt.Recovering("callback QSerialPort::size")

	if signal := qt.GetSignal(C.GoString(ptrName), "size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQSerialPortFromPointer(ptr).SizeDefault())
}

func (ptr *QSerialPort) ConnectSize(f func() int64) {
	defer qt.Recovering("connect QSerialPort::size")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "size", f)
	}
}

func (ptr *QSerialPort) DisconnectSize() {
	defer qt.Recovering("disconnect QSerialPort::size")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "size")
	}
}

func (ptr *QSerialPort) Size() int64 {
	defer qt.Recovering("QSerialPort::size")

	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) SizeDefault() int64 {
	defer qt.Recovering("QSerialPort::size")

	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSerialPort_TimerEvent
func callbackQSerialPort_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSerialPort::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSerialPortFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSerialPort) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSerialPort::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSerialPort) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSerialPort::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QSerialPort) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSerialPort::timerEvent")

	if ptr.Pointer() != nil {
		C.QSerialPort_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSerialPort) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSerialPort::timerEvent")

	if ptr.Pointer() != nil {
		C.QSerialPort_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSerialPort_ChildEvent
func callbackQSerialPort_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSerialPort::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSerialPortFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSerialPort) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSerialPort::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSerialPort) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSerialPort::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QSerialPort) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSerialPort::childEvent")

	if ptr.Pointer() != nil {
		C.QSerialPort_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSerialPort) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSerialPort::childEvent")

	if ptr.Pointer() != nil {
		C.QSerialPort_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSerialPort_ConnectNotify
func callbackQSerialPort_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSerialPort::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSerialPortFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSerialPort) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSerialPort::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QSerialPort) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSerialPort::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QSerialPort) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSerialPort::connectNotify")

	if ptr.Pointer() != nil {
		C.QSerialPort_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSerialPort) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSerialPort::connectNotify")

	if ptr.Pointer() != nil {
		C.QSerialPort_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSerialPort_CustomEvent
func callbackQSerialPort_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSerialPort::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSerialPortFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSerialPort) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSerialPort::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSerialPort) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSerialPort::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QSerialPort) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSerialPort::customEvent")

	if ptr.Pointer() != nil {
		C.QSerialPort_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSerialPort) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSerialPort::customEvent")

	if ptr.Pointer() != nil {
		C.QSerialPort_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSerialPort_DeleteLater
func callbackQSerialPort_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSerialPort::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSerialPortFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSerialPort) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSerialPort::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QSerialPort) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSerialPort::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QSerialPort) DeleteLater() {
	defer qt.Recovering("QSerialPort::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QSerialPort_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSerialPort) DeleteLaterDefault() {
	defer qt.Recovering("QSerialPort::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QSerialPort_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSerialPort_DisconnectNotify
func callbackQSerialPort_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSerialPort::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSerialPortFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSerialPort) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSerialPort::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QSerialPort) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSerialPort::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QSerialPort) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSerialPort::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSerialPort) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSerialPort::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSerialPort_Event
func callbackQSerialPort_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSerialPort::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QSerialPort) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QSerialPort::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QSerialPort) DisconnectEvent() {
	defer qt.Recovering("disconnect QSerialPort::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QSerialPort) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSerialPort::event")

	if ptr.Pointer() != nil {
		return C.QSerialPort_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSerialPort) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSerialPort::event")

	if ptr.Pointer() != nil {
		return C.QSerialPort_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSerialPort_EventFilter
func callbackQSerialPort_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSerialPort::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QSerialPort) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QSerialPort::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QSerialPort) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QSerialPort::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QSerialPort) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSerialPort::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSerialPort_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSerialPort) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSerialPort::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSerialPort_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSerialPort_MetaObject
func callbackQSerialPort_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSerialPort::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSerialPortFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSerialPort) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSerialPort::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QSerialPort) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSerialPort::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QSerialPort) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSerialPort::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSerialPort_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSerialPort) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSerialPort::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSerialPort_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSerialPortInfo struct {
	ptr unsafe.Pointer
}

type QSerialPortInfo_ITF interface {
	QSerialPortInfo_PTR() *QSerialPortInfo
}

func (p *QSerialPortInfo) QSerialPortInfo_PTR() *QSerialPortInfo {
	return p
}

func (p *QSerialPortInfo) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSerialPortInfo) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSerialPortInfo(ptr QSerialPortInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSerialPortInfo_PTR().Pointer()
	}
	return nil
}

func NewQSerialPortInfoFromPointer(ptr unsafe.Pointer) *QSerialPortInfo {
	var n = new(QSerialPortInfo)
	n.SetPointer(ptr)
	return n
}

func newQSerialPortInfoFromPointer(ptr unsafe.Pointer) *QSerialPortInfo {
	var n = NewQSerialPortInfoFromPointer(ptr)
	return n
}

func (ptr *QSerialPortInfo) Swap(other QSerialPortInfo_ITF) {
	defer qt.Recovering("QSerialPortInfo::swap")

	if ptr.Pointer() != nil {
		C.QSerialPortInfo_Swap(ptr.Pointer(), PointerFromQSerialPortInfo(other))
	}
}

func NewQSerialPortInfo() *QSerialPortInfo {
	defer qt.Recovering("QSerialPortInfo::QSerialPortInfo")

	return newQSerialPortInfoFromPointer(C.QSerialPortInfo_NewQSerialPortInfo())
}

func NewQSerialPortInfo2(port QSerialPort_ITF) *QSerialPortInfo {
	defer qt.Recovering("QSerialPortInfo::QSerialPortInfo")

	return newQSerialPortInfoFromPointer(C.QSerialPortInfo_NewQSerialPortInfo2(PointerFromQSerialPort(port)))
}

func NewQSerialPortInfo4(other QSerialPortInfo_ITF) *QSerialPortInfo {
	defer qt.Recovering("QSerialPortInfo::QSerialPortInfo")

	return newQSerialPortInfoFromPointer(C.QSerialPortInfo_NewQSerialPortInfo4(PointerFromQSerialPortInfo(other)))
}

func NewQSerialPortInfo3(name string) *QSerialPortInfo {
	defer qt.Recovering("QSerialPortInfo::QSerialPortInfo")

	return newQSerialPortInfoFromPointer(C.QSerialPortInfo_NewQSerialPortInfo3(C.CString(name)))
}

func (ptr *QSerialPortInfo) Description() string {
	defer qt.Recovering("QSerialPortInfo::description")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSerialPortInfo_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSerialPortInfo) HasProductIdentifier() bool {
	defer qt.Recovering("QSerialPortInfo::hasProductIdentifier")

	if ptr.Pointer() != nil {
		return C.QSerialPortInfo_HasProductIdentifier(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPortInfo) HasVendorIdentifier() bool {
	defer qt.Recovering("QSerialPortInfo::hasVendorIdentifier")

	if ptr.Pointer() != nil {
		return C.QSerialPortInfo_HasVendorIdentifier(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPortInfo) IsNull() bool {
	defer qt.Recovering("QSerialPortInfo::isNull")

	if ptr.Pointer() != nil {
		return C.QSerialPortInfo_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPortInfo) Manufacturer() string {
	defer qt.Recovering("QSerialPortInfo::manufacturer")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSerialPortInfo_Manufacturer(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSerialPortInfo) PortName() string {
	defer qt.Recovering("QSerialPortInfo::portName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSerialPortInfo_PortName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSerialPortInfo) SerialNumber() string {
	defer qt.Recovering("QSerialPortInfo::serialNumber")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSerialPortInfo_SerialNumber(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSerialPortInfo) SystemLocation() string {
	defer qt.Recovering("QSerialPortInfo::systemLocation")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSerialPortInfo_SystemLocation(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSerialPortInfo) DestroyQSerialPortInfo() {
	defer qt.Recovering("QSerialPortInfo::~QSerialPortInfo")

	if ptr.Pointer() != nil {
		C.QSerialPortInfo_DestroyQSerialPortInfo(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSerialPortInfo) IsBusy() bool {
	defer qt.Recovering("QSerialPortInfo::isBusy")

	if ptr.Pointer() != nil {
		return C.QSerialPortInfo_IsBusy(ptr.Pointer()) != 0
	}
	return false
}
