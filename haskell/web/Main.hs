{-# LANGUAGE OverloadedStrings #-}

import Web.Scotty

main :: IO ()
main = scotty 3000 $ do
  get "/" $ do
    html $ mconcat ["<h1>Hello Haskell, Hello Scotty</h1>"]
