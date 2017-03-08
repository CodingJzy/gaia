package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/shared/golang/gaiaconstants"

// PolicyTypeValue represents the possible values for attribute "type".
type PolicyTypeValue string

const (
	// PolicyTypeApiauthorization represents the value APIAuthorization.
	PolicyTypeApiauthorization PolicyTypeValue = "APIAuthorization"

	// PolicyTypeFile represents the value File.
	PolicyTypeFile PolicyTypeValue = "File"

	// PolicyTypeNamespacemapping represents the value NamespaceMapping.
	PolicyTypeNamespacemapping PolicyTypeValue = "NamespaceMapping"

	// PolicyTypeNetwork represents the value Network.
	PolicyTypeNetwork PolicyTypeValue = "Network"

	// PolicyTypeServer represents the value Server.
	PolicyTypeServer PolicyTypeValue = "Server"

	// PolicyTypeSyscall represents the value Syscall.
	PolicyTypeSyscall PolicyTypeValue = "Syscall"
)

// PolicyIdentity represents the Identity of the object
var PolicyIdentity = elemental.Identity{
	Name:     "policy",
	Category: "policies",
}

// PoliciesList represents a list of Policies
type PoliciesList []*Policy

// ContentIdentity returns the identity of the objects in the list.
func (o PoliciesList) ContentIdentity() elemental.Identity {
	return PolicyIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o PoliciesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// Policy represents the model of a policy
type Policy struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id"`

	// Action defines set of actions that must be enforced when a dependency is met.
	Action map[string]map[string]string `json:"action" bson:"action"`

	// This is a set of all object tags for matching in the DB
	AllObjectTags []string `json:"-" bson:"allobjecttags"`

	// This is a set of all subject tags for matching in the DB
	AllSubjectTags []string `json:"-" bson:"allsubjecttags"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" bson:"createdat"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags"`

	// Object represents set of entities that another entity depends on. As subjects, objects are identified as logical operations on tags when a policy is defined.
	Object [][]string `json:"object" bson:"object"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID" bson:"parentid"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType" bson:"parenttype"`

	// Propagate will propagate the policy to all of its children.
	Propagate bool `json:"propagate" bson:"propagate"`

	// If set to true while the policy is propagating, it won't be visible to children namespace, but still used for policy resolution.
	PropagationHidden bool `json:"propagationHidden" bson:"propagationhidden"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected"`

	// Relation describes the required operation to be performed between subjects and objects
	Relation []string `json:"relation" bson:"relation"`

	// Status of an entity
	Status gaiaconstants.EntityStatus `json:"status" bson:"status"`

	// Subject represent sets of entities that will have a dependency other entities. Subjects are defined as logical operations on tags. Logical operations can includes AND/OR
	Subject [][]string `json:"subject" bson:"subject"`

	// Type of the policy
	Type PolicyTypeValue `json:"type" bson:"type"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedat"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`
}

// NewPolicy returns a new *Policy
func NewPolicy() *Policy {

	return &Policy{
		ModelVersion:   1.0,
		AllObjectTags:  []string{},
		AllSubjectTags: []string{},
		AssociatedTags: []string{},
		NormalizedTags: []string{},
		Status:         gaiaconstants.Active,
	}
}

// Identity returns the Identity of the object.
func (o *Policy) Identity() elemental.Identity {

	return PolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Policy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Policy) SetIdentifier(ID string) {

	o.ID = ID
}

func (o *Policy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Policy) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Policy) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *Policy) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetName returns the name of the receiver
func (o *Policy) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *Policy) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *Policy) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Policy) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *Policy) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *Policy) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetParentID returns the parentID of the receiver
func (o *Policy) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *Policy) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *Policy) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *Policy) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetPropagate returns the propagate of the receiver
func (o *Policy) GetPropagate() bool {
	return o.Propagate
}

// SetPropagate set the given propagate of the receiver
func (o *Policy) SetPropagate(propagate bool) {
	o.Propagate = propagate
}

// GetPropagationHidden returns the propagationHidden of the receiver
func (o *Policy) GetPropagationHidden() bool {
	return o.PropagationHidden
}

// SetPropagationHidden set the given propagationHidden of the receiver
func (o *Policy) SetPropagationHidden(propagationHidden bool) {
	o.PropagationHidden = propagationHidden
}

// GetProtected returns the protected of the receiver
func (o *Policy) GetProtected() bool {
	return o.Protected
}

// GetStatus returns the status of the receiver
func (o *Policy) GetStatus() gaiaconstants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Policy) SetStatus(status gaiaconstants.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Policy) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Policy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("action", o.Action); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("action", o.Action); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("subject", o.Subject); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("subject", o.Subject); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"APIAuthorization", "File", "NamespaceMapping", "Network", "Server", "Syscall"}, false); err != nil {
		errors = append(errors, err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (Policy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return PolicyAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (Policy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PolicyAttributesMap
}

// PolicyAttributesMap represents the map of attribute for Policy.
var PolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Action defines set of actions that must be enforced when a dependency is met.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Stored:         true,
		SubType:        "actions_list",
		Type:           "external",
	},
	"AllObjectTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `This is a set of all object tags for matching in the DB`,
		Name:           "allObjectTags",
		Required:       true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"AllSubjectTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `This is a set of all subject tags for matching in the DB`,
		Name:           "allSubjectTags",
		Required:       true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"Annotation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AssociatedTags are the list of tags attached to an entity`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"CreatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `CreatedAt is the time at which an entity was created`,
		Exposed:        true,
		Name:           "createdAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Name is the name of the entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Index:          true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `NormalizedTags contains the list of normalized tags of the entities`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
	},
	"Object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Object represents set of entities that another entity depends on. As subjects, objects are identified as logical operations on tags when a policy is defined.`,
		Exposed:        true,
		Name:           "object",
		Stored:         true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentID is the ID of the parent, if any,`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Format:         "free",
		Getter:         true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ParentType": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"PropagationHidden": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `If set to true while the policy is propagating, it won't be visible to children namespace, but still used for policy resolution.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagationHidden",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Relation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Relation describes the required operation to be performed between subjects and objects`,
		Exposed:        true,
		Name:           "relation",
		Stored:         true,
		SubType:        "relations_list",
		Type:           "external",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		DefaultValue:   gaiaconstants.Active,
		Description:    `Status of an entity`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "status",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		SubType:        "status_enum",
		Type:           "external",
	},
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Subject represent sets of entities that will have a dependency other entities. Subjects are defined as logical operations on tags. Logical operations can includes AND/OR`,
		Exposed:        true,
		Name:           "subject",
		Required:       true,
		Stored:         true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"APIAuthorization", "File", "NamespaceMapping", "Network", "Server", "Syscall"},
		CreationOnly:   true,
		Description:    `Type of the policy`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"UpdatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UpdatedAt is the time at which an entity was updated.`,
		Exposed:        true,
		Name:           "updatedAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}
