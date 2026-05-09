This is a dumb CLI tool which prints out airport, city, and station
info given an IATA or ICAO code.

Don't look at the code.  It's gross.

## Install

    go get github.com/joeshaw/iata

## Usage

    $ iata LAX KIAD FRA sjc CMH BOS NYC QQP BUTT
    LAX (KLAX) - Los Angeles International Airport (Los Angeles, CA, US)
    IAD (KIAD) - Washington Dulles International Airport (Dulles, VA, US)
    FRA (EDDF) - Frankfurt Main Airport (Frankfurt am Main, DE)
    SJC (KSJC) - Norman Y. Mineta San Jose International Airport (San Jose, CA, US)
    CMH (KCMH) - John Glenn Columbus International Airport (Columbus, OH, US)
    BOS (KBOS) - Boston Logan International Airport (Boston, MA, US)
    NYC - New York City (New York City, NY, US)
    QQP - London Paddington station (London, GB)
    BUTT - not found

## Airport data

Airport data comes from https://github.com/datasets/airport-codes.
The airports.csv file is copied into the `tools` subdirectory.

IATA codes for everything that isn't an airport — cities (e.g. NYC,
LON, PAR), rail and bus stations (e.g. QQP for London Paddington),
ferry ports, and off-line points (cities without their own airport,
like JRS for Jerusalem) — come from OpenTravelData's curated POR
dataset at https://github.com/opentraveldata/opentraveldata.  The
optd_por_public.csv file is copied into the `tools` subdirectory.
On collisions, the OurAirports entry wins.

Running `go generate .` from the toplevel directory will regenerate
the Go file containing the data from these sources.

## License

MIT
