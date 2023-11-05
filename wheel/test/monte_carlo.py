import random
import math 

def f(x):
    return (1.0 / math.sqrt(2 * math.pi)) * math.exp(-x**2 / 2.0)

def monte_carlo_integration(num_samples=10000000):
    count = 0
    for _ in range(num_samples):
        x = random.uniform(0, 1)
        y = random.uniform(0, 1)
        if y <= f(x):
            count += 1
    return count / num_samples

print(monte_carlo_integration())
