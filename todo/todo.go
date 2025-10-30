package todo

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type ByPri []Item //implements sort.Interface for []Item based on priority and position field

type Item struct {
	Text     string
	Priority int
	position int
	Done     bool
}

func (s ByPri) Len() int { return len(s) }

func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Priority == s[j].Priority {
		return s[i].Priority < s[j].Priority
	}
	if s[i].Done != s[j].Done {
		return s[i].Done
	}
	return s[i].position < s[j].position
}
func SaveItems(filename string, items []Item) error {
	//kuch to krna h yha
	b, err := json.Marshal(items)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)

	if err != nil {
		return err
	}
	// fmt.Println(string(b))
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	items := []Item{}
	//unmarshal the file
	err = json.Unmarshal(b, &items)
	if err != nil {
		return []Item{}, err
	}

	for i := range items { //using the index
		items[i].position = i + 1
	}
	return items, nil
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(H)"
	}
	if i.Priority == 3 {
		return "(L)"
	}
	return "(M)"
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "✅"
	}
	return ""
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}
