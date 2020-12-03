package webrtc

// RTPCapabilities represents the capabilities of a transceiver
//
// https://w3c.github.io/webrtc-pc/#rtcrtpcapabilities
type RTPCapabilities struct {
	Codecs           []RTPCodecCapability
	HeaderExtensions []RTPHeaderExtensionCapability
}

// RTPCodecCapability provides information about codec capabilities.
//
// https://w3c.github.io/webrtc-pc/#dictionary-rtcrtpcodeccapability-members
type RTPCodecCapability struct {
	MimeType     string
	ClockRate    uint32
	Channels     uint16
	SDPFmtpLine  string
	RTCPFeedback []RTCPFeedback
}

// RTPHeaderExtensionCapability is used to define a RFC5285 RTP header extension supported by the codec.
//
// https://w3c.github.io/webrtc-pc/#dom-rtcrtpcapabilities-headerextensions
type RTPHeaderExtensionCapability struct {
	URI string
}
