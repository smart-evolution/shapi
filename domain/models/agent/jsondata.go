package agent

// Type1DataJSON - entity representing agent data
type Type1DataJSON struct {
	Time        []string `json:"time"`
	Temperature []string `json:"temperature"`
	Presence    []string `json:"presence"`
	Gas         []string `json:"gas"`
	Sound       []string `json:"sound"`
}
