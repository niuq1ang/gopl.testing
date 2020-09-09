package main

import (
	"log"
	"os/exec"
)

func main() {
	replaceCmd := func(dbName, newDBName, path string) {
		cmd := exec.Command("sed", "-i", "s/USE `"+dbName+"`;/USE `"+newDBName+"`;/g", path)
		stdout, err := cmd.Output()
		if err != nil {
			log.Fatalf("replace db name error: %s%v", string(stdout), err)
		}
	}
	replaceCmd("project", "project_test", "./ones-data/project.sql")
}
