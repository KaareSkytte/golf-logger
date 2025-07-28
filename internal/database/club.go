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

func (db *DB) UpdateClubStatus(userID, clubName string, inBag bool) error {
	_, err := db.conn.Exec(
		`UPDATE user_clubs
		SET in_bag = ?
		WHERE user_id = ? AND club_name = ?`,
		inBag, userID, clubName)
	return err
}

func (db *DB) UpdateClubDistance(userID, clubName string, distance int) error {
	_, err := db.conn.Exec(
		`UPDATE user_clubs
		SET distance = ?
		WHERE user_id = ? AND club_name = ?`,
		distance, userID, clubName)
	return err
}
