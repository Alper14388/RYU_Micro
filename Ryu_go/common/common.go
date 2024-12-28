package common

type PacketInWrapper struct {
	Buffer uint32 `json:"Buffer"`
	Length uint16 `json:"Length"`
	Reason uint8  `json:"Reason"`
	Table  uint8  `json:"Table"`
	Cookie uint64 `json:"Cookie"`
	Match  struct {
		Type   uint16 `json:"Type"`
		Fields []struct {
			Class uint16      `json:"Class"`
			Type  uint8       `json:"Type"`
			Value string      `json:"Value"` // Base64 string
			Mask  interface{} `json:"Mask"`  // Optional, null olabilir
		} `json:"Fields"`
	} `json:"Match"`
	Data string `json:"Data"` // Base64 string
}
