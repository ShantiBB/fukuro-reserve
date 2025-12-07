package postgres

const (
	UserCreate = `
    INSERT INTO users (email, password)
    VALUES ($1, $2)
    RETURNING id, role, is_active, created_at, updated_at`

	UserGetByID = `
    SELECT username, email, role, is_active, created_at, updated_at
    FROM users 
    WHERE id = $1`

	UserGetCredentialsByEmail = `
    SELECT id, role, password
    FROM users 
    WHERE email = $1`

	UserGetAll = `
    SELECT id, username, email, role, is_active
    FROM users
    ORDER BY id
    LIMIT $1 OFFSET $2;`

	UserUpdate = `
    UPDATE users
    SET username = $1, email = $2, updated_at = CURRENT_TIMESTAMP
    WHERE id = $3`

	UserUpdateRoleStatus = `
    UPDATE users
    SET role = $1, updated_at = CURRENT_TIMESTAMP
    WHERE id = $2`

	UserUpdateActiveStatus = `
    UPDATE users
    SET is_active = $1, updated_at = CURRENT_TIMESTAMP
    WHERE id = $2`

	UserDelete = `
    DELETE FROM users
    WHERE id = $1`

	UserGetCountRows = `SELECT COUNT(*) FROM users;`
)
