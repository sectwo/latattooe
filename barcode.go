package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
)

func rand2Word() string {
	rand.Seed(time.Now().UnixNano())

	// 첫번째 자리와 두번째 자리에 알파벳을 생성
	letter1 := byte(rand.Intn(26) + 65)
	letter2 := byte(rand.Intn(26) + 65)
	code := string([]byte{letter1, letter2})

	return code
}

func hashBox(id string, size int) string {
	word := rand2Word()
	// 문자열을 숫자로 변환
	// 문자열 정의
	slice := []byte(id)

	// SHA256 해시 값 생성
	hash := sha256.Sum256(slice)
	hashString := hex.EncodeToString(hash[:])

	// 문자열을 자르기
	length := len(hashString) / 2
	if size < length {
		leftHashString := hashString[:size]
		return word + leftHashString
	}

	return "error"
}

func GenerateBarcode(reqMsg reqMsgGenBarcode) respMsgGenBarcode {
	barcodeList := viewAllBarcodeInfo()
	//fmt.Println(barcodeList)
	barcode := hashBox(reqMsg.ID, reqMsg.Size)
	for key, value := range barcodeList {
		if key == reqMsg.ID {
			res := &Result{
				Code: "7000",
				Msg:  "Existied ID!! Can not generate Barcode",
			}
			result := &respMsgGenBarcode{
				Barcode: "",
				Result:  *res,
			}
			return *result
		}
		if value == barcode {
			res := &Result{
				Code: "7000",
				Msg:  "Existied Barcode!! Can not generate Barcode",
			}
			result := &respMsgGenBarcode{
				Barcode: "",
				Result:  *res,
			}
			return *result
		}
	}
	res := &Result{
		Code: "9000",
		Msg:  "Success",
	}
	result := &respMsgGenBarcode{
		Barcode: barcode,
		Result:  *res,
	}
	SaveBarcode(reqMsg.ID, barcode)

	return *result
}

func SearchBarcode(reqMsg reqMsgSearchBarcode) respMsgSearchBarcode {
	barcode := viewBarcode(reqMsg.ID)
	res := &Result{
		Code: "9000",
		Msg:  "Success",
	}
	result := &respMsgSearchBarcode{
		Barcode: barcode,
		Result:  *res,
	}

	fmt.Println("barcode : ", barcode)
	return *result
}

func GenerateBarcodeImg(id string) {
	barcode := viewBarcode(id)
	writer := oned.NewCode128Writer()
	fmt.Println(barcode)
	img, err := writer.Encode(barcode, gozxing.BarcodeFormat_CODE_128, 250, 50, nil)
	if err != nil {
		log.Fatalf("impossible to encode barcode: %s", err)
	}
	// create a file that will hold our barcode
	file, err := os.Create("barcode.png")
	if err != nil {
		log.Fatalf("impossible to create file: %s", err)
	}
	defer file.Close()
	// Encode the image in PNG
	err = png.Encode(file, img)
	if err != nil {
		log.Fatalf("impossible to encode barcode in PNG: %s", err)
	}

}
