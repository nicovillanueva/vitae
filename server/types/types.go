package types

// CVData is the whole data of the CV, unmarshalled in the config phase by the CLI
type CVData struct {
	PDFpayload   []byte
	PathToPDF    string          `yaml:"path_to_pdf"`
	Statement    string          `yaml:"statement"`
	Skills       []string        `yaml:"skills"`
	Achievements []string        `yaml:"achievements"`
	WorkHistory  []WorkSpan      `yaml:"work_history"`
	Education    []EducationSpan `yaml:"education"`
	Hobbies      string          `yaml:"hobbies"`
	References   []Reference     `yaml:"references"`
}

// SpanDetails - is not used as it breaks the swagger docs generation :(
type SpanDetails struct {
	Title       string `json:"title"`
	Location    string `json:"location"`
	Span        string `json:"span"`
	Description string `json:"description"`
}

// WorkSpan is a span of time in which I worked somewhere
type WorkSpan struct {
	Title       string `json:"title"`
	Location    string `json:"location"`
	Span        string `json:"span"`
	Description string `json:"description"`
	Company     string `json:"company"`
}

// EducationSpan is a span of time in which I attended some course or educational thingy
type EducationSpan struct {
	Title       string `json:"title"`
	Location    string `json:"location"`
	Span        string `json:"span"`
	Description string `json:"description"`
	Institute   string `json:"institute"`
}

// Reference ...
type Reference struct {
	Name    string `json:"name"`
	Company string `json:"company"`
	Contact string `json:"contact"`
}
