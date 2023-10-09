package alumni_mod

type IsAlumniRequest struct {
	Username string `json:"username"`
}

type IsAlumniResponse struct {
	IsAlumni bool   `json:"is_alumni"`
	Message  string `json:"message"`
}
