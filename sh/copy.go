package sh

import (
	"fmt"
	"io"
	"os"
)

// Copy copies the given src file to dst. dst and src must be filepaths, not
// directories. Existing files will be overwritten.
func Copy(dst, src string) (err error) {
	fi, err := os.Stat(dst)
	if !os.IsNotExist(err) && !os.IsExist(err) {
		// If the destination file already exists, that's ok. It will be overwritten.
		// If the destination file doesn't exist, that's ok. It will be created.
		// If it's any other kind of error, bail.
		return err
	}

	if err == nil && fi.IsDir() {
		return fmt.Errorf("sh.Copy() destination must be a file, not a directory. %q is a directory", dst)
	}

	outf, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := outf.Close(); err == nil {
			err = cerr
		}
	}()

	inf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := inf.Close(); err == nil {
			err = cerr
		}
	}()

	_, err = io.Copy(outf, inf)
	return err
}
