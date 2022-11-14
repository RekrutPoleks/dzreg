package main

import (
	"reflect"
	"testing"
)

func Test_easySol(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"Cong_1",
			args{[]string{"1-2=?", "1", "-", "2"}},
			[]byte("1-2=-1\n"),
			false,
		},
		{
			"Cong_2",
			args{[]string{"4+2=?", "4", "+", "2"}},
			[]byte("4+2=6\n"),
			false,
		},
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := easySol(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("easySol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {

				t.Errorf("easySol() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
