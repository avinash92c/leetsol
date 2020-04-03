class Solution:
    def findTheDifference(self, s: str, t: str) -> str:
        # s = 'abcd'
        # t = 'abcde'
        ssum = 0
        
        for alp in s:
            ssum+=ord(alp)
            
        for alp in t:
            ssum-=ord(alp)
        
        return chr(abs(ssum))
