package main

import (
  "net/http"
  "encoding/json"
//  "strings"
  "log"
//  "io/ioutil"
  "github.com/gorilla/mux"
  "studentProject/driver"
)

type Student struct {
  ID string `json:"id"`
  Name string `json:"name"`
  GPA int `json:"gpa"`
}

func getStudents(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","application/json")

  db := driver.Connect()

  err := db.SQL.Ping()

  if err != nil {
    panic(err)
  }

  results, err := db.SQL.Query("SELECT id, name, gpa FROM students")
    if err != nil {
        panic(err.Error())
    }

    for results.Next() {
      var student Student

      err = results.Scan(&student.ID, &student.Name, &student.GPA)

      if err != nil {
          panic(err.Error())
      }

      json.NewEncoder(w).Encode(student)
    }

}

func getStudent(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","application/json")
  db := driver.Connect()

  err := db.SQL.Ping()

  if err != nil {
    panic(err)
  }

  vars := mux.Vars(r)
  id := vars["id"]

  var student Student

  err = db.SQL.QueryRow("SELECT id, name, gpa FROM students WHERE id = $1", id).Scan(&student.ID, &student.Name, &student.GPA)

  if err != nil {
    panic(err)
  }

  json.NewEncoder(w).Encode(student)
}

func addStudent(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","application/json")
  db := driver.Connect()

  err := db.SQL.Ping()

  if err != nil {
    panic(err)
  }


}

func updateStudent(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","application/json")
  db := driver.Connect()

  err := db.SQL.Ping()

  if err != nil {
    panic(err)
  }

  vars := mux.Vars(r)

  id := vars["id"]

  newGPA := vars["gpa"]

  updateStatement := `UPDATE students
  SET gpa = $1
  WHERE id = $2;`

  _, err = db.SQL.Exec(updateStatement, newGPA, id)

  if err != nil {
    panic(err)
  }

}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","application/json")
  db := driver.Connect()

  err := db.SQL.Ping()

  if err != nil {
    panic(err)
  }

}

func httpRequest() {
  r := mux.NewRouter().StrictSlash(true)

  r.HandleFunc("/students", getStudents).Methods("GET")
  r.HandleFunc("/student/{id}", getStudent).Methods("GET")
  r.HandleFunc("/newstudent", addStudent).Methods("POST")
  r.HandleFunc("/fixedstudent/{id}/{gpa}", updateStudent).Methods("PUT")
  r.HandleFunc("/nonstudent/{id}", deleteStudent).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {

  httpRequest()

}
