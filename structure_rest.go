package main

type Result struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type reqMsgGenBarcode struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

type respMsgGenBarcode struct {
	Barcode string `json:"barcode"`
	Result  Result `json:"result"`
}

type reqMsgSearchBarcode struct {
	ID string `json:"id"`
}

type respMsgSearchBarcode struct {
	Barcode string `json:"barcode"`
	Result  Result `json:"result"`
}

type respMsgAllBarcodeList struct {
	ID      string `json:"id"`
	Barcode string `json:"barcode"`
	Result  Result `json:"result"`
}

type reqMsgGetMypage struct {
	Barcode string `json:"Barcode"`
}

type respMsgGetMypage struct {
	URL    string `json:"url"`
	Result Result `json:"result"`
}
