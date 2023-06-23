package emailnormalizer

import "strings"

// GoogleRule : email normalization rule for Google
type GoogleRule struct {
}

func (rule *GoogleRule) ProcessUsername(username string) string {
	result := strings.ToLower(username)
	result = strings.Replace(result, ".", "", -1)

	plusSignIndex := strings.Index(result, "+")
	if plusSignIndex != -1 {
		result = result[0:plusSignIndex]
	}

	return result
}

func (rule *GoogleRule) ProcessDomain(domain string) string {
	switch domain {
	case "google.com":
		return domain
	default:
		return "gmail.com" // googlemail.com/gmail.com => gmail.com
	}
}
