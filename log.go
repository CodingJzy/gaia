package gaia

import (
	"fmt"
	"sync"

	"go.aporeto.io/elemental"
)

// LogIdentity represents the Identity of the object.
var LogIdentity = elemental.Identity{
	Name:     "log",
	Category: "logs",
	Package:  "highwind",
	Private:  false,
}

// LogsList represents a list of Logs
type LogsList []*Log

// Identity returns the identity of the objects in the list.
func (o LogsList) Identity() elemental.Identity {

	return LogIdentity
}

// Copy returns a pointer to a copy the LogsList.
func (o LogsList) Copy() elemental.Identifiables {

	copy := append(LogsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the LogsList.
func (o LogsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(LogsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Log))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o LogsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o LogsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o LogsList) Version() int {

	return 1
}

// Log represents the model of a log
type Log struct {
	// Data contains all logs data.
	Data map[string]string `json:"data" bson:"-" mapstructure:"data,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewLog returns a new *Log
func NewLog() *Log {

	return &Log{
		ModelVersion: 1,
		Data:         map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *Log) Identity() elemental.Identity {

	return LogIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Log) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Log) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Log) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Log) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Log) Doc() string {
	return `Retrieves the log of a deployed app.`
}

func (o *Log) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Log) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Log) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := LogAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return LogLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Log) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return LogAttributesMap
}

// LogAttributesMap represents the map of attribute for Log.
var LogAttributesMap = map[string]elemental.AttributeSpecification{
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Data",
		Description:    `Data contains all logs data.`,
		Exposed:        true,
		Name:           "data",
		ReadOnly:       true,
		SubType:        "logs",
		Type:           "external",
	},
}

// LogLowerCaseAttributesMap represents the map of attribute for Log.
var LogLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Data",
		Description:    `Data contains all logs data.`,
		Exposed:        true,
		Name:           "data",
		ReadOnly:       true,
		SubType:        "logs",
		Type:           "external",
	},
}
