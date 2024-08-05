input = gets.chomp
stack = []

input.each_char do |c|
  case c
  when 'H'
    stack.push(1)
  when 'C'
    stack.push(12)
  when 'O'
    stack.push(16)
  when '('
    stack.push(-1)
  when ')'
    sum = 0
    until (top = stack.pop) == -1
      sum += top
    end
    stack.push(sum)
  else
    stack.push(stack.pop * c.to_i)
  end
end

puts stack.sum
