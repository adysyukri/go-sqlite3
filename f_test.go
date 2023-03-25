package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	fmt.Println("start")
	a := selectAccount(db, 1)

	t.Logf("\na: %+v\n\n", a)

	assert.Equal(t, 1, a.id)
	assert.Equal(t, "abu", a.name)
	fmt.Println("end")
}

func TestTransfer(t *testing.T) {
	var fromBal float64 = 900
	var toBal float64 = 200
	transferMoney(db, 1, 2, 100)

	a := selectAccount(db, 1)

	assert.NotNil(t, a)

	b := selectAccount(db, 2)

	assert.NotNil(t, b)

	assert.Equal(t, fromBal, a.balance)
	assert.Equal(t, toBal, b.balance)
}

func TestTransferMany(t *testing.T) {
	db.SetMaxOpenConns(1)
	n := 5
	amount := 100.0
	var fromBal float64 = 1000 - (amount * float64(n))
	var toBal float64 = 100 + (amount * float64(n))

	errs := make(chan error)

	for i := 0; i < n; i++ {
		go func() {
			err := transferMoney(db, 1, 2, 100)

			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		fmt.Printf("\ni----%d\n", i+1)
		err := <-errs
		assert.Nil(t, err)
		a := selectAccount(db, 1)

		assert.NotNil(t, a)
		fmt.Printf("\na----%+v\n", a)

		b := selectAccount(db, 2)

		assert.NotNil(t, b)
		fmt.Printf("\nb----%+v\n", b)

		// fromBal := 1000 - (amount * float64(i+1))
		// toBal := 100 + (amount * float64(i+1))

		// assert.Equal(t, fromBal, a.balance)
		// assert.Equal(t, toBal, b.balance)
	}

	a := selectAccount(db, 1)

	assert.NotNil(t, a)

	b := selectAccount(db, 2)

	assert.NotNil(t, b)
	fmt.Printf("\n\n\na----%+v\n", a)
	fmt.Printf("\nb----%+v\n", b)
	assert.Equal(t, fromBal, a.balance)
	assert.Equal(t, toBal, b.balance)
}
