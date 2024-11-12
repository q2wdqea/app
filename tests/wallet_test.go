package tests

import (
	"app/pkg/model"
	"app/pkg/service"
	"context"
	"testing"
)

func TestWithdraw(t *testing.T) {
	ctx := context.Background()
	db, cache := NewDBAndCache(t)
	walletService := service.NewWallet(db, cache)
	// try test
	w1 := &model.WithdrawRequest{
		UserId: 1,
		Amount: 10,
	}
	if err := walletService.Withdraw(ctx, w1); err != nil {
		t.Errorf("withdraw err: %v", err)
	}
	t.Log(w1)
	w2 := &model.WithdrawRequest{
		UserId: -1,
		Amount: 10,
	}
	// try user err
	if err := walletService.Withdraw(ctx, w2); err != nil {
		t.Errorf("withdraw err: %v", err)
	}
	t.Log(w2)
	w3 := &model.WithdrawRequest{
		UserId: 1,
		Amount: -10,
	}
	// try amount err
	if err := walletService.Withdraw(ctx, w3); err != nil {
		t.Errorf("withdraw err: %v", err)
	}
	t.Log(w3)
	// try test
	w4 := &model.WithdrawRequest{
		UserId: 1,
		Amount: 8.9,
	}
	if err := walletService.Withdraw(ctx, w4); err != nil {
		t.Errorf("withdraw err: %v", err)
	}
	t.Log(w4)
}

func TestDeposit(t *testing.T) {
	ctx := context.Background()
	db, cache := NewDBAndCache(t)
	walletService := service.NewWallet(db, cache)
	// try test
	d1 := &model.DepositRequest{
		UserId: 1,
		Amount: 10,
	}
	if err := walletService.Deposit(ctx, d1); err != nil {
		t.Errorf("deposit err: %v", err)
	}
	t.Log(d1)
	// try user err
	d2 := &model.DepositRequest{
		UserId: -1,
		Amount: 10,
	}
	if err := walletService.Deposit(ctx, d2); err != nil {
		t.Errorf("deposit err: %v", err)
	}
	t.Log(d2)
	d3 := &model.DepositRequest{
		UserId: 1,
		Amount: -10,
	}
	// try amount err
	if err := walletService.Deposit(ctx, d3); err != nil {
		t.Errorf("deposit err: %v", err)
	}
	t.Log(d3)
	// try test
	d4 := &model.DepositRequest{
		UserId: 1,
		Amount: 13.1,
	}
	if err := walletService.Deposit(ctx, d4); err != nil {
		t.Errorf("deposit err: %v", err)
	}
	t.Log(d4)
}

func TestBalance(t *testing.T) {
	ctx := context.Background()
	db, cache := NewDBAndCache(t)
	walletService := service.NewWallet(db, cache)
	// try test
	wallet1, err := walletService.FindOne(ctx, 1)
	if err != nil {
		t.Errorf("balance err: %v", err)
	}
	t.Log(wallet1)
	// try test
	wallet2, err := walletService.FindOne(ctx, 2)
	if err != nil {
		t.Errorf("balance err: %v", err)
	}
	t.Log(wallet2)
	// try test
	wallet3, err := walletService.FindOne(ctx, 3)
	if err != nil {
		t.Errorf("balance err: %v", err)
	}
	t.Log(wallet3)
	// try test
	wallet4, err := walletService.FindOne(ctx, -3)
	if err != nil {
		t.Errorf("balance err: %v", err)
	}
	t.Log(wallet4)
}

func TestTransfer(t *testing.T) {
	ctx := context.Background()
	db, cache := NewDBAndCache(t)
	walletService := service.NewWallet(db, cache)
	// try test
	d1 := &model.TransferRequest{
		FromId: 1,
		ToId:   2,
		Amount: 10,
	}
	if err := walletService.Transfer(ctx, d1); err != nil {
		t.Errorf("transfer err: %v", err)
	}
	t.Log(d1)
	// try user err
	d2 := &model.TransferRequest{
		FromId: -1,
		ToId:   2,
		Amount: 10,
	}
	if err := walletService.Transfer(ctx, d2); err != nil {
		t.Errorf("transfer err: %v", err)
	}
	t.Log(d2)
	d3 := &model.TransferRequest{
		FromId: -1,
		ToId:   2,
		Amount: -10,
	}
	// try amount err
	if err := walletService.Transfer(ctx, d3); err != nil {
		t.Errorf("transfer err: %v", err)
	}
	t.Log(d3)

	// try test
	d4 := &model.TransferRequest{
		FromId: -1,
		ToId:   -2,
		Amount: 10,
	}
	if err := walletService.Transfer(ctx, d4); err != nil {
		t.Errorf("transfer err: %v", err)
	}
	t.Log(d1)
	// try user err
	d5 := &model.TransferRequest{
		FromId: 1,
		ToId:   2,
		Amount: -10,
	}
	if err := walletService.Transfer(ctx, d5); err != nil {
		t.Errorf("transfer err: %v", err)
	}
	t.Log(d2)
	d6 := &model.TransferRequest{
		FromId: -1,
		ToId:   -2,
		Amount: -10,
	}
	// try amount err
	if err := walletService.Transfer(ctx, d6); err != nil {
		t.Errorf("transfer err: %v", err)
	}
	t.Log(d3)
	// try test
	d7 := &model.TransferRequest{
		FromId: 1,
		ToId:   2,
		Amount: 10.8,
	}
	if err := walletService.Transfer(ctx, d7); err != nil {
		t.Errorf("transfer err: %v", err)
	}
	t.Log(d1)
}
