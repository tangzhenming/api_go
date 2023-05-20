// $ go run . migrate

package cmd

// var migrateCmd = &cobra.Command{
// 	Use:   "migrate",
// 	Short: "Migrate the database",
// 	Long:  `Migrate the database by creating or updating tables.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		switch modelName {
// 		case "User":
// 			err = DB.AutoMigrate(&models.User{})
// 		default:
// 			log.Fatalf("Unknown model: %s", modelName)
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Printf("Successfully migrate %s \n", modelName)
// 	},
// }

// func init() {
// 	migrateCmd.Flags().StringVarP(&modelName, "migrate", "m", "", "The name of the model to migrate")
// }
