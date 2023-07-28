package initialiser

import (
	"bytes"
	"fmt"
	"html/template"
	"mime"
	"net/http"
)

func ConfigLoader(location string) (http.Handler, error) {
	content, err := Assets.ReadFile("swagger-initializer.js")
	if err != nil {
		return nil, fmt.Errorf("failed to read file content")
	}

	tpl, err := template.New("js").Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	data := map[string]string{
		"Location": location,
	}
	buffer := bytes.NewBuffer(nil)
	err = tpl.Execute(buffer, data)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template data: %w", err)
	}
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctype := mime.TypeByExtension("js")
		writer.Header().Set("Content-Type", ctype)
		writer.Header().Set("Content-Length", fmt.Sprintf("%d", buffer.Len()))
		writer.WriteHeader(200)
		_, _ = writer.Write(buffer.Bytes())
	}), nil
}
