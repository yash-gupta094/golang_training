package main

import "fmt"

func main() {
	pointersExercise()

	structureExercise()

	interfaceExercise()

	employeeExercise()
}

type DataMap map[string]string

func employeeExercise() {
	var dataMap DataMap
	dataMap = make(map[string]string)
	dataMap.Insert("20", "Yash")
	dataMap.Insert("10", "Shan")

	fmt.Println("Data in map: ", dataMap.ReadAll())
	fmt.Println("Data ToByte: ", dataMap.ToByte())

	fmt.Println("Keys of the map:", dataMap.GetKeys())
	fmt.Println("Valuesof the map:", dataMap.GetVals())

	fmt.Println(dataMap.Update("20", "Tejas"))
	fmt.Println(dataMap.Update("23", "Akshat"))
	fmt.Println("Valuesof the map:", dataMap.GetVals())

}
func (dm DataMap) ToByte() []byte {
	return []byte(fmt.Sprint(dm))
}

func (dm DataMap) Read(key string) string {
	return dm[key]
}

func (dm DataMap) ReadAll() DataMap {
	return dm
}

func (dm DataMap) Delete(key string) {
	delete(dm, key)
}

func (dm DataMap) ToString() string {
	return fmt.Sprint(dm)
}

func (dm DataMap) Insert(key, val string) {
	if dm != nil {
		dm[key] = val
	}
}

func (dm DataMap) Update(key, val string) bool {
	if dm == nil {
		return false
	}
	_, ok := dm[key]
	if !ok {
		return ok
	}

	dm[key] = val
	return true
}

func (dm DataMap) GetKeys() []string {
	if dm == nil {
		return nil
	}
	var slice []string
	for k, _ := range dm {
		slice = append(slice, k)
	}
	return slice
}

func (dm DataMap) GetVals() []string {
	if dm == nil {
		return nil
	}
	var slice []string
	for _, val := range dm {
		slice = append(slice, val)
	}
	return slice
}

func interfaceExercise() {
	emp1 := &Employee{id: 2, age: 16, name: "shan"}
	var iStringer Stringer
	iStringer = emp1
	fmt.Println(iStringer.ToString())
}

type Stringer interface {
	ToString() string
}

func (e *Employee) ToString() string {
	return fmt.Sprint(e)
}

type Employee struct {
	id     int
	age    uint8
	name   string
	salary int
}

func structureExercise() {
	var e1 Employee
	e1 = Employee{id: 101, age: 24, name: "Yash", salary: 300}
	e2 := Employee{id: 102, age: 24, name: "Shan", salary: 500}
	printAllDetails(e2)
	var e5 *Employee
	e5 = &Employee{id: 106, name: "Happy"}
	fmt.Println(e5)
	fmt.Println(e1.name)
}

func printAllDetails(e Employee) {
	fmt.Println("Employee details of employee no:", e.id)
	fmt.Println("Id: ", e.id, ", Age: ", e.age, ", Name: ", e.name, ", Salary: ", e.salary)
}

func pointersExercise() {
	// var d *int
	x := 100
	fmt.Println(x)
	increment(&x)
	fmt.Println(x)
}

func increment(d *int) {
	*d++
}
