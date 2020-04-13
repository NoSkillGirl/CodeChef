inpput = STDIN.gets
n, k = inpput.split(' ')
n = n.to_i
k = l.to_i
# .map(&:to_i)
count = 0

n.times do |i|
	x = STDIN.gets.to_i
	count = count +1 if x % k == 0
end

puts count