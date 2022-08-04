package helpers

import "fmt"

// ConRoute is a
func ConRoute(base string, uri string) string {
	return fmt.Sprintf("%s%s", base, uri)
}
