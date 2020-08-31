// Package sharetunnel implements a client library for https://sharetunnel.work
//
// In addition to providing the sharetunnel client which will forward requests
// from subdomain.sharetunnel.work to a port on localhost. This package also
// provides an implementation of net.Listener which exposes connections from
// sharetunnel. This enables users to serve http requests directly, without
// listening to a port on localhost.
//
//   // Setup a listener for sharetunnel
//   listener, err := sharetunnel.Listen(sharetunnel.Options{})
//
//   // Create your server...
//   server := http.Server{
//       Handler: http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
//           w.WriteHeader(200)
//           ...
//       })
//   }
//
//   // Handle request from sharetunnel
//   server.Serve(listener)
package sharetunnel
