class Trie:

    def __init__(self):
        self.root = {}

    def insert(self, word: str) -> None:
        node = self.root

        for i, c in enumerate(word):
            if c in node:
                node = node[c]
            else:
                node[c] = {}
                node = node[c]

            if i == len(word) - 1:
                node["end"] = True

    def search(self, word: str) -> bool:
        node = self.root
        for c in word:
            if not c in node:
                return False
            else:
                node = node[c]
        return "end" in node

    def startsWith(self, prefix: str) -> bool:
        node = self.root
        for c in prefix:
            if not c in node:
                return False
            else:
                node = node[c]
        return True



if __name__ == '__main__':
    t = Trie()
    methods = ["Trie","insert","search","search","startsWith","insert","search"]
    params = [[],["apple"],["apple"],["app"],["app"],["app"],["app"]]

    obj = None
    for m, param in zip(methods, params):
        if m == "Trie":
            obj = Trie()
        elif m == "insert":
            print(obj.insert(param[0]))
        elif m == "search":
            print(obj.search(param[0]))
        elif m == "startsWith":
            print(obj.startsWith(param[0]))

