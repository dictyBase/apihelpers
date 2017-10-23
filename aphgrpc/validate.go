package aphgrpc

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/dictyBase/apihelpers/aphcollection"
	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	"google.golang.org/grpc/metadata"
)

var re = regexp.MustCompile(`(\w+)(\=\=|\!\=|\=\@|\!\@)(\w+)(\,|\;)?`)

// JSONAPIParams is a container for various JSON API query parameters
type JSONAPIParams struct {
	// contain include query paramters
	Includes []string
	// contain fields query paramters
	Fields []string
	// check for presence of fields parameters
	HasFields bool
	// check for presence of include parameters
	HasIncludes bool
	// check for presence of filter parameters
	HasFilter bool
	// slice of filters
	Filter []*APIFilter
}

// APIFilter is a container for filter parameters
type APIFilter struct {
	// Attribute of the resource on which the filter will be applied
	Attribute string
	// Type of filter for matching or exclusion
	Operator string
	// The value to match or exclude
	Expression string
	//
	Logic string
}

func hasInclude(r *jsonapi.GetRequest) bool {
	if len(r.Include) > 0 {
		return true
	}
	return false
}

func hasFields(r *jsonapi.GetRequest) bool {
	if len(r.Fields) > 0 {
		return true
	}
	return false
}

func hasListInclude(r *jsonapi.ListRequest) bool {
	if len(r.Include) > 0 {
		return true
	}
	return false
}

func hasListFields(r *jsonapi.ListRequest) bool {
	if len(r.Fields) > 0 {
		return true
	}
	return false
}

func hasFilter(r *jsonapi.ListRequest) bool {
	if len(r.Filter) > 0 {
		return true
	}
	return false
}

// ValidateAndParseListParams validate and parse the JSON API include, fields, filter parameters
func ValidateAndParseListParams(jsapi JSONAPIAllowedParams, r *jsonapi.ListRequest) (*JSONAPIParams, metadata.MD, error) {
	params := &JSONAPIParams{
		HasFields:   false,
		HasIncludes: false,
		HasFilter:   false,
	}
	if hasListInclude(r) {
		if strings.Contains(r.Include, ",") {
			params.Includes = strings.Split(r.Include, ",")
		} else {
			params.Includes = []string{r.Include}
		}
		for _, v := range params.Includes {
			if !aphcollection.Contains(jsapi.AllowedInclude(), v) {
				return params, ErrIncludeParam, fmt.Errorf("include %s relationship is not allowed", v)
			}
		}
		params.HasIncludes = true
	}

	if hasListFields(r) {
		if strings.Contains(r.Fields, ",") {
			params.Fields = strings.Split(r.Fields, ",")
		} else {
			params.Fields = []string{r.Fields}
		}
		for _, v := range params.Fields {
			if !aphcollection.Contains(jsapi.AllowedFields(), v) {
				return params, ErrFields, fmt.Errorf("%s fields attribute is not allowed", v)
			}
		}
		params.HasFields = true
	}
	if HasFilter(r) {
		m := re.FindAllStringSubmatch(r.Filter)
		if len(m) > 0 {
			var filters []*APIFilter
			for _, n := range m {
				if !aphcollection.Contains(jsapi.AllowedFilter(), n[1]) {
					return params, ErrFilterParam, fmt.Errorf("%s filter attribute is not allowed", n[1])
				}
				f := &APIFilter{
					Attribute:  n[1],
					Operator:   n[2],
					Expression: n[3],
					Logic:      n[4],
				}
				filters = append(filters, f)
			}
			params.HasFilter = true
			params.Filter = filters
		}
	}
	return params, metadata.Pairs("errors", "none"), nil
}

// ValidateAndParseGetParams validate and parse the JSON API include and fields parameters
// that are used for singular resources
func ValidateAndParseGetParams(jsapi JSONAPIAllowedParams, r *jsonapi.GetRequest) (*JSONAPIParams, metadata.MD, error) {
	params := &JSONAPIParams{}
	if hasInclude(r) {
		if strings.Contains(r.Include, ",") {
			params.Includes = strings.Split(r.Include, ",")
		} else {
			params.Includes = []string{r.Include}
		}
		for _, v := range params.Includes {
			if !aphcollection.Contains(jsapi.AllowedInclude(), v) {
				return params, ErrIncludeParam, fmt.Errorf("include %s relationship is not allowed", v)
			}
		}
	} else {
		params.HasIncludes = false
	}

	if hasFields(r) {
		if strings.Contains(r.Fields, ",") {
			params.Fields = strings.Split(r.Fields, ",")
		} else {
			params.Fields = []string{r.Fields}
		}
		for _, v := range params.Fields {
			if !aphcollection.Contains(jsapi.AllowedFields(), v) {
				return params, ErrFilterParam, fmt.Errorf("%s value in fields is not allowed", v)
			}
		}
	} else {
		params.HasFields = true
	}
	return params, metadata.Pairs("errors", "none"), nil
}
