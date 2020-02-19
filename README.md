# Function optimization based on Nelder-Mead's downhill simplex (1965)

This package, `nmoptim` implements a variant of the Nelder-Mead downhill simplex optimization heuristic. The idea is that we have a function we want to optimize and a simplex in a search space. The algorithm iteratively changes the simplex through contractions, expansions and reflections until the simplex is extremely small and contains the optimal minimum point in the function.

It doesn't work for functions with several optimal points and there are other problems with the approach (infinite cycling). But in many cases, the method works.

# License

	Copyright 2014-2020 Jesper Louis Andersen

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	    http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
