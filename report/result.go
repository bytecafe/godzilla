// Package report collects the scanning result, and make a report for it.
package report

import "go/token"

// ResultItem describes a vulnerability that the analyzer identified.
type ResultItem struct {
	Name        string
	Description string
	Flow        []*token.Position
}
