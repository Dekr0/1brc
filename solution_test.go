package main

import (
    "testing"
)

func BenchmarkSolution(b *testing.B) {
    if err := Solution("./measurements.txt"); err != nil {
        b.Fatal(err)
    }
}
