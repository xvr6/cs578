# ElGamals

1. ElGamal Encryption - Encrypt the following messages using ElGamal encryption (Z∗409 and g = 19):

> Perform both encryption and decryption. Show every intermediate step of the ElGamal Scheme.

   1. Private key x = 301, ephemeral key k = 23, message m = 59.
      1. `go run .\main.go 301 23 59`

            ```plaintext
             Test values: q=409, g=19, x=301, k=23, m=59
             h (g^x mod q): 398
             g^k used : 351
             g^ak used : 258
             Encrypted message: 15222
             Ephemeral public key p: 351
             Decrypted message: 59
            ```

   2. Private key x = 301, ephemeral key k = 135, message m = 59.
      1.  `go run .\main.go 301 135 59`

            ```plaintext
             Test values: q=409, g=19, x=301, k=135, m=59
             h (g^x mod q): 398
             g^k used : 366
             g^ak used : 223
             Encrypted message: 13157
             Ephemeral public key p: 366
             Decrypted message: 59
            ```

