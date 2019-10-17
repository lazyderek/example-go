package main

import "fmt"

// 组合模式：又叫部分整体模式，是用于把一组相似的对象当作一个单一的对象。组合模式依据树形结构来组合对象，用来表示部分以及整体层次。
// 这种类型的设计模式属于结构型模式，它创建了对象组的树形结构。
//这种模式创建了一个包含自己对象组的类。该类提供了修改相同对象组的方式

type Employee struct {
	Name       string
	Department string
	Salary     float32
	Employees  []*Employee
}

func NewEmployee(name, dept string, salary float32) *Employee {
	return &Employee{
		Name:       name,
		Department: dept,
		Salary:     salary,
	}
}

func (e *Employee) Add(em *Employee) {
	e.Employees = append(e.Employees, em)
}

func main() {
	boss := NewEmployee("alice", "Boss", 0)

	headHR := NewEmployee("娜塔莎", "Head HR", 0)
	headIT := NewEmployee("captain", "Head IT", 0)

	hr1 := NewEmployee("hr1", "HR", 0)
	hr2 := NewEmployee("hr2", "HR", 0)

	it1 := NewEmployee("banner", "IT", 0)
	it2 := NewEmployee("peter", "IT", 0)

	boss.Add(headHR)
	boss.Add(headIT)

	headHR.Add(hr1)
	headHR.Add(hr2)

	headIT.Add(it1)
	headIT.Add(it2)

	fmt.Println(boss.Name, boss.Department)
	for _, headE := range boss.Employees {
		fmt.Println(headE.Name, headE.Department)
		for _, e := range headE.Employees {
			fmt.Println(e.Name, e.Department)
		}
	}
}
