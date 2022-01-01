# 2021 Day 24 Arithmetic Logic Unit

Brute forcing the problem will take too long as the possible space is 9^14.

https://github.com/mrphlip/aoc/blob/master/2021/24.md

The puzzle input is divided into 14 modules, 1 for each input digit.
7 of the modules push onto the stack `z`. These are the ones that do 
`div z 1` (essentially a no-op), and add a positive number just before
the `eql` checks. This ensures that the condition cannot be met.

The other 7 modules pop a value from the stack, but the condition can be
met. If it is not met, some value with an addition is pushed onto the
stack, so it will not be 0 at the end of the program.

So none of the conditional stack pushes can be allowed to happen.

The following pseudocode is for my puzzle input.

```
  z.push(A+5)
  z.push(B+14)
  z.push(C+15)
  z.push(D+16)

  // none of the conditional stack pushes must be allowed to happen
  if E != z.pop() - 16 ? z.push(E+8)
  if F != z.pop() - 11 ? z.push(F+9)
  if G != z.pop() - 6 ? z.push(G+2)
  z.push(H+13)
  z.push(I+16)
  if J != z.pop() - 10 ? z.push(J+6)
  if K != z.pop() - 8 ? z.push(K+6)
  if L != z.pop() - 11 ? z.push(L+9)
  z.push(M+11)
  if N!= z.pop() - 15 ? z.push(N+5)
```

This produces the following digit pair equalities.

```
  E == D
  F == C + 4
  G == B + 8
  J == I + 6
  K == H + 5
  L == A - 6
  N == M - 4
```

  // largest, 1 of the digit pair is 9
  ABCDEFGHIJKLMN
  91599994399395

  // smallest, 1 of the digit pair is 1
  ABCDEFGHIJKLMN
  71111591176151