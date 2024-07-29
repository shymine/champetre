package champetre

// The type that must implement the types that will be registered
// as collection in the database
type Model interface {
	Id() string
	SetId(string)
	// Must return the default construction of the struct
	Default() Model
	// Must return the name of the struct
	Kind() string
	// Return a map[string]any where string is the name of the parameter and any its value
	Parameters() map[string]any
	// Replace the elements inside the model given an element of the same type
	Update(Model)
}