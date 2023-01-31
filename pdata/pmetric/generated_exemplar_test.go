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

package pmetric

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/internal/data"
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestExemplar_MoveTo(t *testing.T) {
	ms := generateTestExemplar()
	dest := NewExemplar()
	ms.MoveTo(dest)
	assert.Equal(t, NewExemplar(), ms)
	assert.Equal(t, generateTestExemplar(), dest)
}

func TestExemplar_CopyTo(t *testing.T) {
	ms := NewExemplar()
	orig := NewExemplar()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestExemplar()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestExemplar_Timestamp(t *testing.T) {
	ms := NewExemplar()
	assert.Equal(t, pcommon.Timestamp(0), ms.Timestamp())
	testValTimestamp := pcommon.Timestamp(1234567890)
	ms.SetTimestamp(testValTimestamp)
	assert.Equal(t, testValTimestamp, ms.Timestamp())
}

func TestExemplar_ValueType(t *testing.T) {
	tv := NewExemplar()
	assert.Equal(t, ExemplarValueTypeEmpty, tv.ValueType())
}

func TestExemplar_DoubleValue(t *testing.T) {
	ms := NewExemplar()
	assert.Equal(t, float64(0.0), ms.DoubleValue())
	ms.SetDoubleValue(float64(17.13))
	assert.Equal(t, float64(17.13), ms.DoubleValue())
	assert.Equal(t, ExemplarValueTypeDouble, ms.ValueType())
}

func TestExemplar_IntValue(t *testing.T) {
	ms := NewExemplar()
	assert.Equal(t, int64(0), ms.IntValue())
	ms.SetIntValue(int64(17))
	assert.Equal(t, int64(17), ms.IntValue())
	assert.Equal(t, ExemplarValueTypeInt, ms.ValueType())
}

func TestExemplar_FilteredAttributes(t *testing.T) {
	ms := NewExemplar()
	assert.Equal(t, pcommon.NewMap(), ms.FilteredAttributes())
	internal.FillTestMap(internal.Map(ms.FilteredAttributes()))
	assert.Equal(t, pcommon.Map(internal.GenerateTestMap()), ms.FilteredAttributes())
}

func TestExemplar_TraceID(t *testing.T) {
	ms := NewExemplar()
	assert.Equal(t, pcommon.TraceID(data.TraceID([16]byte{})), ms.TraceID())
	testValTraceID := pcommon.TraceID(data.TraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1}))
	ms.SetTraceID(testValTraceID)
	assert.Equal(t, testValTraceID, ms.TraceID())
}

func TestExemplar_SpanID(t *testing.T) {
	ms := NewExemplar()
	assert.Equal(t, pcommon.SpanID(data.SpanID([8]byte{})), ms.SpanID())
	testValSpanID := pcommon.SpanID(data.SpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1}))
	ms.SetSpanID(testValSpanID)
	assert.Equal(t, testValSpanID, ms.SpanID())
}

func generateTestExemplar() Exemplar {
	tv := NewExemplar()
	fillTestExemplar(tv)
	return tv
}

func fillTestExemplar(tv Exemplar) {
	tv.orig.TimeUnixNano = 1234567890
	tv.orig.Value = &otlpmetrics.Exemplar_AsInt{AsInt: int64(17)}
	internal.FillTestMap(internal.NewMap(&tv.orig.FilteredAttributes))
	tv.orig.TraceId = data.TraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})
	tv.orig.SpanId = data.SpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1})
}
