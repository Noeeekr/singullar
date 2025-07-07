package cmd

import (
	"fmt"
	"strconv"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/operations"

	"github.com/spf13/cobra"
)

var seedCmd *cobra.Command = &cobra.Command{
	Use:   "seed [-e \"production\"|\"development\"] [-f env_file_path] [-t table] [-n amount] [arguments]",
	Short: "Seeds the table the specified amount of times. If the table needs some dependency it may seed it too, if it needs some parameters it will ask for.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tablename, _ := cmd.Flags().GetString("table")
		amount, _ := cmd.Flags().GetInt("amount")
		mode, _ := cmd.Flags().GetString("environment")
		path, _ := cmd.Flags().GetString("environmentFile")

		if res := environment.Parse(path); res != nil {
			fmt.Println(res.ParseToString())
			return
		}

		db, res := connections.ConnectWithEnvironment(connections.ConnectionEnvironment(mode))
		if res != nil {
			fmt.Println(res.ParseToString())
			return
		}

		seeder := Seeder{
			ops: operations.New(db),
		}

		switch models.TableName(tablename) {
		case models.TablesInfo.Users.Name():
			if len(args) == 0 {
				fmt.Println("Please specify the id of the institution to insert users into")
				return
			}

			institution_id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			res := seeder.Users(institution_id, amount)
			if res != nil {
				fmt.Println(res.ParseToString())
				return
			}
		}
	},
}

type Seeder struct {
	ops *operations.Operations
}

func (s *Seeder) Users(institution_id, amount int) *common.Response {
	requests := make([]*models.CreateUsers, amount)
	for i := range amount {
		requests[i] = &models.CreateUsers{
			Name:          common.GenerateRandomStrings(10),
			Email:         common.GenerateRandomStrings(10),
			Password:      common.GenerateRandomStrings(10),
			InstitutionId: institution_id,
			Role:          models.Student,
		}
	}
	_, tx := s.ops.InsertManyUsers(nil, requests...)
	return tx.Response
}

func init() {
	seedCmd.Flags().IntP("amount", "n", 50, "The amount of times to seed the table. Defaults to 50")
	seedCmd.MarkFlagRequired("amount")

	seedCmd.Flags().StringP("table", "t", "", "The table to be seeded")
	seedCmd.MarkFlagRequired("table")

	rootCmd.AddCommand(seedCmd)
}
