package common

import (
	"github.com/dustin/go-humanize"
)

// FormatAmount returns string-formatted of the given float64 number.
func FormatAmount(amt float64) string {
	return humanize.FormatFloat("#,###.##", amt)
}
