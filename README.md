## Bureaucr.at Corporate Directory
Bureaucr.at is a typical hierarchical organisation. Claire, its CEO, has a hierarchy of employees reporting to her, and each employee can have a list of other employees reporting to him/her. An employee with at least one report is called a Manager.

This project implements a corporate directory for Bureaucr.at with an interface to find the closest common Manager (i.e., farthest from the CEO) between two employees. You may assume that all employees eventually report up to the CEO.

## Features
1. In-memory directory structure 
2. Find the closest common Manager between two employees

## Implementation Details
The implementation is done in Go, following these guidelines:

1. Resolve ambiguity with assumptions 
2. The directory should be an in-memory structure 
3. A Manager should link to Employees and not the other way around

## Getting Started
To get started with the project, follow these steps:

1. Clone the repository
```json
git clone https://github.com/mhsnrafi/optiopay-task.git
```
2. Navigate to the project directory 
3. Run go build to compile the project 
4. Run ./main to execute the program

## Usage
The corporate directory is implemented using an EmployeeNode struct, which represents an employee in the organization's hierarchy. The EmployeeNode struct has the following properties:

1. ID: Employee's ID (integer)
2. Name: Employee's name (string)
3. Reports: A list of employees that directly report to this employee (slice of *EmployeeNode)
4. Manager: The employee that this employee directly reports to (*EmployeeNode)

You can use the AddReport method to add an employee as a direct report to a Manager and set the Manager for the employee. The findCommonManager function is used to find the closest common Manager for two employees in the hierarchy.


## Code Logic
The corporate directory is built using the EmployeeNode struct, which represents an employee in the organization's hierarchy. Each employee node contains an ID, a name, a list of direct reports, and a reference to their manager. The organization's hierarchy is established by connecting employees using the AddReport method, which appends a direct report to the employee's list of direct reports and sets the manager for the direct report.

The main function in finding the closest common manager between two employees is findCommonManager. The function takes two employee nodes as input and returns a reference to the closest common manager node. The function's logic can be broken down into the following steps:

1. Initialize a map called visited to store visited employee IDs.
2. Traverse up the hierarchy for the first employee (a), marking all visited employee IDs in the visited map.
3. Traverse up the hierarchy for the second employee (b). If an employee ID is found in the visited map, return the corresponding manager as the closest common manager.
4. If no common manager is found, return nil.

The time complexity of the findCommonManager function is O(N), where N is the maximum depth of the organization's hierarchy. This complexity is due to the two separate traversals up the hierarchy for both employees. The space complexity is also O(N), as the visited map stores the IDs of the visited employees during the traversal.

In the example provided in the main.go file, a sample employee hierarchy is set up, and the findCommonManager function is used to find the closest common manager for two employees, E1 and E2. The output of the program displays the name of the closest common manager
### Output of the program is:
The hierarchy is set up using the AddReport method, and the closest common manager for employees E1 and E2 is found using the findCommonManager function. The output of the program is:
```json
Closest common manager for E1 and E2: Claire
```
## Tests
The project includes tests for the following scenarios:

1. Test **EmployeeNode**.AddReport method 
2. Test **findCommonManager** when both employees have the same Manager 
3. Test **findCommonManager** when both employees have different Managers 
4. Test **findCommonManager** when there is no common Manager 
5. Test **findCommonManager** when one employee is the Manager of the other 
6. Test **findCommonManager** when one employee is the CEO

To run the tests, simply execute the **go test** command in the project directory.
