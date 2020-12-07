## Manipulating the Input

The original input looks like the following:

```
donpevkjhymzl
ezyopckdlnvmj

tqwfdoxvim
imzpqwruhynlca
wkemdqfvigo

kztsuvpgoheyfdrqbnxmlw
onzbdgsrqxhetkvypfulwm
eomvusqtbnwghfzdyxplkr
lthqvfxbnrpdoekugmzyws

tlsgumiveq
suevmtgrix
```

It would be easier to parse if each block (representing a group) was on its own line.

This can be done in Vim by recording a macro and then running it, same as the technique used for day 4.

1. record a macro: `qa`
1. select line: `V`
1. select to end of paragraph: `)`
1. join all the lines: `J`
1. move down 1 line: `j`
1. stop recording: `q`
1. run macro: `999@a`

Part 1 is essentially looking for the number of unique letters, so all the strings can be joined.
This can be done with the substitute command `:%s/ //g`

Part 2 needs to maintain this information so the spaces must remain, since the code for Part 1 still works, only the file with spaces is retained.
