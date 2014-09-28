package httpUtil_test

func ExampleLogToFile() {
	f, err := httpUtil.LogToFile("./logs", "logger.log")
	if err != nil {
		log.Println(err.Error())
		return
	}
}
