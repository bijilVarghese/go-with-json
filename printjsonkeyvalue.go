package main
import (
	"fmt"
	"encoding/json"
    "io/ioutil"
    "os"
)

//Employees struct which contains an array of Employees
type Employees struct {
    Employees []Employee `json : "employees" `
}

//Employee struct which contains different field 
//corresponding to json object key name
type Employee struct {
    Fname string `json : "EmployeeFirstName"`
    Lname string `json : "EmployeeLastname"`
    Mailid string `json : "EmailId"`
    Phone string `json : "PhoneNumber"`
    Dob string `json : "DateOfBirth"`
}
    
func main() {
    //open our json file
	jsonFile, err := os.Open("employees.json")

    //if we os.Open returns an error the handle it
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("successfully opened employees.json")

    //defer the closing of json file
    defer jsonFile.Close()

    //read our opened file as a byte array
    byteValue, _ := ioutil.ReadAll(jsonFile)

    //we initialize our Employees array
    var employees Employees

    //we unmarshall our byte array to employees
    json.Unmarshal(byteValue, &employees)
    //we iterate through every employee within our employee array 
    //and print out each employee details
    for i := 0; i < len(employees.Employees); i++ {
        fmt.Println("EmployeeFirstName :" + employees.Employees[i].Fname)
        fmt.Println("EmployeeLastName :" + employees.Employees[i].Lname)
        fmt.Println("EmployeeMailId :" + employees.Employees[i].Mailid)
        fmt.Println("EmployeePhoneNo :" + employees.Employees[i].Phone)
        fmt.Println("EmployeeDateOfBirth :" + employees.Employees[i].Dob)
    }
}