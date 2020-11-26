package handler

//Action action type
type Action struct {
	Device     string `json:"device"`
	ActionType string `json:"type"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
}
