package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/shared/golang/gaiaconstants"

// IntegrationAuthTypeValue represents the possible values for attribute "authType".
type IntegrationAuthTypeValue string

const (
	// IntegrationAuthTypeBasic represents the value Basic.
	IntegrationAuthTypeBasic IntegrationAuthTypeValue = "Basic"

	// IntegrationAuthTypeNone represents the value None.
	IntegrationAuthTypeNone IntegrationAuthTypeValue = "None"

	// IntegrationAuthTypeOauth represents the value OAuth.
	IntegrationAuthTypeOauth IntegrationAuthTypeValue = "OAuth"
)

// IntegrationTypeValue represents the possible values for attribute "type".
type IntegrationTypeValue string

const (
	// IntegrationTypeRegistry represents the value Registry.
	IntegrationTypeRegistry IntegrationTypeValue = "Registry"

	// IntegrationTypeVulnerabilityscanner represents the value VulnerabilityScanner.
	IntegrationTypeVulnerabilityscanner IntegrationTypeValue = "VulnerabilityScanner"
)

// IntegrationIdentity represents the Identity of the object
var IntegrationIdentity = elemental.Identity{
	Name:     "integration",
	Category: "integrations",
}

// IntegrationsList represents a list of Integrations
type IntegrationsList []*Integration

// ContentIdentity returns the identity of the objects in the list.
func (o IntegrationsList) ContentIdentity() elemental.Identity {
	return IntegrationIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o IntegrationsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// Integration represents the model of a integration
type Integration struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags"`

	// AuthType refers to the type of HTTP authentication used to query endpoints
	AuthType IntegrationAuthTypeValue `json:"authType" bson:"authtype"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" bson:"createdat"`

	// Endpoint is the API end point of the service
	Endpoint string `json:"endpoint" bson:"endpoint"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID" bson:"parentid"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType" bson:"parenttype"`

	// Password is the password of the user to be used in the HTTP Authorization header
	Password string `json:"password" bson:"password"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected"`

	// Status of an entity
	Status gaiaconstants.EntityStatus `json:"status" bson:"status"`

	// Type refers to type of the server
	Type IntegrationTypeValue `json:"type" bson:"type"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedat"`

	// Username refers to the username to be used in the HTTP Authorization header
	UserName string `json:"userName" bson:"username"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`
}

// NewIntegration returns a new *Integration
func NewIntegration() *Integration {

	return &Integration{
		ModelVersion:   1.0,
		AssociatedTags: []string{},
		AuthType:       "None",
		NormalizedTags: []string{},
		Status:         gaiaconstants.Active,
		Type:           "Registry",
	}
}

// Identity returns the Identity of the object.
func (o *Integration) Identity() elemental.Identity {

	return IntegrationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Integration) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Integration) SetIdentifier(ID string) {

	o.ID = ID
}

func (o *Integration) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Integration) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Integration) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *Integration) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetNamespace returns the namespace of the receiver
func (o *Integration) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Integration) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *Integration) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *Integration) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetParentID returns the parentID of the receiver
func (o *Integration) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *Integration) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *Integration) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *Integration) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetProtected returns the protected of the receiver
func (o *Integration) GetProtected() bool {
	return o.Protected
}

// GetStatus returns the status of the receiver
func (o *Integration) GetStatus() gaiaconstants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Integration) SetStatus(status gaiaconstants.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Integration) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Integration) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("authType", string(o.AuthType), []string{"Basic", "None", "OAuth"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Registry", "VulnerabilityScanner"}, false); err != nil {
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
func (Integration) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return IntegrationAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (Integration) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return IntegrationAttributesMap
}

// IntegrationAttributesMap represents the map of attribute for Integration.
var IntegrationAttributesMap = map[string]elemental.AttributeSpecification{
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
	"AuthType": elemental.AttributeSpecification{
		AllowedChoices: []string{"Basic", "None", "OAuth"},
		DefaultValue:   IntegrationAuthTypeValue("None"),
		Description:    `AuthType refers to the type of HTTP authentication used to query endpoints`,
		Exposed:        true,
		Filterable:     true,
		Name:           "authType",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "enum",
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
	"Endpoint": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Endpoint is the API end point of the service`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "endpoint",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
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
	"Password": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Password is the password of the user to be used in the HTTP Authorization header`,
		Exposed:        true,
		Format:         "free",
		Name:           "password",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
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
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"Registry", "VulnerabilityScanner"},
		DefaultValue:   IntegrationTypeValue("Registry"),
		Description:    `Type refers to type of the server`,
		Exposed:        true,
		Name:           "type",
		Orderable:      true,
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
	"UserName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Username refers to the username to be used in the HTTP Authorization header`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "userName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}
