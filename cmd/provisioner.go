package main

import (
	"context"
	"cooder-assist/pkg/agent"
	"cooder-assist/pkg/config"
	"cooder-assist/pkg/log"
	"cooder-assist/pkg/scanner"
	"cooder-assist/pkg/tools"

	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"google.golang.org/genai"
)

var (
	cfgPath string
	cfgFile string
	rootCmd = &cobra.Command{
		Use:     "codingAssist [flags] [command]",
		Short:   "coding assist client",
		Example: `codingAssist --cfgPath=../`,
		Long:    ``,
		Run: func(cmd *cobra.Command, args []string) {
			newProvisioner()
		},
	}
)

func newProvisioner() {

	errExit := false
	defer func() {
		if errExit {
			os.Exit(1)
			return
		}
		os.Exit(0)
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	cfg, err := config.InitConfig(cfgFile, cfgPath)
	if err != nil {
		fmt.Println("Invalid configuration")
		errExit = true
		return
	}
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Backend: genai.BackendGeminiAPI,
	})
	logger := log.Init("provisioner", "./logdump.log")
	scanner := scanner.New()
	tools := tools.New()
	agent := agent.New(cfg.ModelConfig.Model, logger, client, scanner, tools)

	systemInstr := &genai.Content{
		Parts: []*genai.Part{
			{Text: "Answer concisely. Ask clarifying questions, if necessary."},
		},
		Role: "user",
	}
	config := &genai.GenerateContentConfig{
		SystemInstruction: systemInstr,
		CandidateCount:    1,
		Tools:             tools,
	}
	chat, err := client.Chats.Create(ctx, cfg.ModelConfig.Model, config, nil)

	agent.Run(ctx, chat)

}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgPath, "cfgPath", "", "config file path (default is '.')")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "cfgFile", "", "config file name (default is 'agent-config-default.yml')")
	rootCmd.AddCommand(versionCmd)
}

func ExecuteRootCommand() {
	fmt.Println("Cooder Assist")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing root command")
		os.Exit(1)
	}
}
