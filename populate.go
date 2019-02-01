package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func PopulateFromCSV(dir string) error {

	if _, err := os.Stat("./" + dir); os.IsNotExist(err) {
		return err
	}

	ctnts, err := ioutil.ReadFile(dir)
	if err != nil {
		return err
	}

	r := csv.NewReader(strings.NewReader(string(ctnts)))

	for {
		rec, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		fmt.Println(rec)

	}

	return nil
}
