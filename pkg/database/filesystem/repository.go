package filesystem

import (
	"encoding/json"
	"errors"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/database"
	"log"
	"os"
)

type DataRepository struct {
	path string
}

func NewDataRepository(path string) database.DataRepository {
	return DataRepository{path: path}
}

func (r DataRepository) Get() (*database.Data, error) {
	b, err := os.ReadFile(r.path)

	if err != nil {
		log.Printf("os.ReadFile no such file: %s", r.path)
	}

	if len(b) <= 0 {
		err := errors.New("os.ReadFile: data.json is empty")
		log.Printf(err.Error())
		return nil, err
	}

	data, err := formatData(b)

	return data, err
}

func (r DataRepository) Store(d *database.Data) error {
	f, err := os.OpenFile(r.path, os.O_RDWR, 0644)

	if err != nil {
		log.Printf("could not os.OpenFile: %s", err)
		return err
	}

	defer f.Close()

	b, err := json.Marshal(d)

	if err != nil {
		log.Printf("could not json.Marshal: %s", err)
		return err
	}

	if _, err := f.WriteAt(b, 0); err != nil {
		log.Printf("could not f.WriteAt: %s", err)
		return err
	}

	return nil
}

func formatData(b []byte) (*database.Data, error) {
	d := database.Data{}

	if err := json.Unmarshal(b, &d); err != nil {
		log.Printf("json.Unmarshal database.Data from data.json: %s", err)
		return nil, err
	}

	return &d, nil
}
