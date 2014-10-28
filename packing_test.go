// Copyright 2013-2014 Aerospike, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
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

package aerospike

import (
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// . "github.com/aerospike/aerospike-client-go/logger"
	// . "github.com/aerospike/aerospike-client-go/types"
	// ParticleType "github.com/aerospike/aerospike-client-go/types/particle_type"
	// . "github.com/aerospike/aerospike-client-go/utils/buffer"
)

func testPackingFor(v interface{}) interface{} {
	packer := newPacker()

	err := packer.PackObject(v)
	Expect(err).ToNot(HaveOccurred())

	unpacker := newUnpacker(packer.buffer.Bytes(), 0, len(packer.buffer.Bytes()))
	unpackedValue, err := unpacker.unpackObject()

	return unpackedValue
}

var _ = Describe("Packing Test", func() {

	Context("Simple Value Types", func() {

		It("should pack and unpack nil values", func() {
			Expect(testPackingFor(nil)).To(BeNil())
		})

		It("should pack and unpack int8 values", func() {
			v := int8(math.MaxInt8)
			Expect(testPackingFor(v)).To(Equal(int(v)))

			v = int8(math.MinInt8)
			Expect(testPackingFor(v)).To(Equal(int(v)))
		})

		It("should pack and unpack uint8 values", func() {
			v := uint8(math.MaxUint8)
			Expect(testPackingFor(v)).To(Equal(int(v)))
		})

		It("should pack and unpack int16 values", func() {
			v := int16(math.MaxInt16)
			Expect(testPackingFor(v)).To(Equal(int(v)))

			v = int16(math.MinInt16)
			Expect(testPackingFor(v)).To(Equal(int(v)))
		})

		It("should pack and unpack uint16 values", func() {
			v := uint16(math.MaxUint16)
			Expect(testPackingFor(v)).To(Equal(int(v)))
		})

		It("should pack and unpack int32 values", func() {
			v := int32(math.MaxInt32)
			Expect(testPackingFor(v)).To(Equal(int(v)))

			v = int32(math.MinInt32)
			Expect(testPackingFor(v)).To(Equal(int(v)))
		})

		It("should pack and unpack uint32 values", func() {
			v := uint32(math.MaxUint32)
			Expect(testPackingFor(v)).To(Equal(int(v)))
		})

		It("should pack and unpack int64 values", func() {
			v := int64(math.MaxInt64)
			Expect(testPackingFor(v)).To(Equal(v))

			v = int64(math.MinInt64)
			Expect(testPackingFor(v)).To(Equal(v))
		})

		It("should pack and unpack uint64 values", func() {
			v := uint64(math.MaxUint64)
			Expect(testPackingFor(v)).To(Equal(v))
		})

		It("should pack and unpack string values", func() {
			v := "string123456789\n"
			Expect(testPackingFor(v)).To(Equal(v))
		})

		// It("should pack and unpack boolean: true values", func() {
		// 	v := true
		// 	Expect(testPackingFor(v)).To(Equal(v))
		// })

		// It("should pack and unpack boolean: false values", func() {
		// 	v := false
		// 	Expect(testPackingFor(v)).To(Equal(v))
		// })
	})

	Context("Array Value Types", func() {

		It("should pack and unpack empty array of int8", func() {
			v := []int8{}
			Expect(testPackingFor(v)).To(Equal([]interface{}{}))
		})

		It("should pack and unpack an array of int8", func() {
			v := []int8{1, 2, 3}
			Expect(testPackingFor(v)).To(Equal([]interface{}{1, 2, 3}))
		})

		It("should pack and unpack empty array of uint8", func() {
			// Note: An array of uint8 ends up as being a ByteArrayValue
			v := []uint8{}
			Expect(testPackingFor(v)).To(Equal([]byte{}))
		})

		It("should pack and unpack an array of uint8", func() {
			// Note: An array of uint8 ends up as being a ByteArrayValue
			v := []uint8{1, 2, 3}
			Expect(testPackingFor(v)).To(Equal([]byte{1, 2, 3}))
		})

		It("should pack and unpack empty array of int16", func() {
			v := []int16{}
			Expect(testPackingFor(v)).To(Equal([]interface{}{}))
		})

		It("should pack and unpack an array of int16", func() {
			v := []int16{1, 2, 3}
			Expect(testPackingFor(v)).To(Equal([]interface{}{1, 2, 3}))
		})

		It("should pack and unpack empty array of uint16", func() {
			v := []uint16{}
			Expect(testPackingFor(v)).To(Equal([]interface{}{}))
		})

		It("should pack and unpack an array of uint16", func() {
			v := []uint16{1, 2, 3}
			Expect(testPackingFor(v)).To(Equal([]interface{}{1, 2, 3}))
		})

		It("should pack and unpack empty array of int32", func() {
			v := []int32{}
			Expect(testPackingFor(v)).To(Equal([]interface{}{}))
		})

		It("should pack and unpack an array of int32", func() {
			v := []int32{1, 2, 3}
			Expect(testPackingFor(v)).To(Equal([]interface{}{1, 2, 3}))
		})

		It("should pack and unpack empty array of uint32", func() {
			v := []uint32{}
			Expect(testPackingFor(v)).To(Equal([]interface{}{}))
		})

		It("should pack and unpack an array of uint32", func() {
			v := []uint32{1, 2, 3}
			Expect(testPackingFor(v)).To(Equal([]interface{}{1, 2, 3}))
		})

		It("should pack and unpack empty array of int64", func() {
			v := []int64{}
			Expect(testPackingFor(v)).To(Equal([]interface{}{}))
		})

		It("should pack and unpack an array of int64", func() {
			v := []int64{1, 2, 3}
			Expect(testPackingFor(v)).To(Equal([]interface{}{1, 2, 3}))
		})

		It("should pack and unpack empty array of uint64", func() {
			v := []uint64{}
			Expect(testPackingFor(v)).To(Equal([]interface{}{}))
		})

		It("should pack and unpack an array of uint64", func() {
			v := []uint64{1, 2, 3}
			Expect(testPackingFor(v)).To(Equal([]interface{}{uint64(1), uint64(2), uint64(3)}))
		})

		It("should pack and unpack empty array of string", func() {
			v := []string{}
			Expect(testPackingFor(v)).To(Equal([]interface{}{}))
		})

		It("should pack and unpack an array of string", func() {
			v := []string{"this", "is", "an", "array", "of", "strings"}
			Expect(testPackingFor(v)).To(Equal([]interface{}{"this", "is", "an", "array", "of", "strings"}))
		})

	})

	Context("Map Value Types", func() {

		It("should pack and unpack empty map", func() {
			v := map[interface{}]interface{}{}
			Expect(testPackingFor(v)).To(Equal(map[interface{}]interface{}{}))
		})

		It("should pack and unpack a complex map", func() {
			v := map[interface{}]interface{}{
				"uint8":  uint8(math.MaxUint8),
				"int8":   int8(math.MaxInt8),
				"mint8":  int8(math.MinInt8),
				"uint16": uint16(math.MaxUint16),
				"int16":  int16(math.MaxInt16),
				"mint16": int16(math.MinInt16),
				"uint32": uint32(math.MaxUint32),
				"int32":  int32(math.MaxInt32),
				"mint32": int32(math.MinInt32),
				"uint64": uint64(math.MaxUint64),
				"int64":  int64(math.MaxInt64),
				"mint64": int64(math.MinInt64),
				"str":    "this is a string",
				"nil":    nil,
			}

			v_res := map[interface{}]interface{}{
				"uint8":  int(math.MaxUint8),
				"int8":   int(math.MaxInt8),
				"mint8":  int(math.MinInt8),
				"uint16": int(math.MaxUint16),
				"int16":  int(math.MaxInt16),
				"mint16": int(math.MinInt16),
				"uint32": int(math.MaxUint32),
				"int32":  int(math.MaxInt32),
				"mint32": int(math.MinInt32),
				"uint64": uint64(math.MaxUint64),
				"int64":  int64(math.MaxInt64),
				"mint64": int64(math.MinInt64),
				"str":    "this is a string",
				"nil":    nil,
			}

			Expect(testPackingFor(v)).To(Equal(v_res))
		})

	})
})
