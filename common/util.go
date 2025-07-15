package common

const BaseIndent = 2

func Indent(s string, n int) string {
	output := ""

	for range n {
		output += " "
	}

	output += s
	return output
}
