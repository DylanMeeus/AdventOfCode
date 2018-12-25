input = File.read("input.txt")

chars = input.split("").take_while{|c| c != "\n"}
puts chars.reduce(0){|s,c| s + (c == "(" ? 1 : -1)}
