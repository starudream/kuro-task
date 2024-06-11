package main

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/utils/fmtutil"
	"github.com/starudream/go-lib/core/v2/utils/sliceutil"
	"github.com/starudream/go-lib/tablew/v2"

	"github.com/starudream/kuro-task/api/kuro"
	"github.com/starudream/kuro-task/config"
	"github.com/starudream/kuro-task/util"
)

var (
	accountCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "account"
		c.Short = "Manage accounts"
	})

	accountAddCmd = cobra.NewCommand(func(c *cobra.Command) {
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

	accountLoginCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "login <account phone>"
		c.Short = "Login account"
		c.RunE = func(cmd *cobra.Command, args []string) error {
			phone, _ := sliceutil.GetValue(args, 0)
			if phone == "" {
				return fmt.Errorf("requires account phone")
			}

			geetest, err := json.UnmarshalTo[*kuro.GeeTestData](fmtutil.Scan("please enter GeeTest json string: "))
			if err != nil {
				return err
			}

			resp1, err := kuro.SendPhoneCodeGeeTest(phone, geetest)
			if err != nil {
				return err
			}

			if resp1.GeeTest {
				return fmt.Errorf("wrong GeeTest validation")
			}

			account := config.Account{
				Phone:   phone,
				DevCode: util.RandString(util.CharsetHex, 40),
				Source:  kuro.SourceAndroid,
			}

			resp2, err := kuro.LoginByPhoneCode(phone, fmtutil.Scan("please enter the verification code you received: "), account.DevCode)
			if err != nil {
				return err
			}

			account.Token = resp2.Token

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
	accountCmd.AddCommand(accountAddCmd)
	accountCmd.AddCommand(accountLoginCmd)
	accountCmd.AddCommand(accountListCmd)

	rootCmd.AddCommand(accountCmd)
}
