package events

type eventRecvMsgType string
type eventSendMsgType string

const (
	EventConnected eventRecvMsgType = "connected" // initiates the handshake
	EventStart     eventRecvMsgType = "start"     // contains stream metadata
	EventRMedia    eventRecvMsgType = "media"     // contains base64 encoded audio data
	EventRStop     eventRecvMsgType = "stop"      // terminates the stream
	EventDTMF      eventRecvMsgType = "dtmf"      // na
	EventRMark     eventRecvMsgType = "mark"      // end of input
	EventSMedia    eventSendMsgType = "media"     // contains base64 encoded audio data
	EventSMark     eventSendMsgType = "mark"      // end of input
	EventClear     eventSendMsgType = "clear"     // resets the stream

)
