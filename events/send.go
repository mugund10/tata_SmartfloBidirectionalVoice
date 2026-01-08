package events

// payload must be encoded to mulaw with sample rate of 8000
// must be encoded with base64 pcm mono audio

// outgoing media metadata and payload
type mediaout struct {
	Chunk   string `json:"chunk"`
	Payload string `json:"payload"`
}

// default outgoing payload structure
type outgoing struct {
	Event eventSendMsgType `json:"event"`
	StSid string           `json:"streamSid"`
	Media *mediaout        `json:"media"`
	Mark  *mark            `json:"mark"`
}
