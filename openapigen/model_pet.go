/*
 * Swagger Petstore
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

//a
type Pet struct {
	Id int64 `json:"id"`

	Name string `json:"name"`

	Tag string `json:"tag,omitempty"`
}

// AssertPetRequired checks if the required fields are not zero-ed
func AssertPetRequired(obj Pet) error {
	elements := map[string]interface{}{
		"id":   obj.Id,
		"name": obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecursePetRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Pet (e.g. [][]Pet), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePetRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPet, ok := obj.(Pet)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPetRequired(aPet)
	})
}
