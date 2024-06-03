{-# LANGUAGE RecordWildCards #-}

import System.IO


-- calculate min sq ft of paper needed. return this + area of smallest side
-- surface area: 2lw + 2wh + 2hl
-- input : LxWxH


data Box = Box {
    l :: Int,
    w :: Int,
    h :: Int
}

instance Show Box where
    show (Box l w h) = "(L: " <> show l <> " W: " <> show w <> " H: " <> show h <> ")\n"


getSides :: Box -> (Int,Int,Int)
getSides b = 
    let x = (2 * l b * w b)  
        y = (2 * w b * h b) 
        z = (2 * h b * l b)
    in (x,y,z)


surfaceArea :: Box -> Int
surfaceArea b = 
    let (x,y,z) = getSides b 
    in x + y + z

smallestSide :: Box -> Int
smallestSide b = 
    let (x,y,z) = getSides b
    in div (minimum [x,y,z]) 2

neededPaper :: Box -> Int
neededPaper b = surfaceArea b + smallestSide b


split str = case break (== 'x') str of
    (a, 'x': b) -> a : split b
    (a, "") -> [a]

--lineToBox :: String -> Box
lineToBox line = 
    let [a,b,c] = map(\c -> read c :: Int) (split line)
    in Box { l = a, w = b, h = c }


--solve1 :: [String] -> Int
solve1 lines = foldl1 (+) $ map(\x -> neededPaper (lineToBox x)) lines

main :: IO ()
main = do
    handle <- openFile "input.txt" ReadMode
    contents <- hGetContents handle
    let linesOfFile = lines contents
    let result = (solve1 linesOfFile)
    putStrLn $ show result
    hClose handle
