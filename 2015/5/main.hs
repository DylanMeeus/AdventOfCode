import System.IO
import Data.List

vowels :: [Char]
vowels = ['a', 'e', 'i', 'o', 'u']

blocklist :: [String]
blocklist = ["ab", "cd", "pq", "xy"]


-- functions for 'nice string check'

isBlocked :: String -> Bool
isBlocked s = (length $ filter(\x -> x `isInfixOf` s) blocklist) > 0

threeVowels :: String -> Bool
threeVowels s = (length $ filter(\x -> elem x vowels) s) >= 3

doubleLetter :: String -> Bool
doubleLetter "" = False 
doubleLetter (a:b:"") = a == b
doubleLetter (a:b:c:rem) = a == b || b == c || doubleLetter (c : rem)
doubleLetter _ = undefined

isNice :: String -> Bool
isNice s = (threeVowels s && doubleLetter s) && (not $ isBlocked s)

solve1 :: [String] -> Int
solve1 input = length $ filter(\s -> isNice s) input

main :: IO () 
main = do
    handle <- openFile "input.txt" ReadMode
    contents <- hGetContents handle
    let input = lines contents
    let result1 = solve1 input
    putStrLn $ show result1
    hClose handle
