package gomarkov

import (
	"reflect"
	"testing"
)

func TestNGram_key(t *testing.T) {
	tests := []struct {
		name  string
		ngram NGram
		want  string
	}{
		{"One word", NGram{"One"}, "One"},
		{"Two words", NGram{"Two", "words"}, "Two_words"},
		{"No words", NGram{""}, ""},
		{"Empty NGram", NGram{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ngram.key(); got != tt.want {
				t.Errorf("NGram.key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sparseArray_sum(t *testing.T) {
	tests := []struct {
		name string
		s    sparseArray
		want int
	}{
		{"One element", sparseArray{1: 1}, 1},
		{"Two elements", sparseArray{1: 1, 2: 1}, 2},
		{"No elements", sparseArray{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.sum(); got != tt.want {
				t.Errorf("sparseArray.sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_array(t *testing.T) {
	type args struct {
		value string
		count int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array(tt.args.value, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("array() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakePairs(t *testing.T) {
	type args struct {
		tokens []string
		order  int
	}
	tests := []struct {
		name string
		args args
		want []Pair
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakePairs(tt.args.tokens, tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
