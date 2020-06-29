# BrainyLog

Your brain is pretty impressive! But sometimes it has trouble remembering things. **BrainyLog** is the extension to your brain, combining it's amazing cognitive abilities with a super-convenient tool that lets you access anything you wanted to remember, without having to remember the specifics of what you wanted to remember.  

**BrainyLog** keeps track of any info you want to remember for posterity, be it specific research, or tasks that you don't want to leave un-completed.  

And there's no need to be selfish. **BrainyLog** even allows you to share and collaborate memories - it's almost if you could share your brains (the memory part only, of course).  

Sounds too good to be true? It is right now. Because we're in version v 0.1. But it won't be once it's given some time to mature.

## Quickstart

You didn't master your mind in a few days. But you could get started the day you were born. **BrainyLog** works the same way, only faster.  

BrainyLog commands run on the command line in your favorite Windows, Linux, or MacOS terminal:

Add information:
> $ brainylog a l The small bananas are my favorite. 

The a command tells **brainyLog** to add some info. The l is an argument that tells **brainyLog** to add whatever comes after it to the logs.

Add tasks:
>$ brainyLog a t l Bake a cake.  
This works the same as the above command, except it adds a task "Bake a cake." The task is in the created state.

Retrieve information:
>$ brainylog g l bake cake.  
> BrainyLog: 1 Result.   
> [1] (2020-06-29T00:24:09+05:30)[1593370449183.0](T-0)541763e0-e38d-4e10-9522-0e5212ced4b7>Bake a cake.  

The g command tells **brainyLog** to retrieve all lines that match some of the keywords that come after l.
The lines are retrieved along with their uuid and some other metadata.

You can use the nm argument to tell **brainyLog** not to display any metadata:
>$ brainyLog g nm l cake
>Bake a cake.

You can also retrieve tasks in any given state:
>$ brainyLog g t progress l bake cake

Also, given a line, you can get 5 lines above and below it, including the line itself:
>$ brainyLog g nm u 541763e0-e38d-4e10-9522-0e5212ced4b7 (using the line's uuid).

>$ brainyLog g n 3 (using the line's temporary positional number - be careful, we haven't specified nm here, so it would display the metadata for these lines also). 

You can also process tasks, moving them to a different state:
>$ brainyLog t complete n 5


## Local Setup

Want to fiddle around? Or give back? Get started with the source code on your own system:

##### Clone the repo
>$ git clone https://github.com/thehamzarocks/brainylog.git

##### Install go for your env (look it up)

##### Set the environment for go commands
>$ mkdir bin/

>$ go env -w GOBIN=/some_path/bin

##### Build and install
>$ go build

>$ go install

##### Create a default log.bl file
>$ touch log.bl

##### You're good to go!
>$ brainyLog v


