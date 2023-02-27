package account

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "getAccount",
	Short: "Get information about account",
	Long:  "Get status or information about account",
	RunE:  run,
}

func run(cmd *cobra.Command, argv []string) error {
	response, err := http.Get("http://localhost:8080/v1/account")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

	return nil
}
