// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pcommon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewByteSlice(t *testing.T) {
	ms := NewByteSlice()
	assert.Equal(t, 0, ms.Len())
	ms.FromRaw([]byte{1, 2, 3})
	assert.Equal(t, 3, ms.Len())
	assert.Equal(t, []byte{1, 2, 3}, ms.AsRaw())
	ms.SetAt(1, byte(5))
	assert.Equal(t, []byte{1, 5, 3}, ms.AsRaw())
	ms.FromRaw([]byte{3})
	assert.Equal(t, 1, ms.Len())
	assert.Equal(t, byte(3), ms.At(0))

	cp := NewByteSlice()
	ms.CopyTo(cp)
	ms.SetAt(0, byte(2))
	assert.Equal(t, byte(2), ms.At(0))
	assert.Equal(t, byte(3), cp.At(0))
	ms.CopyTo(cp)
	assert.Equal(t, byte(2), cp.At(0))

	mv := NewByteSlice()
	ms.MoveTo(mv)
	assert.Equal(t, 0, ms.Len())
	assert.Equal(t, 1, mv.Len())
	assert.Equal(t, byte(2), mv.At(0))
	ms.FromRaw([]byte{1, 2, 3})
	ms.MoveTo(mv)
	assert.Equal(t, 3, mv.Len())
	assert.Equal(t, byte(1), mv.At(0))
}

func TestByteSliceAppend(t *testing.T) {
	ms := NewByteSlice()
	ms.FromRaw([]byte{1, 2, 3})
	ms.Append(4, 5)
	assert.Equal(t, 5, ms.Len())
	assert.Equal(t, byte(5), ms.At(4))
}

func TestByteSliceEnsureCapacity(t *testing.T) {
	ms := NewByteSlice()
	ms.EnsureCapacity(4)
	assert.Equal(t, 4, cap(*ms.getOrig()))
	ms.EnsureCapacity(2)
	assert.Equal(t, 4, cap(*ms.getOrig()))
}
