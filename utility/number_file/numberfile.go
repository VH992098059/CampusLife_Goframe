package number_file

import (
	"demo3/internal/consts"
	"fmt"
	"os"
)

// 从文件中读取编号
func readNumberFromFile(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var number int
	_, err = fmt.Fscanf(file, "%d", &number)
	if err != nil {
		return 0, err
	}
	return number, nil
}

// 将编号写入文件
func writeNumberToFile(filePath string, number int) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%d", number)
	return err
}

func FileCor() {
	filePath := consts.NumberFile

	// 读取当前编号
	currentNumber, err := readNumberFromFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error reading number:", err)
		return
	}
	fmt.Println("Current number:", currentNumber)

	// 增加编号并写入文件
	newNumber := currentNumber + 1
	err = writeNumberToFile(filePath, newNumber)
	if err != nil {
		fmt.Println("Error writing number:", err)
		return
	}
	fmt.Println("New number:", newNumber)
}
