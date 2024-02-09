def canConstruct(ransomNote: str, magazine: str) -> bool:
    if len(magazine) < len(ransomNote):
        return False

    r = list(ransomNote)
    m = list(magazine)
    for c in ransomNote:
        if c in m:
            m.remove(c)
            r.remove(c)

    return len(r) == 0

if __name__ == "__main__":
    ransomNote = "fffbfg"
    magazine = "effjfggbffjdgbjjhhdegh"

    print(canConstruct(ransomNote, magazine))
