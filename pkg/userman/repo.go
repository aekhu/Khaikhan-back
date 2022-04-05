package userman

import (
	"git.bolor.net/orgil/go-gin-template/pkg/common"
)

func GetUserByBolorID(userID int) (customer *Customer, err error) {
	c := new(Customer)
	result := common.Db.QueryRow(`
		SELECT 
			customer_id, user_id, name, email, phone, reg_at,
		    role, branch_id
		FROM customer
		WHERE user_id = $1;`, userID)

	err = result.Scan(
		&c.CustomerID, &c.UserID, &c.Name,
		&c.Email, &c.Phone, &c.RegAt,
		&c.Role, &c.BranchID)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func UserIDExist(id int) (*Customer, bool) {
	c := new(Customer)
	result := common.Db.QueryRow(`
		SELECT 
		       customer_id, user_id, name, 
		       email, phone, reg_at, 
		       role, branch_id 
		FROM customer 
		WHERE user_id = $1;`, id)
	err := result.Scan(&c.CustomerID, &c.UserID, &c.Name,
		&c.Email, &c.Phone, &c.RegAt, &c.Role,
		&c.BranchID)
	if err != nil {
		return nil, false
	}
	return c, true
}

func CreateUser(UserID int, name, email, phone string) (int, error) {
	var id int
	row := common.Db.QueryRow(`
		INSERT INTO customer(user_id, name, email, phone, role) 
		VALUES	($1,$2,$3,$4,0) RETURNING customer_id ;`,
		UserID, name, email, phone)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func GetUserInfo(CustomerID int) (*Customer, error) {
	c := new(Customer)
	result := common.Db.QueryRow(`SELECT customer_id,user_id,name,email,phone,reg_at,role,branch_id FROM customer WHERE customer_id =$1 ;`, CustomerID)
	err := result.Scan(&c.CustomerID, &c.UserID, &c.Name, &c.Email, &c.Phone, &c.RegAt, &c.Role, &c.BranchID)
	if err != nil {
		return nil, err
	}
	return c, nil
}
