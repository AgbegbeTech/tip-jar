package contracts

import (
	"fmt"
)

type TipJarContract struct {
	creatorPublicKey string
	duration         int
	totalTips        uint64
	contributions    map[string]uint64
}

func NewTipJarContract(creatorPublicKey string, duration int, totalTips uint64) *TipJarContract {
	return &TipJarContract{
		creatorPublicKey: creatorPublicKey,
		duration:         duration,
		totalTips:        totalTips,
		contributions:    make(map[string]uint64),
	}
}

func (tc *TipJarContract) Deposit(userPublicKey string, amount uint64) {
	tc.contributions[userPublicKey] += amount
}

func (tc *TipJarContract) GetTotalTips() uint64 {
	return tc.totalTips
}

func (tc *TipJarContract) DistributeTips(creatorPublicKey string) {
	if tc.creatorPublicKey != creatorPublicKey {
		fmt.Println("Only the creator can distribute tips.")
		return
	}

	totalTips := tc.totalTips

	if totalTips == 0 {
		fmt.Println("No tips to distribute.")
		return
	}

	fmt.Printf("Tips distributed to creator (%s): %d DeSo\n", creatorPublicKey, totalTips)

	tc.totalTips = 0
}

func main() {
	tipJar := NewTipJarContract("content_creator_public_key", 30, 0)
	tipJar.Deposit("alice_public_key", 50)
	tipJar.Deposit("bob_public_key", 30)
	tipJar.Deposit("carol_public_key", 20)
	tipJar.DistributeTips("content_creator_public_key")
}
