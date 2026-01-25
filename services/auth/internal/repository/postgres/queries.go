package postgres

const (
	InsertUser = `
    INSERT INTO users (email, username, password)
    VALUES ($1, $2, $3)
    RETURNING id, role, is_active, created_at, updated_at`

	SelectUserByID = `
    SELECT username, email, role, is_active, created_at, updated_at
    FROM users 
    WHERE id = $1`

	SelectUserCredentialsByEmail = `
    SELECT id, role, password
    FROM users 
    WHERE email = $1`

	SelectUsers = `
    SELECT id, 
           username, 
           email, 
           role, 
           is_active, 
           COUNT(*) OVER() as total_count
    FROM users
    ORDER BY username
    LIMIT $1 OFFSET $2;`

	UpdateUser = `
    UPDATE users
    SET username = $2, email = $3, updated_at = CURRENT_TIMESTAMP
    WHERE id = $1`

	UpdateUserRoleStatus = `
    UPDATE users
    SET role = $2, updated_at = CURRENT_TIMESTAMP
    WHERE id = $1`

	UpdateUserActiveStatus = `
    UPDATE users
    SET is_active = $2, updated_at = CURRENT_TIMESTAMP
    WHERE id = $1`

	DeleteUser = `
    DELETE FROM users
    WHERE id = $1`
)
