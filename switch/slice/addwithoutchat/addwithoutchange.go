package main 


func main(){


	//here adding a ellement midle  without changing the original slice 
	slc := []int{1,2,4,5}

	//here what we doing, frist we are taking the first two ellement from the original slice
	//then we are appending a new slice which contain the new ellement we want to add
	//finally we are appending the rest of the original slice after the new ellement
	slc = append(slc[:2], append([]int{3},slc[2:]...)...)
}