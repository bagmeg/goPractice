## waitgroup
### this example show how to use waitgroup


when you run basic1 several times, you will notice that  \
print result are different. It's because you don't know  \
when goroutine will be scheduled.                        \
so, use waitgroup to make sure goroutine begins and wait \
for goroutine to finish