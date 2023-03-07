//go:build !windows
// +build !windows

package opts

// defaultHost constant defines the default host string used by docker on other hosts than Windows
const defaultHost = "unix://" + defaultUnixSocket

// defaultHTTPHost Default HTTP Host used if only port is provided to -H flag e.g. dockerd -H tcp://:8080
const defaultHTTPHost = "localhost"
