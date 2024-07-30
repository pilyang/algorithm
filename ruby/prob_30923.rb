# frozen_string_literal: true

num = gets.chomp.to_i
heights = gets.chomp.split.map(&:to_i)

min_height_sum = heights.each_cons(2).sum { |a, b| [a, b].min }

surface_area = 2 * num + 4 * heights.sum - 2 * min_height_sum
puts surface_area
