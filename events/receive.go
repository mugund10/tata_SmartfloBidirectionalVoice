package events

// format of media in payload
type mediafmt struct {
	Encoding   string `json:"encoding"`
	SampleRate int    `json:"sampleRate"`
	BitRate    int    `json:"bitRate"`
	BitDepth   int    `json:"bitDepth"`
}

// stream metadata
type streamInfo struct {
	AccountSid   string            `json:"accountSid"`
	StreamSid    string            `json:"streamSid"`
	CallSid      string            `json:"callSid"`
	From         string            `json:"from"`
	To           string            `json:"to"`
	Direction    string            `json:"outbound"`
	MediaFormat  mediafmt          `json:"mediaFormat"`
	CustomParams map[string]string `json:"customParameters"`
}

// incoming media metadata and payload
type mediain struct {
	Chunk     string `json:"chunk"`
	Timestamp string `json:"timestamp"`
	Payload   string `json:"payload"`
}

// stoped stream metadata
type stopin struct {
	AccSid  string `json:"accountSid"`
	CallSid string `json:"callSid"`
	Reason  string `json:"reason"`
}

// detected key tone
type dtmf struct {
	Digit string `json:"digit"`
}

// mark metadata
type mark struct {
	Name string `json:"name"`
}

// default incoming payload structure
type incoming struct {
	Event eventRecvMsgType `json:"event"`
	Seq   string       `json:"sequenceNumber"`
	Start *streamInfo  `json:"start"`
	Media *mediain     `json:"media"`
	Stop  *stopin      `json:"stop"`
	StSid string       `json:"streamSid"`
	Dtmf  *dtmf        `json:"dtmf"`
	Mark  *mark        `json:"mark"`
}
