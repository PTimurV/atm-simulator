package services

import (
	"atm-simulator/models"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

type AccountService struct {
    Accounts map[string]*models.Account
    Logins   map[string]string
    mu       sync.Mutex
}

func NewAccountService() *AccountService {
    return &AccountService{
        Accounts: make(map[string]*models.Account),
        Logins:   make(map[string]string),
    }
}

func (s *AccountService) CreateAccount(login string) (*models.Account, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    if _, exists := s.Logins[login]; exists {
        return nil, fmt.Errorf("login already exists")
    }

    id := uuid.New().String()
    account := &models.Account{ID: id}
    s.Accounts[id] = account
    s.Logins[login] = id

    log.Printf("Account created: %s (login: %s) at %v", id, login, time.Now())
    return account, nil
}

func (s *AccountService) Deposit(id string, amount float64) error {
    account, exists := s.Accounts[id]
    if !exists {
        return fmt.Errorf("account not found")
    }
    go func() {
        err := account.Deposit(amount)
        if err == nil {
            log.Printf("Deposited %f to account %s at %v", amount, id, time.Now())
        } else {
            log.Printf("Failed to deposit %f to account %s at %v: %v", amount, id, time.Now(), err)
        }
    }()
    return nil
}

func (s *AccountService) Withdraw(id string, amount float64) error {
    account, exists := s.Accounts[id]
    if !exists {
        return fmt.Errorf("account not found")
    }
    go func() {
        err := account.Withdraw(amount)
        if err == nil {
            log.Printf("Withdrew %f from account %s at %v", amount, id, time.Now())
        } else {
            log.Printf("Failed to withdraw %f from account %s at %v: %v", amount, id, time.Now(), err)
        }
    }()
    return nil
}

func (s *AccountService) GetBalance(id string) (float64, error) {
    account, exists := s.Accounts[id]
    if !exists {
        return 0, fmt.Errorf("account not found")
    }
    result := make(chan float64)
    go func() {
        result <- account.GetBalance()
    }()
    balance := <-result
    log.Printf("Checked balance for account %s at %v: %f", id, time.Now(), balance)
    return balance, nil
}
