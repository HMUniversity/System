package certification_mod

type GenerationRequest struct {
	Programme string `json:"programme"`
	Date      string `json:"date"`
	Receiver  string `json:"receiver"`
	Degree    string `json:"degree"`
}
