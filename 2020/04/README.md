## Manipulating the Input

The original input looks like the following:

```
eyr:2027
hcl:#602927
hgt:186cm byr:1939 iyr:2019 pid:552194973 ecl:hzl

pid:657988073 eyr:2020 byr:1996
ecl:brn
hcl:#866857 iyr:2015
hgt:164cm

hcl:#fffffd
byr:1951 cid:321 iyr:2017 eyr:2022 ecl:brn hgt:62in pid:#6ef4e1

eyr:2025 iyr:2011
byr:1980
hcl:#fffffd cid:129 pid:420023864
hgt:150cm
ecl:brn
```

It would be easier to parse if each block (representing a passport) was on its own line.

This can be done in Vim by recording a macro and then running it.

1. record a macro: `qa`
1. select line: `V`
1. select to end of paragraph: `)`
1. join all the lines: `J`
1. move down 1 line: `j`
1. stop recording: `q`
1. run macro: `999@a`
