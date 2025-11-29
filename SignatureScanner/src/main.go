package main

import (
	"SignatureScanner"
	"fmt"
	"log"
)

func main() {
	avScanner := scanner.NewSignatureScanner()
	err := avScanner.Load("signatures.txt")
	if err != nil {
		log.Fatalf("Ошибка загрузки базы: %v", err)
	}
	fmt.Println("База загружена.")
	targetFile := "malware.txt"
	detected := avScanner.Scan(targetFile)
	if len(detected) == 0 {
		fmt.Println("Файл чист.")
	} else {
		fmt.Printf("Найдено угроз: %d\n", len(detected))
		for _, pos := range detected {
			fmt.Printf(" -> Смещение: 0x%x (0x%X), Сигнатура: %s\n",
				pos.Offset, pos.Offset, pos.GetSignature())
		}
	}
}
