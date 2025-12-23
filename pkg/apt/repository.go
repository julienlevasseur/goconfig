package apt

import (
	"fmt"
	"log"

	"github.com/julienlevasseur/goconfig/pkg/file"
)

func NewSource(srcURL, name string) error {
	dstPath := fmt.Sprintf("/etc/apt/sources.list.d/%v.sources", name)

	if !file.Exists(dstPath) {
		log.Printf("[apt] Download source file %v\n", srcURL)
		err := file.Download(
			srcURL,
			fmt.Sprintf("/etc/apt/sources.list.d/%v.sources", name),
		)
		if err != nil {
			return err
		}

		err = Update(true)
		if err != nil {
			return err
		}
	}

	return nil
}
