_ = gets
drinks = gets.split.map(&:to_f)
puts (drinks.max + drinks.sum) / 2
