package archive

import (
	"archive/tar"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/afero"
)

// Writes Tar
func WriteTar(filePaths []string, destination string, directory string) (target string, err error) {
	fs := afero.NewOsFs()
	afs := &afero.Afero{Fs: fs}

	// TODO: verify paths
	//if _, err := fs.Stat(filePath); err != nil {
	//	return "", fmt.Errorf("unable to package files - %v", err.Error())
	//}

	file, err := afs.Create(destination)
	if err != nil {
		return "", err
	}
	defer func() {
		if ferr := file.Close(); ferr != nil {
			err = ferr
		}
	}()

	err = TarPaths(afs, filePaths, file, directory)
	return target, err
}

// Tars Paths
func TarPaths(fs afero.Fs, filePaths []string, w io.Writer, directory string) (err error) {
	tw := tar.NewWriter(w)
	defer tw.Close()

	for _, file := range filePaths {
		fi, err := fs.Stat(file)

		// return on any error
		if err != nil {
			return err
		}

		// return on non-regular files.  We don't add directories without files and symlinks
		if !fi.Mode().IsRegular() {
			return nil
		}

		// create a new dir/file header
		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return err
		}

		// update the name to correctly reflect the desired destination when untaring
		header.Name = strings.TrimPrefix(directory+string(filepath.Separator)+fi.Name(), string(filepath.Separator))

		// change certain header metadata to make the build reproducible
		header.ModTime = time.Time{}
		header.Uid = 0
		header.Gid = 0
		header.Uname = ""
		header.Gname = ""
		header.Format = tar.FormatPAX

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// open files for taring
		f, err := fs.Open(file)
		if err != nil {
			return err
		}

		defer func() {
			if ferr := f.Close(); ferr != nil {
				err = ferr
			}
		}()

		// copy file data into tar writer
		if _, err := io.Copy(tw, f); err != nil {
			return err
		}

	}
	return err
}
