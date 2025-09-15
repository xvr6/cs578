import os
import subprocess
import numpy as np

def xor_bytes(a, b):
  return bytes(x ^ y for x, y in zip(a, b))

key = b'\x00' * 16  # Example key
nonce = '0x12345678'
num_tests = 100
bit_positions = [0, 1, 2, 3, 4, 5, 6, 7]
diff_counts = []

for bit in bit_positions:
  diffs = []
  for _ in range(num_tests):
    pt1 = os.urandom(32)
    pt2 = bytearray(pt1)
    pt2[0] ^= (1 << bit)  # Flip one bit in first byte
    with open('pt1.bin', 'wb') as f: f.write(pt1)
    with open('pt2.bin', 'wb') as f: f.write(pt2)
    subprocess.run(['./chacbin', '-n', nonce, 'pt1.bin', 'ct1.bin'])
    subprocess.run(['./chacbin', '-n', nonce, 'pt2.bin', 'ct2.bin'])
    with open('ct1.bin', 'rb') as f: ct1 = f.read()
    with open('ct2.bin', 'rb') as f: ct2 = f.read()
    diff = sum(b1 != b2 for b1, b2 in zip(ct1, ct2))
    diffs.append(diff)
  diff_counts.append((bit, np.mean(diffs), np.std(diffs)))

for bit, mean, std in diff_counts:
  print(f'Bit {bit}: Mean diff = {mean:.2f}, Std = {std:.2f}')