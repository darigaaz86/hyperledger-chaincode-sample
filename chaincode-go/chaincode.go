package main

import (
  "fmt"
  "encoding/json"
//   "log"
  "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
  contractapi.Contract
}

type Lottery struct {
	TransactionId  string `json:"TransactionId"`
	LotteryNo      string `json:"LotteryNo"`
}

func (s *SmartContract) CreateLottery(ctx contractapi.TransactionContextInterface, transactionId string, lotteryNo string) error {
	exists, err := s.LotteryExists(ctx, transactionId)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the lottery %s already exists", transactionId)
	}
  
	lottery := Lottery{
		TransactionId: transactionId,
		LotteryNo: lotteryNo,
	}
	lotteryJSON, err := json.Marshal(lottery)
	if err != nil {
		return err
	}
  
	return ctx.GetStub().PutState(transactionId, lotteryJSON)
}

func (s *SmartContract) ReadLottery(ctx contractapi.TransactionContextInterface, transactionId string) (*Lottery, error) {
	lotteryJSON, err := ctx.GetStub().GetState(transactionId)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if lotteryJSON == nil {
		return nil, fmt.Errorf("the lottery %s does not exist", transactionId)
	}
  
	var lottery Lottery
	err = json.Unmarshal(lotteryJSON, &lottery)
	if err != nil {
		return nil, err
	}
  
	return &lottery, nil
}

func (s *SmartContract) LotteryExists(ctx contractapi.TransactionContextInterface, transactionId string) (bool, error) {
	lotteryJSON, err := ctx.GetStub().GetState(transactionId)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
  
	return lotteryJSON != nil, nil
}