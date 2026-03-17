# #creates a function
# def is_palindrome(s):
#     #remove all spaces if there is any
#     s=s.replace(" ", "")
#     #change all string char to lowercase
#     s= s.lower()
#     #reverse the string and compare with the original string passed
#     if s[::-1]==s:
#         return True
#     #catches cases where user just enter empty text
#     if s==" " or s=="":
#         return False
#     return False
# print(is_palindrome("racecar"))
# print(is_palindrome("A man a plan a canal Panama"))
# print(is_palindrome("Hello"))


# #alt
# #create the function
# def is_palindrome(s):
#     #catches cases where user just enter empty text
#     if s==" " or s=="":
#         return False
#     # creates a clean string
#     clean_string=""
#     #loops through s
#     for char in s:
#         #filter the character and check for only numbers and letters
#         #if its num or al add to the empty string and change to lowercase
#         if char.isalnum():
#             clean_string+=char.lower()
#     #reverse the string and check if its same as the string entered
#     if clean_string[::-1]==clean_string:
#         return True
#     return False 
        

# #creates a function
# def is_palindrome(s):
#    # reverse the string and compare
#    if s[::-1]==s:
#        return True
#    return False
# print(is_palindrome("racecar"))
# print(is_palindrome("A man a plan a canal Panama"))
# print(is_palindrome("Hello"))
# # the above works but doesn't ignore characters

#creates function
def is_palindrome(s):
    if s==" " or s=="":
        return False
    #declare a new string variables to store the string
    new_word=""
    #loop through the string passed by the user
    for char in s:
        #check that its an alphabet or number
        if char.isalnum():
            #if it is al/num...store in the "new_word" string and make them all lowercase
            new_word+=char.lower()
    # reverse the new stored string and compare      
    if new_word[::-1]==new_word:
        return True
    return False
print(is_palindrome("racecar"))
print(is_palindrome("A man a plan a canal Panama"))
print(is_palindrome("Hello"))

# it works but doesn't return position where it stops being a palindrome

#AI

# #creates function
# def is_palindrome(s):
#     #declare a new string variables to store the string
#     new_word=""
#     #loop through the string passed by the user
#     for char in s:
#         #check that its an alphabet or number
#         if char.isalnum():
#             #if it is al/num...store in the "new_word" string and make them all lowercase
#             new_word+=char.lower()
#     #declare the first and last index of the string
#     begin_index, end_index = 0,len(new_word)-1
#     #a while loop that iterates until the index of the first string is less than that of the last string
#     while begin_index < end_index:
#         #checks if they are the same foward or backwards
#         if new_word[begin_index]!= new_word[end_index]:
#             #Return the position where the string stops being a palindrome (if not one)
#             return False, begin_index, end_index
#         #adds to the begin_index each time
#         begin_index += 1
#         #adds to the end_index each time
#         end_index -= 1
#     return True
# print(is_palindrome("racecar"))
# print(is_palindrome("A man a plan a canal Panama"))
# print(is_palindrome("Hello"))
