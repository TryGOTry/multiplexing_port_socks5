package protocols

import "regexp"

// NewHTTPProtocol initializes a Protocol with a HTTP signature.
func NewHTTPProtocol(targetAddress string) *Protocol {
	regexes := []*regexp.Regexp{
		regexp.MustCompile("^[A-Z]+ .+ HTTP/"),
	}

	return &Protocol{
		Name:                    "HTTP",
		Target:                  targetAddress,
		MatchRegexes:            regexes,
		NoComparisonBeforeBytes: 11, // GET / HTTP/
	}
}
