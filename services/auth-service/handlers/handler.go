package handlers

import db "github.com/shehbaazsl/go-commerce-microservices/auth-service/db/queries"

func AuthenticateUser(q *db.Queries, e mail, password string) (bool, error) {
	user, err := q.(email)
	if err != nil {
		return false, err
	}

	if user == nil {
		return false, nil // User does not exist
	}

	if user.PasswordHash != hashPassword(password) {
		return false, nil // Incorrect password
	}

	return true, nil // Authentication successful
}
