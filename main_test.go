package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 7, 1000} {
		data := generateRandomElements(size)
		expected := size
		if expected < 0 {
			expected = 0
		}
		if len(data) != expected {
			t.Fatalf("size %d: length %d", size, len(data))
		}
		for _, value := range data {
			if value <= 0 {
				t.Fatalf("size %d: nonpositive value %d", size, value)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	cases := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{12}, 12},
		{"first", []int{9, 3, 1}, 9},
		{"last", []int{1, 3, 9}, 9},
		{"repeated", []int{9, 9, 9}, 9},
		{"negative", []int{-8, -2, -3}, -2},
		{"zero", []int{0, 0}, 0},
		{"tail", []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, 99},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			if got := maximum(test.data); got != test.want {
				t.Fatalf("got %d, want %d", got, test.want)
			}
			if got := maxChunks(test.data); got != test.want {
				t.Fatalf("parallel got %d, want %d", got, test.want)
			}
		})
	}
}

func TestMaxChunksBoundaries(t *testing.T) {
	for _, size := range []int{7, 8, 9, 15, 16, 17, 1001} {
		for peak := 0; peak < size; peak++ {
			data := make([]int, size)
			for i := range data {
				data[i] = -2
			}
			data[peak] = -1
			if got := maxChunks(data); got != -1 {
				t.Fatalf("size %d peak %d: got %d", size, peak, got)
			}
		}
	}
}
