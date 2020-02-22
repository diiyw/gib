package strings

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Format(s string, types ...Type) string {
	for _, t := range types {
		s = t(s)
	}
	return s
}

// Has reports whether substr is within s.
func Has(s, sub string) bool {
	return strings.Contains(s, sub)
}

func GenCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
