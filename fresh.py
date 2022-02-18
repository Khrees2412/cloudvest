def get_result():
    result = int(input("Enter the result: "))

    if result == 999:
        print(result)
        return;

    if result >= 70:
        print("A")
    elif result >= 60:
        print("B")
    elif result >= 50:
        print("C")
    elif result >= 40:
        print("D")
    elif result < 40:
        print("F")
  
    get_result()
get_result()


def get_avg():
    first = input("Enter the first set of three numbers e.g 1,2,3 : ")
    second = input("Enter the second set of three numbers e.g 1,2,3 : ")
    third = input("Enter the third set of three numbers e.g 1,2,3 : ")

    a = first.split(",")
    first_total = 0
    for num in a:
        first_total += int(num)
        
    a_avg = first_total / 3

    b = second.split(",")
    second_total = 0
    for num in b:
        second_total += int(num)
    
    b_avg = second_total / 3

    c = third.split(",")
    third_total = 0
    for num in c:
       third_total += int(num)
        
    c_avg = third_total / 3

    all = a_avg + b_avg + c_avg 
    avg = all / 3
    print(avg)

# get_avg()