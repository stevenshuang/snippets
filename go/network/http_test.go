package network

import (
    "fmt"
    "testing"
)

func TestHttpbinGet(t *testing.T) {
    fmt.Println("start httpbin get")
    Get()
}

func TestHttbinPost(t *testing.T) {
    fmt.Println("start httpbin post")
    Post()
}
