package dig

import (
    "context"
    "fmt"
    "testing"
)

func TestGet(t *testing.T) {
    for i := 0; i < 10; i++ {
        stone, _ := Get(context.Background(), "88888")
        if stone != nil {
            fmt.Printf("Hash:%s,前缀:%s,时间戳:%d,次数:%d\n", stone.Hash, stone.Prefix, stone.Time, stone.Ec)
        }
    }
}
