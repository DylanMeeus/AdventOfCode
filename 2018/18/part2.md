## solution to part 2

I figured the cellular automata would reach a steady state where it can just cycle between
solutions.

to test this I just created a sample of 2000 iterations. Then checked those iterations (output.txt)
for repetition. 

It seems to repeat in cycles of 28. (Easy to eyeball, examine that iter:2000 gives us 19830 and look upward to where you find it again)

Working out the remaining cycles gives us: `100.000.000 - 2000 = 999.998.000`.

Thus we can find out how close we'd be to completing 28 full cycles when we have reached our final
goal: `999.998.000 mod 28 = 8`.

Again, opening output.txt you can count 8 lines down from where you first encounter 19830.

(Or be fancy in vim and go:
 `/19830 -> 7j`)

(Keeping in mind you need to go 7 down from the start of the repetition, because you want mod 8
but counting the start as 1).
