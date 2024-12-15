package common

type PacketInWrapper struct {
	OFPPacketIn struct {
		BufferID uint32 `json:"buffer_id"`
		Data     string `json:"data"`
		InPort   uint32 `json:"in_port"`
		Reason   uint8  `json:"reason"`
		TotalLen uint16 `json:"total_len"`
	} `json:"OFPPacketIn"`
	DatapathID uint32 `json:"dpid"`
}
