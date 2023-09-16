package spec

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func NewSpecFromFile(path string) (*Spec, error) {
	log.Info("Parsing spec from file", "path", path)
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %v", path, err)
	}

	s := &Spec{}
	if err := yaml.Unmarshal(b, s); err != nil {
		return nil, fmt.Errorf("could not unmarshal file contents: %v", err)
	}

	return s, nil
}
