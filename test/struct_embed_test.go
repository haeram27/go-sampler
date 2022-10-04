package test

import (
	"testing"
)

type Inner struct {
	name_over    string
	name_no_over string
}

func (i Inner) func_override() string {
	return i.name_over
}

func (i Inner) func_no_override() string {
	return i.name_over
}

type Outter struct {
	Inner           // struct embedding
	nest      Inner // struct nesting
	name_over string
}

func (o Outter) func_override() string {
	return o.name_over
}

func TestStructEmbed(t *testing.T) {
	o := Outter{
		// initialize embedded struct
		Inner: Inner{
			name_over:    "Embedded",
			name_no_over: "Embedded",
		},

		// initialize nested struct
		nest: Inner{
			name_over:    "Nested",
			name_no_over: "Nested",
		},

		name_over: "Outter",
	}

	t.Log("\n//**** outter ****")
	t.Log("o.oname: ", o.name_over)                       // Outter
	t.Log("o.iname: ", o.name_no_over)                    // Embedded
	t.Log("o.func_override(): ", o.func_override())       // Outter
	t.Log("o.func_no_override(): ", o.func_no_override()) // Embedded

	t.Log("\n//**** embedded ****")
	t.Log("o.Inner.name_over: ", o.Inner.name_over)                   // Embedded
	t.Log("o.Inner.name_no_over: ", o.Inner.name_no_over)             // Embedded
	t.Log("o.Inner.func_override(): ", o.Inner.func_override())       // Embedded
	t.Log("o.Inner.func_no_override(): ", o.Inner.func_no_override()) // Embedded

	t.Log("\n//**** nested ****")
	t.Log("o.nest.name_over: ", o.nest.name_over)                   // Nested
	t.Log("o.nest.name_no_over: ", o.nest.name_no_over)             // Nested
	t.Log("o.nest.func_override(): ", o.nest.func_override())       // Nested
	t.Log("o.nest.func_no_override(): ", o.nest.func_no_override()) // Nested

}
