package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

// PolicyRuleIdentity represents the Identity of the object
var PolicyRuleIdentity = elemental.Identity{
	Name:     "policyrule",
	Category: "policyrules",
}

// PolicyRulesList represents a list of PolicyRules
type PolicyRulesList []*PolicyRule

// PolicyRule represents the model of a policyrule
type PolicyRule struct {
	// ID is the identifier of the object.
	ID string `json:"ID,omitempty" cql:"-"`

	// Action defines set of actions that must be enforced when a dependency is met.
	Action map[string]map[string]string `json:"action,omitempty" cql:"action,omitempty"`

	// Policy target networks
	Files FilePathsList `json:"files,omitempty" cql:"files,omitempty"`

	// Policy target networks
	Namespaces NamespacesList `json:"namespaces,omitempty" cql:"namespaces,omitempty"`

	// Policy target networks
	Networks ExternalServicesList `json:"networks,omitempty" cql:"networks,omitempty"`

	// Relation describes the required operation to be performed between subjects and objects
	Relation []string `json:"relation,omitempty" cql:"relation,omitempty"`

	// Policy target networks
	Syscalls SystemCallsList `json:"syscalls,omitempty" cql:"syscalls,omitempty"`

	// Policy target tags
	Tagclauses [][]string `json:"tagclauses,omitempty" cql:"tagclauses,omitempty"`
}

// NewPolicyRule returns a new *PolicyRule
func NewPolicyRule() *PolicyRule {

	return &PolicyRule{}
}

// Identity returns the Identity of the object.
func (o *PolicyRule) Identity() elemental.Identity {

	return PolicyRuleIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyRule) Identifier() string {

	return o.ID
}

func (o *PolicyRule) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyRule) SetIdentifier(ID string) {

	o.ID = ID
}

// Validate valides the current information stored into the structure.
func (o *PolicyRule) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o PolicyRule) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return PolicyRuleAttributesMap[name]
}

// PolicyRuleAttributesMap represents the map of attribute for PolicyRule.
var PolicyRuleAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "action",
		Stored:         true,
		SubType:        "actions_list",
		Type:           "external",
	},
	"Files": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "files",
		Stored:         true,
		SubType:        "file_entities",
		Type:           "external",
	},
	"Namespaces": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "namespaces",
		Stored:         true,
		SubType:        "namespace_entities",
		Type:           "external",
	},
	"Networks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "networks",
		Stored:         true,
		SubType:        "network_entities",
		Type:           "external",
	},
	"Relation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "relation",
		Stored:         true,
		SubType:        "relations_list",
		Type:           "external",
	},
	"Syscalls": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "syscalls",
		Stored:         true,
		SubType:        "syscall_entities",
		Type:           "external",
	},
	"Tagclauses": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "tagclauses",
		Stored:         true,
		SubType:        "target_tags",
		Type:           "external",
	},
}
