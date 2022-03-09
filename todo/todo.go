package todo

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Task struct {
	Name     string
	Done     bool
	position int
}

func (i *Task) Label() string {
	return strconv.Itoa(i.position) + ": "
}

// save tasks
func SaveItems(filename string, items []Task) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

// read tasks
func ReadItems(filename string) ([]Task, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var items []Task
	if err := json.Unmarshal(b, &items); err != nil {
		return []Task{}, nil
	}
	for i := range items {
		items[i].position = i + 1
	}
	return items, nil
}
