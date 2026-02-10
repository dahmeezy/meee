#creates a function
def is_palindrome(s):
    #remove all spaces if there is any
    s=s.replace(" ", "")
    #change all string char to lowercase
    s= s.lower()
    #reverse the string and compare with the original string passed
    if s[::-1]==s:
        return True
    #catches cases where user just enter empty text
    if s==" " or s=="":
        return False
    return False
print(is_palindrome("racecar"))
print(is_palindrome("A man a plan a canal Panama"))
print(is_palindrome("Hello"))


#alt
#create the function
def is_palindrome(s):
    #catches cases where user just enter empty text
    if s==" " or s=="":
        return False
    # creates a clean string
    clean_string=""
    #loops through s
    for char in s:
        #filter the character and check for only numbers and letters
        #if its num or al add to the empty string and change to lowercase
        if char.isalnum():
            clean_string+=char.lower()
    #reverse the string and check if its same as the string entered
    if clean_string[::-1]==clean_string:
        return True
    return False 
        