package apiutils

import "errors"

type Dependencies map[string]interface{}

func ValidateDependencies(depMap Dependencies) error {
	if depMap == nil {
		return errors.New("dependencies map is nil")
	}

	for dependency, source := range depMap {
		if source == nil {
			return errors.New(dependency + " is nil")
		}
	}

	return nil
}
