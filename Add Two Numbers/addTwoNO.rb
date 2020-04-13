n = gets.to_i
while n > 0
    n = n -1
    a,b = STDIN.gets.split(' ').map(&:to_i)
    sum = a + b
    puts sum
end
