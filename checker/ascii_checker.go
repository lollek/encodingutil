package checker

import (
	"github.com/lollek/encodingutil/encoding"
)

type AsciiChecker struct {
	validates bool
}

func NewAsciiChecker() *AsciiChecker {
	return &AsciiChecker{
		validates: true,
	}
}

func (checker *AsciiChecker) Encoding() encoding.Encoding {
	return encoding.ASCII
}

func (checker *AsciiChecker) Probability() Probability {
	return MEDIUM
}

func (checker *AsciiChecker) Validates() bool {
	return checker.validates
}

func (checker *AsciiChecker) CheckNext(character byte) {
	checker.validates = checker.validates && character&0x80 == 0
}
