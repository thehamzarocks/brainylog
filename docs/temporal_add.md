# What happens on an add

The add command is simple, but hugely powerful.

Adding a simple log of type info looks like this:

>$ brainyLog a I love bananas

This would create a line in log.bl:
>(2020-06-26T22:10:58+00:00) [1593209458000.0] vy665dg> I love bananas

The timestamp with() is the actual UTC timestamp at which this log was added.
The number within [] is the corresponding epoch with a .0 appended to it.

Now, let's say I added lots more data unrelated to this log. The data would look something like this:

>(2020-06-26T22:10:58+00:00) [1593209678000.0] vy665dg> I love bananas
>(2020-06-27T22:09:40+00:00) [1593209987000.0] bd65gh8> Summers are hot
>(2020-06-27T22:11:34+00:00) [1593209998000.0] p999iut> Winters are cold
>... and lots more weather stuff.

Now, I want to add "I love pineapples too", but this is related to the bananas log, not the weather stuff. So ideally, it should be next to it. That's where temporal locality comes in,

First, we get the hash of the bananas log:
>$ brainyLog g bananas
>(2020-06-26T22:10:58+00:00) [1593209458000.0] vy665dg> I love bananas

Now, we take this hash and specify we want to temporally place our current log after this one:
> $ brainyLog a -t vy665dg I love pineapples too
> (2020-06-27T22:14:58+00:00) [1593209458000.10] bg675a0s> Winters are cold

This changed the .0 at the end of the [] to .10. Trying to temporally add after this new position would result in .20 and so on.

But what if we tried to add after .0 again? This would create a [] with a .010.
Trying to add after .10 would result in .110.


## The algorithm:
##### Generating the next temporally-located log:
Get the epoch corrresponding to the log we want to add this log after.
Let this be abcd.efgh0
If abcd.efg(h+1) doesn't exist, add abcd.efg(h+1)0 as the temporal epoch for the new log.
Else, add abcd.efgh10 as the new temporal epoch.

(Special case where temporal epoch is abcd.0: Generate abcd.10 as the new temporal epoch)
