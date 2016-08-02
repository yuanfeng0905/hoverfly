package views

type PayloadViewData struct {
	Data []PayloadView `json:"data"`
}

// PayloadView is used when marshalling and unmarshalling payloads.
type PayloadView struct {
	Response ResponseDetailsView `json:"response"`
	Request  RequestDetailsView  `json:"request"`
}

// PayloadMiddlewareView is used when marshalling and unmarshalling payloads via Middleware.
type PayloadMiddlewareView struct {
	Response              ResponseDetailsView       `json:"response"`
	Request               RequestDetailsView        `json:"request"`
	HoverflyConfiguration HoverflyConfigurationView `json:"hoverflyConfiguration"`
}

type HoverflyConfigurationView struct {
	AdminUrl string `json:"url"`
}

// RequestDetailsView is used when marshalling and unmarshalling RequestDetails
type RequestDetailsView struct {
	Path        string              `json:"path"`
	Method      string              `json:"method"`
	Destination string              `json:"destination"`
	Scheme      string              `json:"scheme"`
	Query       string              `json:"query"`
	Body        string              `json:"body"`
	Headers     map[string][]string `json:"headers"`
}

// ResponseDetailsView is used when marshalling and
// unmarshalling requests. This struct's Body may be Base64
// encoded based on the EncodedBody field.
type ResponseDetailsView struct {
	Status      int                 `json:"status"`
	Body        string              `json:"body"`
	EncodedBody bool                `json:"encodedBody"`
	Headers     map[string][]string `json:"headers"`
}
