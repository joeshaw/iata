This is a dumb CLI tool which prints out airport info given an IATA
or ICAO code.

Don't look at the code.  It's gross.

## Install

    go get github.com/joeshaw/iata

## Usage

    $ iata LAX KIAD FRA sjc CMH BOS BUTT
    LAX (KLAX) - Los Angeles International Airport (Los Angeles, CA, US)
    IAD (KIAD) - Washington Dulles International Airport (Washington, DC, VA, US)
    FRA (EDDF) - Frankfurt am Main Airport (Frankfurt am Main, DE)
    SJC (KSJC) - Norman Y. Mineta San Jose International Airport (San Jose, CA, US)
    CMH (KCMH) - John Glenn Columbus International Airport (Columbus, OH, US)
    BOS (KBOS) - General Edward Lawrence Logan International Airport (Boston, MA, US)
    BUTT - not found

## Airport data

Airport data comes from https://ourairports.com/data/.
The airports.csv file is copied into the `tools` subdirectory.
Running `go generate .` from the toplevel directory will regenerate
the Go file containing the airport data from it.

## License

MIT
