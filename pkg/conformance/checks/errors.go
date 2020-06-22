package checks

import (
	"fmt"

	"github.com/envoyproxy/xds-conformance/pkg/conformance"
)

func connectionFailure(addr string) conformance.TestResult {
	return conformance.TestResult{
		Error:       fmt.Errorf("failed to conect to the XDS server at %q", addr),
		Information: "Verify the XDS server is running and is reachable at the provided address. If TLS is required by the server, ensure TLS configuration is set.",
	}
}

func requestFailure(err error) conformance.TestResult {
	return conformance.TestResult{
		Error: fmt.Errorf("failed to send discovery request: %v", err),
	}
}

func responseFailure(err error) conformance.TestResult {
	return conformance.TestResult{
		Error: fmt.Errorf("failed to recieve a discovery response: %v", err),
	}
}

func AddError(e1, e2 error) error {
	if e1 == nil {
		return e2
	}
	if e2 == nil {
		return e1
	}
	return fmt.Errorf("%v and %v", e1, e2)
}
