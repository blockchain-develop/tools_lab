package solana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type Request struct {
	Query string `json:"query""`
}

type Current struct {
	Decimals int    `json:"decimals"`
	Name     string `json:"name"`
}

type Date struct {
	Date string `json:"date"`
}

type Receiver struct {
	Address string `json:"address"`
}

type Sender struct {
	Address string `json:"address"`
}

type Program struct {
	Name string `json:"name"`
}

type TimeStamp struct {
	Time string `json:"time"`
}

type Block struct {
	Hash      string    `json:"hash"`
	Height    int64     `json:"height"`
	TimeStamp TimeStamp `json:"timestamp"`
}

type Transfer struct {
	Amount       float64  `json:"amount"`
	Currency     Current  `json:"currency"`
	Date         Date     `json:"date"`
	TransferType string   `json:"transferType"`
	Success      bool     `json:"success"`
	Receiver     Receiver `json:"receiver"`
	Sender       Sender   `json:"sender"`
	FeePayer     string   `json:"feePayer"`
	Program      Program  `json:"program"`
	Block        Block    `json:"block"`
}

type TransferViaSendersAddress struct {
	Transfers []Transfer `json:"transfers"`
}

type Data struct {
	TransferViaSendersAddress TransferViaSendersAddress `json:"transfer_via_senders_address"`
}

type Response struct {
	Data Data `json:"data"`
}

func TestSolanaQuery(t *testing.T) {
	req := Request{
		Query: "query MyQuery {\n  transfer_via_senders_address: solana(network: solana) {\n    transfers(\n      signature: {is: \"3Gimy7iesyo9jfGYciqJiSUzddC8nhS9QbWXLw6QvK6FdB6FfThwn5a9KxWuimXkvP8DQ8an1HkwK8QNehicANe1\"}\n    ) {\n      amount\n      currency {\n        decimals\n        name\n      }\n      date {\n        date\n      }\n      transferType\n      success\n      receiver {\n        address\n      }\n      sender {\n        address\n      }\n      feePayer\n      program(programId: {is: \"11111111111111111111111111111111\"}) {\n        name\n      }\n      block {\n        hash\n        height\n        timestamp {\n          time\n        }\n      }\n    }\n  }\n}",
	}
	reqJson, _ := json.Marshal(req)
	httpRequest, err := http.NewRequest("POST", "https://graphql.bitquery.io", bytes.NewBuffer(reqJson))
	if err != nil {
		panic(err)
	}
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Add("X-API-KEY", "BQYvhnv04csZHaprIBZNwtpRiDIwEIW9")

	client := &http.Client{}
	rsp, err := client.Do(httpRequest)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	statuscode := rsp.StatusCode
	body, _ := ioutil.ReadAll(rsp.Body)
	fmt.Printf("%d %s\n", statuscode, string(body))
	var xx Response
	err = json.Unmarshal(body, &xx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("amount: %g, height: %d\n",
		xx.Data.TransferViaSendersAddress.Transfers[0].Amount,
		xx.Data.TransferViaSendersAddress.Transfers[0].Block.Height,
	)
}
