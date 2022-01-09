package swagger

// swagger request model
type Message struct {
	SrcId string `json:"srcId"`
	DstId string `json:"dstId"`

	MsgType int    `json:"msgType"`
	Content string `json:"content"`

	ResType int    `json:"resType"`
	ResUrl  string `json:"resUrl"`
}
