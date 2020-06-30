// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/v1alpha1/access_policy.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for AccessPolicySpec
func (this *AccessPolicySpec) MarshalJSON() ([]byte, error) {
	str, err := AccessPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AccessPolicySpec
func (this *AccessPolicySpec) UnmarshalJSON(b []byte) error {
	return AccessPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AccessPolicyStatus
func (this *AccessPolicyStatus) MarshalJSON() ([]byte, error) {
	str, err := AccessPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AccessPolicyStatus
func (this *AccessPolicyStatus) UnmarshalJSON(b []byte) error {
	return AccessPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AccessPolicyStatus_TranslatorError
func (this *AccessPolicyStatus_TranslatorError) MarshalJSON() ([]byte, error) {
	str, err := AccessPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AccessPolicyStatus_TranslatorError
func (this *AccessPolicyStatus_TranslatorError) UnmarshalJSON(b []byte) error {
	return AccessPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	AccessPolicyMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	AccessPolicyUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
