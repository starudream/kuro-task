package main

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/utils/fmtutil"
	"github.com/starudream/go-lib/core/v2/utils/sliceutil"
	"github.com/starudream/go-lib/tablew/v2"

	"github.com/starudream/kuro-task/api/kuro"
	"github.com/starudream/kuro-task/config"
)

var (
	accountCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "account"
		c.Short = "Manage accounts"
	})

	accountLoginCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "add <account phone>"
		c.Short = "Add account"
		c.RunE = func(cmd *cobra.Command, args []string) error {
			phone, _ := sliceutil.GetValue(args, 0)
			if phone == "" {
				return fmt.Errorf("requires account phone")
			}

			account := config.Account{Phone: phone}

			account.DevCode = fmtutil.Scan("please enter device code: ")
			account.Token = fmtutil.Scan("please enter token: ")

			user, err := kuro.GetUser(account)
			if err != nil {
				return fmt.Errorf("get user error: %w", err)
			}
			fmt.Println(tablew.Structs([]kuro.Mine{user.Mine}))

			config.AddAccount(account)
			return config.Save()
		}
	})

	accountListCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "list"
		c.Short = "List accounts"
		c.Run = func(cmd *cobra.Command, args []string) {
			fmt.Println(tablew.Structs(config.C().Accounts))
		}
	})
)

func init() {
	accountCmd.AddCommand(accountLoginCmd)
	accountCmd.AddCommand(accountListCmd)

	rootCmd.AddCommand(accountCmd)
}
