keywords = %w[pi ka chu]

input = gets.chomp

index = 0
while index < input.length
  found = false
  keywords.each do |keyword|
    next unless input[index..].start_with?(keyword)

    index += keyword.length
    found = true
    break
  end
  break unless found
end

if index == input.length
  puts 'YES'
else
  puts 'NO'
end
