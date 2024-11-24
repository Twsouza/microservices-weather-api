package services

import (
	"context"
	"regexp"
	"servicea/tracer"

	"github.com/sirupsen/logrus"
)

type ZipCodeService interface {
	IsValidCEP(ctx context.Context, cep string) bool
}

type ZipCodeValidator struct{}

func NewZipCodeService() ZipCodeService {
	return &ZipCodeValidator{}
}

func (z *ZipCodeValidator) IsValidCEP(ctx context.Context, cep string) bool {
	_, span := tracer.Tracer.Start(ctx, "ZipCodeValidator.IsValidCEP")
	defer span.End()

	matched, err := regexp.MatchString(`^\d{8}$`, cep)
	if err != nil {
		logrus.Error("Error while validating CEP: ", err)

		return false
	}

	return matched
}
