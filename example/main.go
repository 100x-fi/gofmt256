package main

import (
	"fmt"
	"github.com/100x-fi/gofmt256"
)

type SubMerchantReportHeader struct {
	RecordType     string `gofmt256:"from=1,to=1"`
	SequenceNo     int    `gofmt256:"from=2,to=7,align=R,padding='0'"`
	BankCode       string `gofmt256:"from=8,to=10"`
	CompanyAccount string `gofmt256:"from=11,to=20"`
	CompanyName    string `gofmt256:"from=21,to=60"`
	EffectiveDate  string `gofmt256:"from=61,to=68"`
	ServiceCode    string `gofmt256:"from=69,to=76"`
	Spare          string `gofmt256:"from=77,to=256"`
}

type SubMerchantReportBody struct {
	RecordType     string `gofmt256:"from=1,to=1"`
	SequenceNo     int    `gofmt256:"from=2,to=7,align=R,padding='0'"`
	BankCode       string `gofmt256:"from=8,to=10"`
	CompanyAccount string `gofmt256:"from=11,to=20"`
	PaymentDate    string `gofmt256:"from=21,to=28"`
	PaymentTime    string `gofmt256:"from=29,to=34"`
	CustomerName   string `gofmt256:"from=35,to=84"`
	Ref1           string `gofmt256:"from=85,to=104"`
	Ref2           string `gofmt256:"from=105,to=124"`
	Ref3           string `gofmt256:"from=125,to=144"`
	BranchNo       string `gofmt256:"from=145,to=148"`
	TellerNo       string `gofmt256:"from=149,to=152"`
	KindOfTx       string `gofmt256:"from=153,to=153"`
	TxCode         string `gofmt256:"from=154,to=156"`
	ChequeNo       string `gofmt256:"from=157,to=163"`
	Amount         string `gofmt256:"from=164,to=176,align=R,padding='0'"`
	Spare          string `gofmt256:"from=177,to=256"`
}

type SubMerchantReportFooter struct {
	RecordType             string `gofmt256:"from=1,to=1"`
	SequenceNo             int    `gofmt256:"from=2,to=7,align=R,padding='0'"`
	BankCode               string `gofmt256:"from=8,to=10"`
	CompanyAccount         string `gofmt256:"from=11,to=20"`
	TotalDebitAmount       string `gofmt256:"from=21,to=33,align=R,padding='0'"`
	TotalDebitTransaction  int    `gofmt256:"from=34,to=39,align=R,padding='0'"`
	TotalCreditAmount      string `gofmt256:"from=40,to=52,align=R,padding='0'"`
	TotalCreditTransaction int    `gofmt256:"from=53,to=58,align=R,padding='0'"`
	Spare                  string `gofmt256:"from=59,to=256"`
}

func main() {
	header := SubMerchantReportHeader{
		RecordType:     "H",
		SequenceNo:     1,
		BankCode:       "888",
		CompanyAccount: "8888888888",
		CompanyName:    "100X-FI",
		EffectiveDate:  "03092020",
		ServiceCode:    "",
		Spare:          "",
	}

	body := []SubMerchantReportBody{{
		RecordType:     "D",
		SequenceNo:     2,
		BankCode:       "888",
		CompanyAccount: "8888888888",
		PaymentDate:    "03092020",
		PaymentTime:    "100337",
		CustomerName:   "John Doe",
		Ref1:           "8888888",
		Ref2:           "888888888888",
		Ref3:           "00000000000000000000",
		BranchNo:       "0000",
		TellerNo:       "0000",
		KindOfTx:       "C",
		TxCode:         "BTC",
		ChequeNo:       "0000000",
		Amount:         "51500",
		Spare:          "",
	}, {
		RecordType:     "D",
		SequenceNo:     3,
		BankCode:       "888",
		CompanyAccount: "8888888888",
		PaymentDate:    "03092020",
		PaymentTime:    "100739",
		CustomerName:   "John Doe",
		Ref1:           "9999999",
		Ref2:           "9999999999999",
		Ref3:           "00000000000000000000",
		BranchNo:       "0000",
		TellerNo:       "0000",
		KindOfTx:       "C",
		TxCode:         "BTC",
		ChequeNo:       "0000000",
		Amount:         "746000",
		Spare:          "",
	}, {
		RecordType:     "D",
		SequenceNo:     4,
		BankCode:       "888",
		CompanyAccount: "8888888888",
		PaymentDate:    "03092020",
		PaymentTime:    "101056",
		CustomerName:   "John Doe",
		Ref1:           "7777777",
		Ref2:           "7777777777777",
		Ref3:           "00000000000000000000",
		BranchNo:       "0000",
		TellerNo:       "0000",
		KindOfTx:       "C",
		TxCode:         "BTC",
		ChequeNo:       "0000000",
		Amount:         "4880700",
		Spare:          "",
	}}

	footer := SubMerchantReportFooter{
		RecordType:             "T",
		SequenceNo:             5,
		BankCode:               "888",
		CompanyAccount:         "8888888888",
		TotalDebitAmount:       "0000",
		TotalDebitTransaction:  0,
		TotalCreditAmount:      "5678200",
		TotalCreditTransaction: 3,
		Spare:                  "",
	}

	builder := gofmt256.New(header, body, footer)
	out, err := builder.Build()
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}
