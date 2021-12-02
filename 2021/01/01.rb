#!/usr/bin/env ruby

def countIncrease(depths)
  numIncrease = 0
  currDepth = depths[0]

  depths.drop(1).each do |depth|
    if depth > currDepth
      numIncrease += 1
    end
    currDepth = depth
  end

  return numIncrease
end

file = File.open("01.txt")
depths = file.readlines.map(&:chomp).map(&:to_i)
puts countIncrease(depths)

windowDepths = []
(1..depths.size-2).each do |index|
  windowDepths.push(depths[index-1] + depths[index] + depths[index+1])
end

puts countIncrease(windowDepths)
