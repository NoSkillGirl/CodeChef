def fact(n)
    if(n==1)
        return 1 
    end 
    return n*fact(n-1)
end

n = gets.to_i
n.times do |i|
    a = gets.to_i    
    puts fact(a)
end