package main

func GetCollectionWithoutElements[T comparable](collection []T, excludes ...T) []T {
	r := []T{}
	for _, v := range collection {
		for _, ex := range excludes {
			if ex == v {
				continue
			}
			r = append(r, v)
		}
	}
	return r
}
