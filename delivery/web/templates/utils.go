package templates

import "fmt"

func CssId(id string) string {
	return fmt.Sprintf("#%s", id)
}
