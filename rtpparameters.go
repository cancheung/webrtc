// +build !js

package webrtc

import "strings"

// RTPParameters is a list of negotiated codecs and header extensions
//
// https://w3c.github.io/webrtc-pc/#dictionary-rtcrtpparameters-members
type RTPParameters struct {
	HeaderExtensions []RTPHeaderExtensionParameter
	Codecs           []RTPCodecParameters
}

// RTPCodecParameters is a sequence containing the media codecs that an RtpSender
// will choose from, as well as entries for RTX, RED and FEC mechanisms. This also
// includes the PayloadType that has been negotiated
//
// https://w3c.github.io/webrtc-pc/#rtcrtpcodecparameters
type RTPCodecParameters struct {
	RTPCodecCapability
	PayloadType PayloadType

	statsID string
}

// RTPHeaderExtensionParameter represents a negotiated RFC5285 RTP header extension.
//
// https://w3c.github.io/webrtc-pc/#dictionary-rtcrtpheaderextensionparameters-members
type RTPHeaderExtensionParameter struct {
	URI string
	ID  int
}

// Do a fuzzy find for a codec in the list of codecs
// Used for lookup up a codec in an existing list to find a match
func codecParametersFuzzySearch(needle RTPCodecParameters, haystack []RTPCodecParameters) (RTPCodecParameters, error) {
	// First attempt to match on MimeType + SDPFmtpLine
	for _, c := range haystack {
		if strings.EqualFold(c.RTPCodecCapability.MimeType, needle.RTPCodecCapability.MimeType) &&
			c.RTPCodecCapability.SDPFmtpLine == needle.RTPCodecCapability.SDPFmtpLine {
			return c, nil
		}
	}

	// Fallback to just MimeType
	for _, c := range haystack {
		if strings.EqualFold(c.RTPCodecCapability.MimeType, needle.RTPCodecCapability.MimeType) {
			return c, nil
		}
	}

	return RTPCodecParameters{}, ErrCodecNotFound
}
