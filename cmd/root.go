package cmd

import (
	"fmt"
	"strconv"
	"task-tracker/internal/tasks"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "Task CLI helps manage your tasks efficiently",
	Long:  "A simple CLI application to track, update, and manage tasks.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add Command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "add [description]",
		Short: "Add a new task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := tasks.AddTask(args[0])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Printf("Task added (ID: %d)\n", id)
		},
	})

	// Update Command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "update [id] [new description]",
		Short: "Update a task description",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid ID")
				return
			}
			err = tasks.UpdateTask(id, args[1])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Task updated.")
		},
	})

	// Delete Command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid ID")
				return
			}
			err = tasks.DeleteTask(id)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Task deleted.")
		},
	})

	// Mark Task Command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "mark [id] [status]",
		Short: "Mark a task as 'todo', 'in-progress', or 'done'",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid ID")
				return
			}
			err = tasks.MarkTask(id, args[1])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Task status updated.")
		},
	})

	// List Tasks Command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "list [status]",
		Short: "List all tasks or filter by status",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			status := ""
			if len(args) > 0 {
				status = args[0]
			}
			tasks.ListTasks(status)
		},
	})
}
