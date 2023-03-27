package main

import (
	"fmt"
)

// EmployeeNode represents an employee in an organization's hierarchy.
type EmployeeNode struct {
	ID      int
	Name    string
	Reports []*EmployeeNode // A list of employees that directly report to this employee.
	Manager *EmployeeNode   // The employee that this employee directly reports to.
}

// AddReport adds an employee as a direct report to this employee and sets the manager.
func (e *EmployeeNode) AddReport(report *EmployeeNode) {
	e.Reports = append(e.Reports, report) // Add the report to the list of direct reports.
	report.Manager = e                    // Set this employee as the manager for the report.
}

// findCommonManager finds the closest common manager for two employees in the hierarchy.
func findCommonManager(a, b *EmployeeNode) *EmployeeNode {
	visited := make(map[int]bool) // Create a map to store visited employee IDs.

	// Traverse up the hierarchy for employee a and mark all visited employee IDs.
	for a != nil {
		visited[a.ID] = true
		a = a.Manager
	}

	// Traverse up the hierarchy for employee b.
	// If an employee ID is already visited, return the corresponding manager.
	for b != nil {
		if visited[b.ID] {
			return b
		}
		b = b.Manager
	}

	return nil // Return nil if no common manager is found.
}

func main() {
	CEO := &EmployeeNode{ID: 1, Name: "Claire"}
	VP1 := &EmployeeNode{ID: 2, Name: "VP1"}
	VP2 := &EmployeeNode{ID: 3, Name: "VP2"}
	M1 := &EmployeeNode{ID: 4, Name: "M1"}
	M2 := &EmployeeNode{ID: 5, Name: "M2"}
	M3 := &EmployeeNode{ID: 6, Name: "M3"}
	E1 := &EmployeeNode{ID: 7, Name: "E1"}
	E2 := &EmployeeNode{ID: 8, Name: "E2"}
	E3 := &EmployeeNode{ID: 9, Name: "E3"}

	// Set up the hierarchy
	CEO.AddReport(VP1)
	CEO.AddReport(VP2)
	VP1.AddReport(M1)
	VP1.AddReport(M2)
	VP2.AddReport(M3)
	M1.AddReport(E1)
	M3.AddReport(E2)
	M3.AddReport(E3)

	// Find the closest common manager for employees E1 and E2.
	commonManager := findCommonManager(E1, E2)

	// Output: Closest common manager for E1 and E2: VP2
	fmt.Printf("Closest common manager for E1 and E2: %s\n", commonManager.Name)
}
