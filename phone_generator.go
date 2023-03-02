package us_phone_generator

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrAreaCodeNotFound = errors.New("AREA_CODE_NOT_FOUND")
var ErrStateNotFound = errors.New("STATE_NOT_FOUND")

func GenerateByState(state string) (string, error) {
	var areaCode int
	var prefixCode int

	areaCodes, err := areaCodesByState(state)
	if err != nil {
		return "", fmt.Errorf("searching area code by state error: %w", err)
	}

	areaCodesCount := len(areaCodes)

	for i := 0; i < areaCodesCount; i++ {
		areaCode = randomFromSlice(areaCodes)
		prefixCodes, err := prefixCodesByAreaCode(areaCode, true)
		if err != nil {
			if errors.Is(err, ErrAreaCodeNotFound) && i+1 < areaCodesCount {
				areaCodes = removeFromSlice(areaCodes, areaCode)
				continue
			}

			return "", fmt.Errorf("available prefix code %s not found by area code: %w", strconv.Itoa(areaCode), err)
		}

		prefixCode = randomFromSlice(prefixCodes)

		break
	}

	//fmt.Printf("+1-%d-%d-%d"+"\n", areaCode, prefixCode, randomInt(1111, 8888))
	return fmt.Sprintf("+1%d%d%d", areaCode, prefixCode, randomInt(1111, 8888)), nil
}

func GenerateCodes() {
	npaCodes := generateFromNpa()
	generateFromFoneFinder(npaCodes)
}
