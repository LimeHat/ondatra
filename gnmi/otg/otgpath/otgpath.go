/*
Package otgpath is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema. The generated paths are
based on a compressed form of the schema.

This package was generated by ygnmi version: v0.7.4: (ygot: v0.25.4)
using the following YANG input files:
  - models-yang/models/bgp/open-traffic-generator-bgp.yang
  - models-yang/models/discovery/open-traffic-generator-discovery.yang
  - models-yang/models/flow/open-traffic-generator-flow.yang
  - models-yang/models/interface/open-traffic-generator-port.yang
  - models-yang/models/isis/open-traffic-generator-isis.yang
  - models-yang/models/lacp/open-traffic-generator-lacp.yang
  - models-yang/models/lag/open-traffic-generator-lag.yang
  - models-yang/models/lldp/open-traffic-generator-lldp.yang
  - models-yang/models/rsvp/open-traffic-generator-rsvp.yang
  - models-yang/models/types/open-traffic-generator-types.yang

Imported modules were sourced from:
  - models-yang/models/...
*/
package otgpath

import (
	oc "github.com/openconfig/ondatra/gnmi/otg"
	"github.com/openconfig/ondatra/gnmi/otg/bgp"
	"github.com/openconfig/ondatra/gnmi/otg/discovery"
	"github.com/openconfig/ondatra/gnmi/otg/flow"
	"github.com/openconfig/ondatra/gnmi/otg/isis"
	"github.com/openconfig/ondatra/gnmi/otg/lacp"
	"github.com/openconfig/ondatra/gnmi/otg/lag"
	"github.com/openconfig/ondatra/gnmi/otg/lldp"
	"github.com/openconfig/ondatra/gnmi/otg/port"
	"github.com/openconfig/ondatra/gnmi/otg/rsvp"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// RootPath represents the /root YANG schema element.
type RootPath struct {
	*ygnmi.DeviceRootBase
}

// Root returns a root path object from which YANG paths can be constructed.
func Root() *RootPath {
	return &RootPath{ygnmi.NewDeviceRootBase()}
}

// BgpPeerAny (list): Each BGP peer is identified by an arbitrary string
// identifier.
//
//	Defining module:      "open-traffic-generator-bgp"
//	Instantiating module: "open-traffic-generator-bgp"
//	Path from parent:     "bgp-peers/bgp-peer"
//	Path from root:       "/bgp-peers/bgp-peer"
func (n *RootPath) BgpPeerAny() *bgp.BgpPeerPathAny {
	return &bgp.BgpPeerPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"bgp-peers", "bgp-peer"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// BgpPeer (list): Each BGP peer is identified by an arbitrary string
// identifier.
//
//	Defining module:      "open-traffic-generator-bgp"
//	Instantiating module: "open-traffic-generator-bgp"
//	Path from parent:     "bgp-peers/bgp-peer"
//	Path from root:       "/bgp-peers/bgp-peer"
//
//	Name: string
func (n *RootPath) BgpPeer(Name string) *bgp.BgpPeerPath {
	return &bgp.BgpPeerPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"bgp-peers", "bgp-peer"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// FlowAny (list): A flow of packets between one or more internal and external sources
// and one or more internal and external destinations that the target
// is able to track and report statistics on. Each flow is identified by
// an arbitrary string identifier.
//
//	Defining module:      "open-traffic-generator-flow"
//	Instantiating module: "open-traffic-generator-flow"
//	Path from parent:     "flows/flow"
//	Path from root:       "/flows/flow"
func (n *RootPath) FlowAny() *flow.FlowPathAny {
	return &flow.FlowPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"flows", "flow"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Flow (list): A flow of packets between one or more internal and external sources
// and one or more internal and external destinations that the target
// is able to track and report statistics on. Each flow is identified by
// an arbitrary string identifier.
//
//	Defining module:      "open-traffic-generator-flow"
//	Instantiating module: "open-traffic-generator-flow"
//	Path from parent:     "flows/flow"
//	Path from root:       "/flows/flow"
//
//	Name: string
func (n *RootPath) Flow(Name string) *flow.FlowPath {
	return &flow.FlowPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"flows", "flow"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// InterfaceAny (list): An individual interface defined by an OTG.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "interfaces/interface"
//	Path from root:       "/interfaces/interface"
func (n *RootPath) InterfaceAny() *discovery.InterfacePathAny {
	return &discovery.InterfacePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Interface (list): An individual interface defined by an OTG.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "interfaces/interface"
//	Path from root:       "/interfaces/interface"
//
//	Name: string
func (n *RootPath) Interface(Name string) *discovery.InterfacePath {
	return &discovery.InterfacePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// IsisRouterAny (list): Each ISIS router is identified by an arbitrary string
// identifier.
//
//	Defining module:      "open-traffic-generator-isis"
//	Instantiating module: "open-traffic-generator-isis"
//	Path from parent:     "isis-routers/isis-router"
//	Path from root:       "/isis-routers/isis-router"
func (n *RootPath) IsisRouterAny() *isis.IsisRouterPathAny {
	return &isis.IsisRouterPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"isis-routers", "isis-router"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// IsisRouter (list): Each ISIS router is identified by an arbitrary string
// identifier.
//
//	Defining module:      "open-traffic-generator-isis"
//	Instantiating module: "open-traffic-generator-isis"
//	Path from parent:     "isis-routers/isis-router"
//	Path from root:       "/isis-routers/isis-router"
//
//	Name: string
func (n *RootPath) IsisRouter(Name string) *isis.IsisRouterPath {
	return &isis.IsisRouterPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"isis-routers", "isis-router"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// Lacp (container): LACP telemetry collected by the ATE device.
//
//	Defining module:      "open-traffic-generator-lacp"
//	Instantiating module: "open-traffic-generator-lacp"
//	Path from parent:     "lacp"
//	Path from root:       "/lacp"
func (n *RootPath) Lacp() *lacp.LacpPath {
	return &lacp.LacpPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp"},
			map[string]interface{}{},
			n,
		),
	}
}

// LagAny (list): An individual LAG defined by an OTG.
//
//	Defining module:      "open-traffic-generator-lag"
//	Instantiating module: "open-traffic-generator-lag"
//	Path from parent:     "lags/lag"
//	Path from root:       "/lags/lag"
func (n *RootPath) LagAny() *lag.LagPathAny {
	return &lag.LagPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"lags", "lag"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Lag (list): An individual LAG defined by an OTG.
//
//	Defining module:      "open-traffic-generator-lag"
//	Instantiating module: "open-traffic-generator-lag"
//	Path from parent:     "lags/lag"
//	Path from root:       "/lags/lag"
//
//	Name: string
func (n *RootPath) Lag(Name string) *lag.LagPath {
	return &lag.LagPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"lags", "lag"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// LldpInterfaceAny (list): Each LLDP interface is identified by an arbitrary string
// identifier.
//
//	Defining module:      "open-traffic-generator-lldp"
//	Instantiating module: "open-traffic-generator-lldp"
//	Path from parent:     "lldps/lldp-interface"
//	Path from root:       "/lldps/lldp-interface"
func (n *RootPath) LldpInterfaceAny() *lldp.LldpInterfacePathAny {
	return &lldp.LldpInterfacePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"lldps", "lldp-interface"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// LldpInterface (list): Each LLDP interface is identified by an arbitrary string
// identifier.
//
//	Defining module:      "open-traffic-generator-lldp"
//	Instantiating module: "open-traffic-generator-lldp"
//	Path from parent:     "lldps/lldp-interface"
//	Path from root:       "/lldps/lldp-interface"
//
//	Name: string
func (n *RootPath) LldpInterface(Name string) *lldp.LldpInterfacePath {
	return &lldp.LldpInterfacePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"lldps", "lldp-interface"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// PortAny (list): An individual port defined by an OTG.
//
//	Defining module:      "open-traffic-generator-port"
//	Instantiating module: "open-traffic-generator-port"
//	Path from parent:     "ports/port"
//	Path from root:       "/ports/port"
func (n *RootPath) PortAny() *port.PortPathAny {
	return &port.PortPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"ports", "port"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Port (list): An individual port defined by an OTG.
//
//	Defining module:      "open-traffic-generator-port"
//	Instantiating module: "open-traffic-generator-port"
//	Path from parent:     "ports/port"
//	Path from root:       "/ports/port"
//
//	Name: string
func (n *RootPath) Port(Name string) *port.PortPath {
	return &port.PortPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"ports", "port"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// RsvpteRouterAny (list): Each RSVP-TE router is identified by a unique string
// identifier.
//
//	Defining module:      "open-traffic-generator-rsvp"
//	Instantiating module: "open-traffic-generator-rsvp"
//	Path from parent:     "rsvpte-routers/rsvpte-router"
//	Path from root:       "/rsvpte-routers/rsvpte-router"
func (n *RootPath) RsvpteRouterAny() *rsvp.RsvpteRouterPathAny {
	return &rsvp.RsvpteRouterPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"rsvpte-routers", "rsvpte-router"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// RsvpteRouter (list): Each RSVP-TE router is identified by a unique string
// identifier.
//
//	Defining module:      "open-traffic-generator-rsvp"
//	Instantiating module: "open-traffic-generator-rsvp"
//	Path from parent:     "rsvpte-routers/rsvpte-router"
//	Path from root:       "/rsvpte-routers/rsvpte-router"
//
//	Name: string
func (n *RootPath) RsvpteRouter(Name string) *rsvp.RsvpteRouterPath {
	return &rsvp.RsvpteRouterPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"rsvpte-routers", "rsvpte-router"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// Batch contains a collection of paths.
// Calling State() or Config() on the batch returns a query
// that can use to Lookup, Watch, etc on multiple paths at once.
type Batch struct {
	paths []ygnmi.PathStruct
}

// AddPaths adds the paths to the batch.
func (b *Batch) AddPaths(paths ...ygnmi.PathStruct) *Batch {
	b.paths = append(b.paths, paths...)
	return b
}

// State returns a Query that can be used in gNMI operations.
// The returned query is immutable, adding paths does not modify existing queries.
func (b *Batch) State() ygnmi.SingletonQuery[*oc.Root] {
	queryPaths := make([]ygnmi.PathStruct, len(b.paths))
	copy(queryPaths, b.paths)
	return ygnmi.NewNonLeafSingletonQuery[*oc.Root](
		"Root",
		true,
		ygnmi.NewDeviceRootBase(),
		queryPaths,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
// The returned query is immutable, adding paths does not modify existing queries.
func (b *Batch) Config() ygnmi.SingletonQuery[*oc.Root] {
	queryPaths := make([]ygnmi.PathStruct, len(b.paths))
	copy(queryPaths, b.paths)
	return ygnmi.NewNonLeafSingletonQuery[*oc.Root](
		"Root",
		false,
		ygnmi.NewDeviceRootBase(),
		queryPaths,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

func binarySliceToFloatSlice(in []oc.Binary) []float32 {
	converted := make([]float32, 0, len(in))
	for _, binary := range in {
		converted = append(converted, ygot.BinaryToFloat32(binary))
	}
	return converted
}

// State returns a Query that can be used in gNMI operations.
func (n *RootPath) State() ygnmi.SingletonQuery[*oc.Root] {
	return ygnmi.NewNonLeafSingletonQuery[*oc.Root](
		"Root",
		true,
		n,
		nil,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *RootPath) Config() ygnmi.ConfigQuery[*oc.Root] {
	return ygnmi.NewNonLeafConfigQuery[*oc.Root](
		"Root",
		false,
		n,
		nil,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}
