package cmd

import (
	"fmt"
	"log"
	

	"github.com/spf13/cobra"
)

// timezoneCmd represents the timezone command
var arcCmd = &cobra.Command{
	Use:   "archiver",
	Short: "Create zip",
	Long: `Get the current time in a given timezone.
				  This command takes one argument, the timezone you want to get the current time in.
				  It returns the current time in RFC1123 format.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
         archiver := args[0] // это для создания файла
		 
		  arcCreateFlag, _ := cmd.Flags().GetString("archrCreate")
		   var archrCreate string
		   if arcCreateFlag !=""{
			archrCreate = arcCreateFlag
		   }
           
		   
		 currentZip,err := osCreate(archiver,archrCreate)
		 if err != nil {
			log.Fatalln("Что-то не так")
		   }
		   

		fmt.Println(currentZip,archrCreate,archiver)

	}}


	func init() {
		rootCmd.AddCommand(arcCmd)
	
		// Here you will define your flags and configuration settings.
	
		// Cobra supports Persistent Flags which will work for this command
		// and all subcommands, e.g.:
		// timezoneCmd.PersistentFlags().String("foo", "", "A help for foo")
	
		// Cobra supports local flags which will only run when this command
		// is called directly, e.g.:
		// timezoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	}
	
	func init() {
		rootCmd.AddCommand(arcCmd)
		arcCmd.PersistentFlags().String("archrCreate", "", "create zip archive")
	  }