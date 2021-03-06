// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package networkserver

import (
	"context"
	"testing"

	pbtypes "github.com/gogo/protobuf/types"
	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/band"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestNeedsTxParamSetupReq(t *testing.T) {
	type TestCase struct {
		Name        string
		InputDevice *ttnpb.EndDevice
		Band        *band.Band
		Needs       bool
	}
	tcs := []TestCase{
		{
			Name:        "no MAC state",
			InputDevice: &ttnpb.EndDevice{},
			Band:        LoRaWANBands[band.EU_863_870][ttnpb.PHY_V1_1_REV_B],
		},
	}
	for _, conf := range []struct {
		Suffix                               string
		CurrentParameters, DesiredParameters ttnpb.MACParameters
		Needs                                bool
	}{
		{
			Suffix: "current(EIRP:26,downlink:nil,uplink:nil),desired(EIRP:26,downlink:nil,uplink:nil)",
			CurrentParameters: ttnpb.MACParameters{
				MaxEIRP: 26,
			},
			DesiredParameters: ttnpb.MACParameters{
				MaxEIRP: 26,
			},
		},
		{
			Suffix: "current(EIRP:26,downlink:true,uplink:false),desired(EIRP:26,downlink:nil,uplink:nil)",
			CurrentParameters: ttnpb.MACParameters{
				MaxEIRP:           26,
				DownlinkDwellTime: &pbtypes.BoolValue{Value: true},
				UplinkDwellTime:   &pbtypes.BoolValue{Value: false},
			},
			DesiredParameters: ttnpb.MACParameters{
				MaxEIRP: 26,
			},
		},
		{
			Suffix: "current(EIRP:26,downlink:nil,uplink:nil),desired(EIRP:26,downlink:true,uplink:true)",
			CurrentParameters: ttnpb.MACParameters{
				MaxEIRP: 26,
			},
			DesiredParameters: ttnpb.MACParameters{
				MaxEIRP:           26,
				DownlinkDwellTime: &pbtypes.BoolValue{Value: true},
				UplinkDwellTime:   &pbtypes.BoolValue{Value: true},
			},
			Needs: true,
		},
		{
			Suffix: "current(EIRP:26,downlink:nil,uplink:nil),desired(EIRP:26,downlink:false,uplink:false)",
			CurrentParameters: ttnpb.MACParameters{
				MaxEIRP: 26,
			},
			DesiredParameters: ttnpb.MACParameters{
				MaxEIRP:           26,
				DownlinkDwellTime: &pbtypes.BoolValue{Value: false},
				UplinkDwellTime:   &pbtypes.BoolValue{Value: false},
			},
			Needs: true,
		},
		{
			Suffix: "current(EIRP:26,downlink:true,uplink:nil),desired(EIRP:26,downlink:false,uplink:false)",
			CurrentParameters: ttnpb.MACParameters{
				MaxEIRP:           26,
				DownlinkDwellTime: &pbtypes.BoolValue{Value: true},
			},
			DesiredParameters: ttnpb.MACParameters{
				MaxEIRP:           26,
				DownlinkDwellTime: &pbtypes.BoolValue{Value: false},
				UplinkDwellTime:   &pbtypes.BoolValue{Value: false},
			},
			Needs: true,
		},
		{
			Suffix: "current(EIRP:24,downlink:true,uplink:false),desired(EIRP:26,downlink:true,uplink:false)",
			CurrentParameters: ttnpb.MACParameters{
				MaxEIRP:           24,
				DownlinkDwellTime: &pbtypes.BoolValue{Value: true},
				UplinkDwellTime:   &pbtypes.BoolValue{Value: false},
			},
			DesiredParameters: ttnpb.MACParameters{
				MaxEIRP:           26,
				DownlinkDwellTime: &pbtypes.BoolValue{Value: true},
				UplinkDwellTime:   &pbtypes.BoolValue{Value: false},
			},
			Needs: true,
		},
	} {
		ForEachBandMACVersion(t, func(makeName func(parts ...string) string, phy *band.Band, phyVersion ttnpb.PHYVersion, macVersion ttnpb.MACVersion) {
			tcs = append(tcs,
				TestCase{
					Name: makeName(conf.Suffix),
					InputDevice: &ttnpb.EndDevice{
						LoRaWANVersion:    macVersion,
						LoRaWANPHYVersion: phyVersion,
						MACState: &ttnpb.MACState{
							LoRaWANVersion:    macVersion,
							CurrentParameters: conf.CurrentParameters,
							DesiredParameters: conf.DesiredParameters,
						},
					},
					Band:  phy,
					Needs: phy.TxParamSetupReqSupport && conf.Needs && macVersion.Compare(ttnpb.MAC_V1_0_2) >= 0,
				},
			)
		})
	}

	for _, tc := range tcs {
		test.RunSubtest(t, test.SubtestConfig{
			Name:     tc.Name,
			Parallel: true,
			Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
				dev := CopyEndDevice(tc.InputDevice)
				res := deviceNeedsTxParamSetupReq(dev, tc.Band)
				if tc.Needs {
					a.So(res, should.BeTrue)
				} else {
					a.So(res, should.BeFalse)
				}
				a.So(dev, should.Resemble, tc.InputDevice)
			},
		})
	}
}

func TestEnqueueTxParamSetupReq(t *testing.T) {
	for _, tc := range []struct {
		Name                        string
		InputDevice, ExpectedDevice *ttnpb.EndDevice
		MaxDownlinkLength           uint16
		MaxUplinkLength             uint16
		State                       macCommandEnqueueState
	}{
		{
			Name: "payload fits/EIRP 26/dwell time both",
			InputDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					LoRaWANVersion: ttnpb.MAC_V1_1,
					CurrentParameters: ttnpb.MACParameters{
						MaxEIRP: 26,
					},
					DesiredParameters: ttnpb.MACParameters{
						MaxEIRP:           26,
						DownlinkDwellTime: &pbtypes.BoolValue{Value: true},
						UplinkDwellTime:   &pbtypes.BoolValue{Value: true},
					},
				},
			},
			ExpectedDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					LoRaWANVersion: ttnpb.MAC_V1_1,
					CurrentParameters: ttnpb.MACParameters{
						MaxEIRP: 26,
					},
					DesiredParameters: ttnpb.MACParameters{
						MaxEIRP:           26,
						DownlinkDwellTime: &pbtypes.BoolValue{Value: true},
						UplinkDwellTime:   &pbtypes.BoolValue{Value: true},
					},
					PendingRequests: []*ttnpb.MACCommand{
						(&ttnpb.MACCommand_TxParamSetupReq{
							MaxEIRPIndex:      ttnpb.DEVICE_EIRP_26,
							DownlinkDwellTime: true,
							UplinkDwellTime:   true,
						}).MACCommand(),
					},
				},
			},
			MaxDownlinkLength: 42,
			MaxUplinkLength:   24,
			State: macCommandEnqueueState{
				MaxDownLen: 40,
				MaxUpLen:   23,
				Ok:         true,
				QueuedEvents: events.Builders{
					evtEnqueueTxParamSetupRequest.With(events.WithData(&ttnpb.MACCommand_TxParamSetupReq{
						MaxEIRPIndex:      ttnpb.DEVICE_EIRP_26,
						DownlinkDwellTime: true,
						UplinkDwellTime:   true,
					})),
				},
			},
		},
		{
			Name: "payload fits/EIRP 26/no dwell time limitations",
			InputDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					LoRaWANVersion: ttnpb.MAC_V1_1,
					CurrentParameters: ttnpb.MACParameters{
						MaxEIRP: 26,
					},
					DesiredParameters: ttnpb.MACParameters{
						MaxEIRP: 26,
					},
				},
			},
			ExpectedDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					LoRaWANVersion: ttnpb.MAC_V1_1,
					CurrentParameters: ttnpb.MACParameters{
						MaxEIRP: 26,
					},
					DesiredParameters: ttnpb.MACParameters{
						MaxEIRP: 26,
					},
				},
			},
			MaxDownlinkLength: 42,
			MaxUplinkLength:   24,
			State: macCommandEnqueueState{
				MaxDownLen: 42,
				MaxUpLen:   24,
				Ok:         true,
			},
		},
		{
			Name: "downlink does not fit/EIRP 26/dwell time both",
			InputDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					LoRaWANVersion: ttnpb.MAC_V1_1,
					CurrentParameters: ttnpb.MACParameters{
						MaxEIRP: 26,
					},
					DesiredParameters: ttnpb.MACParameters{
						MaxEIRP:           26,
						DownlinkDwellTime: &pbtypes.BoolValue{Value: true},
						UplinkDwellTime:   &pbtypes.BoolValue{Value: true},
					},
				},
			},
			ExpectedDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					LoRaWANVersion: ttnpb.MAC_V1_1,
					CurrentParameters: ttnpb.MACParameters{
						MaxEIRP: 26,
					},
					DesiredParameters: ttnpb.MACParameters{
						MaxEIRP:           26,
						DownlinkDwellTime: &pbtypes.BoolValue{Value: true},
						UplinkDwellTime:   &pbtypes.BoolValue{Value: true},
					},
				},
			},
			MaxDownlinkLength: 1,
			MaxUplinkLength:   24,
			State: macCommandEnqueueState{
				MaxDownLen: 1,
				MaxUpLen:   24,
			},
		},
		{
			Name: "uplink does not fit/EIRP 26/dwell time both",
			InputDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					LoRaWANVersion: ttnpb.MAC_V1_1,
					CurrentParameters: ttnpb.MACParameters{
						MaxEIRP: 26,
					},
					DesiredParameters: ttnpb.MACParameters{
						MaxEIRP:           26,
						DownlinkDwellTime: &pbtypes.BoolValue{Value: true},
						UplinkDwellTime:   &pbtypes.BoolValue{Value: true},
					},
				},
			},
			ExpectedDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					LoRaWANVersion: ttnpb.MAC_V1_1,
					CurrentParameters: ttnpb.MACParameters{
						MaxEIRP: 26,
					},
					DesiredParameters: ttnpb.MACParameters{
						MaxEIRP:           26,
						DownlinkDwellTime: &pbtypes.BoolValue{Value: true},
						UplinkDwellTime:   &pbtypes.BoolValue{Value: true},
					},
				},
			},
			MaxDownlinkLength: 42,
			State: macCommandEnqueueState{
				MaxDownLen: 42,
			},
		},
	} {
		test.RunSubtest(t, test.SubtestConfig{
			Name:     tc.Name,
			Parallel: true,
			Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
				dev := CopyEndDevice(tc.InputDevice)

				st := enqueueTxParamSetupReq(ctx, dev, tc.MaxDownlinkLength, tc.MaxUplinkLength, LoRaWANBands[band.AS_923][ttnpb.PHY_V1_1_REV_B])
				a.So(dev, should.Resemble, tc.ExpectedDevice)
				a.So(st.QueuedEvents, should.ResembleEventBuilders, tc.State.QueuedEvents)
				st.QueuedEvents = tc.State.QueuedEvents
				a.So(st, should.Resemble, tc.State)
			},
		})
	}
}

func TestHandleTxParamSetupAns(t *testing.T) {
	for _, tc := range []struct {
		Name                        string
		InputDevice, ExpectedDevice *ttnpb.EndDevice
		Events                      events.Builders
		Error                       error
	}{
		{
			Name: "no request",
			InputDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{},
			},
			ExpectedDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{},
			},
			Events: events.Builders{
				evtReceiveTxParamSetupAnswer,
			},
			Error: errMACRequestNotFound,
		},
		{
			Name: "EIRP 26, dwell time both",
			InputDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					PendingRequests: []*ttnpb.MACCommand{
						(&ttnpb.MACCommand_TxParamSetupReq{
							MaxEIRPIndex:      ttnpb.DEVICE_EIRP_26,
							DownlinkDwellTime: true,
							UplinkDwellTime:   true,
						}).MACCommand(),
					},
				},
			},
			ExpectedDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					CurrentParameters: ttnpb.MACParameters{
						MaxEIRP:           26,
						DownlinkDwellTime: &pbtypes.BoolValue{Value: true},
						UplinkDwellTime:   &pbtypes.BoolValue{Value: true},
					},
					PendingRequests: []*ttnpb.MACCommand{},
				},
			},
			Events: events.Builders{
				evtReceiveTxParamSetupAnswer,
			},
		},
	} {
		test.RunSubtest(t, test.SubtestConfig{
			Name:     tc.Name,
			Parallel: true,
			Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
				dev := CopyEndDevice(tc.InputDevice)

				evs, err := handleTxParamSetupAns(ctx, dev)
				if tc.Error != nil && !a.So(err, should.EqualErrorOrDefinition, tc.Error) ||
					tc.Error == nil && !a.So(err, should.BeNil) {
					t.FailNow()
				}
				a.So(dev, should.Resemble, tc.ExpectedDevice)
				a.So(evs, should.ResembleEventBuilders, tc.Events)
			},
		})
	}
}
