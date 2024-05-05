package structures

const (
	DEFAULT_TYPE    = "core"
	DEFAULT_VERSION = "1.0.0"
)

type Request struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Hash    string `json:"hash"`
}

func (r Request) UpdateRequestPayloadWithDefaults() {
	if r.Type == "" {
		r.Type = DEFAULT_TYPE
	}
	if r.Version == "" {
		r.Version = DEFAULT_VERSION
	}
}
