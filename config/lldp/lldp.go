/*
Package lldp is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema. The generated paths are
based on a compressed form of the schema.

This package was generated by /usr/local/google/home/hines/go/pkg/mod/github.com/openconfig/ygot@v0.16.2/genutil/names.go
using the following YANG input files:
	- gnmi-collector-metadata.yang
	- public/release/models/acl/openconfig-acl.yang
	- public/release/models/acl/openconfig-packet-match.yang
	- public/release/models/aft/openconfig-aft.yang
	- public/release/models/ate/openconfig-ate-flow.yang
	- public/release/models/ate/openconfig-ate-intf.yang
	- public/release/models/bfd/openconfig-bfd.yang
	- public/release/models/bgp/openconfig-bgp-policy.yang
	- public/release/models/bgp/openconfig-bgp-types.yang
	- public/release/models/interfaces/openconfig-if-aggregate.yang
	- public/release/models/interfaces/openconfig-if-ethernet.yang
	- public/release/models/interfaces/openconfig-if-ip-ext.yang
	- public/release/models/interfaces/openconfig-if-ip.yang
	- public/release/models/interfaces/openconfig-interfaces.yang
	- public/release/models/isis/openconfig-isis.yang
	- public/release/models/lacp/openconfig-lacp.yang
	- public/release/models/lldp/openconfig-lldp-types.yang
	- public/release/models/lldp/openconfig-lldp.yang
	- public/release/models/local-routing/openconfig-local-routing.yang
	- public/release/models/mpls/openconfig-mpls-types.yang
	- public/release/models/multicast/openconfig-pim.yang
	- public/release/models/network-instance/openconfig-network-instance.yang
	- public/release/models/openconfig-extensions.yang
	- public/release/models/optical-transport/openconfig-transport-types.yang
	- public/release/models/ospf/openconfig-ospfv2.yang
	- public/release/models/platform/openconfig-platform-cpu.yang
	- public/release/models/platform/openconfig-platform-integrated-circuit.yang
	- public/release/models/platform/openconfig-platform-software.yang
	- public/release/models/platform/openconfig-platform-transceiver.yang
	- public/release/models/platform/openconfig-platform.yang
	- public/release/models/policy-forwarding/openconfig-policy-forwarding.yang
	- public/release/models/policy/openconfig-policy-types.yang
	- public/release/models/qos/openconfig-qos-elements.yang
	- public/release/models/qos/openconfig-qos-interfaces.yang
	- public/release/models/qos/openconfig-qos-types.yang
	- public/release/models/qos/openconfig-qos.yang
	- public/release/models/rib/openconfig-rib-bgp.yang
	- public/release/models/segment-routing/openconfig-segment-routing-types.yang
	- public/release/models/system/openconfig-system.yang
	- public/release/models/types/openconfig-inet-types.yang
	- public/release/models/types/openconfig-types.yang
	- public/release/models/types/openconfig-yang-types.yang
	- public/release/models/vlan/openconfig-vlan.yang
	- public/third_party/ietf/iana-if-type.yang
	- public/third_party/ietf/ietf-inet-types.yang
	- public/third_party/ietf/ietf-interfaces.yang
	- public/third_party/ietf/ietf-yang-types.yang
Imported modules were sourced from:
	- public/release/models/...
	- public/third_party/ietf/...
*/
package lldp

import (
	"github.com/openconfig/ygot/ygot"
)

// LldpPath represents the /openconfig-lldp/lldp YANG schema element.
type LldpPath struct {
	*ygot.NodePath
}

// LldpPathAny represents the wildcard version of the /openconfig-lldp/lldp YANG schema element.
type LldpPathAny struct {
	*ygot.NodePath
}

// Lldp_ChassisIdPath represents the /openconfig-lldp/lldp/config/chassis-id YANG schema element.
type Lldp_ChassisIdPath struct {
	*ygot.NodePath
}

// Lldp_ChassisIdPathAny represents the wildcard version of the /openconfig-lldp/lldp/config/chassis-id YANG schema element.
type Lldp_ChassisIdPathAny struct {
	*ygot.NodePath
}

// Lldp_ChassisIdTypePath represents the /openconfig-lldp/lldp/config/chassis-id-type YANG schema element.
type Lldp_ChassisIdTypePath struct {
	*ygot.NodePath
}

// Lldp_ChassisIdTypePathAny represents the wildcard version of the /openconfig-lldp/lldp/config/chassis-id-type YANG schema element.
type Lldp_ChassisIdTypePathAny struct {
	*ygot.NodePath
}

// Lldp_EnabledPath represents the /openconfig-lldp/lldp/config/enabled YANG schema element.
type Lldp_EnabledPath struct {
	*ygot.NodePath
}

// Lldp_EnabledPathAny represents the wildcard version of the /openconfig-lldp/lldp/config/enabled YANG schema element.
type Lldp_EnabledPathAny struct {
	*ygot.NodePath
}

// Lldp_HelloTimerPath represents the /openconfig-lldp/lldp/config/hello-timer YANG schema element.
type Lldp_HelloTimerPath struct {
	*ygot.NodePath
}

// Lldp_HelloTimerPathAny represents the wildcard version of the /openconfig-lldp/lldp/config/hello-timer YANG schema element.
type Lldp_HelloTimerPathAny struct {
	*ygot.NodePath
}

// Lldp_SuppressTlvAdvertisementPath represents the /openconfig-lldp/lldp/config/suppress-tlv-advertisement YANG schema element.
type Lldp_SuppressTlvAdvertisementPath struct {
	*ygot.NodePath
}

// Lldp_SuppressTlvAdvertisementPathAny represents the wildcard version of the /openconfig-lldp/lldp/config/suppress-tlv-advertisement YANG schema element.
type Lldp_SuppressTlvAdvertisementPathAny struct {
	*ygot.NodePath
}

// Lldp_SystemDescriptionPath represents the /openconfig-lldp/lldp/config/system-description YANG schema element.
type Lldp_SystemDescriptionPath struct {
	*ygot.NodePath
}

// Lldp_SystemDescriptionPathAny represents the wildcard version of the /openconfig-lldp/lldp/config/system-description YANG schema element.
type Lldp_SystemDescriptionPathAny struct {
	*ygot.NodePath
}

// Lldp_SystemNamePath represents the /openconfig-lldp/lldp/config/system-name YANG schema element.
type Lldp_SystemNamePath struct {
	*ygot.NodePath
}

// Lldp_SystemNamePathAny represents the wildcard version of the /openconfig-lldp/lldp/config/system-name YANG schema element.
type Lldp_SystemNamePathAny struct {
	*ygot.NodePath
}

// ChassisId (leaf): The Chassis ID is a mandatory TLV which identifies the
// chassis component of the endpoint identifier associated with
// the transmitting LLDP agent
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/chassis-id"
// Path from root: "/lldp/config/chassis-id"
func (n *LldpPath) ChassisId() *Lldp_ChassisIdPath {
	return &Lldp_ChassisIdPath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "chassis-id"},
			map[string]interface{}{},
			n,
		),
	}
}

// ChassisId (leaf): The Chassis ID is a mandatory TLV which identifies the
// chassis component of the endpoint identifier associated with
// the transmitting LLDP agent
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/chassis-id"
// Path from root: "/lldp/config/chassis-id"
func (n *LldpPathAny) ChassisId() *Lldp_ChassisIdPathAny {
	return &Lldp_ChassisIdPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"config", "chassis-id"},
			map[string]interface{}{},
			n,
		),
	}
}

// ChassisIdType (leaf): This field identifies the format and source of the chassis
// identifier string. It is an enumerator defined by the
// LldpChassisIdSubtype object from IEEE 802.1AB MIB.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/chassis-id-type"
// Path from root: "/lldp/config/chassis-id-type"
func (n *LldpPath) ChassisIdType() *Lldp_ChassisIdTypePath {
	return &Lldp_ChassisIdTypePath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "chassis-id-type"},
			map[string]interface{}{},
			n,
		),
	}
}

// ChassisIdType (leaf): This field identifies the format and source of the chassis
// identifier string. It is an enumerator defined by the
// LldpChassisIdSubtype object from IEEE 802.1AB MIB.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/chassis-id-type"
// Path from root: "/lldp/config/chassis-id-type"
func (n *LldpPathAny) ChassisIdType() *Lldp_ChassisIdTypePathAny {
	return &Lldp_ChassisIdTypePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"config", "chassis-id-type"},
			map[string]interface{}{},
			n,
		),
	}
}

// Enabled (leaf): System level state of the LLDP protocol.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/enabled"
// Path from root: "/lldp/config/enabled"
func (n *LldpPath) Enabled() *Lldp_EnabledPath {
	return &Lldp_EnabledPath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "enabled"},
			map[string]interface{}{},
			n,
		),
	}
}

// Enabled (leaf): System level state of the LLDP protocol.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/enabled"
// Path from root: "/lldp/config/enabled"
func (n *LldpPathAny) Enabled() *Lldp_EnabledPathAny {
	return &Lldp_EnabledPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"config", "enabled"},
			map[string]interface{}{},
			n,
		),
	}
}

// HelloTimer (leaf): System level hello timer for the LLDP protocol.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/hello-timer"
// Path from root: "/lldp/config/hello-timer"
func (n *LldpPath) HelloTimer() *Lldp_HelloTimerPath {
	return &Lldp_HelloTimerPath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "hello-timer"},
			map[string]interface{}{},
			n,
		),
	}
}

// HelloTimer (leaf): System level hello timer for the LLDP protocol.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/hello-timer"
// Path from root: "/lldp/config/hello-timer"
func (n *LldpPathAny) HelloTimer() *Lldp_HelloTimerPathAny {
	return &Lldp_HelloTimerPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"config", "hello-timer"},
			map[string]interface{}{},
			n,
		),
	}
}

// InterfaceAny (list): List of interfaces on which LLDP is enabled / available
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "interfaces/interface"
// Path from root: "/lldp/interfaces/interface"
// Name (wildcarded): string
func (n *LldpPath) InterfaceAny() *Lldp_InterfacePathAny {
	return &Lldp_InterfacePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// InterfaceAny (list): List of interfaces on which LLDP is enabled / available
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "interfaces/interface"
// Path from root: "/lldp/interfaces/interface"
// Name (wildcarded): string
func (n *LldpPathAny) InterfaceAny() *Lldp_InterfacePathAny {
	return &Lldp_InterfacePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Interface (list): List of interfaces on which LLDP is enabled / available
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "interfaces/interface"
// Path from root: "/lldp/interfaces/interface"
// Name: string
func (n *LldpPath) Interface(Name string) *Lldp_InterfacePath {
	return &Lldp_InterfacePath{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// Interface (list): List of interfaces on which LLDP is enabled / available
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "interfaces/interface"
// Path from root: "/lldp/interfaces/interface"
// Name: string
func (n *LldpPathAny) Interface(Name string) *Lldp_InterfacePathAny {
	return &Lldp_InterfacePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// SuppressTlvAdvertisement (leaf): Indicates whether the local system should suppress the
// advertisement of particular TLVs with the LLDP PDUs that it
// transmits. Where a TLV type is specified within this list, it
// should not be included in any LLDP PDU transmitted by the
// local agent.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/suppress-tlv-advertisement"
// Path from root: "/lldp/config/suppress-tlv-advertisement"
func (n *LldpPath) SuppressTlvAdvertisement() *Lldp_SuppressTlvAdvertisementPath {
	return &Lldp_SuppressTlvAdvertisementPath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "suppress-tlv-advertisement"},
			map[string]interface{}{},
			n,
		),
	}
}

// SuppressTlvAdvertisement (leaf): Indicates whether the local system should suppress the
// advertisement of particular TLVs with the LLDP PDUs that it
// transmits. Where a TLV type is specified within this list, it
// should not be included in any LLDP PDU transmitted by the
// local agent.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/suppress-tlv-advertisement"
// Path from root: "/lldp/config/suppress-tlv-advertisement"
func (n *LldpPathAny) SuppressTlvAdvertisement() *Lldp_SuppressTlvAdvertisementPathAny {
	return &Lldp_SuppressTlvAdvertisementPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"config", "suppress-tlv-advertisement"},
			map[string]interface{}{},
			n,
		),
	}
}

// SystemDescription (leaf): The system description field shall contain an alpha-numeric
// string that is the textual description of the network entity.
// The system description should include the full name and
// version identification of the system's hardware type,
// software operating system, and networking software. If
// implementations support IETF RFC 3418, the sysDescr object
// should be used for this field.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/system-description"
// Path from root: "/lldp/config/system-description"
func (n *LldpPath) SystemDescription() *Lldp_SystemDescriptionPath {
	return &Lldp_SystemDescriptionPath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "system-description"},
			map[string]interface{}{},
			n,
		),
	}
}

// SystemDescription (leaf): The system description field shall contain an alpha-numeric
// string that is the textual description of the network entity.
// The system description should include the full name and
// version identification of the system's hardware type,
// software operating system, and networking software. If
// implementations support IETF RFC 3418, the sysDescr object
// should be used for this field.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/system-description"
// Path from root: "/lldp/config/system-description"
func (n *LldpPathAny) SystemDescription() *Lldp_SystemDescriptionPathAny {
	return &Lldp_SystemDescriptionPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"config", "system-description"},
			map[string]interface{}{},
			n,
		),
	}
}

// SystemName (leaf): The system name field shall contain an alpha-numeric string
// that indicates the system's administratively assigned name.
// The system name should be the system's fully qualified domain
// name. If implementations support IETF RFC 3418, the sysName
// object should be used for this field.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/system-name"
// Path from root: "/lldp/config/system-name"
func (n *LldpPath) SystemName() *Lldp_SystemNamePath {
	return &Lldp_SystemNamePath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "system-name"},
			map[string]interface{}{},
			n,
		),
	}
}

// SystemName (leaf): The system name field shall contain an alpha-numeric string
// that indicates the system's administratively assigned name.
// The system name should be the system's fully qualified domain
// name. If implementations support IETF RFC 3418, the sysName
// object should be used for this field.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/system-name"
// Path from root: "/lldp/config/system-name"
func (n *LldpPathAny) SystemName() *Lldp_SystemNamePathAny {
	return &Lldp_SystemNamePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"config", "system-name"},
			map[string]interface{}{},
			n,
		),
	}
}

// Lldp_InterfacePath represents the /openconfig-lldp/lldp/interfaces/interface YANG schema element.
type Lldp_InterfacePath struct {
	*ygot.NodePath
}

// Lldp_InterfacePathAny represents the wildcard version of the /openconfig-lldp/lldp/interfaces/interface YANG schema element.
type Lldp_InterfacePathAny struct {
	*ygot.NodePath
}

// Lldp_Interface_EnabledPath represents the /openconfig-lldp/lldp/interfaces/interface/config/enabled YANG schema element.
type Lldp_Interface_EnabledPath struct {
	*ygot.NodePath
}

// Lldp_Interface_EnabledPathAny represents the wildcard version of the /openconfig-lldp/lldp/interfaces/interface/config/enabled YANG schema element.
type Lldp_Interface_EnabledPathAny struct {
	*ygot.NodePath
}

// Lldp_Interface_NamePath represents the /openconfig-lldp/lldp/interfaces/interface/config/name YANG schema element.
type Lldp_Interface_NamePath struct {
	*ygot.NodePath
}

// Lldp_Interface_NamePathAny represents the wildcard version of the /openconfig-lldp/lldp/interfaces/interface/config/name YANG schema element.
type Lldp_Interface_NamePathAny struct {
	*ygot.NodePath
}

// Enabled (leaf): Enable or disable the LLDP protocol on the interface.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/enabled"
// Path from root: "/lldp/interfaces/interface/config/enabled"
func (n *Lldp_InterfacePath) Enabled() *Lldp_Interface_EnabledPath {
	return &Lldp_Interface_EnabledPath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "enabled"},
			map[string]interface{}{},
			n,
		),
	}
}

// Enabled (leaf): Enable or disable the LLDP protocol on the interface.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/enabled"
// Path from root: "/lldp/interfaces/interface/config/enabled"
func (n *Lldp_InterfacePathAny) Enabled() *Lldp_Interface_EnabledPathAny {
	return &Lldp_Interface_EnabledPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"config", "enabled"},
			map[string]interface{}{},
			n,
		),
	}
}

// Name (leaf): Reference to the LLDP Ethernet interface
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/name"
// Path from root: "/lldp/interfaces/interface/config/name"
func (n *Lldp_InterfacePath) Name() *Lldp_Interface_NamePath {
	return &Lldp_Interface_NamePath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "name"},
			map[string]interface{}{},
			n,
		),
	}
}

// Name (leaf): Reference to the LLDP Ethernet interface
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "config/name"
// Path from root: "/lldp/interfaces/interface/config/name"
func (n *Lldp_InterfacePathAny) Name() *Lldp_Interface_NamePathAny {
	return &Lldp_Interface_NamePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"config", "name"},
			map[string]interface{}{},
			n,
		),
	}
}
