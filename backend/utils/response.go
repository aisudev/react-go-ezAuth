package utils

import "fmt"

func Response(success bool, msg *string, data interface{}, err error) map[string]interface{} {

	resMap := map[string]interface{}{
		"success": success,
	}

	if msg != nil {
		resMap["msg"] = msg
	}

	if data != nil {
		resMap["data"] = data
	}

	if err != nil {
		resMap["error"] = err.Error()
	}

	return resMap

}

func Log(msg string, err error) {
	fmt.Println("------------------------------")
	fmt.Println("msg ======>", msg)
	fmt.Println("error ====>", err.Error())
}
