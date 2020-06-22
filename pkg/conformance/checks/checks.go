package checks

import "github.com/envoyproxy/xds-conformance/pkg/conformance"

var All = []conformance.Check{}

func RegisterCheck(c conformance.Check) conformance.Check {
	All = append(All, c)
	return c
}
