This is a dumb CLI tool which prints out airport info given an IATA
code.

Don't look at the code.  It's gross.

## Install

    go get github.com/joeshaw/iata

## Usage

    $ iata LAX FRA sjc CMH BOS BUT
    LAX - Los Angeles International Airport (Los Angeles, US)
    FRA - Frankfurt am Main International Airport (Frankfurt-am-Main, DE)
    SJC - Norman Y. Mineta San Jose International Airport (San Jose, US)
    CMH - Port Columbus International Airport (Columbus, US)
    BOS - General Edward Lawrence Logan International Airport (Boston, US)
    BUT - not found

## Airport data

Airport data comes from https://ourairports.com/data/.
The airports.csv file is copied into the `tools` subdirectory.
Running `go generate .` from the toplevel directory will regenerate
the Go file containing the airport data from it.

## License

MIT
