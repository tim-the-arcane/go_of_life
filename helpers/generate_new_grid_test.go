package helpers

import (
	"reflect"
	"testing"
)

func TestGenerateNewGrid(t *testing.T) {
	type args struct {
		rows   int
		colums int
	}

	tests := []struct {
		name              string
		args              args
		wantGeneratedGrid [][]bool
	}{
		{
			name: "GenerateNewGrid(3, 3)",
			args: args{
				rows:   3,
				colums: 3,
			},
			wantGeneratedGrid: [][]bool{
				{false, false, false},
				{false, false, false},
				{false, false, false},
			},
		},
		{
			name: "GenerateNewGrid(3, 4)",
			args: args{
				rows:   3,
				colums: 4,
			},
			wantGeneratedGrid: [][]bool{
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
			},
		},
		{
			name: "GenerateNewGrid(4, 3)",
			args: args{
				rows:   4,
				colums: 3,
			},
			wantGeneratedGrid: [][]bool{
				{false, false, false},
				{false, false, false},
				{false, false, false},
				{false, false, false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotGeneratedGrid := GenerateNewGrid(tt.args.rows, tt.args.colums); !reflect.DeepEqual(gotGeneratedGrid, tt.wantGeneratedGrid) {
				t.Errorf("GenerateNewGrid() = %v, want %v", gotGeneratedGrid, tt.wantGeneratedGrid)
			}
		})
	}
}
