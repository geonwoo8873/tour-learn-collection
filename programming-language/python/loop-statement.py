import time

# Day, Minute, Second 단위로 각 초기값 설정
set_time_table = {  
    "day": 1, 
    "minute": 3, 
    "second": 0
    }

while True:
    set_time_table["second"] -= 1

    if set_time_table["second"] < 0:
        set_time_table["minute"] -= 1
        set_time_table["second"] = 59

    if set_time_table["minute"] < 0:
        set_time_table["day"] += 1
        set_time_table["minute"] = 2
        set_time_table["second"] = 59

    if set_time_table["day"] == 7:
        print("Countdown finished")
        break

    print(f"{set_time_table['day']:02}Days {set_time_table['minute']:02}:{set_time_table['second']:02}")
    time.sleep(1)