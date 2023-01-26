package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/invopop/jsonschema"
	"github.com/spf13/cobra"

	"github.com/authelia/authelia/v4/internal/configuration/schema"
)

func newDocsJSONSchemaCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "json-schema",
		Short: "Generate docs JSON schema",
		RunE:  docsJSONSchemaRunE,

		DisableAutoGenTag: true,
	}

	return cmd
}

func docsJSONSchemaRunE(cmd *cobra.Command, args []string) (err error) {
	var schemaDir string

	if schemaDir, err = getPFlagPath(cmd.Flags(), cmdFlagRoot, cmdFlagDirSchema); err != nil {
		return err
	}

	r := &jsonschema.Reflector{
		RequiredFromJSONSchemaTags: true,
		Mapper:                     mapper,
	}

	if runtime.GOOS == "windows" {
		mapComments := map[string]string{}

		if err = jsonschema.ExtractGoComments("github.com/authelia/authelia/v4", schemaDir, mapComments); err != nil {
			return err
		}

		if r.CommentMap == nil {
			r.CommentMap = map[string]string{}
		}

		for key, comment := range mapComments {
			r.CommentMap[strings.ReplaceAll(key, `\`, `/`)] = comment
		}
	} else {
		if err = r.AddGoComments("github.com/authelia/authelia/v4", schemaDir); err != nil {
			return err
		}
	}

	s := r.Reflect(&schema.Configuration{})
	s.ID = "https://github.com/authelia/authelia/v4"
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	return nil
}

func mapper(t reflect.Type) *jsonschema.Schema {
	switch t.String() {
	case "regexp.Regexp", "*regexp.Regexp":
		return &jsonschema.Schema{
			Type: "string",
		}
	}
	return nil
}
