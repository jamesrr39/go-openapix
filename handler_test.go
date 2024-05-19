package openapix

import "testing"

func Test_createEndpointName(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "byID",
			args: args{
				title: "Get User By ID",
			},
			want: "getUserById",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createEndpointName(tt.args.title); got != tt.want {
				t.Errorf("createEndpointName() = %v, want %v", got, tt.want)
			}
		})
	}
}
