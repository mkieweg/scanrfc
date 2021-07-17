package scanrfc

import (
	"os"

	"github.com/spf13/viper"
)

func WriteFile(entries [][]byte) error {
	filename := viper.GetString("bib-file")

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, entry := range entries {
		_, err := f.Write(entry)
		if err != nil {
			return err
		}
	}
	return nil
}
