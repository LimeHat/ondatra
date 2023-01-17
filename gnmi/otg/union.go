/*
Package otg is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was true
in this case).

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
package otg

import (
	"fmt"
)

// RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelIn_Union is an interface that is implemented by valid types for the union
// for the leaf /open-traffic-generator-rsvp/rsvpte-routers/rsvpte-router/state/label-switched-path-database/lsps/ipv4-lsp/state/label-in within the YANG schema.
// Union type can be one of [E_Ipv4Lsp_LabelIn, UnionUint32].
type RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelIn_Union interface {
	// Union type can be one of [E_Ipv4Lsp_LabelIn, UnionUint32]
	Documentation_for_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelIn_Union()
}

// Documentation_for_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelIn_Union ensures that E_Ipv4Lsp_LabelIn
// implements the RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelIn_Union interface.
func (E_Ipv4Lsp_LabelIn) Documentation_for_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelIn_Union() {
}

// Documentation_for_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelIn_Union ensures that UnionUint32
// implements the RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelIn_Union interface.
func (UnionUint32) Documentation_for_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelIn_Union() {}

// To_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelIn_Union takes an input interface{} and attempts to convert it to a struct
// which implements the RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelIn_Union union. It returns an error if the interface{} supplied
// cannot be converted to a type within the union.
func (t *RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp) To_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelIn_Union(i interface{}) (RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelIn_Union, error) {
	if v, ok := i.(RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelIn_Union); ok {
		return v, nil
	}
	switch v := i.(type) {
	case uint32:
		return UnionUint32(v), nil
	}
	return nil, fmt.Errorf("cannot convert %v to RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelIn_Union, unknown union type, got: %T, want any of [E_Ipv4Lsp_LabelIn, uint32]", i, i)
}

// RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelOut_Union is an interface that is implemented by valid types for the union
// for the leaf /open-traffic-generator-rsvp/rsvpte-routers/rsvpte-router/state/label-switched-path-database/lsps/ipv4-lsp/state/label-out within the YANG schema.
// Union type can be one of [E_Ipv4Lsp_LabelOut, UnionUint32].
type RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelOut_Union interface {
	// Union type can be one of [E_Ipv4Lsp_LabelOut, UnionUint32]
	Documentation_for_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelOut_Union()
}

// Documentation_for_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelOut_Union ensures that E_Ipv4Lsp_LabelOut
// implements the RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelOut_Union interface.
func (E_Ipv4Lsp_LabelOut) Documentation_for_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelOut_Union() {
}

// Documentation_for_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelOut_Union ensures that UnionUint32
// implements the RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelOut_Union interface.
func (UnionUint32) Documentation_for_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelOut_Union() {
}

// To_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelOut_Union takes an input interface{} and attempts to convert it to a struct
// which implements the RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelOut_Union union. It returns an error if the interface{} supplied
// cannot be converted to a type within the union.
func (t *RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp) To_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelOut_Union(i interface{}) (RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelOut_Union, error) {
	if v, ok := i.(RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelOut_Union); ok {
		return v, nil
	}
	switch v := i.(type) {
	case uint32:
		return UnionUint32(v), nil
	}
	return nil, fmt.Errorf("cannot convert %v to RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_LabelOut_Union, unknown union type, got: %T, want any of [E_Ipv4Lsp_LabelOut, uint32]", i, i)
}

// RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro_ReportedLabel_Union is an interface that is implemented by valid types for the union
// for the leaf /open-traffic-generator-rsvp/rsvpte-routers/rsvpte-router/state/label-switched-path-database/lsps/ipv4-lsp/state/rro/state/reported-label within the YANG schema.
// Union type can be one of [E_Rro_ReportedLabel, UnionUint32].
type RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro_ReportedLabel_Union interface {
	// Union type can be one of [E_Rro_ReportedLabel, UnionUint32]
	Documentation_for_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro_ReportedLabel_Union()
}

// Documentation_for_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro_ReportedLabel_Union ensures that E_Rro_ReportedLabel
// implements the RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro_ReportedLabel_Union interface.
func (E_Rro_ReportedLabel) Documentation_for_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro_ReportedLabel_Union() {
}

// Documentation_for_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro_ReportedLabel_Union ensures that UnionUint32
// implements the RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro_ReportedLabel_Union interface.
func (UnionUint32) Documentation_for_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro_ReportedLabel_Union() {
}

// To_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro_ReportedLabel_Union takes an input interface{} and attempts to convert it to a struct
// which implements the RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro_ReportedLabel_Union union. It returns an error if the interface{} supplied
// cannot be converted to a type within the union.
func (t *RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro) To_RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro_ReportedLabel_Union(i interface{}) (RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro_ReportedLabel_Union, error) {
	if v, ok := i.(RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro_ReportedLabel_Union); ok {
		return v, nil
	}
	switch v := i.(type) {
	case uint32:
		return UnionUint32(v), nil
	}
	return nil, fmt.Errorf("cannot convert %v to RsvpteRouter_LabelSwitchedPathDatabase_Ipv4Lsp_Rro_ReportedLabel_Union, unknown union type, got: %T, want any of [E_Rro_ReportedLabel, uint32]", i, i)
}
