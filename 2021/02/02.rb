#!/usr/bin/env ruby

file = File.open("02.txt")
commands = file.readlines.map(&:chomp)

horizPos = 0
depth = 0

commands.each do |command|
  subCommands = command.split
  case subCommands[0]
  when 'forward'
    horizPos += subCommands[1].to_i
  when 'down'
    depth += subCommands[1].to_i
  when 'up'
    depth -= subCommands[1].to_i
  end
end

puts horizPos * depth

# Part 2

aim = 0
horizPos = 0
depth = 0

commands.each do |command|
  subCommands = command.split
  case subCommands[0]
  when 'forward'
    commandAmount = subCommands[1].to_i
    horizPos += commandAmount
    depth += aim * commandAmount
  when 'down'
    aim += subCommands[1].to_i
  when 'up'
    aim -= subCommands[1].to_i
  end
end

puts horizPos * depth
