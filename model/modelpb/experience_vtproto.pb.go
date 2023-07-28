// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.4.0
// source: experience.proto

package modelpb

import (
	binary "encoding/binary"
	fmt "fmt"
	io "io"
	math "math"
	sync "sync"

	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *UserExperience) CloneVT() *UserExperience {
	if m == nil {
		return (*UserExperience)(nil)
	}
	r := &UserExperience{
		CumulativeLayoutShift: m.CumulativeLayoutShift,
		FirstInputDelay:       m.FirstInputDelay,
		TotalBlockingTime:     m.TotalBlockingTime,
		LongTask:              m.LongTask.CloneVT(),
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *UserExperience) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *LongtaskMetrics) CloneVT() *LongtaskMetrics {
	if m == nil {
		return (*LongtaskMetrics)(nil)
	}
	r := &LongtaskMetrics{
		Count: m.Count,
		Sum:   m.Sum,
		Max:   m.Max,
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *LongtaskMetrics) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *UserExperience) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserExperience) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *UserExperience) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.LongTask != nil {
		size, err := m.LongTask.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if m.TotalBlockingTime != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.TotalBlockingTime))))
		i--
		dAtA[i] = 0x19
	}
	if m.FirstInputDelay != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.FirstInputDelay))))
		i--
		dAtA[i] = 0x11
	}
	if m.CumulativeLayoutShift != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.CumulativeLayoutShift))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func (m *LongtaskMetrics) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LongtaskMetrics) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *LongtaskMetrics) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.Max != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Max))))
		i--
		dAtA[i] = 0x19
	}
	if m.Sum != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Sum))))
		i--
		dAtA[i] = 0x11
	}
	if m.Count != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

var vtprotoPool_UserExperience = sync.Pool{
	New: func() interface{} {
		return &UserExperience{}
	},
}

func (m *UserExperience) ResetVT() {
	m.LongTask.ReturnToVTPool()
	m.Reset()
}
func (m *UserExperience) ReturnToVTPool() {
	if m != nil {
		m.ResetVT()
		vtprotoPool_UserExperience.Put(m)
	}
}
func UserExperienceFromVTPool() *UserExperience {
	return vtprotoPool_UserExperience.Get().(*UserExperience)
}

var vtprotoPool_LongtaskMetrics = sync.Pool{
	New: func() interface{} {
		return &LongtaskMetrics{}
	},
}

func (m *LongtaskMetrics) ResetVT() {
	m.Reset()
}
func (m *LongtaskMetrics) ReturnToVTPool() {
	if m != nil {
		m.ResetVT()
		vtprotoPool_LongtaskMetrics.Put(m)
	}
}
func LongtaskMetricsFromVTPool() *LongtaskMetrics {
	return vtprotoPool_LongtaskMetrics.Get().(*LongtaskMetrics)
}
func (m *UserExperience) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CumulativeLayoutShift != 0 {
		n += 9
	}
	if m.FirstInputDelay != 0 {
		n += 9
	}
	if m.TotalBlockingTime != 0 {
		n += 9
	}
	if m.LongTask != nil {
		l = m.LongTask.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *LongtaskMetrics) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Count != 0 {
		n += 1 + sov(uint64(m.Count))
	}
	if m.Sum != 0 {
		n += 9
	}
	if m.Max != 0 {
		n += 9
	}
	n += len(m.unknownFields)
	return n
}

func (m *UserExperience) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserExperience: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserExperience: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumulativeLayoutShift", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.CumulativeLayoutShift = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstInputDelay", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.FirstInputDelay = float64(math.Float64frombits(v))
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalBlockingTime", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.TotalBlockingTime = float64(math.Float64frombits(v))
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LongTask", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LongTask == nil {
				m.LongTask = LongtaskMetricsFromVTPool()
			}
			if err := m.LongTask.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LongtaskMetrics) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LongtaskMetrics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LongtaskMetrics: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Sum = float64(math.Float64frombits(v))
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Max", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Max = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
