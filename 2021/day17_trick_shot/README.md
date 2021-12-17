# 2021 Day 17 Trick Shot

## Input

target area: x=155..182, y=-117..-67

## Part 1

Only the y-axis needs to be taken into consideration.
What goes up must come down, so the probe will eventually return to yPos 0 at with negative y launch velocity.
On the next step, it cannot be lower than the minimum point of the target area.

So its y-velocity must be -117 after coming back to yPos 0, which means it must be launched at y-velocity 116
so the 2nd last step before coming to yPos 0 is 116.

The maximum height it reaches is the sum of the Arithmetic Progression
`116 + 115 + 114 + ... + 1`, which is `0.5 * 116 * 117 = 6786`.
