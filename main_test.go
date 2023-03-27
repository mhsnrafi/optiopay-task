package main

import (
	"testing"
)

func TestEmployeeNode_AddReport(t *testing.T) {
	manager := &EmployeeNode{ID: 1, Name: "Manager"}
	report := &EmployeeNode{ID: 2, Name: "Report"}

	manager.AddReport(report)

	if len(manager.Reports) != 1 || manager.Reports[0] != report {
		t.Errorf("AddReport failed to add report to manager")
	}

	if report.Manager != manager {
		t.Errorf("AddReport failed to set manager for report")
	}
}

func TestFindCommonManager_SameManager(t *testing.T) {
	manager := &EmployeeNode{ID: 1, Name: "Manager"}
	employee1 := &EmployeeNode{ID: 2, Name: "Employee1"}
	employee2 := &EmployeeNode{ID: 3, Name: "Employee2"}

	manager.AddReport(employee1)
	manager.AddReport(employee2)

	commonManager := findCommonManager(employee1, employee2)

	if commonManager != manager {
		t.Errorf("Expected common manager to be %s, but got %s", manager.Name, commonManager.Name)
	}
}

func TestFindCommonManager_DifferentManagers(t *testing.T) {
	// Set up a tree structure with different managers for two employees
	CEO := &EmployeeNode{ID: 1, Name: "CEO"}
	VP1 := &EmployeeNode{ID: 2, Name: "VP1"}
	VP2 := &EmployeeNode{ID: 3, Name: "VP2"}
	M1 := &EmployeeNode{ID: 4, Name: "M1"}
	M2 := &EmployeeNode{ID: 5, Name: "M2"}

	CEO.AddReport(VP1)
	CEO.AddReport(VP2)
	VP1.AddReport(M1)
	VP2.AddReport(M2)

	commonManager := findCommonManager(M1, M2)

	if commonManager != CEO {
		t.Errorf("Expected common manager to be %s, but got %s", CEO.Name, commonManager.Name)
	}
}

func TestFindCommonManager_NoCommonManager(t *testing.T) {
	employee1 := &EmployeeNode{ID: 1, Name: "Employee1"}
	employee2 := &EmployeeNode{ID: 2, Name: "Employee2"}

	commonManager := findCommonManager(employee1, employee2)

	if commonManager != nil {
		t.Errorf("Expected no common manager, but got %s", commonManager.Name)
	}
}

func TestFindCommonManager_EmployeeAndManager(t *testing.T) {
	manager := &EmployeeNode{ID: 1, Name: "Manager"}
	employee := &EmployeeNode{ID: 2, Name: "Employee"}

	manager.AddReport(employee)

	commonManager := findCommonManager(manager, employee)

	if commonManager != manager {
		t.Errorf("Expected common manager to be %s, but got %s", manager.Name, commonManager.Name)
	}
}

func TestFindCommonManager_EmployeeIsCEO(t *testing.T) {
	CEO := &EmployeeNode{ID: 1, Name: "CEO"}
	VP := &EmployeeNode{ID: 2, Name: "VP"}

	CEO.AddReport(VP)

	commonManager := findCommonManager(CEO, VP)

	if commonManager != CEO {
		t.Errorf("Expected common manager to be %s, but got %s", CEO.Name, commonManager.Name)
	}
}
