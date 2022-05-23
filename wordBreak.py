class Solution(object):
    def wordBreak(self, s, wordDict):
        """
        :type s: str
        :type wordDict: List[str]
        :rtype: bool
        """
        words_set = set(wordDict)
        
        def check_cut_possible(word):
            #從左邊開始蒐
            for i in range(1, min(len(word), 20)):
                word1 = word[: -i]
                if word1 in words_set:
                    word2 = word[-i: ]
                    if word2 in words_set:
                        words_set.add(word)         
        #從最小開始搜
        for i in range(1, len(s)+1):
            check_cut_possible(s[:i])
            
        return s in words_set
