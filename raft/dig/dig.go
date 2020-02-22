package dig

import (
    "context"
    "crypto/sha1"
    "fmt"
    "math"
    "time"
)

type Stone struct {
    Time   uint64
    Ec     uint32
    Hash   string
    Prefix string
}

func Check(stone *Stone) bool {
    return getHash(fmt.Sprintf("%d__%d", stone.Time, stone.Ec)) == stone.Hash && stone.Hash[:len(stone.Prefix)] == stone.Prefix
}

func Get(ctx context.Context, prefix string) (*Stone, bool) {
    t := time.Now().UnixNano()
    i := uint32(32)
    hash := ""
    for {
        select {
        case <-ctx.Done():
            return nil, false
        default:
            if i >= math.MaxUint32-10 {
                t = time.Now().UnixNano()
            }
            hash = getHash(fmt.Sprintf("%d__%d", t, i))
            if hash[:len(prefix)] == prefix {
                return &Stone{
                    Time:   uint64(t),
                    Ec:     i,
                    Hash:   hash,
                    Prefix: prefix,
                }, true
            } else {
                i++
            }
        }
    }
}

func getHash(s string) string {
    _sha := sha1.New()
    _sha.Write([]byte(s))
    hash := _sha.Sum(nil)
    return fmt.Sprintf("%x", hash)
}
