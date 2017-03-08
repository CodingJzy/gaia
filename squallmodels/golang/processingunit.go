package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/shared/golang/gaiaconstants"

// ProcessingUnitOperationalStatusValue represents the possible values for attribute "operationalStatus".
type ProcessingUnitOperationalStatusValue string

const (
	// ProcessingUnitOperationalStatusInitialized represents the value Initialized.
	ProcessingUnitOperationalStatusInitialized ProcessingUnitOperationalStatusValue = "Initialized"

	// ProcessingUnitOperationalStatusPaused represents the value Paused.
	ProcessingUnitOperationalStatusPaused ProcessingUnitOperationalStatusValue = "Paused"

	// ProcessingUnitOperationalStatusRunning represents the value Running.
	ProcessingUnitOperationalStatusRunning ProcessingUnitOperationalStatusValue = "Running"

	// ProcessingUnitOperationalStatusStopped represents the value Stopped.
	ProcessingUnitOperationalStatusStopped ProcessingUnitOperationalStatusValue = "Stopped"

	// ProcessingUnitOperationalStatusTerminated represents the value Terminated.
	ProcessingUnitOperationalStatusTerminated ProcessingUnitOperationalStatusValue = "Terminated"
)

// ProcessingUnitTypeValue represents the possible values for attribute "type".
type ProcessingUnitTypeValue string

const (
	// ProcessingUnitTypeDocker represents the value Docker.
	ProcessingUnitTypeDocker ProcessingUnitTypeValue = "Docker"

	// ProcessingUnitTypeLinuxservice represents the value LinuxService.
	ProcessingUnitTypeLinuxservice ProcessingUnitTypeValue = "LinuxService"

	// ProcessingUnitTypeRkt represents the value RKT.
	ProcessingUnitTypeRkt ProcessingUnitTypeValue = "RKT"
)

// ProcessingUnitIdentity represents the Identity of the object
var ProcessingUnitIdentity = elemental.Identity{
	Name:     "processingunit",
	Category: "processingunits",
}

// ProcessingUnitsList represents a list of ProcessingUnits
type ProcessingUnitsList []*ProcessingUnit

// ContentIdentity returns the identity of the objects in the list.
func (o ProcessingUnitsList) ContentIdentity() elemental.Identity {
	return ProcessingUnitIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o ProcessingUnitsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// ProcessingUnit represents the model of a processingunit
type ProcessingUnit struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" bson:"createdat"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// LastSyncTime is the time when the policy was last resolved
	LastSyncTime time.Time `json:"lastSyncTime" bson:"lastsynctime"`

	// Metadata are list of tags associated to the processing unit
	Metadata []string `json:"metadata" bson:"metadata"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace"`

	// NativeContextID is the Docker UUID or service PID
	NativeContextID string `json:"nativeContextID" bson:"nativecontextid"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags"`

	// OperationalStatus of the processing unit
	OperationalStatus ProcessingUnitOperationalStatusValue `json:"operationalStatus" bson:"operationalstatus"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected"`

	// serverID is the ID of the server associated with the processing unit
	ServerID string `json:"serverID" bson:"serverid"`

	// Status of an entity
	Status gaiaconstants.EntityStatus `json:"status" bson:"status"`

	// Type of the container ecosystem
	Type ProcessingUnitTypeValue `json:"type" bson:"type"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedat"`

	// Vulnerabilities contains the list of vulnerabilities of the processing unit.
	Vulnerabilities []string `json:"-" bson:"vulnerabilities"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`
}

// NewProcessingUnit returns a new *ProcessingUnit
func NewProcessingUnit() *ProcessingUnit {

	return &ProcessingUnit{
		ModelVersion:      1.0,
		AssociatedTags:    []string{},
		NormalizedTags:    []string{},
		OperationalStatus: "Initialized",
		Status:            gaiaconstants.Active,
	}
}

// Identity returns the Identity of the object.
func (o *ProcessingUnit) Identity() elemental.Identity {

	return ProcessingUnitIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ProcessingUnit) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ProcessingUnit) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *ProcessingUnit) Version() float64 {

	return 1.0
}

func (o *ProcessingUnit) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *ProcessingUnit) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *ProcessingUnit) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *ProcessingUnit) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetName returns the name of the receiver
func (o *ProcessingUnit) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *ProcessingUnit) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *ProcessingUnit) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *ProcessingUnit) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *ProcessingUnit) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *ProcessingUnit) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetProtected returns the protected of the receiver
func (o *ProcessingUnit) GetProtected() bool {
	return o.Protected
}

// GetStatus returns the status of the receiver
func (o *ProcessingUnit) GetStatus() gaiaconstants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *ProcessingUnit) SetStatus(status gaiaconstants.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *ProcessingUnit) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *ProcessingUnit) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("operationalStatus", string(o.OperationalStatus), []string{"Initialized", "Paused", "Running", "Stopped", "Terminated"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Docker", "LinuxService", "RKT"}, false); err != nil {
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
func (ProcessingUnit) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return ProcessingUnitAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (ProcessingUnit) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ProcessingUnitAttributesMap
}

// ProcessingUnitAttributesMap represents the map of attribute for ProcessingUnit.
var ProcessingUnitAttributesMap = map[string]elemental.AttributeSpecification{
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
	"LastSyncTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `LastSyncTime is the time when the policy was last resolved`,
		Exposed:        true,
		Name:           "lastSyncTime",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Metadata are list of tags associated to the processing unit`,
		Exposed:        true,
		Filterable:     true,
		Name:           "metadata",
		Stored:         true,
		SubType:        "metadata_list",
		Type:           "external",
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
	"NativeContextID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `NativeContextID is the Docker UUID or service PID`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "nativeContextID",
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
	"OperationalStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Initialized", "Paused", "Running", "Stopped", "Terminated"},
		DefaultValue:   ProcessingUnitOperationalStatusValue("Initialized"),
		Description:    `OperationalStatus of the processing unit`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		Stored:         true,
		Type:           "enum",
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
	"ServerID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `serverID is the ID of the server associated with the processing unit`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "serverID",
		Stored:         true,
		Type:           "string",
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
		AllowedChoices: []string{"Docker", "LinuxService", "RKT"},
		CreationOnly:   true,
		Description:    `Type of the container ecosystem`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
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
	"Vulnerabilities": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Vulnerabilities contains the list of vulnerabilities of the processing unit.`,
		Name:           "vulnerabilities",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}
