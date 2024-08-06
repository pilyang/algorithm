alphabet_count = Hash.new(0)

while (line = gets)
  line.chomp.tr(' ', '').each_char do |c|
    alphabet_count[c] += 1
  end
end

max_count = alphabet_count.values.max
alphabet_count.sort.each do |k, v|
  print k if v == max_count
end
