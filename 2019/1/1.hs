import System.IO
import Data.List

f = "input.txt"

fuel :: Int -> Int
fuel x = (div x 3) - 2

ints :: [String] -> [Int]
ints = map read

solve1 :: [Int] -> Int
solve1 xs = sum $ map(\x -> fuel x) xs 

main :: IO()
main = do
	content <- readFile f
	let l = lines content
	print $ solve1 $ ints l 
