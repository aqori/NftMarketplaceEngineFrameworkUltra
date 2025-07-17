// cmd/nftmarketplaceengineframeworkultra/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplaceengineframeworkultra/internal/nftmarketplaceengineframeworkultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplaceengineframeworkultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
