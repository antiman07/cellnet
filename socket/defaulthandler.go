package socket

import "github.com/davyxu/cellnet"

// socket.EncodePacketHandler -> socket.MsgLogHandler -> socket.WritePacketHandler
func BuildSendHandler(useMsgLog bool) cellnet.EventHandler {

	if useMsgLog {

		return cellnet.LinkHandler(NewEncodePacketHandler(), NewMsgLogHandler(), NewWritePacketHandler())
	} else {

		return cellnet.LinkHandler(NewEncodePacketHandler(), NewWritePacketHandler())
	}

}

// socket.ReadPacketHandler -> socket.MsgLogHandler -> socket.DispatcherHandler -> socket.DecodePacketHandler -> socket.CallbackHandler
func BuildWriteHandler(useMsgLog bool, dispatcher *cellnet.DispatcherHandler, q cellnet.EventQueue) cellnet.EventHandler {

	if useMsgLog {

		return cellnet.LinkHandler(NewReadPacketHandler(q), NewMsgLogHandler(), dispatcher)

	} else {
		return cellnet.LinkHandler(NewReadPacketHandler(q), dispatcher)
	}

}