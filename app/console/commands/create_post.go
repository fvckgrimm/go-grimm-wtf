package commands

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "time"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
)

type CreatePost struct {
}

//Signature The name and signature of the console command.
func (receiver *CreatePost) Signature() string {
	return "create:post"
}

//Description The console command description.
func (receiver *CreatePost) Description() string {
	return "Crate a new blog post with default front matter."
}

//Extend The console command extend.
func (receiver *CreatePost) Extend() command.Extend {
	return command.Extend{
        Category: "blog",
    }
}

//Handle Execute the console command.
func (receiver *CreatePost) Handle(ctx console.Context) error {
    filename := ctx.Argument(0)
    if filename == "" {
        return fmt.Errorf("filename is required")
    }

    title := strings.ReplaceAll(filename, "-", " ")
    title = strings.Title(strings.ToLower(title))

    currentTime := time.Now().Format("2006-01-02T15:04:05-07:00")

    content := fmt.Sprintf(`---
title: "%s"
date: %s
description: 
tags: []
categories: []
draft: false
---

Write your content here.
`, title, currentTime)

    filePath := filepath.Join("content", "posts", filename+".md")

    err := os.WriteFile(filePath, []byte(content), 0644)
    if err != nil {
        return fmt.Errorf("error creating file: %v", err)
    }

    fmt.Printf("Created new post: %s\n", filePath)
	
	return nil
}
