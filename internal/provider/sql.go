package provider

import (
	"database/sql"
	"errors"
	"github.com/IP-lo/web-rk2/internal/entities"
)

func (p *Provider) InsertUser(user entities.User) (*entities.User, error) {
	var id int

	err := p.conn.QueryRow(`INSERT INTO "user" (name, email) VALUES ($1, $2) RETURNING id`, user.Name, user.Email).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:    id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (p *Provider) SelectAllUsers() ([]*entities.User, error) {
	users := []*entities.User{}

	rows, err := p.conn.Query(`SELECT id, name, email FROM "user"`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return users, nil
		}
		return nil, err
	}

	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (p *Provider) SelectUserByID(id int) (*entities.User, error) {
	var user entities.User
	err := p.conn.QueryRow(`SELECT id, name, email FROM "user" WHERE id = $1 LIMIT 1`, id).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (p *Provider) SelectUserByName(name string) (*entities.User, error) {
	var user entities.User
	err := p.conn.QueryRow(`SELECT id, name, email FROM "user" WHERE name = $1 LIMIT 1`, name).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (p *Provider) SelectUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := p.conn.QueryRow(`SELECT id, name, email FROM "user" WHERE email = $1 LIMIT 1`, email).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (p *Provider) UpdateUserByID(id int, user entities.User) (*entities.User, error) {
	var updatedUser entities.User
	err := p.conn.QueryRow(`UPDATE "user" SET name = $1, email = $2 WHERE id = $3 RETURNING id, name, email`,
		user.Name, user.Email, id).
		Scan(&updatedUser.ID, &updatedUser.Name, &updatedUser.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (p *Provider) DeleteUserByID(id int) error {
	_, err := p.conn.Exec(`DELETE FROM "user" WHERE id = $1`,
		id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.ErrUserNotFound
		}
		return err
	}

	return nil
}
