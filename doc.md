# BrainyLog Doc:

**BrainyLog** lets you log stuff so you can access it quickly when you can't remember it later. It lets you worry less about the interface, and more about what you're actually logging and retrieving.

### The log.bl file:  
Before you can run any commands though, create an empty log.bl file in your current directory. **BrainyLog** will warn if this file does not exist in the current directory. This file is used to store all your logs, so be careful with it!  

### Command structure:
All **BrainyLog** commands have the following structure:  
brainyLog < command > < flags > < key - value pairs >  

The command is always required. The flags and key-value pairs may be optional and vary depending on the command.  

command - one of v, a, g, t. The main command used to determine the action to perform.  

flags - command-dependent. Used to enable/disable functionality depending on the command. Always optional.  

key - command-dependent. Used to provide parameters to a command. Followed by a single value (single-valued key) or multiple space - separated values (mutlivalued key) depending on the key itself. May be optional. Multivalued keys always come in the end.  

value - Either a single value, or multiple space - separated values. Always follow a key. Mandatory if a key is entered.  


### Version:
The v command is used to display info about the current version of **BrainyLog**.  
It does use any flags or keys.  

Eg:  
>$ brainyLog v  


### Add logs:
The a command is used to add logs.  
Flags:  
t - adds a task log instead of a normal info log. The task would be in the created state by default.  

Single-Valued Key-Value pairs:  
None  

Multi-Valued Key-Value pairs:  
l - Used to specify that anything that comes after this will be added directly as new line content to the log.  

Eg  
>$ brainyLog a t l This is a task.  

### Get logs:
The g command is used to retrieve logs. **BrainyLog** tries to retrieve as many matching logs as possible, while ordering the most relevant logs first.
Flags:  
nm - Tells brainyLog not to show positional numbers next to the displayed line. No new positional mappings will be created if this flag is specified.  

Single-Value Key-Value pairs:  
t - Followed by a task state (create/progress/suspend/cancel/complete/allTasks). Filters retrieved logs to match the specified task state. Opional.  

n - If this key is specified, brainyLog uses the value as a positional number to match a specific line. Only this line would be displayed if the optional m key is not specified, else all lines in the range specified by m would be displayed. If the nm flag is used, no positional numbers would be displayed, nor would the existing positional numbers be overwritten. Otherwise they would be displayed and overwrite the existing mappings. Optional.    

m - used along with the n key to tell **brainyLog** to display m lines above and below the current line including the line matched by n. The lines would be displayed in order. Refer n for more details. Optional.  


Multi-Valued Key-Value pairs:  
l - Used to specify that anything that comes after this will be used as search text to retrieve logs. Optional.  

Note: If conflicting keys (eg n and l) are specified in the same g command, the behaviour is undefined. 

Eg:  
>$ brainyLog g l search for this  
>$ brainyLog g n 3 m 4 nm


### Process a task:
The t command is used to change the state of a task in the line specified by the positional number.  
Flags:  
None.  

Single-Value Key-Value pairs:  
t - used to specify the task state to change to (create/progress/suspend/cancel/complete). Required.  

n - the positional number of the task line to change the state of. Retrieved using the g command. Required.  

Multi-Valued Key-Value pairs:  
None.  

Eg:  
>$ brainyLog t t progress n 4


### Delete a log line:  
THe d command is used to delete a log. Useful for when a log was inadvertently entered, or to correct typos.  
Flags: None  

Single-Value Key-Value pairs:  
n - the positional number of the task line to change the state of. Retrieved using the g command. Required.  

Multi-Valued Key-Value pairs:  
None.  

Eg:  
>$ brainyLog d n 8  






