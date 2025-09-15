# ChaCha Analysis Report

## 1. Introduction

This report analyzes a modified implementation of the ChaCha stream cipher, named "ChaC." The goal is to understand the changes made to the original design, and apply cryptanalytic methods to evaluate its security.

## 2. The Original ChaCha Cipher

ChaCha is a stream cipher designed by Daniel J. Bernstein, known for its speed and security. It operates on a 512-bit state matrix (16 x 32-bit words) composed of constants, a 256-bit key, a 64-bit nonce, and a 64-bit block counter. The cipher uses a series of rounds (typically 20), each applying a quarter-round function that mixes the state using addition, XOR, and rotation operations. The key schedule ensures that the key and nonce are properly integrated, and unique nonces are critical for security. The output keystream is XORed with plaintext to produce ciphertext.

### Diagram: ChaCha State Matrix

```
| constant | constant | constant | constant |
| key      | key      | key      | key      |
| key      | key      | key      | key      |
| counter  | nonce    | nonce    | nonce    |
```

**Quarter-Round Function:**

- The quarter-round function operates on four 32-bit words:

$$
a = a + b;\ d = d \oplus a;\ d = \text{ROTL}(d, 16)\\
c = c + d;\ b = b \oplus c;\ b = \text{ROTL}(b, 12)\\
a = a + b;\ d = d \oplus a;\ d = \text{ROTL}(d, 8)\\
c = c + d;\ b = b \oplus c;\ b = \text{ROTL}(b, 7)
$$

## 3. The Modified Version

The provided implementation (`chac.c`), described as "ChaC," introduces several changes compared to the standard ChaCha design. Key differences include:

- **Reduced State Size:** ChaC uses a 256-bit state (8 x 32-bit words) instead of ChaCha's 512-bit state (16 x 32-bit words). The state consists of a 128-bit key (4 x 32-bit), a 32-bit counter, a 64-bit nonce, and a 32-bit constant (binary for "ChaC").

- **Modified Key Schedule:** The key schedule is simplified. The block is initialized as follows:
  - block[0-3]: key
  - block[4]: counter
  - block[5-6]: nonce
  - block[7]: constant

- **Quarter-Round Function:** The QR macro is similar to ChaCha, but operates on only 8 words (instead of 16). The rotations used are 16, 12, 8, and 7 bits, matching ChaCha's quarter-round.

- **Keystream Generation:** The keystream function first runs 20 rounds (like ChaCha20), but only on the 8-word state. Then, it performs three additional sets of 2 rounds each ("cheap rotations") to generate more output blocks from the same state, aiming for efficiency.

- **Output Expansion:** Each block generates 4 x 128-bit keystreams (total 512 bits), instead of generating a new 128-bit keystream per block. This is done by reusing the state and applying extra rounds.

### Diagram: ChaC Block Structure

```
+-----+-----+-----+-----+-----+-----+-----+-----+
| k1  | k2  | k3  | k4  | ct  | n1  | n2  |  C  |
+-----+-----+-----+-----+-----+-----+-----+-----+
```

> Where k1-k4 are key words, ct is the counter, n1/n2 are nonce, and C is the constant.

**Implications:**

- Reducing the state size may decrease diffusion and security, as fewer words are mixed per round.
- Reusing the state for multiple keystream outputs could introduce correlations between outputs, potentially weakening security.
- The simplified key schedule and output expansion may make the cipher more efficient, but at the cost of cryptographic strength.

## 4. Cryptanalytic Methods Applied

### 4.1 Differential Cryptanalysis

Differential cryptanalysis examines how differences in plaintext input affect the differences in ciphertext output. For this analysis, pairs of plaintexts with controlled differences were encrypted using the modified cipher. The output differences were recorded and analyzed for patterns that could indicate weaknesses.

**Experimental Setup:**

- Generate plaintext pairs with specific bit differences (e.g., flipping one bit).

- Encrypt using the ChaC implementation, keeping key and nonce constant.

- Compare ciphertexts and analyze the distribution of output differences.

**Sample Code Snippet:**

```python
for delta in test_deltas:
    pt1 = random_plaintext()
    pt2 = xor_bytes(pt1, delta)
    ct1 = chac_encrypt(pt1, key, nonce)
    ct2 = chac_encrypt(pt2, key, nonce)
    diff = xor_bytes(ct1, ct2)
    record(diff)
```

**Results:**

Initial experiments show that the reduced state size and output expansion in ChaC may lead to less diffusion compared to ChaCha. Some output blocks exhibited higher-than-expected correlation when the same state was reused for multiple keystreams. No catastrophic bias was found, but the avalanche effect was weaker than in the original ChaCha.

### 4.2 Statistical/Randomness Testing

Randomness tests were performed to ensure the keystream output is indistinguishable from random data. Tools such as Dieharder or NIST STS were used to analyze large samples of keystream output.

**Experimental Setup:**

- Generate a large keystream using the ChaC cipher (e.g., 1 million bytes).

- Run statistical tests (Dieharder, NIST STS) on the output.

### Sample Command

```bash
**./chac -n 0x12345678 input.txt keystream.bin
**dieharder -a -g 202 -f keystream.bin
```

### Results

The keystream passed most basic randomness tests, but some tests showed slightly lower entropy and more frequent patterns than expected. This may be due to the reuse of state for multiple output blocks. While not immediately exploitable, it suggests caution for cryptographic use.

### Python Example: Entropy and Frequency Test

```python
import collections
import math

def calculate_entropy(data):
    counter = collections.Counter(data)
    total = len(data)
    entropy = -sum((count/total) * math.log2(count/total) for count in counter.values())
    return entropy

def frequency_test(data):
    ones = sum(bin(byte).count('1') for byte in data)
    zeros = len(data) * 8 - ones
    print(f"Ones: {ones}, Zeros: {zeros}")

# Example usage:
with open('keystream.bin', 'rb') as f:
    keystream = f.read()
    print(f"Entropy: {calculate_entropy(keystream):.4f} bits/byte")
    frequency_test(keystream)
```

## 5. Critical Discussion

### Strengths

- ChaC retains the basic structure and round function of ChaCha, which is well-studied and robust against many attacks.
- The cipher is efficient and produces a large keystream per block, which may be useful for certain applications.

### Weaknesses

- Reduced state size and output expansion may decrease diffusion and introduce correlations between output blocks.
- The avalanche effect is weaker, and some statistical tests show lower entropy than ChaCha.
- The simplified key schedule and reuse of state could make the cipher more vulnerable to advanced cryptanalysis.

### Literature Comparison

Academic studies of ChaCha (Bernstein, 2008; RFC 8439) show that its security relies on a large state and strong diffusion. Reducing the state or reusing it for multiple outputs is generally discouraged, as it can introduce weaknesses. Similar stream ciphers (e.g., Salsa20) also emphasize the importance of state size and round count for security.

## 6. Role of AI Tools

AI tools (e.g., Copilot, ChatGPT) were used to:

- Summarize cryptanalytic methods.
- Suggested experiment setups and analysis techniques.

### Reflection

AI accelerated the research and coding process, provided broad ideas, and helped clarify concepts. However, human expertise was essential for critical analysis and interpretation of results. AI tools were especially useful for quickly reviewing the differences between ChaCha and ChaC, generating experimental code, and summarizing academic literature.

## 7. Conclusion

The modified ChaCha cipher (ChaC) was evaluated using differential and statistical methods. While some strengths remain, the modifications introduce potential vulnerabilities, especially regarding diffusion and output correlations. AI tools were valuable for support but did not replace my own analysis. Further cryptanalysis and caution are recommended before deploying ChaC in security-critical applications.
