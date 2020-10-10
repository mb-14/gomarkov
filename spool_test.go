package gomarkov

import "testing"

func Test_spool_add(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		s    *spool
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.add(tt.args.str); got != tt.want {
				t.Errorf("spool.add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_spool_get(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		s     *spool
		args  args
		want  int
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.get(tt.args.str)
			if got != tt.want {
				t.Errorf("spool.get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("spool.get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
