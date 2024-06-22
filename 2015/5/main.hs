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


-- problem 2

countOccurrences :: (Ord a) => [a] -> [(a, Int)]
countOccurrences xs = map(\l -> (head l, length l)) . group . sort $ xs

indexes :: Int -> String -> String -> [Int]
indexes acc stack needle 
    | stack == "" = []
    | isPrefixOf needle stack = [acc] <> indexes (succ acc) (tail stack) needle
    | otherwise = indexes (succ acc) (tail stack) needle


hasNonOverlappingIndex :: Int -> [Int] -> Bool
hasNonOverlappingIndex i lst = 1 <= length (filter(\x -> x /= i && x /= (succ i) && x /= (pred i)) lst)

findNonOverlappingIndexes :: [Int] -> [Int]
findNonOverlappingIndexes list = filter(\l -> hasNonOverlappingIndex l list) list

findDuplicatePairs :: String -> [String]
findDuplicatePairs "" = []
findDuplicatePairs input = 
    let zipped = tail $ init $ zip (" " <> input) (input <> " ")
        instances = countOccurrences zipped
    in map(\i -> [fst (fst i)] ++ [snd (fst i)]) $ filter(\i -> (snd i) > 1) instances

hasValidPairs :: String -> Bool
hasValidPairs s = 
    let pairs = map (\pair -> indexes 0 s pair) (findDuplicatePairs s)
        idx = map (\pair -> findNonOverlappingIndexes pair) pairs
        valid = length $ filter (\x -> x /= []) idx
    in valid >= 1


validRepetition :: String -> Bool
validRepetition s = 
    let idx = [x | x <- [0..(pred $ length s)]]
    in 1 <= length (filter(\y -> s !! y == s !! (y+2)) (init $ init idx))
        

solve2 :: [String] -> Int
solve2 input = length $ filter(\s -> (hasValidPairs s) && validRepetition s) input

main :: IO () 
main = do
    handle <- openFile "input.txt" ReadMode
    contents <- hGetContents handle
    let input = lines contents
    let result1 = solve1 input
    let result2 = solve2 input
    putStrLn $ show result1
    putStrLn $ show result2
    hClose handle
