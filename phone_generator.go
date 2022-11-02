package us_phone_generator

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrAreaCodeNotFound = errors.New("AREA_CODE_NOT_FOUND")

func GenerateByState(state string) string {
	var areaCode int
	var prefixCode int

	areaCodes := areaCodesByState(state)
	areaCodesCount := len(areaCodes)

	for i := 0; i < areaCodesCount; i++ {
		areaCode = randomFromSlice(areaCodes)
		prefixCodes, err := prefixCodesByAreaCode(areaCode, true)
		if err != nil {
			if errors.Is(err, ErrAreaCodeNotFound) && i+1 < areaCodesCount {
				areaCodes = removeFromSlice(areaCodes, areaCode)
				continue
			}

			panic("available prefix code not found by area code " + strconv.Itoa(areaCode))
		}

		prefixCode = randomFromSlice(prefixCodes)

		break
	}

	//fmt.Printf("+1-%d-%d-%d"+"\n", areaCode, prefixCode, randomInt(1111, 8888))
	return fmt.Sprintf("+1%d%d%d", areaCode, prefixCode, randomInt(1111, 8888))
}

func GenerateCodes() {
	npaCodes := generateFromNpa()
	generateFromFoneFinder(npaCodes)
}
