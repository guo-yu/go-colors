package colors

import "strings"

func Yellow(color string) string {
  return strings.Join([]string{"\x1B[33m", color, "\x1B[39m"}, "")
}