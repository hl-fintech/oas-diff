package formatters

import (
	"github.com/tufin/oasdiff/checker"
	"strings"
)

const (
	AddType       = "add"
	DelType       = "del"
	ModifyType    = "mod"
	DeprecateType = "dep"
)

type Endpoint struct {
	Path      string
	Operation string
	ApiStatus string
}

type ChangesByEndpoint map[Endpoint]*Changes

func GroupChanges(changes checker.Changes, l checker.Localizer) ChangesByEndpoint {

	apiChanges := ChangesByEndpoint{}

	for _, change := range changes {
		switch change.(type) {
		case checker.ApiChange:
			ep := Endpoint{Path: change.GetPath(), Operation: change.GetOperation(), ApiStatus: getApiStatus(change.GetTag())}
			if c, ok := apiChanges[ep]; ok {
				*c = append(*c, Change{
					IsBreaking: change.IsBreaking(),
					Text:       change.GetUncolorizedText(l),
					Tag:        change.GetTag(),
				})
			} else {
				apiChanges[ep] = &Changes{Change{
					IsBreaking: change.IsBreaking(),
					Text:       change.GetUncolorizedText(l),
					Tag:        change.GetTag(),
				}}
			}
		}
	}

	return apiChanges
}

func getApiStatus(tag string) string {
	if strings.HasPrefix(tag, "del-api") {
		return "deleted"
	}
	if strings.HasPrefix(tag, "deprecate-api") {
		return "deprecated"
	}
	if strings.HasPrefix(tag, "add-api") {
		return "added"
	}
	return "updated"
}
