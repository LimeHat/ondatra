//go:build go1.18

/*
Package discovery is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema. The generated paths are
based on a compressed form of the schema.

This package was generated by ygnmi version: v0.2.1: (ygot: v0.23.1)
using the following YANG input files:
  - models-yang/models/isis/open-traffic-generator-isis.yang
  - models-yang/models/types/open-traffic-generator-types.yang
  - models-yang/models/flow/open-traffic-generator-flow.yang
  - models-yang/models/discovery/open-traffic-generator-discovery.yang
  - models-yang/models/interface/open-traffic-generator-port.yang
  - models-yang/models/bgp/open-traffic-generator-bgp.yang
  - models-yang/models/lag/open-traffic-generator-lag.yang
  - models-yang/models/lacp/open-traffic-generator-lacp.yang

Imported modules were sourced from:
  - models-yang/models/...
*/
package discovery

import (
	oc "github.com/openconfig/ondatra/gnmi/otg"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// Interface_NamePath represents the /open-traffic-generator-discovery/interfaces/interface/state/name YANG schema element.
type Interface_NamePath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Interface_NamePathAny represents the wildcard version of the /open-traffic-generator-discovery/interfaces/interface/state/name YANG schema element.
type Interface_NamePathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
func (n *InterfacePath) State() ygnmi.SingletonQuery[*oc.Interface] {
	return ygnmi.NewNonLeafSingletonQuery[*oc.Interface](
		"Interface",
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

// State returns a Query that can be used in gNMI operations.
func (n *InterfacePathAny) State() ygnmi.WildcardQuery[*oc.Interface] {
	return ygnmi.NewNonLeafWildcardQuery[*oc.Interface](
		"Interface",
		true,
		n,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "state/name"
//	Path from root:       "/interfaces/interface/state/name"
func (n *Interface_NamePath) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewLeafSingletonQuery[string](
		"Interface",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "name"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface).Name
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "state/name"
//	Path from root:       "/interfaces/interface/state/name"
func (n *Interface_NamePathAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewLeafWildcardQuery[string](
		"Interface",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "name"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface).Name
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "name"
//	Path from root:       ""
func (n *Interface_NamePath) Config() ygnmi.ConfigQuery[string] {
	return ygnmi.NewLeafConfigQuery[string](
		"Interface",
		false,
		true,
		ygnmi.NewNodePath(
			[]string{"name"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface).Name
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "name"
//	Path from root:       ""
func (n *Interface_NamePathAny) Config() ygnmi.WildcardQuery[string] {
	return ygnmi.NewLeafWildcardQuery[string](
		"Interface",
		false,
		true,
		ygnmi.NewNodePath(
			[]string{"name"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface).Name
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// InterfacePath represents the /open-traffic-generator-discovery/interfaces/interface YANG schema element.
type InterfacePath struct {
	*ygnmi.NodePath
}

// InterfacePathAny represents the wildcard version of the /open-traffic-generator-discovery/interfaces/interface YANG schema element.
type InterfacePathAny struct {
	*ygnmi.NodePath
}

// Ipv4NeighborAny (list): The interface neighbor state or ARP cache entry.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "ipv4-neighbors/ipv4-neighbor"
//	Path from root:       "/interfaces/interface/ipv4-neighbors/ipv4-neighbor"
func (n *InterfacePath) Ipv4NeighborAny() *Interface_Ipv4NeighborPathAny {
	return &Interface_Ipv4NeighborPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"ipv4-neighbors", "ipv4-neighbor"},
			map[string]interface{}{"ipv4-address": "*"},
			n,
		),
	}
}

// Ipv4NeighborAny (list): The interface neighbor state or ARP cache entry.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "ipv4-neighbors/ipv4-neighbor"
//	Path from root:       "/interfaces/interface/ipv4-neighbors/ipv4-neighbor"
func (n *InterfacePathAny) Ipv4NeighborAny() *Interface_Ipv4NeighborPathAny {
	return &Interface_Ipv4NeighborPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"ipv4-neighbors", "ipv4-neighbor"},
			map[string]interface{}{"ipv4-address": "*"},
			n,
		),
	}
}

// Ipv4Neighbor (list): The interface neighbor state or ARP cache entry.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "ipv4-neighbors/ipv4-neighbor"
//	Path from root:       "/interfaces/interface/ipv4-neighbors/ipv4-neighbor"
//
//	Ipv4Address: string
func (n *InterfacePath) Ipv4Neighbor(Ipv4Address string) *Interface_Ipv4NeighborPath {
	return &Interface_Ipv4NeighborPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"ipv4-neighbors", "ipv4-neighbor"},
			map[string]interface{}{"ipv4-address": Ipv4Address},
			n,
		),
	}
}

// Ipv4Neighbor (list): The interface neighbor state or ARP cache entry.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "ipv4-neighbors/ipv4-neighbor"
//	Path from root:       "/interfaces/interface/ipv4-neighbors/ipv4-neighbor"
//
//	Ipv4Address: string
func (n *InterfacePathAny) Ipv4Neighbor(Ipv4Address string) *Interface_Ipv4NeighborPathAny {
	return &Interface_Ipv4NeighborPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"ipv4-neighbors", "ipv4-neighbor"},
			map[string]interface{}{"ipv4-address": Ipv4Address},
			n,
		),
	}
}

// Ipv6NeighborAny (list): The interface neighbor state or NDISC cache entry.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "ipv6-neighbors/ipv6-neighbor"
//	Path from root:       "/interfaces/interface/ipv6-neighbors/ipv6-neighbor"
func (n *InterfacePath) Ipv6NeighborAny() *Interface_Ipv6NeighborPathAny {
	return &Interface_Ipv6NeighborPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"ipv6-neighbors", "ipv6-neighbor"},
			map[string]interface{}{"ipv6-address": "*"},
			n,
		),
	}
}

// Ipv6NeighborAny (list): The interface neighbor state or NDISC cache entry.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "ipv6-neighbors/ipv6-neighbor"
//	Path from root:       "/interfaces/interface/ipv6-neighbors/ipv6-neighbor"
func (n *InterfacePathAny) Ipv6NeighborAny() *Interface_Ipv6NeighborPathAny {
	return &Interface_Ipv6NeighborPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"ipv6-neighbors", "ipv6-neighbor"},
			map[string]interface{}{"ipv6-address": "*"},
			n,
		),
	}
}

// Ipv6Neighbor (list): The interface neighbor state or NDISC cache entry.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "ipv6-neighbors/ipv6-neighbor"
//	Path from root:       "/interfaces/interface/ipv6-neighbors/ipv6-neighbor"
//
//	Ipv6Address: string
func (n *InterfacePath) Ipv6Neighbor(Ipv6Address string) *Interface_Ipv6NeighborPath {
	return &Interface_Ipv6NeighborPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"ipv6-neighbors", "ipv6-neighbor"},
			map[string]interface{}{"ipv6-address": Ipv6Address},
			n,
		),
	}
}

// Ipv6Neighbor (list): The interface neighbor state or NDISC cache entry.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "ipv6-neighbors/ipv6-neighbor"
//	Path from root:       "/interfaces/interface/ipv6-neighbors/ipv6-neighbor"
//
//	Ipv6Address: string
func (n *InterfacePathAny) Ipv6Neighbor(Ipv6Address string) *Interface_Ipv6NeighborPathAny {
	return &Interface_Ipv6NeighborPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"ipv6-neighbors", "ipv6-neighbor"},
			map[string]interface{}{"ipv6-address": Ipv6Address},
			n,
		),
	}
}

// Name (leaf): An arbitary name of an OTG interface determined by the OTG
// configuration.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "*/name"
//	Path from root:       "/interfaces/interface/*/name"
func (n *InterfacePath) Name() *Interface_NamePath {
	return &Interface_NamePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "name"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Name (leaf): An arbitary name of an OTG interface determined by the OTG
// configuration.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "*/name"
//	Path from root:       "/interfaces/interface/*/name"
func (n *InterfacePathAny) Name() *Interface_NamePathAny {
	return &Interface_NamePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "name"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Interface_Ipv4Neighbor_Ipv4AddressPath represents the /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address YANG schema element.
type Interface_Ipv4Neighbor_Ipv4AddressPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Interface_Ipv4Neighbor_Ipv4AddressPathAny represents the wildcard version of the /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address YANG schema element.
type Interface_Ipv4Neighbor_Ipv4AddressPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
func (n *Interface_Ipv4NeighborPath) State() ygnmi.SingletonQuery[*oc.Interface_Ipv4Neighbor] {
	return ygnmi.NewNonLeafSingletonQuery[*oc.Interface_Ipv4Neighbor](
		"Interface_Ipv4Neighbor",
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

// State returns a Query that can be used in gNMI operations.
func (n *Interface_Ipv4NeighborPathAny) State() ygnmi.WildcardQuery[*oc.Interface_Ipv4Neighbor] {
	return ygnmi.NewNonLeafWildcardQuery[*oc.Interface_Ipv4Neighbor](
		"Interface_Ipv4Neighbor",
		true,
		n,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "state/ipv4-address"
//	Path from root:       "/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address"
func (n *Interface_Ipv4Neighbor_Ipv4AddressPath) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewLeafSingletonQuery[string](
		"Interface_Ipv4Neighbor",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "ipv4-address"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface_Ipv4Neighbor).Ipv4Address
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface_Ipv4Neighbor) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "state/ipv4-address"
//	Path from root:       "/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address"
func (n *Interface_Ipv4Neighbor_Ipv4AddressPathAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewLeafWildcardQuery[string](
		"Interface_Ipv4Neighbor",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "ipv4-address"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface_Ipv4Neighbor).Ipv4Address
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface_Ipv4Neighbor) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "ipv4-address"
//	Path from root:       ""
func (n *Interface_Ipv4Neighbor_Ipv4AddressPath) Config() ygnmi.ConfigQuery[string] {
	return ygnmi.NewLeafConfigQuery[string](
		"Interface_Ipv4Neighbor",
		false,
		true,
		ygnmi.NewNodePath(
			[]string{"ipv4-address"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface_Ipv4Neighbor).Ipv4Address
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface_Ipv4Neighbor) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "ipv4-address"
//	Path from root:       ""
func (n *Interface_Ipv4Neighbor_Ipv4AddressPathAny) Config() ygnmi.WildcardQuery[string] {
	return ygnmi.NewLeafWildcardQuery[string](
		"Interface_Ipv4Neighbor",
		false,
		true,
		ygnmi.NewNodePath(
			[]string{"ipv4-address"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface_Ipv4Neighbor).Ipv4Address
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface_Ipv4Neighbor) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "state/link-layer-address"
//	Path from root:       "/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address"
func (n *Interface_Ipv4Neighbor_LinkLayerAddressPath) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewLeafSingletonQuery[string](
		"Interface_Ipv4Neighbor",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "link-layer-address"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface_Ipv4Neighbor).LinkLayerAddress
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface_Ipv4Neighbor) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "state/link-layer-address"
//	Path from root:       "/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address"
func (n *Interface_Ipv4Neighbor_LinkLayerAddressPathAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewLeafWildcardQuery[string](
		"Interface_Ipv4Neighbor",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "link-layer-address"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface_Ipv4Neighbor).LinkLayerAddress
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface_Ipv4Neighbor) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Interface_Ipv4Neighbor_LinkLayerAddressPath represents the /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address YANG schema element.
type Interface_Ipv4Neighbor_LinkLayerAddressPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Interface_Ipv4Neighbor_LinkLayerAddressPathAny represents the wildcard version of the /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address YANG schema element.
type Interface_Ipv4Neighbor_LinkLayerAddressPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Interface_Ipv4NeighborPath represents the /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor YANG schema element.
type Interface_Ipv4NeighborPath struct {
	*ygnmi.NodePath
}

// Interface_Ipv4NeighborPathAny represents the wildcard version of the /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor YANG schema element.
type Interface_Ipv4NeighborPathAny struct {
	*ygnmi.NodePath
}

// Ipv4Address (leaf): The IPv4 address of the neighbor.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "*/ipv4-address"
//	Path from root:       "/interfaces/interface/ipv4-neighbors/ipv4-neighbor/*/ipv4-address"
func (n *Interface_Ipv4NeighborPath) Ipv4Address() *Interface_Ipv4Neighbor_Ipv4AddressPath {
	return &Interface_Ipv4Neighbor_Ipv4AddressPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "ipv4-address"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Ipv4Address (leaf): The IPv4 address of the neighbor.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "*/ipv4-address"
//	Path from root:       "/interfaces/interface/ipv4-neighbors/ipv4-neighbor/*/ipv4-address"
func (n *Interface_Ipv4NeighborPathAny) Ipv4Address() *Interface_Ipv4Neighbor_Ipv4AddressPathAny {
	return &Interface_Ipv4Neighbor_Ipv4AddressPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "ipv4-address"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// LinkLayerAddress (leaf): The link layer address or MAC address of the neighbor.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "state/link-layer-address"
//	Path from root:       "/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address"
func (n *Interface_Ipv4NeighborPath) LinkLayerAddress() *Interface_Ipv4Neighbor_LinkLayerAddressPath {
	return &Interface_Ipv4Neighbor_LinkLayerAddressPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"state", "link-layer-address"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// LinkLayerAddress (leaf): The link layer address or MAC address of the neighbor.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "state/link-layer-address"
//	Path from root:       "/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/link-layer-address"
func (n *Interface_Ipv4NeighborPathAny) LinkLayerAddress() *Interface_Ipv4Neighbor_LinkLayerAddressPathAny {
	return &Interface_Ipv4Neighbor_LinkLayerAddressPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"state", "link-layer-address"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Interface_Ipv6Neighbor_Ipv6AddressPath represents the /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address YANG schema element.
type Interface_Ipv6Neighbor_Ipv6AddressPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Interface_Ipv6Neighbor_Ipv6AddressPathAny represents the wildcard version of the /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address YANG schema element.
type Interface_Ipv6Neighbor_Ipv6AddressPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
func (n *Interface_Ipv6NeighborPath) State() ygnmi.SingletonQuery[*oc.Interface_Ipv6Neighbor] {
	return ygnmi.NewNonLeafSingletonQuery[*oc.Interface_Ipv6Neighbor](
		"Interface_Ipv6Neighbor",
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

// State returns a Query that can be used in gNMI operations.
func (n *Interface_Ipv6NeighborPathAny) State() ygnmi.WildcardQuery[*oc.Interface_Ipv6Neighbor] {
	return ygnmi.NewNonLeafWildcardQuery[*oc.Interface_Ipv6Neighbor](
		"Interface_Ipv6Neighbor",
		true,
		n,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "state/ipv6-address"
//	Path from root:       "/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address"
func (n *Interface_Ipv6Neighbor_Ipv6AddressPath) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewLeafSingletonQuery[string](
		"Interface_Ipv6Neighbor",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "ipv6-address"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface_Ipv6Neighbor).Ipv6Address
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface_Ipv6Neighbor) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "state/ipv6-address"
//	Path from root:       "/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/ipv6-address"
func (n *Interface_Ipv6Neighbor_Ipv6AddressPathAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewLeafWildcardQuery[string](
		"Interface_Ipv6Neighbor",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "ipv6-address"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface_Ipv6Neighbor).Ipv6Address
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface_Ipv6Neighbor) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "ipv6-address"
//	Path from root:       ""
func (n *Interface_Ipv6Neighbor_Ipv6AddressPath) Config() ygnmi.ConfigQuery[string] {
	return ygnmi.NewLeafConfigQuery[string](
		"Interface_Ipv6Neighbor",
		false,
		true,
		ygnmi.NewNodePath(
			[]string{"ipv6-address"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface_Ipv6Neighbor).Ipv6Address
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface_Ipv6Neighbor) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "ipv6-address"
//	Path from root:       ""
func (n *Interface_Ipv6Neighbor_Ipv6AddressPathAny) Config() ygnmi.WildcardQuery[string] {
	return ygnmi.NewLeafWildcardQuery[string](
		"Interface_Ipv6Neighbor",
		false,
		true,
		ygnmi.NewNodePath(
			[]string{"ipv6-address"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface_Ipv6Neighbor).Ipv6Address
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface_Ipv6Neighbor) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "state/link-layer-address"
//	Path from root:       "/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/link-layer-address"
func (n *Interface_Ipv6Neighbor_LinkLayerAddressPath) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewLeafSingletonQuery[string](
		"Interface_Ipv6Neighbor",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "link-layer-address"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface_Ipv6Neighbor).LinkLayerAddress
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface_Ipv6Neighbor) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "state/link-layer-address"
//	Path from root:       "/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/link-layer-address"
func (n *Interface_Ipv6Neighbor_LinkLayerAddressPathAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewLeafWildcardQuery[string](
		"Interface_Ipv6Neighbor",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "link-layer-address"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Interface_Ipv6Neighbor).LinkLayerAddress
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Interface_Ipv6Neighbor) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Interface_Ipv6Neighbor_LinkLayerAddressPath represents the /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/link-layer-address YANG schema element.
type Interface_Ipv6Neighbor_LinkLayerAddressPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Interface_Ipv6Neighbor_LinkLayerAddressPathAny represents the wildcard version of the /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/link-layer-address YANG schema element.
type Interface_Ipv6Neighbor_LinkLayerAddressPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Interface_Ipv6NeighborPath represents the /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor YANG schema element.
type Interface_Ipv6NeighborPath struct {
	*ygnmi.NodePath
}

// Interface_Ipv6NeighborPathAny represents the wildcard version of the /open-traffic-generator-discovery/interfaces/interface/ipv6-neighbors/ipv6-neighbor YANG schema element.
type Interface_Ipv6NeighborPathAny struct {
	*ygnmi.NodePath
}

// Ipv6Address (leaf): The IPv6 address of the neighbor.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "*/ipv6-address"
//	Path from root:       "/interfaces/interface/ipv6-neighbors/ipv6-neighbor/*/ipv6-address"
func (n *Interface_Ipv6NeighborPath) Ipv6Address() *Interface_Ipv6Neighbor_Ipv6AddressPath {
	return &Interface_Ipv6Neighbor_Ipv6AddressPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "ipv6-address"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Ipv6Address (leaf): The IPv6 address of the neighbor.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "*/ipv6-address"
//	Path from root:       "/interfaces/interface/ipv6-neighbors/ipv6-neighbor/*/ipv6-address"
func (n *Interface_Ipv6NeighborPathAny) Ipv6Address() *Interface_Ipv6Neighbor_Ipv6AddressPathAny {
	return &Interface_Ipv6Neighbor_Ipv6AddressPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "ipv6-address"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// LinkLayerAddress (leaf): The link layer address or MAC address of the neighbor.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "state/link-layer-address"
//	Path from root:       "/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/link-layer-address"
func (n *Interface_Ipv6NeighborPath) LinkLayerAddress() *Interface_Ipv6Neighbor_LinkLayerAddressPath {
	return &Interface_Ipv6Neighbor_LinkLayerAddressPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"state", "link-layer-address"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// LinkLayerAddress (leaf): The link layer address or MAC address of the neighbor.
//
//	Defining module:      "open-traffic-generator-discovery-interfaces"
//	Instantiating module: "open-traffic-generator-discovery"
//	Path from parent:     "state/link-layer-address"
//	Path from root:       "/interfaces/interface/ipv6-neighbors/ipv6-neighbor/state/link-layer-address"
func (n *Interface_Ipv6NeighborPathAny) LinkLayerAddress() *Interface_Ipv6Neighbor_LinkLayerAddressPathAny {
	return &Interface_Ipv6Neighbor_LinkLayerAddressPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"state", "link-layer-address"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}
