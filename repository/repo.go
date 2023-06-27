package repository

import (
	"database/sql"
	"fmt"

	M "github.com/farhaanfsk/Go-Gorilla-Mux-CRUD/models"
	"github.com/google/uuid"
)

type EmployeeRepo struct {
	db *sql.DB
}

func (r *EmployeeRepo) Save(e M.Employee) M.Employee {
	var aid uuid.UUID
	if (e.Address != M.Address{}) {
		err := r.db.QueryRow("INSERT INTO \"Employee\".\"Address\" (city, state, country) values($1,$2,$3) RETURNING id",
			e.Address.City, e.Address.State, e.Address.Country).Scan(&aid)

		if err != nil {
			fmt.Println("An error occured while executing query: ", err)
		}
	}
	e.Address.Id = aid
	var eid uuid.UUID
	err := r.db.QueryRow("INSERT INTO \"Employee\".\"Employee\"(name, address_id) values($1,$2) RETURNING id", e.Name, aid).Scan(&eid)
	if err != nil {
		fmt.Println("An error occured while executing query: ", err)
	}
	e.Id = eid
	return e
}

func (r *EmployeeRepo) GetAllEmployees() []M.Employee {
	e := []M.Employee{}
	a := []M.Address{}
	rows := executeQuery("SELECT * FROM \"Employee\".\"Address\"", r.db)
	for i := 0; rows.Next(); i++ {
		var id uuid.UUID
		var city string
		var state string
		var country string
		rows.Scan(&id, &city, &state, &country)
		a = append(a, M.Address{
			Id:      id,
			City:    city,
			State:   state,
			Country: country,
		})
	}
	rows = executeQuery("SELECT * FROM \"Employee\".\"Employee\"", r.db)
	for i := 0; rows.Next(); i++ {
		var id uuid.UUID
		var name string
		var aid uuid.UUID
		rows.Scan(&id, &name, &aid)
		e = append(e, M.Employee{
			Id:   id,
			Name: name,
		})
		for j := range a {
			if a[j].Id == aid {
				e[i].Address = a[j]
			}
		}
	}
	return e
}

func (r *EmployeeRepo) GetEmployeeById(id uuid.UUID) M.Employee {
	e := M.Employee{}
	a := M.Address{}
	var aid uuid.UUID

	rows := executeQueryWithId("SELECT * FROM \"Employee\".\"Employee\" where id = $1", id, r.db)
	if rows.Next() {
		rows.Scan(&e.Id, &e.Name, &aid)
	}

	rows = executeQueryWithId("SELECT * FROM \"Employee\".\"Address\" where id = $1", aid, r.db)
	if rows.Next() {
		rows.Scan(&a.Id, &a.City, &a.State, &a.Country)
	}
	e.Address = a
	return e
}

func executeQueryWithId(query string, id uuid.UUID, db *sql.DB) *sql.Rows {
	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Println("An error occured while executing query: ", err)
	}
	return rows
}

func executeQuery(query string, db *sql.DB) *sql.Rows {
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("An error occured while executing query: ", err)
	}
	return rows
}

func (r *EmployeeRepo) DeleteEmployeeById(id uuid.UUID) bool {
	var aid uuid.UUID
	rows := executeQueryWithId("DELETE FROM \"Employee\".\"Employee\" where id = $1 RETURNING address_id", id, r.db)
	if rows.Next() {
		rows.Scan(&aid)
	}
	executeQueryWithId("DELETE FROM \"Employee\".\"Address\" where id = $1 RETURNING id", aid, r.db)
	return true
}

func (r *EmployeeRepo) UpdateEmployee(e M.Employee) bool {
	existingEmp := r.GetEmployeeById(e.Id)
	if existingEmp.Id == e.Id {
		_, err := r.db.Query("update \"Employee\".\"Employee\" set name = $1 where id = $2", e.Name, e.Id)
		if err != nil {
			fmt.Println("An error occured while executing query: ", err)
			return false
		}
		_, err = r.db.Query("update \"Employee\".\"Address\" set city = $1 , state = $2, country = $3 where id = $4", e.Address.City, e.Address.State, e.Address.Country, e.Address.Id)
		if err != nil {
			fmt.Println("An error occured while executing query: ", err)
			return false
		}
		return true
	} else {
		return false
	}

}

func NewRepo(db *sql.DB) EmployeeRepo {
	return EmployeeRepo{db}
}
