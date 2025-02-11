package x86_gen

func handleDB(args []interface{}) []byte {
	var binary []byte
	for _, arg := range args {
		binary = append(binary, byte(arg.(int)))
	}
	return binary
}

func handleDW(args []interface{}) []byte {
	var binary []byte
	for _, arg := range args {
		val := arg.(int)
		binary = append(binary, byte(val&0xFF), byte((val>>8)&0xFF))
	}
	return binary
}

func handleDD(args []interface{}) []byte {
	var binary []byte
	for _, arg := range args {
		val := arg.(int)
		binary = append(binary, byte(val&0xFF), byte((val>>8)&0xFF), byte((val>>16)&0xFF), byte((val>>24)&0xFF))
	}
	return binary
}
