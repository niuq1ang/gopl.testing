package main

func main() {

}

func remove(slice []int, i int) []int {
	copy(slice[:i], slice[i+1:])
	return slice[:len(slice)-1]
}
