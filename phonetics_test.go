package gophonetics

import "testing"

type phonetic struct {
	str                          string
	phonCode                     string
	phonCodeWithoutDeduplication string
}

var phonetics = []phonetic{
	{str: "Müller-Lüdenscheidt", phonCode: "65752682", phonCodeWithoutDeduplication: "65575268822"},
	{str: "Wikipedia", phonCode: "3412", phonCodeWithoutDeduplication: "3412"},
	{str: "Breschnew", phonCode: "17863", phonCodeWithoutDeduplication: "178863"},
	{str: "c", phonCode: "8", phonCodeWithoutDeduplication: "8"},
	{str: "ca", phonCode: "4", phonCodeWithoutDeduplication: "4"},
	{str: "dc", phonCode: "8", phonCodeWithoutDeduplication: "88"},
	{str: "", phonCode: "", phonCodeWithoutDeduplication: ""},
	{str: "Ł", phonCode: "5", phonCodeWithoutDeduplication: "5"},
}

func TestPhonetics(t *testing.T) {
	for _, p := range phonetics {
		code := NewPhoneticCode(p.str)

		if code != p.phonCode {
			t.Errorf("Error: expected %v, got %v",
				p.phonCode, code)
		}
	}
}
func TestPhoneticsWithoutDeduplication(t *testing.T) {
	for _, p := range phonetics {
		code := NewPhoneticCodeWithoutDeduplication(p.str)

		if code != p.phonCodeWithoutDeduplication {
			t.Errorf("Error: expected %v, got %v",
				p.phonCodeWithoutDeduplication, code)
		}
	}
}

// --------------------------------------------------------------------------

type duplicate struct {
	strFull     string
	strExpected string
}

var duplicates = []duplicate{
	{strFull: "65575268822", strExpected: "65752682"},
	{strFull: "6557526882", strExpected: "65752682"},
	{strFull: "1234567899", strExpected: "123456789"},
	{strFull: "1123456789", strExpected: "123456789"},
	{strFull: "1111111111", strExpected: "1"},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, dup := range duplicates {
		strRem := removeDuplicates(dup.strFull)

		if strRem != dup.strExpected {
			t.Errorf("Error: expected %v, got %v",
				dup.strExpected, strRem)
		}
	}
}
