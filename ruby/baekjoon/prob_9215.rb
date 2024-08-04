def gcd(a, b)
  a, b = b, a if a < b
  a, b = b, a % b until b.zero?
  a
end

tc = 1
while (n = gets.chomp.to_i)
  break if n.zero?

  s_d = 0
  s_c = 0
  s_p = 1
  n.times do
    input = gets.chomp
    is_contain_slash = input.include?("/")
    is_contain_comma = input.include?(",")

    if !is_contain_slash && !is_contain_comma
      d = input.to_i
      c = 0
      p = 1
    elsif is_contain_slash && !is_contain_comma
      d = 0
      c, p = input.split("/").map(&:to_i)
    else
      d = input.split(",").first.to_i
      c, p = input.split(",").last.split("/").map(&:to_i)
    end

    s_d += d
    s_c = s_c * p + c * s_p
    s_p *= p
    g = gcd(s_c, s_p)
    s_c /= g
    s_p /= g
  end

  s_d += s_c / s_p
  s_c %= s_p

  print "Test #{tc}: "
  if s_d.zero? && s_c.zero?
    puts 0
  elsif s_c.zero?
    puts s_d
  else
    puts "#{s_d.to_s + "," unless s_d.zero?}#{s_c}#{"/" + s_p.to_s unless s_p == 1}"
  end

  tc += 1

end
