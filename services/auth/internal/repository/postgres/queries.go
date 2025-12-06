package postgres

const UserCreate = `
    INSERT INTO users (username, email, password)
    VALUES ($1, $2, $3)
    RETURNING id, role, is_active, created_at, updated_at
`

const UserGetByID = `
    SELECT username, email, role, is_active, created_at, updated_at
    FROM users 
    WHERE id = $1
`

const UserGetCredentialsByEmail = `
    SELECT id, role, password
    FROM users 
    WHERE email = $1
`

const UserGetAll = `
    SELECT id, username, email, role, is_active, created_at, updated_at
    FROM users
    ORDER BY id
    LIMIT $1 OFFSET $2;
`

const UserUpdate = `
    UPDATE users
    SET username = $1, email = $2, updated_at = CURRENT_TIMESTAMP
    WHERE id = $3
`

const UserDelete = `
    DELETE FROM users
    WHERE id = $1
`
