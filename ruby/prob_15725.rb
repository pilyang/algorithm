# frozen_string_literal: true

input = gets.chomp

if input.include?("x")
  case a = input.split("x").first
  when "", nil
    puts "1"
  when "-"
    puts "-1"
  else
    puts a
  end
else
  puts "0"
end
