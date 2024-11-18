package queries

const (
	INSERT_USER      = "INSERT INTO users (id,username, email, password, role_id) values ($1,$2,$3,$4,$5)"
	SELECT_ALL_USERS = "SELECT * FROM users"
)
