# Function optimization based on Nelder-Mead's downhill simplex (1965)

This package, `nmoptim` implements a variant of the Nelder-Mead downhill simplex optimization heuristic. The idea is that we have a function we want to optimize and a simplex in a search space. The algorithm iteratively changes the simplex through contractions, expansions and reflections until the simplex is extremely small and contains the optimal minimum point in the function.

It doesn't work for functions with several optimal points and there are other problems with the approach (infinite cycling). But in many cases, the method works.
