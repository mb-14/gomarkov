package gomarkov

import (
	"reflect"
	"testing"
)

func TestChain_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		order   int
		data    [][]string
		want    string
		wantErr bool
	}{
		{"Empty chain", 2, [][]string{}, `{"int":2,"spool_map":{},"freq_mat":{}}`, false},
		{"Empty chain, order 1", 1, [][]string{}, `{"int":1,"spool_map":{},"freq_mat":{}}`, false},
		{"Trained once", 1, [][]string{{"Test"}}, `{"int":1,"spool_map":{"$":0,"Test":1,"^":2},"freq_mat":{"0":{"1":1},"1":{"2":1}}}`, false},
		{"Trained on more data", 1, [][]string{{"test", "data"}, {"test", "data"}, {"test", "node"}}, `{"int":1,"spool_map":{"$":0,"^":3,"data":2,"node":4,"test":1},"freq_mat":{"0":{"1":3},"1":{"2":2,"4":1},"2":{"3":2},"4":{"3":1}}}`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chain := NewChain(tt.order)
			for _, data := range tt.data {
				chain.Add(data)
			}

			got, err := chain.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Chain.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("Chain.MarshalJSON() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestChain_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		args    []byte
		wantErr bool
	}{
		{"Empty chain", []byte(`{"int":2,"spool_map":{},"freq_mat":{}}`), false},
		{"More complex chain", []byte(`{"int":1,"spool_map":{"$":0,"^":3,"data":2,"node":4,"test":1},"freq_mat":{"0":{"1":3},"1":{"2":2,"4":1},"2":{"3":2},"4":{"3":1}}}`), false},
		{"Invalid json", []byte(`{{"int":2,"spool_map":{},"freq_mat":{}}`), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chain := NewChain(1)

			if err := chain.UnmarshalJSON(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Chain.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewChain(t *testing.T) {
	type args struct {
		order int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Order 1", args{order: 1}, 1},
		{"Order 2", args{order: 2}, 2},
		{"Order 50", args{order: 50}, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChain(tt.args.order); got.Order != tt.want {
				t.Errorf("NewChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChain_Add(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name  string
		chain *Chain
		args  args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.chain.Add(tt.args.input)
		})
	}
}

func TestChain_TransitionProbability(t *testing.T) {
	type args struct {
		next    string
		current NGram
	}
	tests := []struct {
		name    string
		chain   []byte
		args    args
		want    float64
		wantErr bool
	}{
		{
			"Simple transition positive",
			[]byte(`{"int":1,"spool_map":{"$":0,"Test":1,"^":2},"freq_mat":{"0":{"1":1},"1":{"2":1}}}`),
			args{next: "Test", current: NGram{"$"}},
			1,
			false,
		},
		{
			"Simple transition negative",
			[]byte(`{"int":1,"spool_map":{"$":0,"Test":1,"^":2},"freq_mat":{"0":{"1":1},"1":{"2":1}}}`),
			args{next: "Test", current: NGram{"Test"}},
			0,
			false,
		},
		{
			"Unknown next Ngram",
			[]byte(`{"int":1,"spool_map":{"$":0,"Test":1,"^":2},"freq_mat":{"0":{"1":1},"1":{"2":1}}}`),
			args{next: "Unknown", current: NGram{"Test"}},
			0,
			false,
		},
		{
			"Unknown ncurent Ngram",
			[]byte(`{"int":1,"spool_map":{"$":0,"Test":1,"^":2},"freq_mat":{"0":{"1":1},"1":{"2":1}}}`),
			args{next: "Test", current: NGram{"Unknown"}},
			0,
			false,
		},
		{
			"Invalid Ngram",
			[]byte(`{"int":1,"spool_map":{"$":0,"Test":1,"^":2},"freq_mat":{"0":{"1":1},"1":{"2":1}}}`),
			args{next: "Unknown", current: NGram{"Test", "data"}},
			0,
			true,
		},
		{
			"More than 1 option",
			[]byte(`{"int":1,"spool_map":{"$":0,"^":3,"data":2,"node":4,"test":1},"freq_mat":{"0":{"1":3},"1":{"2":2,"4":2},"2":{"3":2},"4":{"3":1}}}`),
			args{next: "node", current: NGram{"test"}},
			0.5,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chain := NewChain(1)
			chain.UnmarshalJSON(tt.chain)

			got, err := chain.TransitionProbability(tt.args.next, tt.args.current)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chain.TransitionProbability() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Chain.TransitionProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChain_Generate(t *testing.T) {
	type args struct {
		current NGram
	}
	tests := []struct {
		name    string
		chain   []byte
		args    args
		want    string
		wantErr bool
	}{
		{
			"Start of simple chain",
			[]byte(`{"int":1,"spool_map":{"$":0,"Test":1,"^":2},"freq_mat":{"0":{"1":1},"1":{"2":1}}}`),
			args{current: NGram{"$"}},
			"Test",
			false,
		},
		{
			"End of simple chain",
			[]byte(`{"int":1,"spool_map":{"$":0,"Test":1,"^":2},"freq_mat":{"0":{"1":1},"1":{"2":1}}}`),
			args{current: NGram{"Test"}},
			"^",
			false,
		},
		{
			"Complex chain",
			[]byte(`{"int":1,"spool_map":{"$":0,"^":3,"data":2,"node":4,"test":1},"freq_mat":{"0":{"1":3},"1":{"2":2,"4":0},"2":{"3":2},"4":{"3":1}}}`),
			args{current: NGram{"test"}},
			"data",
			false,
		},
		{
			"Invalid Ngram",
			[]byte(`{"int":1,"spool_map":{"$":0,"Test":1,"^":2},"freq_mat":{"0":{"1":1},"1":{"2":1}}}`),
			args{current: NGram{"Invalid", "Ngram"}},
			"",
			true,
		},
		{
			"Unknown Ngram",
			[]byte(`{"int":1,"spool_map":{"$":0,"Test":1,"^":2},"freq_mat":{"0":{"1":1},"1":{"2":1}}}`),
			args{current: NGram{"Unknown"}},
			"",
			true,
		},
		{
			"No next state",
			[]byte(`{"int":1,"spool_map":{"$":0,"Test":1,"^":2},"freq_mat":{"0":{"1":1},"1":{"2":1}}}`),
			args{current: NGram{"^"}},
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chain := NewChain(1)
			chain.UnmarshalJSON(tt.chain)

			got, err := chain.Generate(tt.args.current)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chain.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Chain.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
