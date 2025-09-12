# Assignment 1

Xavier Bonavita
> Due: Sep 15, 2025

## Substitution Cypher

Code for this assignment can be found [on my github.](https://github.com/xvr6/cs578/tree/main/assignments/1.%20Substitution%20Cipher/src)

### a. Provide the relative frequency of all letters A...Z in the ciphertext

1. t (217)
2. z (150)
3. q (147)
4. l (130)
5. o (124)
6. k (101)
7. f (92)
8. g (80)
9. e (77)
10. i (77)
11. s (72)
12. r (55)
13. x (55)
14. y (51)
15. h (49)
16. n (47)
17. d (44)
18. v (41)
19. w (40)
20. u (30)
21. a (20)
22. c (12)
23. m (5)
24. b (4)
25. j (1)

### b. Decrypt the ciphertext with help of the relative letter frequency of the English language

```txt
Why did Caesar love his cipher so much?
Because it made him feel like a real shift leader.
Fun Fact: Caesar’s cipher shifted each letter by 3. Simple, but it kept casual snoops away from his military
orders.

Why did the Spartan message look like nonsense?
Because it was taken out of context (literally off the stick).
Fun Fact: A strip of parchment wrapped around a rod revealed the text. Without the right rod, the letters were
scrambled.

Why did the Atbash cipher fail its test?
Because it always gave the opposite answer.
Fun Fact: Atbash, from ancient Hebrew, simply flips the alphabet (A->J, B->Y, etc.). It’s one of the oldest known
substitution ciphers.

Why don’t Vigenere users ever lose arguments?
Because they always have the last key.
Fun Fact: Using a repeating keyword, Vigenere disguised letter frequencies, baffling cryptanalysts for centuries.

Why was the Playfair cipher bad at sports?
Because it couldn’t handle singles.
Fun Fact: Playfair (1854) encrypts pairs of letters instead of singles. It was used in WWI and WWII for field
communications.

Why is ROT13 like a boomerang?
Because what you send out always comes back the same.
Fun Fact: ROT13 shifts letters by 13. Encrypting twice returns the original text--so decryption is the same as
encryption.

Why was the one-time pad always invited to parties?
Because it was perfectly secure.
Fun Fact: If truly random and never reused, the one-time pad is unbreakable. The problem: sharing and keeping the
pads secret.

Why did the Enigma machine hate small talk?
Because everything it said came out complicated.
Fun Fact: With billions of possible settings, Enigma seemed unbeatable. Allied cryptanalysis cracked it, shaping
the outcome of WWII.

Why did the Lorenj cipher break up with Enigma?
Because it wanted a bigger wheel of fortune.
Fun Fact: Lorenj SJ machines, used by Naji high command, were even more complex than Enigma. British
cryptanalysts built Colossus to crack them.

Why do cryptographers always carry snacks?
Because they believe in secure chips.
Fun Fact: Today’s ciphers rely on hard math problems (like factoring big numbers or modular arithmetic) instead
of simple letter tricks.
```

### c. Find the Plaintext/Ciphertext letter pairs, alphabetized by plaintext

1. z -> t
2. q -> a
3. o -> i
4. k -> r
5. g -> o
6. e -> c
7. i -> h
8. s -> l
9. y -> f
10. v -> w
11. a -> k
12. c -> v
13. m -> j
14. b -> x
15. x -> u
16. l -> s
17. r -> d
18. h -> p
19. n -> y
20. w -> b
21. t -> e
22. f -> n
23. d -> m
24. u -> g
25. j -> q

### d. Provide letter frequency for the given plaintext

1. e (217)
2. t (150)
3. a (147)
4. s (130)
5. i (124)
6. r (101)
7. n (92)
8. o (80)
9. c (77)
10. h (77)
11. l (72)
12. d (55)
13. u (55)
14. f (51)
15. p (49)
16. y (47)
17. m (44)
18. w (41)
19. b (40)
20. g (30)
21. k (20)
22. v (12)
23. j (5)
24. x (4)
25. q (1)

## Usage of AI in the project

- I use AI conservatively in my workflow as I find that it actively gets in my way and diminishes my learning. That being said, I do use it as a 'enhanced search engine' for if i get stumped and need help debugging.
- All that AI (in this case, github copilot via VSC plugin) was used for was a question i had about Go as i am still currently learning this language.

```txt
Question:
    Order of runes being passed into function are changing on different runs

AI Response:
  This happens because Go maps do not guarantee iteration order; each run may produce a different order for keys with equal counts. To ensure a consistent order for runes with duplicate counts, you should:

    1. Collect runes and their counts into a slice.
    2. Sort by count (descending), and for ties, sort by rune value (ascending or descending as you prefer).
```
