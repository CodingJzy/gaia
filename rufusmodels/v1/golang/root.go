package rufusmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

// RootIdentity represents the Identity of the object
var RootIdentity = elemental.Identity{
	Name:     "root",
	Category: "root",
}

// Root represents the model of a root
type Root struct {
	Token        string `json:"APIKey,omitempty"`
	Organization string `json:"enterprise,omitempty"`
}

// NewRoot returns a new *Root
func NewRoot() *Root {

	return &Root{}
}

// Identity returns the Identity of the object.
func (o *Root) Identity() elemental.Identity {

	return RootIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Root) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Root) SetIdentifier(ID string) {

}

func (o *Root) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Root) Validate() error {

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

// APIKey returns a the API Key
func (o *Root) APIKey() string {

	return o.Token
}

// SetAPIKey sets a the API Key
func (o *Root) SetAPIKey(key string) {

	o.Token = key
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (Root) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return RootAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (Root) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RootAttributesMap
}

// RootAttributesMap represents the map of attribute for Root.
var RootAttributesMap = map[string]elemental.AttributeSpecification{}
