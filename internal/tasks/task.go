package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

const taskFile = "tasks.json"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func LoadTasks() ([]Task, error) {
	file, err := os.ReadFile(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	var tasks []Task
	json.Unmarshal(file, &tasks)
	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	return os.WriteFile(taskFile, data, 0644)
}

func AddTask(description string) (int, error) {
	tasks, _ := LoadTasks()
	id := len(tasks) + 1
	newTask := Task{id, description, "todo", time.Now(), time.Now()}
	tasks = append(tasks, newTask)
	SaveTasks(tasks)
	return id, nil
}

func UpdateTask(id int, description string) error {
	tasks, _ := LoadTasks()
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			SaveTasks(tasks)
			return nil
		}
	}
	return errors.New("task not found")
}

func DeleteTask(id int) error {
	tasks, _ := LoadTasks()
	newTasks := []Task{}
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}
	SaveTasks(newTasks)
	return nil
}

func MarkTask(id int, status string) error {
	tasks, _ := LoadTasks()
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			SaveTasks(tasks)
			return nil
		}
	}
	return errors.New("task not found")
}

func ListTasks(filter string) {
	tasks, _ := LoadTasks()
	for _, task := range tasks {
		if filter == "" || task.Status == filter {
			fmt.Printf("[%d] %s - %s - Created: %s\n", task.ID, task.Description, task.Status, task.CreatedAt.Local().UTC().Format(time.RFC822))
		}
	}
}
