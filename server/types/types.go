package types

// CVData is the whole data of the CV, unmarshalled by Viper
type CVData struct {
	CVinPDF      []byte
	Statement    string          `mapstructure:"statement"`
	Skills       []string        `mapstructure:"skills"`
	Achievements []string        `mapstructure:"achievements"`
	WorkHistory  []WorkSpan      `mapstructure:"work_history"`
	Education    []EducationSpan `mapstructure:"education"`
	Hobbies      string          `mapstructure:"hobbies"`
	References   []Reference     `mapstructure:"references"`
}

// TODO: Extend other spans from this one
// AbstractSpan ...
// type AbstractSpan struct {
// 	Title       string `json:"title"`
// 	Location    string `json:"location"`
// 	Span        string `json:"span"`
// 	Description string `json:"description"`
// }

// WorkSpan ...
type WorkSpan struct {
	// AbstractSpan
	Title       string `json:"title"`
	Location    string `json:"location"`
	Span        string `json:"span"`
	Description string `json:"description"`
	Company     string `json:"company"`
}

// EducationSpan ...
type EducationSpan struct {
	// AbstractSpan
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
