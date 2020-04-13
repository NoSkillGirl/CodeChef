n = gets.to_i

n.times do
    num = gets.to_i
    s = 0
    while num > 0
        s = s + num%10
        num = num/10
    end
    puts s
end
