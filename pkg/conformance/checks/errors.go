package checks

import (
	"fmt"
)

func connectionFailure(addr string) error {
	return fmt.Errorf("failed to conect to the XDS server at %q", addr)
}

func requestFailure(err error) error {
	return fmt.Errorf("failed to send discovery request: %v", err)
}

func responseFailure(err error) error {
	return fmt.Errorf("failed to recieve a discovery response: %v", err)
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
