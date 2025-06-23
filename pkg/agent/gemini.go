package agent

import (
	"context"
	"cooder-assist/pkg/log"
	"cooder-assist/pkg/scanner"
	"cooder-assist/pkg/tools"
	"fmt"

	"google.golang.org/genai"
)

type Agent struct {
	Client  *genai.Client
	Model   string
	Logger  log.Logger
	Scanner scanner.Scanner
	Tools   tools.Tools
}

type AgentState int

const (
	StateUndefined AgentState = iota
	UserInput
	RunInference
	Process
	Error
)

func New(model string, l log.Logger, client *genai.Client, scanner scanner.Scanner, tools tools.Tools) *Agent {

	return &Agent{
		Client:  client,
		Model:   model,
		Logger:  l,
		Scanner: scanner,
		Tools:   tools,
	}
}

func (a *Agent) Run(ctx context.Context, chat *genai.Chat) error {
	var conversation []genai.Part
	readUserInput := true
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goodbye!")
			return nil
		default:
		}
		if readUserInput {
			fmt.Print("\u001b[94mYou\u001b[0m:")
			userInput, ok := a.Scanner.GetUserMessage()
			if !ok {
				fmt.Println("Error getting user input")
				continue
			}
			conversation = append(conversation, genai.Part{Text: userInput})
			readUserInput = false // Only set to false after successful input
		}

		if len(conversation) != 0 {
			response, err := chat.SendMessage(ctx, conversation...)
			if err != nil {
				fmt.Printf("Error sending message: %v. Try again\n", err)
				readUserInput = true // Reset to allow user input again
				continue
			}

			if len(response.Candidates) == 0 {
				fmt.Println("No response from Gemini")
				readUserInput = true
				continue
			}

			content := response.Candidates[0].Content
			numParts := len(content.Parts)
			if numParts == 0 {
				fmt.Println("I need more context")
				readUserInput = true
				continue
			}

			conversation = nil
			hasToolCalls := false

			for _, part := range content.Parts {
				if len(part.Text) != 0 {
					fmt.Printf("\u001b[93mGemini\u001b[0m: %s\n", part.Text)
				} else if part.FunctionCall != nil {
					conversation = append(conversation, genai.Part{
						FunctionResponse: a.Tools.ExecuteTool(part.FunctionCall),
					})
					hasToolCalls = true
				}
			}

			readUserInput = !hasToolCalls
		}
	}
}
