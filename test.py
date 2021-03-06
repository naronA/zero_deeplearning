import numpy as np

'''
dot(x[0][0], y)
[[ 7 10]
 [15 22]]
'''

'''
dot(x[0], y)
[[[ 7 10]
  [15 22]]

 [[19 28]
  [17 26]]]
'''

'''
dot(x, y)
[[[[ 7 10]
   [15 22]]

  [[19 28]
   [17 26]]]

 [[[23 34]
   [35 52]]

  [[13 20]
   [ 5  8]]]]
'''
x = np.array([
    [
        [[1, 2], [3, 4]],
        [[4, 5], [5, 4]],
    ],
    [
        [[5, 6], [8, 9]],
        [[4, 3], [2, 1]],
    ]
])

y = np.array([[1, 2], [3, 4]])

print(np.dot(x, y))
print(np.dot(x[0], y))
print(np.dot(x[0][0], y))
'''
[[[[[[  7  10]
     [ 14  13]]

    [[ 21  24]
     [  8   5]]]


   [[[ 15  22]
     [ 32  31]]

    [[ 47  54]
     [ 20  13]]]]



  [[[[ 19  28]
     [ 41  40]]

    [[ 60  69]
     [ 26  17]]]


   [[[ 17  26]
     [ 40  41]]

    [[ 57  66]
     [ 28  19]]]]]




 [[[[[ 23  34]
     [ 50  49]]

    [[ 73  84]
     [ 32  21]]]


   [[[ 35  52]
     [ 77  76]]

    [[112 129]
     [ 50  33]]]]



  [[[[ 13  20]
     [ 31  32]]

    [[ 44  51]
     [ 22  15]]]


   [[[  5   8]
     [ 13  14]]

    [[ 18  21]
     [ 10   7]]]]]]
'''
