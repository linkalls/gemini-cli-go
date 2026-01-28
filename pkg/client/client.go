// Copyright 2026 Google LLC
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/linkalls/gemini-cli-go/pkg/config"
	"google.golang.org/api/option"
)

// Client is a wrapper around the Gemini API client
type Client struct {
	genaiClient *genai.Client
	model       *genai.GenerativeModel
	ctx         context.Context
	history     []*genai.Content
}

// NewClient creates a new Gemini client
func NewClient(cfg *config.Config) (*Client, error) {
	ctx := context.Background()

	var genaiClient *genai.Client
	var err error

	if cfg.UseVertex {
		// Vertex AI client (would require different setup)
		return nil, fmt.Errorf("Vertex AI support not yet implemented in Go version")
	} else {
		// Use API key
		genaiClient, err = genai.NewClient(ctx, option.WithAPIKey(cfg.APIKey))
		if err != nil {
			return nil, fmt.Errorf("failed to create genai client: %w", err)
		}
	}

	model := genaiClient.GenerativeModel(cfg.Model)
	
	// Configure safety settings to be permissive (matching TypeScript version behavior)
	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockNone,
		},
	}

	return &Client{
		genaiClient: genaiClient,
		model:       model,
		ctx:         ctx,
		history:     make([]*genai.Content, 0),
	}, nil
}

// GenerateContent sends a prompt to Gemini and returns the response
func (c *Client) GenerateContent(prompt string) (string, error) {
	resp, err := c.model.GenerateContent(c.ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("GenerateContent failed: %w", err)
	}

	if len(resp.Candidates) == 0 {
		return "", fmt.Errorf("no response candidates returned")
	}

	var result strings.Builder
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			result.WriteString(string(txt))
		}
	}

	return result.String(), nil
}

// StartChat starts a new chat session with history
func (c *Client) StartChat() *ChatSession {
	return &ChatSession{
		chat: c.model.StartChat(),
		ctx:  c.ctx,
	}
}

// ChatSession wraps a genai.ChatSession with a context
type ChatSession struct {
	chat *genai.ChatSession
	ctx  context.Context
}

// SendMessage sends a message in the chat session
func (cs *ChatSession) SendMessage(parts ...genai.Part) (*genai.GenerateContentResponse, error) {
	return cs.chat.SendMessage(cs.ctx, parts...)
}

// Close closes the client
func (c *Client) Close() error {
	if c.genaiClient != nil {
		return c.genaiClient.Close()
	}
	return nil
}
