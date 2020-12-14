// Code generated by kube-linter template codegen. DO NOT EDIT.
// +build !templatecodegen

package params

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"golang.stackrox.io/kube-linter/internal/check"
	"golang.stackrox.io/kube-linter/internal/templates/util"
)

var (
	// Use some imports in case they don't get used otherwise.
	_ = util.MustParseParameterDesc
	_ = fmt.Sprintf

	requirementsTypeParamDesc = util.MustParseParameterDesc(`{
	"Name": "requirementsType",
	"Type": "string",
	"Description": "The type of requirement. Use any to apply to both requests and limits.",
	"Examples": null,
	"Enum": [
		"request",
		"limit",
		"any"
	],
	"SubParameters": null,
	"ArrayElemType": "",
	"Required": true,
	"NoRegex": false,
	"NotNegatable": false,
	"XXXStructFieldName": "RequirementsType",
	"XXXIsPointer": false
}
`)

	lowerBoundMBParamDesc = util.MustParseParameterDesc(`{
	"Name": "lowerBoundMB",
	"Type": "integer",
	"Description": "The lower bound of the requirement (inclusive), specified as a number of MB.",
	"Examples": null,
	"Enum": null,
	"SubParameters": null,
	"ArrayElemType": "",
	"Required": false,
	"NoRegex": false,
	"NotNegatable": false,
	"XXXStructFieldName": "LowerBoundMB",
	"XXXIsPointer": false
}
`)

	upperBoundMBParamDesc = util.MustParseParameterDesc(`{
	"Name": "upperBoundMB",
	"Type": "integer",
	"Description": "The upper bound of the requirement (inclusive), specified as a number of MB. If not specified, it is treated as \"no upper bound\".",
	"Examples": null,
	"Enum": null,
	"SubParameters": null,
	"ArrayElemType": "",
	"Required": false,
	"NoRegex": false,
	"NotNegatable": false,
	"XXXStructFieldName": "UpperBoundMB",
	"XXXIsPointer": true
}
`)

	ParamDescs = []check.ParameterDesc{
		requirementsTypeParamDesc,
		lowerBoundMBParamDesc,
		upperBoundMBParamDesc,
	}
)

func (p *Params) Validate() error {
	var validationErrors []string
	if p.RequirementsType == "" {
		validationErrors = append(validationErrors, "required param requirementsType not found")
	}
	var found bool
	for _, allowedValue := range []string{
		"request",
		"limit",
		"any",
	}{
		if p.RequirementsType == allowedValue {
			found = true
			break
		}
	}
	if !found {
		validationErrors = append(validationErrors, fmt.Sprintf("param requirementsType has invalid value %q, must be one of [request limit any]", p.RequirementsType))
	}
	if len(validationErrors) > 0 {
		return errors.Errorf("invalid parameters: %s", strings.Join(validationErrors, ", "))
    }
	return nil
}

// ParseAndValidate instantiates a Params object out of the passed map[string]interface{},
// validates it, and returns it.
// The return type is interface{} to satisfy the type in the Template struct.
func ParseAndValidate(m map[string]interface{}) (interface{}, error) {
	var p Params
	if err := util.DecodeMapStructure(m, &p); err != nil {
		return nil, err
	}
	if err := p.Validate(); err != nil {
		return nil, err
	}
	return p, nil
}

// WrapInstantiateFunc is a convenience wrapper that wraps an untyped instantiate function
// into a typed one.
func WrapInstantiateFunc(f func(p Params) (check.Func, error)) func (interface{}) (check.Func, error) {
	return func(paramsInt interface{}) (check.Func, error) {
		return f(paramsInt.(Params))
	}
}
