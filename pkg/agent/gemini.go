package agent

import (
	"context"
	"cooder-assist/pkg/log"
	"cooder-assist/pkg/scanner"
	"cooder-assist/pkg/tools"
	"fmt"
	"strings"

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
func colorizeDiff(diff string) string {
	lines := strings.Split(diff, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
			lines[i] = "\033[32m" + line + "\033[0m" // Green
		} else if strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "---") {
			lines[i] = "\033[31m" + line + "\033[0m" // Red
		}
	}
	return strings.Join(lines, "\n")
}

/*
Run function manages the interaction loop between the user and the Gemini model. It takes a context and a Gemini chat object as input.

It orchestrates the conversation between the user and the Gemini model, handling user input, sending messages to Gemini, processing Gemini's responses (including tool calls), and displaying the results to the user.
*/
/*
Run function manages the interaction loop between the user and the Gemini model. It takes a context and a Gemini chat object as input.

It orchestrates the conversation between the user and the Gemini model, handling user input, sending messages to Gemini, processing Gemini's responses (including tool calls), and displaying the results to the user.
*/
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
				fmt.Println("Empty user input is not accepted")
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
					a.Logger.Info("Executing tool", "tool_name", part.FunctionCall.Name, "prompt", part.FunctionCall.Args)
					resp := a.Tools.ExecuteTool(part.FunctionCall)

					if output, ok := resp.Response["output"].(string); ok {
						fmt.Printf("\n\033[1;94m[Tool Output]\033[0m\n%s\n", colorizeDiff(output))
					} else if errMsg, ok := resp.Response["error"].(string); ok {
						fmt.Printf("\n\033[1;91m[Tool Error]\033[0m %s\n", errMsg)
					}
					conversation = append(conversation, genai.Part{FunctionResponse: resp})
					hasToolCalls = true
				}
			}

			readUserInput = !hasToolCalls
		}
	}
}
