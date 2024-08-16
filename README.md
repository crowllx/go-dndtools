## go dnd tools

set of tools for managing dnd characters written in go.

current goals:
* store/view character data as json
* track characters stats/items at runtime without modifing json 
* add tool to help updating the json data after a session
* setup system to log events happening during encounters

some ideas of what to log
* attacks/misses
* number of enemies/types of enemeies
* damage dealt
* exploration events/occurences

unsure how these logs could be formatted for easy searching.

```
ENCOUNTER ATTACK miss werewolf
ENCOUNTER HIT miss werewolf
EXPLORATION ???

or

ENCOUNTER ATTACK self miss werewolf
ENCOUNTER ATTACK enemy hit werewolf 4
```


