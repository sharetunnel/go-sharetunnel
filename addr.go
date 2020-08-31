package sharetunnel

// Network return "sharetunnel"
func (a Addr) Network() string {
	return "sharetunnel"
}

// String returns the URL
func (a Addr) String() string {
	return a.URL
}

// URL returns the URL that the listener is exposed on.
func (l *Listener) URL() string {
	return l.url
}
