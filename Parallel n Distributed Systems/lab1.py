import random
import time
import multiprocessing


n = 500
proc_n = 16


class Matrix:
    def __init__(self, n, m):
        self.n = n
        self.m = m
        self.matrix = [[random.randint(-5, 5) for i in range(m)] for j in range(n)]

        self.MAXPRINT = 7

    def print(self):

        for i in range(min(self.MAXPRINT, self.n)):
            for j in range(min(self.MAXPRINT, self.m)):
                print(self.matrix[i][j], end=' ')
            if self.m > self.MAXPRINT:
                print('...', end='')
            print()
        if self.n > self.MAXPRINT:
            print('...')

    def is_equal(self, mat):
        if self.n != mat.n or self.m != mat.m:
            return False

        for i in range(self.n):
            for j in range(self.m):
                if self.matrix[i][j] != mat.matrix[i][j]:
                    return False
        return True


def multiplicate_rows(a, b):
    res = Matrix(a.n, b.m)
    for i in range(a.n):
        for j in range(b.m):
            res.matrix[i][j] = 0
            for k in range(a.m):
                res.matrix[i][j] += a.matrix[i][k] * b.matrix[k][j]
    return res


def multiplicate_columns(a, b):
    res = Matrix(a.n, b.m)
    for i in range(a.n):
        for j in range(b.m):
            res.matrix[i][j] = 0
            for k in range(a.m):
                res.matrix[i][j] += a.matrix[i][k] * b.matrix[k][j]
    return res


def multi_run_wrapper(args):
    return multi_thread(*args)


def multi_thread(x, y):
    res = Matrix(len(x), len(y[0]))
    for i in range(len(x)):
        for j in range(len(y[0])):
            res.matrix[i][j] = 0
            for k in range(len(x[0])):
                res.matrix[i][j] += x[i][k] * y[k][j]
    return res


def multiplicate_threads(a, b):
    if a.n != b.n:
        raise Exception('Dimensions must be equal!')

    res = Matrix(n, n)

    x = a.matrix
    y = b.matrix

    processes = []
    manager = multiprocessing.Manager()
    return_matrix = manager.list()
    return_matrix.append(res.matrix)

    args = []
    for i in range(proc_n):
        if i != proc_n-1:
            args.append((x[int(i*n/proc_n):int((i+1)*n/proc_n)], y))
        else:
            args.append((x[int((proc_n-1) * n / proc_n):], y))

    with multiprocessing.Pool(proc_n) as p:
        res_mat = p.map(multi_run_wrapper, args)
    for i in range(proc_n):
        if i != proc_n-1:
            res.matrix[int(i*n/proc_n):int((i+1)*n/proc_n)] = res_mat[i].matrix
        else:
            res.matrix[int((proc_n-1) * n / proc_n):] = res_mat[i].matrix
    return res


if __name__ == '__main__':
    mat = Matrix(n, n)

    mat1 = Matrix(n, n)
    mat2 = Matrix(n, n)

    start = time.time()
    mat1 = multiplicate_rows(mat, mat)
    print('By rows:', time.time() - start)

    start = time.time()
    mat2 = multiplicate_columns(mat, mat)
    print('By columns:', time.time() - start)

    start = time.time()
    mat3 = multiplicate_threads(mat, mat)
    print('Multi Rows:', time.time()-start)
    print('Correct result:', mat3.is_equal(mat1))
