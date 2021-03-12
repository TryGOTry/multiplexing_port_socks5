package protocols

import "regexp"

// Protocol is the implementation of a protocol signature.
type Protocol struct {
	Name                    string           // the protocol name for auditing
	Target                  string           // the proxy target
	MatchStartBytes         [][]byte         // the bytestrings by which to match this protocol (prefixes)
	MatchBytes              [][]byte         // the bytestrings by which to match this protocol (contains)
	MatchRegexes            []*regexp.Regexp // the regexes by which to match this protocol
	NoComparisonBeforeBytes int              // we know we won't match before this many bytes, set to 0 to ignore
	NoComparisonAfterBytes  int              // we know we won't match after this many bytes, set to 0 to ignore
}
