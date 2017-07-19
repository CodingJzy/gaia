package vincemodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// PlanIdentity represents the Identity of the object
var PlanIdentity = elemental.Identity{
	Name:     "plan",
	Category: "plans",
}

// PlansList represents a list of Plans
type PlansList []*Plan

// ContentIdentity returns the identity of the objects in the list.
func (o PlansList) ContentIdentity() elemental.Identity {

	return PlanIdentity
}

// List converts the object to an elemental.IdentifiablesList.
func (o PlansList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PlansList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o PlansList) Version() float64 {

	return 1.0
}

// Plan represents the model of a plan
type Plan struct {
	// Description contains the description of the Plan.
	Description string `json:"description" bson:"description"`

	// EnforcerQuota contains the maximum number of enforcers available in the Plan.
	EnforcersQuota int `json:"enforcersQuota" bson:"enforcersquota"`

	// Key contains the key identifier of the Plan.
	Key string `json:"key" bson:"key"`

	// Name contains the name of the Plan.
	Name string `json:"name" bson:"name"`

	// PoliciesQuota contains the maximum number of policies available in the Plan.
	PoliciesQuota int `json:"policiesQuota" bson:"policiesquota"`

	// ProcessingUnitsQuota contains the maximum PUs available in the Plan.
	ProcessingUnitsQuota int `json:"processingUnitsQuota" bson:"processingunitsquota"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewPlan returns a new *Plan
func NewPlan() *Plan {

	return &Plan{
		ModelVersion: 1.0,
	}
}

// Identity returns the Identity of the object.
func (o *Plan) Identity() elemental.Identity {

	return PlanIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Plan) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Plan) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *Plan) Version() float64 {

	return 1.0
}

// DefaultOrder returns the list of default ordering fields.
func (o *Plan) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Plan) Doc() string {
	return `Plan contains the various billing plans available`
}

func (o *Plan) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Plan) Validate() error {

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
func (*Plan) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PlanAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PlanLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Plan) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PlanAttributesMap
}

// PlanAttributesMap represents the map of attribute for Plan.
var PlanAttributesMap = map[string]elemental.AttributeSpecification{
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Description contains the description of the Plan.`,
		Exposed:        true,
		Format:         "free",
		Name:           "description",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcersQuota": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `EnforcerQuota contains the maximum number of enforcers available in the Plan.`,
		Exposed:        true,
		Name:           "enforcersQuota",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
	"Key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Key contains the key identifier of the Plan.`,
		Exposed:        true,
		Format:         "free",
		Name:           "key",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Name contains the name of the Plan.`,
		Exposed:        true,
		Format:         "free",
		Name:           "name",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"PoliciesQuota": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `PoliciesQuota contains the maximum number of policies available in the Plan.`,
		Exposed:        true,
		Name:           "policiesQuota",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
	"ProcessingUnitsQuota": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ProcessingUnitsQuota contains the maximum PUs available in the Plan.`,
		Exposed:        true,
		Name:           "processingUnitsQuota",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
}

// PlanLowerCaseAttributesMap represents the map of attribute for Plan.
var PlanLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Description contains the description of the Plan.`,
		Exposed:        true,
		Format:         "free",
		Name:           "description",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"enforcersquota": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `EnforcerQuota contains the maximum number of enforcers available in the Plan.`,
		Exposed:        true,
		Name:           "enforcersQuota",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
	"key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Key contains the key identifier of the Plan.`,
		Exposed:        true,
		Format:         "free",
		Name:           "key",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Name contains the name of the Plan.`,
		Exposed:        true,
		Format:         "free",
		Name:           "name",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"policiesquota": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `PoliciesQuota contains the maximum number of policies available in the Plan.`,
		Exposed:        true,
		Name:           "policiesQuota",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
	"processingunitsquota": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ProcessingUnitsQuota contains the maximum PUs available in the Plan.`,
		Exposed:        true,
		Name:           "processingUnitsQuota",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
}
