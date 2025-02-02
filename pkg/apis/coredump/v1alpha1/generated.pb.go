/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/WanLinghao/fujitsu-coredump/pkg/apis/coredump/v1alpha1/generated.proto

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		github.com/WanLinghao/fujitsu-coredump/pkg/apis/coredump/v1alpha1/generated.proto

	It has these top-level messages:
		CoredumpEndpoint
		CoredumpEndpointDump
		CoredumpEndpointDumpList
		CoredumpEndpointList
		CoredumpEndpointSpec
		CoredumpEndpointStatus
		CoredumpGetOptions
*/
package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_apimachinery_pkg_types "k8s.io/apimachinery/pkg/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *CoredumpEndpoint) Reset()                    { *m = CoredumpEndpoint{} }
func (*CoredumpEndpoint) ProtoMessage()               {}
func (*CoredumpEndpoint) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *CoredumpEndpointDump) Reset()                    { *m = CoredumpEndpointDump{} }
func (*CoredumpEndpointDump) ProtoMessage()               {}
func (*CoredumpEndpointDump) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *CoredumpEndpointDumpList) Reset()      { *m = CoredumpEndpointDumpList{} }
func (*CoredumpEndpointDumpList) ProtoMessage() {}
func (*CoredumpEndpointDumpList) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{2}
}

func (m *CoredumpEndpointList) Reset()                    { *m = CoredumpEndpointList{} }
func (*CoredumpEndpointList) ProtoMessage()               {}
func (*CoredumpEndpointList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func (m *CoredumpEndpointSpec) Reset()                    { *m = CoredumpEndpointSpec{} }
func (*CoredumpEndpointSpec) ProtoMessage()               {}
func (*CoredumpEndpointSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{4} }

func (m *CoredumpEndpointStatus) Reset()                    { *m = CoredumpEndpointStatus{} }
func (*CoredumpEndpointStatus) ProtoMessage()               {}
func (*CoredumpEndpointStatus) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{5} }

func (m *CoredumpGetOptions) Reset()                    { *m = CoredumpGetOptions{} }
func (*CoredumpGetOptions) ProtoMessage()               {}
func (*CoredumpGetOptions) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{6} }

func init() {
	proto.RegisterType((*CoredumpEndpoint)(nil), "github.com.WanLinghao.fujitsu_coredump.pkg.apis.coredump.v1alpha1.CoredumpEndpoint")
	proto.RegisterType((*CoredumpEndpointDump)(nil), "github.com.WanLinghao.fujitsu_coredump.pkg.apis.coredump.v1alpha1.CoredumpEndpointDump")
	proto.RegisterType((*CoredumpEndpointDumpList)(nil), "github.com.WanLinghao.fujitsu_coredump.pkg.apis.coredump.v1alpha1.CoredumpEndpointDumpList")
	proto.RegisterType((*CoredumpEndpointList)(nil), "github.com.WanLinghao.fujitsu_coredump.pkg.apis.coredump.v1alpha1.CoredumpEndpointList")
	proto.RegisterType((*CoredumpEndpointSpec)(nil), "github.com.WanLinghao.fujitsu_coredump.pkg.apis.coredump.v1alpha1.CoredumpEndpointSpec")
	proto.RegisterType((*CoredumpEndpointStatus)(nil), "github.com.WanLinghao.fujitsu_coredump.pkg.apis.coredump.v1alpha1.CoredumpEndpointStatus")
	proto.RegisterType((*CoredumpGetOptions)(nil), "github.com.WanLinghao.fujitsu_coredump.pkg.apis.coredump.v1alpha1.CoredumpGetOptions")
}
func (m *CoredumpEndpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoredumpEndpoint) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
	n2, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Status.Size()))
	n3, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *CoredumpEndpointDump) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoredumpEndpointDump) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n4, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func (m *CoredumpEndpointDumpList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoredumpEndpointDumpList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n5, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *CoredumpEndpointList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoredumpEndpointList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n6, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *CoredumpEndpointSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoredumpEndpointSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.PodUID)))
	i += copy(dAtA[i:], m.PodUID)
	return i, nil
}

func (m *CoredumpEndpointStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoredumpEndpointStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *CoredumpGetOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoredumpGetOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Container)))
	i += copy(dAtA[i:], m.Container)
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CoredumpEndpoint) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *CoredumpEndpointDump) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *CoredumpEndpointDumpList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *CoredumpEndpointList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *CoredumpEndpointSpec) Size() (n int) {
	var l int
	_ = l
	l = len(m.PodUID)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *CoredumpEndpointStatus) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *CoredumpGetOptions) Size() (n int) {
	var l int
	_ = l
	l = len(m.Container)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CoredumpEndpoint) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CoredumpEndpoint{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "CoredumpEndpointSpec", "CoredumpEndpointSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "CoredumpEndpointStatus", "CoredumpEndpointStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CoredumpEndpointDump) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CoredumpEndpointDump{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CoredumpEndpointDumpList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CoredumpEndpointDumpList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "CoredumpEndpointDump", "CoredumpEndpointDump", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CoredumpEndpointList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CoredumpEndpointList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "CoredumpEndpoint", "CoredumpEndpoint", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CoredumpEndpointSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CoredumpEndpointSpec{`,
		`PodUID:` + fmt.Sprintf("%v", this.PodUID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CoredumpEndpointStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CoredumpEndpointStatus{`,
		`}`,
	}, "")
	return s
}
func (this *CoredumpGetOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CoredumpGetOptions{`,
		`Container:` + fmt.Sprintf("%v", this.Container) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CoredumpEndpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CoredumpEndpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoredumpEndpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CoredumpEndpointDump) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CoredumpEndpointDump: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoredumpEndpointDump: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CoredumpEndpointDumpList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CoredumpEndpointDumpList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoredumpEndpointDumpList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, CoredumpEndpointDump{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CoredumpEndpointList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CoredumpEndpointList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoredumpEndpointList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, CoredumpEndpoint{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CoredumpEndpointSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CoredumpEndpointSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoredumpEndpointSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PodUID = k8s_io_apimachinery_pkg_types.UID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CoredumpEndpointStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CoredumpEndpointStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoredumpEndpointStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CoredumpGetOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CoredumpGetOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoredumpGetOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Container", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Container = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/WanLinghao/coredump-detector/pkg/apis/coredump/v1alpha1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xcf, 0x8b, 0xd3, 0x4e,
	0x18, 0xc6, 0x9b, 0xfd, 0x45, 0x3b, 0xdf, 0xaf, 0xb2, 0x06, 0x91, 0xd0, 0x43, 0xba, 0xe6, 0xe4,
	0xa5, 0x13, 0xab, 0x22, 0x1e, 0x35, 0xdb, 0x45, 0x8a, 0x2b, 0xab, 0x59, 0x96, 0x45, 0x11, 0x74,
	0x9a, 0xcc, 0xa6, 0xb3, 0xdd, 0x64, 0x86, 0xcc, 0xa4, 0x6c, 0xc1, 0x83, 0xfe, 0x07, 0xfe, 0x59,
	0x3d, 0xee, 0x71, 0x4f, 0xc5, 0xc6, 0xbb, 0x17, 0xc1, 0x83, 0x27, 0xc9, 0x74, 0xd2, 0xd4, 0x66,
	0x8b, 0x82, 0xd4, 0x5b, 0xdf, 0x37, 0x79, 0x9f, 0xcf, 0xf3, 0xbc, 0x33, 0x0d, 0x78, 0x19, 0x10,
	0xd1, 0x4b, 0xba, 0xd0, 0xa3, 0xa1, 0x7d, 0x8c, 0xa2, 0x7d, 0x12, 0x05, 0x3d, 0x44, 0xed, 0x93,
	0xe4, 0x94, 0x08, 0x9e, 0x34, 0x3d, 0x1a, 0x63, 0x3f, 0x09, 0x99, 0xcd, 0xfa, 0x81, 0x8d, 0x18,
	0xe1, 0xf6, 0xac, 0x33, 0x68, 0xa1, 0x33, 0xd6, 0x43, 0x2d, 0x3b, 0xc0, 0x11, 0x8e, 0x91, 0xc0,
	0x3e, 0x64, 0x31, 0x15, 0x54, 0x7f, 0x52, 0x48, 0xc2, 0x42, 0x12, 0x2a, 0xc9, 0xb7, 0xb9, 0x00,
	0x64, 0xfd, 0x00, 0x66, 0x92, 0x70, 0xd6, 0xc9, 0x25, 0xeb, 0xcd, 0x39, 0x57, 0x01, 0x0d, 0xa8,
	0x2d, 0x95, 0xbb, 0xc9, 0x89, 0xac, 0x64, 0x21, 0x7f, 0x4d, 0x89, 0xf5, 0x07, 0xfd, 0x47, 0x1c,
	0x12, 0x9a, 0x19, 0x0c, 0x91, 0xd7, 0x23, 0x11, 0x8e, 0x87, 0x85, 0xe3, 0x10, 0x0b, 0x64, 0x0f,
	0x4a, 0x3e, 0xeb, 0xf6, 0xb2, 0xa9, 0x38, 0x89, 0x04, 0x09, 0x71, 0x69, 0xe0, 0xe1, 0xef, 0x06,
	0xb8, 0xd7, 0xc3, 0x21, 0x2a, 0xcd, 0x3d, 0xe6, 0x24, 0xe0, 0xb0, 0x18, 0xe6, 0x38, 0x1e, 0xe0,
	0xb8, 0xd9, 0x4d, 0xc8, 0x99, 0x8f, 0xe3, 0xa6, 0xcc, 0x2c, 0x75, 0x54, 0x87, 0x2f, 0x2a, 0x58,
	0xdf, 0xd6, 0xc0, 0xf6, 0xae, 0xda, 0xd2, 0x5e, 0xe4, 0x33, 0x4a, 0x22, 0xa1, 0xbf, 0x03, 0xd5,
	0x2c, 0x9a, 0x8f, 0x04, 0x32, 0xb4, 0x1d, 0xed, 0xce, 0x7f, 0xf7, 0xee, 0x2a, 0x08, 0x9c, 0x77,
	0x58, 0xec, 0x39, 0x7b, 0x1b, 0x0e, 0x5a, 0xf0, 0xa0, 0x7b, 0x8a, 0x3d, 0xf1, 0x1c, 0x0b, 0xe4,
	0xe8, 0xa3, 0x71, 0xa3, 0x92, 0x8e, 0x1b, 0xa0, 0xe8, 0xb9, 0x33, 0x55, 0x7d, 0x08, 0x36, 0x38,
	0xc3, 0x9e, 0xb1, 0x26, 0xd5, 0x8f, 0xe1, 0x5f, 0x1f, 0x2c, 0x5c, 0x0c, 0x71, 0xc8, 0xb0, 0xe7,
	0xfc, 0xaf, 0x4c, 0x6c, 0x64, 0x95, 0x2b, 0x91, 0xfa, 0x47, 0x0d, 0x6c, 0x71, 0x81, 0x44, 0xc2,
	0x8d, 0x75, 0x49, 0x7f, 0xb5, 0x0a, 0xba, 0x04, 0x38, 0xd7, 0x15, 0x7f, 0x6b, 0x5a, 0xbb, 0x0a,
	0x6c, 0x9d, 0x83, 0x9b, 0x8b, 0x13, 0xed, 0x24, 0x64, 0xab, 0x5f, 0xbc, 0xf5, 0x5d, 0x03, 0xc6,
	0x55, 0xe8, 0x7d, 0xc2, 0x85, 0xfe, 0xa6, 0x84, 0x87, 0x7f, 0x86, 0xcf, 0xa6, 0x25, 0x7c, 0x5b,
	0xc1, 0xab, 0x79, 0x67, 0xee, 0xcc, 0xdf, 0x83, 0x4d, 0x22, 0x70, 0xc8, 0x8d, 0xb5, 0x9d, 0xf5,
	0x15, 0x1d, 0x7a, 0x96, 0xc4, 0xb9, 0xa6, 0x3c, 0x6c, 0x76, 0x32, 0x9a, 0x3b, 0x85, 0x5a, 0x5f,
	0xb5, 0xf2, 0xce, 0xff, 0x41, 0xe8, 0xf3, 0x5f, 0x43, 0x1f, 0xae, 0x20, 0xf4, 0x92, 0xc0, 0x5e,
	0x39, 0x6f, 0xf6, 0x2f, 0xd0, 0x9f, 0x81, 0x2d, 0x46, 0xfd, 0x84, 0xf8, 0x32, 0x6d, 0xcd, 0xb9,
	0x9f, 0xdf, 0xd1, 0x17, 0xd4, 0x3f, 0xea, 0xb4, 0x7f, 0x8c, 0x1b, 0xb7, 0x97, 0x7d, 0x8d, 0xc4,
	0x90, 0x61, 0x0e, 0x8f, 0x3a, 0x6d, 0x57, 0x49, 0x58, 0x06, 0xb8, 0x75, 0xf5, 0xd5, 0xb7, 0xf6,
	0x80, 0x9e, 0x3f, 0x79, 0x8a, 0xc5, 0x01, 0x13, 0x84, 0x46, 0x5c, 0xb7, 0x41, 0xcd, 0xa3, 0x91,
	0x40, 0x99, 0xa6, 0xe2, 0xdf, 0x50, 0xfc, 0xda, 0x6e, 0xfe, 0xc0, 0x2d, 0xde, 0x71, 0xe0, 0x68,
	0x62, 0x56, 0x2e, 0x26, 0x66, 0xe5, 0x72, 0x62, 0x56, 0x3e, 0xa4, 0xa6, 0x36, 0x4a, 0x4d, 0xed,
	0x22, 0x35, 0xb5, 0xcb, 0xd4, 0xd4, 0x3e, 0xa7, 0xa6, 0xf6, 0xe9, 0x8b, 0x59, 0x79, 0x5d, 0xcd,
	0x97, 0xf3, 0x33, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xcb, 0x2a, 0x4f, 0x76, 0x06, 0x00, 0x00,
}
