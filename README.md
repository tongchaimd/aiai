# aiai
## Goal
To prove that neural networks can be used to navigate and
utilize a gravity simulation.
## Core
Create a simulation that exhibits these properties:
- Continuity of spacetime
- Containing rigid body with mass
- Obeying Newtonian mechanics
- Containing one controllable body
- Time navigable
Create an agent that uses machine learning.
Make the agent learns to navigate through the simulation.
Output the progress of the agent.
## META
The real goal of this project is, actually, to experiment
__software engineering with extreme DDD(Domain Driven Design)
and extreme separation of concerns__. I hypothesize that
the end product will be the most maintainable software, ever.
Each file contains exactly one concern. Therefore, if any
part needs to be changed, the logic should be in only one
file (and, maybe, some lower level files).
### Separation of Abstractions
Starting from the purpose of the project which is thought of,
by me, as the highest level of abstraction. The project is, then,
constructed level of abstraction by level of abstraction down
from there. Code will not be written, before the module in the
higher level uses it. Each level contains details of higher
level and specifies the general function of the lower level.
This way, the dependencies only point up (no module depends on
modules lower than itself).
### Separation of Purposes
This is just normal separation of concerns, I guess.
