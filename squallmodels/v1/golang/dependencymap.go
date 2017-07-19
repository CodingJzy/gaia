package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "github.com/aporeto-inc/gaia/shared/golang/gaiatypes"

// DependencyMapIdentity represents the Identity of the object
var DependencyMapIdentity = elemental.Identity{
	Name:     "dependencymap",
	Category: "dependencymaps",
}

// DependencyMapsList represents a list of DependencyMaps
type DependencyMapsList []*DependencyMap

// ContentIdentity returns the identity of the objects in the list.
func (o DependencyMapsList) ContentIdentity() elemental.Identity {

	return DependencyMapIdentity
}

// List converts the object to an elemental.IdentifiablesList.
func (o DependencyMapsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o DependencyMapsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o DependencyMapsList) Version() float64 {

	return 1.0
}

// DependencyMap represents the model of a dependencymap
type DependencyMap struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-"`

	// edges are the edges of the map
	Edges gaiatypes.GraphEdgeMap `json:"edges" bson:"-"`

	// Groups provide information about the group values
	Groups gaiatypes.GraphGroupMap `json:"groups" bson:"-"`

	// nodes refers to the nodes of the map
	Nodes gaiatypes.GraphNodeMap `json:"nodes" bson:"-"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewDependencyMap returns a new *DependencyMap
func NewDependencyMap() *DependencyMap {

	return &DependencyMap{
		ModelVersion: 1.0,
		Edges:        gaiatypes.GraphEdgeMap{},
		Groups:       gaiatypes.GraphGroupMap{},
		Nodes:        gaiatypes.GraphNodeMap{},
	}
}

// Identity returns the Identity of the object.
func (o *DependencyMap) Identity() elemental.Identity {

	return DependencyMapIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DependencyMap) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DependencyMap) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *DependencyMap) Version() float64 {

	return 1.0
}

// DefaultOrder returns the list of default ordering fields.
func (o *DependencyMap) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *DependencyMap) Doc() string {
	return `This api returns a data structure representing the graph of all processing units and their connections in a particular namespace, in a given time window. To pass the time window you can use the query parameters "startAbsolute", "endAbsolute", "startRelative", "endRelative".  For example "https://squall.aporeto.com/dependencymaps?startAbsolute=1489132800000&endAbsolute=1489219200000"`
}

func (o *DependencyMap) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *DependencyMap) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("edges", o.Edges); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("edges", o.Edges); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("groups", o.Groups); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("groups", o.Groups); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("nodes", o.Nodes); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("nodes", o.Nodes); err != nil {
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
func (*DependencyMap) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := DependencyMapAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return DependencyMapLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*DependencyMap) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DependencyMapAttributesMap
}

// DependencyMapAttributesMap represents the map of attribute for DependencyMap.
var DependencyMapAttributesMap = map[string]elemental.AttributeSpecification{
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
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"Edges": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `edges are the edges of the map`,
		Exposed:        true,
		Name:           "edges",
		ReadOnly:       true,
		Required:       true,
		SubType:        "graphedges_map",
		Type:           "external",
	},
	"Groups": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Groups provide information about the group values`,
		Exposed:        true,
		Name:           "groups",
		ReadOnly:       true,
		Required:       true,
		SubType:        "graphgroups_map",
		Type:           "external",
	},
	"Nodes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `nodes refers to the nodes of the map`,
		Exposed:        true,
		Name:           "nodes",
		ReadOnly:       true,
		Required:       true,
		SubType:        "graphnodes_map",
		Type:           "external",
	},
}

// DependencyMapLowerCaseAttributesMap represents the map of attribute for DependencyMap.
var DependencyMapLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID is the identifier of the object.`,
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
	"edges": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `edges are the edges of the map`,
		Exposed:        true,
		Name:           "edges",
		ReadOnly:       true,
		Required:       true,
		SubType:        "graphedges_map",
		Type:           "external",
	},
	"groups": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Groups provide information about the group values`,
		Exposed:        true,
		Name:           "groups",
		ReadOnly:       true,
		Required:       true,
		SubType:        "graphgroups_map",
		Type:           "external",
	},
	"nodes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `nodes refers to the nodes of the map`,
		Exposed:        true,
		Name:           "nodes",
		ReadOnly:       true,
		Required:       true,
		SubType:        "graphnodes_map",
		Type:           "external",
	},
}
