package license

import (
	"fmt"

	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate"
)

type LicenseBuilder interface {
	kpbtemplate.Builder
	SourceFileHeader() string
}

func ChooseLicenseBuilder(licenseKind string, year uint, author string) (LicenseBuilder, error) {
	switch licenseKind {
	case MITLicense:
		data := &MITLicenseData{Year: year, Author: author}
		return NewMITBuilder("LICENSE", data), nil
	}

	return nil, fmt.Errorf("unsupported license kind %s", licenseKind)
}
