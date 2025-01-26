package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var FFinalizado = false

func main() {
	go loadingPrint()
	backupDataBase()
	FFinalizado = true
	time.Sleep(200 * time.Millisecond)
}

func backupDataBase() {
	// docker exec -t 8a63432d4f7d pg_dump -U postgres postgres > "C:\BACKUP\BANCO\backup_postgres.sql"
	cmd := exec.Command(
		"docker",
		"exec",
		"-t",
		"8a63432d4f7d",
		"pg_dump",
		"-U",
		"postgres",
		"postgres",
	)
	fileName := fmt.Sprintf("C:\\BACKUP\\BANCO\\bkp_postgres_%s.sql", time.Now().Format("0102"))
	outFile, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer outFile.Close()
	cmd.Stdout = outFile

	if err := cmd.Run(); err != nil {
		return
	}

	fmt.Println("\nBackup concluído com sucesso!")

}

func loadingPrint() {
	msg := "Realizando Backup do Banco de Dados "
	loadPipis := []string{".", "..", "...", "o", "oo", "ooo"}
	for i := 0; ; i++ {
		fmt.Printf("\r%s%s", msg, loadPipis[i%len(loadPipis)])
		time.Sleep(100 * time.Millisecond)
		if FFinalizado {
			break
		}
	}
}
