import System.IO
import Data.List

f = "input.txt"

fuel :: Int -> Int
fuel x = if y > 0
		then y
		else 0
	where y = (div x 3) - 2

ints :: [String] -> [Int]
ints = map read

solve1 :: [Int] -> Int
solve1 xs = sum $ map(\x -> fuel x) xs 


fuelFM :: Int -> Int
fuelFM x 
 | fuel x == 0 = 0
 | otherwise = fuel x + (fuelFM $ fuel x)

solve2 :: [Int] -> Int
solve2 xs = sum $ map(\x -> fuelFM x) xs

main :: IO()
main = do
	content <- readFile f
	let l = lines content
	print $ solve1 $ ints l 
	print $ solve2 $ ints l 
