package system_tests

import (
	"fmt"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("mysql-restore", func() {
	var dbDumpPath string
	var configPath string
	var databaseName string
	var dbJob, brJob JobInstance

	BeforeEach(func() {
		brJob = JobInstance{
			deployment:    "mysql-dev",
			instance:      "database-backup-restorer",
			instanceIndex: "0",
		}

		dbJob = JobInstance{
			deployment:    "mysql-dev",
			instance:      "mysql",
			instanceIndex: "0",
		}

		configPath = "/tmp/config.json" + strconv.FormatInt(time.Now().Unix(), 10)
		dbDumpPath = "/tmp/sql_dump" + strconv.FormatInt(time.Now().Unix(), 10)
		databaseName = "db" + strconv.FormatInt(time.Now().Unix(), 10)

		dbJob.runOnVMAndSucceed(fmt.Sprintf(`echo 'CREATE DATABASE %s;' | /var/vcap/packages/mariadb/bin/mysql -u root -h localhost --password='%s'`, databaseName, MustHaveEnv("MYSQL_PASSWORD")))
		dbJob.runMysqlSqlCommand("CREATE TABLE people (name varchar(255));", databaseName)
		dbJob.runMysqlSqlCommand("INSERT INTO people VALUES ('Derik');", databaseName)

		ip := dbJob.getIPOfInstance()
		configJson := fmt.Sprintf(
			`{"username":"root","password":"%s","host":"%s","port":3306,"database":"%s","adapter":"mysql"}`,
			MustHaveEnv("MYSQL_PASSWORD"),
			ip,
			databaseName,
		)

		brJob.RunOnInstance(fmt.Sprintf("echo '%s' > %s", configJson, configPath))
	})

	AfterEach(func() {
		dbJob.runOnVMAndSucceed(fmt.Sprintf(`echo 'DROP DATABASE %s;' | /var/vcap/packages/mariadb/bin/mysql -u root -h localhost --password='%s'`, databaseName, MustHaveEnv("MYSQL_PASSWORD")))
		brJob.RunOnInstance(fmt.Sprintf("rm -rf %s %s", configPath, dbDumpPath))
	})

	Context("database-backup-restorer lives on its own instance", func() {
		It("restores the MySQL database", func() {
			backupSession := brJob.RunOnInstance(fmt.Sprintf("/var/vcap/jobs/database-backup-restorer/bin/backup --artifact-file %s --config %s", dbDumpPath, configPath))
			Expect(backupSession).To(gexec.Exit(0))

			dbJob.runMysqlSqlCommand("UPDATE people SET NAME = 'Dave';", databaseName)

			restoreSession := brJob.RunOnInstance(fmt.Sprintf("/var/vcap/jobs/database-backup-restorer/bin/restore --artifact-file %s --config %s", dbDumpPath, configPath))
			Expect(restoreSession).To(gexec.Exit(0))

			Expect(dbJob.runMysqlSqlCommand("SELECT name FROM people;", databaseName)).To(gbytes.Say("Derik"))
			Expect(dbJob.runMysqlSqlCommand("SELECT name FROM people;", databaseName)).NotTo(gbytes.Say("Dave"))
		})
	})
})