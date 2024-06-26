package checker

import (
	"github.com/tufin/oasdiff/diff"
	"github.com/tufin/oasdiff/load"
)

const (
	RequestParameterDefaultValueChangedId = "request-parameter-default-value-changed"
	RequestParameterDefaultValueAddedId   = "request-parameter-default-value-added"
	RequestParameterDefaultValueRemovedId = "request-parameter-default-value-removed"
)

func RequestParameterDefaultValueChangedCheck(diffReport *diff.Diff, operationsSources *diff.OperationsSourcesMap, config *Config) Changes {
	result := make(Changes, 0)
	if diffReport.PathsDiff == nil {
		return result
	}
	for path, pathItem := range diffReport.PathsDiff.Modified {
		if pathItem.OperationsDiff == nil {
			continue
		}
		for operation, operationItem := range pathItem.OperationsDiff.Modified {
			if operationItem.ParametersDiff == nil {
				continue
			}
			source := (*operationsSources)[operationItem.Revision]
			appendResultItem := func(messageId string, tag string, a ...any) {
				result = append(result, ApiChange{
					Id:          messageId,
					Level:       ERR,
					Args:        a,
					Operation:   operation,
					OperationId: operationItem.Revision.OperationID,
					Path:        path,
					Source:      load.NewSource(source),
					Tag:         tag,
				})
			}
			for paramLocation, paramDiffs := range operationItem.ParametersDiff.Modified {
				for paramName, paramDiff := range paramDiffs {

					baseParam := operationItem.Base.Parameters.GetByInAndName(paramLocation, paramName)
					if baseParam == nil || baseParam.Required {
						continue
					}

					revisionParam := operationItem.Revision.Parameters.GetByInAndName(paramLocation, paramName)
					if revisionParam == nil || revisionParam.Required {
						continue
					}

					if paramDiff.SchemaDiff == nil {
						continue
					}

					defaultValueDiff := paramDiff.SchemaDiff.DefaultDiff
					if defaultValueDiff.Empty() {
						continue
					}

					if defaultValueDiff.From == nil {
						appendResultItem(RequestParameterDefaultValueAddedId, AddParameterTag, paramLocation, paramName, defaultValueDiff.To)
					} else if defaultValueDiff.To == nil {
						appendResultItem(RequestParameterDefaultValueRemovedId, DelParameterTag, paramLocation, paramName, defaultValueDiff.From)
					} else {
						appendResultItem(RequestParameterDefaultValueChangedId, ModifyParameterTag, paramLocation, paramName, defaultValueDiff.From, defaultValueDiff.To)
					}
				}
			}
		}
	}
	return result
}
