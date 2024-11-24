package services

import (
	"regexp"

	"github.com/sirupsen/logrus"
)

type ZipCodeService interface {
	IsValidCEP(cep string) bool
}

type ZipCodeValidator struct{}

func NewZipCodeService() ZipCodeService {
	return &ZipCodeValidator{}
}

func (z *ZipCodeValidator) IsValidCEP(cep string) bool {
	matched, err := regexp.MatchString(`^\d{8}$`, cep)
	if err != nil {
		logrus.Error("Error while validating CEP: ", err)

		return false
	}

	return matched
}
