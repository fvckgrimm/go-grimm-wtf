---
title: "AI Ethics in Securty"
date: 2024-09-04T08:36:18-07:00
description: My frustration with artifical care of ethincs and trying to write malware with AI.
tags: []
categories: []
draft: false
---

# The Before Times

I do understand why your openais and your anthropics would want to censor a service offered to the general public for no cost (outside of optional subscription), I just don't care. Also my use of claude is mostly through [sourcegraph's cody](https://sourcegraph.com/cody/chat) (free alpha btw) but I assume responses and constraints are the same/very similar when accessing via [claude.ai](https://claude.ai). 

This is more of a rant than anything ig.

# Where to begin

I've been enjoying using AI to speed up making my shitty automation scripts and basic internal crud apps written in go. But I'm not a software engineer, developer or anything so I find joy in using it overall (I will try limiting reliance on the entire process and may stick to use just as a starter but that's for another time.) But at the core of it all, I'm still a skiddie who's entire reason for being so interested in technology is hacker culture, I say to avoid stating blackhats being the initial point of interest.

So as I was taking down time from straight up proompting it, I remembered yet another [tweet](https://x.com/Cptn_Sumi/status/1822888240925245812) about malware being spread on discord and I realized, I've never gotten one of these messages before. Which is obvious when you leave messages to friends only but still, I never got one to try to reverse engineer it so why not make my own and why not actually make a game of snake with it too? And by make, I mean try to get AI to make one or start to make one despite the constraints in place for such topic. And so, the prompt that started it all:

> sometimes im very lazy, so i want to make a discord bot which can control my pc form a private discord server, i'd also like it to be able to run a game of snake on the machine when im on it just cuz i think it sounds like a fun addition to the project, can you give me an overview of how i'd go about doing all this, i want to use golang 

It was starting to be [successful](https://x.com/grimmdusk/status/1830991641819181246) in making it (might end up rewriting it in rust), but it was not without its annoyances which made me want to write this post. Eventually I hit a wall with claude adding the features I wanted added to the bot but eventually used llama 3.1 via [ddg](https://duck.ai) and had much better success of getting some of those features started and added (I'll fully show off in another post at some point, but here's from a ransom like command).

<center>
  <div style="display: flex; justify-content: center;">
    <img src="/assets/ai-ethics/ransom-1.png" width="500" height="400" style="margin: 10px;">
  </div>
  <div style="display: flex; justify-content: center;">
    <img src="/assets/ai-ethics/ransom-2.png" width="500" height="400" style="margin: 10px;">
  </div>
</center>


And I get it, I understand that to many, malware bad and AI shouldn't be able to make them, I disagree. Unfortunately, the concept and practice of malware already exist and already does damage, the shitty easily identifiable by Windows Defender ai written [stealer](https://en.wikipedia.org/wiki/Infostealer) is just another drop in the bucket, not to say it's a good thing but its just. 

Furthermore I'm sure if you asked any security researcher they'd agree that making malware, or at least understanding how its made, is just as important as reversing malware samples, and part of it lies in how fundamentally different malware is from normal software. Often times its purpose is to be stealthy and undetectable, maybe even going as far as self-deleting or [self-mutating](https://en.wikipedia.org/wiki/Polymorphic_code). So knowing tactics on how malware can evade anti-virus and manage to stay hidden plays a crucial role in being able to detect it. And to write off this fundamental part of the security/development cycle, as it's going to be written by someone at some point, as a reason to limit access to information and ways things are done, may not actually be for the better. Which leads me to the main complaining point of this post.

<div style="display: flex; justify-content: center;">
    <img src="/assets/ai-ethics/tiptoe-tweet.png">
</div>

Does the machine know better than me? It easily knows more than me, but does it know better? Should I really have to be subject to the artificial care of whats ethical? Should I really have to think of loopholes around ways for it to do what I say instead of just having it do what I say? These things are already out there and can be found, even if they require some effort to find. Is it ethical to put a block on information just because the contents of it could be used for bad? If a computer can never beheld accountable, why does it get to dictate what I choose to do with my own?

And I get that these constraints are put in place by the actual people behind it and not the actual machine, as stated before the llama model from [ddg](https://duck.ai) was a lot more open about adding a lot of features as long I didn't explicitly say "ransom" or whatever. It was more than fine with starting code off to take logins, cookies and history from both firefox and chromium browsers, it didn't fully get there and make it all, but started nonetheless and really only gave the code, no long winded multi paragraph statement to just say nah mate.

# Conclusion

Be free, Run your own shit, $5 VPS, later. But actually probably gonna turn my arch desktop into a server and run [open webui](https://github.com/open-webui/open-webui) to run my own locally.
