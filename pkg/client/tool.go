package client

type Tool struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Supports    int    `json:"supports,omitempty"`
	Latest      string `json:"latest,omitempty"`
}

type ToolVersion struct {
}
