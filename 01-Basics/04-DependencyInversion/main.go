/*
	This principle state that a High Level Module should not depend on Low Level Module. Let's take an example.
*/

package main

// We are going to track relationships between persons in this example.

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

func (r *Relationships) Add(from *Person, relationship Relationship, to *Person) {
	r.Relations = append(r.Relations, Info{from, relationship, to})
}

// High Level Module
type Research struct {
	relationships Relationships
}

func (r *Research) Investigate() {
	for _, relation := range r.relationships.Relations {
		// Do Some research here.
		_ = relation
	}
}

// In above code issue is that Research which is High Level Module depends on Relationships which is low level module.
// In case if anything changes in low level module it will affect high level module.
// So the solution for such kind of problem is that move the Investigate logic to low level module and let high level module just call the logic.
