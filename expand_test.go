package expand

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExpand(t *testing.T) {
	Convey("Get home test", t, func() {
		u, err := user.Current()
		So(err, ShouldBeNil)

		dir, err := HomeDir()
		So(err, ShouldBeNil)

		So(dir, ShouldEqual, u.HomeDir)

		hbackup := os.Getenv("HOME")
		os.Setenv("HOME", "")

		dir, err = HomeDir()
		So(err, ShouldBeNil)

		So(dir, ShouldEqual, u.HomeDir)

		// A nasty hack :)
		pbackup := os.Getenv("PATH")
		os.Setenv("PATH", "")

		dir, err = HomeDir()
		So(err, ShouldNotBeNil)

		os.Setenv("HOME", hbackup)
		os.Setenv("PATH", pbackup)
	})

	Convey("get pwd", t, func() {
		pwd, err := Pwd()
		So(err, ShouldBeNil)
		So(pwd, ShouldNotBeEmpty)
	})

	Convey("expand", t, func() {
		pwd, err := Pwd()
		So(err, ShouldBeNil)

		u, err := user.Current()
		So(err, ShouldBeNil)

		p, err := Path("~")
		So(err, ShouldBeNil)
		So(p, ShouldEqual, u.HomeDir)

		p, err = Path("~/a/b/c")
		So(err, ShouldBeNil)
		So(p, ShouldEqual, filepath.Join(u.HomeDir, "a", "b", "c"))

		p, err = Path("$PWD")
		So(err, ShouldBeNil)
		So(p, ShouldEqual, pwd)

		p, err = Path("$PWD/a/b/c")
		So(err, ShouldBeNil)
		So(p, ShouldEqual, filepath.Join(pwd, "a", "b", "c"))

		p, err = Path("$HOME")
		So(err, ShouldBeNil)
		So(p, ShouldEqual, u.HomeDir)

		p, err = Path("$HOME/a/b/c")
		So(err, ShouldBeNil)
		So(p, ShouldEqual, filepath.Join(u.HomeDir, "a", "b", "c"))
	})
}
