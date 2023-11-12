import matplotlib.pyplot as plt
import numpy as np

# 生成随机数
np.random.seed(0)
N = 10000
X = np.random.uniform(0, 1, (N, 12)).sum(axis=1) - 6

# 绘制直方图
plt.hist(X, bins=50, density=True, alpha=0.6, color='g')

# 设置图像标题和坐标轴标签
plt.title('Distribution of Sum of 12 U(0,1) Minus 6')
plt.xlabel('Value')
plt.ylabel('Frequency')

# 显示图像
plt.show()
