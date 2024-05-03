package random

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Code(n int) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%4v", rand.Intn(int(math.Pow10(n))))
}
