package openai

import (
	"fmt"
	"os"
	"testing"
)

func TestOpenaiTranslate(t *testing.T) {
	for _, unit := range []struct {
		text, from, to string
	}{
		{`Oh yeah! I'm a translator!`, "", "zh-CN"},
		{`Oh yeah! I'm a translator!`, "", "zh-TW"},
		{`Oh yeah! I'm a translator!`, "", "ja"},
		{`Oh yeah! I'm a translator!`, "", "de"},
		{`Oh yeah! I'm a translator!`, "", "fr"},
	} {
		fmt.Println("--------------------------")
		result, err := (&OpenAI{
			APIKey:  os.Getenv("OPENAI_API_KEY"),
			BaseURL: os.Getenv("OPENAI_API_URL"),
			Model:   os.Getenv("OPENAI_API_MODEL"),
		}).Translate(unit.text, unit.from, unit.to)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}
}
