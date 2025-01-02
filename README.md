# 0xBlade

### Terminal based melee game with almost no dependencies

## Useful Links

ANSI Escape Codes : [here](https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797)

## Weird Issues I've run into

1. For some reason in gameManager.drawScreen() it was throwing a nil pointer dereference. But changing it from a `for i=0;i<len;i++` loop to a `for _, idx := range()` and dereferencing the val pointer worked.
2. Need to make udev rule to access inputs
