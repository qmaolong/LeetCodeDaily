package main

import "testing"

func Test_findLongestWord(t *testing.T) {
	type args struct {
		s          string
		dictionary []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{
			name: "t1",
			args: args{
				s:          "abpcplea",
				dictionary: []string{"ale", "apple", "monkey", "plea"},
			},
			wantRes: "apple",
		},
		{
			name: "t2",
			args: args{
				s:          "abpcplea",
				dictionary: []string{"a", "b", "c"},
			},
			wantRes: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findLongestWord(tt.args.s, tt.args.dictionary); gotRes != tt.wantRes {
				t.Errorf("findLongestWord() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
