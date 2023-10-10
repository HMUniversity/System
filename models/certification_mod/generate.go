package certification_mod

type GenerationRequest struct {
	Programme string `json:"programme"`
	Receiver  string `json:"receiver"`
	Degree    string `json:"degree"`
}
