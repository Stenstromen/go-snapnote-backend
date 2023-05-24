package model

type FormData struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Delta struct {
		Ops []Operation `json:"ops,omitempty"`
	} `json:"delta,omitempty"`
}

type Operation struct {
	Insert     interface{} `json:"insert,omitempty"`
	Attributes struct {
		Font       string `json:"font,omitempty"`
		Size       string `json:"size,omitempty"`
		Header     int    `json:"header,omitempty"`
		Color      string `json:"color,omitempty"`
		Background string `json:"background,omitempty"`
		Blockquote bool   `json:"blockquote,omitempty"`
		CodeBlock  bool   `json:"code-block,omitempty"`
		Bold       bool   `json:"bold,omitempty"`
		Italic     bool   `json:"italic,omitempty"`
		Underline  bool   `json:"underline,omitempty"`
		Strike     bool   `json:"strike,omitempty"`
		List       string `json:"list,omitempty"`
		Align      string `json:"align,omitempty"`
		Link       string `json:"link,omitempty"`
	} `json:"attributes,omitempty"`
	Image map[string]string `json:"image,omitempty"`
}
