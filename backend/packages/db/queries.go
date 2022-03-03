package db

const (
	CheckUserExists                 = `SELECT true from users WHERE email = $1`
	LoginQuery                      = `SELECT * from users WHERE email = $1`
	UpdateUserPasswordQuery         = `UPDATE users SET password = $2 WHERE id = $1`
	DeleteUser                      = `DELETE FROM users WHERE email = $1`
	CreateUserQuery                 = `INSERT INTO users(id, name, password, email) VALUES (DEFAULT, $1 , $2, $3);`
	GetUserByIDQuery                = `SELECT * FROM users WHERE id = $1`
	GetUserByEmailQuery             = `SELECT * FROM users WHERE email = $1`
	CreateRestaurant                = `INSERT INTO restaurants(id, rname, location, rating, rtype, phone, zipcode) VALUES (DEFAULT, $1 , $2, $3, $4, $5, $6);`
	GetAllRestaurantsQuery          = `SELECT * FROM restaurants`
	GetAllRestaurantsQueryByZipCode = `SELECT * FROM restaurants WHERE zipcode > $1 and zipcode < $2`
)
