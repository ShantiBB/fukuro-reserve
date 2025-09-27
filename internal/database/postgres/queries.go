package postgres

var UserCreate = `
	INSERT INTO users (username, first_name, last_name, email, description, password)
	VALUES (:username, :first_name, :last_name, :email, :description, :password) 
	RETURNING id
`

var UserGet = `
	SELECT id, username, first_name, last_name, email, description 
	FROM users 
	WHERE id = $1
`

var UserGetAll = `
	SELECT id, username, first_name, last_name, email, description 
	FROM users
`

var UserUpdate = `
	UPDATE users
	SET
	    username = :username, 
	    first_name = :first_name, 
	    last_name = :last_name, 
	    email = :email, 
	    description = :description
	WHERE id = :id
`

var UserDelete = `
	DELETE FROM users
	WHERE id = $1
`
