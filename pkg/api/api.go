package api

import (
	"fmt"

	constants "github.com/shaksham08/temo_go_mod/constants"
)

func CallApi() {
	fmt.Println("Calling API.....")
	fmt.Println("LOG LEVEL is " + constants.LOG_LEVEL)
}
