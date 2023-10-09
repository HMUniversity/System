package utils

func DeRefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
