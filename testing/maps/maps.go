package maps

func Search(dic map[string]string, s string) string {
	return dic[s]
}

//never initialize empty maps like var m map[string]string
//do it this wat var dictionaty = map[string]string{}
//thats a correct empty map
