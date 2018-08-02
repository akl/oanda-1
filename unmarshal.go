package oanda

import (
	"encoding/json"
	"time"
)

/*
***************************
prices
***************************
*/

//Heartbeat is returned from the Oanda streaming endpoint
type Heartbeat struct {
	Time time.Time `json:"time"`
	Type string    `json:"type"`
}

//UnmarshalHeartbeat is a method of Heartbeat
func (h Heartbeat) UnmarshalHeartbeat(heartbeatByte []byte) *Heartbeat {

	err := json.Unmarshal(heartbeatByte, &h)

	if err != nil {
		panic(err)
	}

	return &h
}

//Pricing is returned from the Oanda pricing endpoint
type Pricing struct {
	Prices []Prices  `json:"prices"`
	Time   time.Time `json:"time"`
}

//Prices is embedded within each Pricing struct and  is returned object from
//the Oanda streaming endpoint
type Prices struct {
	Type        string    `json:"type"`
	Bids        []Bid     `json:"bids"`
	Asks        []Ask     `jons:"asks"`
	CloseOutAsk string    `json:"closeoutAsk"`
	CloseOutBid string    `json:"closeoutBid"`
	Instrument  string    `json:"instrument"`
	Status      string    `json:"status"`
	Time        time.Time `json:"time"`
}

//Ask represents one element in the Asks list of a Prices Struct
type Ask struct {
	Price     string `json:"price"`
	Liquidity int64  `json:"liquidity"`
}

//Bid represents one element in the Bids list of a Prices Struct
type Bid struct {
	Price     string `json:"price"`
	Liquidity int64  `json:"liquidity"`
}

//UnmarshalPrices used by StreamPricing
func (p Prices) UnmarshalPrices(priceByte []byte) *Prices {

	err := json.Unmarshal(priceByte, &p)

	if err != nil {
		panic(err)
	}

	return &p
}

//UnmarshalPricing unmarshals the Pricing data byte slice from Oanda
func (p Pricing) UnmarshalPricing(priceByte []byte) *Pricing {

	err := json.Unmarshal(priceByte, &p)

	if err != nil {
		panic(err)
	}

	return &p
}

/*
***************************
history
***************************
*/

//Candles represents the data structure returned by Oanda when requesting
//multiple Candles
type Candles struct {
	Instrument  string   `json:"instrument"`
	Granularity string   `json:"granularity"`
	Candles     []Candle `json:"candles"`
}

//Candle represents a single data point in an instrument's pricing history
type Candle struct {
	Complete bool      `json:"complete"`
	Volume   int64     `json:"volume"`
	Time     time.Time `json:"time"`
	Mid      Mid       `json:"mid"`
}

//Mid represents the actual quotes/prices in a Candle
type Mid struct {
	Open  string `json:"o"`
	High  string `json:"h"`
	Low   string `json:"l"`
	Close string `json:"c"`
}

//UnmarshalCandles unmarshals History data byte slice from Oanda
func (c Candles) UnmarshalCandles(priceByte []byte) *Candles {

	err := json.Unmarshal(priceByte, &c)

	if err != nil {
		panic(err)
	}

	return &c
}

/*
***************************
orders
***************************
*/

//OrderCreateTransaction represents the data structure returned by oanda after
//submiting an order
type OrderCreateTransaction struct {
	OrderCreateTransaction OrderCreateTransactionData `json:"orderCreateTransaction"`
	OrderFillTransaction   OrderFillTransactionData   `json:"orderFillTransaction"`
}

//OrderCreateTransactionData represents the data structure embedded in
//OrderCreateTransaction
type OrderCreateTransactionData struct {
	Type             string           `json:"type"`
	Instrument       string           `json:"instrument"`
	Units            string           `json:"units"`
	TimeInForce      string           `json:"timeInForce"`
	PositionFill     string           `json:"positionFill"`
	TakeProfitOnFill TakeProfitOnFill `json:"takeProfitOnFill"` //see orders.go
	StopLossOnFill   StopLossOnFill   `json:"stopLossOnFill"`   //see orders.go
	Reason           string           `json:"reason"`
	ID               string           `json:"id"`
	UserID           int              `json:"userID"`
	AccountID        string           `json:"accountID"`
	BatchID          string           `json:"batchID"`
	RequestID        string           `json:"requestID"`
	Time             time.Time        `json:"time"`
}

//OrderFillTransactionData represents the data structure embedded in
//OrderCreateTransaction
type OrderFillTransactionData struct {
	Type                          string          `json:"type"`
	OrderID                       string          `json:"orderID"`
	Instrument                    string          `json:"instrument"`
	Units                         string          `json:"units"`
	Price                         string          `json:"price"`
	PL                            string          `json:"pl"`
	Financing                     string          `json:"financing"`
	Commission                    string          `json:"commission"`
	AccountBalance                string          `json:"accountBalance"`
	GainQuoteHomeConversionFactor string          `json:"gainQuoteHomeConversionFactor"`
	LossQuoteHomeConversionFactor string          `json:"lossQuoteHomeConversionFactor"`
	HalfSpreadCost                string          `json:"halfSpreadCost"`
	Reason                        string          `json:"reason"`
	TradeOpened                   TradeOpenedData `json:"tradeOpened"`
	FullPrice                     FullPrice       `json:"fullPrice"`
	RelatedTransactionIDs         []string        `json:"relatedTransactionIDs"`
	LastTransactionID             string          `json:"lastTransactionID"`
}

//TradeOpenedData represents the data structure embedded in OrderFillTransactionData
type TradeOpenedData struct {
	Price                  string `json:"price"`
	TradeID                string `json:"tradeID"`
	Units                  string `json:"units"`
	GuaranteedExecutionFee string `json:"guaranteedExecutionFee"`
	HalfSpreadCost         string `json:"halfSpreadCost"`
	InitialMarginRequired  string `json:"initialMarginRequired"`
	LastTransactionID      string `json:"lastTransactionID"`
}

//FullPrice represents the data structure embedded in OrderFillTransactionData
type FullPrice struct {
	CloseoutBid string         `json:"closeoutBid"`
	CloseoutAsk string         `json:"closeoutAsk"`
	Time        time.Time      `json:"timestamp"`
	Bids        []FullPriceBid `json:"bids"`
	Asks        []FullPriceAsk `json:"asks"`
	ID          string         `json:"id"`
	UserID      string         `json:"userID"`
	AccountID   string         `json:"accountID"`
	BatchID     string         `json:"batchID"`
}

//FullPriceBid represents one element in the Bids list of a Prices Struct
//this differs from Bid which has an int for Liquidity
type FullPriceBid struct {
	Price     string `json:"price"`
	Liquidity string `json:"liquidity"`
}

//FullPriceAsk represents one element in the Asks list of a Prices Struct
//this differs from Ask which has an int for Liquidity
type FullPriceAsk struct {
	Price     string `json:"price"`
	Liquidity string `json:"liquidity"`
}

//UnmarshalOrderCreateTransaction unmarshals the returned data byte slice from Oanda
//that contains the order data
func (o OrderCreateTransaction) UnmarshalOrderCreateTransaction(
	ordersResponseByte []byte) *OrderCreateTransaction {

	err := json.Unmarshal(ordersResponseByte, &o)

	if err != nil {
		panic(err)
	}

	return &o
}
