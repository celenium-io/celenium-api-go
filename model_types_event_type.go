/*
Celenium API

This is docs of Celenium API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package celenium

import (
	"encoding/json"
	"fmt"
)

// TypesEventType the model 'TypesEventType'
type TypesEventType string

// List of types.EventType
const (
	EventTypeUnknown TypesEventType = "unknown"
	EventTypeCoinReceived TypesEventType = "coin_received"
	EventTypeCoinbase TypesEventType = "coinbase"
	EventTypeCoinSpent TypesEventType = "coin_spent"
	EventTypeBurn TypesEventType = "burn"
	EventTypeMint TypesEventType = "mint"
	EventTypeMessage TypesEventType = "message"
	EventTypeProposerReward TypesEventType = "proposer_reward"
	EventTypeRewards TypesEventType = "rewards"
	EventTypeCommission TypesEventType = "commission"
	EventTypeLiveness TypesEventType = "liveness"
	EventTypeTransfer TypesEventType = "transfer"
	EventTypeCelestiablobv1EventPayForBlobs TypesEventType = "celestia.blob.v1.EventPayForBlobs"
	EventTypeRedelegate TypesEventType = "redelegate"
	EventTypeAttestationRequest TypesEventType = "AttestationRequest"
	EventTypeWithdrawRewards TypesEventType = "withdraw_rewards"
	EventTypeWithdrawCommission TypesEventType = "withdraw_commission"
	EventTypeSetWithdrawAddress TypesEventType = "set_withdraw_address"
	EventTypeCreateValidator TypesEventType = "create_validator"
	EventTypeDelegate TypesEventType = "delegate"
	EventTypeEditValidator TypesEventType = "edit_validator"
	EventTypeUnbond TypesEventType = "unbond"
	EventTypeTx TypesEventType = "tx"
	EventTypeCompleteRedelegation TypesEventType = "complete_redelegation"
	EventTypeCompleteUnbonding TypesEventType = "complete_unbonding"
	EventTypeUseFeegrant TypesEventType = "use_feegrant"
	EventTypeRevokeFeegrant TypesEventType = "revoke_feegrant"
	EventTypeSetFeegrant TypesEventType = "set_feegrant"
	EventTypeUpdateFeegrant TypesEventType = "update_feegrant"
	EventTypeSlash TypesEventType = "slash"
	EventTypeProposalVote TypesEventType = "proposal_vote"
	EventTypeProposalDeposit TypesEventType = "proposal_deposit"
	EventTypeSubmitProposal TypesEventType = "submit_proposal"
	EventTypeCosmosauthzv1beta1EventGrant TypesEventType = "cosmos.authz.v1beta1.EventGrant"
	EventTypeSendPacket TypesEventType = "send_packet"
	EventTypeIbcTransfer TypesEventType = "ibc_transfer"
	EventTypeFungibleTokenPacket TypesEventType = "fungible_token_packet"
	EventTypeAcknowledgePacket TypesEventType = "acknowledge_packet"
	EventTypeCreateClient TypesEventType = "create_client"
	EventTypeUpdateClient TypesEventType = "update_client"
	EventTypeConnectionOpenTry TypesEventType = "connection_open_try"
	EventTypeConnectionOpenInit TypesEventType = "connection_open_init"
	EventTypeConnectionOpenConfirm TypesEventType = "connection_open_confirm"
	EventTypeConnectionOpenAck TypesEventType = "connection_open_ack"
	EventTypeChannelOpenTry TypesEventType = "channel_open_try"
	EventTypeChannelOpenInit TypesEventType = "channel_open_init"
	EventTypeChannelOpenConfirm TypesEventType = "channel_open_confirm"
	EventTypeChannelOpenAck TypesEventType = "channel_open_ack"
	EventTypeRecvPacket TypesEventType = "recv_packet"
	EventTypeWriteAcknowledgement TypesEventType = "write_acknowledgement"
	EventTypeTimeout TypesEventType = "timeout"
	EventTypeTimeoutPacket TypesEventType = "timeout_packet"
	EventTypeCosmosauthzv1beta1EventRevoke TypesEventType = "cosmos.authz.v1beta1.EventRevoke"
	EventTypeCosmosauthzv1EventRevoke TypesEventType = "cosmos.authz.v1.EventRevoke"
	EventTypeCancelUnbondingDelegation TypesEventType = "cancel_unbonding_delegation"
)

// All allowed values of TypesEventType enum
var AllowedTypesEventTypeEnumValues = []TypesEventType{
	"unknown",
	"coin_received",
	"coinbase",
	"coin_spent",
	"burn",
	"mint",
	"message",
	"proposer_reward",
	"rewards",
	"commission",
	"liveness",
	"transfer",
	"celestia.blob.v1.EventPayForBlobs",
	"redelegate",
	"AttestationRequest",
	"withdraw_rewards",
	"withdraw_commission",
	"set_withdraw_address",
	"create_validator",
	"delegate",
	"edit_validator",
	"unbond",
	"tx",
	"complete_redelegation",
	"complete_unbonding",
	"use_feegrant",
	"revoke_feegrant",
	"set_feegrant",
	"update_feegrant",
	"slash",
	"proposal_vote",
	"proposal_deposit",
	"submit_proposal",
	"cosmos.authz.v1beta1.EventGrant",
	"send_packet",
	"ibc_transfer",
	"fungible_token_packet",
	"acknowledge_packet",
	"create_client",
	"update_client",
	"connection_open_try",
	"connection_open_init",
	"connection_open_confirm",
	"connection_open_ack",
	"channel_open_try",
	"channel_open_init",
	"channel_open_confirm",
	"channel_open_ack",
	"recv_packet",
	"write_acknowledgement",
	"timeout",
	"timeout_packet",
	"cosmos.authz.v1beta1.EventRevoke",
	"cosmos.authz.v1.EventRevoke",
	"cancel_unbonding_delegation",
}

func (v *TypesEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesEventType(value)
	for _, existing := range AllowedTypesEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesEventType", value)
}

// NewTypesEventTypeFromValue returns a pointer to a valid TypesEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesEventTypeFromValue(v string) (*TypesEventType, error) {
	ev := TypesEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesEventType: valid values are %v", v, AllowedTypesEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesEventType) IsValid() bool {
	for _, existing := range AllowedTypesEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.EventType value
func (v TypesEventType) Ptr() *TypesEventType {
	return &v
}

type NullableTypesEventType struct {
	value *TypesEventType
	isSet bool
}

func (v NullableTypesEventType) Get() *TypesEventType {
	return v.value
}

func (v *NullableTypesEventType) Set(val *TypesEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesEventType(val *TypesEventType) *NullableTypesEventType {
	return &NullableTypesEventType{value: val, isSet: true}
}

func (v NullableTypesEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
