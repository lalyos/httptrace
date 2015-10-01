package httptrace

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

var DEBUG_ENV_NAME = "DEBUG"

type debugTransport struct {
	transport http.RoundTripper
}

func NewDebugTransport() *debugTransport {
	return &debugTransport{
		transport: http.DefaultTransport,
	}
}

func (d *debugTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if os.Getenv(DEBUG_ENV_NAME) != "" {
		fmt.Fprintf(os.Stderr, "\n=====> %s\n", r.URL)
		dump, _ := httputil.DumpRequest(r, true)
		fmt.Fprintf(os.Stderr, string(dump[:len(dump)]))
	}
	resp, err := d.transport.RoundTrip(r)
	return resp, err
}

func init() {
	http.DefaultTransport = NewDebugTransport()
}
