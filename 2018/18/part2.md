## solution to part 2

I figured the cellular automata would reach a steady state where it can just cycle between
solutions.

to test this I just created a sample of 2000 iterations. Then checked those iteration (output.txt)
for repetition. 

It seems to repeat in cycles of 28, with cycle 2000 being 19830. (Easy to eyeball, so didn't code
that part though the repetition check would be trivial)

Working out the remaining cycles (100.000.000 - 2000) leaving us with (999.998.000) we can figure
out how many cycles it'd complete before ending.  
Thus we can find out how close we'd be to completing 28 full cycles when we have reached our final
goal: `999.998.000 mod 28 = 8`.

Again, opening output.txt you can count 8 lines upwards from where you first encounter 19830.

(Or be fancy in vim and go:
 `/19830 -> 7j`)

(Keeping in mind you need to go 7 up from the start of the repetition, because you want 8 further
but counting the start as 1)
