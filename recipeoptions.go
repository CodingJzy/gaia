package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// RecipeOptionsAppCrendentialFormatValue represents the possible values for attribute "appCrendentialFormat".
type RecipeOptionsAppCrendentialFormatValue string

const (
	// RecipeOptionsAppCrendentialFormatJSON represents the value JSON.
	RecipeOptionsAppCrendentialFormatJSON RecipeOptionsAppCrendentialFormatValue = "JSON"

	// RecipeOptionsAppCrendentialFormatYAML represents the value YAML.
	RecipeOptionsAppCrendentialFormatYAML RecipeOptionsAppCrendentialFormatValue = "YAML"
)

// RecipeOptionsDeploymentModeValue represents the possible values for attribute "deploymentMode".
type RecipeOptionsDeploymentModeValue string

const (
	// RecipeOptionsDeploymentModeNamespaceUnique represents the value NamespaceUnique.
	RecipeOptionsDeploymentModeNamespaceUnique RecipeOptionsDeploymentModeValue = "NamespaceUnique"

	// RecipeOptionsDeploymentModeUnique represents the value Unique.
	RecipeOptionsDeploymentModeUnique RecipeOptionsDeploymentModeValue = "Unique"

	// RecipeOptionsDeploymentModeUnrestricted represents the value Unrestricted.
	RecipeOptionsDeploymentModeUnrestricted RecipeOptionsDeploymentModeValue = "Unrestricted"
)

// RecipeOptions represents the model of a recipeoptions
type RecipeOptions struct {
	// Indicates the format of the app credential.
	AppCrendentialFormat RecipeOptionsAppCrendentialFormatValue `json:"appCrendentialFormat" msgpack:"appCrendentialFormat" bson:"appcrendentialformat" mapstructure:"appCrendentialFormat,omitempty"`

	// Indicates if the recipe can be run multiple time.
	DeploymentMode RecipeOptionsDeploymentModeValue `json:"deploymentMode" msgpack:"deploymentMode" bson:"deploymentmode" mapstructure:"deploymentMode,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewRecipeOptions returns a new *RecipeOptions
func NewRecipeOptions() *RecipeOptions {

	return &RecipeOptions{
		ModelVersion:         1,
		AppCrendentialFormat: RecipeOptionsAppCrendentialFormatJSON,
		DeploymentMode:       RecipeOptionsDeploymentModeUnrestricted,
	}
}

// BleveType implements the bleve.Classifier Interface.
func (o *RecipeOptions) BleveType() string {

	return "recipeoptions"
}

// DeepCopy returns a deep copy if the RecipeOptions.
func (o *RecipeOptions) DeepCopy() *RecipeOptions {

	if o == nil {
		return nil
	}

	out := &RecipeOptions{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *RecipeOptions.
func (o *RecipeOptions) DeepCopyInto(out *RecipeOptions) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy RecipeOptions: %s", err))
	}

	*out = *target.(*RecipeOptions)
}

// Validate valides the current information stored into the structure.
func (o *RecipeOptions) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("appCrendentialFormat", string(o.AppCrendentialFormat), []string{"JSON", "YAML"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("deploymentMode", string(o.DeploymentMode), []string{"Unrestricted", "Unique", "NamespaceUnique"}, false); err != nil {
		errors = errors.Append(err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
