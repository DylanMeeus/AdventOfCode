import System.IO
import System.Environment
import Data.List
-- Get the valid passphrases (no repeated words)


file1 = "passphrase.txt"
solve :: [String] -> Int
solve input = length $ filter(\x -> length (words x) == length(nub (words(x)))) input

solve2 :: [String] -> [String]
solve2 input = filter(\y -> hasAnagrams(words y)) input

-- get a [String]
hasAnagrams :: [String] -> Bool
hasAnagrams xs = length (filter(\y -> y == 1) [length (getAnagrams x xs) | x <- xs]) == length xs


-- count anagrams of X in XS
getAnagrams :: String -> [String] -> [String]
getAnagrams x xs = filter(\y -> isAnagram x y) xs

isAnagram :: String -> String -> Bool
isAnagram a b = (sort a) == (sort b) 

main :: IO()
main = do
       content <- readFile file1
       let fileLines = lines content
       --mapM_ putStrLn (solve2 fileLines)
       putStrLn (show $ length (solve2 fileLines))
      
