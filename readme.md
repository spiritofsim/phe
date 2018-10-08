# Poker Hand Evaluator

This is a Texas Holdem poker hand evaluation go library.

Can evaluate 5,6 and 7 card combinations.

It uses TwoPlusTwo algorithm with pre-build ranks table (see ranks.dat.gz).

On my machine (i7-7700HQ 2.80 GHz) it can evaluate 53 millions 7-cards hands per second.
I think it fastest go implementation for now.

## Installation

    go get github.com/spiritofsim/phe


## Usage

Eval func returns rank and hand type. Bigger rank - better hand, so you can just compare two ranks to know which hand is better.

See test fo examples.