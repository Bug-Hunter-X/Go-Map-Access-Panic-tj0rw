# Go Map Access Panic

This repository demonstrates a common error in Go: panicking due to accessing a non-existent key in a map.

The `bug.go` file shows the problematic code.  Attempting to access `m[0]` where `m` is an uninitialized map will cause a runtime panic.

The `bugSolution.go` file provides a corrected version that safely handles the case where the key might not exist.

This is a crucial concept in Go to prevent unexpected program termination. Always check if a key exists before accessing it in a map to avoid panics.