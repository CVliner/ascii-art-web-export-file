package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func Generate(args string, filename string) (string, bool) {
	src, err := os.ReadFile("asciiart/" + filename + ".txt")
	if err != nil {
		fmt.Println("Bad request", err.Error())
		return "2", false
		// os.Exit(1)
	}
	for _, letter := range args {
		if (letter != '\n' && letter != '\r') && (letter < 32 || letter > 127) {
			return "1", false
		}
	}
	var ans string
	data := strings.Split(args, "\n")
	strarr := strings.Split(string(src), "\n")

	for k := range data {
		if data[k] == "" {
			ans = ans + "\n"
			continue
		}
		for i := 1; i < 9; i++ {
			for j := 0; j < len(data[k]); j++ {
				num := data[k][j] - 32
				line := int(num)*9 + i

				// Check if the line is within bounds before accessing strarr
				if line >= 0 && line < len(strarr) {
					ans = ans + strings.TrimSuffix(strarr[line], "\r")
				} else {
					// Handle the out-of-range error (e.g., append a default value or skip)
					ans = ans + ""
				}
			}
			ans = ans + "\n"
		}
	}

	return ans, true
}
