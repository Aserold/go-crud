package repository

const (
	createUserQuery = `INSERT INTO users (username, email, age)
	VALUES ($1, $2, $3)
	RETURNING *`

	updateUserQuery = `UPDATE users
	SET username = COALESCE(NULLIF($1, ''), username),
		email = COALESCE(NULLIF($2, ''), email),
		age = COALESCE(NULLIF($3, 0), age)
	WHERE id = $4
	RETURNING *`

	getUserQuery = `SELECT * FROM users WHERE id = $1`

	listUserQuery = `SELECT * FROM users`

	deleteUserQuery = `DELETE FROM users WHERE id = $1`

	getTotal = `SELECT COUNT(id) FROM users`

	findUserByEmail = `SELECT * FROM users 
				 		WHERE email = $1`
)