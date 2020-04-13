amt, accBal = STDIN.gets.split(' ').map(&:to_f)
puts (amt % 5 == 0 and amt+0.5 < accBal) ? accBal-amt-0.5 : accBal