package main

import (
    "os"
    "fmt"
    "encoding/json"
    "github.com/lvshaco/bencode"
)

func assert(err interface{}) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func main() {
    if len(os.Args) < 2 {
        assert(fmt.Sprintf("Usage: %s bt_file", os.Args[0]))
    }
    f, err := os.Open(os.Args[1])
    defer f.Close()
    assert(err)

    offset, err := f.Seek(0, 2)
    assert(err)

    fmt.Println("bt file size:", offset)

    f.Seek(0, 0)
    s := make([]byte, offset)
    _, err = f.Read(s)
    assert(err)

    r, err := bencode.Decode(s)
    assert(err)

    o, err := json.MarshalIndent(r, "", " ")
    assert(err)

    fmt.Println(string(o))
}
