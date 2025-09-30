package postgres

var UserCreate = `
    INSERT INTO users (username, email, password)
    VALUES ($1, $2, $3)
    RETURNING id, role, is_active, created_at, updated_at
`

var UserGetByID = `
    SELECT username, email, role, is_active, created_at, updated_at
    FROM users 
    WHERE id = $1
`

var UserGetByUsername = `
    SELECT id, email, role, is_active, created_at, updated_at
    FROM users 
    WHERE username = $1
`

var UserGetByEmail = `
    SELECT id, username, role, is_active, created_at, updated_at
    FROM users 
    WHERE email = $1
`

var UserGetAll = `
    SELECT id, username, email, role, is_active, created_at, updated_at
    FROM users
`

var UserUpdate = `
    UPDATE users
    SET username = $1, email = $2, updated_at = CURRENT_TIMESTAMP
    WHERE id = $3
`

var UserDelete = `
    DELETE FROM users
    WHERE id = $1
`
