package mymath

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Positive number", args{16}, 4},
		{"Zero", args{0}, 0},
		{"Negative number", args{-4}, math.NaN()},
	}
	const tolerance = 1e-10
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sqrt(tt.args.x); tt.args.x < 0 {
				if !math.IsNaN(got) {
					t.Errorf("Sqrt() = %v, want %v", got, math.NaN())
				}
			} else if math.Abs(got-tt.want) > tolerance {
				t.Errorf("Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCeil(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Positive number", args{3.14}, 4},
		{"Negative number", args{-3.14}, -3},
		{"Integer", args{5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ceil(tt.args.x); got != tt.want {
				t.Errorf("Ceil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloor(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Positive number", args{3.14}, 3},
		{"Negative number", args{-3.14}, -4},
		{"Integer", args{5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Floor(tt.args.x); got != tt.want {
				t.Errorf("Floor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPow(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Positive base and exponent", args{2, 3}, 8},
		{"Negative base and positive exponent", args{-2, 3}, -8},
		{"Positive base and negative exponent", args{2, -3}, 0.125},
		{"Zero base", args{0, 5}, 0},
		{"Non-zero base and zero exponent", args{5, 0}, 1},
	}
	const tolerance = 1e-10
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pow(tt.args.x, tt.args.y); math.Abs(got-tt.want) > tolerance {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"x is greater", args{10, 5}, 10},
		{"y is greater", args{5, 10}, 10},
		{"x and y are equal", args{7, 7}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"x is greater", args{10, 5}, 5},
		{"y is greater", args{5, 10}, 5},
		{"x and y are equal", args{7, 7}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Positive number", args{5}, 5},
		{"Negative number", args{-5}, 5},
		{"Zero", args{0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.x); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYn(t *testing.T) {
	type args struct {
		x int
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"n=0, x=2", args{0, 2}, 0.51},
		{"n=1, x=2", args{1, 2}, -0.11},
		{"n=2, x=2", args{2, 2}, -0.62},
		{"n=3, x=5", args{3, 5}, 0.15},
		{"n=4, x=10", args{4, 10}, -0.14},
	}
	const tolerance = 1e-2
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Yn(tt.args.x, tt.args.y); math.Abs(got-tt.want) > tolerance {
				t.Errorf("Yn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAcos(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Valid input (0)", args{0}, 1.57},
		{"Valid input (0.5)", args{0.5}, 1.05},
		{"Valid input (-0.5)", args{-0.5}, 2.09},
		{"Valid input (0.9)", args{0.9}, 0.45},
		{"Valid input (-0.9)", args{-0.9}, 2.69},
	}
	const tolerance = 1e-2
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Acos(tt.args.x); math.Abs(got-tt.want) > tolerance {
				t.Errorf("Acos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAcosh(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Valid input (1)", args{1}, 0.00},
		{"Valid input (2)", args{2}, 1.32},
		{"Valid input (5)", args{5}, 2.29},
		{"Valid input (10)", args{10}, 2.99},
		{"Valid input (20)", args{20}, 3.69},
	}
	const tolerance = 1e-2
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Acosh(tt.args.x); math.Abs(got-tt.want) > tolerance {
				t.Errorf("Acosh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsin(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Valid input (0)", args{0}, 0.00},
		{"Valid input (0.5)", args{0.5}, 0.52},
		{"Valid input (-0.5)", args{-0.5}, -0.52},
		{"Valid input (0.9)", args{0.9}, 1.12},
		{"Valid input (-0.9)", args{-0.9}, -1.12},
	}
	const tolerance = 1e-2
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Asin(tt.args.x); math.Abs(got-tt.want) > tolerance {
				t.Errorf("Asin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsinh(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Positive number (2.5)", args{2.5}, 1.65},
		{"Negative number (-2.5)", args{-2.5}, -1.65},
		{"Positive number (5)", args{5}, 2.31},
		{"Negative number (-5)", args{-5}, -2.31},
		{"Positive number (10)", args{10}, 2.99},
		{"Negative number (-10)", args{-10}, -2.99},
	}
	const tolerance = 1e-2
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Asinh(tt.args.x); math.Abs(got-tt.want) > tolerance {
				t.Errorf("Asinh() = %v, want %v", got, tt.want)
			}
		})
	}
}
