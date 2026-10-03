# Frist create array `a` with elements 1 to 5 [0, 1, 2, 3, 4]
a = [1, 2, 3, 4, 5]
# `b`is a array index copy
b = a.copy()
# `b`is copyed array index add value `6`
b.append(6)
# `c`of `b` equal array index element
c = b

d = c.copy()
d.remove(6)

print(a)
print(b)
print(c)
print(d)