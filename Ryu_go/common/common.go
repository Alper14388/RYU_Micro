package common

type PacketInWrapper struct {
	OFPPacketIn struct {
		BufferID uint32 `json:"buffer_id"`
		Cookie   uint64 `json:"cookie"`
		Data     string `json:"data"`
		Match    struct {
			OFPMatch struct {
				OxmFields []struct {
					OXMTlv struct {
						Field string      `json:"field"`
						Value interface{} `json:"value"`
						Mask  interface{} `json:"mask"`
					} `json:"OXMTlv"`
				} `json:"oxm_fields"`
				Length uint16 `json:"length"`
				Type   uint16 `json:"type"`
			} `json:"OFPMatch"`
		} `json:"match"`
		Reason   uint8  `json:"reason"`
		TableID  uint8  `json:"table_id"`
		TotalLen uint16 `json:"total_len"`
	} `json:"OFPPacketIn"`
	DatapathID uint64 `json:"dpid"`
}
