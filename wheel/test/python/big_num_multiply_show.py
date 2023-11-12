import matplotlib.pyplot as plt

# 假设你的数据如下：
size = [1, 2, 3, 4, 5]  # 这应该是你的size数组
naive = [1.0, 1.2, 1.4, 1.6, 1.8]  # 这应该是你的naive数组
karatsuba = [0.8, 1.0, 1.2, 1.4, 1.6]  # 这应该是你的karatsuba数组
fft = [0.6, 0.8, 1.0, 1.2, 1.4]  # 这应该是你的fft数组

# 绘制折线图
plt.plot(size, naive, label='Naive')
plt.plot(size, karatsuba, label='Karatsuba')
plt.plot(size, fft, label='FFT')

# 设置图像标题和坐标轴标签
plt.title('Time vs Size')
plt.xlabel('Size')
plt.ylabel('Time (ms)')

# 显示图例
plt.legend()

# 显示图像
plt.show()
