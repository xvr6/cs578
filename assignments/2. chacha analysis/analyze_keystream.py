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

with open('keystream.bin', 'rb') as f:
    keystream = f.read()
    print(f"Entropy: {calculate_entropy(keystream):.4f} bits/byte")
    frequency_test(keystream)