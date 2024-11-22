module Main where

import System.Environment
import Text.ParserCombinators.Parsec hiding (spaces)

-- Entry point of the program
main :: IO ()
main = do
    args <- getArgs
    -- Check if there are arguments, otherwise print a default message
    if null args
        then putStrLn "No argument provided"
        else putStrLn (readExpr (args !! 0))

-- Define the parser for a symbol
symbol :: Parser Char
symbol = oneOf "!$%&|*+-/:<=?>@^_~#"

-- Function to parse an input string
readExpr :: String -> String
readExpr input = case parse symbol "lisp" input of
    Left err  -> "No match: " ++ show err  -- Handle parsing errors
    Right val -> "Found value: " ++ [val]  -- Successfully parsed a symbol
