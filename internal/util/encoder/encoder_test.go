package encoder

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeUrl(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_1",
			args: args{id: 1},
			want: "aaaaaaaaab",
		},
		{
			name: "test_2",
			args: args{id: 100},
			want: "aaaaaaaab" + string([]byte(alphabet)[37]),
		},
		{
			name: "test_3",
			args: args{id: 0},
			want: "aaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, EncodeUrl(tt.args.id, false),
			fmt.Sprintf("EncodeUrl() = %v, want %v", EncodeUrl(tt.args.id, false), tt.want))
	}
}
