package render

import "github.com/nas-fish/common/json"

func Render(object any) string {
	return json.MarshalWithoutError[string](object)
}
