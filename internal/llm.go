package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"text/template"

	"github.com/invopop/jsonschema"
	"github.com/ollama/ollama/api"
)

const systemTemplateString = "You are a JSON API with the following response jsonschema:"

func codeSnippet(code string, langs ...string) string {
	lang := ""
	if len(langs) > 0 {
		lang = langs[0]
	}
	return fmt.Sprintf("```%s\n%s\n```\n", lang, code)
}

func templatePrompt(tmpl string, i interface{}) (string, error) {
	t, err := template.New("prompt").Parse(tmpl)
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, i); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func jsonSystemPrompt(i interface{}) (string, error) {
	schema, err := json.MarshalIndent(jsonschema.Reflect(i), "", "  ")
	return systemTemplateString + "\n" + codeSnippet(string(schema)), err
}

type Model interface {
	GenerateJSON(ctx context.Context, tmpl string, input interface{}, output interface{}) error
	Generate(ctx context.Context, tmpl string, input interface{}) (string, error)
}

type ollama struct {
	model  string
	client *api.Client
}

func DefaultModel() (Model, error) {
	client, err := api.ClientFromEnvironment()
	return ollama{model: "llama3", client: client}, err
}

func (o ollama) generate(ctx context.Context, req *api.GenerateRequest) (string, error) {
	var res string
	err := o.client.Generate(ctx, req, func(r api.GenerateResponse) error {
		res += r.Response
		return nil
	})
	return res, err
}

func (o ollama) Generate(ctx context.Context, tmpl string, input interface{}) (string, error) {
	prompt, err := templatePrompt(tmpl, input)
	if err != nil {
		return "", err
	}
	return o.generate(ctx, &api.GenerateRequest{Prompt: prompt, Model: o.model})
}

func (o ollama) GenerateJSON(
	ctx context.Context,
	tmpl string,
	input interface{}, output interface{},
) error {
	prompt, err := templatePrompt(tmpl, input)
	if err != nil {
		return err
	}
	system, err := jsonSystemPrompt(output)
	if err != nil {
		return err
	}
	res, err := o.generate(ctx, &api.GenerateRequest{Prompt: prompt, System: system, Model: o.model})
	code := extractCode(res)
	if err := json.Unmarshal([]byte(code), output); err != nil {
		return fmt.Errorf("unable to unmarshal: %q", res)
	}
	return nil
}
