package bloom


/*
  Some bloom filter stuff is going to be used here to make some algorithms a bit
  cheaper. Should be a fast way of detecting if we're going down a path we
  already know is bad.

  One thing I'd like to play around with is using a "bloom tree" so that I can
  reconstruct previous states of the bloom filter, and similar operations.
*/


type Bloom [4]uint64

type BloomStack []Bloom
