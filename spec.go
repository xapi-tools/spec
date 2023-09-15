package api

type Spec struct {
	Components Components `yaml:"components,omitempty"`
}

type Components struct {
	Schemas Schemas `yaml:"schemas,omitempty"`
}

type Schemas map[string]Schema

type Schema struct {
	Description string     `yaml:"description,omitempty"`
	Properties  Properties `yaml:"properties,omitempty"`
	Required    []string   `yaml:"required,omitempty"`
	// Possible value is `object`
	Type string `yaml:"type,omitempty"`

	XApiStatus XApiStatus `yaml:"x-api-status,omitempty"`

	// Names of schemas of type object from each of which all properties shall
	// be embedded. If a property to be embedded already exists, then it shall
	// be skipped
	XApiEmbed []string `yaml:"x-api-embed,omitempty"`

	// When set to true
	// - exactly one property in the schema MUST be set at all times
	// - list of required properties in the schema MUST be empty
	XApiUnion bool `yaml:"x-api-union,omitempty"`
}

type Properties map[string]Property

type Property struct {
	Description string `yaml:"description,omitempty"`

	Ref *string `yaml:"$ref,omitempty"`

	// Possible values are `string`, `integer`, `number`, `boolean` and `array`
	Type    string  `yaml:"type,omitempty"`
	Format  *string `yaml:"format,omitempty"`
	Pattern *string `yaml:"pattern,omitempty"`
	Min     *string `yaml:"min,omitempty"`
	Max     *string `yaml:"max,omitempty"`
	Default *string `yaml:"default,omitempty"`

	Enum []string `yaml:"enum,omitempty"`

	Items       *Items `yaml:"items,omitempty"`
	MinItems    *uint  `yaml:"minItems,omitempty"`
	MaxItems    *uint  `yaml:"maxItems,omitempty"`
	UniqueItems bool   `yaml:"uniqueItems,omitempty"`

	ReadOnly  bool `yaml:"readOnly,omitempty"`
	WriteOnly bool `yaml:"writeOnly,omitempty"`

	XApiEnum   XApiEnum   `yaml:"x-api-enum,omitempty"`
	XApiStatus XApiStatus `yaml:"x-api-status,omitempty"`

	// Applicable only for a property of type `string`
	// Possible values are `global` and `local`
	XApiUid *string `yaml:"x-api-uid,omitempty"`
}

type Items struct {
	Description string `yaml:"description,omitempty"`

	Ref *string `yaml:"$ref,omitempty"`

	// Possible values are `string`, `integer`, `number` and `boolean`
	Type    string  `yaml:"type,omitempty"`
	Format  *string `yaml:"format,omitempty"`
	Pattern *string `yaml:"pattern,omitempty"`
	Min     *string `yaml:"min,omitempty"`
	Max     *string `yaml:"max,omitempty"`

	Enum []string `yaml:"enum,omitempty"`

	XApiEnum XApiEnum `yaml:"x-api-enum,omitempty"`
}

// This should accompany enums defined in a property or items
// The key is an enum variant name and value is details of enum variant
type XApiEnum map[string]XApiEnumVariant
type XApiEnumVariant struct {
	Description string `yaml:"description,omitempty"`
	Value       *uint  `yaml:"value,omitempty"`

	XApiStatus XApiStatus `yaml:"x-api-status,omitempty"`
}

// Possible values of key are `experimental` and `deprecated`
// Value shall contain further details of the status
type XApiStatus map[string]string
