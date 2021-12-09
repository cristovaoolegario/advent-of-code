package main

import "testing"

func Test_main(t *testing.T) {
	t.Run(`Should return 1 when number is bigger than the previous one
				 Should return 0 when number is equal to the last one
				 Should return 0 when number is smaller then the previous one`, func(t *testing.T) {
		testItems := []struct {
			last     int
			actual   int
			expected int
		}{{0, 1, 1}, {-1, 0, 1}, {0, -1, 0}, {0, 0, 0}}

		for _, item := range testItems {
			response := isBiggerThanTheLastNumber(item.last, item.expected)

			if response != item.expected {
				t.Errorf("Expected %v, got %v", item.expected, response)
			}
		}
	})
}
