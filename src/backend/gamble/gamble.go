package gamble

import (
	"database/sql"
	"errors"
	"itacademy/domain"
	"itacademy/raffle"
	"itacademy/util"

	"github.com/lib/pq"
)

const createSQL = "INSERT INTO gambles (name, cpf, numbers, raffle_id) VALUES ($1, $2, $3, $4) RETURNING gamble_id"

func Create(db *sql.DB, gambler *domain.Gamble) (int64, error) {
	raffle, err := raffle.LastRaffle(db)
	gambler.RaffleID = raffle.RaffleId

	if err != nil {
		return 0, err
	}

	if !raffle.Active {
		return 0, errors.New("raffle not active")
	}

	var gambleId int64

	err = db.QueryRow(createSQL, gambler.Name, gambler.Cpf, pq.Array(gambler.Numbers), gambler.RaffleID).Scan(&gambleId)

	if err != nil {
		return 0, err
	}

	return gambleId, nil
}

const listSQL = "SELECT gamble_id, name, cpf, numbers, raffle_id FROM gambles"

func List(db *sql.DB) ([]domain.GambleDetail, error) {
	rows, err := db.Query(listSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.GambleDetail

	for rows.Next() {
		var (
			id       int64
			name     string
			cpf      string
			numbers  []sql.NullInt64
			raffleId int64
		)

		if err := rows.Scan(&id, &name, &cpf, pq.Array(&numbers), &raffleId); err != nil {
			return nil, err
		}

		parsedNumbers, err := util.ConvertNullArrayToInt(numbers)

		if err != nil {
			return nil, err
		}

		result = append(result, domain.GambleDetail{GambleId: id, Name: name, Cpf: cpf, Numbers: parsedNumbers, RaffleId: raffleId})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
