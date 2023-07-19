// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package binding

import (
	"errors"
	"fmt"
	"io"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	"github.com/open-traffic-generator/snappi/gosnappi"
	"google.golang.org/grpc"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	bpb "github.com/openconfig/gnoi/bgp"
	cpb "github.com/openconfig/gnoi/cert"
	dpb "github.com/openconfig/gnoi/diag"
	frpb "github.com/openconfig/gnoi/factory_reset"
	fpb "github.com/openconfig/gnoi/file"
	hpb "github.com/openconfig/gnoi/healthz"
	lpb "github.com/openconfig/gnoi/layer2"
	mpb "github.com/openconfig/gnoi/mpls"
	ospb "github.com/openconfig/gnoi/os"
	otpb "github.com/openconfig/gnoi/otdr"
	plqpb "github.com/openconfig/gnoi/packet_link_qualification"
	spb "github.com/openconfig/gnoi/system"
	wpb "github.com/openconfig/gnoi/wavelength_router"
	accpb "github.com/openconfig/gnsi/accounting"
	authzpb "github.com/openconfig/gnsi/authz"
	certzpb "github.com/openconfig/gnsi/certz"
	credzpb "github.com/openconfig/gnsi/credentialz"
	pathzpb "github.com/openconfig/gnsi/pathz"

	grpb "github.com/openconfig/gribi/v1/proto/service"
	opb "github.com/openconfig/ondatra/proto"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"
)

var _ DUT = &AbstractDUT{}

// AbstractDUT is a reserved DUT.
// All implementations of the DUT interface must embed this type.
type AbstractDUT struct {
	*Dims
}

// Name returns the name of the DUT.
func (d *AbstractDUT) Name() string {
	return d.Dims.Name
}

// Vendor returns the vendor of the DUT.
func (d *AbstractDUT) Vendor() opb.Device_Vendor {
	return d.Dims.Vendor
}

// SoftwareVersion returns the software version of the DUT.
func (d *AbstractDUT) SoftwareVersion() string {
	return d.Dims.SoftwareVersion
}

// HardwareModel returns the hardware model of the DUT.
func (d *AbstractDUT) HardwareModel() string {
	return d.Dims.HardwareModel
}

// Ports returns the reserved ports on the DUT.
func (d *AbstractDUT) Ports() map[string]*Port {
	return d.Dims.Ports
}

// CustomData returns custom data for the DUT.
func (d *AbstractDUT) CustomData() map[string]any {
	return d.Dims.CustomData
}

func (d *AbstractDUT) String() string {
	return fmt.Sprintf("DUT%+v", *d)
}

// PushConfig returns an unimplemented error.
func (*AbstractDUT) PushConfig(ctx context.Context, config string, reset bool) error {
	return errors.New("PushConfig unimplemented")
}

// DialCLI returns an unimplemented error.
func (*AbstractDUT) DialCLI(context.Context) (StreamClient, error) {
	return nil, errors.New("DialCLI unimplemented")
}

// DialConsole returns an unimplemented error.
func (*AbstractDUT) DialConsole(context.Context) (StreamClient, error) {
	return nil, errors.New("DialConsole unimplemented")
}

// DialGNMI returns an unimplemented error.
func (*AbstractDUT) DialGNMI(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	return nil, errors.New("DialGNMI unimplemented")
}

// DialGNOI returns an unimplemented error.
func (*AbstractDUT) DialGNOI(context.Context, ...grpc.DialOption) (GNOIClients, error) {
	return nil, errors.New("DialGNOI unimplemented")
}

// DialGNSI returns an unimplemented error.
func (*AbstractDUT) DialGNSI(context.Context, ...grpc.DialOption) (GNSIClients, error) {
	return nil, errors.New("DialGNSI unimplemented")
}

// DialGRIBI returns an unimplemented error.
func (*AbstractDUT) DialGRIBI(context.Context, ...grpc.DialOption) (grpb.GRIBIClient, error) {
	return nil, errors.New("DialGRIBI unimplemented")
}

// DialP4RT returns an unimplemented error.
func (*AbstractDUT) DialP4RT(context.Context, ...grpc.DialOption) (p4pb.P4RuntimeClient, error) {
	return nil, errors.New("DialP4RT unimplemented")
}

func (*AbstractDUT) mustEmbedAbstractDUT() {}

var _ ATE = &AbstractATE{}

// AbstractATE is implementation support for the ATE interface.
// All implementations of the ATE interface must embed this type.
type AbstractATE struct {
	Dims *Dims
}

// Name returns the name of the ATE.
func (a *AbstractATE) Name() string {
	return a.Dims.Name
}

// Vendor returns the vendor of the ATE.
func (a *AbstractATE) Vendor() opb.Device_Vendor {
	return a.Dims.Vendor
}

// SoftwareVersion returns the software version of the ATE.
func (a *AbstractATE) SoftwareVersion() string {
	return a.Dims.SoftwareVersion
}

// HardwareModel returns the hardware model of the ATE.
func (a *AbstractATE) HardwareModel() string {
	return a.Dims.HardwareModel
}

// Ports returns the reserved ports on the ATE.
func (a *AbstractATE) Ports() map[string]*Port {
	return a.Dims.Ports
}

// CustomData returns custom data for the ATE.
func (a *AbstractATE) CustomData() map[string]any {
	return a.Dims.CustomData
}

func (a *AbstractATE) String() string {
	return fmt.Sprintf("ATE%+v", *a)
}

// DialIxNetwork returns an unimplemented error.
func (*AbstractATE) DialIxNetwork(context.Context) (*IxNetwork, error) {
	return nil, errors.New("DialIxNetwork unimplemented")
}

// DialGNMI returns an unimplemented error.
func (*AbstractATE) DialGNMI(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error) {
	return nil, errors.New("DialGNMI unimplemented")
}

// DialOTG returns an unimplemented error.
func (*AbstractATE) DialOTG(context.Context, ...grpc.DialOption) (gosnappi.GosnappiApi, error) {
	return nil, errors.New("DialOTG unimplemented")
}

func (*AbstractATE) mustEmbedAbstractATE() {}

var _ GNOIClients = &AbstractGNOIClients{}

// AbstractGNOIClients is implementation support for the GNOIClients interface.
type AbstractGNOIClients struct{}

// BGP logs a fatal unimplemented error.
func (*AbstractGNOIClients) BGP() bpb.BGPClient {
	log.Fatal("BGP unimplemented")
	return nil
}

// CertificateManagement logs a fatal unimplemented error.
func (*AbstractGNOIClients) CertificateManagement() cpb.CertificateManagementClient {
	log.Fatal("CertificateManagement unimplemented")
	return nil
}

// Diag logs a fatal unimplemented error.
func (*AbstractGNOIClients) Diag() dpb.DiagClient {
	log.Fatal("Diag unimplemented")
	return nil
}

// FactoryReset logs a fatal unimplemented error.
func (*AbstractGNOIClients) FactoryReset() frpb.FactoryResetClient {
	log.Fatal("FactoryReset unimplemented")
	return nil
}

// File logs a fatal unimplemented error.
func (*AbstractGNOIClients) File() fpb.FileClient {
	log.Fatal("File unimplemented")
	return nil
}

// Healthz logs a fatal unimplemented error.
func (*AbstractGNOIClients) Healthz() hpb.HealthzClient {
	log.Fatal("Healthz unimplemented")
	return nil
}

// Layer2 logs a fatal unimplemented error.
func (*AbstractGNOIClients) Layer2() lpb.Layer2Client {
	log.Fatal("Layer2 unimplemented")
	return nil
}

// LinkQualification logs a fatal unimplemented error.
func (*AbstractGNOIClients) LinkQualification() plqpb.LinkQualificationClient {
	log.Fatal("LinkQualification unimplemented")
	return nil
}

// MPLS logs a fatal unimplemented error.
func (*AbstractGNOIClients) MPLS() mpb.MPLSClient {
	log.Fatal("MPLS unimplemented")
	return nil
}

// OS logs a fatal unimplemented error.
func (*AbstractGNOIClients) OS() ospb.OSClient {
	log.Fatal("OS unimplemented")
	return nil
}

// OTDR logs a fatal unimplemented error.
func (*AbstractGNOIClients) OTDR() otpb.OTDRClient {
	log.Fatal("OTDR unimplemented")
	return nil
}

// System logs a fatal unimplemented error.
func (*AbstractGNOIClients) System() spb.SystemClient {
	log.Fatal("System unimplemented")
	return nil
}

// WavelengthRouter logs a fatal unimplemented error.
func (*AbstractGNOIClients) WavelengthRouter() wpb.WavelengthRouterClient {
	log.Fatal("WavelengthRouter unimplemented")
	return nil
}

func (g *AbstractGNOIClients) mustEmbedAbstractGNOIClients() {}

var _ GNSIClients = &AbstractGNSIClients{}

// AbstractGNSIClients is implementation support for the GNSIClients interface.
type AbstractGNSIClients struct{}

// Authz logs a fatal unimplemented error.
func (*AbstractGNSIClients) Authz() authzpb.AuthzClient {
	log.Fatal("Authz unimplemented")
	return nil
}

// Pathz logs a fatal unimplemented error.
func (*AbstractGNSIClients) Pathz() pathzpb.PathzClient {
	log.Fatal("Pathz unimplemented")
	return nil
}

// Certz logs a fatal unimplemented error.
func (*AbstractGNSIClients) Certz() certzpb.CertzClient {
	log.Fatal("Certz unimplemented")
	return nil
}

// Credentialz logs a fatal unimplemented error.
func (*AbstractGNSIClients) Credentialz() credzpb.CredentialzClient {
	log.Fatal("Credentialz unimplemented")
	return nil
}

// AccountingPull logs a fatal unimplemented error.
func (*AbstractGNSIClients) AccountingPull() accpb.AccountingPullClient {
	log.Fatal("AccountingPull unimplemented")
	return nil
}

// AccountingPush logs a fatal unimplemented error.
func (*AbstractGNSIClients) AccountingPush() accpb.AccountingPushClient {
	log.Fatal("AccountingPush unimplemented")
	return nil
}

func (*AbstractGNSIClients) mustEmbedAbstractGNSIClients() {}

var _ StreamClient = &AbstractStreamClient{}

// AbstractStreamClient is implementation support for the StreamClient interface.
type AbstractStreamClient struct{}

// SendCommand returns an unimplemented error.
func (*AbstractStreamClient) SendCommand(ctx context.Context, cmd string) (string, error) {
	return "", errors.New("SendCommand unimplemented")
}

// Stdin logs a fatal unimplemented error.
func (*AbstractStreamClient) Stdin() io.WriteCloser {
	log.Fatal("Stdin unimplemented")
	return nil
}

// Stdout logs a fatal unimplemented error.
func (*AbstractStreamClient) Stdout() io.ReadCloser {
	log.Fatal("Stdout unimplemented")
	return nil
}

// Stderr logs a fatal unimplemented error.
func (*AbstractStreamClient) Stderr() io.ReadCloser {
	log.Fatal("Stderr unimplemented")
	return nil
}

// Close returns an unimplemented error.
func (*AbstractStreamClient) Close() error {
	return errors.New("Close unimplemented")
}

func (*AbstractStreamClient) mustEmbedAbstractStreamClient() {}
