// Greeter computes a greeting message

package greeter

import (
	"fmt"
	"strconv"
	"strings"
)

func greeter1(n string, a int) string {
	const format = "Hello, %s! You are %d old today..."
	return fmt.Sprintf(format, n, a)
}

func greeter2(n string, a int) string {
	return "Hello, " + n + "! You are " + strconv.Itoa(a) + " old today..."
}

func greeter3(n string, a int) string {
	var b strings.Builder
	b.Grow(39)
	b.WriteString("Hello, ")
	b.WriteString(n)
	b.WriteString("! You are ")
	b.WriteString(strconv.Itoa(a))
	b.WriteString(" old today...")
	return b.String()
}

func greeter4(n string, a int) string {
	b := make([]byte, 0, 39)
	b = append(b, "Hello, "...)
	b = append(b, n...)
	b = append(b, "! You are "...)
	b = append(b, strconv.Itoa(a)...)
	b = append(b, " old today..."...)
	return string(b)
}
