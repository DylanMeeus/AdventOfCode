
input = File.read("input.txt")

chars = input.split("").take_while{|c| c != "\n"}
puts chars
puts chars.map{|c| c == "(" ? 1 : -1}.sum()
