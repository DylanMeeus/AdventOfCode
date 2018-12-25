input = File.read("input.txt")

chars = input.split("").take_while{|c| c != "\n"}

# solve 1
puts chars.reduce(0){|s,c| s + (c == "(" ? 1 : -1)}

# solve 2
sum = 0
seen = 0
chars.each { |c|
    sum += c == "(" ? 1 : -1
    seen += 1
    break if sum < 0
}
puts seen 
