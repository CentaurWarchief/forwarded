package forwarded

// Forwarded is a representation of all information contained in the
// RFC 7239 Forwarded header
type Forwarded struct {
	Host  string
	By    string
	For   string
	Proto string
}
