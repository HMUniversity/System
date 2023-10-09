package err_handler

import "log"

func HandleError(err error) {
	log.Println("Receive err: " + err.Error())
}
