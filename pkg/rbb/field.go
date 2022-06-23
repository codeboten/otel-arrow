// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rbb

import "strings"

// Field is a scalar or a composite named value.
type Field struct {
	Name  string
	Value Value
}

// Normalize normalizes the field name and value.
func (f Field) Normalize() {
	f.Value.Normalize()
}

func (f Field) WriteSignature(sig *strings.Builder) {
	sig.WriteString(f.Name)
	sig.WriteString(":")
	switch v := f.Value.(type) {
	case *Bool:
		sig.WriteString(BOOL_SIG)
	case *I8:
		sig.WriteString(I8_SIG)
	case *I16:
		sig.WriteString(I16_SIG)
	case *I32:
		sig.WriteString(I32_SIG)
	case *I64:
		sig.WriteString(I64_SIG)
	case *U8:
		sig.WriteString(U8_SIG)
	case *U16:
		sig.WriteString(U16_SIG)
	case *U32:
		sig.WriteString(U32_SIG)
	case *U64:
		sig.WriteString(U64_SIG)
	case *F32:
		sig.WriteString(F32_SIG)
	case *F64:
		sig.WriteString(F64_SIG)
	case *String:
		sig.WriteString(STRING_SIG)
	case *Binary:
		sig.WriteString(BINARY_SIG)
	case *Struct:
		sig.WriteString("{")
		for i, f := range v.fields {
			if i > 0 {
				sig.WriteByte(',')
			}
			f.WriteSignature(sig)
		}
		sig.WriteString("}")
	case *List:
		sig.WriteString("[")
		sig.WriteString(DataTypeSignature(ListDataType(v.values)))
		sig.WriteString("]")
	default:
		panic("unknown field value type")
	}
}
