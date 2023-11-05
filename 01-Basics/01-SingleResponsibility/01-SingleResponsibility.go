/*
	Single Responsibility Principle state that a every class, module and function should have only one responsibility.
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

var entryCount = 0

type Journal struct {
	notes []string
}

// These function follow single concern with Journal. We can perform CRUD operations on Journal and that's ok with Single Responsibility Principle.
func (j *Journal) AddNote(text string) {
	entryCount++
	newNote := fmt.Sprintf("%d : %s", entryCount, text)

	j.notes = append(j.notes, newNote)
}

func (j *Journal) RemoveNote(entry string) {
	newNotes := []string{}

	for _, val := range j.notes {
		if string(val[0]) != entry {
			newNotes = append(newNotes, val)
		}
	}

	j.notes = newNotes
}

func (j *Journal) Print() {
	for _, val := range j.notes {
		fmt.Println(val)
	}
}

// Seperation of Concern
// These might seem okay but they broke the rule of Single Responsibility.
// Suppose we have another Type of Quotes and we also need to Save that to file, html and word than we need to rewrite all these functions.
func (j *Journal) Save()       {}
func (j *Journal) SaveToHTML() {}
func (j *Journal) SaveToWord() {}

// Solution -
// Instead of mixing concerns what we can do is that we can have Persistance struct that will have all the methods to save.
type Persistence struct {
	seprator string
}

func (p *Persistence) Save(filename string, data []byte) {
	_ = os.WriteFile(filename, data, 0644)
}
func (p *Persistence) SaveToHTML() {}
func (p *Persistence) SaveToWord() {}

func main() {
	j := Journal{}

	j.AddNote("I'm learning Design Patterns")
	j.AddNote("This is very interesting")
	j.AddNote("This will help me write better code")

	j.Print()

	j.RemoveNote("2")
	fmt.Println()

	j.Print()

	p := Persistence{seprator: "\n"}

	p.Save("journal.txt", []byte(strings.Join(j.notes, p.seprator)))
}
