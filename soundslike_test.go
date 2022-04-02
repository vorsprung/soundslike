package soundslike

import (
	"testing"

	"github.com/dlclark/metaphone3"
)

// TestPhonic  tests the Phonic struct
func TestPhonic_FindEndingToken(t *testing.T) {
	type fields struct {
		Setup []setup
		enc   *metaphone3.Encoder
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Phonic{
				Setup: tt.fields.Setup,
				enc:   tt.fields.enc,
			}
			if got := e.FindEndingToken(tt.args.s); got != tt.want {
				t.Errorf("Phonic.FindEndingToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
