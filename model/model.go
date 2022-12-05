package model

type Inspector struct {
	Pattern        string   `json:"pattern"`
	GroupPattern   []string `json:"group_pattern"`
	Fixes          []Fix    `json:"fixes"`
	HightLightType string   `json:"hightlight_type"`
}
type Fix struct {
	Name     string   `json:"name"`
	Patterns []string `json:"patterns"`
	Strings  []string `json:"strings"`
}
type Algorithms struct {
	Id               string    `json:"id"`
	BriefDescription string    `json:"brief_description"`
	Inspector        Inspector `json:"inspector"`
}
