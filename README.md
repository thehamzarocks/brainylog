# BrainyLog

Your brain is pretty impressive! But sometimes it has trouble remembering things. **BrainyLog** is the extension to your brain, combining it's amazing cognitive abilities with a super-convenient tool that lets you access anything you wanted to remember, without having to remember the specifics of what you wanted to remember.

**BrainyLog** keeps track of any info you want to remember for posterity, be it specific research, or tasks that you don't want to leave un-completed.

And there's no need to be selfish. **BrainyLog** even allows you to share and collaborate memories - it's almost if you could share your brains (the memory part only, of course).

## Quickstart

You didn't master your mind in a few days. But you could get started the day you were born. **BrainyLog** works the same way, only faster.

Add information:
> $ brainylog -a The small bananas are my favorite
> BrainyLog: Log added at 2020-06-26T18:37:32

Retrieve information:
>$ brainylog -g favorite
> BrainyLog: 1 Result
> 45ad34v: The small bananas are my favorite (2020-06-26T18:37:32)

## A Bit More:

That was brain-dead easy! But not very helpful if you were looking for contextual information. **BrainyLog** can help you with that, and what's even more awesome - you can do it with exactly the same commands. Plus **BrainyLog** even sorts logs by relevance.

So if this is your data:
>I love fruits
>Especially the sweet ones
>The small bananas are my favorite

Then,
>$ brainylog -g small fruits
>BranyLog: 2 Results
>89hjlv: I love fruits
>32gli4: The small bananas are my favorite
>99yto: Especially the sweet ones

Wait, how did the third line get in there? That's because **BrainyLog** uses temporal locality to determine relevance - things you remembered together are probably related to each other. Unless your thoughts are completely random of course.

You can read more about **BrainyLog**'s temporal locality here.

## Quite a Bit More

Here are a few more things **BrainyLog** can do. There's more advanced stuff too, but we can come to that when you're ready. Unless of course you need it **right now**.

View logs surrounding a given log.
Add a task log.
Complete a task log (or something else on the task log).
Add a new entry today to something you discovered years ago (temporal locality at play).
Create **BrainyLog** files to organize your thoughts.
Share your logs with others.


## Scary Stuff

These are the things that would scary away the noobs. But don't worry, once you master them, **BrainyLog** will become a true memory powerhouse for your brain.

Understand **BrainyLog** hashes and manipulate them.
Master temporal locality to ensure.

## BrainyLog Best Practices

Even the best minds turn to mush if riddled with garbage. **BrainyLog** is resilient to garbage, but following these best practices will help ensure **BrainyLog** is always there for you.

1. Be a little descriptive in each log. Not so much that there's a wall of text, but at least include the key words you believe **you'd** remember.
2. Exploit temporal locality. You'll appreciate it when you see what you're looking for next to what you actually searched for.
3. Log everything that warrants remembering!
4. Use **BrainyLog** files when you want to organize your thoughts. But don't clutter them up - that's on you. Plus, files are incredibly useful for sharing.
5. Take backups. Period.
