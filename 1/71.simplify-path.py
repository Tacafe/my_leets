class Solution:
    def simplifyPath(self, path: str) -> str:
        path = path.replace('//', '/')
        dirs = [p for p in path.split('/') if p != '']

        absolute_pathes = []
        for d in dirs:
            if d == '..':
                if len(absolute_pathes) > 0:
                    absolute_pathes.pop()
            elif d == '.':
                continue
            else:
                absolute_pathes.append(d)

        return f"/{'/'.join(absolute_pathes)}"

if __name__ == '__main__':
    path = '/home/'
    print(Solution().simplifyPath(path))
