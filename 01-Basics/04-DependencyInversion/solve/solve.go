/*
	This principle state that a High Level Module should not depend on Low Level Module. Let's take an example.
*/

package main

type Relationship int

const (
	Son = iota
	Father
	Mother
)

type Person struct {
	Name string
}

type Info struct {
	From         *Person
	Relationship Relationship
	To           *Person
}

// Low Level Module - As it interacts with database
type Relationships struct {
	Relations []Info
}

func (r *Relationships) Investigate() {
	for _, relation := range r.Relations {
		// Do Some research here.
		_ = relation
	}
}

func (r *Relationships) Add(from *Person, relationship Relationship, to *Person) {
	r.Relations = append(r.Relations, Info{from, relationship, to})
}

// High Level Module
type Research struct {
	relationships Relationships
}

func (r *Research) Investigate() {
	r.relationships.Investigate()
}
