# 2020 Day 13 Shuttle Search

Part 1 was done with Excel.

The bus numbers are 23, 41, 449, 13, 19, 29, 991, 37, 17.

Put the starting timestamp in cell A2, then put the bus numbers in the first row from B1 to J1.

Set B2 to `$A2/B$1` and drag across the row to copy the formula across.
Create a few rows where the timestamp auto increments.
Select row 2 and drag down across the range to autocomplete, then find the first cell with an integer.

It's quite easy to spot since it's shorter than the other cells.
The answer for this input is bus 991 at time 108838, giving a final answer of `991*(108838/108832) = 5946`.

