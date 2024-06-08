package pretty

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	formatting "github.com/tom773/gopolish/pkg"
)

func Pretty(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	counter := 0
	full := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		full = append(full, line)
		counter++
	}
	new := replace(full)
	test := generateTestStatus(new)
	finalOutput(test, new, w)

	return scanner.Err()
}

type TestStatus struct {
	packageName string
	passed      bool
	failed      bool
}

func generateTestStatus(full []string) TestStatus {

	test := &TestStatus{}

	last := len(full) - 1
	re := regexp.MustCompile(`\b\w+\b\s{2,6}(\S+)`)
	matches := re.FindAllStringSubmatch(full[last], -1)
	if len(matches) > 0 {
		test.packageName = matches[0][1]
	}
	if strings.Contains(full[last], "FAIL") {
		test.failed = true
	}
	if strings.Contains(full[last], "ok") {
		test.passed = true
	}
	return *test
}
func replace(full []string) []string {

	new := []string{}
	for _, l := range full {
		upd := strings.ReplaceAll(l, "===", "")
		upd = strings.ReplaceAll(upd, "---", "")
		upd = strings.ReplaceAll(upd, "RUN", formatting.Colors.Cyan+formatting.Icons.Subtest+formatting.Colors.CadetBlue+"Running"+formatting.Colors.Reset)
		upd = strings.ReplaceAll(upd, "PASS", formatting.Colors.Green+formatting.Icons.Success+"Passed"+formatting.Colors.Reset)
		upd = strings.ReplaceAll(upd, "FAIL", formatting.Colors.LightPink+formatting.Icons.Error+"Fail"+formatting.Colors.Reset)
		upd = strings.ReplaceAll(upd, "ERROR", formatting.Colors.Red+formatting.Icons.Error+"Error"+formatting.Colors.Reset)
		upd = strings.ReplaceAll(upd, "SKIP", formatting.Colors.Magenta+formatting.Icons.Output+"Skipped"+formatting.Colors.Reset)
		if (strings.Contains(l, "PASS") || strings.Contains(l, "FAIL")) && len(l) == 4 {
			upd = ""
		}
		new = append(new, upd)
	}
	return new
}

func fpretty(line string) string {
	keywords := []string{"Running", "Passed", "Fail", "Error", "Skipped", "Testing"}
	for _, k := range keywords {
		if strings.Contains(line, k) {
			return line
		}
	}
	if strings.Contains(line, "ok") {
		line = formatting.Colors.Green + formatting.Icons.Success + "All Tests Passed" + formatting.Colors.Reset
		return line
	}
	if strings.Contains(line, "exit status") {
		line = formatting.Colors.LightPink + "Exit Status: 1" + formatting.Colors.Reset
		return line
	}
	if line == "" {
		line = "=========================="
		return line
	}
	line = "\t" + formatting.Colors.Yellow + formatting.Icons.Output + "Output:\u00a0" + formatting.Colors.Reset + line
	return line
}

func finalOutput(test TestStatus, full []string, w io.Writer) {
	trimSpaces := func(r rune) bool {
		return r == ' '
	}
	fmt.Fprintf(w, "\n%s%s Testing Package: %s%s\n\n", formatting.Colors.Brown, formatting.Icons.Pkg, test.packageName, formatting.Colors.Reset)
	for _, l := range full {
		l = fpretty(l)
		l = strings.TrimLeftFunc(l, trimSpaces)
		l += "\n"
		fmt.Fprintln(w, l)
	}
}
