// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v1alpha1

import (
	proto "github.com/gogo/protobuf/proto"
)

// DeepCopyInto for the VirtualMeshCertificateSigningRequest.Spec
func (in *VirtualMeshCertificateSigningRequestSpec) DeepCopyInto(out *VirtualMeshCertificateSigningRequestSpec) {
	p := proto.Clone(in).(*VirtualMeshCertificateSigningRequestSpec)
	*out = *p
}

// DeepCopyInto for the VirtualMeshCertificateSigningRequest.Status
func (in *VirtualMeshCertificateSigningRequestStatus) DeepCopyInto(out *VirtualMeshCertificateSigningRequestStatus) {
	p := proto.Clone(in).(*VirtualMeshCertificateSigningRequestStatus)
	*out = *p
}
