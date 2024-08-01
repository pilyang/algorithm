# frozen_string_literal: true

_n = gets.chomp.to_i
open1, open2 = gets.split.map(&:to_i)
m = gets.chomp.to_i
sequence = m.times.map { gets.chomp.to_i }

@dp = {}

def min_move(open1, open2, target, sequence)
  return 0 if target == sequence.size

  open1, open2 = open2, open1 if open1 > open2

  key = [open1, open2, target].hash
  return @dp[key] if @dp.key?(key)

  move1 = (open1 - sequence[target]).abs + min_move(sequence[target], open2, target + 1, sequence)
  move2 = (open2 - sequence[target]).abs + min_move(open1, sequence[target], target + 1, sequence)
  @dp[key] = [move1, move2].min
end

puts min_move(open1, open2, 0, sequence)
