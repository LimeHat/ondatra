/*
Package otg is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was true
in this case).

This package was generated by ygnmi version: v0.8.1: (ygot: v0.29.0)
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
package otg

import (
	"reflect"

	"github.com/openconfig/ygot/ygot"
)

// ΛEnum is a map, keyed by the name of the type defined for each enum in the
// generated Go code, which provides a mapping between the constant int64 value
// of each value of the enumeration, and the string that is used to represent it
// in the YANG schema. The map is named ΛEnum in order to avoid clash with any
// valid YANG identifier.
var ΛEnum = map[string]map[int64]ygot.EnumDefinition{
	"E_BgpPeer_SessionState": {
		1: {Name: "IDLE"},
		2: {Name: "CONNECT"},
		3: {Name: "ACTIVE"},
		4: {Name: "OPEN_SENT"},
		5: {Name: "OPEN_CONFIRM"},
		6: {Name: "ESTABLISHED"},
	},
	"E_Capabilities_Name": {
		1:  {Name: "MAC_BRIDGE"},
		2:  {Name: "TWO_PORT_MAC_RELAY"},
		3:  {Name: "REPEATER"},
		4:  {Name: "DOCSIS_CABLE_DEVICE"},
		5:  {Name: "S_VLAN"},
		6:  {Name: "TELEPHONE"},
		7:  {Name: "OTHER"},
		8:  {Name: "ROUTER"},
		9:  {Name: "C_VLAN"},
		10: {Name: "STATION_ONLY"},
		11: {Name: "WLAN_ACCESS_POINT"},
	},
	"E_Ero_Type": {
		1: {Name: "IPV4"},
		2: {Name: "IPV6"},
		3: {Name: "ASN"},
		4: {Name: "ASN4"},
		5: {Name: "LABEL"},
		6: {Name: "UNNUMBERED_INTERFACE"},
	},
	"E_ExtendedIpv4Reachability_Prefix_RedistributionType": {
		1: {Name: "UP"},
		2: {Name: "DOWN"},
	},
	"E_Ipv4ExternalReachability_Prefix_OriginType": {
		1: {Name: "INTERNAL"},
		2: {Name: "EXTERNAL"},
	},
	"E_Ipv4ExternalReachability_Prefix_RedistributionType": {
		1: {Name: "UP"},
		2: {Name: "DOWN"},
	},
	"E_Ipv4Lsp_LabelIn": {
		1: {Name: "IPV4_EXPLICIT_NULL"},
		2: {Name: "ROUTER_ALERT"},
		3: {Name: "IPV6_EXPLICIT_NULL"},
		4: {Name: "IMPLICIT_NULL"},
		8: {Name: "ENTROPY_LABEL_INDICATOR"},
		9: {Name: "NO_LABEL"},
	},
	"E_Ipv4Lsp_LabelOut": {
		1: {Name: "IPV4_EXPLICIT_NULL"},
		2: {Name: "ROUTER_ALERT"},
		3: {Name: "IPV6_EXPLICIT_NULL"},
		4: {Name: "IMPLICIT_NULL"},
		8: {Name: "ENTROPY_LABEL_INDICATOR"},
		9: {Name: "NO_LABEL"},
	},
	"E_Ipv4Lsp_LastFlapReason": {
		1: {Name: "RESVTEAR"},
		2: {Name: "PATHTEAR"},
		3: {Name: "PATHTIMEOUT"},
	},
	"E_Ipv4Lsp_SessionStatus": {
		1: {Name: "UP"},
		2: {Name: "DOWN"},
	},
	"E_Ipv6Reachability_Prefix_OriginType": {
		1: {Name: "INTERNAL"},
		2: {Name: "EXTERNAL"},
	},
	"E_Ipv6Reachability_Prefix_RedistributionType": {
		1: {Name: "UP"},
		2: {Name: "DOWN"},
	},
	"E_Lacp_LacpActivityType": {
		1: {Name: "ACTIVE"},
		2: {Name: "PASSIVE"},
	},
	"E_Lacp_LacpSynchronizationType": {
		1: {Name: "IN_SYNC"},
		2: {Name: "OUT_SYNC"},
	},
	"E_Lacp_LacpTimeoutType": {
		1: {Name: "LONG"},
		2: {Name: "SHORT"},
	},
	"E_Lag_OperStatus": {
		1: {Name: "UP"},
		2: {Name: "DOWN"},
	},
	"E_LldpNeighbor_ChassisIdType": {
		1: {Name: "PORT_COMPONENT"},
		2: {Name: "NETWORK_ADDRESS"},
		3: {Name: "CHASSIS_COMPONENT"},
		4: {Name: "MAC_ADDRESS"},
		5: {Name: "INTERFACE_NAME"},
		6: {Name: "LOCAL"},
		7: {Name: "INTERFACE_ALIAS"},
	},
	"E_LldpNeighbor_PortIdType": {
		1: {Name: "PORT_COMPONENT"},
		2: {Name: "NETWORK_ADDRESS"},
		3: {Name: "AGENT_CIRCUIT_ID"},
		4: {Name: "MAC_ADDRESS"},
		5: {Name: "INTERFACE_NAME"},
		6: {Name: "LOCAL"},
		7: {Name: "INTERFACE_ALIAS"},
	},
	"E_Lsps_Flags": {
		1: {Name: "PARTITION_REPAIR"},
		2: {Name: "ATTACHED_ERROR"},
		3: {Name: "ATTACHED_EXPENSE"},
		4: {Name: "ATTACHED_DELAY"},
		5: {Name: "ATTACHED_DEFAULT"},
		6: {Name: "OVERLOAD"},
	},
	"E_Lsps_PduType": {
		1: {Name: "LEVEL_1"},
		2: {Name: "LEVEL_2"},
	},
	"E_Port_Link": {
		1: {Name: "UP"},
		2: {Name: "DOWN"},
	},
	"E_Rro_ReportedLabel": {
		1: {Name: "IPV4_EXPLICIT_NULL"},
		2: {Name: "ROUTER_ALERT"},
		3: {Name: "IPV6_EXPLICIT_NULL"},
		4: {Name: "IMPLICIT_NULL"},
		8: {Name: "ENTROPY_LABEL_INDICATOR"},
		9: {Name: "NO_LABEL"},
	},
	"E_State_CommunityType": {
		1: {Name: "MANUAL_AS_NUMBER"},
		2: {Name: "NO_EXPORT"},
		3: {Name: "NO_ADVERTISE"},
		4: {Name: "NO_EXPORT_SUBCONFED"},
		5: {Name: "LLGR_STALE"},
		6: {Name: "NO_LLGR"},
	},
	"E_State_Flags": {
		1: {Name: "EXTERNAL_FLAG"},
		2: {Name: "READVERTISEMENT_FLAG"},
		3: {Name: "NODE_FLAG"},
	},
	"E_State_SegmentType": {
		1: {Name: "AS_SEQUENCE"},
		2: {Name: "AS_SET"},
		3: {Name: "AS_CONFED_SEQUENCE"},
		4: {Name: "AS_CONFED_SET"},
	},
	"E_Tags_ValueType": {
		1: {Name: "STRING"},
		2: {Name: "HEX"},
		3: {Name: "BOOL"},
		4: {Name: "COUNTER64"},
		5: {Name: "FLOAT32"},
		6: {Name: "IPV4"},
		7: {Name: "IPV6"},
		8: {Name: "MAC"},
	},
	"E_UnicastIpv4Prefix_Origin": {
		1: {Name: "IGP"},
		2: {Name: "EGP"},
		3: {Name: "INCOMPLETE"},
	},
	"E_UnicastIpv6Prefix_Origin": {
		1: {Name: "IGP"},
		2: {Name: "EGP"},
		3: {Name: "INCOMPLETE"},
	},
}

// ΛEnumTypes is a map, keyed by a YANG schema path, of the enumerated types that
// correspond with the leaf. The type is represented as a reflect.Type. The naming
// of the map ensures that there are no clashes with valid YANG identifiers.
func initΛEnumTypes() {
	ΛEnumTypes = map[string][]reflect.Type{
		"/bgp-peers/bgp-peer/state/session-state": {
			reflect.TypeOf((E_BgpPeer_SessionState)(0)),
		},
		"/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/segment-type": {
			reflect.TypeOf((E_State_SegmentType)(0)),
		},
		"/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/community-type": {
			reflect.TypeOf((E_State_CommunityType)(0)),
		},
		"/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/origin": {
			reflect.TypeOf((E_UnicastIpv4Prefix_Origin)(0)),
		},
		"/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/as-path/segment-type": {
			reflect.TypeOf((E_State_SegmentType)(0)),
		},
		"/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/community/community-type": {
			reflect.TypeOf((E_State_CommunityType)(0)),
		},
		"/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/origin": {
			reflect.TypeOf((E_UnicastIpv6Prefix_Origin)(0)),
		},
		"/flows/flow/tagged-metrics/tagged-metric/state/tags/tag-value/value-type": {
			reflect.TypeOf((E_Tags_ValueType)(0)),
		},
		"/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/flags": {
			reflect.TypeOf((E_Lsps_Flags)(0)),
		},
		"/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-type": {
			reflect.TypeOf((E_Lsps_PduType)(0)),
		},
		"/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes/flags": {
			reflect.TypeOf((E_State_Flags)(0)),
		},
		"/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/redistribution-type": {
			reflect.TypeOf((E_ExtendedIpv4Reachability_Prefix_RedistributionType)(0)),
		},
		"/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/origin-type": {
			reflect.TypeOf((E_Ipv4ExternalReachability_Prefix_OriginType)(0)),
		},
		"/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/redistribution-type": {
			reflect.TypeOf((E_Ipv4ExternalReachability_Prefix_RedistributionType)(0)),
		},
		"/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/origin-type": {
			reflect.TypeOf((E_Ipv4ExternalReachability_Prefix_OriginType)(0)),
		},
		"/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/redistribution-type": {
			reflect.TypeOf((E_Ipv4ExternalReachability_Prefix_RedistributionType)(0)),
		},
		"/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/origin-type": {
			reflect.TypeOf((E_Ipv6Reachability_Prefix_OriginType)(0)),
		},
		"/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes/flags": {
			reflect.TypeOf((E_State_Flags)(0)),
		},
		"/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/redistribution-type": {
			reflect.TypeOf((E_Ipv6Reachability_Prefix_RedistributionType)(0)),
		},
		"/lacp/lag-members/lag-member/state/activity": {
			reflect.TypeOf((E_Lacp_LacpActivityType)(0)),
		},
		"/lacp/lag-members/lag-member/state/synchronization": {
			reflect.TypeOf((E_Lacp_LacpSynchronizationType)(0)),
		},
		"/lacp/lag-members/lag-member/state/timeout": {
			reflect.TypeOf((E_Lacp_LacpTimeoutType)(0)),
		},
		"/lags/lag/state/oper-status": {
			reflect.TypeOf((E_Lag_OperStatus)(0)),
		},
		"/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/name": {
			reflect.TypeOf((E_Capabilities_Name)(0)),
		},
		"/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id-type": {
			reflect.TypeOf((E_LldpNeighbor_ChassisIdType)(0)),
		},
		"/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id-type": {
			reflect.TypeOf((E_LldpNeighbor_PortIdType)(0)),
		},
		"/ports/port/state/link": {
			reflect.TypeOf((E_Port_Link)(0)),
		},
		"/rsvpte-routers/rsvpte-router/state/label-switched-path-database/lsps/ipv4-lsp/state/ero/state/type": {
			reflect.TypeOf((E_Ero_Type)(0)),
		},
		"/rsvpte-routers/rsvpte-router/state/label-switched-path-database/lsps/ipv4-lsp/state/label-in": {
			reflect.TypeOf((E_Ipv4Lsp_LabelIn)(0)),
		},
		"/rsvpte-routers/rsvpte-router/state/label-switched-path-database/lsps/ipv4-lsp/state/label-out": {
			reflect.TypeOf((E_Ipv4Lsp_LabelOut)(0)),
		},
		"/rsvpte-routers/rsvpte-router/state/label-switched-path-database/lsps/ipv4-lsp/state/last-flap-reason": {
			reflect.TypeOf((E_Ipv4Lsp_LastFlapReason)(0)),
		},
		"/rsvpte-routers/rsvpte-router/state/label-switched-path-database/lsps/ipv4-lsp/state/rro/state/reported-label": {
			reflect.TypeOf((E_Rro_ReportedLabel)(0)),
		},
		"/rsvpte-routers/rsvpte-router/state/label-switched-path-database/lsps/ipv4-lsp/state/session-status": {
			reflect.TypeOf((E_Ipv4Lsp_SessionStatus)(0)),
		},
	}
}
