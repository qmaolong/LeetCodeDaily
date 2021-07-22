package main

import (
	"reflect"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	n1 := Node{
		Val: 7,
	}
	n2 := Node{
		Val: 13,
	}
	n3 := Node{
		Val: 11,
	}
	n4 := Node{
		Val: 10,
	}

	n5 := Node{
		Val:  1,
		Next: nil,
	}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5
	n1.Random = nil
	n2.Random = &n1
	n3.Random = &n5
	n4.Random = &n3
	n5.Random = &n1

	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "t1",
			args: args{
				head: &n1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}
