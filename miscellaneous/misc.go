package main

import (
	"reflect"
	"runtime"
)

func detect_os()(bool){
	if runtime.GOOS == "windows" {
		return false
	} else {
		return true
	}
	return false
}

func swap(x, y int)(int, int){
	x, y = y, x
	return x, y
}

func is_even(number int) bool {
	if (number % 2) == 0{
		return true
	} else {
		return false
	}
}

func how_many_69(number int) int {
	count := 0
	for true{
		if number >= 69 {
			number = number - 69
			count++
		} else {
			return count
		}
	}
	return 0
}

func how_many_420(number int) int {
	count := 0
	for true{
		if number >= 420 {
			number = number - 420
			count++
		} else {
			return count
		}
	}
	return 0
}

func how_many_x_in_y(x, y int) int {
	count := 0
	for true{
		if x < y{
			return count
		} else {
			x = x - y
			count++
		}
	}
	return 0
}

func second_over_the_first(first_value, second_value, return_value string) string {
	/*
	case 1 : first = "";  second = "";  -> returns return value
	case 2 : first = "x"; second = "";  -> returns first
	case 3 : first = "";  second = "y"; -> returns second
	case 4 : first = "x"; second = "y"; -> returns second
	 */
	if first_value == "" {
		if second_value == "" {
			return return_value
		} else {
			return second_value
		}
	} else {
		if (second_value == "") {
			return first_value;
		}else{
				return second_value;
			}
	}
}

func go_version()(string){
	return runtime.Version()
}

func join_list(word []string, sep string)(string){
	/*
		:param word : list of word that will be transformed to a single string
		:param sep  : separator that will be place in between each word
		:return 	: stringed word
	 */
	var return_string string
	for i := 0; i < len(word); i++{
		if i == 0{
			return_string =  word[i]
		}else{
			return_string = return_string + sep	 + word[i]
		}
	}
	return return_string
}

func clear_duplicates_in_list(src interface{}) interface{} {
	/*
	example :
				var intSlice []interface{}
				var = []interface{}{1, 2, 1,4, "hello", 5, 7, 8, "9"}
				fmt.Println(intSlice)
				uniqueSlice := clear_duplicates_in_list(intSlice).([]interface{})
				fmt.Println(uniqueSlice)

	 */
	srcv := reflect.ValueOf(src)
	dstv := reflect.MakeSlice(srcv.Type(), 0, 0)
	visited := make(map[interface{}]struct{})
	for i := 0; i < srcv.Len(); i++ {
		elemv := srcv.Index(i)
		if _, ok := visited[elemv.Interface()]; ok {
			continue
		}
		visited[elemv.Interface()] = struct{}{}
		dstv = reflect.Append(dstv, elemv)
	}
	return dstv.Interface()
}

func remove_to_index(s []string, index int) []string {
	/*
	x = [0 1 2 3 4 5 6 7 8 9]
	x = remove_to_index(x, 4)
	x = [4 5 6 7 8 9]

	 */
	return append(s[index:len(s)], s[len(s):]...)
}

func main()(){

}
























