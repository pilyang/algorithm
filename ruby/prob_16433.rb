# frozen_string_literal: true

n, i_r, i_c = gets.split.map(&:to_i)
i_remain = (i_r + i_c) % 2

(1..n).each do |r|
  (1..n).each do |c|
    print ((r + c) % 2 == i_remain) ? "v" : "."
  end
  puts
end
