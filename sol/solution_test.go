package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	points := [][]int{{1, 3}, {-2, 2}}
	k := 1
	for idx := 0; idx < b.N; idx++ {
		kClosest(points, k)
	}
}
func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "points = [[1,3],[-2,2]], k = 1",
			args: args{points: [][]int{{1, 3}, {-2, 2}}, k: 1},
			want: [][]int{{-2, 2}},
		},
		{
			name: "Input: points = [[3,3],[5,-1],[-2,4]], k = 2",
			args: args{points: [][]int{{3, 3}, {5, -1}, {-2, 4}}, k: 2},
			want: [][]int{{-2, 4}, {3, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosest(tt.args.points, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
