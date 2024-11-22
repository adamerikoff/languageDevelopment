module Main where

import Control.Monad
import System.Environment
import Text.ParserCombinators.Parsec hiding (spaces)

-- Entry point of the program
main :: IO ()
main = getArgs >>= putStrLn . show . eval . readExpr . (!! 0)

-- Parser for symbols
symbol :: Parser Char
symbol = oneOf "!$%&|*+-/:<=?>@^_~#"

-- Function to parse an input string
readExpr :: String -> LispVal
readExpr input = case parse parseExpr "lisp" input of
    Left err  -> String $ "No match: " ++ show err  -- Handle parsing errors
    Right val -> val  -- Successfully parsed a Lisp expression

-- Parser for spaces
spaces :: Parser ()
spaces = skipMany1 space

-- Data type to represent Lisp values
data LispVal
    = Atom String
    | List [LispVal]
    | DottedList [LispVal] LispVal
    | Number Integer
    | String String
    | Bool Bool

-- Parser for strings
parseString :: Parser LispVal
parseString = do
    _ <- char '"'  -- Start of the string
    x <- many (noneOf "\"")  -- Capture all characters until the closing quote
    _ <- char '"'  -- End of the string
    return $ String x

-- Parser for atoms
parseAtom :: Parser LispVal
parseAtom = do
    first <- letter <|> symbol  -- First character must be a letter or symbol
    rest <- many (letter <|> digit <|> symbol)  -- Remaining characters
    let atom = [first] ++ rest  -- Combine the first and rest into a single string
    return $ case atom of
        "#t" -> Bool True  -- Special case for true
        "#f" -> Bool False  -- Special case for false
        otherwise -> Atom atom  -- All other cases

-- Parser for numbers
parseNumber :: Parser LispVal
parseNumber = liftM (Number . read) $ many1 digit

-- Parser for a list
parseList :: Parser LispVal
parseList = liftM List $ sepBy parseExpr spaces

-- Parser for dotted lists
parseDottedList :: Parser LispVal
parseDottedList = do
    head <- endBy parseExpr spaces
    tail <- char '.' >> spaces >> parseExpr
    return $ DottedList head tail

-- Parser for quoted expressions
parseQuoted :: Parser LispVal
parseQuoted = do
    _ <- char '\''
    x <- parseExpr
    return $ List [Atom "quote", x]

-- Parser for Lisp expressions
parseExpr :: Parser LispVal
parseExpr = parseAtom
    <|> parseString
    <|> parseNumber
    <|> parseQuoted
    <|> do
        _ <- char '('
        x <- (try parseList) <|> parseDottedList
        _ <- char ')'
        return x

showVal :: LispVal -> String
showVal (String contents) =  "\"" ++ contents ++ "\""
showVal (Atom name) = name
showVal (Number contents) = show contents
showVal (Bool True) = "#t"
showVal (Bool False) = "#f"
showVal (List contents) = "(" ++ unwordsList contents ++ ")"
showVal (DottedList head tail) = "(" ++ unwordsList head ++ " . " ++ showVal tail ++ ")"

unwordsList :: [LispVal] -> String
unwordsList = unwords . map showVal

instance Show LispVal where show = showVal

eval :: LispVal -> LispVal
eval val@(String _) = val
eval val@(Number _) = val
eval val@(Bool _) = val
eval (List [Atom "quote", val]) = val
eval (List (Atom func : args)) = apply func $ map eval args

apply :: String -> [LispVal] -> LispVal
apply func args = maybe (Bool False) ($ args) $ lookup func primitives

primitives :: [(String, [LispVal] -> LispVal)]
primitives = [
    ("+", numericBinop (+)),
    ("-", numericBinop (-)),
    ("*", numericBinop (*)),
    ("/", numericBinop div),
    ("mod", numericBinop mod),
    ("quotient", numericBinop quot),
    ("remainder", numericBinop rem)]

numericBinop :: (Integer -> Integer -> Integer) -> [LispVal] -> LispVal
numericBinop op params = Number $ foldl1 op $ map unpackNum params

unpackNum :: LispVal -> Integer
unpackNum (Number n) = n
unpackNum (String n) = let parsed = reads n in
    if null parsed
        then 0
        else fst $ parsed !! 0
unpackNum (List [n]) = unpackNum n
unpackNum _ = 0