n = gets.chomp.to_i

words = []
n.times { words << gets.chomp }

words.sort!.reverse!

pre_x_words = []
pre_x_words << words.shift

words.each do |word|
  pre_x_words << word unless pre_x_words.last.start_with? word
end

puts pre_x_words.length
