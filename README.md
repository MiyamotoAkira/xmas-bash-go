# Christmas Bash

So it is Christmas time and Santa Claws has found that there is an issue: both
the Grinch and Jack Skellington are planning "stuff" and he cannot allow that.

# Initial set of Rules

This is the easy part :-) Is just Rock Paper Scissors, and the triangle is:
- `Jack Skellington` beats `Santa Claws`, because he is resorceful and crafty
- `The Grinch` beats `Jack Skellington`, because he is evil and devious
- `Santa Claws` beats `The Grinch`, because Good Triumphs over Evil (well, we wish)

So, you have to create a game, best of three hands, and there could be ties. 
Two players can play. let's say CLI just because is easier and quicker.

# First set of modifications

We are changing the rules!!.

First, each player selects what they want to play (Grinch, Jack or Santa). But
doesn't get revealed immediately. Each player reveals one of the two they have
not selected randomly (ie: I have chosen Jack, but I keep that hidden, and I
show one of the two other cards, which this time happens to be Santa).

Second, each player selects and plays one of three possible modifiers:
- `Barrel, Lock and Shock` - They are mischievous, so they inverse the resolution triangle
- `Elves in the Shelves` - They work extremely hard, and they convert a tied game into a win
- `Mama Who` - She has no time for all this non-sense, the hand end in a tie.

This events are revealed at the same time as the selected character. But they resolve
in the same order as above. That it is, first `Barrel, Lock and Shock` do their thing, then
`Elves in the Shelves` do they thing, finally `Mama Who`.

# Second set of modifications

The changes are as thus:
- Each player has a deck of 9 cards: four `Barrel, Lock and Shock`, three `Elves in the Shelves`,
and two `Mama Who`
- Each hand a player will draw (randomly) three cards.
- After the hand is finished, whatever cards are not played will be discarded
without being revealed.
- Each player in turn can play one of their cards. A player can choose not to play
a card
- This will continue until either both players pass or they have used all three cards.
- You resolve them in order, the first pair of cards played, then the second, then the third.
- For each pair you resolve same as on the first set of modifications.


And may the best win.
