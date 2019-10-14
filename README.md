
# bounce

## Bounce stereo files down to mono.

I'm deaf in one ear, so before I transfer songs to my MP3 player, I
bounce them down to mono with a call to sox. That way, I don't need
any special gear; I just use ordinary hardware and the same thing
plays in both channels. I've been doing this for some time now with a
Perl script that does all the songs in a given directory
sequentially. This version, written in Go, does them concurrently.

## Example 1

Convert all the .flac files in the current working directory to mono
.mp3 files in the `bounced` directory. This is what I do when I want
to transfer files to my MP3 player.

```
bounce
```

## Example 2

If I want to make a mono CD, I first have to convert to mono, then convert to stereo. Both channels end up the same, of course, but [this is required by the CDDA standard](https://en.wikipedia.org/wiki/Compact_Disc_Digital_Audio).

1. Convert all the .flac files in the current working directory to mono
.wav files in the `bounced` directory.
```
bounce --outext=.wav
```
2. Change to the bounced directory.
```
cd bounced
```
3. Convert all the .wav files to stereo in the `CD` directory
```
bounce --inext=.wav --outext=.wav --outdir=CD --numchan=2
```
4. Change to the `CD` directory.
```
cd CD
```
5. Write all of the .wav files to CD.
```
wodim -v -dao -audio -pad *.wav
```
