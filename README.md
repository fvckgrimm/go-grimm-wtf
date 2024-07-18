# Grimm Site

<img src="/repo-assets/showcase.gif"/>

This is a project using [goravel](https://github.com/goravel/goravel) and [claude](https://claude.ai/) (actually surprised claude knew what goravel was), as I was just working on a project using [Laravel](https://laravel.com/) and have been wanting to rework my personal site using go. Blog posts are written in markdown and can be stored in [/content/posts/](./content/posts/) folder for local file rendering, or in [/storage/app/blog-posts/](./storage/app/blog-posts/) for use with a database if you so please.

If you want to use a database, you'll need to edit the [blog_posts.go](./app/models/blog_post.go) model, and just simply rename the structs. (At least I'm pretty sure, once I started working on local file rendering I kinda stopped caring about having the option for both to work based on what you want, was tired of trying to deal with errors(wtf are test?)).

## Blog Post Generating

```bash
go run . artisan create:post example-title
```

This command creates a post/markdown file with the given name, in the content/posts/ directory with some default Front Matter. The title is taken from the filename and the date is the time created.

## Blog Post Heading / Front Matter

```yaml
---
title: "Example Post"
date: 2024-07-16T15:08:20-07:00
description: Example markdown post for local file rendering.
tags: [tag1, tag2]
categories: [ category1, category2]
draft: false
---

```

Setting draft to `true` hides the post from being shown/accessed in all the following: 

- The blog index for site.
- Viewing by going to `/blog/{slug}`.
- The rss feed.
- The `/api/recent-posts` response.

Obviously, if the git repo is public, what does it really matter if it's shown on the site or not, but who cares, chances are if it's in the [content/posts](./content/posts/) or [storage/app/blog-posts](./storage/app/blog-posts/) folder, it'd end up on the site anyways. If I didn't want it on there I'd use my obsidian vault to note it down or something.

Both tags and categories currently aren't used in anything, might make a view to show posts containing the selected tag/category at some point. Might also move some things to a layout so only one instance has to be edited for the header/footer for example. 

*But I also mostly made this for myself so I probably won't care to since shit's already set, but maybe I'll force myself to on some refactoring shit, idk.*

# Terminal Commands

Terminal commands are all handled by the [script.js](./resources/javascript/script.js) file. Adding new commands just requires adding the name to the list of availableCommands, adding a new case and handling the logic for that command in there.


## Compiling 

Different ways to compile the project.

```bash 
# Select the system to compile
go run . artisan build

# Specify the system to compile
go run . artisan build --os=linux
go run . artisan build -o=linux

# Static compilation
go run . artisan build --static
go run . artisan build -s

// Specify the output file name
go run . artisan build --name=goravel
go run . artisan build -n=goravel 

# Regular Compilation 
go build .
```

The Following files and folders need to be uploaded to the server during deployment:

```bash
./main // Binary file
.env
./database
./public
./storage
./resources 
./content 
```


## About Goravel

Goravel is a web application framework with complete functions and good scalability. As a starting scaffolding to help
Gopher quickly build their own applications.

The framework style is consistent with [Laravel](https://github.com/laravel/laravel), let Php developer don't need to learn a
new framework, but also happy to play around Golang! In tribute to Laravel!

Welcome to star, PR and issues！

## Getting started

```
// Generate APP_KEY
go run . artisan key:generate

// Route
facades.Route().Get("/", userController.Show)

// ORM
facades.Orm().Query().With("Author").First(&user)

// Task Scheduling
facades.Schedule().Command("send:emails name").EveryMinute()

// Log
facades.Log().Debug(message)

// Cache
value := facades.Cache().Get("goravel", "default")

// Queues
err := facades.Queue().Job(&jobs.Test{}, []queue.Arg{}).Dispatch()
```

## Documentation

Online documentation [https://www.goravel.dev](https://www.goravel.dev)

Example [https://github.com/goravel/example](https://github.com/goravel/example)

> To optimize the documentation, please submit a PR to the documentation
> repository [https://github.com/goravel/docs](https://github.com/goravel/docs)

## License

The Goravel framework is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
