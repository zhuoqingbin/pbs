// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/zhuoqingbin/pbs/evsepb/evse_ctrl.proto

package evsepb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *APDU) Validate() error {
	if this.Extrend != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Extrend); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Extrend", err)
		}
	}
	return nil
}
func (this *Extrend) Validate() error {
	return nil
}
func (this *DeviceRegistrationReq) Validate() error {
	return nil
}
func (this *DeviceRegistrationConf) Validate() error {
	return nil
}
func (this *BootNotificationReq) Validate() error {
	if this.EvseModel != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EvseModel); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EvseModel", err)
		}
	}
	return nil
}
func (this *BootNotificationConf) Validate() error {
	return nil
}
func (this *HeartbeatReq) Validate() error {
	return nil
}
func (this *HeartbeatResp) Validate() error {
	return nil
}
func (this *SampledValue) Validate() error {
	return nil
}
func (this *TelemetryReq) Validate() error {
	for _, item := range this.Values {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Values", err)
			}
		}
	}
	return nil
}
func (this *SetEvseConfigReq) Validate() error {
	for _, item := range this.ConfiguKey {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ConfiguKey", err)
			}
		}
	}
	return nil
}
func (this *SetEvseConfigConf) Validate() error {
	return nil
}
func (this *GetEvseConfigReq) Validate() error {
	for _, item := range this.Keys {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Keys", err)
			}
		}
	}
	return nil
}
func (this *GetEvseConfigConf) Validate() error {
	for _, item := range this.Values {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Values", err)
			}
		}
	}
	return nil
}
func (this *RemoteControlReq) Validate() error {
	return nil
}
func (this *RemoteControlConf) Validate() error {
	return nil
}
func (this *RemoteControlStartParam) Validate() error {
	return nil
}
func (this *TransactionReq) Validate() error {
	if this.Cost != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Cost); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Cost", err)
		}
	}
	return nil
}
func (this *TransactionCost) Validate() error {
	for _, item := range this.Details {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Details", err)
			}
		}
	}
	return nil
}
func (this *TransactionCostDetail) Validate() error {
	return nil
}
func (this *TransactionConf) Validate() error {
	return nil
}
func (this *OTAReq) Validate() error {
	return nil
}
func (this *OTAConf) Validate() error {
	return nil
}
func (this *EvseWarningReq) Validate() error {
	return nil
}
func (this *WarningConf) Validate() error {
	return nil
}
func (this *RemoteStartReq) Validate() error {
	return nil
}
func (this *RemoteStartConf) Validate() error {
	return nil
}
func (this *StartTransactionReq) Validate() error {
	return nil
}
func (this *StartTransactionConf) Validate() error {
	return nil
}
