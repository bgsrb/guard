package guard

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bgsrb/guard/messages"
)

func Test_MinLength(t *testing.T) {
	type fields struct {
		name      string
		value     interface{}
		minlength int
	}
	tests := []struct {
		name   string
		fields fields
		want   Argument
	}{
		{
			name: "not_min_length",
			fields: fields{
				name:      "",
				value:     "1",
				minlength: 2,
			},
			want: Argument{
				Err:   fmt.Errorf(messages.MinLength, "", 2),
				Name:  "",
				Value: "1",
			},
		}, {
			name: "min_length",
			fields: fields{
				name:      "",
				value:     "01",
				minlength: 2,
			},
			want: Argument{
				Err:   nil,
				Name:  "",
				Value: "01",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Guard(tt.fields.value, tt.fields.name)
			if got := g.MinLength(tt.fields.minlength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("guard.MinLength() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
