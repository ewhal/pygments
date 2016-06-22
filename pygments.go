// Pygments wrapper for golang. Pygments is a syntax highlighter

package pygments

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	bin = "pygmentize"
)

func Binary(path string) {
	bin = path
}

func Which() string {
	return bin
}

func Highlight(code string, lexer string, format string, options string, enc string) (string, error) {

	if _, err := exec.LookPath(bin); err != nil {
		fmt.Println("You do not have " + bin + " installed!")
		os.Exit(0)
	}

	cmd := exec.Command(bin, "-l"+lexer, "-f"+format, "-O"+options+"encoding="+enc)
	cmd.Stdin = strings.NewReader(code)

	var out bytes.Buffer
	cmd.Stdout = &out

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return out.String(), nil
}
