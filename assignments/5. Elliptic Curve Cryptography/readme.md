# Assignment 4

```sh
> src % go run main.go
==============================
Assignment 4: Elliptic Curve Cryptography
==============================

Part 1: ECC Arithmetic (mod 29)
(a) Example point α on curve: (x=0, y=6)
(b) Scalar multiples of α:
    2α = (x=28, y=11)
    3α = (x=26, y=8)
    4α = (x=26, y=21)
    5α = (x=28, y=18)
(c) 8α computed two ways:
    3α + 5α = (x=0, y=6)
    4α + 4α = (x=0, y=6)
    Both methods yield the same point.
(d) Total number of points on curve: 35

------------------------------
Part 2: Elliptic Curve Diffie-Hellman (mod 751)
(a) Alice's public key: (x=750, y=375)
(a) Bob's public key:   (x=188, y=657)
(b) Shared key computed by Alice: (x=39, y=349)
(b) Shared key computed by Bob:   (x=39, y=349)
    Shared keys match!
==============================
```