package command

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"notsybil/api"
	"notsybil/config"
	"notsybil/csvreader"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	fmt.Println("#### Made With ʕ•́ᴥ•̀ʔっ♡")
	fmt.Println("#### by Alucard")
	fmt.Println()
}

type Commands struct {
}

func (cmds *Commands) Run(ctx context.Context) error {
	rootCmd := &cobra.Command{
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
	setupCmd, err := cmds.setup(ctx)
	if err != nil {
		return err
	}
	rootCmd.AddCommand(setupCmd)

	withdrawCmd, err := cmds.withdraw(ctx)
	if err != nil {
		return err
	}
	rootCmd.AddCommand(withdrawCmd)

	return rootCmd.ExecuteContext(ctx)
}

func (cmds *Commands) setup(ctx context.Context) (*cobra.Command, error) {
	cmd := cobra.Command{
		Use:     "setup",
		Short:   "Setup credentials",
		Example: "notsybil setup",
		RunE: func(cmd *cobra.Command, args []string) error {
			var conf config.Config
			cmd.Print("Enter API Key: ")
			fmt.Scanln(&conf.APIKey)
			cmd.Print("Enter Secret Key: ")
			fmt.Scanln(&conf.SecretKey)
			cmd.Print("Enter Passphrase: ")
			fmt.Scanln(&conf.Passphrase)

			configJSON, err := json.MarshalIndent(conf, "", "  ")
			if err != nil {
				cmd.Println("Error encoding JSON:", err)
				return err
			}

			if err = ioutil.WriteFile("config.json", configJSON, 0644); err != nil {
				cmd.Println("Error writing to file:", err)
				return err
			}
			cmd.Println("File config.json updated successfully.")

			return nil
		},
	}
	return &cmd, nil
}

func (cmds *Commands) withdraw(ctx context.Context) (*cobra.Command, error) {
	cmd := cobra.Command{
		Use:     "withdraw",
		Short:   "Withdraw funds",
		Example: "notsybil withdraw",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := config.Load(".", "config", "json")
			if err != nil {
				return err
			}
			c := api.NewClient(conf.APIKey, conf.SecretKey, conf.Passphrase)

			assets, err := csvreader.Parse("withdraw.csv")
			if err != nil {
				return err
			}

			bal, err := c.AssetsBalance(assets)
			if err != nil {
				return err
			}

			cmd.Println("Available Funds:")
			avaccy := make(map[string]struct{})
			for _, b := range bal {
				avaccy[b.Currency] = struct{}{}
				cmd.Printf("%s: %s\n",
					color.CyanString(b.Currency),
					color.GreenString(b.Available),
				)
			}
			cmd.Println()

			wds, err := c.ComposeWithdraw(assets)
			if err != nil {
				return err
			}

			cmd.Println("Withdraw:")
			for _, wd := range wds {
				a := wd.Asset
				if a == nil {
					return err
				}
				if _, ok := avaccy[wd.Asset.Ccy]; !ok {
					continue
				}
				cmd.Printf("%s %s %s(fee) => [%s]\n",
					color.CyanString(a.Chain),
					color.GreenString(a.Amt),
					color.BlueString(wd.Fee),
					a.ToAddr,
				)

				var confirm string
				cmd.Print("confirm? (y/n):")
				fmt.Scanln(&confirm)
				if confirm != "y" {
					continue
				}
				wdid, err := c.Withdraw(wd)
				if err != nil {
					cmd.Printf("withdraw: %s", err)
				}
				cmd.Printf("success wdid=%s\n", wdid)
			}
			return nil
		},
	}
	return &cmd, nil
}
