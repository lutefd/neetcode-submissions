class Solution:
    def isValid(self, s: str) -> bool:
       ## we can use a stack to track the last opened type of bracket
       ## if the closing ones matches the last opened one then it's fine
       ## if not then it's false, if the stack is empty at the end then 
       ## it's true, we can see the match via a dict

       stack = []
       source = {'}' : '{', ')':'(', ']': '['}

       for c in s:
        if c == '{':
            stack.append('{')
        elif c == '[':
           stack.append('[')
        elif c == '(':
            stack.append('(')
        else:
            ## logic here should be parse remaining chars looking at the
            ## source dict, if it matches with the last element in the
            ## stack we pop it and continue, if not return false
            if len(stack) == 0:
                return False
            if source[c] == stack[-1]:
                stack.pop()
                continue
            else:
                return False
       return len(stack) == 0
    