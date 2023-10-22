package test

type OspfV6Config struct {
	OspfV6        OspfV6            `xml:"ipv6-ospf,omitempty"`
	InterfaceList []OspfV6Interface `xml:"interface,omitempty"`
}

type OspfV6 struct {
	Router    OspfV6Router `xml:"router,omitempty"`
	DryRun    string       `xml:"dry-run,attr,omitempty"`
	Operation string       `xml:"operation,attr,omitempty"`
}

type OspfV6Router struct {
	RouterId  string `xml:"ROUTER-ID,attr,omitempty"`
	DryRun    string `xml:"dry-run,attr,omitempty"`
	Operation string `xml:"operation,attr,omitempty"`
}

type OspfV6Interface struct {
	Name          string                       `xml:"INTERFACE-NAME,attr,omitempty"`
	Area          OspfV6InterfaceArea          `xml:"ipv6-ospf-area,omitempty"`
	Network       OspfV6InterfaceNetwork       `xml:"ipv6-ospf-network,omitempty"`
	Cost          OspfV6InterfaceCost          `xml:"ipv6-ospf-cost,omitempty"`
	HelloInterval OspfV6InterfaceHelloInterval `xml:"ipv6-ospf-hello-interval,omitempty"`
}

type OspfV6InterfaceArea struct {
	Name      string `xml:"INTERFACE-NAME,attr,omitempty"`
	Area      string `xml:"AREA,attr,omitempty"`
	DryRun    string `xml:"dry-run,attr,omitempty"`
	Operation string `xml:"operation,attr,omitempty"`
}

type OspfV6InterfaceNetwork struct {
	Name      string `xml:"INTERFACE-NAME,attr,omitempty"`
	Network   string `xml:"NETWORK,attr,omitempty"`
	DryRun    string `xml:"dry-run,attr,omitempty"`
	Operation string `xml:"operation,attr,omitempty"`
}

type OspfV6InterfaceCost struct {
	Name      string `xml:"INTERFACE-NAME,attr,omitempty"`
	Cost      int    `xml:"COST,attr,omitempty"`
	DryRun    string `xml:"dry-run,attr,omitempty"`
	Operation string `xml:"operation,attr,omitempty"`
}

type OspfV6InterfaceHelloInterval struct {
	Name          string `xml:"INTERFACE-NAME,attr,omitempty"`
	HelloInterval string `xml:"HELLO-INTERVAL,attr,omitempty"`
	DryRun        string `xml:"dry-run,attr,omitempty"`
	Operation     string `xml:"operation,attr,omitempty"`
}
