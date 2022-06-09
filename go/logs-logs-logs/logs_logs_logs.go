package logs

import "fmt"

func Application(log string) string {
	for _, r := range log {
		switch r {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

func Replace(log string, oldRune, newRune rune) string {
	r := []rune(log)
	for i, v := range r {
		if v == oldRune {
			r[i] = newRune
		}
	}
	return fmt.Sprintf(string(r))
}

func WithinLimit(log string, limit int) bool {
	r := []rune(log)
  return len(r) <= limit
}
