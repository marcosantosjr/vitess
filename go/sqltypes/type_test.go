/*
Copyright 2019 The Vitess Authors.

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

package sqltypes

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	querypb "vitess.io/vitess/go/vt/proto/query"
)

func TestTypeValues(t *testing.T) {
	testcases := []struct {
		defined  querypb.Type
		expected int
	}{{
		defined:  Null,
		expected: 0,
	}, {
		defined:  Int8,
		expected: 1 | flagIsIntegral,
	}, {
		defined:  Uint8,
		expected: 2 | flagIsIntegral | flagIsUnsigned,
	}, {
		defined:  Int16,
		expected: 3 | flagIsIntegral,
	}, {
		defined:  Uint16,
		expected: 4 | flagIsIntegral | flagIsUnsigned,
	}, {
		defined:  Int24,
		expected: 5 | flagIsIntegral,
	}, {
		defined:  Uint24,
		expected: 6 | flagIsIntegral | flagIsUnsigned,
	}, {
		defined:  Int32,
		expected: 7 | flagIsIntegral,
	}, {
		defined:  Uint32,
		expected: 8 | flagIsIntegral | flagIsUnsigned,
	}, {
		defined:  Int64,
		expected: 9 | flagIsIntegral,
	}, {
		defined:  Uint64,
		expected: 10 | flagIsIntegral | flagIsUnsigned,
	}, {
		defined:  Float32,
		expected: 11 | flagIsFloat,
	}, {
		defined:  Float64,
		expected: 12 | flagIsFloat,
	}, {
		defined:  Timestamp,
		expected: 13 | flagIsQuoted,
	}, {
		defined:  Date,
		expected: 14 | flagIsQuoted,
	}, {
		defined:  Time,
		expected: 15 | flagIsQuoted,
	}, {
		defined:  Datetime,
		expected: 16 | flagIsQuoted,
	}, {
		defined:  Year,
		expected: 17 | flagIsIntegral | flagIsUnsigned,
	}, {
		defined:  Decimal,
		expected: 18,
	}, {
		defined:  Text,
		expected: 19 | flagIsQuoted | flagIsText,
	}, {
		defined:  Blob,
		expected: 20 | flagIsQuoted | flagIsBinary,
	}, {
		defined:  VarChar,
		expected: 21 | flagIsQuoted | flagIsText,
	}, {
		defined:  VarBinary,
		expected: 22 | flagIsQuoted | flagIsBinary,
	}, {
		defined:  Char,
		expected: 23 | flagIsQuoted | flagIsText,
	}, {
		defined:  Binary,
		expected: 24 | flagIsQuoted | flagIsBinary,
	}, {
		defined:  Bit,
		expected: 25 | flagIsQuoted,
	}, {
		defined:  Enum,
		expected: 26 | flagIsQuoted,
	}, {
		defined:  Set,
		expected: 27 | flagIsQuoted,
	}, {
		defined:  Geometry,
		expected: 29 | flagIsQuoted,
	}, {
		defined:  TypeJSON,
		expected: 30 | flagIsQuoted,
	}, {
		defined:  Expression,
		expected: 31,
	}, {
		defined:  HexNum,
		expected: 32 | flagIsText,
	}, {
		defined:  HexVal,
		expected: 33 | flagIsText,
	}, {
		defined:  BitNum,
		expected: 34 | flagIsText,
	}}
	for _, tcase := range testcases {
		if int(tcase.defined) != tcase.expected {
			t.Errorf("Type %s: %d, want: %d", tcase.defined, int(tcase.defined), tcase.expected)
		}
	}
}

// TestCategory verifies that the type categorizations
// are non-overlapping and complete.
func TestCategory(t *testing.T) {
	alltypes := []querypb.Type{
		Null,
		Int8,
		Uint8,
		Int16,
		Uint16,
		Int24,
		Uint24,
		Int32,
		Uint32,
		Int64,
		Uint64,
		Float32,
		Float64,
		Timestamp,
		Date,
		Time,
		Datetime,
		Year,
		Decimal,
		Text,
		Blob,
		VarChar,
		VarBinary,
		Char,
		Binary,
		Bit,
		Enum,
		Set,
		Geometry,
		TypeJSON,
		Expression,
		HexNum,
		HexVal,
		BitNum,
	}
	for _, typ := range alltypes {
		matched := false
		if IsSigned(typ) {
			if !IsIntegral(typ) {
				t.Errorf("Signed type %v is not an integral", typ)
			}
			matched = true
		}
		if IsUnsigned(typ) {
			if !IsIntegral(typ) {
				t.Errorf("Unsigned type %v is not an integral", typ)
			}
			if matched {
				t.Errorf("%v matched more than one category", typ)
			}
			matched = true
		}
		if IsFloat(typ) {
			if matched {
				t.Errorf("%v matched more than one category", typ)
			}
			matched = true
		}
		if IsQuoted(typ) {
			if matched {
				t.Errorf("%v matched more than one category", typ)
			}
			matched = true
		}
		if typ == Null || typ == Decimal || typ == Expression || typ == Bit ||
			typ == HexNum || typ == HexVal || typ == BitNum {
			if matched {
				t.Errorf("%v matched more than one category", typ)
			}
			matched = true
		}
		if !matched {
			t.Errorf("%v matched no category", typ)
		}
	}
}

func TestIsFunctions(t *testing.T) {
	if IsIntegral(Null) {
		t.Error("Null: IsIntegral, must be false")
	}
	if !IsIntegral(Int64) {
		t.Error("Int64: !IsIntegral, must be true")
	}
	if IsSigned(Uint64) {
		t.Error("Uint64: IsSigned, must be false")
	}
	if !IsSigned(Int64) {
		t.Error("Int64: !IsSigned, must be true")
	}
	if IsUnsigned(Int64) {
		t.Error("Int64: IsUnsigned, must be false")
	}
	if !IsUnsigned(Uint64) {
		t.Error("Uint64: !IsUnsigned, must be true")
	}
	if IsFloat(Int64) {
		t.Error("Int64: IsFloat, must be false")
	}
	if !IsFloat(Float64) {
		t.Error("Uint64: !IsFloat, must be true")
	}
	if IsQuoted(Int64) {
		t.Error("Int64: IsQuoted, must be false")
	}
	if !IsQuoted(Binary) {
		t.Error("Binary: !IsQuoted, must be true")
	}
	if IsText(Int64) {
		t.Error("Int64: IsText, must be false")
	}
	if !IsText(Char) {
		t.Error("Char: !IsText, must be true")
	}
	if IsBinary(Int64) {
		t.Error("Int64: IsBinary, must be false")
	}
	if !IsBinary(Binary) {
		t.Error("Char: !IsBinary, must be true")
	}
	if !IsNumber(Int64) {
		t.Error("Int64: !isNumber, must be true")
	}
}

func TestTypeToMySQL(t *testing.T) {
	v, f := TypeToMySQL(Bit)
	if v != 16 {
		t.Errorf("Bit: %d, want 16", v)
	}
	if f != mysqlUnsigned {
		t.Errorf("Bit flag: %x, want %x", f, mysqlUnsigned)
	}
	v, f = TypeToMySQL(Date)
	if v != 10 {
		t.Errorf("Bit: %d, want 10", v)
	}
	if f != mysqlBinary {
		t.Errorf("Bit flag: %x, want %x", f, mysqlBinary)
	}
}

func TestMySQLToType(t *testing.T) {
	testcases := []struct {
		intype  byte
		inflags int64
		outtype querypb.Type
	}{{
		intype:  1,
		outtype: Int8,
	}, {
		intype:  1,
		inflags: mysqlUnsigned,
		outtype: Uint8,
	}, {
		intype:  2,
		outtype: Int16,
	}, {
		intype:  2,
		inflags: mysqlUnsigned,
		outtype: Uint16,
	}, {
		intype:  3,
		outtype: Int32,
	}, {
		intype:  3,
		inflags: mysqlUnsigned,
		outtype: Uint32,
	}, {
		intype:  4,
		outtype: Float32,
	}, {
		intype:  5,
		outtype: Float64,
	}, {
		intype:  6,
		outtype: Null,
	}, {
		intype:  7,
		outtype: Timestamp,
	}, {
		intype:  8,
		outtype: Int64,
	}, {
		intype:  8,
		inflags: mysqlUnsigned,
		outtype: Uint64,
	}, {
		intype:  9,
		outtype: Int24,
	}, {
		intype:  9,
		inflags: mysqlUnsigned,
		outtype: Uint24,
	}, {
		intype:  10,
		outtype: Date,
	}, {
		intype:  11,
		outtype: Time,
	}, {
		intype:  12,
		outtype: Datetime,
	}, {
		intype:  13,
		outtype: Year,
	}, {
		intype:  16,
		outtype: Bit,
	}, {
		intype:  245,
		outtype: TypeJSON,
	}, {
		intype:  246,
		outtype: Decimal,
	}, {
		intype:  249,
		outtype: Text,
	}, {
		intype:  250,
		outtype: Text,
	}, {
		intype:  251,
		outtype: Text,
	}, {
		intype:  252,
		outtype: Text,
	}, {
		intype:  252,
		inflags: mysqlBinary,
		outtype: Blob,
	}, {
		intype:  253,
		outtype: VarChar,
	}, {
		intype:  253,
		inflags: mysqlBinary,
		outtype: VarBinary,
	}, {
		intype:  254,
		outtype: Char,
	}, {
		intype:  254,
		inflags: mysqlBinary,
		outtype: Binary,
	}, {
		intype:  254,
		inflags: mysqlEnum,
		outtype: Enum,
	}, {
		intype:  254,
		inflags: mysqlSet,
		outtype: Set,
	}, {
		intype:  255,
		outtype: Geometry,
	}, {
		// Binary flag must be ignored.
		intype:  8,
		inflags: mysqlUnsigned | mysqlBinary,
		outtype: Uint64,
	}, {
		// Unsigned flag must be ignored
		intype:  252,
		inflags: mysqlUnsigned | mysqlBinary,
		outtype: Blob,
	}}
	for _, tcase := range testcases {
		got, err := MySQLToType(tcase.intype, tcase.inflags)
		if err != nil {
			t.Error(err)
		}
		if got != tcase.outtype {
			t.Errorf("MySQLToType(%d, %x): %v, want %v", tcase.intype, tcase.inflags, got, tcase.outtype)
		}
	}
}

func TestTypeError(t *testing.T) {
	_, err := MySQLToType(50, 0)
	want := "unsupported type: 50"
	if err == nil || err.Error() != want {
		t.Errorf("MySQLToType: %v, want %s", err, want)
	}
}

func TestTypeEquivalenceCheck(t *testing.T) {
	if !AreTypesEquivalent(Int16, Int16) {
		t.Errorf("Int16 and Int16 are same types.")
	}
	if AreTypesEquivalent(Int16, Int24) {
		t.Errorf("Int16 and Int24 are not same types.")
	}
	if !AreTypesEquivalent(VarChar, VarBinary) {
		t.Errorf("VarChar in binlog and VarBinary in schema are equivalent types.")
	}
	if AreTypesEquivalent(VarBinary, VarChar) {
		t.Errorf("VarBinary in binlog and VarChar in schema are not equivalent types.")
	}
	if !AreTypesEquivalent(Int16, Uint16) {
		t.Errorf("Int16 in binlog and Uint16 in schema are equivalent types.")
	}
	if AreTypesEquivalent(Uint16, Int16) {
		t.Errorf("Uint16 in binlog and Int16 in schema are not equivalent types.")
	}
}

func TestPrintTypeChecks(t *testing.T) {
	var funcs = []struct {
		name string
		f    func(p Type) bool
	}{
		{"IsSigned", IsSigned},
		{"IsFloat", IsFloat},
		{"IsUnsigned", IsUnsigned},
		{"IsIntegral", IsIntegral},
		{"IsText", IsText},
		{"IsNumber", IsNumber},
		{"IsQuoted", IsQuoted},
		{"IsBinary", IsBinary},
		{"IsDate", IsDate},
		{"IsNull", IsNull},
	}
	var types = []Type{
		Null,
		Int8,
		Uint8,
		Int16,
		Uint16,
		Int24,
		Uint24,
		Int32,
		Uint32,
		Int64,
		Uint64,
		Float32,
		Float64,
		Timestamp,
		Date,
		Time,
		Datetime,
		Year,
		Decimal,
		Text,
		Blob,
		VarChar,
		VarBinary,
		Char,
		Binary,
		Bit,
		Enum,
		Set,
		Geometry,
		TypeJSON,
		Expression,
		HexNum,
		HexVal,
		Tuple,
		BitNum,
	}

	for _, f := range funcs {
		var match []string
		for _, tt := range types {
			if f.f(tt) {
				match = append(match, tt.String())
			}
		}
		t.Logf("%s(): %s", f.name, strings.Join(match, ", "))
	}
}

func TestIsTextOrBinary(t *testing.T) {
	tests := []struct {
		name           string
		ty             querypb.Type
		isTextorBinary bool
	}{
		{
			name:           "null type",
			ty:             querypb.Type_NULL_TYPE,
			isTextorBinary: false,
		},
		{
			name:           "blob type",
			ty:             querypb.Type_BLOB,
			isTextorBinary: true,
		},
		{
			name:           "text type",
			ty:             querypb.Type_TEXT,
			isTextorBinary: true,
		},
		{
			name:           "binary type",
			ty:             querypb.Type_BINARY,
			isTextorBinary: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.isTextorBinary, IsTextOrBinary(tt.ty))
		})
	}
}

func TestIsDateOrTime(t *testing.T) {
	tests := []struct {
		name         string
		ty           querypb.Type
		isDateOrTime bool
	}{
		{
			name:         "null type",
			ty:           querypb.Type_NULL_TYPE,
			isDateOrTime: false,
		},
		{
			name:         "blob type",
			ty:           querypb.Type_BLOB,
			isDateOrTime: false,
		},
		{
			name:         "timestamp type",
			ty:           querypb.Type_TIMESTAMP,
			isDateOrTime: true,
		},
		{
			name:         "date type",
			ty:           querypb.Type_DATE,
			isDateOrTime: true,
		},
		{
			name:         "time type",
			ty:           querypb.Type_TIME,
			isDateOrTime: true,
		},
		{
			name:         "date time type",
			ty:           querypb.Type_DATETIME,
			isDateOrTime: true,
		},
		{
			name:         "year type",
			ty:           querypb.Type_YEAR,
			isDateOrTime: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.isDateOrTime, IsDateOrTime(tt.ty))
		})
	}
}
