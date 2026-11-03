class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        freq1 = {}
        freq2 = {}
        if len(s) != len(t):
            return False
        for ch in s:
            if ch in freq1:
                freq1[ch] += 1
                continue
            freq1[ch] = 1
        for ch in t:
            if ch in freq2:
                freq2[ch] += 1
                continue
            freq2[ch] = 1
        return freq1 == freq2