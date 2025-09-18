# Assignment 1

Xavier Bonavita
> Due: Sep 15, 2025

## Substitution Cypher

Code for this assignment can be found [on my github.](https://github.com/xvr6/cs578/tree/main/assignments/1.%20Substitution%20Cipher/src)

### a. Provide the relative frequency of all letters A...Z in the ciphertext

01. t (217)
02. z (150)
03. q (147)
04. l (130)
05. o (124)
06. k (101)
07. f (92)
08. g (80)
09. e (77)
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

00. Cipher -> Plain
01. q -> a
02. w -> b
03. e -> c
04. r -> d
05. t -> e
06. y -> f
07. u -> g
08. i -> h
09. o -> i
10. -- no mapping for j --
11. a -> k
12. s -> l
13. d -> m
14. f -> n
15. g -> o
16. h -> p
17. j -> q
18. k -> r
19. l -> s
20. z -> t
21. x -> u
22. c -> v
23. v -> w
24. b -> x
25. n -> y
26. m -> z

### d. Provide letter frequency for the given plaintext

01. e (217)
02. t (150)
03. a (147)
04. s (130)
05. i (124)
06. r (101)
07. n (92)
08. o (80)
09. c (77)
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
23. z (5)
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

## Bonus Question: LFSR Gate Positions

Consider an LFSR with $m$ registers. If you have $2m$ consecutive output bits, you can determine the gate (feedback tap) positions as follows:

The LFSR generates bits according to a linear recurrence:
$$
s_{i+m} = c_1 s_{i+m-1} + c_2 s_{i+m-2} + \cdots + c_m s_i \pmod{2}
$$

Given $2m$ output bits $s_0, s_1, \ldots, s_{2m-1}$, you can set up $m$ equations in $m$ unknowns ($c_1, \ldots, c_m$):
$$
\begin{align*}
s_{m} &= c_1 s_{m-1} + c_2 s_{m-2} + \cdots + c_m s_0 \\
s_{m+1} &= c_1 s_{m} + c_2 s_{m-1} + \cdots + c_m s_1 \\
&\vdots \\
s_{2m-1} &= c_1 s_{2m-2} + c_2 s_{2m-3} + \cdots + c_m s_{m-1}
\end{align*}
$$

This system can be solved (mod 2) to find the feedback coefficients, i.e., the gate positions. Thus, $2m$ consecutive output bits are sufficient to uniquely determine the gate positions of an $m$-stage LFSR.
