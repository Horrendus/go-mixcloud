package mixcloud

type Link struct {
	PreviousURL string `json:"previous,omitempty"`
	NextURL		string `json:"next,omitempty"`
}
