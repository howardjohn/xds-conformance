package label

type Instance struct {
	Name        string
	Description string
}

var (
	Server = Instance{
		Name:        "Server",
		Description: "Associated tests will test the conformance of an XDS server, such as go-control-plane.",
	}
	Client = Instance{
		Name:        "Client",
		Description: "Associated tests will test the conformance of an XDS client, such as Envoy.",
	}
)
