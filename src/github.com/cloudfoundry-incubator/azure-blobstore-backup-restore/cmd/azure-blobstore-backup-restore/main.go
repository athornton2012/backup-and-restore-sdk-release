package main

import (
	"flag"
	"log"

	"github.com/cloudfoundry-incubator/azure-blobstore-backup-restore"
)

func main() {
	var artifactFilePath = flag.String("artifact-file", "", "Path to the artifact file")
	var configFilePath = flag.String("config", "", "Path to JSON config file")
	//var backup = flag.Bool("backup", false, "Run blobstore backup")
	//var restore = flag.Bool("restore", false, "Run blobstore restore")

	flag.Parse()

	config, err := azure.ParseConfig(*configFilePath)
	if err != nil {
		log.Fatalf(err.Error())
	}

	backuper := azure.NewBackuper(config)
	backups, err := backuper.Backup()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = azure.NewArtifact(*artifactFilePath).Write(backups)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
