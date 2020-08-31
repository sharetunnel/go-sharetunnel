package sharetunnel

import (
	"fmt"
	"io"
	"net"
)

// sharetunnel forwards remote requests to a port on localhost
type sharetunnel struct {
	listener  *Listener
	localAddr string
}

// New returns a sharetunnel forwarding requests to port on host
//
// host defaults to 'localhost', and options defaults to using sharetunnel.work
func New(port int, host string, options Options) (*sharetunnel, error) {
	if host == "" {
		host = "localhost"
	}

	l, err := Listen(options)
	if err != nil {
		return nil, err
	}

	lt := &sharetunnel{
		listener:  l,
		localAddr: fmt.Sprintf("%s:%d", host, port),
	}
	go lt.listen()
	return lt, nil
}

// URL returns the URL at which the sharetunnel is exposed
func (lt *sharetunnel) URL() string {
	return lt.listener.URL()
}

func (lt *sharetunnel) listen() {
	for {
		remoteConn, err := lt.listener.Accept()
		if err != nil {
			break
		}

		go lt.forward(remoteConn)
	}
}

func (lt *sharetunnel) forward(remoteConn net.Conn) {
	localConn, err := net.Dial("tcp", lt.localAddr)
	if err != nil {
		remoteConn.Close()
		return
	}

	go func() {
		io.Copy(remoteConn, localConn)
		remoteConn.Close()
	}()
	go func() {
		io.Copy(localConn, remoteConn)
		localConn.Close()
	}()
}

// Close the sharetunnel aborting all connections
func (lt *sharetunnel) Close() error {
	return lt.listener.Close()
}
