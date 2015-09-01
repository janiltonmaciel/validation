package validation

import (
	"regexp"

	"gitlab.globoi.com/bastian/falkor/errors"
)

const formatUUID = "^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$"

var reUUID = regexp.MustCompile(formatUUID)

type UUIDValidator struct{}

func (u UUIDValidator) Validate(key string, value string) error {
	if isOK := UUIDMath(value); !isOK {
		return errors.HttpParamInvalidError("%s invalid param", key)
	}
	return nil
}

func UUIDMath(value string) bool {
	return reUUID.MatchString(value)
}
