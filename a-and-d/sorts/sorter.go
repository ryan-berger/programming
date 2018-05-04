package sorts


func swap(dataset []int, i, j int)  {
	temp := dataset[i]
	dataset[i] = dataset[j]
	dataset[j] = temp
}

func InsertionSort(dataset []int) {
	for i := 1; i < len(dataset); i++ {
		for j := i; j > 0; j-- {
			if dataset[j] < dataset[j - 1] {
				swap(dataset, j, j - 1)
			}
		}
	}
}

func MergeSort(dataset []int, l, r int)  {
	
}

func split(dataset []int) (first, second []int) {
	return first, second
}

func sort()  {
	
}

func merge(l, r []int)  {
	
}


