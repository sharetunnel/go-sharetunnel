sharetunnel Client Library for Go [![Build Status](https://travis-ci.org/jonasfj/go-sharetunnel.svg?branch=master)](https://travis-ci.org/jonasfj/go-sharetunnel)
=================================

A [sharetunnel.work](https://sharetunnel.work) client library exposing sharetunnel
connections through a `net.Listener` implementation. While sharetunnel only
supports forwarding HTTP(S) connections, this is useful as you can hook it up
to `http.Server` directly. Neat, if writing test-suites or command-line
utilities exposing web-hooks of sharetunnel.

```go
// Setup a listener for sharetunnel
listener, err := sharetunnel.Listen(sharetunnel.Options{})

// Create your server...
server := http.Server{
    Handler: http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(200)
        ...
    })
}

// Handle request from sharetunnel
server.Serve(listener)
```

See [documentation](https://godoc.org/github.com/jonasfj/go-sharetunnel) for
more details.


License
-------
This package is released under [MPLv2](https://www.mozilla.org/en-US/MPL/2.0/).
