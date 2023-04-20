package main

func BarcodeRouter(barcode string) respMsgGetMypage {
	url := "www.naver.com"
	res := &Result{
		Code: "9000",
		Msg:  "Success",
	}
	result := &respMsgGetMypage{
		URL:    url,
		Result: *res,
	}

	return *result
}
