package users

import (
	"github.com/jmoiron/sqlx"
	"github.com/oskov/megabot/internal/db"
)

type User struct {
	Id   int64 `db:"id"`
	TgId int64 `db:"tgId"`

	Name   string  `db:"name"`
	Power  float64 `db:"power"`
	Title  int64   `db:"title"`
	Energy float64 `db:"energy"`

	Money    float64 `db:"money"`
	TotalAlc float64 `db:"total_alc"`
}

type BattleRecord struct {
	Id            int64   `db:"id"`
	AttackerId    int64   `db:"attackerId"`
	AttackerPower float64 `db:"attackerPower"`
	DefenderId    int64   `db:"defenderId"`
	DefenderPower float64 `db:"defenderPower"`
	Winner        int64   `db:"winner"` // 0 attacker, 1 defender
}

func (b *BattleRecord) AttackerWon() {
	b.Winner = 0
}

func (b *BattleRecord) DefenderWon() {
	b.Winner = 1
}

type UserRepository struct {
	db *sqlx.DB
}

var uRep *UserRepository

func GetUserRepository() (*UserRepository, error) {
	if uRep != nil {
		return uRep, nil
	}

	d, err := db.GetDb()
	if err != nil {
		return nil, err
	}
	uRep = NewUserRepository(d)
	return uRep, nil
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user *User) error {
	query := "INSERT INTO users (tgId, name, power, title, energy) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, user.TgId, user.Name, user.Power, user.Title, user.Energy)
	return err
}

func (r *UserRepository) ReadAll() ([]User, error) {
	var users []User
	query := "SELECT * FROM users ORDER BY title DESC , power DESC, total_alc DESC"
	err := r.db.Select(&users, query)
	return users, err
}

func (r *UserRepository) Read(id int64) (*User, error) {
	var user User
	query := "SELECT * FROM users WHERE id = ?"
	err := r.db.Get(&user, query, id)
	return &user, err
}

func (r *UserRepository) ReadByTgId(tgId int64) (*User, error) {
	var user User
	query := "SELECT * FROM users WHERE tgId = ?"
	err := r.db.Get(&user, query, tgId)
	return &user, err
}

func (r *UserRepository) Update(user *User) error {
	query := "UPDATE users SET tgId = ?, name = ?, power = ?, title = ?, energy = ? WHERE id = ?"
	_, err := r.db.Exec(query, user.TgId, user.Name, user.Power, user.Title, user.Energy, user.Id)
	return err
}

func (r *UserRepository) Delete(id int64) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UserRepository) IncreaseEnergy(userId int64, increaseBy float64) error {
	query := "UPDATE users SET energy = energy + ? WHERE id = ?"
	_, err := r.db.Exec(query, increaseBy, userId)
	return err
}

func (r *UserRepository) IncreaseMoney(userId int64, increaseBy float64) error {
	query := "UPDATE users SET money = money + ? WHERE id = ?"
	_, err := r.db.Exec(query, increaseBy, userId)
	return err
}

func (r *UserRepository) IncreaseTotalAlc(userId int64, increaseBy float64) error {
	query := "UPDATE users SET total_alc = total_alc + ? WHERE id = ?"
	_, err := r.db.Exec(query, increaseBy, userId)
	return err
}

func (r *UserRepository) IncreaseEnergyGlobal() error {
	query := "UPDATE users SET energy = energy + 0.5 WHERE energy < 10"
	_, err := r.db.Exec(query)
	return err
}

func (r *UserRepository) IncreasePower(userId int64, increaseBy float64) error {
	query := "UPDATE users SET power = power + ? WHERE id = ?"
	_, err := r.db.Exec(query, increaseBy, userId)
	return err
}

func (r *UserRepository) IncreaseTitle(userId int64) error {
	query := "UPDATE users SET title = title + 1 WHERE id = ?"
	_, err := r.db.Exec(query, userId)
	return err
}

type BattleRecordRepository struct {
	db *sqlx.DB
}

var brRep *BattleRecordRepository

func GetBattleRecordRepository() (*BattleRecordRepository, error) {
	if brRep != nil {
		return brRep, nil
	}

	d, err := db.GetDb()
	if err != nil {
		return nil, err
	}

	brRep = NewBattleRecordRepository(d)

	return brRep, nil
}

func NewBattleRecordRepository(db *sqlx.DB) *BattleRecordRepository {
	return &BattleRecordRepository{db}
}

func (r *BattleRecordRepository) Create(record *BattleRecord) error {
	query := "INSERT INTO battle_records (attackerId, attackerPower, defenderId, defenderPower, winner) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, record.AttackerId, record.AttackerPower, record.DefenderId, record.DefenderPower, record.Winner)
	return err
}

func (r *BattleRecordRepository) ReadForUser(id int64) ([]BattleRecord, error) {
	var records []BattleRecord
	query := "SELECT * FROM battle_records WHERE attackerId = ? OR defenderId = ?"
	err := r.db.Select(&records, query, id, id)
	return records, err
}

func (r *BattleRecordRepository) Update(record *BattleRecord) error {
	query := "UPDATE battle_records SET attackerId = ?, attackerPower = ?, defenderId = ?, defenderPower = ?, winner = ? WHERE id = ?"
	_, err := r.db.Exec(query, record.AttackerId, record.AttackerPower, record.DefenderId, record.DefenderPower, record.Winner, record.Id)
	return err
}

func (r *BattleRecordRepository) Delete(id int64) error {
	query := "DELETE FROM battle_records WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *BattleRecordRepository) GetRecordsForUser(userId int64) ([]BattleRecord, error) {
	var records []BattleRecord
	query := "SELECT * FROM battle_records WHERE attackerId = ? OR defenderId = ?;"
	err := r.db.Select(&records, query, userId, userId)
	return records, err
}
