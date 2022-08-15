package repository

import "database/sql"

type TransactionRepositoryDB struct {
	db *sql.DB
}

func NewTransactionRepositoryDB(db *sql.DB) *TransactionRepositoryDB{
	return &TransactionRepositoryDB(db: db)
}

func (t *TransactionRepositoryDB) SaveTransaction(transaction domain.Transaction, creditCard domain.CreditCard) error {
	stmt, err := t.db.Prepare(query: `insert into transactions (id, credit_card_id, amount, status, description, store, created_at)
										VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		transaction.ID,
		transaction.CreditCardId,
		transaction.Amount,
		transaction.Status,
		transaction.Description,
		transaction.Store,
		transaction.CreatedAt
	)
	if err != nil {
		return err
	}
	if transaction.Status == "approved"{
		err = t.updateBalance(creditCard)
		if err != nil {
			return err
		}
	}
	
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepositoryDB) updateBalance (creditCard domain.CreditCard) error {
	_, err := t.db.Exec(query: `update credit_cards set balance = $1 where id = $2`),
	creditCard.Balance, creditCard.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepositoryDB) CreateCreditCard(creditCard domain.CreditCard) error {
	stmt, err := t.db.Prepare(query: `insert into credit_cards(id, name, number, expiration_month, expiration_year, cvv, balance, balance_limit)
										values ($1,$2,$3,$4,$5,$6,$7,$8)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		creditCard.ID,
		creditCard.Name,
		creditCard.Number,
		creditCard.ExpirationMonth,
		creditCard.ExpirationYear,
		creditCard.CVV,
		creditCard.Balance,
		creditCard.Limi
	)

	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepositoryDB) GetCreditCard(creditCard domain.CreditCard) (domain.CreditCard, error) {
	
}