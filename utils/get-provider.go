package utils

import (
	"errors"

	"github.com/spf13/pflag"
)

const (
	GURU = "guru"
)

func GetProvider(flags *pflag.FlagSet) (string, error) {
	isGuru, _ := flags.GetBool(GURU)
	if isGuru == true {
		return "guru", nil
	} else {
		return "", errors.New("Unknown provider.")
	} 
}