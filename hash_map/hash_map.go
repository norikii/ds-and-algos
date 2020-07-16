package hash_map


//func Insert(key string, value interface{}) [...]interface{} {
//	var a [...]interface{}
//
//	// compute the key's hashcode
//	hashCode := 0
//	for _, b := range []byte(key) {
//		hashCode += int(b)
//	}
//
//	index := hashCode % len(hashMap{})
//	a[index] = value
//
//	return a
//}

func HashFunction(input string) [10]string {
	var storage [10]string
	var total int32

	for _, char := range input{
		total += char
	}

	position := total % int32(len(storage)-1)
	storage[position] = input

	return storage
}
