n = gets.to_i

n.times do
  origin = gets.chomp
  sum = origin.to_i + origin.reverse.to_i
  puts sum.to_s == sum.to_s.reverse ? 'YES' : 'NO'
end
