package guard

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bgsrb/guard/messages"
)

func Test_Empty(t *testing.T) {
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
			name: "empty",
			fields: fields{
				name:  "",
				value: "",
			},
			want: Argument{
				Error: fmt.Errorf(messages.NotEmpty, ""),
				Name:  "",
				Value: "",
			},
		}, {
			name: "not_empty",
			fields: fields{
				name:  "",
				value: "not_empty",
			},
			want: Argument{
				Error: nil,
				Name:  "",
				Value: "not_empty",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Guard(tt.fields.value, tt.fields.name)
			if got := g.NotEmpty(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("guard.NotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
