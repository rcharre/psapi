// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Pokemon Studio API
 *
 * API for the Pokemon Studio
 *
 * API version: 0.0.1
 */

package psapigen

// PokemonDetails - An object containing detailed data of a pokemon with its main form
type PokemonDetails struct {

	// The symbol of the pokemon
	Symbol string `json:"symbol,omitempty"`

	// The translated name of the first pokemon's form
	Name string `json:"name,omitempty"`

	// The translated description of the first pokemon's form
	Description string `json:"description,omitempty"`

	// The number of the pokemon
	Number int32 `json:"number,omitempty"`

	MainForm FormDetails `json:"main_form,omitempty"`
}

// AssertPokemonDetailsRequired checks if the required fields are not zero-ed
func AssertPokemonDetailsRequired(obj PokemonDetails) error {
	if err := AssertFormDetailsRequired(obj.MainForm); err != nil {
		return err
	}
	return nil
}

// AssertPokemonDetailsConstraints checks if the values respects the defined constraints
func AssertPokemonDetailsConstraints(obj PokemonDetails) error {
	if err := AssertFormDetailsConstraints(obj.MainForm); err != nil {
		return err
	}
	return nil
}