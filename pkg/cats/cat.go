package cats

import "database/sql"

type Cat struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// AddCat creates a new cat
func (c *Cat) AddCat(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO cats(name, description) VALUES($1, $2) RETURNING id",
		c.Name, c.Description).Scan(&c.Id)

	if err != nil {
		return err
	}

	return nil
}

// GetAllCats returns all the cats from db
func GetAllCats(db *sql.DB) ([]Cat, error) {
	rows, err := db.Query("SELECT id, name, description FROM cats")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cats []Cat

	for rows.Next() {
		var c Cat
		if err := rows.Scan(&c.Id, &c.Name, &c.Description); err != nil {
			return nil, err
		}

		cats = append(cats, c)
	}

	return cats, nil
}
