// Copyright (c) 2008-2018, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"fmt"
	"unicode/utf8"

	"github.com/hazelcast/hazelcast-go-client/core"
	"github.com/hazelcast/hazelcast-go-client/internal/proto/bufutil"
	"github.com/hazelcast/hazelcast-go-client/serialization"
	"github.com/hazelcast/hazelcast-go-client/serialization/internal/bufferutil"
)

type ObjectDataOutput struct {
	buffer    []byte
	service   *Service
	bigEndian bool
	position  int32
}

func NewObjectDataOutput(length int, service *Service, bigEndian bool) *ObjectDataOutput {
	return &ObjectDataOutput{make([]byte, length), service, bigEndian, 0}
}

func (o *ObjectDataOutput) Available() int {
	if o.buffer == nil {
		return 0
	}
	return len(o.buffer) - int(o.position)
}

func (o *ObjectDataOutput) Position() int32 {
	return o.position
}

func (o *ObjectDataOutput) SetPosition(pos int32) {
	o.position = pos
}

func (o *ObjectDataOutput) ToBuffer() []byte {
	if o.position == 0 {
		return make([]byte, 0)
	}
	snapBuffer := make([]byte, o.position)
	copy(snapBuffer, o.buffer)
	return snapBuffer
}

func (o *ObjectDataOutput) WriteZeroBytes(count int) {
	for i := 0; i < count; i++ {
		o.WriteByte(0)
	}
}

func (o *ObjectDataOutput) EnsureAvailable(size int) {
	if o.Available() < size {
		temp := make([]byte, int(o.position)+size)
		copy(temp, o.buffer)
		o.buffer = temp
	}
}

func (o *ObjectDataOutput) WriteByte(v byte) {
	o.EnsureAvailable(bufutil.ByteSizeInBytes)
	bufferutil.WriteUInt8(o.buffer, o.position, v)
	o.position += bufutil.ByteSizeInBytes
}

func (o *ObjectDataOutput) WriteBool(v bool) {
	o.EnsureAvailable(bufutil.BoolSizeInBytes)
	bufferutil.WriteBool(o.buffer, o.position, v)
	o.position += bufutil.BoolSizeInBytes
}

func (o *ObjectDataOutput) WriteUInt16(v uint16) {
	o.EnsureAvailable(bufutil.Uint16SizeInBytes)
	bufferutil.WriteUInt16(o.buffer, o.position, v, o.bigEndian)
	o.position += bufutil.Uint16SizeInBytes
}

func (o *ObjectDataOutput) WriteInt16(v int16) {
	o.EnsureAvailable(bufutil.Int16SizeInBytes)
	bufferutil.WriteInt16(o.buffer, o.position, v, o.bigEndian)
	o.position += bufutil.Int16SizeInBytes
}

func (o *ObjectDataOutput) WriteInt32(v int32) {
	o.EnsureAvailable(bufutil.Int32SizeInBytes)
	bufferutil.WriteInt32(o.buffer, o.position, v, o.bigEndian)
	o.position += bufutil.Int32SizeInBytes
}

func (o *ObjectDataOutput) WriteInt64(v int64) {
	o.EnsureAvailable(bufutil.Int64SizeInBytes)
	bufferutil.WriteInt64(o.buffer, o.position, v, o.bigEndian)
	o.position += bufutil.Int64SizeInBytes
}

func (o *ObjectDataOutput) WriteFloat32(v float32) {
	o.EnsureAvailable(bufutil.Float32SizeInBytes)
	bufferutil.WriteFloat32(o.buffer, o.position, v, o.bigEndian)
	o.position += bufutil.Float32SizeInBytes
}

func (o *ObjectDataOutput) WriteFloat64(v float64) {
	o.EnsureAvailable(bufutil.Float64SizeInBytes)
	bufferutil.WriteFloat64(o.buffer, o.position, v, o.bigEndian)
	o.position += bufutil.Float64SizeInBytes
}

func (o *ObjectDataOutput) WriteUTF(v string) {
	length := int32(utf8.RuneCountInString(v))
	o.WriteInt32(length)

	if length > 0 {
		o.EnsureAvailable(len(v))
		for _, s := range v {
			o.position += int32(utf8.EncodeRune(o.buffer[o.position:], s))
		}
	}
}

func (o *ObjectDataOutput) WriteObject(object interface{}) error {
	return o.service.WriteObject(o, object)

}

func (o *ObjectDataOutput) WriteByteArray(v []byte) {
	var length int32
	if v != nil {
		length = int32(len(v))
	} else {
		length = bufutil.NilArrayLength
	}
	o.WriteInt32(length)
	for j := int32(0); j < length; j++ {
		o.WriteByte(v[j])
	}
}

func (o *ObjectDataOutput) WriteBoolArray(v []bool) {
	var length int32
	if v != nil {
		length = int32(len(v))
	} else {
		length = bufutil.NilArrayLength
	}
	o.WriteInt32(length)
	for j := int32(0); j < length; j++ {
		o.WriteBool(v[j])
	}
}

func (o *ObjectDataOutput) WriteUInt16Array(v []uint16) {
	var length int32
	if v != nil {
		length = int32(len(v))
	} else {
		length = bufutil.NilArrayLength
	}
	o.WriteInt32(length)
	for j := int32(0); j < length; j++ {
		o.WriteUInt16(v[j])
	}
}

func (o *ObjectDataOutput) WriteInt16Array(v []int16) {
	var length int32
	if v != nil {
		length = int32(len(v))
	} else {
		length = bufutil.NilArrayLength
	}
	o.WriteInt32(length)
	for j := int32(0); j < length; j++ {
		o.WriteInt16(v[j])
	}
}

func (o *ObjectDataOutput) WriteInt32Array(v []int32) {
	var length int32
	if v != nil {
		length = int32(len(v))
	} else {
		length = bufutil.NilArrayLength
	}
	o.WriteInt32(length)
	for j := int32(0); j < length; j++ {
		o.WriteInt32(v[j])
	}
}

func (o *ObjectDataOutput) WriteInt64Array(v []int64) {
	var length int32
	if v != nil {
		length = int32(len(v))
	} else {
		length = bufutil.NilArrayLength
	}
	o.WriteInt32(length)
	for j := int32(0); j < length; j++ {
		o.WriteInt64(v[j])
	}
}

func (o *ObjectDataOutput) WriteFloat32Array(v []float32) {
	var length int32
	if v != nil {
		length = int32(len(v))
	} else {
		length = bufutil.NilArrayLength
	}
	o.WriteInt32(length)
	for j := int32(0); j < length; j++ {
		o.WriteFloat32(v[j])
	}
}

func (o *ObjectDataOutput) WriteFloat64Array(v []float64) {
	var length int32
	if v != nil {
		length = int32(len(v))
	} else {
		length = bufutil.NilArrayLength
	}
	o.WriteInt32(length)
	for j := int32(0); j < length; j++ {
		o.WriteFloat64(v[j])
	}
}

func (o *ObjectDataOutput) WriteUTFArray(v []string) {
	var length int32
	if v != nil {
		length = int32(len(v))
	} else {
		length = bufutil.NilArrayLength
	}
	o.WriteInt32(length)
	for j := int32(0); j < length; j++ {
		o.WriteUTF(v[j])
	}
}

func (o *ObjectDataOutput) WriteBytes(v string) {
	for _, char := range v {
		o.WriteByte(uint8(char))
	}
}

func (o *ObjectDataOutput) WriteData(data serialization.Data) {
	var length int32
	if data == nil {
		length = bufutil.NilArrayLength
	} else {
		length = int32(data.TotalSize())
	}
	o.WriteInt32(length)
	if length > 0 {
		o.EnsureAvailable(int(length))
		copy(o.buffer[o.position:], data.Buffer())
		o.position += length
	}
}

//// ObjectDataInput ////

type ObjectDataInput struct {
	buffer    []byte
	offset    int32
	service   *Service
	bigEndian bool
	position  int32
}

func NewObjectDataInput(buffer []byte, offset int32, service *Service, bigEndian bool) *ObjectDataInput {
	return &ObjectDataInput{buffer, offset, service, bigEndian, offset}
}

func (i *ObjectDataInput) Available() int32 {
	return int32(len(i.buffer)) - i.position
}

func (i *ObjectDataInput) AssertAvailable(k int) error {
	if i.position < 0 {
		return core.NewHazelcastIllegalArgumentError(fmt.Sprintf("negative pos -> %v", i.position), nil)
	}
	if len(i.buffer) < int(i.position)+k {
		return core.NewHazelcastEOFError(fmt.Sprintf("cannot read %v bytes", k), nil)
	}
	return nil
}

func (i *ObjectDataInput) Position() int32 {
	return i.position
}

func (i *ObjectDataInput) SetPosition(pos int32) {
	i.position = pos
}

func (i *ObjectDataInput) ReadByte() (byte, error) {
	var err = i.AssertAvailable(bufutil.ByteSizeInBytes)
	var ret byte
	if err == nil {
		ret = bufferutil.ReadUInt8(i.buffer, i.position)
		i.position += bufutil.ByteSizeInBytes
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadByteWithPosition(pos int32) (byte, error) {
	var err = i.AssertAvailable(bufutil.ByteSizeInBytes)
	var ret byte
	if err == nil {
		ret = bufferutil.ReadUInt8(i.buffer, pos)
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadBool() (bool, error) {
	var err = i.AssertAvailable(bufutil.BoolSizeInBytes)
	var ret bool
	if err == nil {
		ret = bufferutil.ReadBool(i.buffer, i.position)
		i.position += bufutil.BoolSizeInBytes
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadBoolWithPosition(pos int32) (bool, error) {
	var err = i.AssertAvailable(bufutil.BoolSizeInBytes)
	var ret bool
	if err == nil {
		ret = bufferutil.ReadBool(i.buffer, pos)
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadUInt16() (uint16, error) {
	var err = i.AssertAvailable(bufutil.Uint16SizeInBytes)
	var ret uint16
	if err == nil {
		ret = bufferutil.ReadUInt16(i.buffer, i.position, i.bigEndian)
		i.position += bufutil.Uint16SizeInBytes
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadUInt16WithPosition(pos int32) (uint16, error) {
	var err = i.AssertAvailable(bufutil.Uint16SizeInBytes)
	var ret uint16
	if err == nil {
		ret = bufferutil.ReadUInt16(i.buffer, pos, i.bigEndian)
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadInt16() (int16, error) {
	var err = i.AssertAvailable(bufutil.Int16SizeInBytes)
	var ret int16
	if err == nil {
		ret = bufferutil.ReadInt16(i.buffer, i.position, i.bigEndian)
		i.position += bufutil.Int16SizeInBytes
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadInt16WithPosition(pos int32) (int16, error) {
	var err = i.AssertAvailable(bufutil.Int16SizeInBytes)
	var ret int16
	if err == nil {
		ret = bufferutil.ReadInt16(i.buffer, pos, i.bigEndian)
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadInt32() (int32, error) {
	var err = i.AssertAvailable(bufutil.Int32SizeInBytes)
	var ret int32
	if err == nil {
		ret = bufferutil.ReadInt32(i.buffer, i.position, i.bigEndian)
		i.position += bufutil.Int32SizeInBytes
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadInt32WithPosition(pos int32) (int32, error) {
	var err = i.AssertAvailable(bufutil.Int32SizeInBytes)
	var ret int32
	if err == nil {
		ret = bufferutil.ReadInt32(i.buffer, pos, i.bigEndian)
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadInt64() (int64, error) {
	var err = i.AssertAvailable(bufutil.Int64SizeInBytes)
	var ret int64
	if err == nil {
		ret = bufferutil.ReadInt64(i.buffer, i.position, i.bigEndian)
		i.position += bufutil.Int64SizeInBytes
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadInt64WithPosition(pos int32) (int64, error) {
	var err = i.AssertAvailable(bufutil.Int64SizeInBytes)
	var ret int64
	if err == nil {
		ret = bufferutil.ReadInt64(i.buffer, pos, i.bigEndian)
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadFloat32() (float32, error) {
	var err = i.AssertAvailable(bufutil.Float32SizeInBytes)
	var ret float32
	if err == nil {
		ret = bufferutil.ReadFloat32(i.buffer, i.position, i.bigEndian)
		i.position += bufutil.Float32SizeInBytes
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadFloat32WithPosition(pos int32) (float32, error) {
	var err = i.AssertAvailable(bufutil.Float32SizeInBytes)
	var ret float32
	if err == nil {
		ret = bufferutil.ReadFloat32(i.buffer, pos, i.bigEndian)
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadFloat64() (float64, error) {
	var err = i.AssertAvailable(bufutil.Float64SizeInBytes)
	var ret float64
	if err == nil {
		ret = bufferutil.ReadFloat64(i.buffer, i.position, i.bigEndian)
		i.position += bufutil.Float64SizeInBytes
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadFloat64WithPosition(pos int32) (float64, error) {
	var err = i.AssertAvailable(bufutil.Float64SizeInBytes)
	var ret float64
	if err == nil {
		ret = bufferutil.ReadFloat64(i.buffer, pos, i.bigEndian)
		return ret, err
	}
	return ret, err
}

func (i *ObjectDataInput) ReadUTF() (string, error) {
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return "", err
	}
	var ret = make([]rune, length)
	for j := 0; j < int(length); j++ {
		r, n := utf8.DecodeRune(i.buffer[i.position:])
		i.position += int32(n)
		ret[j] = r
	}
	return string(ret), nil
}

func (i *ObjectDataInput) ReadUTFWithPosition(pos int32) (string, error) {
	length, err := i.ReadInt32WithPosition(pos)
	if err != nil || length == bufutil.NilArrayLength {
		return "", err
	}
	pos += bufutil.Int32SizeInBytes
	var ret = make([]rune, length)
	for j := 0; j < int(length); j++ {
		r, n := utf8.DecodeRune(i.buffer[pos:])
		pos += int32(n)
		ret[j] = r
	}
	return string(ret), nil
}

func (i *ObjectDataInput) ReadObject() (interface{}, error) {
	return i.service.ReadObject(i)
}

func (i *ObjectDataInput) ReadByteArray() ([]byte, error) {
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]byte, length)
	for j := int32(0); j < length; j++ {
		arr[j], err = i.ReadByte()
		if err != nil {
			return nil, err
		}
	}
	return arr, nil
}

func (i *ObjectDataInput) ReadByteArrayWithPosition(pos int32) ([]byte, error) {
	backupPos := i.position
	i.position = pos
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]byte, length)
	for j := int32(0); j < length; j++ {
		arr[j], err = i.ReadByte()
		if err != nil {
			return nil, err
		}
	}
	i.position = backupPos
	return arr, nil
}

func (i *ObjectDataInput) ReadBoolArray() ([]bool, error) {
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]bool, length)
	for j := int32(0); j < length; j++ {
		arr[j], err = i.ReadBool()
		if err != nil {
			return nil, err
		}
	}
	return arr, nil
}

func (i *ObjectDataInput) ReadBoolArrayWithPosition(pos int32) ([]bool, error) {
	backupPos := i.position
	i.position = pos
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]bool, length)
	for j := int32(0); j < length; j++ {
		arr[j], err = i.ReadBool()
		if err != nil {
			return nil, err
		}
	}
	i.position = backupPos
	return arr, nil
}

func (i *ObjectDataInput) ReadUInt16Array() ([]uint16, error) {
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]uint16, length)
	for j := int32(0); j < length; j++ {
		arr[j], err = i.ReadUInt16()
		if err != nil {
			return nil, err
		}
	}
	return arr, nil
}

func (i *ObjectDataInput) ReadUInt16ArrayWithPosition(pos int32) ([]uint16, error) {
	backupPos := i.position
	i.position = pos
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]uint16, length)
	for j := int32(0); j < length; j++ {
		arr[j], err = i.ReadUInt16()
		if err != nil {
			return nil, err
		}
	}
	i.position = backupPos
	return arr, nil
}

func (i *ObjectDataInput) ReadInt16Array() ([]int16, error) {
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]int16, length)
	for j := int32(0); j < length; j++ {
		arr[j], err = i.ReadInt16()
		if err != nil {
			return nil, err
		}
	}
	return arr, nil
}

func (i *ObjectDataInput) ReadInt16ArrayWithPosition(pos int32) ([]int16, error) {
	backupPos := i.position
	i.position = pos
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]int16, length)
	for j := int32(0); j < length; j++ {
		arr[j], err = i.ReadInt16()
		if err != nil {
			return nil, err
		}
	}
	i.position = backupPos
	return arr, nil
}

func (i *ObjectDataInput) ReadInt32Array() ([]int32, error) {
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]int32, length)
	for j := int32(0); j < length; j++ {
		arr[j], err = i.ReadInt32()
		if err != nil {
			return nil, err
		}
	}
	return arr, nil
}

func (i *ObjectDataInput) ReadInt32ArrayWithPosition(pos int32) ([]int32, error) {
	backupPos := i.position
	i.position = pos
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]int32, length)
	for j := int32(0); j < length; j++ {
		arr[j], err = i.ReadInt32()
		if err != nil {
			return nil, err
		}
	}
	i.position = backupPos
	return arr, nil
}

func (i *ObjectDataInput) ReadInt64Array() ([]int64, error) {
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]int64, length)
	for j := int32(0); j < length; j++ {
		arr[j], _ = i.ReadInt64()
		if err != nil {
			return nil, err
		}
	}
	return arr, nil
}

func (i *ObjectDataInput) ReadInt64ArrayWithPosition(pos int32) ([]int64, error) {
	backupPos := i.position
	i.position = pos
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]int64, length)
	for j := int32(0); j < length; j++ {
		arr[j], _ = i.ReadInt64()
		if err != nil {
			return nil, err
		}
	}
	i.position = backupPos
	return arr, nil
}

func (i *ObjectDataInput) ReadFloat32Array() ([]float32, error) {
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]float32, length)
	for j := int32(0); j < length; j++ {
		arr[j], _ = i.ReadFloat32()
		if err != nil {
			return nil, err
		}
	}
	return arr, nil
}

func (i *ObjectDataInput) ReadFloat32ArrayWithPosition(pos int32) ([]float32, error) {
	backupPos := i.position
	i.position = pos
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]float32, length)
	for j := int32(0); j < length; j++ {
		arr[j], _ = i.ReadFloat32()
		if err != nil {
			return nil, err
		}
	}
	i.position = backupPos
	return arr, nil
}

func (i *ObjectDataInput) ReadFloat64Array() ([]float64, error) {
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]float64, length)
	for j := int32(0); j < length; j++ {
		arr[j], _ = i.ReadFloat64()
		if err != nil {
			return nil, err
		}
	}
	return arr, nil
}

func (i *ObjectDataInput) ReadFloat64ArrayWithPosition(pos int32) ([]float64, error) {
	backupPos := i.position
	i.position = pos
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]float64, length)
	for j := int32(0); j < length; j++ {
		arr[j], _ = i.ReadFloat64()
		if err != nil {
			return nil, err
		}
	}
	i.position = backupPos
	return arr, nil
}

func (i *ObjectDataInput) ReadUTFArray() ([]string, error) {
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]string, length)
	for j := int32(0); j < length; j++ {
		arr[j], err = i.ReadUTF()
		if err != nil {
			return nil, err
		}
	}
	return arr, nil
}

func (i *ObjectDataInput) ReadUTFArrayWithPosition(pos int32) ([]string, error) {
	backupPos := i.position
	i.position = pos
	length, err := i.ReadInt32()
	if err != nil || length == bufutil.NilArrayLength {
		return nil, err
	}
	var arr = make([]string, length)
	for j := int32(0); j < length; j++ {
		arr[j], err = i.ReadUTF()
		if err != nil {
			return nil, err
		}
	}
	i.position = backupPos
	return arr, nil
}

func (i *ObjectDataInput) ReadData() (serialization.Data, error) {
	array, err := i.ReadByteArray()
	if err != nil {
		return nil, err
	}
	if array == nil {
		return nil, nil
	}
	return &Data{array}, nil
}

type PositionalObjectDataOutput struct {
	*ObjectDataOutput
}

func NewPositionalObjectDataOutput(length int, service *Service, bigEndian bool) *PositionalObjectDataOutput {
	return &PositionalObjectDataOutput{NewObjectDataOutput(length, service, bigEndian)}
}

func (p *PositionalObjectDataOutput) PWriteByte(pos int32, v byte) {
	bufferutil.WriteUInt8(p.buffer, pos, v)
}

func (p *PositionalObjectDataOutput) PWriteBool(pos int32, v bool) {
	bufferutil.WriteBool(p.buffer, pos, v)
}

func (p *PositionalObjectDataOutput) PWriteUInt16(pos int32, v uint16) {
	bufferutil.WriteUInt16(p.buffer, pos, v, p.bigEndian)
}

func (p *PositionalObjectDataOutput) PWriteInt16(pos int32, v int16) {
	bufferutil.WriteInt16(p.buffer, pos, v, p.bigEndian)
}

func (p *PositionalObjectDataOutput) PWriteInt32(pos int32, v int32) {
	bufferutil.WriteInt32(p.buffer, pos, v, p.bigEndian)
}

func (p *PositionalObjectDataOutput) PWriteInt64(pos int32, v int64) {
	bufferutil.WriteInt64(p.buffer, pos, v, p.bigEndian)
}

func (p *PositionalObjectDataOutput) PWriteFloat32(pos int32, v float32) {
	bufferutil.WriteFloat32(p.buffer, pos, v, p.bigEndian)
}

func (p *PositionalObjectDataOutput) PWriteFloat64(pos int32, v float64) {
	bufferutil.WriteFloat64(p.buffer, pos, v, p.bigEndian)
}
