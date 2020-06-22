package label

type Instance struct {
	Name        string
	Description string
}

var (
	Server = Instance{
		Name:        "server",
		Description: "Associated tests will test the conformance of an XDS server, such as go-control-plane.",
	}
	XdsV2 = Instance{
		Name:        "v2",
		Description: "Associated tests will test XDS v2.",
	}
	XdsV3 = Instance{
		Name:        "v3",
		Description: "Associated tests will test XDS v3.",
	}
	XdsV4 = Instance{
		Name:        "v3",
		Description: "Associated tests will test XDS v4.",
	}
	Client = Instance{
		Name:        "client",
		Description: "Associated tests will test the conformance of an XDS client, such as Envoy.",
	}
)
