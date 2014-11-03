
Bounce stereo files down to mono.

I'm deaf in one ear, so before I transfer songs to my MP3 player, I
bounce them down to mono with a call to sox. That way, I don't need
any special gear; I just use ordinary hardware and the same thing
plays in both chanels. I've been doing this for some time now with a
Perl script that does all the songs in a given directory
sequentially. This version, written in Go, does them concurrently.
