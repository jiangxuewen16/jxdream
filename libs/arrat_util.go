package libs

func StringArrayHasElement(arr []string, elem string) bool {
	for _,arrElem := range arr {
		if arrElem == elem {
			return true
		}
	}
	return false
}