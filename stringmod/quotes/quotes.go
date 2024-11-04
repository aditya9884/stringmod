// for the quotes

package quotes

import (
	"github.com/hackebrot/turtle"
)

func Getemoji(name string) string {
	emoji, ok := turtle.Emojis[name]
	if !ok {
		return ""
	}
	return emoji.Char
}
