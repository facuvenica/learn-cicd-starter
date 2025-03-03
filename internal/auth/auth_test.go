package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"empty":      {input: "", want: []string{"", "no authorization header included"}},
		"correct":    {input: "ApiKey TestingKey", want: []string{"TestingKey", ""}},
		"no keyword": {input: "TestingKey", want: []string{"a", "malformed authorization header"}},
		"no apikey":  {input: "ApiKey", want: []string{"", "malformed authorization header"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			heads := http.Header{
				"Authorization": []string{tc.input},
			}

			msg := ""
			got, err := GetAPIKey(heads)
			if err != nil {
				t.Logf("Failed to get ApiKey, %s", err)
				msg = err.Error()
			}

			if !reflect.DeepEqual(tc.want, []string{got, msg}) {
				t.Fatalf("Test %s, expected: %#v, got: %#v", name, tc.want, []string{got, msg})
			}
		})
	}
}
