package main

import "database/sql"

const update = `
UPDATE accounts
SET balance = ?
WHERE id = ?
`

func transferMoney(db *sql.DB, fromID int, toID int, amount float64) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	row := tx.QueryRowContext(ctx, select1, fromID)

	var fromAccount account
	row.Scan(&fromAccount.id, &fromAccount.name, &fromAccount.balance)

	row = tx.QueryRowContext(ctx, select1, toID)

	var toAccount account
	row.Scan(&toAccount.id, &toAccount.name, &toAccount.balance)

	fromAccount.balance = fromAccount.balance - amount
	toAccount.balance = toAccount.balance + amount

	_, err = tx.ExecContext(ctx, update, fromAccount.balance, fromAccount.id)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, update, toAccount.balance, toAccount.id)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func selectAccount(db *sql.DB, id int) account {
	row := db.QueryRowContext(ctx, select1, id)

	var a account
	row.Scan(&a.id, &a.name, &a.balance)

	return a
}
