A small superoptimizer I'm writing in golang, and planning to use Z3 for verification.

The goal is to try to trim down the search space by:
* Doing dependency analysis on a skeleton of a code fragment to detect if the dataflow is even reasonable for an optimization
* Detecting simple cases with false dependencies/constant results (e.g, xor x x, sub x x, eq x x, neq x x, etc.)
* Avoiding double-checking commutative and associative operations
* Simple dead code analysis (no need to have a random operation whose result is never used, so why even try it?)
* Common subexpression eliminiation (an optimization with redundant code is obviously not the most efficient code)

I also would like to have a more sophisticated cost model to at least approximate factors such as instruction latency, instruction-level parallelism, and decode bottlenecks.

I'm not sure how I'll handle vector operations and constant values, though I would like to support them eventually.


Running the search on the GPU (OpenCL) is also a long-term goal.
