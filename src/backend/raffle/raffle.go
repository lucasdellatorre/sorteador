package raffle

import (
	"database/sql"
	"errors"
	"itacademy/domain"
	"itacademy/util"
	"math/rand"

	"github.com/lib/pq"
)

const checkWinners = `
SELECT 
	gamble_id, 
	name, 
	cpf, 
	numbers, 
	raffle_id
FROM 
	gambles
WHERE 
	$1 @> numbers 
	AND 
	raffle_id = (
    	SELECT raffle_id
    	FROM raffles
    	ORDER BY raffle_id DESC
    	LIMIT 1
	)
`

func checkWinner(db *sql.DB, raffleNumbers []int) ([]domain.WinnerDetail, error) {
	rows, err := db.Query(checkWinners, pq.Array(raffleNumbers))

	if err != nil {
		return nil, err
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
		)

		if err := rows.Scan(&gambleId, &name, &cpf, pq.Array(&numbers), &raffleId); err != nil {
			return nil, err
		}

		parsedNumbers, err := util.ConvertNullArrayToInt(numbers)

		if err != nil {
			return nil, err
		}

		result = append(result, domain.WinnerDetail{GambleId: gambleId, Name: name, Cpf: cpf, Numbers: parsedNumbers, RaffleId: raffleId})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

const saveRaffleSQL = `
UPDATE 
	raffles
SET 
	numbers = $1
WHERE 
	raffle_id = (
		SELECT raffle_id
		FROM raffles
		ORDER BY raffle_id DESC
		LIMIT 1
	)	
RETURNING 
	raffle_id, prize
`

func GenerateRaffle(db *sql.DB) (domain.RaffleDetail, error) {
	tx, err := db.Begin()
	if err != nil {
		return domain.RaffleDetail{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	var raffleId int64
	var raffleNumbers []int
	var winners []domain.WinnerDetail
	var prize int
	used := make(map[int]bool)

	for i := 0; i < 5; i++ {
		generateUniqueRaffleNumbers(&raffleNumbers, &used)
	}

	winners, err = checkWinner(db, raffleNumbers)

	if err != nil {
		return domain.RaffleDetail{}, err
	}

	for i := 0; i < 20 && len(winners) == 0; i++ {
		// Generate raffle numbers
		generateUniqueRaffleNumbers(&raffleNumbers, &used)

		// Check for winners
		winners, err = checkWinner(db, raffleNumbers)
		if err != nil {
			return domain.RaffleDetail{}, err
		}
	}

	err = tx.QueryRow(saveRaffleSQL, pq.Array(raffleNumbers)).Scan(&raffleId, &prize)

	if err != nil {
		return domain.RaffleDetail{}, err
	}

	if len(winners) == 0 {
		return domain.RaffleDetail{RaffleId: raffleId, Active: false, Raffle: raffleNumbers, Prize: prize}, nil
	}

	err = saveWinners(tx, winners)

	if err != nil {
		return domain.RaffleDetail{}, err
	}

	if err != nil {
		return domain.RaffleDetail{}, err
	}

	return domain.RaffleDetail{RaffleId: raffleId, Active: false, Raffle: raffleNumbers, Prize: prize}, nil
}

const saveWinnersSQL = "INSERT INTO winners (gamble_id, prize) VALUES ($1, $2)"

const getPrizeSQL = "SELECT prize_amount FROM accumulated_prize WHERE prize_id = 1"

const resetAccumulatedPrizeSQL = "UPDATE accumulated_prize SET prize_amount = 0 where prize_id = 1"

func saveWinners(tx *sql.Tx, winners []domain.WinnerDetail) error {
	stmt, err := tx.Prepare(saveWinnersSQL)

	if err != nil {
		return err
	}

	var prizeAmount int

	err = tx.QueryRow(getPrizeSQL).Scan(&prizeAmount)

	if err != nil {
		return err
	}

	_, err = tx.Exec(resetAccumulatedPrizeSQL)

	if err != nil {
		return err
	}

	splittedPrize := prizeAmount / len(winners) // winners must not be 0 to enter in this func

	for i := range winners {
		// winners[i].Prize = splittedPrize
		_, err := stmt.Exec(winners[i].GambleId, splittedPrize)

		if err != nil {
			return err
		}
	}

	return nil
}

const closeRaffleSQL = `
UPDATE 
	raffles
SET 
	active = false
WHERE 
	raffle_id = (
    	SELECT raffle_id
    	FROM raffles
    	ORDER BY raffle_id DESC
    	LIMIT 1
	)
	AND active = true
RETURNING 
	raffle_id, numbers, prize
`

func CloseRaffle(db *sql.DB) (domain.RaffleDetail, error) {
	var raffleId int64
	var numbers []sql.NullInt64
	var prize int
	err := db.QueryRow(closeRaffleSQL).Scan(&raffleId, pq.Array(&numbers), &prize)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.RaffleDetail{}, errors.New("there is no raffle registered yet")
		}
		return domain.RaffleDetail{}, err
	}

	parsedNumber, err := util.ConvertNullArrayToInt(numbers)

	if err != nil {
		return domain.RaffleDetail{}, err
	}

	return domain.RaffleDetail{RaffleId: raffleId, Active: false, Raffle: parsedNumber, Prize: prize}, nil
}

const updateAccumulatedPrize = "UPDATE accumulated_prize SET prize_amount = prize_amount + $1 WHERE prize_id = 1 RETURNING prize_amount"

const createRaffleSQL = "INSERT INTO raffles (raffle_id, active, prize) VALUES (DEFAULT, DEFAULT, $1) RETURNING raffle_id, active"

func CreateRaffle(db *sql.DB) (domain.RaffleDetail, error) {
	tx, err := db.Begin()
	if err != nil {
		return domain.RaffleDetail{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	var (
		id     int64
		active bool
		prize  int
	)

	err = tx.QueryRow(updateAccumulatedPrize, util.DEFAULT_PRIZE).Scan(&prize)

	if err != nil {
		return domain.RaffleDetail{}, err
	}

	err = tx.QueryRow(createRaffleSQL, prize).Scan(&id, &active)

	raffle := domain.RaffleDetail{RaffleId: id, Active: active, Prize: prize}

	if err != nil {
		return domain.RaffleDetail{}, err
	}

	return raffle, nil
}

const startRaffleSQL = `
UPDATE 
	raffles
SET 
	active = false
WHERE 
	raffle_id = (
    	SELECT raffle_id
    	FROM raffles
    	ORDER BY raffle_id DESC
    	LIMIT 1
	)
RETURNING 
	raffle_id, prize
`

func StartRaffle(db *sql.DB) (domain.RaffleDetail, error) {
	var raffleId int64
	var prize int
	err := db.QueryRow(startRaffleSQL).Scan(&raffleId, &prize)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.RaffleDetail{}, errors.New("there is no raffle registered yet")
		}
		return domain.RaffleDetail{}, err
	}

	return domain.RaffleDetail{RaffleId: raffleId, Active: false, Raffle: nil, Prize: prize}, nil
}

const lastRaffleSQL = "SELECT raffle_id, numbers, active, prize from raffles ORDER BY raffle_id DESC LIMIT 1"

func LastRaffle(db *sql.DB) (domain.RaffleDetail, error) {
	var (
		raffleId int64
		numbers  []sql.NullInt64
		active   bool
		prize    int
	)
	err := db.QueryRow(lastRaffleSQL).Scan(&raffleId, pq.Array(&numbers), &active, &prize)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.RaffleDetail{}, errors.New("no raffle exist")
		} else {
			return domain.RaffleDetail{}, err
		}
	}
	parsedNumbers, err := util.ConvertNullArrayToInt(numbers)

	if err != nil {
		return domain.RaffleDetail{}, err
	}

	return domain.RaffleDetail{RaffleId: raffleId, Raffle: parsedNumbers, Active: active, Prize: prize}, nil
}

func generateUniqueRaffleNumbers(numbers *[]int, used *map[int]bool) {
	generated := false
	for !generated {
		num := rand.Intn(util.MAX_GAMBLE_RANGE) + 1
		if !(*used)[num] {
			*numbers = append(*numbers, num)
			(*used)[num] = true
			generated = true
		}
	}
}
