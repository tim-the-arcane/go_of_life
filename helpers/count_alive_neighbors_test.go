package helpers

import "testing"

func TestCountAliveNeighbors(t *testing.T) {
	type args struct {
		grid        *[][]bool
		rowIndex    int
		columnIndex int
	}

	tests := []struct {
		name      string
		args      args
		wantAlive int
	}{
		{"Should return 0 when there are no alive neighbours", args{&[][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, false},
		}, 1, 1}, 0},
		{"Should return 1 when there is 1 alive neighbour", args{&[][]bool{
			{false, false, false},
			{false, true, false},
			{false, true, false},
		}, 1, 1}, 1},
		{"Should return 8 when there are 8 alive neighbours", args{&[][]bool{
			{true, true, true},
			{true, true, true},
			{true, true, true},
		}, 1, 1}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAlive := CountAliveNeighbors(tt.args.grid, tt.args.rowIndex, tt.args.columnIndex); gotAlive != tt.wantAlive {
				t.Errorf("CountAliveNeighbors() = %v, want %v", gotAlive, tt.wantAlive)
			}
		})
	}
}
