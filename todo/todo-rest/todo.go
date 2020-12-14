package main

import "github.com/google/uuid"

func newId() string {
	id, _ := uuid.NewUUID()
	return id.String()
}

type TodoItem struct {
	ID    string `json:"id"`
	User  string `json:"user"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Order int    `json:"order"`
	Text  string `json:"text"`
}

func (i *TodoItem) Update(item TodoItem) *TodoItem {
	i.Title = item.Title
	i.User = item.User
	i.Done = item.Done
	i.Order = item.Order
	i.Text = item.Text
	return i
}

type Todo map[string]*TodoItem

func (t Todo) All() []*TodoItem {
	items := []*TodoItem{}
	for _, item := range t {
		items = append(items, item)
	}
	return items
}

func (t Todo) Find(id string) *TodoItem {
	for _, item := range t {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (t Todo) Create(item TodoItem) *TodoItem {
	item.ID = newId()
	t[item.ID] = &item
	return &item
}

func (t Todo) Update(id string, updatedItem TodoItem) *TodoItem {
	if item := t.Find(id); item != nil {
		return item.Update(updatedItem)
	}
	return nil
}

func (t Todo) DeleteAll() {
	for k := range t {
		delete(t, k)
	}
}

func (t Todo) Delete(id string) {
	for k := range t {
		if k == id {
			delete(t, k)
		}
	}
}
