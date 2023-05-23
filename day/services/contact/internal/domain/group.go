package domain

import "fmt"

type Group struct {
	ID   int
	Name string
}

func (g *Group) SetName(name string) error {
	if len(name) > 250 {
		return fmt.Errorf("Group name must be no more than 250 characters")
	}

	g.Name = name
	return nil
}
