package array

import "fmt"

// массивы
func arrayAndSlice() {
	msgArr := [3]string{"a", "r", "r"}                     // У массива нельзя менять длину
	msgSlice := []string{"s", "l", "i", "c", "e"}          // У слайса можно менять длину и передается по ссылке
	msgSliceByMake := make([]string, 5)                    // Создание слайса через make
	msgSliceByMakeAppended := append(msgSliceByMake, "22") // Если не хватает длины при добавлении, то длина увеличивается на x2

	var changeSlice = func(msgSlice []string) []string {
		msgSlice[0] = "22"
		return msgSlice
	}

	fmt.Println(msgArr, " | msgArr")
	fmt.Println(msgSlice, " | msgSlice")
	fmt.Println(changeSlice(msgSlice), " | changeSlice")

	fmt.Println(msgSliceByMake, " | msgSliceByMake", len(msgSliceByMake), cap(msgSliceByMake))
	fmt.Println(msgSliceByMakeAppended, " | msgSliceByMakeAppended", len(msgSliceByMakeAppended), cap(msgSliceByMakeAppended))

	// matrix - матрица
	fmt.Println("matrix\n=======================\n'")
	matrix := make([][]int, 10)

	var counter int = 0
	for i := 0; i < 10; i++ {
		matrix[i] = make([]int, 10)

		for j := 0; j < 10; j++ {
			counter++
			matrix[i][j] = counter
		}

		fmt.Println(matrix[i])

	}

	// range - цикл
	fmt.Println("range\n=======================\n'")
	for index, value := range matrix {
		if index == 6 {
			fmt.Println("index равен: ", 6, " остановка цикла")
			break
		}
		fmt.Println(index, value)
	}

}
