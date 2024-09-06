package translit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransliterateEnRu(t *testing.T) {
	type TestCase struct {
		Input    string
		Expected string
	}

	cases := []TestCase{
		{
			Input:    "Chance",
			Expected: "Чансе",
		},
		{
			Input:    "Shanghai",
			Expected: "Шангхаи",
		},
		{
			Input:    "Shenyang",
			Expected: "Шенянг",
		},
		{
			Input:    "Shendzhen",
			Expected: "Шенджен",
		},
		{
			Input:    "Shendzhen Китайский",
			Expected: "Шенджен Китайский",
		},
	}

	for i := range cases {
		assert.Equal(t, cases[i].Expected, TransliterateEnRu(cases[i].Input))
	}
}
