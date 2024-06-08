import System.IO
import Data.List

-- turn input into "move instructions"
-- run through each move instruction, and record the new location (list)
-- count duplicates in the list (tuples)


-- v => south
-- ^ => north
-- < => east
-- > => west


pos:: (Int, Int)
pos = (0,0)

updatePos :: (Int, Int) -> Char -> (Int, Int)
updatePos (x, y) move = 
    case move of
    '^' -> (succ x, y)
    '>' -> (x, succ y)
    'v' -> (pred x, y)
    '<' -> (x, pred y)
    _ -> (x, y)



walk :: String -> (Int, Int) -> [(Int, Int)]
walk [] _ =  []
walk (x:xs) current = 
    let newPos = updatePos current x
    in newPos : (walk xs newPos)

solve1 :: String -> Int
solve1 input  = (length . nub) $ (pos : (walk input pos))


main :: IO ()
main = do
    handle <- openFile "input.txt" ReadMode
    contents <- hGetContents handle
    let result1 = solve1 contents
    putStrLn $ show result1
