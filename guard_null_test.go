package guard

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bgsrb/guard/messages"
)

func Test_Null(t *testing.T) {
	type fields struct {
		name  string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   Argument
	}{
		{
			name: "null",
			fields: fields{
				name:  "",
				value: nil,
			},
			want: Argument{
				Error: fmt.Errorf(messages.NotEmpty, ""),
				Name:  "",
				Value: nil,
			},
		}, {
			name: "not_null",
			fields: fields{
				name:  "",
				value: "not_null",
			},
			want: Argument{
				Error: nil,
				Name:  "",
				Value: "not_null",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Guard(tt.fields.value, tt.fields.name)
			if got := g.NotNull(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("guard.NotNull() = %v, want %v", got, tt.want)
			}
		})
	}
}
