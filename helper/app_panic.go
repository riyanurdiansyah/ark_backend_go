package helper

func PanicIfError(err error) {
	if err != nil {
		println("ERROR ", err.Error())
		panic(err)
	}
}
