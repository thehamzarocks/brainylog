# File Structure

Although you can create **BrainyLog** files to organize your thoughts, **BrainyLog** creates a special file for you by default. Any commands without a specified filename operate on this special "default" file.

The file is called brainylog.bl and here's what the file looks like:

#### Header
Every brainylog file (.bl extension) has the following headers:
name: The filename (brainylog.bl) in this case
createdBy: The user who created this file
createdTime: The creation timestamp


#### Body
The body of a brainylog file consists of several lines. Each line has the following:
timeCreated: Time when the line was first created
timeUpdated: Time when the line was last edited
hash: A unique hash identifying the line
type: type of the line (information/task)
content: the actual content to be logged
taskStatus: The status if it's a task line. Information lines contain null for task
deleted: flag identifying if the line has been deleted
position: temporal position of the line in the file. If the line was added normally, it would take the latest position. If it was a temporal update, it would come in the position it was supposed to be added temporally.


