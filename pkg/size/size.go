package size

import (
	"fmt"
	"os"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
)

func FileSize(filePath string) (err error) {
	fi, err := os.Stat(filePath)

	// return on any error
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		return err
	}
	fmt.Fprintf(os.Stdout, "Size: %dB\n", fi.Size())

	rf, err := bazel.RunfilesPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		return err
	}
	fmt.Fprintf(os.Stdout, "Runpath: %v\n", rf)

	return nil
}
