package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var FFinalizado = false

func main() {
	go loading_print()
	bakcup_database()
	FFinalizado = true
	time.Sleep(200 * time.Millisecond)
}

func bakcup_database() {
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
	outFile, err := os.Create("C:\\BACKUP\\BANCO\\backup_postgres.sql")
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

func loading_print() {
	msg := "Realizando Backup do Banco de Dados "
	load_pipis := []string{".", "..", "...", "o", "oo", "ooo"}
	for i := 0; ; i++ {
		fmt.Printf("\r%s%s", msg, load_pipis[i%len(load_pipis)])
		time.Sleep(100 * time.Millisecond)
		if FFinalizado {
			break
		}
	}
}
