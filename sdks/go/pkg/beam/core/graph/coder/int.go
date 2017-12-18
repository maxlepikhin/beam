// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package coder

import (
	"encoding/binary"
	"io"

	"github.com/apache/beam/sdks/go/pkg/beam/core/util/ioutilx"
	"github.com/apache/beam/sdks/go/pkg/beam/core/util/reflectx"
)

var (
	// Fixed-sized custom coders for integers.

	Uint32 *CustomCoder
	Int32  *CustomCoder
	Uint64 *CustomCoder
	Int64  *CustomCoder
)

func init() {
	var err error
	Uint32, err = NewCustomCoder("uint32", reflectx.Uint32, encUint32, decUint32)
	if err != nil {
		panic(err)
	}
	Int32, err = NewCustomCoder("int32", reflectx.Int32, encInt32, decInt32)
	if err != nil {
		panic(err)
	}
	Uint64, err = NewCustomCoder("uint64", reflectx.Uint64, encUint64, decUint64)
	if err != nil {
		panic(err)
	}
	Int64, err = NewCustomCoder("int64", reflectx.Int64, encInt64, decInt64)
	if err != nil {
		panic(err)
	}
}

func encUint32(v uint32) []byte {
	ret := make([]byte, 4)
	binary.BigEndian.PutUint32(ret, v)
	return ret
}

func decUint32(data []byte) uint32 {
	return binary.BigEndian.Uint32(data)
}

func encInt32(v int32) []byte {
	return encUint32(uint32(v))
}

func decInt32(data []byte) int32 {
	return int32(decUint32(data))
}

func encUint64(v uint64) []byte {
	ret := make([]byte, 8)
	binary.BigEndian.PutUint64(ret, v)
	return ret
}

func decUint64(data []byte) uint64 {
	return binary.BigEndian.Uint64(data)
}

func encInt64(v int64) []byte {
	return encUint64(uint64(v))
}

func decInt64(data []byte) int64 {
	return int64(decUint64(data))
}

// EncodeUint64 encodes an uint64 in big endian format.
func EncodeUint64(value uint64, w io.Writer) error {
	ret := make([]byte, 8)
	binary.BigEndian.PutUint64(ret, value)
	_, err := w.Write(ret)
	return err
}

// DecodeUint64 decodes an uint64 in big endian format.
func DecodeUint64(r io.Reader) (uint64, error) {
	data, err := ioutilx.ReadN(r, 8)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(data), nil
}

// EncodeUint32 encodes an uint32 in big endian format.
func EncodeUint32(value uint32, w io.Writer) error {
	ret := make([]byte, 4)
	binary.BigEndian.PutUint32(ret, value)
	_, err := w.Write(ret)
	return err
}

// DecodeUint32 decodes an uint32 in big endian format.
func DecodeUint32(r io.Reader) (uint32, error) {
	data, err := ioutilx.ReadN(r, 4)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(data), nil
}

// EncodeInt32 encodes an int32 in big endian format.
func EncodeInt32(value int32, w io.Writer) error {
	return EncodeUint32(uint32(value), w)
}

// DecodeInt32 decodes an int32 in big endian format.
func DecodeInt32(r io.Reader) (int32, error) {
	ret, err := DecodeUint32(r)
	if err != nil {
		return 0, err
	}
	return int32(ret), nil
}
