class Value:
    def __init__(self, data: float, children: tuple["Value"]=(), op: str=""):
        self.data = data
        self.children = set(children)
        self.op = op

    def __repr__(self):
        return f"({self.data})"

    def __add__(self, other: "Value"):
        return Value(self.data + other.data, (self, other), "+")

    def __mul__(self, other: "Value"):
        return Value(self.data * other.data, (self, other), "*")

a = Value(2.0)
b = Value(-3.0)
c = Value(10.0)

LEFT OFF: https://youtu.be/VMj-3S1tku0?feature=shared&t=1521
