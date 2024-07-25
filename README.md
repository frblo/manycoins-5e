# manycoins-5e
Simple yet complicated tool for converting currencies in Dungeons &amp; Dragons 5e

Name based on [Manycoins services](https://forgottenrealms.fandom.com/wiki/Manycoins_services) in the Forgotten Realms. This is a service for exchanging coins in Dungeons & Dragons 5e. This program will convert a purse, a collection of coins, which you put into it to a *rebased* purse of the same value, but in a more sorted manner.

## User instruction

It will try to distribute the coins in descending order of value, meaning you will get as many platinum pieces (pp) as possible. You can also choose to only include certain types of coins by using the `-i` flag, followed by the types of pieces. This program also allows for addition and subtraction of coins by placing a `-` before any number of coins.

### Example 1

```sh
./manycoins 302gp 12sp 3cp
```

```
rebasing purse 302gp 12sp 3cp...
30pp 3gp 2sp 3cp
303gp 2sp 3cp
606ep 2sp 3cp
3032sp 3cp
30323cp
```

### Example 2

```sh
./manycoins -i gp ep 302gp 12sp 3cp
```

```
rebasing purse 302gp 12sp 3cp...
only using gp ep cp
303gp 23cp
606ep 23cp
30323cp
```
