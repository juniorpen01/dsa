package algorithms2

import (
	"fmt"

	datastructures "github.com/juniorpen01/dsa/internal/data_structures"
)

type User struct {
	ID     int
	Action string
}

func searchAndRemove(queue *datastructures.Queue, id int) *int {
	for i, item := range queue.Items {
		if item == id {
			queue.Items = append(queue.Items[:i], queue.Items[i+1:]...)
			return &id
		}
	}
	return nil
}

func Matchmake(queue *datastructures.Queue, user User) string {
	switch user.Action {
	case "join":
		queue.Push(user.ID)
		if queue.Size() == 4 {
			id1, _ := queue.Pop()
			id2, _ := queue.Pop()
			return fmt.Sprintf("%d matched %d!", id1, id2)
		}

	case "leave":
		searchAndRemove(queue, user.ID)
	}
	return "No match found"
}
