package srv

import (
	"fmt"
)

type info struct {
	id string
	name string
}

type Info interface {
	Add() error
	Get() (string, error)
	Delete() error
}

func NewInfo(id, name string) Info {
	return &info{id: id, name: name}
}

func (i *info) Add() error {
	fmt.Println("Added %s:%s", i.id, i.name)
	return nil
}

func (i *info) Get() (string, error) {
	if i.id == nil {
		return fmt.Errorf("Nothing to return")
	}

	fmt.Println("Found %s", i.id)
	return i.name
}

func (i *info) Delete() error {
	if i.id == nil {
		return fmt.Errorf("Nothing to delete")
	}

	fmt.Println("Deleted %s", i.id)
	return nil
}
