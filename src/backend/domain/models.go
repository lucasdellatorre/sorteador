package domain

type Gamble struct {
	Name     string `json:"name"      validate:"required,min=2,max=20"`
	Cpf      string `json:"cpf"       validate:"required,cpf,len=11"`
	Numbers  []int  `json:"numbers"   validate:"required,numbers"`
	RaffleID int64  `json:"raffle_id""`
}

type GambleDetail struct {
	GambleId int64  `json:"gamble_id"`
	Name     string `json:"name"`
	Cpf      string `json:"cpf"`
	Numbers  []int  `json:"numbers"`
	RaffleId int64  `json:"raffle_id"`
}

type RaffleDetail struct {
	RaffleId int64 `json:"raffle_id"`
	Raffle   []int `json:"raffle"`
	Prize    int   `json:"prize"`
	Active   bool  `json:"active"`
}

type WinnerDetail struct {
	GambleId int64  `json:"gamble_id"`
	Name     string `json:"name"`
	Cpf      string `json:"cpf"`
	Numbers  []int  `json:"numbers"`
	Prize    int    `json:"prize"`
	RaffleId int64  `json:"raffle_id"`
}

type WinnersStatsDetail struct {
	Numbers      []int          `json:"numbers"`
	Rounds       int            `json:"rounds"`
	WinnersCount int            `json:"winners_count"`
	Winners      []WinnerDetail `json:"winners"`
}
