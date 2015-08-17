/*Package expand is a simple utility for expanding some variables in path
its base on https://github.com/mitchellh/go-homedir , with some more functionality
also there is a compiler build tag for getting home dir, not checking the
GOOS variable.

*/
package expand

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`^(\$(PWD|HOME)|~)`)

// HomeDir return the home directory of he current user without need to use CGO
// so the cross compiling is not a pain just for getting a home directory.
func HomeDir() (string, error) {
	// A simple trick to get correct version compiled in the code
	return getHomeDir()
}

// Pwd return the path to the current executable directory
func Pwd() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return dir, nil
}

// Path is a function to expand path with variables into real address.
// supported variables are $HOME, $PWD and ~ if and only if its at
// the begin of string.
func Path(in string) (string, error) {
	tok := re.FindString(in)

	var dir string
	var err error
	if tok == "~" || tok == "$HOME" {
		dir, err = getHomeDir()
		if err != nil {
			return "", err
		}
	}

	if tok == "$PWD" {
		dir, err = Pwd()
		if err != nil {
			return "", err
		}
	}

	// There is no way thet the regexp pass another option.
	return strings.Replace(in, tok, dir, 1), nil
}
