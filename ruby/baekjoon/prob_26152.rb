_ = gets.chomp.to_i
up = gets.chomp.split(' ').map(&:to_i)
down = gets.chomp.split(' ').map(&:to_i)

offsets = []
last_min_value = Float::INFINITY

up.each_with_index do |p, i|
  diff = p - down[i]
  if diff < last_min_value
    offsets << [diff, i + 1]
    last_min_value = diff
  end
end

_ = gets.chomp.to_i
birds = gets.chomp.split(' ').map(&:to_i)

birds.each do |b|
  offset = offsets.bsearch { |o| o[0] < b }
  puts offset.nil? ? up.length : offset[1] - 1
end
