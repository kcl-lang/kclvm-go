package gen

import (
	"encoding/json"
	"io"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"kcl-lang.io/kcl-go/pkg/3rdparty/jsonschema"
	"kcl-lang.io/kcl-go/pkg/logger"
)

type convertContext struct {
	rootSchema *jsonschema.Schema
	imports    map[string]struct{}
	resultMap  map[string]convertResult
	paths      []string
}

type convertResult struct {
	IsSchema    bool
	Name        string
	Description string
	schema
	property
}

func (k *kclGenerator) genSchemaFromJsonSchema(w io.Writer, filename string, src interface{}) error {
	code, err := readSource(filename, src)
	if err != nil {
		return err
	}
	js := &jsonschema.Schema{}
	if err = js.UnmarshalJSON(code); err != nil {
		return err
	}

	// convert json schema to kcl schema
	ctx := convertContext{
		rootSchema: js,
		resultMap:  make(map[string]convertResult),
		imports:    make(map[string]struct{}),
		paths:      []string{},
	}
	result := convertSchemaFromJsonSchema(&ctx, js,
		strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename)))
	if !result.IsSchema {
		panic("result is not schema")
	}
	kclSch := kclFile{
		Schemas: []schema{result.schema},
	}
	for _, imp := range getSortedKeys(ctx.imports) {
		kclSch.Imports = append(kclSch.Imports, kImport{PkgPath: imp})
	}
	for _, key := range getSortedKeys(ctx.resultMap) {
		if ctx.resultMap[key].IsSchema {
			kclSch.Schemas = append(kclSch.Schemas, ctx.resultMap[key].schema)
		}
	}

	// generate kcl schema code
	return k.genKcl(w, kclSch)
}

func convertSchemaFromJsonSchema(ctx *convertContext, s *jsonschema.Schema, name string) convertResult {
	// in jsonschema, type is one of True, False and Object
	// we only convert Object type
	if s.SchemaType != jsonschema.SchemaTypeObject {
		return convertResult{IsSchema: false}
	}

	// for the name of the result, we prefer $id, then title, then name in parameter
	// if none of them exists, "MyType" as default
	if id, ok := s.Keywords["$id"].(*jsonschema.ID); ok {
		lastSlashIndex := strings.LastIndex(string(*id), "/")
		name = strings.Replace(string(*id)[lastSlashIndex+1:], ".json", "", -1)
	} else if title, ok := s.Keywords["title"].(*jsonschema.Title); ok {
		name = string(*title)
	}
	if name == "" {
		name = "MyType"
	}
	result := convertResult{IsSchema: false, Name: name}
	ctx.paths = append(ctx.paths, name)

	isArray := false
	typeList := typeUnion{}
	required := make(map[string]struct{})
	for i := 0; i < len(s.OrderedKeywords); i++ {
		k := s.OrderedKeywords[i]
		switch v := s.Keywords[k].(type) {
		case *jsonschema.Title:
		case *jsonschema.Comment:
		case *jsonschema.SchemaURI:
		case *jsonschema.ID:
		case *jsonschema.Description:
			result.Description = string(*v)
		case *jsonschema.Type:
			if len(v.Vals) == 1 {
				switch v.Vals[0] {
				case "object":
					result.IsSchema = true
					continue
				case "array":
					isArray = true
					continue
				}
			}
			typeList.Items = append(typeList.Items, jsonTypesToKclTypes(v.Vals))
		case *jsonschema.Items:
			if !v.Single {
				logger.GetLogger().Warningf("unsupported multiple items: %#v", v)
				break
			}
			for i, val := range v.Schemas {
				item := convertSchemaFromJsonSchema(ctx, val, "items"+strconv.Itoa(i))
				if item.IsSchema {
					ctx.resultMap[item.schema.Name] = item
					typeList.Items = append(typeList.Items, typeCustom{Name: item.schema.Name})
				} else {
					typeList.Items = append(typeList.Items, item.Type)
				}
			}
		case *jsonschema.Required:
			for _, key := range []string(*v) {
				required[key] = struct{}{}
			}
		case *jsonschema.Properties:
			for _, prop := range *v {
				key := prop.Key
				val := prop.Value
				propSch := convertSchemaFromJsonSchema(ctx, val, key)
				_, propSch.Required = required[key]
				if propSch.IsSchema {
					ctx.resultMap[propSch.schema.Name] = propSch
				}
				result.Properties = append(result.Properties, propSch.property)
				if !propSch.IsSchema {
					for _, validate := range propSch.Validations {
						validate.Name = propSch.property.Name
						result.Validations = append(result.Validations, validate)
					}
				}
			}
		case *jsonschema.PatternProperties:
			result.IsSchema = true
			canConvert := true
			if result.HasIndexSignature {
				canConvert = false
				logger.GetLogger().Warningf("failed to convert patternProperties: already has index signature.")
			}
			if len(*v) != 1 {
				canConvert = false
				logger.GetLogger().Warningf("unsupported multiple patternProperties.")
			}
			result.HasIndexSignature = true
			result.IndexSignature = indexSignature{
				Type: typePrimitive(typAny),
			}
			for _, prop := range *v {
				val := prop.Schema
				propSch := convertSchemaFromJsonSchema(ctx, val, "patternProperties")
				if propSch.IsSchema {
					ctx.resultMap[propSch.schema.Name] = propSch
				}
				if canConvert {
					result.IndexSignature = indexSignature{
						Alias: "key",
						Type:  propSch.property.Type,
					}
					result.Validations = append(result.Validations, validation{
						Name:  "key",
						Regex: prop.Re,
					})
				}
			}
		case *jsonschema.Default:
			result.HasDefault = true
			result.DefaultValue = v.Data
		case *jsonschema.Enum:
			typeList.Items = make([]typeInterface, 0, len(*v))
			for _, val := range *v {
				unmarshalledVal := interface{}(nil)
				err := json.Unmarshal(val, &unmarshalledVal)
				if err != nil {
					logger.GetLogger().Warningf("failed to unmarshal enum value: %s", err)
					continue
				}
				typeList.Items = append(typeList.Items, typeValue{
					Value: unmarshalledVal,
				})
			}
		case *jsonschema.Const:
			unmarshalledVal := interface{}(nil)
			err := json.Unmarshal(*v, &unmarshalledVal)
			if err != nil {
				logger.GetLogger().Warningf("failed to unmarshal const value: %s", err)
				continue
			}
			typeList.Items = []typeInterface{typeValue{Value: unmarshalledVal}}
			result.HasDefault = true
			result.DefaultValue = unmarshalledVal
		case *jsonschema.Defs:
		case *jsonschema.Ref:
			refSch := v.ResolveRef(ctx.rootSchema)
			if refSch == nil || refSch.OrderedKeywords == nil {
				logger.GetLogger().Warningf("failed to resolve ref: %s", v.Reference)
			}
			for _, key := range refSch.OrderedKeywords {
				if _, ok := s.Keywords[key]; ok {
					logger.GetLogger().Warningf("keyword %s already exists when resolving ref %s", key, v.Reference)
					continue
				}
				s.OrderedKeywords = append(s.OrderedKeywords, key)
				s.Keywords[key] = refSch.Keywords[key]
			}
			sort.SliceStable(s.OrderedKeywords[i+1:], func(i, j int) bool {
				return jsonschema.GetKeywordOrder(s.OrderedKeywords[i]) < jsonschema.GetKeywordOrder(s.OrderedKeywords[j])
			})
		case *jsonschema.AdditionalProperties:
			switch v.SchemaType {
			case jsonschema.SchemaTypeObject:
				sch := convertSchemaFromJsonSchema(ctx, (*jsonschema.Schema)(v), "additionalProperties")
				result.HasIndexSignature = true
				result.IndexSignature = indexSignature{
					Type: sch.Type,
				}
			case jsonschema.SchemaTypeTrue:
				result.HasIndexSignature = true
				result.IndexSignature = indexSignature{
					Type: typePrimitive(typAny),
				}
			case jsonschema.SchemaTypeFalse:
				result.HasIndexSignature = false
			}
		case *jsonschema.Minimum:
			result.Validations = append(result.Validations, validation{
				Minimum:          (*float64)(v),
				ExclusiveMinimum: false,
			})
		case *jsonschema.Maximum:
			result.Validations = append(result.Validations, validation{
				Maximum:          (*float64)(v),
				ExclusiveMaximum: false,
			})
		case *jsonschema.ExclusiveMinimum:
			result.Validations = append(result.Validations, validation{
				Minimum:          (*float64)(v),
				ExclusiveMinimum: true,
			})
		case *jsonschema.ExclusiveMaximum:
			result.Validations = append(result.Validations, validation{
				Maximum:          (*float64)(v),
				ExclusiveMaximum: true,
			})
		case *jsonschema.MinLength:
			result.Validations = append(result.Validations, validation{
				MinLength: (*int)(v),
			})
		case *jsonschema.MaxLength:
			result.Validations = append(result.Validations, validation{
				MaxLength: (*int)(v),
			})
		case *jsonschema.Pattern:
			result.Validations = append(result.Validations, validation{
				Regex: (*regexp.Regexp)(v),
			})
			ctx.imports["regex"] = struct{}{}
		case *jsonschema.MultipleOf:
			vInt := int(*v)
			if float64(vInt) != float64(*v) {
				logger.GetLogger().Warningf("unsupported multipleOf value: %f", *v)
				continue
			}
			result.Validations = append(result.Validations, validation{
				MultiplyOf: &vInt,
			})
		case *jsonschema.UniqueItems:
			if *v {
				result.Validations = append(result.Validations, validation{
					Unique: true,
				})
			}
		case *jsonschema.MinItems:
			result.Validations = append(result.Validations, validation{
				MinLength: (*int)(v),
			})
		case *jsonschema.MaxItems:
			result.Validations = append(result.Validations, validation{
				MaxLength: (*int)(v),
			})
		case *jsonschema.OneOf:
			for i, val := range *v {
				item := convertSchemaFromJsonSchema(ctx, val, "oneOf"+strconv.Itoa(i))
				if item.IsSchema {
					ctx.resultMap[item.schema.Name] = item
					typeList.Items = append(typeList.Items, typeCustom{Name: item.schema.Name})
				} else {
					typeList.Items = append(typeList.Items, item.Type)
				}
			}
		case *jsonschema.AllOf:
			schs := *v
			for i := 0; i < len(schs); i++ {
				sch := schs[i]
				for _, key := range sch.OrderedKeywords {
					if _, ok := s.Keywords[key]; !ok {
						s.OrderedKeywords = append(s.OrderedKeywords, key)
						s.Keywords[key] = sch.Keywords[key]
					} else {
						switch v := sch.Keywords[key].(type) {
						case *jsonschema.Type:
						case *jsonschema.Ref:
							refSch := v.ResolveRef(ctx.rootSchema)
							if refSch == nil || refSch.OrderedKeywords == nil {
								logger.GetLogger().Warningf("failed to resolve ref: %s", v.Reference)
							}
							schs = append(schs, refSch)
						case *jsonschema.Properties:
							props := *s.Keywords[key].(*jsonschema.Properties)
							props = append(props, *v...)
							s.Keywords[key] = &props
						case *jsonschema.Required:
							reqs := *s.Keywords[key].(*jsonschema.Required)
							reqs = append(reqs, *v...)
							s.Keywords[key] = &reqs
						default:
							logger.GetLogger().Warningf("failed to merge allOf: unsupported keyword %s", key)
						}
					}
				}
			}
			sort.SliceStable(s.OrderedKeywords[i+1:], func(i, j int) bool {
				return jsonschema.GetKeywordOrder(s.OrderedKeywords[i]) < jsonschema.GetKeywordOrder(s.OrderedKeywords[j])
			})
		default:
			logger.GetLogger().Warningf("unknown Keyword: %s", k)
		}
	}

	if result.IsSchema {
		var s strings.Builder
		for _, p := range ctx.paths {
			s.WriteString(strcase.ToCamel(p))
		}
		result.schema.Name = s.String()
		result.schema.Description = result.Description
		result.Type = typeCustom{Name: strcase.ToCamel(result.schema.Name)}
		if len(result.Properties) == 0 && !result.HasIndexSignature {
			result.HasIndexSignature = true
			result.IndexSignature = indexSignature{Type: typePrimitive(typAny)}
		}
	} else {
		if len(typeList.Items) != 0 {
			if isArray {
				result.Type = typeArray{Items: typeList}
			} else {
				result.Type = typeList
			}
		} else {
			result.Type = typePrimitive(typAny)
		}
	}
	result.property.Name = strcase.ToSnake(result.Name)
	result.property.Description = result.Description
	ctx.paths = ctx.paths[:len(ctx.paths)-1]
	return result
}

func jsonTypesToKclTypes(t []string) typeInterface {
	var kclTypes typeUnion
	for _, v := range t {
		kclTypes.Items = append(kclTypes.Items, jsonTypeToKclType(v))
	}
	return kclTypes
}

func jsonTypeToKclType(t string) typeInterface {
	switch t {
	case "string":
		return typePrimitive(typStr)
	case "boolean", "bool":
		return typePrimitive(typBool)
	case "integer":
		return typePrimitive(typInt)
	case "number":
		return typePrimitive(typFloat)
	default:
		logger.GetLogger().Warningf("unknown type: %s", t)
		return typePrimitive(typStr)
	}
}
