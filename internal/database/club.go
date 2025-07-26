package database

import "github.com/kaareskytte/golf-logger/pkg/clubs"

func (db *DB) GetUserBag(userID string) ([]clubs.Club, error) {
	rows, err := db.conn.Query(
		`SELECT club_name, club_type, distance, in_bag
		FROM user_clubs
		WHERE user_id = ? AND in_bag = 1`,
		userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bag []clubs.Club
	for rows.Next() {
		var c clubs.Club
		err := rows.Scan(&c.ClubName, &c.ClubType, &c.Distance, &c.InBag)
		if err != nil {
			return nil, err
		}
		bag = append(bag, c)
	}
	return bag, nil
}
