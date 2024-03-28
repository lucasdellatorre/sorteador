package raffle

import (
	"database/sql"
	"itacademy/domain"
	"itacademy/util"

	"github.com/lib/pq"
)

// a. a lista de números sorteados
// b. quantas rodadas de sorteio foram realizadas
// c. a quantidade de apostas vencedoras

const getWinners = `
SELECT
	g.gamble_id,
    g.name,
    g.cpf,
    g.numbers,
    g.raffle_id,
	w.prize
FROM
    gambles g
JOIN
    winners w ON g.gamble_id = w.gamble_id
WHERE
	g.raffle_id = $1
`

func WinnersStats(db *sql.DB) (domain.WinnersStatsDetail, error) {
	raffle, err := LastRaffle(db)

	if err != nil {
		return domain.WinnersStatsDetail{}, err
	}

	rows, err := db.Query(getWinners, raffle.RaffleId)

	if err != nil {
		return domain.WinnersStatsDetail{}, err
	}

	defer rows.Close()

	var result []domain.WinnerDetail

	for rows.Next() {
		var (
			gambleId int64
			name     string
			cpf      string
			numbers  []sql.NullInt64
			raffleId int64
			prize    int
		)

		if err := rows.Scan(&gambleId, &name, &cpf, pq.Array(&numbers), &raffleId, &prize); err != nil {
			return domain.WinnersStatsDetail{}, err
		}

		parsedNumbers, err := util.ConvertNullArrayToInt(numbers)

		if err != nil {
			return domain.WinnersStatsDetail{}, err
		}

		result = append(result, domain.WinnerDetail{GambleId: gambleId, Name: name, Cpf: cpf, Numbers: parsedNumbers, RaffleId: raffleId, Prize: prize})
	}
	if err := rows.Err(); err != nil {
		return domain.WinnersStatsDetail{}, err
	}

	return domain.WinnersStatsDetail{
		Numbers:      raffle.Raffle,
		Rounds:       (len(raffle.Raffle) - 4),
		WinnersCount: len(result),
		Winners:      result,
	}, nil
}

const allGambleNumbers = `
SELECT 
	numbers
FROM
	gambles
`

func AllGambleNumbers(db *sql.DB) ([]int, error) {
	rows, err := db.Query(allGambleNumbers)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]int, 51)

	for rows.Next() {
		var (
			numbers []sql.NullInt64
		)

		if err := rows.Scan(pq.Array(&numbers)); err != nil {
			return nil, err
		}

		parsedNumbers, err := util.ConvertNullArrayToInt(numbers)

		if err != nil {
			return nil, err
		}

		for _, num := range parsedNumbers {
			result[num]++
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
