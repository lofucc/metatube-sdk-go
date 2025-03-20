package openai

import (
	openai "github.com/zijiren233/openai-translator"

	"github.com/lofucc/metatube-sdk-go/translate"
)

var _ translate.Translator = (*OpenAI)(nil)

type OpenAI struct {
	APIKey  string `json:"openai-api-key"`
	BaseURL string `json:"openai-base-url"`
}

func (oa *OpenAI) Translate(q, source, target string) (result string, err error) {
	opts := []openai.Option{openai.WithFrom(source)}
	if oa.BaseURL != "" {
		opts = append(opts, openai.WithBaseURL(oa.BaseURL))
	}
	return openai.Translate(q, target, oa.APIKey, opts...)
}

func init() {
	translate.Register(&OpenAI{})
}
